// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsCacheDomainRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDnsCacheDomainRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDnsCacheDomainRemarkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDnsCacheDomainRemarkResponseBody) *UpdateDnsCacheDomainRemarkResponse
	GetBody() *UpdateDnsCacheDomainRemarkResponseBody
}

type UpdateDnsCacheDomainRemarkResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDnsCacheDomainRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDnsCacheDomainRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsCacheDomainRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateDnsCacheDomainRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDnsCacheDomainRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDnsCacheDomainRemarkResponse) GetBody() *UpdateDnsCacheDomainRemarkResponseBody {
	return s.Body
}

func (s *UpdateDnsCacheDomainRemarkResponse) SetHeaders(v map[string]*string) *UpdateDnsCacheDomainRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateDnsCacheDomainRemarkResponse) SetStatusCode(v int32) *UpdateDnsCacheDomainRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDnsCacheDomainRemarkResponse) SetBody(v *UpdateDnsCacheDomainRemarkResponseBody) *UpdateDnsCacheDomainRemarkResponse {
	s.Body = v
	return s
}

func (s *UpdateDnsCacheDomainRemarkResponse) Validate() error {
	return dara.Validate(s)
}
