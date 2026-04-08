// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFileVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFileVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListFileVersionsResponseBody) *ListFileVersionsResponse
	GetBody() *ListFileVersionsResponseBody
}

type ListFileVersionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFileVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFileVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFileVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListFileVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFileVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFileVersionsResponse) GetBody() *ListFileVersionsResponseBody {
	return s.Body
}

func (s *ListFileVersionsResponse) SetHeaders(v map[string]*string) *ListFileVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListFileVersionsResponse) SetStatusCode(v int32) *ListFileVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFileVersionsResponse) SetBody(v *ListFileVersionsResponseBody) *ListFileVersionsResponse {
	s.Body = v
	return s
}

func (s *ListFileVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
