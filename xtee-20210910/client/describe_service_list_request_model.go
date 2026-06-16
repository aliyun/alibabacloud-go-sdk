// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeServiceListRequest
	GetLang() *string
	SetRegId(v string) *DescribeServiceListRequest
	GetRegId() *string
}

type DescribeServiceListRequest struct {
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

func (s DescribeServiceListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeServiceListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeServiceListRequest) SetLang(v string) *DescribeServiceListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeServiceListRequest) SetRegId(v string) *DescribeServiceListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeServiceListRequest) Validate() error {
	return dara.Validate(s)
}
