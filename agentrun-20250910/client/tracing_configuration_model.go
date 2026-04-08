// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTracingConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *TracingConfiguration
	GetEnabled() *bool
}

type TracingConfiguration struct {
	// 是否启用链路追踪功能
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s TracingConfiguration) String() string {
	return dara.Prettify(s)
}

func (s TracingConfiguration) GoString() string {
	return s.String()
}

func (s *TracingConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *TracingConfiguration) SetEnabled(v bool) *TracingConfiguration {
	s.Enabled = &v
	return s
}

func (s *TracingConfiguration) Validate() error {
	return dara.Validate(s)
}
