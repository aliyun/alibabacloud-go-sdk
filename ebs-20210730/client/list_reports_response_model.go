// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListReportsResponse
	GetStatusCode() *int32
	SetBody(v *ListReportsResponseBody) *ListReportsResponse
	GetBody() *ListReportsResponseBody
}

type ListReportsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListReportsResponse) GoString() string {
	return s.String()
}

func (s *ListReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListReportsResponse) GetBody() *ListReportsResponseBody {
	return s.Body
}

func (s *ListReportsResponse) SetHeaders(v map[string]*string) *ListReportsResponse {
	s.Headers = v
	return s
}

func (s *ListReportsResponse) SetStatusCode(v int32) *ListReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListReportsResponse) SetBody(v *ListReportsResponseBody) *ListReportsResponse {
	s.Body = v
	return s
}

func (s *ListReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
