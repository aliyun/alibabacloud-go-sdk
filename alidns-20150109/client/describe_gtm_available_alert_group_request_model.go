// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmAvailableAlertGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeGtmAvailableAlertGroupRequest
	GetLang() *string
}

type DescribeGtmAvailableAlertGroupRequest struct {
	// The language used by the user.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeGtmAvailableAlertGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAvailableAlertGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmAvailableAlertGroupRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGtmAvailableAlertGroupRequest) SetLang(v string) *DescribeGtmAvailableAlertGroupRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmAvailableAlertGroupRequest) Validate() error {
	return dara.Validate(s)
}
