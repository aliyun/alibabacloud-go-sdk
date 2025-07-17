// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiVersionConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *HttpApiVersionConfig
	GetEnable() *bool
	SetHeaderName(v string) *HttpApiVersionConfig
	GetHeaderName() *string
	SetQueryName(v string) *HttpApiVersionConfig
	GetQueryName() *string
	SetScheme(v string) *HttpApiVersionConfig
	GetScheme() *string
	SetVersion(v string) *HttpApiVersionConfig
	GetVersion() *string
}

type HttpApiVersionConfig struct {
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// my-version
	HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty"`
	// example:
	//
	// myVersion
	QueryName *string `json:"queryName,omitempty" xml:"queryName,omitempty"`
	// example:
	//
	// Query
	Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s HttpApiVersionConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiVersionConfig) GoString() string {
	return s.String()
}

func (s *HttpApiVersionConfig) GetEnable() *bool {
	return s.Enable
}

func (s *HttpApiVersionConfig) GetHeaderName() *string {
	return s.HeaderName
}

func (s *HttpApiVersionConfig) GetQueryName() *string {
	return s.QueryName
}

func (s *HttpApiVersionConfig) GetScheme() *string {
	return s.Scheme
}

func (s *HttpApiVersionConfig) GetVersion() *string {
	return s.Version
}

func (s *HttpApiVersionConfig) SetEnable(v bool) *HttpApiVersionConfig {
	s.Enable = &v
	return s
}

func (s *HttpApiVersionConfig) SetHeaderName(v string) *HttpApiVersionConfig {
	s.HeaderName = &v
	return s
}

func (s *HttpApiVersionConfig) SetQueryName(v string) *HttpApiVersionConfig {
	s.QueryName = &v
	return s
}

func (s *HttpApiVersionConfig) SetScheme(v string) *HttpApiVersionConfig {
	s.Scheme = &v
	return s
}

func (s *HttpApiVersionConfig) SetVersion(v string) *HttpApiVersionConfig {
	s.Version = &v
	return s
}

func (s *HttpApiVersionConfig) Validate() error {
	return dara.Validate(s)
}
