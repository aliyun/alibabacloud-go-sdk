// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCORSConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAllowCredentials(v bool) *CORSConfig
	GetAllowCredentials() *bool
	SetAllowHeaders(v []*string) *CORSConfig
	GetAllowHeaders() []*string
	SetAllowMethods(v []*string) *CORSConfig
	GetAllowMethods() []*string
	SetAllowOrigins(v []*string) *CORSConfig
	GetAllowOrigins() []*string
	SetExposeHeaders(v []*string) *CORSConfig
	GetExposeHeaders() []*string
	SetMaxAge(v int32) *CORSConfig
	GetMaxAge() *int32
}

type CORSConfig struct {
	// example:
	//
	// true
	AllowCredentials *bool `json:"allowCredentials,omitempty" xml:"allowCredentials,omitempty"`
	// example:
	//
	// ["Content-Type", "Authorization"]
	AllowHeaders []*string `json:"allowHeaders" xml:"allowHeaders" type:"Repeated"`
	// example:
	//
	// ["GET", "POST", "OPTIONS"]
	AllowMethods []*string `json:"allowMethods" xml:"allowMethods" type:"Repeated"`
	// example:
	//
	// ["https://example.com", "https://app.example.com"]
	AllowOrigins []*string `json:"allowOrigins" xml:"allowOrigins" type:"Repeated"`
	// example:
	//
	// ["X-Custom-Header"]
	ExposeHeaders []*string `json:"exposeHeaders" xml:"exposeHeaders" type:"Repeated"`
	// example:
	//
	// 3600
	MaxAge *int32 `json:"maxAge,omitempty" xml:"maxAge,omitempty"`
}

func (s CORSConfig) String() string {
	return dara.Prettify(s)
}

func (s CORSConfig) GoString() string {
	return s.String()
}

func (s *CORSConfig) GetAllowCredentials() *bool {
	return s.AllowCredentials
}

func (s *CORSConfig) GetAllowHeaders() []*string {
	return s.AllowHeaders
}

func (s *CORSConfig) GetAllowMethods() []*string {
	return s.AllowMethods
}

func (s *CORSConfig) GetAllowOrigins() []*string {
	return s.AllowOrigins
}

func (s *CORSConfig) GetExposeHeaders() []*string {
	return s.ExposeHeaders
}

func (s *CORSConfig) GetMaxAge() *int32 {
	return s.MaxAge
}

func (s *CORSConfig) SetAllowCredentials(v bool) *CORSConfig {
	s.AllowCredentials = &v
	return s
}

func (s *CORSConfig) SetAllowHeaders(v []*string) *CORSConfig {
	s.AllowHeaders = v
	return s
}

func (s *CORSConfig) SetAllowMethods(v []*string) *CORSConfig {
	s.AllowMethods = v
	return s
}

func (s *CORSConfig) SetAllowOrigins(v []*string) *CORSConfig {
	s.AllowOrigins = v
	return s
}

func (s *CORSConfig) SetExposeHeaders(v []*string) *CORSConfig {
	s.ExposeHeaders = v
	return s
}

func (s *CORSConfig) SetMaxAge(v int32) *CORSConfig {
	s.MaxAge = &v
	return s
}

func (s *CORSConfig) Validate() error {
	return dara.Validate(s)
}
