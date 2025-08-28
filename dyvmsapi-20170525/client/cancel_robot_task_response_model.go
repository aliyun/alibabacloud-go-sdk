// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRobotTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelRobotTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelRobotTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelRobotTaskResponseBody) *CancelRobotTaskResponse
	GetBody() *CancelRobotTaskResponseBody
}

type CancelRobotTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelRobotTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelRobotTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelRobotTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelRobotTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelRobotTaskResponse) GetBody() *CancelRobotTaskResponseBody {
	return s.Body
}

func (s *CancelRobotTaskResponse) SetHeaders(v map[string]*string) *CancelRobotTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelRobotTaskResponse) SetStatusCode(v int32) *CancelRobotTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelRobotTaskResponse) SetBody(v *CancelRobotTaskResponseBody) *CancelRobotTaskResponse {
	s.Body = v
	return s
}

func (s *CancelRobotTaskResponse) Validate() error {
	return dara.Validate(s)
}
