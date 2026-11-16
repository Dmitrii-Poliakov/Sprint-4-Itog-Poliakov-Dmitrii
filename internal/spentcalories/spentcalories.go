package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	Split := strings.Split(data, ",")
	if len(Split) != 3 {
		return 0, "", 0, errors.New("wrong value count")
	}
	StepStrConv, err := strconv.Atoi(Split[0])
	if err != nil {
		return 0, "", 0, errors.New("conversion error string-int")
	}
	if StepStrConv <= 0 {
		return 0, "", 0, errors.New("steps count: 0")
	}
	TimeStrConv, err := time.ParseDuration(Split[2])
	if err != nil {
		return 0, "", 0, errors.New("conversion error string-time")
	}
	if TimeStrConv <= 0 {
		return 0, "", 0, errors.New("duration value: 0")
	}
	return StepStrConv, Split[1], TimeStrConv, nil
}

func distance(steps int, height float64) float64 {
	CalcDistance := (float64(steps) * (height * stepLengthCoefficient)) / float64(mInKm)
	return CalcDistance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	Distance := distance(steps, height)
	AverageSpeed := Distance / duration.Hours()
	return AverageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, ActivityType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if steps <= 0 {
		err := errors.New("steps: 0")
		log.Println("error:", err)
		return "", err
	}
	if weight <= 0 {
		err := errors.New("weight: 0")
		log.Println("error:", err)
		return "", err
	}
	if height <= 0 {
		err := errors.New("height: 0")
		log.Println("error:", err)
		return "", err
	}
	if duration <= 0 {
		err := errors.New("duration: 0")
		log.Println("error:", err)
		return "", err
	}

	Distance := distance(steps, height)
	AverageSpeed := meanSpeed(steps, height, duration)

	var SpentCalories float64
	var ErrCalories error

	switch ActivityType {
	case "Ходьба":
		SpentCalories, ErrCalories = WalkingSpentCalories(steps, weight, height, duration)
	case "Бег":
		SpentCalories, ErrCalories = RunningSpentCalories(steps, weight, height, duration)
	default:
		err := fmt.Errorf("неизвестный тип тренировки")
		log.Println("error:", err)
		return "", err
	}
	if ErrCalories != nil {
		log.Println(ErrCalories)
		return "", ErrCalories
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", ActivityType, duration.Hours(), Distance, AverageSpeed, SpentCalories), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps: 0")
	}
	if weight <= 0 {
		return 0, errors.New("weight: 0")
	}
	if height <= 0 {
		return 0, errors.New("height: 0")
	}
	if duration <= 0 {
		return 0, errors.New("duration: 0")
	}
	AverageSpeed := meanSpeed(steps, height, duration)
	DurationMin := duration.Minutes()
	CalcSpentCalories := (weight * AverageSpeed * DurationMin) / minInH
	return CalcSpentCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		return 0, errors.New("steps: 0")
	}
	if weight <= 0 {
		return 0, errors.New("weight: 0")
	}
	if height <= 0 {
		return 0, errors.New("height: 0")
	}
	if duration <= 0 {
		return 0, errors.New("duration: 0")
	}
	AverageSpeed := meanSpeed(steps, height, duration)
	DurationMin := duration.Minutes()
	CalcCalories := (weight * AverageSpeed * DurationMin) / minInH
	CalcSpentCalories := CalcCalories * walkingCaloriesCoefficient
	return CalcSpentCalories, nil
}
