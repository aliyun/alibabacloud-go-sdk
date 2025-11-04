// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRtcRobotInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartRtcRobotInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartRtcRobotInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StartRtcRobotInstanceResponseBody) *StartRtcRobotInstanceResponse
	GetBody() *StartRtcRobotInstanceResponseBody
}

type StartRtcRobotInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRtcRobotInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRtcRobotInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartRtcRobotInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartRtcRobotInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartRtcRobotInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartRtcRobotInstanceResponse) GetBody() *StartRtcRobotInstanceResponseBody {
	return s.Body
}

func (s *StartRtcRobotInstanceResponse) SetHeaders(v map[string]*string) *StartRtcRobotInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartRtcRobotInstanceResponse) SetStatusCode(v int32) *StartRtcRobotInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRtcRobotInstanceResponse) SetBody(v *StartRtcRobotInstanceResponseBody) *StartRtcRobotInstanceResponse {
	s.Body = v
	return s
}

func (s *StartRtcRobotInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
