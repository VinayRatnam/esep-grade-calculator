package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator(70)

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator(70)

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator(70)

	gradeCalculator.AddGrade("open source assignment", 0, Assignment)
	gradeCalculator.AddGrade("exam 1", 0, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 0, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator(70)

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 75, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 79, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator(70)

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 65, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 69, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
	expected_value := "exam"

	actual_value := Exam.String()

	if expected_value != actual_value {
		t.Errorf("Expected String to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetPassFailPass(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := NewGradeCalculator(70)

	gradeCalculator.AddGrade("open source assignment", 90, Assignment)
	gradeCalculator.AddGrade("exam 1", 85, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 80, Essay)

	actual_value := gradeCalculator.GetPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetPassFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetPassFailFail(t *testing.T) {
	expected_value := "Fail"

	gradeCalculator := NewGradeCalculator(70)

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 60, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 65, Essay)

	actual_value := gradeCalculator.GetPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetPassFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetPassFailAtThreshold(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := NewGradeCalculator(70)

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 70, Essay)

	actual_value := gradeCalculator.GetPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetPassFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetPassFailBelowThreshold(t *testing.T) {
	expected_value := "Fail"

	gradeCalculator := NewGradeCalculator(70)

	gradeCalculator.AddGrade("open source assignment", 69, Assignment)
	gradeCalculator.AddGrade("exam 1", 69, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 69, Essay)

	actual_value := gradeCalculator.GetPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetPassFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetPassFailCustomThreshold(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := NewGradeCalculator(60)

	gradeCalculator.AddGrade("open source assignment", 65, Assignment)
	gradeCalculator.AddGrade("exam 1", 65, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 65, Essay)

	actual_value := gradeCalculator.GetPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetPassFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetLetterGradeAndPassFailFromSameCalculator(t *testing.T) {
	expected_letter := "C"
	expected_pass_fail := "Pass"

	gradeCalculator := NewGradeCalculator(70)

	gradeCalculator.AddGrade("open source assignment", 75, Assignment)
	gradeCalculator.AddGrade("exam 1", 75, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 75, Essay)

	actual_letter := gradeCalculator.GetFinalGrade()
	actual_pass_fail := gradeCalculator.GetPassFail()

	if expected_letter != actual_letter {
		t.Errorf("Expected GetFinalGrade to return '%s'; got '%s' instead", expected_letter, actual_letter)
	}
	if expected_pass_fail != actual_pass_fail {
		t.Errorf("Expected GetPassFail to return '%s'; got '%s' instead", expected_pass_fail, actual_pass_fail)
	}
}
