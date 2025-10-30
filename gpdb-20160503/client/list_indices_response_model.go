// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIndicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIndicesResponse
	GetStatusCode() *int32
	SetBody(v *ListIndicesResponseBody) *ListIndicesResponse
	GetBody() *ListIndicesResponseBody
}

type ListIndicesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIndicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIndicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIndicesResponse) GoString() string {
	return s.String()
}

func (s *ListIndicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIndicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIndicesResponse) GetBody() *ListIndicesResponseBody {
	return s.Body
}

func (s *ListIndicesResponse) SetHeaders(v map[string]*string) *ListIndicesResponse {
	s.Headers = v
	return s
}

func (s *ListIndicesResponse) SetStatusCode(v int32) *ListIndicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIndicesResponse) SetBody(v *ListIndicesResponseBody) *ListIndicesResponse {
	s.Body = v
	return s
}

func (s *ListIndicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
