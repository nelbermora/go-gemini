package gemini

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	Init()
	out, err := AskForData("Puedes generar preguntas y respuestas frecuentes del Producto de Nombre Ranger XLS de la marca Ford?", "")
	log.Println(out)
	assert.NotNil(t, err)
}
