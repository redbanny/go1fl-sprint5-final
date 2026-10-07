package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

var argException = errors.New("Incorrect input data")
var negativeArgException = errors.New("value must be greater than zero")
var negativeResult = errors.New("The calculation result is negative. Check the input parameters.")
var unknownTrainingType = errors.New("неизвестный тип тренировки")

func (t *Training) Parse(datastring string) (err error) {
	splitData := strings.Split(datastring, ",")
	if len(splitData) != 3 {
		return argException
	}

	stepsCount, err := strconv.Atoi(splitData[0])
	if err != nil {
		return err
	}
	if stepsCount <= 0 {
		return negativeArgException
	}

	trainingDuration, err := time.ParseDuration(splitData[2])
	if err != nil {
		return err
	}
	if trainingDuration <= 0 {
		return negativeArgException
	}

	t.Steps = stepsCount
	t.TrainingType = splitData[1]
	t.Duration = trainingDuration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	var kkalCount float64
	var err error
	distance := spentenergy.Distance(t.Steps, t.Height)
	if distance <= 0 {
		return "", fmt.Errorf("Distance: %w", negativeResult)
	}

	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	if meanSpeed <= 0 {
		return "", fmt.Errorf("Mean speed: %w", negativeResult)
	}

	switch strings.ToLower(t.TrainingType) {
	case "ходьба":
		kkalCount, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "бег":
		kkalCount, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		err = unknownTrainingType
	}

	if err != nil {
		return "", fmt.Errorf("WalkingSpentCalories: %w", err)
	}

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, meanSpeed, kkalCount)

	return result, nil
}
