// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRtcRobotInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopRtcRobotInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopRtcRobotInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StopRtcRobotInstanceResponseBody) *StopRtcRobotInstanceResponse
	GetBody() *StopRtcRobotInstanceResponseBody
}

type StopRtcRobotInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRtcRobotInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopRtcRobotInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopRtcRobotInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopRtcRobotInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopRtcRobotInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopRtcRobotInstanceResponse) GetBody() *StopRtcRobotInstanceResponseBody {
	return s.Body
}

func (s *StopRtcRobotInstanceResponse) SetHeaders(v map[string]*string) *StopRtcRobotInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopRtcRobotInstanceResponse) SetStatusCode(v int32) *StopRtcRobotInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRtcRobotInstanceResponse) SetBody(v *StopRtcRobotInstanceResponseBody) *StopRtcRobotInstanceResponse {
	s.Body = v
	return s
}

func (s *StopRtcRobotInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
