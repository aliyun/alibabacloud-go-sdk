// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsCacheDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDnsCacheDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDnsCacheDomainResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDnsCacheDomainResponseBody) *UpdateDnsCacheDomainResponse
	GetBody() *UpdateDnsCacheDomainResponseBody
}

type UpdateDnsCacheDomainResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDnsCacheDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDnsCacheDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsCacheDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateDnsCacheDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDnsCacheDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDnsCacheDomainResponse) GetBody() *UpdateDnsCacheDomainResponseBody {
	return s.Body
}

func (s *UpdateDnsCacheDomainResponse) SetHeaders(v map[string]*string) *UpdateDnsCacheDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateDnsCacheDomainResponse) SetStatusCode(v int32) *UpdateDnsCacheDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDnsCacheDomainResponse) SetBody(v *UpdateDnsCacheDomainResponseBody) *UpdateDnsCacheDomainResponse {
	s.Body = v
	return s
}

func (s *UpdateDnsCacheDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
