// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOrderRobotTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelOrderRobotTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelOrderRobotTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelOrderRobotTaskResponseBody) *CancelOrderRobotTaskResponse
	GetBody() *CancelOrderRobotTaskResponseBody
}

type CancelOrderRobotTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelOrderRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelOrderRobotTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelOrderRobotTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelOrderRobotTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelOrderRobotTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelOrderRobotTaskResponse) GetBody() *CancelOrderRobotTaskResponseBody {
	return s.Body
}

func (s *CancelOrderRobotTaskResponse) SetHeaders(v map[string]*string) *CancelOrderRobotTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelOrderRobotTaskResponse) SetStatusCode(v int32) *CancelOrderRobotTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelOrderRobotTaskResponse) SetBody(v *CancelOrderRobotTaskResponseBody) *CancelOrderRobotTaskResponse {
	s.Body = v
	return s
}

func (s *CancelOrderRobotTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
