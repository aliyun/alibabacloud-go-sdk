// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeApiLimitRequest
	GetLang() *string
	SetRegId(v string) *DescribeApiLimitRequest
	GetRegId() *string
}

type DescribeApiLimitRequest struct {
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
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeApiLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiLimitRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiLimitRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeApiLimitRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeApiLimitRequest) SetLang(v string) *DescribeApiLimitRequest {
	s.Lang = &v
	return s
}

func (s *DescribeApiLimitRequest) SetRegId(v string) *DescribeApiLimitRequest {
	s.RegId = &v
	return s
}

func (s *DescribeApiLimitRequest) Validate() error {
	return dara.Validate(s)
}
