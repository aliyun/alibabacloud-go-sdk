// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExistNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeExistNameRequest
	GetLang() *string
	SetName(v string) *DescribeExistNameRequest
	GetName() *string
	SetRegId(v string) *DescribeExistNameRequest
	GetRegId() *string
}

type DescribeExistNameRequest struct {
	// Set the language type for request and response messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Variable name
	//
	// example:
	//
	// age
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeExistNameRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExistNameRequest) GoString() string {
	return s.String()
}

func (s *DescribeExistNameRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeExistNameRequest) GetName() *string {
	return s.Name
}

func (s *DescribeExistNameRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeExistNameRequest) SetLang(v string) *DescribeExistNameRequest {
	s.Lang = &v
	return s
}

func (s *DescribeExistNameRequest) SetName(v string) *DescribeExistNameRequest {
	s.Name = &v
	return s
}

func (s *DescribeExistNameRequest) SetRegId(v string) *DescribeExistNameRequest {
	s.RegId = &v
	return s
}

func (s *DescribeExistNameRequest) Validate() error {
	return dara.Validate(s)
}
