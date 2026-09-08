// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebhookContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWebhookContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWebhookContactResponse
	GetStatusCode() *int32
	SetBody(v *CreateWebhookContactResponseBody) *CreateWebhookContactResponse
	GetBody() *CreateWebhookContactResponseBody
}

type CreateWebhookContactResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWebhookContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWebhookContactResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWebhookContactResponse) GoString() string {
	return s.String()
}

func (s *CreateWebhookContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWebhookContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWebhookContactResponse) GetBody() *CreateWebhookContactResponseBody {
	return s.Body
}

func (s *CreateWebhookContactResponse) SetHeaders(v map[string]*string) *CreateWebhookContactResponse {
	s.Headers = v
	return s
}

func (s *CreateWebhookContactResponse) SetStatusCode(v int32) *CreateWebhookContactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWebhookContactResponse) SetBody(v *CreateWebhookContactResponseBody) *CreateWebhookContactResponse {
	s.Body = v
	return s
}

func (s *CreateWebhookContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
