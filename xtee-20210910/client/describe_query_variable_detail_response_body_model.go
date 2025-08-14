// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQueryVariableDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeQueryVariableDetailResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeQueryVariableDetailResponseBodyResultObject) *DescribeQueryVariableDetailResponseBody
	GetResultObject() *DescribeQueryVariableDetailResponseBodyResultObject
}

type DescribeQueryVariableDetailResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	ResultObject *DescribeQueryVariableDetailResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeQueryVariableDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryVariableDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQueryVariableDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQueryVariableDetailResponseBody) GetResultObject() *DescribeQueryVariableDetailResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeQueryVariableDetailResponseBody) SetRequestId(v string) *DescribeQueryVariableDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQueryVariableDetailResponseBody) SetResultObject(v *DescribeQueryVariableDetailResponseBodyResultObject) *DescribeQueryVariableDetailResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeQueryVariableDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeQueryVariableDetailResponseBodyResultObject struct {
	// Data source code.
	//
	// example:
	//
	// ds_vcaoii1697
	DataSourceCode *int64 `json:"dataSourceCode,omitempty" xml:"dataSourceCode,omitempty"`
	// Data source name
	//
	// example:
	//
	// 名称数据源
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// Description.
	//
	// example:
	//
	// 变量描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Event code
	//
	// example:
	//
	// de_aszbjb7236
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Expression.
	//
	// example:
	//
	// SELECT  AVG( $source )\\nFROM ds_vcaoii1697 \\nWHERE  $age > 0
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// Expression title.
	//
	// example:
	//
	// SELECT  AVG( $source )\\nFROM testCase\\nWHERE  $age > 0
	ExpressionTitle *string `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	// Expression variable.
	//
	// example:
	//
	// [96426]
	ExpressionVariable *string `json:"expressionVariable,omitempty" xml:"expressionVariable,omitempty"`
	// Variable ID
	//
	// example:
	//
	// 355
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Outlier
	//
	// example:
	//
	// -1
	Outlier *string `json:"outlier,omitempty" xml:"outlier,omitempty"`
	// Output results.
	//
	// example:
	//
	// DOUBLE
	Outputs *string `json:"outputs,omitempty" xml:"outputs,omitempty"`
	// Title.
	//
	// example:
	//
	// 自定义查询变量标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s DescribeQueryVariableDetailResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryVariableDetailResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) GetDataSourceCode() *int64 {
	return s.DataSourceCode
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) GetExpression() *string {
	return s.Expression
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) GetExpressionVariable() *string {
	return s.ExpressionVariable
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) GetOutlier() *string {
	return s.Outlier
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) GetOutputs() *string {
	return s.Outputs
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) GetTitle() *string {
	return s.Title
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) SetDataSourceCode(v int64) *DescribeQueryVariableDetailResponseBodyResultObject {
	s.DataSourceCode = &v
	return s
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) SetDataSourceName(v string) *DescribeQueryVariableDetailResponseBodyResultObject {
	s.DataSourceName = &v
	return s
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) SetDescription(v string) *DescribeQueryVariableDetailResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) SetEventCode(v string) *DescribeQueryVariableDetailResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) SetExpression(v string) *DescribeQueryVariableDetailResponseBodyResultObject {
	s.Expression = &v
	return s
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) SetExpressionTitle(v string) *DescribeQueryVariableDetailResponseBodyResultObject {
	s.ExpressionTitle = &v
	return s
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) SetExpressionVariable(v string) *DescribeQueryVariableDetailResponseBodyResultObject {
	s.ExpressionVariable = &v
	return s
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) SetId(v int64) *DescribeQueryVariableDetailResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) SetOutlier(v string) *DescribeQueryVariableDetailResponseBodyResultObject {
	s.Outlier = &v
	return s
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) SetOutputs(v string) *DescribeQueryVariableDetailResponseBodyResultObject {
	s.Outputs = &v
	return s
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) SetTitle(v string) *DescribeQueryVariableDetailResponseBodyResultObject {
	s.Title = &v
	return s
}

func (s *DescribeQueryVariableDetailResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
