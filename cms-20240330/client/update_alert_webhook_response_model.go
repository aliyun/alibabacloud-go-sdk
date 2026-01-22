// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAlertWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAlertWebhookResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAlertWebhookResponseBody) *UpdateAlertWebhookResponse
	GetBody() *UpdateAlertWebhookResponseBody
}

type UpdateAlertWebhookResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlertWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlertWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertWebhookResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlertWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAlertWebhookResponse) GetBody() *UpdateAlertWebhookResponseBody {
	return s.Body
}

func (s *UpdateAlertWebhookResponse) SetHeaders(v map[string]*string) *UpdateAlertWebhookResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertWebhookResponse) SetStatusCode(v int32) *UpdateAlertWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertWebhookResponse) SetBody(v *UpdateAlertWebhookResponseBody) *UpdateAlertWebhookResponse {
	s.Body = v
	return s
}

func (s *UpdateAlertWebhookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
