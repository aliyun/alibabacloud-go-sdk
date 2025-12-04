// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRobotTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRobotTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryRobotTaskDetailResponseBody) *QueryRobotTaskDetailResponse
	GetBody() *QueryRobotTaskDetailResponseBody
}

type QueryRobotTaskDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRobotTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRobotTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRobotTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRobotTaskDetailResponse) GetBody() *QueryRobotTaskDetailResponseBody {
	return s.Body
}

func (s *QueryRobotTaskDetailResponse) SetHeaders(v map[string]*string) *QueryRobotTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryRobotTaskDetailResponse) SetStatusCode(v int32) *QueryRobotTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRobotTaskDetailResponse) SetBody(v *QueryRobotTaskDetailResponseBody) *QueryRobotTaskDetailResponse {
	s.Body = v
	return s
}

func (s *QueryRobotTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
