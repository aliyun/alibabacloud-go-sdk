// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressionVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateExpressionVariableRequest
	GetLang() *string
	SetDescription(v string) *CreateExpressionVariableRequest
	GetDescription() *string
	SetEventCode(v string) *CreateExpressionVariableRequest
	GetEventCode() *string
	SetExpression(v string) *CreateExpressionVariableRequest
	GetExpression() *string
	SetExpressionTitle(v string) *CreateExpressionVariableRequest
	GetExpressionTitle() *string
	SetExpressionVariable(v string) *CreateExpressionVariableRequest
	GetExpressionVariable() *string
	SetOutlier(v string) *CreateExpressionVariableRequest
	GetOutlier() *string
	SetOutputs(v string) *CreateExpressionVariableRequest
	GetOutputs() *string
	SetRegId(v string) *CreateExpressionVariableRequest
	GetRegId() *string
	SetTitle(v string) *CreateExpressionVariableRequest
	GetTitle() *string
}

type CreateExpressionVariableRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Description.
	//
	// example:
	//
	// 获取入参的手机号前7位
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Event code
	//
	// This parameter is required.
	//
	// example:
	//
	// de_ahpayh4121
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Calculation expression
	//
	// This parameter is required.
	//
	// example:
	//
	// @ex_GX9rrlTq4b67 + 1001
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// Display value of calculation expression
	//
	// This parameter is required.
	//
	// example:
	//
	// @selfvariable_02 + 1001
	ExpressionTitle *string `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	// Calculation expression variable
	//
	// example:
	//
	// [{"name":"ex_GX9rrlTq4b67","code":"deInvokeSelfVariable(44659)","fieldType":"INT"}]
	ExpressionVariable *string `json:"expressionVariable,omitempty" xml:"expressionVariable,omitempty"`
	// Outlier
	//
	// This parameter is required.
	//
	// example:
	//
	// -1
	Outlier *string `json:"outlier,omitempty" xml:"outlier,omitempty"`
	// Variable return type
	//
	// This parameter is required.
	//
	// example:
	//
	// STRING
	Outputs *string `json:"outputs,omitempty" xml:"outputs,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Title.
	//
	// This parameter is required.
	//
	// example:
	//
	// 获取手机号前7位
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateExpressionVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressionVariableRequest) GoString() string {
	return s.String()
}

func (s *CreateExpressionVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateExpressionVariableRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateExpressionVariableRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *CreateExpressionVariableRequest) GetExpression() *string {
	return s.Expression
}

func (s *CreateExpressionVariableRequest) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *CreateExpressionVariableRequest) GetExpressionVariable() *string {
	return s.ExpressionVariable
}

func (s *CreateExpressionVariableRequest) GetOutlier() *string {
	return s.Outlier
}

func (s *CreateExpressionVariableRequest) GetOutputs() *string {
	return s.Outputs
}

func (s *CreateExpressionVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateExpressionVariableRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateExpressionVariableRequest) SetLang(v string) *CreateExpressionVariableRequest {
	s.Lang = &v
	return s
}

func (s *CreateExpressionVariableRequest) SetDescription(v string) *CreateExpressionVariableRequest {
	s.Description = &v
	return s
}

func (s *CreateExpressionVariableRequest) SetEventCode(v string) *CreateExpressionVariableRequest {
	s.EventCode = &v
	return s
}

func (s *CreateExpressionVariableRequest) SetExpression(v string) *CreateExpressionVariableRequest {
	s.Expression = &v
	return s
}

func (s *CreateExpressionVariableRequest) SetExpressionTitle(v string) *CreateExpressionVariableRequest {
	s.ExpressionTitle = &v
	return s
}

func (s *CreateExpressionVariableRequest) SetExpressionVariable(v string) *CreateExpressionVariableRequest {
	s.ExpressionVariable = &v
	return s
}

func (s *CreateExpressionVariableRequest) SetOutlier(v string) *CreateExpressionVariableRequest {
	s.Outlier = &v
	return s
}

func (s *CreateExpressionVariableRequest) SetOutputs(v string) *CreateExpressionVariableRequest {
	s.Outputs = &v
	return s
}

func (s *CreateExpressionVariableRequest) SetRegId(v string) *CreateExpressionVariableRequest {
	s.RegId = &v
	return s
}

func (s *CreateExpressionVariableRequest) SetTitle(v string) *CreateExpressionVariableRequest {
	s.Title = &v
	return s
}

func (s *CreateExpressionVariableRequest) Validate() error {
	return dara.Validate(s)
}
