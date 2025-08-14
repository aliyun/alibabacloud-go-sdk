// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVariableDetailRequest
	GetLang() *string
	SetId(v int64) *DescribeVariableDetailRequest
	GetId() *int64
	SetRegId(v string) *DescribeVariableDetailRequest
	GetRegId() *string
}

type DescribeVariableDetailRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
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
	// 3144
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeVariableDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVariableDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVariableDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeVariableDetailRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeVariableDetailRequest) SetLang(v string) *DescribeVariableDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVariableDetailRequest) SetId(v int64) *DescribeVariableDetailRequest {
	s.Id = &v
	return s
}

func (s *DescribeVariableDetailRequest) SetRegId(v string) *DescribeVariableDetailRequest {
	s.RegId = &v
	return s
}

func (s *DescribeVariableDetailRequest) Validate() error {
	return dara.Validate(s)
}
