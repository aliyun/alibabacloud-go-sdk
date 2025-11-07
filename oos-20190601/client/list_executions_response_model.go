// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExecutionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExecutionsResponse
	GetStatusCode() *int32
	SetBody(v *ListExecutionsResponseBody) *ListExecutionsResponse
	GetBody() *ListExecutionsResponseBody
}

type ListExecutionsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExecutionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExecutionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExecutionsResponse) GetBody() *ListExecutionsResponseBody {
	return s.Body
}

func (s *ListExecutionsResponse) SetHeaders(v map[string]*string) *ListExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListExecutionsResponse) SetStatusCode(v int32) *ListExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExecutionsResponse) SetBody(v *ListExecutionsResponseBody) *ListExecutionsResponse {
	s.Body = v
	return s
}

func (s *ListExecutionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
