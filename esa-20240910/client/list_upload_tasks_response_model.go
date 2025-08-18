// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUploadTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUploadTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUploadTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListUploadTasksResponseBody) *ListUploadTasksResponse
	GetBody() *ListUploadTasksResponseBody
}

type ListUploadTasksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUploadTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUploadTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUploadTasksResponse) GoString() string {
	return s.String()
}

func (s *ListUploadTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUploadTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUploadTasksResponse) GetBody() *ListUploadTasksResponseBody {
	return s.Body
}

func (s *ListUploadTasksResponse) SetHeaders(v map[string]*string) *ListUploadTasksResponse {
	s.Headers = v
	return s
}

func (s *ListUploadTasksResponse) SetStatusCode(v int32) *ListUploadTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUploadTasksResponse) SetBody(v *ListUploadTasksResponseBody) *ListUploadTasksResponse {
	s.Body = v
	return s
}

func (s *ListUploadTasksResponse) Validate() error {
	return dara.Validate(s)
}
