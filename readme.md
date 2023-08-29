# MoodTracker - Your Personal Mood Tracker in Terminal

MoodTracker is a command-line tool to help you track your moods, because mental health matters, even for developers! With a few keystrokes, log your mood, add additional notes, and generate reports to understand yourself better. All of this, right in your terminal. It's built using Go, and incorporates SQLite for local storage.

---

## Table of Contents

- [Features](#features)
- [Usage](#usage)

---

## Features

- 💡 **Simple:** Easy-to-understand commands.
- 📄 **Note-taking:** Allows adding additional notes for context.
- 📊 **Reporting:** Generate mood reports for different time frames.
- 🎨 **Stylish:** Makes use of colors to distinguish moods visually.

## Usage

```bash
# Log a positive mood
./moodtracker p --note "Had a great day!"
./moodtracker p -m "Had a great day!"

# Log a negative mood
./moodtracker n --note "Felt a bit off."
./moodtracker n -m "Felt a bit off."
```

Generate a report:

```bash
# Generate a report for the last day
./moodtracker report day

# Generate a report for the last week
./moodtracker report week

# Generate a report for the last month
./moodtracker report month
```

List all moods:

```bash
# List all recorded moods
./moodtracker list
```
