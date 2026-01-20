// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOversoldUsageConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDisabled(v bool) *OversoldUsageConfig
	GetDisabled() *bool
	SetDisabledBy(v string) *OversoldUsageConfig
	GetDisabledBy() *string
}

type OversoldUsageConfig struct {
	Disabled   *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	DisabledBy *string `json:"DisabledBy,omitempty" xml:"DisabledBy,omitempty"`
}

func (s OversoldUsageConfig) String() string {
	return dara.Prettify(s)
}

func (s OversoldUsageConfig) GoString() string {
	return s.String()
}

func (s *OversoldUsageConfig) GetDisabled() *bool {
	return s.Disabled
}

func (s *OversoldUsageConfig) GetDisabledBy() *string {
	return s.DisabledBy
}

func (s *OversoldUsageConfig) SetDisabled(v bool) *OversoldUsageConfig {
	s.Disabled = &v
	return s
}

func (s *OversoldUsageConfig) SetDisabledBy(v string) *OversoldUsageConfig {
	s.DisabledBy = &v
	return s
}

func (s *OversoldUsageConfig) Validate() error {
	return dara.Validate(s)
}
