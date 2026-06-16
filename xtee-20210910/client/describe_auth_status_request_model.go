// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAuthStatusRequest
	GetLang() *string
	SetRegId(v string) *DescribeAuthStatusRequest
	GetRegId() *string
}

type DescribeAuthStatusRequest struct {
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeAuthStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuthStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAuthStatusRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAuthStatusRequest) SetLang(v string) *DescribeAuthStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAuthStatusRequest) SetRegId(v string) *DescribeAuthStatusRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAuthStatusRequest) Validate() error {
	return dara.Validate(s)
}
