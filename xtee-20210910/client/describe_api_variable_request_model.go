// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeApiVariableRequest
	GetLang() *string
	SetId(v string) *DescribeApiVariableRequest
	GetId() *string
	SetRegId(v string) *DescribeApiVariableRequest
	GetRegId() *string
}

type DescribeApiVariableRequest struct {
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
	// Variable ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 3144
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeApiVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiVariableRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeApiVariableRequest) GetId() *string {
	return s.Id
}

func (s *DescribeApiVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeApiVariableRequest) SetLang(v string) *DescribeApiVariableRequest {
	s.Lang = &v
	return s
}

func (s *DescribeApiVariableRequest) SetId(v string) *DescribeApiVariableRequest {
	s.Id = &v
	return s
}

func (s *DescribeApiVariableRequest) SetRegId(v string) *DescribeApiVariableRequest {
	s.RegId = &v
	return s
}

func (s *DescribeApiVariableRequest) Validate() error {
	return dara.Validate(s)
}
