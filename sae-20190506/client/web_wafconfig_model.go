// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebWAFConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnableWAF(v bool) *WebWAFConfig
	GetEnableWAF() *bool
}

type WebWAFConfig struct {
	EnableWAF *bool `json:"EnableWAF,omitempty" xml:"EnableWAF,omitempty"`
}

func (s WebWAFConfig) String() string {
	return dara.Prettify(s)
}

func (s WebWAFConfig) GoString() string {
	return s.String()
}

func (s *WebWAFConfig) GetEnableWAF() *bool {
	return s.EnableWAF
}

func (s *WebWAFConfig) SetEnableWAF(v bool) *WebWAFConfig {
	s.EnableWAF = &v
	return s
}

func (s *WebWAFConfig) Validate() error {
	return dara.Validate(s)
}
