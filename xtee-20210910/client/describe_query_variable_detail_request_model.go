// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQueryVariableDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeQueryVariableDetailRequest
	GetLang() *string
	SetId(v int64) *DescribeQueryVariableDetailRequest
	GetId() *int64
	SetRegId(v string) *DescribeQueryVariableDetailRequest
	GetRegId() *string
}

type DescribeQueryVariableDetailRequest struct {
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
	// Variable ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 355
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

func (s DescribeQueryVariableDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryVariableDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeQueryVariableDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeQueryVariableDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeQueryVariableDetailRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeQueryVariableDetailRequest) SetLang(v string) *DescribeQueryVariableDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeQueryVariableDetailRequest) SetId(v int64) *DescribeQueryVariableDetailRequest {
	s.Id = &v
	return s
}

func (s *DescribeQueryVariableDetailRequest) SetRegId(v string) *DescribeQueryVariableDetailRequest {
	s.RegId = &v
	return s
}

func (s *DescribeQueryVariableDetailRequest) Validate() error {
	return dara.Validate(s)
}
