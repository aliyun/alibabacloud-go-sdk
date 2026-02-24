// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleRcaConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDigitalEmployeeName(v string) *AlertRuleRcaConfig
	GetDigitalEmployeeName() *string
	SetEnableRca(v bool) *AlertRuleRcaConfig
	GetEnableRca() *bool
}

type AlertRuleRcaConfig struct {
	DigitalEmployeeName *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	EnableRca           *bool   `json:"enableRca,omitempty" xml:"enableRca,omitempty"`
}

func (s AlertRuleRcaConfig) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleRcaConfig) GoString() string {
	return s.String()
}

func (s *AlertRuleRcaConfig) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *AlertRuleRcaConfig) GetEnableRca() *bool {
	return s.EnableRca
}

func (s *AlertRuleRcaConfig) SetDigitalEmployeeName(v string) *AlertRuleRcaConfig {
	s.DigitalEmployeeName = &v
	return s
}

func (s *AlertRuleRcaConfig) SetEnableRca(v bool) *AlertRuleRcaConfig {
	s.EnableRca = &v
	return s
}

func (s *AlertRuleRcaConfig) Validate() error {
	return dara.Validate(s)
}
