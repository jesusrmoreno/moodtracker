# MoodCLI - Your Personal Mood Tracker in Terminal

MoodCLI is a command-line tool to help you track your moods, because mental health matters, even for developers! With a few keystrokes, log your mood, add additional notes, and generate reports to understand yourself better. All of this, right in your terminal. It's built using Go, and incorporates SQLite for local storage.

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
./moodcli p --note "Had a great day!"
./moodcli p -m "Had a great day!"

# Log a negative mood
./moodcli n --note "Felt a bit off."
./moodcli n -m "Felt a bit off."
```

Generate a report:

```bash
# Generate a report for the last day
./moodcli report day

# Generate a report for the last week
./moodcli report week

# Generate a report for the last month
./moodcli report month
```

List all moods:

```bash
# List all recorded moods
./moodcli list
```
