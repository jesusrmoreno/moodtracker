package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Mood struct {
	ID        uint `gorm:"primaryKey"`
	Feeling   string
	Note      string
	CreatedAt time.Time
}

func UserDir() string {
	// get the user dir programatically
	dir := os.Getenv("HOME")
	if dir == "" {
		dir = os.Getenv("USERPROFILE")
	}
	return dir
}

var dbPath = UserDir() + "/moods.db"

func main() {
	db, _ := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	db.AutoMigrate(&Mood{})

	var note string
	var rootCmd = &cobra.Command{Use: "moodcli"}

	var cmdPositive = &cobra.Command{
		Use:   "p",
		Short: "Log a positive mood",
		Run: func(cmd *cobra.Command, args []string) {
			logMood(db, "positive", note)
		},
	}
	cmdPositive.Flags().StringVarP(&note, "note", "m", "", "Add a note")

	var cmdNegative = &cobra.Command{
		Use:   "n",
		Short: "Log a negative mood",
		Run: func(cmd *cobra.Command, args []string) {
			logMood(db, "negative", note)
		},
	}
	cmdNegative.Flags().StringVarP(&note, "note", "m", "", "Add a note")

	var cmdReport = &cobra.Command{
		Use:   "report [day|week|month]",
		Short: "Generate a mood report",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Specify a time frame: day, week, or month.")
				return
			}
			generateReport(db, args[0])
		},
	}

	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "List all moods",
		Run: func(cmd *cobra.Command, args []string) {
			var moods []Mood
			db.Order("created_at desc").Find(&moods)

			for _, mood := range moods {
				// Log the mood, including the note, created_at, and feeling
				// log the date in the format YYYY-MM-DD HH:MM:SS
				fmt.Print(lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Render(mood.CreatedAt.Format("2006-01-02 15:04")))
				fmt.Print(" ")
				if mood.Feeling == "positive" {
					fmt.Print(lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Render(mood.Feeling))
				} else {
					fmt.Print(lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render(mood.Feeling))
				}
				fmt.Print(" ")
				fmt.Println(mood.Note)

			}
		},
	}

	rootCmd.AddCommand(cmdPositive, cmdNegative, cmdReport, cmdList)
	rootCmd.Execute()
}

func logMood(db *gorm.DB, feeling string, note string) {
	mood := Mood{Feeling: feeling, Note: note, CreatedAt: time.Now()}
	db.Create(&mood)
}

func generateReport(db *gorm.DB, timeFrame string) {
	var moods []Mood
	var fromTime time.Time

	switch timeFrame {
	case "day":
		fromTime = time.Now().AddDate(0, 0, -1)
	case "week":
		fromTime = time.Now().AddDate(0, 0, -7)
	case "month":
		fromTime = time.Now().AddDate(0, -1, 0)
	default:
		fmt.Println("Invalid time frame. Use day, week, or month.")
		return
	}

	db.Where("created_at > ?", fromTime).Find(&moods)

	var positiveCount, negativeCount int
	for _, mood := range moods {
		if mood.Feeling == "positive" {
			positiveCount++
		} else {
			negativeCount++
		}
	}
	fmt.Printf("Mood report for the last %s:\n", timeFrame)
	fmt.Printf("Positive moods: %d\n", positiveCount)
	fmt.Printf("Negative moods: %d\n", negativeCount)
}
