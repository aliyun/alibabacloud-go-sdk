// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotInfoListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRobotInfoListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRobotInfoListResponse
	GetStatusCode() *int32
	SetBody(v *QueryRobotInfoListResponseBody) *QueryRobotInfoListResponse
	GetBody() *QueryRobotInfoListResponseBody
}

type QueryRobotInfoListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRobotInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRobotInfoListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotInfoListResponse) GoString() string {
	return s.String()
}

func (s *QueryRobotInfoListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRobotInfoListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRobotInfoListResponse) GetBody() *QueryRobotInfoListResponseBody {
	return s.Body
}

func (s *QueryRobotInfoListResponse) SetHeaders(v map[string]*string) *QueryRobotInfoListResponse {
	s.Headers = v
	return s
}

func (s *QueryRobotInfoListResponse) SetStatusCode(v int32) *QueryRobotInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRobotInfoListResponse) SetBody(v *QueryRobotInfoListResponseBody) *QueryRobotInfoListResponse {
	s.Body = v
	return s
}

func (s *QueryRobotInfoListResponse) Validate() error {
	return dara.Validate(s)
}
