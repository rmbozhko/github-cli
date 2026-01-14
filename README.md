# GitHub Command-Line Companion

A lightweight, terminal-first GitHub CLI that allows you to **manage GitHub issues** and **inspect recent user activity** directly from the command line.  
This project is designed as a practical learning exercise for working with REST APIs, JSON, authentication, and CLI application design — without relying on external HTTP libraries.

---

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [GitHub User Activity](#github-user-activity)
  - [Issue Management](#issue-management)
- [Authentication](#authentication)
- [Editor Integration](#editor-integration)
- [Error Handling](#error-handling)
- [API Usage](#api-usage)
- [Constraints](#constraints)
- [Future Enhancements](#future-enhancements)
- [Project Goals](#project-goals)

---

## Overview

**GitHub Command-Line Companion** is a CLI tool that enables developers to interact with GitHub without leaving the terminal.

It supports two primary workflows:

1. **Viewing recent GitHub user activity**
2. **Creating, reading, updating, and closing GitHub issues**

The tool emphasizes:
- Clear terminal output
- Predictable commands
- Safe editor-based text input
- Graceful error handling
- Direct use of the GitHub REST API

---

## Features

### ✅ User Activity Viewer
- Fetch recent public activity of any GitHub user
- Human-readable terminal output
- No authentication required (optional for higher rate limits)

### ✅ Issue Management (Authenticated)
- Create new issues
- View existing issues
- Update issue title and body
- Close issues
- Navigate:
  - Bug reports (issues with `bug` label)
  - Milestones
  - Users (authors, assignees)

### ✅ Editor-Based Input
- Automatically opens the user’s preferred text editor for long text input
- No awkward multiline terminal prompts

### ✅ Robust Error Handling
- Invalid usernames
- API failures
- Authentication errors
- Rate limit handling
