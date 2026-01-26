// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIMRobotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIMRobotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIMRobotResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIMRobotResponseBody) *DeleteIMRobotResponse
	GetBody() *DeleteIMRobotResponseBody
}

type DeleteIMRobotResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIMRobotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIMRobotResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIMRobotResponse) GoString() string {
	return s.String()
}

func (s *DeleteIMRobotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIMRobotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIMRobotResponse) GetBody() *DeleteIMRobotResponseBody {
	return s.Body
}

func (s *DeleteIMRobotResponse) SetHeaders(v map[string]*string) *DeleteIMRobotResponse {
	s.Headers = v
	return s
}

func (s *DeleteIMRobotResponse) SetStatusCode(v int32) *DeleteIMRobotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIMRobotResponse) SetBody(v *DeleteIMRobotResponseBody) *DeleteIMRobotResponse {
	s.Body = v
	return s
}

func (s *DeleteIMRobotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
