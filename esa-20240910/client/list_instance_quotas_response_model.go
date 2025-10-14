// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceQuotasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceQuotasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceQuotasResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceQuotasResponseBody) *ListInstanceQuotasResponse
	GetBody() *ListInstanceQuotasResponseBody
}

type ListInstanceQuotasResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceQuotasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceQuotasResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceQuotasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceQuotasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceQuotasResponse) GetBody() *ListInstanceQuotasResponseBody {
	return s.Body
}

func (s *ListInstanceQuotasResponse) SetHeaders(v map[string]*string) *ListInstanceQuotasResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceQuotasResponse) SetStatusCode(v int32) *ListInstanceQuotasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceQuotasResponse) SetBody(v *ListInstanceQuotasResponseBody) *ListInstanceQuotasResponse {
	s.Body = v
	return s
}

func (s *ListInstanceQuotasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
