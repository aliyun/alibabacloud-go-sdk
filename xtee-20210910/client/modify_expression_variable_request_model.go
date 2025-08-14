// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressionVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifyExpressionVariableRequest
	GetLang() *string
	SetDataVersion(v int64) *ModifyExpressionVariableRequest
	GetDataVersion() *int64
	SetDescription(v string) *ModifyExpressionVariableRequest
	GetDescription() *string
	SetEventCode(v string) *ModifyExpressionVariableRequest
	GetEventCode() *string
	SetExpression(v string) *ModifyExpressionVariableRequest
	GetExpression() *string
	SetExpressionTitle(v string) *ModifyExpressionVariableRequest
	GetExpressionTitle() *string
	SetExpressionVariable(v string) *ModifyExpressionVariableRequest
	GetExpressionVariable() *string
	SetId(v int64) *ModifyExpressionVariableRequest
	GetId() *int64
	SetName(v string) *ModifyExpressionVariableRequest
	GetName() *string
	SetOutlier(v string) *ModifyExpressionVariableRequest
	GetOutlier() *string
	SetOutputs(v string) *ModifyExpressionVariableRequest
	GetOutputs() *string
	SetRegId(v string) *ModifyExpressionVariableRequest
	GetRegId() *string
	SetTitle(v string) *ModifyExpressionVariableRequest
	GetTitle() *string
}

type ModifyExpressionVariableRequest struct {
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
	// Data version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	DataVersion *int64 `json:"dataVersion,omitempty" xml:"dataVersion,omitempty"`
	// Description.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Event code
	//
	// This parameter is required.
	//
	// example:
	//
	// de_acytyt7036
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Expression
	//
	// This parameter is required.
	//
	// example:
	//
	// @ex_GX9rrlTq4b67 + 1001
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// Expression display
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
	// ex_GX9rrlTq4b67
	ExpressionVariable *string `json:"expressionVariable,omitempty" xml:"expressionVariable,omitempty"`
	// Variable ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 3144
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Variable name
	//
	// example:
	//
	// ex_NgR6nDVD821c
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Outlier
	//
	// This parameter is required.
	//
	// example:
	//
	// -1
	Outlier *string `json:"outlier,omitempty" xml:"outlier,omitempty"`
	// Output
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
	// 变量标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ModifyExpressionVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressionVariableRequest) GoString() string {
	return s.String()
}

func (s *ModifyExpressionVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyExpressionVariableRequest) GetDataVersion() *int64 {
	return s.DataVersion
}

func (s *ModifyExpressionVariableRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyExpressionVariableRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *ModifyExpressionVariableRequest) GetExpression() *string {
	return s.Expression
}

func (s *ModifyExpressionVariableRequest) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *ModifyExpressionVariableRequest) GetExpressionVariable() *string {
	return s.ExpressionVariable
}

func (s *ModifyExpressionVariableRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyExpressionVariableRequest) GetName() *string {
	return s.Name
}

func (s *ModifyExpressionVariableRequest) GetOutlier() *string {
	return s.Outlier
}

func (s *ModifyExpressionVariableRequest) GetOutputs() *string {
	return s.Outputs
}

func (s *ModifyExpressionVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *ModifyExpressionVariableRequest) GetTitle() *string {
	return s.Title
}

func (s *ModifyExpressionVariableRequest) SetLang(v string) *ModifyExpressionVariableRequest {
	s.Lang = &v
	return s
}

func (s *ModifyExpressionVariableRequest) SetDataVersion(v int64) *ModifyExpressionVariableRequest {
	s.DataVersion = &v
	return s
}

func (s *ModifyExpressionVariableRequest) SetDescription(v string) *ModifyExpressionVariableRequest {
	s.Description = &v
	return s
}

func (s *ModifyExpressionVariableRequest) SetEventCode(v string) *ModifyExpressionVariableRequest {
	s.EventCode = &v
	return s
}

func (s *ModifyExpressionVariableRequest) SetExpression(v string) *ModifyExpressionVariableRequest {
	s.Expression = &v
	return s
}

func (s *ModifyExpressionVariableRequest) SetExpressionTitle(v string) *ModifyExpressionVariableRequest {
	s.ExpressionTitle = &v
	return s
}

func (s *ModifyExpressionVariableRequest) SetExpressionVariable(v string) *ModifyExpressionVariableRequest {
	s.ExpressionVariable = &v
	return s
}

func (s *ModifyExpressionVariableRequest) SetId(v int64) *ModifyExpressionVariableRequest {
	s.Id = &v
	return s
}

func (s *ModifyExpressionVariableRequest) SetName(v string) *ModifyExpressionVariableRequest {
	s.Name = &v
	return s
}

func (s *ModifyExpressionVariableRequest) SetOutlier(v string) *ModifyExpressionVariableRequest {
	s.Outlier = &v
	return s
}

func (s *ModifyExpressionVariableRequest) SetOutputs(v string) *ModifyExpressionVariableRequest {
	s.Outputs = &v
	return s
}

func (s *ModifyExpressionVariableRequest) SetRegId(v string) *ModifyExpressionVariableRequest {
	s.RegId = &v
	return s
}

func (s *ModifyExpressionVariableRequest) SetTitle(v string) *ModifyExpressionVariableRequest {
	s.Title = &v
	return s
}

func (s *ModifyExpressionVariableRequest) Validate() error {
	return dara.Validate(s)
}
