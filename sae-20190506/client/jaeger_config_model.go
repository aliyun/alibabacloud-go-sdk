// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJaegerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoint(v string) *JaegerConfig
	GetEndpoint() *string
}

type JaegerConfig struct {
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
}

func (s JaegerConfig) String() string {
	return dara.Prettify(s)
}

func (s JaegerConfig) GoString() string {
	return s.String()
}

func (s *JaegerConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *JaegerConfig) SetEndpoint(v string) *JaegerConfig {
	s.Endpoint = &v
	return s
}

func (s *JaegerConfig) Validate() error {
	return dara.Validate(s)
}
