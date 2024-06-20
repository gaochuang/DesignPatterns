package main

import "fmt"

// 工厂接口，由具体工厂类实现
type OperatorFactory interface {
	Create() MathOperator
}

// 实际产品实现接口-- 表示数学运算器实现应该哪些行为
type MathOperator interface {
	SetOperatorA(int)
	SetOperatorB(int)
	ComputeResult() int
}

// 所有Operator的基类
type BaseOperator struct {
	operandA, operandB int
}

func (o *BaseOperator) SetOperatorA(operandA int) {
	o.operandA = operandA
}

func (o *BaseOperator) SetOperatorB(operandB int) {
	o.operandB = operandB
}

// PlusOperatorFactory 是 PlusOperator 加法运算器的工厂类
type PlusOperatorFactory struct {
}

func (pl *PlusOperatorFactory) Create() MathOperator {
	return &PlusOperator{
		BaseOperator: &BaseOperator{},
	}
}

// ComputeResult 计算并获取结果
func (p *PlusOperator) ComputeResult() int {
	return p.operandA + p.operandB
}

// PlusOperator 实际的产品类--加法运算器
type PlusOperator struct {
	*BaseOperator
}

// MultiOperatorFactory 是乘法运算器产品的工厂
type MultiOperatorFactory struct{}

func (mf *MultiOperatorFactory) Create() MathOperator {
	return &MultiOperator{
		BaseOperator: &BaseOperator{},
	}
}

func (m *MultiOperator) ComputeResult() int {
	return m.operandA * m.operandB
}

// MultiOperator 实际的产品类--乘法运算器
type MultiOperator struct {
	*BaseOperator
}

func main() {

	var factory OperatorFactory
	var mathOp MathOperator

	factory = &PlusOperatorFactory{}
	mathOp = factory.Create()
	mathOp.SetOperatorA(3)
	mathOp.SetOperatorA(2)
	fmt.Printf("Plus operation reuslt: %d\n", mathOp.ComputeResult())

	factory = &MultiOperatorFactory{}
	mathOp = factory.Create()
	mathOp.SetOperatorA(3)
	mathOp.SetOperatorB(2)
	fmt.Printf("Multiple operation reuslt: %d\n", mathOp.ComputeResult())
}
