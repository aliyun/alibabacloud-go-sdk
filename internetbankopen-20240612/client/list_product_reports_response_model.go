// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProductReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProductReportsResponse
	GetStatusCode() *int32
	SetBody(v *ListProductReportsResponseBody) *ListProductReportsResponse
	GetBody() *ListProductReportsResponseBody
}

type ListProductReportsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProductReportsResponse) GoString() string {
	return s.String()
}

func (s *ListProductReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProductReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProductReportsResponse) GetBody() *ListProductReportsResponseBody {
	return s.Body
}

func (s *ListProductReportsResponse) SetHeaders(v map[string]*string) *ListProductReportsResponse {
	s.Headers = v
	return s
}

func (s *ListProductReportsResponse) SetStatusCode(v int32) *ListProductReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductReportsResponse) SetBody(v *ListProductReportsResponseBody) *ListProductReportsResponse {
	s.Body = v
	return s
}

func (s *ListProductReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
