// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDownloadTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDownloadTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDownloadTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListDownloadTasksResponseBody) *ListDownloadTasksResponse
	GetBody() *ListDownloadTasksResponseBody
}

type ListDownloadTasksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDownloadTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDownloadTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDownloadTasksResponse) GoString() string {
	return s.String()
}

func (s *ListDownloadTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDownloadTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDownloadTasksResponse) GetBody() *ListDownloadTasksResponseBody {
	return s.Body
}

func (s *ListDownloadTasksResponse) SetHeaders(v map[string]*string) *ListDownloadTasksResponse {
	s.Headers = v
	return s
}

func (s *ListDownloadTasksResponse) SetStatusCode(v int32) *ListDownloadTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDownloadTasksResponse) SetBody(v *ListDownloadTasksResponseBody) *ListDownloadTasksResponse {
	s.Body = v
	return s
}

func (s *ListDownloadTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
