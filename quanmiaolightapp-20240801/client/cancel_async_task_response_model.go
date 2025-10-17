// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAsyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelAsyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelAsyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelAsyncTaskResponseBody) *CancelAsyncTaskResponse
	GetBody() *CancelAsyncTaskResponseBody
}

type CancelAsyncTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelAsyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelAsyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelAsyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelAsyncTaskResponse) GetBody() *CancelAsyncTaskResponseBody {
	return s.Body
}

func (s *CancelAsyncTaskResponse) SetHeaders(v map[string]*string) *CancelAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelAsyncTaskResponse) SetStatusCode(v int32) *CancelAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAsyncTaskResponse) SetBody(v *CancelAsyncTaskResponseBody) *CancelAsyncTaskResponse {
	s.Body = v
	return s
}

func (s *CancelAsyncTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
