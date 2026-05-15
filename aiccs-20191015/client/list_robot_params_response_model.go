// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRobotParamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRobotParamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRobotParamsResponse
	GetStatusCode() *int32
	SetBody(v *ListRobotParamsResponseBody) *ListRobotParamsResponse
	GetBody() *ListRobotParamsResponseBody
}

type ListRobotParamsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRobotParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRobotParamsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRobotParamsResponse) GoString() string {
	return s.String()
}

func (s *ListRobotParamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRobotParamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRobotParamsResponse) GetBody() *ListRobotParamsResponseBody {
	return s.Body
}

func (s *ListRobotParamsResponse) SetHeaders(v map[string]*string) *ListRobotParamsResponse {
	s.Headers = v
	return s
}

func (s *ListRobotParamsResponse) SetStatusCode(v int32) *ListRobotParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRobotParamsResponse) SetBody(v *ListRobotParamsResponseBody) *ListRobotParamsResponse {
	s.Body = v
	return s
}

func (s *ListRobotParamsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
