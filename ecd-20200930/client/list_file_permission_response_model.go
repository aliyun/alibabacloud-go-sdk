// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFilePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFilePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFilePermissionResponse
	GetStatusCode() *int32
	SetBody(v *ListFilePermissionResponseBody) *ListFilePermissionResponse
	GetBody() *ListFilePermissionResponseBody
}

type ListFilePermissionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFilePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFilePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFilePermissionResponse) GoString() string {
	return s.String()
}

func (s *ListFilePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFilePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFilePermissionResponse) GetBody() *ListFilePermissionResponseBody {
	return s.Body
}

func (s *ListFilePermissionResponse) SetHeaders(v map[string]*string) *ListFilePermissionResponse {
	s.Headers = v
	return s
}

func (s *ListFilePermissionResponse) SetStatusCode(v int32) *ListFilePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFilePermissionResponse) SetBody(v *ListFilePermissionResponseBody) *ListFilePermissionResponse {
	s.Body = v
	return s
}

func (s *ListFilePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
