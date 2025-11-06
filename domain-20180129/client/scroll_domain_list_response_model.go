// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScrollDomainListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScrollDomainListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScrollDomainListResponse
	GetStatusCode() *int32
	SetBody(v *ScrollDomainListResponseBody) *ScrollDomainListResponse
	GetBody() *ScrollDomainListResponseBody
}

type ScrollDomainListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScrollDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScrollDomainListResponse) String() string {
	return dara.Prettify(s)
}

func (s ScrollDomainListResponse) GoString() string {
	return s.String()
}

func (s *ScrollDomainListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScrollDomainListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScrollDomainListResponse) GetBody() *ScrollDomainListResponseBody {
	return s.Body
}

func (s *ScrollDomainListResponse) SetHeaders(v map[string]*string) *ScrollDomainListResponse {
	s.Headers = v
	return s
}

func (s *ScrollDomainListResponse) SetStatusCode(v int32) *ScrollDomainListResponse {
	s.StatusCode = &v
	return s
}

func (s *ScrollDomainListResponse) SetBody(v *ScrollDomainListResponseBody) *ScrollDomainListResponse {
	s.Body = v
	return s
}

func (s *ScrollDomainListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
