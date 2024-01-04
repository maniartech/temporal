# temporal (WIP)

A golang library for parsing and parsing, formatting and processing dates and times using simple human-friendly formats such as `yesterday`, `tomorrow`, `dd/mm/yyyy`, etc. The temporal does not aims to be a replacement for the standard time package, but rather addtional facilities to make regular date and time operations such as formatting, parsing, relative time, data range, etc. easier.

It does not depend on any third-party libraries and is fully compatible with Tinygo. It uses caching to improve performance and reduce allocations.

💯 **100% test coverage** 💯

✔️ Tinygo Compatible

## Installation

Installation is simple. Just run the following command in your terminal to install the temporal package in your project.

```sh
go get github.com/maniartech/temporal
```

## Usage

The following example shows how to use the temporal package.

```go

import "github.com/maniartech/temporal"

// Parse a date
t, err := temporal.Parse("tomorrow")

// Format a date
s := temporal.Format(t, "yyyy-mm-dd")

// Convert date string to different format
s, err := temporal.Convert("2012-01-01", "yyyy-mm-dd", "dd/mm/yyyy")

// Get the relative date range
start, end, err := temporal.RelativeRange("last-week")

temporal.Today()
temporal.EoD()
temporal.Yesterday()
temporal.Tomorrow()
temporal.LastWeek()
temporal.LastMonth()
temporal.LastYear()
temporal.NextWeek()
temporal.NextMonth()
temporal.NextYear()
```

## Introducing IDF (Intuitive Date Format)

We've developed the Intuitive Date Format (IDF) for Temporal. IDF is a cAsE-insensitive format, eliminating ambiguity often associated with dd-mm-yyyy formats. This intuitive format makes date and time formatting simple and hackable. entry by removing the need to remember upper and lower case attributes, a common issue with other similar formats. For example, %Y represents a four-digit year, while %y denotes a two-digit year in strftime. In contrast, IDF intuitively uses yyyy for a four-digit year and yy for a two-digit year. Typing dates is also more straightforward with IDF, as the format yyyy-mm-dd is easier to remember and input compared to the less intuitive 2006-01-02.

Temporal supports simple, human-friendly date-time formatting. The table below displays the supported formats. Internally, Temporal utilizes time.Time.Format() and converts human-friendly formats into the time.Time format. For instance, it transforms yyyy-mm-dd into 2006-01-02 before using time.Time.Format() to format the date.

### Date Formats

| Format | Output   | Description                                   |
| ------ | -------- | --------------------------------------------- |
| `yy`   | `06`     | Two-digit year with leading zero              |
| `yyyy` | `2006`   | Four-digit year                               |
| `m`    | `1`      | Month without leading zero                    |
| `mm`   | `01`     | Month in two digits with leading zero         |
| `mt`   | `1st`    | Month in ordinal format (not for parsing)     |
| `mmm`  | `Jan`    | Month in short name                           |
| `mmmm` | `January`| Month in full name                            |
| `d`    | `2`      | Day without leading zero                      |
| `dd`   | `02`     | Day in two digits with leading zero           |
| `dt`   | `2nd`    | Day in ordinal format (not for parsing)       |
| `ddd`  | `002`    | Zero padded day of year                       |
| `www`  | `Mon`    | Three-letter weekday name                     |
| `wwww` | `Monday` | Full weekday name                             |

### Time Formats

| Format | Output | Description                                      |
| ------ | ------ | ------------------------------------------------ |
| `h`    | `3`    | Hour in 12-hour format without leading zero      |
| `hh`   | `03`   | Hour in 12-hour format with leading zero         |
| `hhh`  | `15`   | Hour in 24-hour format without leading zero      |
| `a`    | `pm`   | AM/PM in lowercase                               |
| `aa`   | `PM`   | AM/PM in uppercase                               |
| `ii`   | `04`   | Minute with leading zero                         |
| `i`    | `4`    | Minute without leading zero                      |
| `ss`   | `05`   | Second with leading zero                         |
| `s`    | `5`    | Second without leading zero                      |
| `.0`    | `.00`    | Microsecond with leading zero                    |
| `.9`    | `.99`    | Microsecond without leading zero                 |

### Timezone Formats

| Format | Output  | Description                                        |
| ------ | ------- | -------------------------------------------------- |
| `z`    | `Z`     | The Z literal represents UTC                       |
| `zz`   | `MST`   | Timezone abbreviation                              |
| `o`    | `±07`   | Timezone offset with leading zero (only hours)     |
| `oo`   | `±0700` | Timezone offset with leading zero without colon    |
| `ooo`  | `±07:00`| Timezone offset with leading zero with colon       |

### Built-in Formats

| Layout Name        | Output                                    |
| ------------------ | ----------------------------------------- |
| `time.Layout`      | `2006-01-02 15:04:05.999999999 -0700 MST` |
| `time.ANSIC`       | `Mon Jan  2 15:04:05 2006`               |
| `time.UnixDate`    | `Mon Jan  2 15:04:05 MST 2006`           |
| `time.RubyDate`    | `Mon Jan 02 15:04:05 -0700 2006`         |
| `time.RFC822`      | `02 Jan 06 15:04 MST`                    |
| `time.RFC822Z`     | `02 Jan 06 15:04 -0700`                  |
| `time.RFC850`      | `Monday, 02-Jan-06 15:04:05 MST`         |
| `time.RFC1123`     | `Mon, 02 Jan 2006 15:04:05 MST`          |
| `time.RFC1123Z`    | `Mon, 02 Jan 2006 15:04:05 -0700`        |
| `time.RFC3339`     | `2006-01-02T15:04:05Z07:00`              |
| `time.RFC3339Nano` | `2006-01-02T15:04:05.999999999Z07:00`    |
| `time.Kitchen`     | `3:04PM`                                 |

### Handy Time Stamps

| Layout Name        | Output                      |
| ------------------ | --------------------------- |
| `time.Stamp`       | `Jan  2 15:04:05`           |
| `time.StampMilli`  | `Jan  2 15:04:05.000`       |
| `time.StampMicro`  | `Jan  2 15:04:05.000000`    |
| `time.StampNano`   | `Jan  2 15:04:05.000000000` |

### Additional Date and Time Formats

| Layout Name        | Output                  |
| ------------------ | ----------------------- |
| `time.DateTime`    | `2006-01-02 15:04:05`   |
| `time.DateOnly`    | `2006-01-02`            |
| `time.TimeOnly`    | `15:04:05`              |

Temporal provides a comprehensive range of specifiers for all your date and time formatting needs, making it an indispensable tool for Go developers.
