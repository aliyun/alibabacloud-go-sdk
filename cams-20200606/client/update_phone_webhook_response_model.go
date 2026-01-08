// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePhoneWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePhoneWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePhoneWebhookResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePhoneWebhookResponseBody) *UpdatePhoneWebhookResponse
	GetBody() *UpdatePhoneWebhookResponseBody
}

type UpdatePhoneWebhookResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePhoneWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePhoneWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePhoneWebhookResponse) GoString() string {
	return s.String()
}

func (s *UpdatePhoneWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePhoneWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePhoneWebhookResponse) GetBody() *UpdatePhoneWebhookResponseBody {
	return s.Body
}

func (s *UpdatePhoneWebhookResponse) SetHeaders(v map[string]*string) *UpdatePhoneWebhookResponse {
	s.Headers = v
	return s
}

func (s *UpdatePhoneWebhookResponse) SetStatusCode(v int32) *UpdatePhoneWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePhoneWebhookResponse) SetBody(v *UpdatePhoneWebhookResponseBody) *UpdatePhoneWebhookResponse {
	s.Body = v
	return s
}

func (s *UpdatePhoneWebhookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
