// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskFlowNotificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskFlowNotificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskFlowNotificationResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskFlowNotificationResponseBody) *GetTaskFlowNotificationResponse
	GetBody() *GetTaskFlowNotificationResponseBody
}

type GetTaskFlowNotificationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskFlowNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskFlowNotificationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskFlowNotificationResponse) GoString() string {
	return s.String()
}

func (s *GetTaskFlowNotificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskFlowNotificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskFlowNotificationResponse) GetBody() *GetTaskFlowNotificationResponseBody {
	return s.Body
}

func (s *GetTaskFlowNotificationResponse) SetHeaders(v map[string]*string) *GetTaskFlowNotificationResponse {
	s.Headers = v
	return s
}

func (s *GetTaskFlowNotificationResponse) SetStatusCode(v int32) *GetTaskFlowNotificationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskFlowNotificationResponse) SetBody(v *GetTaskFlowNotificationResponseBody) *GetTaskFlowNotificationResponse {
	s.Body = v
	return s
}

func (s *GetTaskFlowNotificationResponse) Validate() error {
	return dara.Validate(s)
}
