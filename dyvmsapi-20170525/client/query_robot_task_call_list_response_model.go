// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotTaskCallListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRobotTaskCallListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRobotTaskCallListResponse
	GetStatusCode() *int32
	SetBody(v *QueryRobotTaskCallListResponseBody) *QueryRobotTaskCallListResponse
	GetBody() *QueryRobotTaskCallListResponseBody
}

type QueryRobotTaskCallListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRobotTaskCallListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRobotTaskCallListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotTaskCallListResponse) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskCallListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRobotTaskCallListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRobotTaskCallListResponse) GetBody() *QueryRobotTaskCallListResponseBody {
	return s.Body
}

func (s *QueryRobotTaskCallListResponse) SetHeaders(v map[string]*string) *QueryRobotTaskCallListResponse {
	s.Headers = v
	return s
}

func (s *QueryRobotTaskCallListResponse) SetStatusCode(v int32) *QueryRobotTaskCallListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRobotTaskCallListResponse) SetBody(v *QueryRobotTaskCallListResponseBody) *QueryRobotTaskCallListResponse {
	s.Body = v
	return s
}

func (s *QueryRobotTaskCallListResponse) Validate() error {
	return dara.Validate(s)
}
