// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRobotTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopRobotTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopRobotTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopRobotTaskResponseBody) *StopRobotTaskResponse
	GetBody() *StopRobotTaskResponseBody
}

type StopRobotTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopRobotTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopRobotTaskResponse) GoString() string {
	return s.String()
}

func (s *StopRobotTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopRobotTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopRobotTaskResponse) GetBody() *StopRobotTaskResponseBody {
	return s.Body
}

func (s *StopRobotTaskResponse) SetHeaders(v map[string]*string) *StopRobotTaskResponse {
	s.Headers = v
	return s
}

func (s *StopRobotTaskResponse) SetStatusCode(v int32) *StopRobotTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRobotTaskResponse) SetBody(v *StopRobotTaskResponseBody) *StopRobotTaskResponse {
	s.Body = v
	return s
}

func (s *StopRobotTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
