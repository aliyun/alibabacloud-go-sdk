// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAsyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAsyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAsyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAsyncTaskResponseBody) *DeleteAsyncTaskResponse
	GetBody() *DeleteAsyncTaskResponseBody
}

type DeleteAsyncTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAsyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteAsyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAsyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAsyncTaskResponse) GetBody() *DeleteAsyncTaskResponseBody {
	return s.Body
}

func (s *DeleteAsyncTaskResponse) SetHeaders(v map[string]*string) *DeleteAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteAsyncTaskResponse) SetStatusCode(v int32) *DeleteAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAsyncTaskResponse) SetBody(v *DeleteAsyncTaskResponseBody) *DeleteAsyncTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteAsyncTaskResponse) Validate() error {
	return dara.Validate(s)
}
