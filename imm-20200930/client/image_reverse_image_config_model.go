// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageReverseImageConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *ImageReverseImageConfig
	GetEnable() *bool
}

type ImageReverseImageConfig struct {
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s ImageReverseImageConfig) String() string {
	return dara.Prettify(s)
}

func (s ImageReverseImageConfig) GoString() string {
	return s.String()
}

func (s *ImageReverseImageConfig) GetEnable() *bool {
	return s.Enable
}

func (s *ImageReverseImageConfig) SetEnable(v bool) *ImageReverseImageConfig {
	s.Enable = &v
	return s
}

func (s *ImageReverseImageConfig) Validate() error {
	return dara.Validate(s)
}
