// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRobotTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartRobotTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartRobotTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartRobotTaskResponseBody) *StartRobotTaskResponse
	GetBody() *StartRobotTaskResponseBody
}

type StartRobotTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRobotTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartRobotTaskResponse) GoString() string {
	return s.String()
}

func (s *StartRobotTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartRobotTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartRobotTaskResponse) GetBody() *StartRobotTaskResponseBody {
	return s.Body
}

func (s *StartRobotTaskResponse) SetHeaders(v map[string]*string) *StartRobotTaskResponse {
	s.Headers = v
	return s
}

func (s *StartRobotTaskResponse) SetStatusCode(v int32) *StartRobotTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRobotTaskResponse) SetBody(v *StartRobotTaskResponseBody) *StartRobotTaskResponse {
	s.Body = v
	return s
}

func (s *StartRobotTaskResponse) Validate() error {
	return dara.Validate(s)
}
