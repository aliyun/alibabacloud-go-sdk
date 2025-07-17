// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowNotificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskFlowNotificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskFlowNotificationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskFlowNotificationResponseBody) *UpdateTaskFlowNotificationResponse
	GetBody() *UpdateTaskFlowNotificationResponseBody
}

type UpdateTaskFlowNotificationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskFlowNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskFlowNotificationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowNotificationResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowNotificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskFlowNotificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskFlowNotificationResponse) GetBody() *UpdateTaskFlowNotificationResponseBody {
	return s.Body
}

func (s *UpdateTaskFlowNotificationResponse) SetHeaders(v map[string]*string) *UpdateTaskFlowNotificationResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskFlowNotificationResponse) SetStatusCode(v int32) *UpdateTaskFlowNotificationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskFlowNotificationResponse) SetBody(v *UpdateTaskFlowNotificationResponseBody) *UpdateTaskFlowNotificationResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskFlowNotificationResponse) Validate() error {
	return dara.Validate(s)
}
