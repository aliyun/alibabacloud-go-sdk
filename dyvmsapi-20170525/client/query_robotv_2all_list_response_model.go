// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotv2AllListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRobotv2AllListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRobotv2AllListResponse
	GetStatusCode() *int32
	SetBody(v *QueryRobotv2AllListResponseBody) *QueryRobotv2AllListResponse
	GetBody() *QueryRobotv2AllListResponseBody
}

type QueryRobotv2AllListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRobotv2AllListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRobotv2AllListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotv2AllListResponse) GoString() string {
	return s.String()
}

func (s *QueryRobotv2AllListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRobotv2AllListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRobotv2AllListResponse) GetBody() *QueryRobotv2AllListResponseBody {
	return s.Body
}

func (s *QueryRobotv2AllListResponse) SetHeaders(v map[string]*string) *QueryRobotv2AllListResponse {
	s.Headers = v
	return s
}

func (s *QueryRobotv2AllListResponse) SetStatusCode(v int32) *QueryRobotv2AllListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRobotv2AllListResponse) SetBody(v *QueryRobotv2AllListResponseBody) *QueryRobotv2AllListResponse {
	s.Body = v
	return s
}

func (s *QueryRobotv2AllListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
