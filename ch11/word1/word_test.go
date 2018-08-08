// TDD: Test Driven Development 
// este é um teste de caixa preta, limitado ao o que sua API e documentacao informam, nao tem acesso privilegiado de funcoes e variaveis globais 
// $	go test 
// reparar que arquivo termina em _test
// chama funcao de teste com resultados PASS / FAIL (analise relatorio de bugs)
// responsavilidade de pacote: fazer assercoes sobre a qualidade da API em desenvolvimento

// nome do projeto
package word
// importa as funcoes 
import "testing"

// testa funcao com parametro verdadeiro retornando falso
func TestPalindrome(t *testing.T) { // o parametro t oferece metodos para informar falhas e log de erros
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	} 
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}
// testa funcao com parametro verdadeiro retornando falso
func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

//  testa com entradas em diferentes dialetos
func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}
//  testa com espacos entre as strings
func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}

