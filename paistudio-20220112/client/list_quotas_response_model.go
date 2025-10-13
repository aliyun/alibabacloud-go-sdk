// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQuotasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQuotasResponse
	GetStatusCode() *int32
	SetBody(v *ListQuotasResponseBody) *ListQuotasResponse
	GetBody() *ListQuotasResponseBody
}

type ListQuotasResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponse) GoString() string {
	return s.String()
}

func (s *ListQuotasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQuotasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQuotasResponse) GetBody() *ListQuotasResponseBody {
	return s.Body
}

func (s *ListQuotasResponse) SetHeaders(v map[string]*string) *ListQuotasResponse {
	s.Headers = v
	return s
}

func (s *ListQuotasResponse) SetStatusCode(v int32) *ListQuotasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotasResponse) SetBody(v *ListQuotasResponseBody) *ListQuotasResponse {
	s.Body = v
	return s
}

func (s *ListQuotasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
