// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedMaxPriority(v int32) *QuotaConfig
	GetAllowedMaxPriority() *int32
	SetEnableDLC(v bool) *QuotaConfig
	GetEnableDLC() *bool
	SetEnableDSW(v bool) *QuotaConfig
	GetEnableDSW() *bool
	SetEnableTideResource(v bool) *QuotaConfig
	GetEnableTideResource() *bool
	SetResourceLevel(v string) *QuotaConfig
	GetResourceLevel() *string
}

type QuotaConfig struct {
	AllowedMaxPriority *int32  `json:"AllowedMaxPriority,omitempty" xml:"AllowedMaxPriority,omitempty"`
	EnableDLC          *bool   `json:"EnableDLC,omitempty" xml:"EnableDLC,omitempty"`
	EnableDSW          *bool   `json:"EnableDSW,omitempty" xml:"EnableDSW,omitempty"`
	EnableTideResource *bool   `json:"EnableTideResource,omitempty" xml:"EnableTideResource,omitempty"`
	ResourceLevel      *string `json:"ResourceLevel,omitempty" xml:"ResourceLevel,omitempty"`
}

func (s QuotaConfig) String() string {
	return dara.Prettify(s)
}

func (s QuotaConfig) GoString() string {
	return s.String()
}

func (s *QuotaConfig) GetAllowedMaxPriority() *int32 {
	return s.AllowedMaxPriority
}

func (s *QuotaConfig) GetEnableDLC() *bool {
	return s.EnableDLC
}

func (s *QuotaConfig) GetEnableDSW() *bool {
	return s.EnableDSW
}

func (s *QuotaConfig) GetEnableTideResource() *bool {
	return s.EnableTideResource
}

func (s *QuotaConfig) GetResourceLevel() *string {
	return s.ResourceLevel
}

func (s *QuotaConfig) SetAllowedMaxPriority(v int32) *QuotaConfig {
	s.AllowedMaxPriority = &v
	return s
}

func (s *QuotaConfig) SetEnableDLC(v bool) *QuotaConfig {
	s.EnableDLC = &v
	return s
}

func (s *QuotaConfig) SetEnableDSW(v bool) *QuotaConfig {
	s.EnableDSW = &v
	return s
}

func (s *QuotaConfig) SetEnableTideResource(v bool) *QuotaConfig {
	s.EnableTideResource = &v
	return s
}

func (s *QuotaConfig) SetResourceLevel(v string) *QuotaConfig {
	s.ResourceLevel = &v
	return s
}

func (s *QuotaConfig) Validate() error {
	return dara.Validate(s)
}
