// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalGuardForBase64Request interface {
	dara.Model
	String() string
	GoString() string
	SetImageBase64Str(v string) *MultiModalGuardForBase64Request
	GetImageBase64Str() *string
	SetService(v string) *MultiModalGuardForBase64Request
	GetService() *string
	SetServiceParameters(v string) *MultiModalGuardForBase64Request
	GetServiceParameters() *string
}

type MultiModalGuardForBase64Request struct {
	// example:
	//
	// {base64}
	ImageBase64Str *string `json:"ImageBase64Str,omitempty" xml:"ImageBase64Str,omitempty"`
	// Service
	//
	// example:
	//
	// query_security_check
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// example:
	//
	// {"content":"test"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s MultiModalGuardForBase64Request) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardForBase64Request) GoString() string {
	return s.String()
}

func (s *MultiModalGuardForBase64Request) GetImageBase64Str() *string {
	return s.ImageBase64Str
}

func (s *MultiModalGuardForBase64Request) GetService() *string {
	return s.Service
}

func (s *MultiModalGuardForBase64Request) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *MultiModalGuardForBase64Request) SetImageBase64Str(v string) *MultiModalGuardForBase64Request {
	s.ImageBase64Str = &v
	return s
}

func (s *MultiModalGuardForBase64Request) SetService(v string) *MultiModalGuardForBase64Request {
	s.Service = &v
	return s
}

func (s *MultiModalGuardForBase64Request) SetServiceParameters(v string) *MultiModalGuardForBase64Request {
	s.ServiceParameters = &v
	return s
}

func (s *MultiModalGuardForBase64Request) Validate() error {
	return dara.Validate(s)
}
