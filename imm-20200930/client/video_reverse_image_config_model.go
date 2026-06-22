// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoReverseImageConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *VideoReverseImageConfig
	GetEnable() *bool
}

type VideoReverseImageConfig struct {
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s VideoReverseImageConfig) String() string {
	return dara.Prettify(s)
}

func (s VideoReverseImageConfig) GoString() string {
	return s.String()
}

func (s *VideoReverseImageConfig) GetEnable() *bool {
	return s.Enable
}

func (s *VideoReverseImageConfig) SetEnable(v bool) *VideoReverseImageConfig {
	s.Enable = &v
	return s
}

func (s *VideoReverseImageConfig) Validate() error {
	return dara.Validate(s)
}
