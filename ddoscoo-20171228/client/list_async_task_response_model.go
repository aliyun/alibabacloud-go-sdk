// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAsyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAsyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListAsyncTaskResponseBody) *ListAsyncTaskResponse
	GetBody() *ListAsyncTaskResponseBody
}

type ListAsyncTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAsyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *ListAsyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAsyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAsyncTaskResponse) GetBody() *ListAsyncTaskResponseBody {
	return s.Body
}

func (s *ListAsyncTaskResponse) SetHeaders(v map[string]*string) *ListAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *ListAsyncTaskResponse) SetStatusCode(v int32) *ListAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAsyncTaskResponse) SetBody(v *ListAsyncTaskResponseBody) *ListAsyncTaskResponse {
	s.Body = v
	return s
}

func (s *ListAsyncTaskResponse) Validate() error {
	return dara.Validate(s)
}
