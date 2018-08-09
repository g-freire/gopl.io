// responsabilidade: contar os bytes escritos antes destes serem descartados, serializacao int para bytecounter
// *ByteCounter sastisfaz o contrato de io.Writer, que por sua vez satizfaz o contrato de Fprint, responsavel pela formatacao em comum
// dos metodos do pacote fmt

package main

import (
	"fmt"
)

type ByteCounter int

// metodo responsavel por contagem de bytes na stream
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}
// debugando v->v,e espacos 
func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")
	
	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
}
