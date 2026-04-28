// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAsyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAsyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetAsyncTaskResponseBody) *GetAsyncTaskResponse
	GetBody() *GetAsyncTaskResponseBody
}

type GetAsyncTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAsyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAsyncTaskResponse) GetBody() *GetAsyncTaskResponseBody {
	return s.Body
}

func (s *GetAsyncTaskResponse) SetHeaders(v map[string]*string) *GetAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncTaskResponse) SetStatusCode(v int32) *GetAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncTaskResponse) SetBody(v *GetAsyncTaskResponseBody) *GetAsyncTaskResponse {
	s.Body = v
	return s
}

func (s *GetAsyncTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
