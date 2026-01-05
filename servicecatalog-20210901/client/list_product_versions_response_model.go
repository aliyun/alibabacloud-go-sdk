// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProductVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProductVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListProductVersionsResponseBody) *ListProductVersionsResponse
	GetBody() *ListProductVersionsResponseBody
}

type ListProductVersionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProductVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListProductVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProductVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProductVersionsResponse) GetBody() *ListProductVersionsResponseBody {
	return s.Body
}

func (s *ListProductVersionsResponse) SetHeaders(v map[string]*string) *ListProductVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListProductVersionsResponse) SetStatusCode(v int32) *ListProductVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductVersionsResponse) SetBody(v *ListProductVersionsResponseBody) *ListProductVersionsResponse {
	s.Body = v
	return s
}

func (s *ListProductVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
