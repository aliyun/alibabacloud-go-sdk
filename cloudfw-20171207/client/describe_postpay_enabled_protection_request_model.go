// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayEnabledProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribePostpayEnabledProtectionRequest
	GetLang() *string
}

type DescribePostpayEnabledProtectionRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribePostpayEnabledProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayEnabledProtectionRequest) GoString() string {
	return s.String()
}

func (s *DescribePostpayEnabledProtectionRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePostpayEnabledProtectionRequest) SetLang(v string) *DescribePostpayEnabledProtectionRequest {
	s.Lang = &v
	return s
}

func (s *DescribePostpayEnabledProtectionRequest) Validate() error {
	return dara.Validate(s)
}
