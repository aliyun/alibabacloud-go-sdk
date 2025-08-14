// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInitDigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeInitDigRequest
	GetLang() *string
	SetRegId(v string) *DescribeInitDigRequest
	GetRegId() *string
	SetType(v string) *DescribeInitDigRequest
	GetType() *string
}

type DescribeInitDigRequest struct {
	// Set the language type for request and response messages. Default value is **zh**. Values:
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
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// Data source type
	//
	// This parameter is required.
	//
	// example:
	//
	// SAF_CONSOLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeInitDigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInitDigRequest) GoString() string {
	return s.String()
}

func (s *DescribeInitDigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInitDigRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeInitDigRequest) GetType() *string {
	return s.Type
}

func (s *DescribeInitDigRequest) SetLang(v string) *DescribeInitDigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInitDigRequest) SetRegId(v string) *DescribeInitDigRequest {
	s.RegId = &v
	return s
}

func (s *DescribeInitDigRequest) SetType(v string) *DescribeInitDigRequest {
	s.Type = &v
	return s
}

func (s *DescribeInitDigRequest) Validate() error {
	return dara.Validate(s)
}
