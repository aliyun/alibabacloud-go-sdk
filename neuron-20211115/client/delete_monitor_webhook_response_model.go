// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMonitorWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMonitorWebhookResponse
	GetStatusCode() *int32
}

type DeleteMonitorWebhookResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteMonitorWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorWebhookResponse) GoString() string {
	return s.String()
}

func (s *DeleteMonitorWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMonitorWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMonitorWebhookResponse) SetHeaders(v map[string]*string) *DeleteMonitorWebhookResponse {
	s.Headers = v
	return s
}

func (s *DeleteMonitorWebhookResponse) SetStatusCode(v int32) *DeleteMonitorWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMonitorWebhookResponse) Validate() error {
	return dara.Validate(s)
}
