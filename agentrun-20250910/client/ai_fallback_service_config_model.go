// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiFallbackServiceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetPassThroughModelName(v bool) *AiFallbackServiceConfig
	GetPassThroughModelName() *bool
	SetServiceId(v string) *AiFallbackServiceConfig
	GetServiceId() *string
	SetTargetModelName(v string) *AiFallbackServiceConfig
	GetTargetModelName() *string
}

type AiFallbackServiceConfig struct {
	PassThroughModelName *bool   `json:"passThroughModelName,omitempty" xml:"passThroughModelName,omitempty"`
	ServiceId            *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	TargetModelName      *string `json:"targetModelName,omitempty" xml:"targetModelName,omitempty"`
}

func (s AiFallbackServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s AiFallbackServiceConfig) GoString() string {
	return s.String()
}

func (s *AiFallbackServiceConfig) GetPassThroughModelName() *bool {
	return s.PassThroughModelName
}

func (s *AiFallbackServiceConfig) GetServiceId() *string {
	return s.ServiceId
}

func (s *AiFallbackServiceConfig) GetTargetModelName() *string {
	return s.TargetModelName
}

func (s *AiFallbackServiceConfig) SetPassThroughModelName(v bool) *AiFallbackServiceConfig {
	s.PassThroughModelName = &v
	return s
}

func (s *AiFallbackServiceConfig) SetServiceId(v string) *AiFallbackServiceConfig {
	s.ServiceId = &v
	return s
}

func (s *AiFallbackServiceConfig) SetTargetModelName(v string) *AiFallbackServiceConfig {
	s.TargetModelName = &v
	return s
}

func (s *AiFallbackServiceConfig) Validate() error {
	return dara.Validate(s)
}
