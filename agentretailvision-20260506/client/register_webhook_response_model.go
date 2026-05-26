// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterWebhookResponse
	GetStatusCode() *int32
	SetBody(v *RegisterWebhookResponseBody) *RegisterWebhookResponse
	GetBody() *RegisterWebhookResponseBody
}

type RegisterWebhookResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterWebhookResponse) GoString() string {
	return s.String()
}

func (s *RegisterWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterWebhookResponse) GetBody() *RegisterWebhookResponseBody {
	return s.Body
}

func (s *RegisterWebhookResponse) SetHeaders(v map[string]*string) *RegisterWebhookResponse {
	s.Headers = v
	return s
}

func (s *RegisterWebhookResponse) SetStatusCode(v int32) *RegisterWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterWebhookResponse) SetBody(v *RegisterWebhookResponseBody) *RegisterWebhookResponse {
	s.Body = v
	return s
}

func (s *RegisterWebhookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
