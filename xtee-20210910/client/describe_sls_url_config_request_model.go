// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsUrlConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSlsUrlConfigRequest
	GetLang() *string
	SetRegId(v string) *DescribeSlsUrlConfigRequest
	GetRegId() *string
}

type DescribeSlsUrlConfigRequest struct {
	// The language type for the request and response messages. Default value: **zh**. Valid values:
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

func (s DescribeSlsUrlConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsUrlConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsUrlConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSlsUrlConfigRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSlsUrlConfigRequest) SetLang(v string) *DescribeSlsUrlConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsUrlConfigRequest) SetRegId(v string) *DescribeSlsUrlConfigRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSlsUrlConfigRequest) Validate() error {
	return dara.Validate(s)
}
