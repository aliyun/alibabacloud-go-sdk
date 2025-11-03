// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceHealthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceHealthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceHealthResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceHealthResponseBody) *ListInstanceHealthResponse
	GetBody() *ListInstanceHealthResponseBody
}

type ListInstanceHealthResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceHealthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceHealthResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceHealthResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceHealthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceHealthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceHealthResponse) GetBody() *ListInstanceHealthResponseBody {
	return s.Body
}

func (s *ListInstanceHealthResponse) SetHeaders(v map[string]*string) *ListInstanceHealthResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceHealthResponse) SetStatusCode(v int32) *ListInstanceHealthResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceHealthResponse) SetBody(v *ListInstanceHealthResponseBody) *ListInstanceHealthResponse {
	s.Body = v
	return s
}

func (s *ListInstanceHealthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
