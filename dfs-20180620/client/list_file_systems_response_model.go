// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileSystemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFileSystemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFileSystemsResponse
	GetStatusCode() *int32
	SetBody(v *ListFileSystemsResponseBody) *ListFileSystemsResponse
	GetBody() *ListFileSystemsResponseBody
}

type ListFileSystemsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFileSystemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *ListFileSystemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFileSystemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFileSystemsResponse) GetBody() *ListFileSystemsResponseBody {
	return s.Body
}

func (s *ListFileSystemsResponse) SetHeaders(v map[string]*string) *ListFileSystemsResponse {
	s.Headers = v
	return s
}

func (s *ListFileSystemsResponse) SetStatusCode(v int32) *ListFileSystemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFileSystemsResponse) SetBody(v *ListFileSystemsResponseBody) *ListFileSystemsResponse {
	s.Body = v
	return s
}

func (s *ListFileSystemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
