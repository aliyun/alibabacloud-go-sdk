// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDnsCacheDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDnsCacheDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDnsCacheDomainResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDnsCacheDomainResponseBody) *DeleteDnsCacheDomainResponse
	GetBody() *DeleteDnsCacheDomainResponseBody
}

type DeleteDnsCacheDomainResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDnsCacheDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDnsCacheDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDnsCacheDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDnsCacheDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDnsCacheDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDnsCacheDomainResponse) GetBody() *DeleteDnsCacheDomainResponseBody {
	return s.Body
}

func (s *DeleteDnsCacheDomainResponse) SetHeaders(v map[string]*string) *DeleteDnsCacheDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDnsCacheDomainResponse) SetStatusCode(v int32) *DeleteDnsCacheDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDnsCacheDomainResponse) SetBody(v *DeleteDnsCacheDomainResponseBody) *DeleteDnsCacheDomainResponse {
	s.Body = v
	return s
}

func (s *DeleteDnsCacheDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
