// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAvailableAlertGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDnsGtmAvailableAlertGroupRequest
	GetLang() *string
}

type DescribeDnsGtmAvailableAlertGroupRequest struct {
	// The language to return some response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeDnsGtmAvailableAlertGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAvailableAlertGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAvailableAlertGroupRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmAvailableAlertGroupRequest) SetLang(v string) *DescribeDnsGtmAvailableAlertGroupRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmAvailableAlertGroupRequest) Validate() error {
	return dara.Validate(s)
}
