// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAsyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAsyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAsyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateAsyncTaskResponseBody) *CreateAsyncTaskResponse
	GetBody() *CreateAsyncTaskResponseBody
}

type CreateAsyncTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAsyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAsyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAsyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAsyncTaskResponse) GetBody() *CreateAsyncTaskResponseBody {
	return s.Body
}

func (s *CreateAsyncTaskResponse) SetHeaders(v map[string]*string) *CreateAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAsyncTaskResponse) SetStatusCode(v int32) *CreateAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAsyncTaskResponse) SetBody(v *CreateAsyncTaskResponseBody) *CreateAsyncTaskResponse {
	s.Body = v
	return s
}

func (s *CreateAsyncTaskResponse) Validate() error {
	return dara.Validate(s)
}
