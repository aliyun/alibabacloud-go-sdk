// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApsWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApsWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApsWebhookResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApsWebhookResponseBody) *UpdateApsWebhookResponse
	GetBody() *UpdateApsWebhookResponseBody
}

type UpdateApsWebhookResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApsWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApsWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApsWebhookResponse) GoString() string {
	return s.String()
}

func (s *UpdateApsWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApsWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApsWebhookResponse) GetBody() *UpdateApsWebhookResponseBody {
	return s.Body
}

func (s *UpdateApsWebhookResponse) SetHeaders(v map[string]*string) *UpdateApsWebhookResponse {
	s.Headers = v
	return s
}

func (s *UpdateApsWebhookResponse) SetStatusCode(v int32) *UpdateApsWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApsWebhookResponse) SetBody(v *UpdateApsWebhookResponseBody) *UpdateApsWebhookResponse {
	s.Body = v
	return s
}

func (s *UpdateApsWebhookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
