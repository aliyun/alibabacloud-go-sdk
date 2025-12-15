// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableFileSystemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAvailableFileSystemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAvailableFileSystemsResponse
	GetStatusCode() *int32
	SetBody(v *ListAvailableFileSystemsResponseBody) *ListAvailableFileSystemsResponse
	GetBody() *ListAvailableFileSystemsResponseBody
}

type ListAvailableFileSystemsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvailableFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvailableFileSystemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableFileSystemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAvailableFileSystemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAvailableFileSystemsResponse) GetBody() *ListAvailableFileSystemsResponseBody {
	return s.Body
}

func (s *ListAvailableFileSystemsResponse) SetHeaders(v map[string]*string) *ListAvailableFileSystemsResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableFileSystemsResponse) SetStatusCode(v int32) *ListAvailableFileSystemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableFileSystemsResponse) SetBody(v *ListAvailableFileSystemsResponseBody) *ListAvailableFileSystemsResponse {
	s.Body = v
	return s
}

func (s *ListAvailableFileSystemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
