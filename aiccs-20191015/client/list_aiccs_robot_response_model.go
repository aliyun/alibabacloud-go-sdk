// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiccsRobotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAiccsRobotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAiccsRobotResponse
	GetStatusCode() *int32
	SetBody(v *ListAiccsRobotResponseBody) *ListAiccsRobotResponse
	GetBody() *ListAiccsRobotResponseBody
}

type ListAiccsRobotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAiccsRobotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAiccsRobotResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAiccsRobotResponse) GoString() string {
	return s.String()
}

func (s *ListAiccsRobotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAiccsRobotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAiccsRobotResponse) GetBody() *ListAiccsRobotResponseBody {
	return s.Body
}

func (s *ListAiccsRobotResponse) SetHeaders(v map[string]*string) *ListAiccsRobotResponse {
	s.Headers = v
	return s
}

func (s *ListAiccsRobotResponse) SetStatusCode(v int32) *ListAiccsRobotResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAiccsRobotResponse) SetBody(v *ListAiccsRobotResponseBody) *ListAiccsRobotResponse {
	s.Body = v
	return s
}

func (s *ListAiccsRobotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
