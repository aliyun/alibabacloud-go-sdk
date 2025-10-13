// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTracingConfig interface {
	dara.Model
	String() string
	GoString() string
	SetJaegerConfig(v *JaegerConfig) *TracingConfig
	GetJaegerConfig() *JaegerConfig
	SetParams(v interface{}) *TracingConfig
	GetParams() interface{}
	SetType(v string) *TracingConfig
	GetType() *string
}

type TracingConfig struct {
	JaegerConfig *JaegerConfig `json:"jaegerConfig,omitempty" xml:"jaegerConfig,omitempty"`
	Params       interface{}   `json:"params,omitempty" xml:"params,omitempty"`
	Type         *string       `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TracingConfig) String() string {
	return dara.Prettify(s)
}

func (s TracingConfig) GoString() string {
	return s.String()
}

func (s *TracingConfig) GetJaegerConfig() *JaegerConfig {
	return s.JaegerConfig
}

func (s *TracingConfig) GetParams() interface{} {
	return s.Params
}

func (s *TracingConfig) GetType() *string {
	return s.Type
}

func (s *TracingConfig) SetJaegerConfig(v *JaegerConfig) *TracingConfig {
	s.JaegerConfig = v
	return s
}

func (s *TracingConfig) SetParams(v interface{}) *TracingConfig {
	s.Params = v
	return s
}

func (s *TracingConfig) SetType(v string) *TracingConfig {
	s.Type = &v
	return s
}

func (s *TracingConfig) Validate() error {
	if s.JaegerConfig != nil {
		if err := s.JaegerConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
