// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskCallChatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TaskCallChatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TaskCallChatsResponse
	GetStatusCode() *int32
	SetBody(v *TaskCallChatsResponseBody) *TaskCallChatsResponse
	GetBody() *TaskCallChatsResponseBody
}

type TaskCallChatsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskCallChatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskCallChatsResponse) String() string {
	return dara.Prettify(s)
}

func (s TaskCallChatsResponse) GoString() string {
	return s.String()
}

func (s *TaskCallChatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TaskCallChatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TaskCallChatsResponse) GetBody() *TaskCallChatsResponseBody {
	return s.Body
}

func (s *TaskCallChatsResponse) SetHeaders(v map[string]*string) *TaskCallChatsResponse {
	s.Headers = v
	return s
}

func (s *TaskCallChatsResponse) SetStatusCode(v int32) *TaskCallChatsResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskCallChatsResponse) SetBody(v *TaskCallChatsResponseBody) *TaskCallChatsResponse {
	s.Body = v
	return s
}

func (s *TaskCallChatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
