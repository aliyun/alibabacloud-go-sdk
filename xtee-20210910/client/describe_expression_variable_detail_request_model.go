// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressionVariableDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeExpressionVariableDetailRequest
	GetLang() *string
	SetId(v int64) *DescribeExpressionVariableDetailRequest
	GetId() *int64
	SetRegId(v string) *DescribeExpressionVariableDetailRequest
	GetRegId() *string
}

type DescribeExpressionVariableDetailRequest struct {
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
	// Primary key ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 334
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeExpressionVariableDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressionVariableDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressionVariableDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeExpressionVariableDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeExpressionVariableDetailRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeExpressionVariableDetailRequest) SetLang(v string) *DescribeExpressionVariableDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeExpressionVariableDetailRequest) SetId(v int64) *DescribeExpressionVariableDetailRequest {
	s.Id = &v
	return s
}

func (s *DescribeExpressionVariableDetailRequest) SetRegId(v string) *DescribeExpressionVariableDetailRequest {
	s.RegId = &v
	return s
}

func (s *DescribeExpressionVariableDetailRequest) Validate() error {
	return dara.Validate(s)
}
