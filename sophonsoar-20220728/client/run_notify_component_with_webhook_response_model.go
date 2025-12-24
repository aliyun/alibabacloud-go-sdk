// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunNotifyComponentWithWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunNotifyComponentWithWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunNotifyComponentWithWebhookResponse
	GetStatusCode() *int32
	SetBody(v *RunNotifyComponentWithWebhookResponseBody) *RunNotifyComponentWithWebhookResponse
	GetBody() *RunNotifyComponentWithWebhookResponseBody
}

type RunNotifyComponentWithWebhookResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunNotifyComponentWithWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunNotifyComponentWithWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s RunNotifyComponentWithWebhookResponse) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunNotifyComponentWithWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunNotifyComponentWithWebhookResponse) GetBody() *RunNotifyComponentWithWebhookResponseBody {
	return s.Body
}

func (s *RunNotifyComponentWithWebhookResponse) SetHeaders(v map[string]*string) *RunNotifyComponentWithWebhookResponse {
	s.Headers = v
	return s
}

func (s *RunNotifyComponentWithWebhookResponse) SetStatusCode(v int32) *RunNotifyComponentWithWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *RunNotifyComponentWithWebhookResponse) SetBody(v *RunNotifyComponentWithWebhookResponseBody) *RunNotifyComponentWithWebhookResponse {
	s.Body = v
	return s
}

func (s *RunNotifyComponentWithWebhookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
