// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMonitorWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMonitorWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMonitorWebhookResponse
	GetStatusCode() *int32
	SetBody(v *MonitorWebhook) *UpdateMonitorWebhookResponse
	GetBody() *MonitorWebhook
}

type UpdateMonitorWebhookResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorWebhook    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMonitorWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMonitorWebhookResponse) GoString() string {
	return s.String()
}

func (s *UpdateMonitorWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMonitorWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMonitorWebhookResponse) GetBody() *MonitorWebhook {
	return s.Body
}

func (s *UpdateMonitorWebhookResponse) SetHeaders(v map[string]*string) *UpdateMonitorWebhookResponse {
	s.Headers = v
	return s
}

func (s *UpdateMonitorWebhookResponse) SetStatusCode(v int32) *UpdateMonitorWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMonitorWebhookResponse) SetBody(v *MonitorWebhook) *UpdateMonitorWebhookResponse {
	s.Body = v
	return s
}

func (s *UpdateMonitorWebhookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
