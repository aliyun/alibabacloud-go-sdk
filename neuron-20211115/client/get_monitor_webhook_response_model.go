// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMonitorWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMonitorWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMonitorWebhookResponse
	GetStatusCode() *int32
	SetBody(v *MonitorWebhook) *GetMonitorWebhookResponse
	GetBody() *MonitorWebhook
}

type GetMonitorWebhookResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorWebhook    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMonitorWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMonitorWebhookResponse) GoString() string {
	return s.String()
}

func (s *GetMonitorWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMonitorWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMonitorWebhookResponse) GetBody() *MonitorWebhook {
	return s.Body
}

func (s *GetMonitorWebhookResponse) SetHeaders(v map[string]*string) *GetMonitorWebhookResponse {
	s.Headers = v
	return s
}

func (s *GetMonitorWebhookResponse) SetStatusCode(v int32) *GetMonitorWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMonitorWebhookResponse) SetBody(v *MonitorWebhook) *GetMonitorWebhookResponse {
	s.Body = v
	return s
}

func (s *GetMonitorWebhookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
