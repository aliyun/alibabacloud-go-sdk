// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDomainResponse
	GetStatusCode() *int32
	SetBody(v *ListDomainResponseBody) *ListDomainResponse
	GetBody() *ListDomainResponseBody
}

type ListDomainResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDomainResponse) GoString() string {
	return s.String()
}

func (s *ListDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDomainResponse) GetBody() *ListDomainResponseBody {
	return s.Body
}

func (s *ListDomainResponse) SetHeaders(v map[string]*string) *ListDomainResponse {
	s.Headers = v
	return s
}

func (s *ListDomainResponse) SetStatusCode(v int32) *ListDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDomainResponse) SetBody(v *ListDomainResponseBody) *ListDomainResponse {
	s.Body = v
	return s
}

func (s *ListDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
