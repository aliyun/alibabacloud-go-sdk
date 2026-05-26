// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonReferenceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *PersonReferenceConfig
	GetEnable() *bool
}

type PersonReferenceConfig struct {
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s PersonReferenceConfig) String() string {
	return dara.Prettify(s)
}

func (s PersonReferenceConfig) GoString() string {
	return s.String()
}

func (s *PersonReferenceConfig) GetEnable() *bool {
	return s.Enable
}

func (s *PersonReferenceConfig) SetEnable(v bool) *PersonReferenceConfig {
	s.Enable = &v
	return s
}

func (s *PersonReferenceConfig) Validate() error {
	return dara.Validate(s)
}
