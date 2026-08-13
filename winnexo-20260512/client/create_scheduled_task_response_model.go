// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScheduledTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScheduledTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateScheduledTaskResponseBody) *CreateScheduledTaskResponse
	GetBody() *CreateScheduledTaskResponseBody
}

type CreateScheduledTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduledTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScheduledTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScheduledTaskResponse) GetBody() *CreateScheduledTaskResponseBody {
	return s.Body
}

func (s *CreateScheduledTaskResponse) SetHeaders(v map[string]*string) *CreateScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduledTaskResponse) SetStatusCode(v int32) *CreateScheduledTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduledTaskResponse) SetBody(v *CreateScheduledTaskResponseBody) *CreateScheduledTaskResponse {
	s.Body = v
	return s
}

func (s *CreateScheduledTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
