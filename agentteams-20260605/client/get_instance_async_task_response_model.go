// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceAsyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceAsyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceAsyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceAsyncTaskResponseBody) *GetInstanceAsyncTaskResponse
	GetBody() *GetInstanceAsyncTaskResponseBody
}

type GetInstanceAsyncTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceAsyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceAsyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceAsyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceAsyncTaskResponse) GetBody() *GetInstanceAsyncTaskResponseBody {
	return s.Body
}

func (s *GetInstanceAsyncTaskResponse) SetHeaders(v map[string]*string) *GetInstanceAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceAsyncTaskResponse) SetStatusCode(v int32) *GetInstanceAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceAsyncTaskResponse) SetBody(v *GetInstanceAsyncTaskResponseBody) *GetInstanceAsyncTaskResponse {
	s.Body = v
	return s
}

func (s *GetInstanceAsyncTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
