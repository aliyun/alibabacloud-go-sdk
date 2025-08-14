// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRtcRobotInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRtcRobotInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRtcRobotInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRtcRobotInstanceResponseBody) *UpdateRtcRobotInstanceResponse
	GetBody() *UpdateRtcRobotInstanceResponseBody
}

type UpdateRtcRobotInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRtcRobotInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRtcRobotInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcRobotInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateRtcRobotInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRtcRobotInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRtcRobotInstanceResponse) GetBody() *UpdateRtcRobotInstanceResponseBody {
	return s.Body
}

func (s *UpdateRtcRobotInstanceResponse) SetHeaders(v map[string]*string) *UpdateRtcRobotInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateRtcRobotInstanceResponse) SetStatusCode(v int32) *UpdateRtcRobotInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRtcRobotInstanceResponse) SetBody(v *UpdateRtcRobotInstanceResponseBody) *UpdateRtcRobotInstanceResponse {
	s.Body = v
	return s
}

func (s *UpdateRtcRobotInstanceResponse) Validate() error {
	return dara.Validate(s)
}
