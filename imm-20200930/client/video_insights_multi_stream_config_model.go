// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoInsightsMultiStreamConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *VideoInsightsMultiStreamConfig
	GetEnable() *bool
}

type VideoInsightsMultiStreamConfig struct {
	// Specifies whether video multi-stream is supported.
	//
	// example:
	//
	// false
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s VideoInsightsMultiStreamConfig) String() string {
	return dara.Prettify(s)
}

func (s VideoInsightsMultiStreamConfig) GoString() string {
	return s.String()
}

func (s *VideoInsightsMultiStreamConfig) GetEnable() *bool {
	return s.Enable
}

func (s *VideoInsightsMultiStreamConfig) SetEnable(v bool) *VideoInsightsMultiStreamConfig {
	s.Enable = &v
	return s
}

func (s *VideoInsightsMultiStreamConfig) Validate() error {
	return dara.Validate(s)
}
