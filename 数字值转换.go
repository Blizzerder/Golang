//两个漂亮的专门用于处理类型转换的函数，确保没有发生精度的丢失


//int型转换成uint8
func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

//float64转换成int
funt IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprint("%g is out of the int32 range", x))
}
