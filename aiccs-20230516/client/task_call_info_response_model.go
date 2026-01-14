// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskCallInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TaskCallInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TaskCallInfoResponse
	GetStatusCode() *int32
	SetBody(v *TaskCallInfoResponseBody) *TaskCallInfoResponse
	GetBody() *TaskCallInfoResponseBody
}

type TaskCallInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskCallInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskCallInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s TaskCallInfoResponse) GoString() string {
	return s.String()
}

func (s *TaskCallInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TaskCallInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TaskCallInfoResponse) GetBody() *TaskCallInfoResponseBody {
	return s.Body
}

func (s *TaskCallInfoResponse) SetHeaders(v map[string]*string) *TaskCallInfoResponse {
	s.Headers = v
	return s
}

func (s *TaskCallInfoResponse) SetStatusCode(v int32) *TaskCallInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskCallInfoResponse) SetBody(v *TaskCallInfoResponseBody) *TaskCallInfoResponse {
	s.Body = v
	return s
}

func (s *TaskCallInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
