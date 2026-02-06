// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskPreparingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TaskPreparingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TaskPreparingResponse
	GetStatusCode() *int32
	SetBody(v *TaskPreparingResponseBody) *TaskPreparingResponse
	GetBody() *TaskPreparingResponseBody
}

type TaskPreparingResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskPreparingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskPreparingResponse) String() string {
	return dara.Prettify(s)
}

func (s TaskPreparingResponse) GoString() string {
	return s.String()
}

func (s *TaskPreparingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TaskPreparingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TaskPreparingResponse) GetBody() *TaskPreparingResponseBody {
	return s.Body
}

func (s *TaskPreparingResponse) SetHeaders(v map[string]*string) *TaskPreparingResponse {
	s.Headers = v
	return s
}

func (s *TaskPreparingResponse) SetStatusCode(v int32) *TaskPreparingResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskPreparingResponse) SetBody(v *TaskPreparingResponseBody) *TaskPreparingResponse {
	s.Body = v
	return s
}

func (s *TaskPreparingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
