// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRobotNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRobotNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRobotNodeResponse
	GetStatusCode() *int32
	SetBody(v *ListRobotNodeResponseBody) *ListRobotNodeResponse
	GetBody() *ListRobotNodeResponseBody
}

type ListRobotNodeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRobotNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRobotNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRobotNodeResponse) GoString() string {
	return s.String()
}

func (s *ListRobotNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRobotNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRobotNodeResponse) GetBody() *ListRobotNodeResponseBody {
	return s.Body
}

func (s *ListRobotNodeResponse) SetHeaders(v map[string]*string) *ListRobotNodeResponse {
	s.Headers = v
	return s
}

func (s *ListRobotNodeResponse) SetStatusCode(v int32) *ListRobotNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRobotNodeResponse) SetBody(v *ListRobotNodeResponseBody) *ListRobotNodeResponse {
	s.Body = v
	return s
}

func (s *ListRobotNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
