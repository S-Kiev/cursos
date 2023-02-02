package cursos

import "fmt"

type Curso struct {
	Nombre     string
	Precio     float64
	EsGratis   bool
	IdUsuarios []uint
	Clases     map[uint]string
}

func (c Curso) MostrarClases() {
	text := "Los cursos son: "
	for _, clase := range c.Clases {
		text += clase + ", "
	}

	fmt.Println(text[:len(text)-2])
}

func (c *Curso) CambiarPrecio(nuevoPrecio float64) {
	c.Precio = nuevoPrecio
}
