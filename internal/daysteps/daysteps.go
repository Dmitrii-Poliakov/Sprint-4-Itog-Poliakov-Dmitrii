package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	Split := strings.Split(data, ",")
	if len(Split) != 2 {
		return 0, 0, errors.New("Неверное количество данных")
	}
	StepStrConv, err := strconv.Atoi(Split[0])
	if err != nil {
		return 0, 0, errors.New("Ошибка конвертации строка-число")
	}
	if StepStrConv <= 0 {
		return 0, 0, errors.New("Количество шагов: 0")
	}
	TimeStrConv, err := time.ParseDuration(Split[1])
	if err != nil {
		return 0, 0, errors.New("Ошибка конвертации строка-продолжительность")
	}
	return StepStrConv, TimeStrConv, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	DistanceKm := (float64(steps) * stepLength) / float64(mInKm)
	SpentCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\n Дистанция составила %.2f км.\n Вы сожгли %.2f ккал.", steps, DistanceKm, SpentCalories)
}
