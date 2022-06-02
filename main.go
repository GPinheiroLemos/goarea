package goarea

import "math"

//Pi
const Pi = 3.1416

//Circ é responsável por calcular a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Calcula área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
