// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHBaseQuotasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHBaseQuotasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHBaseQuotasResponse
	GetStatusCode() *int32
	SetBody(v *ListHBaseQuotasResponseBody) *ListHBaseQuotasResponse
	GetBody() *ListHBaseQuotasResponseBody
}

type ListHBaseQuotasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHBaseQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHBaseQuotasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHBaseQuotasResponse) GoString() string {
	return s.String()
}

func (s *ListHBaseQuotasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHBaseQuotasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHBaseQuotasResponse) GetBody() *ListHBaseQuotasResponseBody {
	return s.Body
}

func (s *ListHBaseQuotasResponse) SetHeaders(v map[string]*string) *ListHBaseQuotasResponse {
	s.Headers = v
	return s
}

func (s *ListHBaseQuotasResponse) SetStatusCode(v int32) *ListHBaseQuotasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHBaseQuotasResponse) SetBody(v *ListHBaseQuotasResponseBody) *ListHBaseQuotasResponse {
	s.Body = v
	return s
}

func (s *ListHBaseQuotasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
