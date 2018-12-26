package main

import "fmt"

const max_len = 1024

func MDhb(strs string) string {
	//strs to bit
	data := []byte(strs)
	data_len := len(data)

	if data_len > max_len {
		return ""
	}
	// fill
	if data_len%32 != 0 {
		c_len := (32 - data_len%32)
		new_data := make([]byte, c_len+data_len)
		copy(new_data, data)
		data = new_data
	}
	var A uint8 = 0
	var B uint8 = 9
	var C uint8 = 1
	var D uint8 = 8
	t := [...]uint8{1, 1, 1, 2}
	s := [...]uint8{0, 9, 1, 2}

	step_count := int(len(data) / 32)

	for i := 0; i != step_count; i++ {
		//step 1
		A = FF(A, B, C, D, b2unit8(data[0:8]), s[0], t[0])
		B = FF(A, B, C, D, b2unit8(data[8:16]), s[1], t[1])
		C = FF(A, B, C, D, b2unit8(data[16:24]), s[2], t[2])
		D = FF(A, B, C, D, b2unit8(data[24:32]), s[3], t[3])
		//step 2
		A = GG(A, B, C, D, b2unit8(data[0:8]), s[0], t[0])
		B = GG(A, B, C, D, b2unit8(data[8:16]), s[1], t[1])
		C = GG(A, B, C, D, b2unit8(data[16:24]), s[2], t[2])
		D = GG(A, B, C, D, b2unit8(data[24:32]), s[3], t[3])
		//step 3
		A = HH(A, B, C, D, b2unit8(data[0:8]), s[0], t[0])
		B = HH(A, B, C, D, b2unit8(data[8:16]), s[1], t[1])
		C = HH(A, B, C, D, b2unit8(data[16:24]), s[2], t[2])
		D = HH(A, B, C, D, b2unit8(data[24:32]), s[3], t[3])
		//step 4
		A = II(A, B, C, D, b2unit8(data[0:8]), s[0], t[0])
		B = II(A, B, C, D, b2unit8(data[8:16]), s[1], t[1])
		C = II(A, B, C, D, b2unit8(data[16:24]), s[2], t[2])
		D = II(A, B, C, D, b2unit8(data[24:32]), s[3], t[3])
	}
	//bit to string
	out_str := fmt.Sprintf("%02X%02X%02X%02X", A, B, C, D)
	return out_str
}
func b2unit8(data []byte) uint8 {
	var value uint8
	value = (uint8)(data[0]*128 +
		data[1]*64 +
		data[2]*32 +
		data[3]*16 +
		data[4]*8 +
		data[5]*4 +
		data[6]*2 +
		data[7]*1)
	return value
}
func F(x uint8, y uint8, z uint8) uint8 {
	return (x & y) | ((^x) & z)
}
func G(x uint8, y uint8, z uint8) uint8 {
	return (x & z) | (y & (^z))
}
func H(x uint8, y uint8, z uint8) uint8 {
	return x ^ y ^ z
}
func I(x uint8, y uint8, z uint8) uint8 {
	return y ^ (x | (^z))
}

func FF(a uint8, b uint8, c uint8, d uint8, Mj uint8, s uint8, ti uint8) uint8 {
	return b + ((a + F(b, c, d) + Mj + ti) << s)
}
func GG(a uint8, b uint8, c uint8, d uint8, Mj uint8, s uint8, ti uint8) uint8 {
	return b + ((a + G(b, c, d) + Mj + ti) << s)
}
func HH(a uint8, b uint8, c uint8, d uint8, Mj uint8, s uint8, ti uint8) uint8 {
	return b + ((a + H(b, c, d) + Mj + ti) << s)
}
func II(a uint8, b uint8, c uint8, d uint8, Mj uint8, s uint8, ti uint8) uint8 {
	return b + ((a + I(b, c, d) + Mj + ti) << s)
}

func main() {
	println(MDhb("hello,world"))
	println(MDhb("hello world"))
}
