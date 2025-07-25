// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmRecoveryPlanAvailableConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeGtmRecoveryPlanAvailableConfigRequest
	GetLang() *string
}

type DescribeGtmRecoveryPlanAvailableConfigRequest struct {
	// The language in which the returned results are displayed. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeGtmRecoveryPlanAvailableConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlanAvailableConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanAvailableConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGtmRecoveryPlanAvailableConfigRequest) SetLang(v string) *DescribeGtmRecoveryPlanAvailableConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmRecoveryPlanAvailableConfigRequest) Validate() error {
	return dara.Validate(s)
}
