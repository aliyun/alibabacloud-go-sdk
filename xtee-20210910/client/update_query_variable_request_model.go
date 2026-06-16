// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQueryVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateQueryVariableRequest
	GetLang() *string
	SetDataSourceCode(v string) *UpdateQueryVariableRequest
	GetDataSourceCode() *string
	SetDescription(v string) *UpdateQueryVariableRequest
	GetDescription() *string
	SetEventCode(v string) *UpdateQueryVariableRequest
	GetEventCode() *string
	SetExpression(v string) *UpdateQueryVariableRequest
	GetExpression() *string
	SetExpressionTitle(v string) *UpdateQueryVariableRequest
	GetExpressionTitle() *string
	SetExpressionVariable(v string) *UpdateQueryVariableRequest
	GetExpressionVariable() *string
	SetId(v int64) *UpdateQueryVariableRequest
	GetId() *int64
	SetOutlier(v string) *UpdateQueryVariableRequest
	GetOutlier() *string
	SetOutputs(v string) *UpdateQueryVariableRequest
	GetOutputs() *string
	SetRegId(v string) *UpdateQueryVariableRequest
	GetRegId() *string
	SetTitle(v string) *UpdateQueryVariableRequest
	GetTitle() *string
}

type UpdateQueryVariableRequest struct {
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The data source code.
	//
	// This parameter is required.
	//
	// example:
	//
	// ds_vcaoii1697
	DataSourceCode *string `json:"dataSourceCode,omitempty" xml:"dataSourceCode,omitempty"`
	// The description.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The event code.
	//
	// This parameter is required.
	//
	// example:
	//
	// de_ajnoqe2016
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// The expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT  AVG( $source )\\nFROM ds_vcaoii1697 \\nWHERE  $age > 0
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// The display expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT  AVG( $source )\\nFROM testCase\\nWHERE  $age > 0
	ExpressionTitle *string `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	// The variable associated with the expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// age
	ExpressionVariable *string `json:"expressionVariable,omitempty" xml:"expressionVariable,omitempty"`
	// The primary key ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3144
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
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
	// 获取年龄大于30的数据
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateQueryVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQueryVariableRequest) GoString() string {
	return s.String()
}

func (s *UpdateQueryVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateQueryVariableRequest) GetDataSourceCode() *string {
	return s.DataSourceCode
}

func (s *UpdateQueryVariableRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateQueryVariableRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *UpdateQueryVariableRequest) GetExpression() *string {
	return s.Expression
}

func (s *UpdateQueryVariableRequest) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *UpdateQueryVariableRequest) GetExpressionVariable() *string {
	return s.ExpressionVariable
}

func (s *UpdateQueryVariableRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateQueryVariableRequest) GetOutlier() *string {
	return s.Outlier
}

func (s *UpdateQueryVariableRequest) GetOutputs() *string {
	return s.Outputs
}

func (s *UpdateQueryVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *UpdateQueryVariableRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateQueryVariableRequest) SetLang(v string) *UpdateQueryVariableRequest {
	s.Lang = &v
	return s
}

func (s *UpdateQueryVariableRequest) SetDataSourceCode(v string) *UpdateQueryVariableRequest {
	s.DataSourceCode = &v
	return s
}

func (s *UpdateQueryVariableRequest) SetDescription(v string) *UpdateQueryVariableRequest {
	s.Description = &v
	return s
}

func (s *UpdateQueryVariableRequest) SetEventCode(v string) *UpdateQueryVariableRequest {
	s.EventCode = &v
	return s
}

func (s *UpdateQueryVariableRequest) SetExpression(v string) *UpdateQueryVariableRequest {
	s.Expression = &v
	return s
}

func (s *UpdateQueryVariableRequest) SetExpressionTitle(v string) *UpdateQueryVariableRequest {
	s.ExpressionTitle = &v
	return s
}

func (s *UpdateQueryVariableRequest) SetExpressionVariable(v string) *UpdateQueryVariableRequest {
	s.ExpressionVariable = &v
	return s
}

func (s *UpdateQueryVariableRequest) SetId(v int64) *UpdateQueryVariableRequest {
	s.Id = &v
	return s
}

func (s *UpdateQueryVariableRequest) SetOutlier(v string) *UpdateQueryVariableRequest {
	s.Outlier = &v
	return s
}

func (s *UpdateQueryVariableRequest) SetOutputs(v string) *UpdateQueryVariableRequest {
	s.Outputs = &v
	return s
}

func (s *UpdateQueryVariableRequest) SetRegId(v string) *UpdateQueryVariableRequest {
	s.RegId = &v
	return s
}

func (s *UpdateQueryVariableRequest) SetTitle(v string) *UpdateQueryVariableRequest {
	s.Title = &v
	return s
}

func (s *UpdateQueryVariableRequest) Validate() error {
	return dara.Validate(s)
}
