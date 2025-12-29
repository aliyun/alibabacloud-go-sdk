// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeRobotPushResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokeRobotPushResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokeRobotPushResponse
	GetStatusCode() *int32
	SetBody(v *InvokeRobotPushResponseBody) *InvokeRobotPushResponse
	GetBody() *InvokeRobotPushResponseBody
}

type InvokeRobotPushResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokeRobotPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeRobotPushResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokeRobotPushResponse) GoString() string {
	return s.String()
}

func (s *InvokeRobotPushResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeRobotPushResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokeRobotPushResponse) GetBody() *InvokeRobotPushResponseBody {
	return s.Body
}

func (s *InvokeRobotPushResponse) SetHeaders(v map[string]*string) *InvokeRobotPushResponse {
	s.Headers = v
	return s
}

func (s *InvokeRobotPushResponse) SetStatusCode(v int32) *InvokeRobotPushResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeRobotPushResponse) SetBody(v *InvokeRobotPushResponseBody) *InvokeRobotPushResponse {
	s.Body = v
	return s
}

func (s *InvokeRobotPushResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
