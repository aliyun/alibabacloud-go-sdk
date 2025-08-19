// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTracingConfig interface {
	dara.Model
	String() string
	GoString() string
	SetParams(v map[string]*string) *TracingConfig
	GetParams() map[string]*string
	SetType(v string) *TracingConfig
	GetType() *string
}

type TracingConfig struct {
	Params map[string]*string `json:"params" xml:"params"`
	// example:
	//
	// Jaeger
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TracingConfig) String() string {
	return dara.Prettify(s)
}

func (s TracingConfig) GoString() string {
	return s.String()
}

func (s *TracingConfig) GetParams() map[string]*string {
	return s.Params
}

func (s *TracingConfig) GetType() *string {
	return s.Type
}

func (s *TracingConfig) SetParams(v map[string]*string) *TracingConfig {
	s.Params = v
	return s
}

func (s *TracingConfig) SetType(v string) *TracingConfig {
	s.Type = &v
	return s
}

func (s *TracingConfig) Validate() error {
	return dara.Validate(s)
}
