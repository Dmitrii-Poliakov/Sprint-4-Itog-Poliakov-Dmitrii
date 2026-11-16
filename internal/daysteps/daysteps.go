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
		return 0, 0, errors.New("wrong value count")
	}
	StepStrConv, err := strconv.Atoi(Split[0])
	if err != nil {
		return 0, 0, errors.New("conversion error string-int")
	}
	if StepStrConv <= 0 {
		return 0, 0, errors.New("steps count: 0")
	}
	TimeStrConv, err := time.ParseDuration(Split[1])
	if err != nil {
		return 0, 0, errors.New("сonversion error string-time")
	}
	if TimeStrConv <= 0 {
		return 0, 0, errors.New("duration value: 0")
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
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, DistanceKm, SpentCalories)
}
