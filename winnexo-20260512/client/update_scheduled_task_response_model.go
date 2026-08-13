// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduledTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateScheduledTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateScheduledTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateScheduledTaskResponseBody) *UpdateScheduledTaskResponse
	GetBody() *UpdateScheduledTaskResponseBody
}

type UpdateScheduledTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScheduledTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateScheduledTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateScheduledTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateScheduledTaskResponse) GetBody() *UpdateScheduledTaskResponseBody {
	return s.Body
}

func (s *UpdateScheduledTaskResponse) SetHeaders(v map[string]*string) *UpdateScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateScheduledTaskResponse) SetStatusCode(v int32) *UpdateScheduledTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScheduledTaskResponse) SetBody(v *UpdateScheduledTaskResponseBody) *UpdateScheduledTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateScheduledTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
