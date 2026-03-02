// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMonitorWebhooksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMonitorWebhooksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMonitorWebhooksResponse
	GetStatusCode() *int32
	SetBody(v *MonitorWebhookPageResult) *ListMonitorWebhooksResponse
	GetBody() *MonitorWebhookPageResult
}

type ListMonitorWebhooksResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorWebhookPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMonitorWebhooksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMonitorWebhooksResponse) GoString() string {
	return s.String()
}

func (s *ListMonitorWebhooksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMonitorWebhooksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMonitorWebhooksResponse) GetBody() *MonitorWebhookPageResult {
	return s.Body
}

func (s *ListMonitorWebhooksResponse) SetHeaders(v map[string]*string) *ListMonitorWebhooksResponse {
	s.Headers = v
	return s
}

func (s *ListMonitorWebhooksResponse) SetStatusCode(v int32) *ListMonitorWebhooksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMonitorWebhooksResponse) SetBody(v *MonitorWebhookPageResult) *ListMonitorWebhooksResponse {
	s.Body = v
	return s
}

func (s *ListMonitorWebhooksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
