// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnsCacheDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDnsCacheDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDnsCacheDomainResponse
	GetStatusCode() *int32
	SetBody(v *AddDnsCacheDomainResponseBody) *AddDnsCacheDomainResponse
	GetBody() *AddDnsCacheDomainResponseBody
}

type AddDnsCacheDomainResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDnsCacheDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDnsCacheDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDnsCacheDomainResponse) GoString() string {
	return s.String()
}

func (s *AddDnsCacheDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDnsCacheDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDnsCacheDomainResponse) GetBody() *AddDnsCacheDomainResponseBody {
	return s.Body
}

func (s *AddDnsCacheDomainResponse) SetHeaders(v map[string]*string) *AddDnsCacheDomainResponse {
	s.Headers = v
	return s
}

func (s *AddDnsCacheDomainResponse) SetStatusCode(v int32) *AddDnsCacheDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDnsCacheDomainResponse) SetBody(v *AddDnsCacheDomainResponseBody) *AddDnsCacheDomainResponse {
	s.Body = v
	return s
}

func (s *AddDnsCacheDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
