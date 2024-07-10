package utils

import "time"

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func ToMap[K string, V any](arr []V, groupingFunc func(val V) K) map[K]V {
	result := make(map[K]V)
	for _, v := range arr {
		key := groupingFunc(v)
		result[key] = v
	}
	return result
}

func SleepWithInterruption(getSleepTimeFunc func() time.Duration, checkInterval time.Duration) {
	sleepFor := getSleepTimeFunc()
	if sleepFor < checkInterval {
		time.Sleep(sleepFor)
	} else {
		slept := time.Duration(0)
		for {
			left := sleepFor - slept
			if left > checkInterval {
				time.Sleep(checkInterval)
				slept += checkInterval
			} else {
				if left > 0 {
					time.Sleep(left)
				}
				return
			}
			sleepFor = getSleepTimeFunc()
		}
	}
}
