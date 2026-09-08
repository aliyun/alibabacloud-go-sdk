// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebhookContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWebhookContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWebhookContactResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWebhookContactResponseBody) *UpdateWebhookContactResponse
	GetBody() *UpdateWebhookContactResponseBody
}

type UpdateWebhookContactResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWebhookContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWebhookContactResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebhookContactResponse) GoString() string {
	return s.String()
}

func (s *UpdateWebhookContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWebhookContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWebhookContactResponse) GetBody() *UpdateWebhookContactResponseBody {
	return s.Body
}

func (s *UpdateWebhookContactResponse) SetHeaders(v map[string]*string) *UpdateWebhookContactResponse {
	s.Headers = v
	return s
}

func (s *UpdateWebhookContactResponse) SetStatusCode(v int32) *UpdateWebhookContactResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWebhookContactResponse) SetBody(v *UpdateWebhookContactResponseBody) *UpdateWebhookContactResponse {
	s.Body = v
	return s
}

func (s *UpdateWebhookContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
