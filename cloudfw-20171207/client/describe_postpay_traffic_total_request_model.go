// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayTrafficTotalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribePostpayTrafficTotalRequest
	GetLang() *string
}

type DescribePostpayTrafficTotalRequest struct {
	// The language of the content within the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribePostpayTrafficTotalRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayTrafficTotalRequest) GoString() string {
	return s.String()
}

func (s *DescribePostpayTrafficTotalRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePostpayTrafficTotalRequest) SetLang(v string) *DescribePostpayTrafficTotalRequest {
	s.Lang = &v
	return s
}

func (s *DescribePostpayTrafficTotalRequest) Validate() error {
	return dara.Validate(s)
}
