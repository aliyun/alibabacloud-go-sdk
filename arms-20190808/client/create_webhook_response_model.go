// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWebhookResponse
	GetStatusCode() *int32
	SetBody(v *CreateWebhookResponseBody) *CreateWebhookResponse
	GetBody() *CreateWebhookResponseBody
}

type CreateWebhookResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWebhookResponse) GoString() string {
	return s.String()
}

func (s *CreateWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWebhookResponse) GetBody() *CreateWebhookResponseBody {
	return s.Body
}

func (s *CreateWebhookResponse) SetHeaders(v map[string]*string) *CreateWebhookResponse {
	s.Headers = v
	return s
}

func (s *CreateWebhookResponse) SetStatusCode(v int32) *CreateWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWebhookResponse) SetBody(v *CreateWebhookResponseBody) *CreateWebhookResponse {
	s.Body = v
	return s
}

func (s *CreateWebhookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
