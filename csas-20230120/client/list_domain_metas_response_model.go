// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainMetasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDomainMetasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDomainMetasResponse
	GetStatusCode() *int32
	SetBody(v *ListDomainMetasResponseBody) *ListDomainMetasResponse
	GetBody() *ListDomainMetasResponseBody
}

type ListDomainMetasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDomainMetasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDomainMetasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDomainMetasResponse) GoString() string {
	return s.String()
}

func (s *ListDomainMetasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDomainMetasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDomainMetasResponse) GetBody() *ListDomainMetasResponseBody {
	return s.Body
}

func (s *ListDomainMetasResponse) SetHeaders(v map[string]*string) *ListDomainMetasResponse {
	s.Headers = v
	return s
}

func (s *ListDomainMetasResponse) SetStatusCode(v int32) *ListDomainMetasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDomainMetasResponse) SetBody(v *ListDomainMetasResponseBody) *ListDomainMetasResponse {
	s.Body = v
	return s
}

func (s *ListDomainMetasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
