// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWebhookResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWebhookResponseBody) *UpdateWebhookResponse
	GetBody() *UpdateWebhookResponseBody
}

type UpdateWebhookResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebhookResponse) GoString() string {
	return s.String()
}

func (s *UpdateWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWebhookResponse) GetBody() *UpdateWebhookResponseBody {
	return s.Body
}

func (s *UpdateWebhookResponse) SetHeaders(v map[string]*string) *UpdateWebhookResponse {
	s.Headers = v
	return s
}

func (s *UpdateWebhookResponse) SetStatusCode(v int32) *UpdateWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWebhookResponse) SetBody(v *UpdateWebhookResponseBody) *UpdateWebhookResponse {
	s.Body = v
	return s
}

func (s *UpdateWebhookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
