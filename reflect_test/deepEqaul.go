package reflecttest

import "reflect"

func reflectDeepCompare(x, y *string) bool {
	return reflect.DeepEqual(x, y)
}

func normalCompare(x, y *string) bool {
	return *x == *y
}
