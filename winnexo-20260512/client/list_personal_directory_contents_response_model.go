// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPersonalDirectoryContentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPersonalDirectoryContentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPersonalDirectoryContentsResponse
	GetStatusCode() *int32
	SetBody(v *ListPersonalDirectoryContentsResponseBody) *ListPersonalDirectoryContentsResponse
	GetBody() *ListPersonalDirectoryContentsResponseBody
}

type ListPersonalDirectoryContentsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPersonalDirectoryContentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPersonalDirectoryContentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPersonalDirectoryContentsResponse) GoString() string {
	return s.String()
}

func (s *ListPersonalDirectoryContentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPersonalDirectoryContentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPersonalDirectoryContentsResponse) GetBody() *ListPersonalDirectoryContentsResponseBody {
	return s.Body
}

func (s *ListPersonalDirectoryContentsResponse) SetHeaders(v map[string]*string) *ListPersonalDirectoryContentsResponse {
	s.Headers = v
	return s
}

func (s *ListPersonalDirectoryContentsResponse) SetStatusCode(v int32) *ListPersonalDirectoryContentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPersonalDirectoryContentsResponse) SetBody(v *ListPersonalDirectoryContentsResponseBody) *ListPersonalDirectoryContentsResponse {
	s.Body = v
	return s
}

func (s *ListPersonalDirectoryContentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
