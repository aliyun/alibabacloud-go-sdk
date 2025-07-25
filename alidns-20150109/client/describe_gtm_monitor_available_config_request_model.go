// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmMonitorAvailableConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeGtmMonitorAvailableConfigRequest
	GetLang() *string
}

type DescribeGtmMonitorAvailableConfigRequest struct {
	// The language of the values of specific response parameters.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeGtmMonitorAvailableConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmMonitorAvailableConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorAvailableConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGtmMonitorAvailableConfigRequest) SetLang(v string) *DescribeGtmMonitorAvailableConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmMonitorAvailableConfigRequest) Validate() error {
	return dara.Validate(s)
}
