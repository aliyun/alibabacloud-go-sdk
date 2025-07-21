// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAccountWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAccountWebhookResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAccountWebhookResponseBody) *UpdateAccountWebhookResponse
	GetBody() *UpdateAccountWebhookResponseBody
}

type UpdateAccountWebhookResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccountWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccountWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountWebhookResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccountWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAccountWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAccountWebhookResponse) GetBody() *UpdateAccountWebhookResponseBody {
	return s.Body
}

func (s *UpdateAccountWebhookResponse) SetHeaders(v map[string]*string) *UpdateAccountWebhookResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccountWebhookResponse) SetStatusCode(v int32) *UpdateAccountWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccountWebhookResponse) SetBody(v *UpdateAccountWebhookResponseBody) *UpdateAccountWebhookResponse {
	s.Body = v
	return s
}

func (s *UpdateAccountWebhookResponse) Validate() error {
	return dara.Validate(s)
}
