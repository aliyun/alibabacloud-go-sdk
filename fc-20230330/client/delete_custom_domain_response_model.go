// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomDomainResponse
	GetStatusCode() *int32
}

type DeleteCustomDomainResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteCustomDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomDomainResponse) SetHeaders(v map[string]*string) *DeleteCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomDomainResponse) SetStatusCode(v int32) *DeleteCustomDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomDomainResponse) Validate() error {
	return dara.Validate(s)
}
