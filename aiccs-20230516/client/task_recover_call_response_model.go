// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskRecoverCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TaskRecoverCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TaskRecoverCallResponse
	GetStatusCode() *int32
	SetBody(v *TaskRecoverCallResponseBody) *TaskRecoverCallResponse
	GetBody() *TaskRecoverCallResponseBody
}

type TaskRecoverCallResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskRecoverCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskRecoverCallResponse) String() string {
	return dara.Prettify(s)
}

func (s TaskRecoverCallResponse) GoString() string {
	return s.String()
}

func (s *TaskRecoverCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TaskRecoverCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TaskRecoverCallResponse) GetBody() *TaskRecoverCallResponseBody {
	return s.Body
}

func (s *TaskRecoverCallResponse) SetHeaders(v map[string]*string) *TaskRecoverCallResponse {
	s.Headers = v
	return s
}

func (s *TaskRecoverCallResponse) SetStatusCode(v int32) *TaskRecoverCallResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskRecoverCallResponse) SetBody(v *TaskRecoverCallResponseBody) *TaskRecoverCallResponse {
	s.Body = v
	return s
}

func (s *TaskRecoverCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
