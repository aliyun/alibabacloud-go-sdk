// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFieldListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeFieldListRequest
	GetLang() *string
	SetCondition(v string) *DescribeFieldListRequest
	GetCondition() *string
	SetInputs(v string) *DescribeFieldListRequest
	GetInputs() *string
	SetRegId(v string) *DescribeFieldListRequest
	GetRegId() *string
}

type DescribeFieldListRequest struct {
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
	// Query input name or title
	//
	// example:
	//
	// ip
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// Selected fields
	//
	// example:
	//
	// sex,ip,id
	Inputs *string `json:"inputs,omitempty" xml:"inputs,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeFieldListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldListRequest) GoString() string {
	return s.String()
}

func (s *DescribeFieldListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeFieldListRequest) GetCondition() *string {
	return s.Condition
}

func (s *DescribeFieldListRequest) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeFieldListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeFieldListRequest) SetLang(v string) *DescribeFieldListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFieldListRequest) SetCondition(v string) *DescribeFieldListRequest {
	s.Condition = &v
	return s
}

func (s *DescribeFieldListRequest) SetInputs(v string) *DescribeFieldListRequest {
	s.Inputs = &v
	return s
}

func (s *DescribeFieldListRequest) SetRegId(v string) *DescribeFieldListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeFieldListRequest) Validate() error {
	return dara.Validate(s)
}
