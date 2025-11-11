// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAsyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAsyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAsyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAsyncTaskResponseBody) *SubmitAsyncTaskResponse
	GetBody() *SubmitAsyncTaskResponseBody
}

type SubmitAsyncTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAsyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitAsyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAsyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAsyncTaskResponse) GetBody() *SubmitAsyncTaskResponseBody {
	return s.Body
}

func (s *SubmitAsyncTaskResponse) SetHeaders(v map[string]*string) *SubmitAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitAsyncTaskResponse) SetStatusCode(v int32) *SubmitAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAsyncTaskResponse) SetBody(v *SubmitAsyncTaskResponseBody) *SubmitAsyncTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitAsyncTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
