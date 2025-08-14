// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeUserInfoRequest
	GetLang() *string
	SetRegId(v string) *DescribeUserInfoRequest
	GetRegId() *string
}

type DescribeUserInfoRequest struct {
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
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeUserInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeUserInfoRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeUserInfoRequest) SetLang(v string) *DescribeUserInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUserInfoRequest) SetRegId(v string) *DescribeUserInfoRequest {
	s.RegId = &v
	return s
}

func (s *DescribeUserInfoRequest) Validate() error {
	return dara.Validate(s)
}
