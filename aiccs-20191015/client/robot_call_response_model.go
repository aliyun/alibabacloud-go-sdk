// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRobotCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RobotCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RobotCallResponse
	GetStatusCode() *int32
	SetBody(v *RobotCallResponseBody) *RobotCallResponse
	GetBody() *RobotCallResponseBody
}

type RobotCallResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RobotCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RobotCallResponse) String() string {
	return dara.Prettify(s)
}

func (s RobotCallResponse) GoString() string {
	return s.String()
}

func (s *RobotCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RobotCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RobotCallResponse) GetBody() *RobotCallResponseBody {
	return s.Body
}

func (s *RobotCallResponse) SetHeaders(v map[string]*string) *RobotCallResponse {
	s.Headers = v
	return s
}

func (s *RobotCallResponse) SetStatusCode(v int32) *RobotCallResponse {
	s.StatusCode = &v
	return s
}

func (s *RobotCallResponse) SetBody(v *RobotCallResponseBody) *RobotCallResponse {
	s.Body = v
	return s
}

func (s *RobotCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
