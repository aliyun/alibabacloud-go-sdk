// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceAsyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceAsyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceAsyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceAsyncTaskResponseBody) *UpdateInstanceAsyncTaskResponse
	GetBody() *UpdateInstanceAsyncTaskResponseBody
}

type UpdateInstanceAsyncTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceAsyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAsyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceAsyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceAsyncTaskResponse) GetBody() *UpdateInstanceAsyncTaskResponseBody {
	return s.Body
}

func (s *UpdateInstanceAsyncTaskResponse) SetHeaders(v map[string]*string) *UpdateInstanceAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceAsyncTaskResponse) SetStatusCode(v int32) *UpdateInstanceAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceAsyncTaskResponse) SetBody(v *UpdateInstanceAsyncTaskResponseBody) *UpdateInstanceAsyncTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceAsyncTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
