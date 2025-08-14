// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressionVariableFunctionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeExpressionVariableFunctionListRequest
	GetLang() *string
	SetRegId(v string) *DescribeExpressionVariableFunctionListRequest
	GetRegId() *string
}

type DescribeExpressionVariableFunctionListRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeExpressionVariableFunctionListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressionVariableFunctionListRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressionVariableFunctionListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeExpressionVariableFunctionListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeExpressionVariableFunctionListRequest) SetLang(v string) *DescribeExpressionVariableFunctionListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeExpressionVariableFunctionListRequest) SetRegId(v string) *DescribeExpressionVariableFunctionListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeExpressionVariableFunctionListRequest) Validate() error {
	return dara.Validate(s)
}
