package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

var argException = errors.New("Incorrect input data")
var negativeArgException = errors.New("value must be greater than zero")
var negativeResult = errors.New("The calculation result is negative. Check the input parameters.")

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	splitData := strings.Split(datastring, ",")
	if len(splitData) != 2 {
		return argException
	}

	stepsCount, err := strconv.Atoi(splitData[0])
	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}
	if stepsCount <= 0 {
		return fmt.Errorf("steps %w", negativeArgException)
	}

	trainingDuration, err := time.ParseDuration(splitData[1])
	if err != nil {
		return fmt.Errorf("invalid training duration format: %w", err)
	}
	if trainingDuration <= 0 {
		return fmt.Errorf("training duration %w", negativeArgException)
	}

	ds.Steps = stepsCount
	ds.Duration = trainingDuration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	var kkalCount float64
	var err error

	distance := spentenergy.Distance(ds.Steps, ds.Height)
	if distance <= 0 {
		return "", fmt.Errorf("ActionInfo: calculate distance error => %w", negativeResult)
	}

	kkalCount, err = spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ActionInfo: calculate kkalCount error: %w", err)
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, kkalCount)

	return result, nil
}
