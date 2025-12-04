// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRobotTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRobotTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRobotTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRobotTaskResponseBody) *DeleteRobotTaskResponse
	GetBody() *DeleteRobotTaskResponseBody
}

type DeleteRobotTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRobotTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRobotTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteRobotTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRobotTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRobotTaskResponse) GetBody() *DeleteRobotTaskResponseBody {
	return s.Body
}

func (s *DeleteRobotTaskResponse) SetHeaders(v map[string]*string) *DeleteRobotTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteRobotTaskResponse) SetStatusCode(v int32) *DeleteRobotTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRobotTaskResponse) SetBody(v *DeleteRobotTaskResponseBody) *DeleteRobotTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteRobotTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
