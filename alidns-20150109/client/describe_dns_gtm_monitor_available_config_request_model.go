// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmMonitorAvailableConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDnsGtmMonitorAvailableConfigRequest
	GetLang() *string
}

type DescribeDnsGtmMonitorAvailableConfigRequest struct {
	// The language of the values of specific response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeDnsGtmMonitorAvailableConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorAvailableConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorAvailableConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmMonitorAvailableConfigRequest) SetLang(v string) *DescribeDnsGtmMonitorAvailableConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmMonitorAvailableConfigRequest) Validate() error {
	return dara.Validate(s)
}
