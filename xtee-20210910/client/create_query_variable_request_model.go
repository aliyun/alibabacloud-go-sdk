// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueryVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateQueryVariableRequest
	GetLang() *string
	SetDataSourceCode(v string) *CreateQueryVariableRequest
	GetDataSourceCode() *string
	SetDescription(v string) *CreateQueryVariableRequest
	GetDescription() *string
	SetEventCode(v string) *CreateQueryVariableRequest
	GetEventCode() *string
	SetExpression(v string) *CreateQueryVariableRequest
	GetExpression() *string
	SetExpressionTitle(v string) *CreateQueryVariableRequest
	GetExpressionTitle() *string
	SetExpressionVariable(v string) *CreateQueryVariableRequest
	GetExpressionVariable() *string
	SetOutlier(v string) *CreateQueryVariableRequest
	GetOutlier() *string
	SetOutputs(v string) *CreateQueryVariableRequest
	GetOutputs() *string
	SetRegId(v string) *CreateQueryVariableRequest
	GetRegId() *string
	SetTitle(v string) *CreateQueryVariableRequest
	GetTitle() *string
}

type CreateQueryVariableRequest struct {
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3527
	DataSourceCode *string `json:"dataSourceCode,omitempty" xml:"dataSourceCode,omitempty"`
	// The description.
	//
	// example:
	//
	// 查询变量描述信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The event code.
	//
	// This parameter is required.
	//
	// example:
	//
	// de_arqbuy7206
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// The calculation expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT  AVG( $source )
	//
	// FROM testCase
	//
	// WHERE  $age > 0
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// The display value of the calculation expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT  AVG( $source )
	//
	// FROM testCase
	//
	// WHERE  $age > 0
	ExpressionTitle *string `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	// The variable of the calculation expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// age
	ExpressionVariable *string `json:"expressionVariable,omitempty" xml:"expressionVariable,omitempty"`
	// The outlier value.
	//
	// This parameter is required.
	//
	// example:
	//
	// -1
	Outlier *string `json:"outlier,omitempty" xml:"outlier,omitempty"`
	// The return type of the variable.
	//
	// This parameter is required.
	//
	// example:
	//
	// STRING
	Outputs *string `json:"outputs,omitempty" xml:"outputs,omitempty"`
	// The region code.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// The title.
	//
	// This parameter is required.
	//
	// example:
	//
	// 获取手机号前7位
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateQueryVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQueryVariableRequest) GoString() string {
	return s.String()
}

func (s *CreateQueryVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateQueryVariableRequest) GetDataSourceCode() *string {
	return s.DataSourceCode
}

func (s *CreateQueryVariableRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateQueryVariableRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *CreateQueryVariableRequest) GetExpression() *string {
	return s.Expression
}

func (s *CreateQueryVariableRequest) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *CreateQueryVariableRequest) GetExpressionVariable() *string {
	return s.ExpressionVariable
}

func (s *CreateQueryVariableRequest) GetOutlier() *string {
	return s.Outlier
}

func (s *CreateQueryVariableRequest) GetOutputs() *string {
	return s.Outputs
}

func (s *CreateQueryVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateQueryVariableRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateQueryVariableRequest) SetLang(v string) *CreateQueryVariableRequest {
	s.Lang = &v
	return s
}

func (s *CreateQueryVariableRequest) SetDataSourceCode(v string) *CreateQueryVariableRequest {
	s.DataSourceCode = &v
	return s
}

func (s *CreateQueryVariableRequest) SetDescription(v string) *CreateQueryVariableRequest {
	s.Description = &v
	return s
}

func (s *CreateQueryVariableRequest) SetEventCode(v string) *CreateQueryVariableRequest {
	s.EventCode = &v
	return s
}

func (s *CreateQueryVariableRequest) SetExpression(v string) *CreateQueryVariableRequest {
	s.Expression = &v
	return s
}

func (s *CreateQueryVariableRequest) SetExpressionTitle(v string) *CreateQueryVariableRequest {
	s.ExpressionTitle = &v
	return s
}

func (s *CreateQueryVariableRequest) SetExpressionVariable(v string) *CreateQueryVariableRequest {
	s.ExpressionVariable = &v
	return s
}

func (s *CreateQueryVariableRequest) SetOutlier(v string) *CreateQueryVariableRequest {
	s.Outlier = &v
	return s
}

func (s *CreateQueryVariableRequest) SetOutputs(v string) *CreateQueryVariableRequest {
	s.Outputs = &v
	return s
}

func (s *CreateQueryVariableRequest) SetRegId(v string) *CreateQueryVariableRequest {
	s.RegId = &v
	return s
}

func (s *CreateQueryVariableRequest) SetTitle(v string) *CreateQueryVariableRequest {
	s.Title = &v
	return s
}

func (s *CreateQueryVariableRequest) Validate() error {
	return dara.Validate(s)
}
