// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperatorListByTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeOperatorListByTypeRequest
	GetLang() *string
	SetRegId(v string) *DescribeOperatorListByTypeRequest
	GetRegId() *string
}

type DescribeOperatorListByTypeRequest struct {
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
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeOperatorListByTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorListByTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeOperatorListByTypeRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOperatorListByTypeRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeOperatorListByTypeRequest) SetLang(v string) *DescribeOperatorListByTypeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOperatorListByTypeRequest) SetRegId(v string) *DescribeOperatorListByTypeRequest {
	s.RegId = &v
	return s
}

func (s *DescribeOperatorListByTypeRequest) Validate() error {
	return dara.Validate(s)
}
