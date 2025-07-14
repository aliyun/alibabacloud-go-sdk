// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWAFConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnableWAF(v bool) *WAFConfig
	GetEnableWAF() *bool
}

type WAFConfig struct {
	EnableWAF *bool `json:"enableWAF,omitempty" xml:"enableWAF,omitempty"`
}

func (s WAFConfig) String() string {
	return dara.Prettify(s)
}

func (s WAFConfig) GoString() string {
	return s.String()
}

func (s *WAFConfig) GetEnableWAF() *bool {
	return s.EnableWAF
}

func (s *WAFConfig) SetEnableWAF(v bool) *WAFConfig {
	s.EnableWAF = &v
	return s
}

func (s *WAFConfig) Validate() error {
	return dara.Validate(s)
}
