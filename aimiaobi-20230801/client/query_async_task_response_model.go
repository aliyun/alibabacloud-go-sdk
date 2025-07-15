// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAsyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAsyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAsyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *QueryAsyncTaskResponseBody) *QueryAsyncTaskResponse
	GetBody() *QueryAsyncTaskResponseBody
}

type QueryAsyncTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAsyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryAsyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAsyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAsyncTaskResponse) GetBody() *QueryAsyncTaskResponseBody {
	return s.Body
}

func (s *QueryAsyncTaskResponse) SetHeaders(v map[string]*string) *QueryAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryAsyncTaskResponse) SetStatusCode(v int32) *QueryAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAsyncTaskResponse) SetBody(v *QueryAsyncTaskResponseBody) *QueryAsyncTaskResponse {
	s.Body = v
	return s
}

func (s *QueryAsyncTaskResponse) Validate() error {
	return dara.Validate(s)
}
