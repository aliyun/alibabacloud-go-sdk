// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAlertWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAlertWebhookResponse
	GetStatusCode() *int32
	SetBody(v *CreateAlertWebhookResponseBody) *CreateAlertWebhookResponse
	GetBody() *CreateAlertWebhookResponseBody
}

type CreateAlertWebhookResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAlertWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlertWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertWebhookResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAlertWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAlertWebhookResponse) GetBody() *CreateAlertWebhookResponseBody {
	return s.Body
}

func (s *CreateAlertWebhookResponse) SetHeaders(v map[string]*string) *CreateAlertWebhookResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertWebhookResponse) SetStatusCode(v int32) *CreateAlertWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlertWebhookResponse) SetBody(v *CreateAlertWebhookResponseBody) *CreateAlertWebhookResponse {
	s.Body = v
	return s
}

func (s *CreateAlertWebhookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
