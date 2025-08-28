// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRobotTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRobotTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRobotTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateRobotTaskResponseBody) *CreateRobotTaskResponse
	GetBody() *CreateRobotTaskResponseBody
}

type CreateRobotTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRobotTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRobotTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRobotTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRobotTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRobotTaskResponse) GetBody() *CreateRobotTaskResponseBody {
	return s.Body
}

func (s *CreateRobotTaskResponse) SetHeaders(v map[string]*string) *CreateRobotTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRobotTaskResponse) SetStatusCode(v int32) *CreateRobotTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRobotTaskResponse) SetBody(v *CreateRobotTaskResponseBody) *CreateRobotTaskResponse {
	s.Body = v
	return s
}

func (s *CreateRobotTaskResponse) Validate() error {
	return dara.Validate(s)
}
