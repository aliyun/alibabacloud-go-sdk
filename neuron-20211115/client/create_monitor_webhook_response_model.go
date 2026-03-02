// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitorWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitorWebhookResponse
	GetStatusCode() *int32
	SetBody(v *MonitorWebhook) *CreateMonitorWebhookResponse
	GetBody() *MonitorWebhook
}

type CreateMonitorWebhookResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorWebhook    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorWebhookResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitorWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitorWebhookResponse) GetBody() *MonitorWebhook {
	return s.Body
}

func (s *CreateMonitorWebhookResponse) SetHeaders(v map[string]*string) *CreateMonitorWebhookResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorWebhookResponse) SetStatusCode(v int32) *CreateMonitorWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorWebhookResponse) SetBody(v *MonitorWebhook) *CreateMonitorWebhookResponse {
	s.Body = v
	return s
}

func (s *CreateMonitorWebhookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
