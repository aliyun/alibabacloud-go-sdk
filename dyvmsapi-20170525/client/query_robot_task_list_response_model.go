// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRobotTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRobotTaskListResponse
	GetStatusCode() *int32
	SetBody(v *QueryRobotTaskListResponseBody) *QueryRobotTaskListResponse
	GetBody() *QueryRobotTaskListResponseBody
}

type QueryRobotTaskListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRobotTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRobotTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotTaskListResponse) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRobotTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRobotTaskListResponse) GetBody() *QueryRobotTaskListResponseBody {
	return s.Body
}

func (s *QueryRobotTaskListResponse) SetHeaders(v map[string]*string) *QueryRobotTaskListResponse {
	s.Headers = v
	return s
}

func (s *QueryRobotTaskListResponse) SetStatusCode(v int32) *QueryRobotTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRobotTaskListResponse) SetBody(v *QueryRobotTaskListResponseBody) *QueryRobotTaskListResponse {
	s.Body = v
	return s
}

func (s *QueryRobotTaskListResponse) Validate() error {
	return dara.Validate(s)
}
