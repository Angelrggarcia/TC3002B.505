package main

import (
	"fmt"
)

// Stack representa una estructura LIFO (último en entrar, primero en salir)
type Stack[T any] struct {
	elementos []T // Slice que almacena los elementos de la pila
}

// NewStack crea una nueva pila vacía
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{elementos: make([]T, 0)}
}

// Push añade un elemento a la parte superior de la pila
func (s *Stack[T]) Push(item T) {
	s.elementos = append(s.elementos, item)
}

// Pop elimina y devuelve el elemento superior de la pila
// Retorna el elemento y un booleano indicando si la operación fue exitosa
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elementos) == 0 {
		var zero T // Valor cero del tipo T
		return zero, false
	}
	item := s.elementos[len(s.elementos)-1]
	s.elementos = s.elementos[:len(s.elementos)-1]
	return item, true
}

// Peek devuelve el elemento superior sin eliminarlo de la pila
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elementos) == 0 {
		var zero T
		return zero, false
	}
	return s.elementos[len(s.elementos)-1], true
}

// IsEmpty verifica si la pila está vacía
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elementos) == 0
}

// Size devuelve el número de elementos en la pila
func (s *Stack[T]) Size() int {
	return len(s.elementos)
}

// Queue representa una estructura FIFO (primero en entrar, primero en salir)
type Queue[T any] struct {
	elementos []T // Slice que almacena los elementos de la cola
}

// NewQueue crea una nueva cola vacía
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{elementos: make([]T, 0)}
}

// Enqueue añade un elemento al final de la cola
func (q *Queue[T]) Enqueue(item T) {
	q.elementos = append(q.elementos, item)
}

// Dequeue elimina y devuelve el primer elemento de la cola
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.elementos) == 0 {
		var zero T
		return zero, false
	}
	item := q.elementos[0]
	q.elementos = q.elementos[1:]
	return item, true
}

// Peek devuelve el primer elemento sin eliminarlo de la cola
func (q *Queue[T]) Peek() (T, bool) {
	if len(q.elementos) == 0 {
		var zero T
		return zero, false
	}
	return q.elementos[0], true
}

// IsEmpty verifica si la cola está vacía
func (q *Queue[T]) IsEmpty() bool {
	return len(q.elementos) == 0
}

// Size devuelve el número de elementos en la cola
func (q *Queue[T]) Size() int {
	return len(q.elementos)
}

// Para el diccionario usaremos el map nativo de Go
// No es necesario implementarlo ya que Go tiene una excelente implementación incorporada

func main() {
	// Demostración del Stack
	fmt.Println("=== Demostración de Stack ===")
	pila := NewStack[int]()

	pila.Push(1)
	pila.Push(2)
	pila.Push(3)

	fmt.Printf("Tamaño de la pila: %d\n", pila.Size())
	if top, ok := pila.Peek(); ok {
		fmt.Printf("Elemento superior: %d\n", top)
	}

	for !pila.IsEmpty() {
		if item, ok := pila.Pop(); ok {
			fmt.Printf("Elemento sacado: %d\n", item)
		}
	}

	fmt.Printf("Pila vacía: %v\n", pila.IsEmpty())

	// Demostración de la Queue
	fmt.Println("\n=== Demostración de Queue ===")
	cola := NewQueue[string]()

	cola.Enqueue("Primero")
	cola.Enqueue("Segundo")
	cola.Enqueue("Tercero")

	fmt.Printf("Tamaño de la cola: %d\n", cola.Size())
	if front, ok := cola.Peek(); ok {
		fmt.Printf("Primer elemento: %s\n", front)
	}

	for !cola.IsEmpty() {
		if item, ok := cola.Dequeue(); ok {
			fmt.Printf("Elemento sacado: %s\n", item)
		}
	}

	fmt.Printf("Cola vacía: %v\n", cola.IsEmpty())

	// Demostración del Map
	fmt.Println("\n=== Demostración de Map ===")
	diccionario := make(map[string]string)

	diccionario["clave1"] = "valor1"
	diccionario["clave2"] = "valor2"
	diccionario["clave3"] = "valor3"

	fmt.Printf("Valor para 'clave2': %s\n", diccionario["clave2"])

	// Verificar si una clave existe
	if valor, existe := diccionario["clave4"]; existe {
		fmt.Println(valor)
	} else {
		fmt.Println("'clave4' no existe")
	}

	// Iterar sobre el map
	for clave, valor := range diccionario {
		fmt.Printf("%s: %s\n", clave, valor)
	}

	// Eliminar un elemento
	delete(diccionario, "clave1")
	fmt.Printf("Tamaño después de eliminar: %d\n", len(diccionario))
}

// Tests para verificar el funcionamiento
func testEstructuras() {
	fmt.Println("\n=== Ejecutando tests ===")

	// Test Stack
	pila := NewStack[int]()
	if !pila.IsEmpty() {
		fmt.Println("ERROR: La pila nueva debería estar vacía")
	}

	pila.Push(10)
	pila.Push(20)
	if pila.Size() != 2 {
		fmt.Println("ERROR: Tamaño incorrecto de la pila")
	}

	if top, _ := pila.Peek(); top != 20 {
		fmt.Println("ERROR: El elemento superior debería ser 20")
	}

	if item, _ := pila.Pop(); item != 20 {
		fmt.Println("ERROR: Debería haber sacado 20")
	}

	if item, _ := pila.Pop(); item != 10 {
		fmt.Println("ERROR: Debería haber sacado 10")
	}

	if _, ok := pila.Pop(); ok {
		fmt.Println("ERROR: No debería poder sacar de pila vacía")
	}

	// Test Queue
	cola := NewQueue[string]()
	if !cola.IsEmpty() {
		fmt.Println("ERROR: La cola nueva debería estar vacía")
	}

	cola.Enqueue("a")
	cola.Enqueue("b")
	if cola.Size() != 2 {
		fmt.Println("ERROR: Tamaño incorrecto de la cola")
	}

	if front, _ := cola.Peek(); front != "a" {
		fmt.Println("ERROR: El primer elemento debería ser 'a'")
	}

	if item, _ := cola.Dequeue(); item != "a" {
		fmt.Println("ERROR: Debería haber sacado 'a'")
	}

	if item, _ := cola.Dequeue(); item != "b" {
		fmt.Println("ERROR: Debería haber sacado 'b'")
	}

	if _, ok := cola.Dequeue(); ok {
		fmt.Println("ERROR: No debería poder sacar de cola vacía")
	}

	// Test Map
	miMap := make(map[int]string)
	if len(miMap) != 0 {
		fmt.Println("ERROR: El map nuevo debería estar vacío")
	}

	miMap[1] = "uno"
	miMap[2] = "dos"
	if len(miMap) != 2 {
		fmt.Println("ERROR: Tamaño incorrecto del map")
	}

	if valor := miMap[1]; valor != "uno" {
		fmt.Println("ERROR: Valor incorrecto para clave 1")
	}

	delete(miMap, 1)
	if len(miMap) != 1 {
		fmt.Println("ERROR: Tamaño incorrecto después de eliminar")
	}

	if _, existe := miMap[1]; existe {
		fmt.Println("ERROR: La clave 1 debería haber sido eliminada")
	}

	fmt.Println("Tests completados")
}

func main() {
	// Ejecutar demostración
	main()

	// Ejecutar tests
	testEstructuras()
}
