// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRobotTaskCallDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRobotTaskCallDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRobotTaskCallDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryRobotTaskCallDetailResponseBody) *QueryRobotTaskCallDetailResponse
	GetBody() *QueryRobotTaskCallDetailResponseBody
}

type QueryRobotTaskCallDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRobotTaskCallDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRobotTaskCallDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRobotTaskCallDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskCallDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRobotTaskCallDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRobotTaskCallDetailResponse) GetBody() *QueryRobotTaskCallDetailResponseBody {
	return s.Body
}

func (s *QueryRobotTaskCallDetailResponse) SetHeaders(v map[string]*string) *QueryRobotTaskCallDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryRobotTaskCallDetailResponse) SetStatusCode(v int32) *QueryRobotTaskCallDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRobotTaskCallDetailResponse) SetBody(v *QueryRobotTaskCallDetailResponseBody) *QueryRobotTaskCallDetailResponse {
	s.Body = v
	return s
}

func (s *QueryRobotTaskCallDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
