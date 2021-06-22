package goarea

import "math"

// Pi é uma proporção numérica definida pela relação entre
// o perímetro de uma cincunferência e seu deâmetro
const Pi = 3.1416

// Circ é responsável por calcular a área da cincunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// React é responsável por calcular a área de um retangulo
func React(base, altura float64) float64 {
	return base * altura
}

// não é visível!
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
