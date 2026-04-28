// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDomainResponse
	GetStatusCode() *int32
}

type DeleteDomainResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDomainResponse) SetHeaders(v map[string]*string) *DeleteDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainResponse) SetStatusCode(v int32) *DeleteDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainResponse) Validate() error {
	return dara.Validate(s)
}
