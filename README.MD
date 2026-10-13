# ESEP Grade Calculator

## Summary + Getting Started
This library is a simple utility to calculate your grades. 
To import this library into your source code, use:

```
import esep/grade-calculator/esepunittests
```

## Key Features

The Grade Calculator contains the state involving all current assignments, essays, and exams that are graded. Along with this, it provides methods to manipulate its state and calculate a final grade.

You can instantiate one with `NewGradeCalculator()`.

### GradeType
A variable of type `gradeType` can either be an `assignment`, `essay`, or `exam`.

### Add Grade
Use 
```
func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType)
```
to add a grade to the Grade Calculator.

### Get Final Grade
Use 
```
func (gc *GradeCalculator) GetFinalGrade() string
```
to calcualate the final grade based on the current grades added to the calculator.

## Runtime requirements:
Golang 1.24+


## Testing
Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```