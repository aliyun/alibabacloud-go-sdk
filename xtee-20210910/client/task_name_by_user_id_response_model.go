// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskNameByUserIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TaskNameByUserIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TaskNameByUserIdResponse
	GetStatusCode() *int32
	SetBody(v *TaskNameByUserIdResponseBody) *TaskNameByUserIdResponse
	GetBody() *TaskNameByUserIdResponseBody
}

type TaskNameByUserIdResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskNameByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskNameByUserIdResponse) String() string {
	return dara.Prettify(s)
}

func (s TaskNameByUserIdResponse) GoString() string {
	return s.String()
}

func (s *TaskNameByUserIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TaskNameByUserIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TaskNameByUserIdResponse) GetBody() *TaskNameByUserIdResponseBody {
	return s.Body
}

func (s *TaskNameByUserIdResponse) SetHeaders(v map[string]*string) *TaskNameByUserIdResponse {
	s.Headers = v
	return s
}

func (s *TaskNameByUserIdResponse) SetStatusCode(v int32) *TaskNameByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskNameByUserIdResponse) SetBody(v *TaskNameByUserIdResponseBody) *TaskNameByUserIdResponse {
	s.Body = v
	return s
}

func (s *TaskNameByUserIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
