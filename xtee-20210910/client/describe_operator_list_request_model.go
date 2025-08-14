// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperatorListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeOperatorListRequest
	GetLang() *string
	SetRegId(v string) *DescribeOperatorListRequest
	GetRegId() *string
}

type DescribeOperatorListRequest struct {
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
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeOperatorListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorListRequest) GoString() string {
	return s.String()
}

func (s *DescribeOperatorListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOperatorListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeOperatorListRequest) SetLang(v string) *DescribeOperatorListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOperatorListRequest) SetRegId(v string) *DescribeOperatorListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeOperatorListRequest) Validate() error {
	return dara.Validate(s)
}
