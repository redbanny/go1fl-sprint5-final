package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

var err = errors.New("Argument exection")

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	kkalCount := (weight * meanSpeed * float64(duration.Minutes())) / minInH

	return kkalCount * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	return (weight * meanSpeed * float64(duration.Minutes())) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// В ТЗ указано проверить значение duration.
	// Но я не совсем понимаю зачем это условие, если его можно проверить
	// перед вызовом метода MeanSpeed()
	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient

	return (stepLength * float64(steps)) / mInKm
}
