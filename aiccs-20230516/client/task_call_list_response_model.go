// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskCallListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TaskCallListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TaskCallListResponse
	GetStatusCode() *int32
	SetBody(v *TaskCallListResponseBody) *TaskCallListResponse
	GetBody() *TaskCallListResponseBody
}

type TaskCallListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskCallListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskCallListResponse) String() string {
	return dara.Prettify(s)
}

func (s TaskCallListResponse) GoString() string {
	return s.String()
}

func (s *TaskCallListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TaskCallListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TaskCallListResponse) GetBody() *TaskCallListResponseBody {
	return s.Body
}

func (s *TaskCallListResponse) SetHeaders(v map[string]*string) *TaskCallListResponse {
	s.Headers = v
	return s
}

func (s *TaskCallListResponse) SetStatusCode(v int32) *TaskCallListResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskCallListResponse) SetBody(v *TaskCallListResponseBody) *TaskCallListResponse {
	s.Body = v
	return s
}

func (s *TaskCallListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
