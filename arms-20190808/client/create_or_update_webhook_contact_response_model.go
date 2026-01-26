// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateWebhookContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrUpdateWebhookContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrUpdateWebhookContactResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrUpdateWebhookContactResponseBody) *CreateOrUpdateWebhookContactResponse
	GetBody() *CreateOrUpdateWebhookContactResponseBody
}

type CreateOrUpdateWebhookContactResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrUpdateWebhookContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrUpdateWebhookContactResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateWebhookContactResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateWebhookContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrUpdateWebhookContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrUpdateWebhookContactResponse) GetBody() *CreateOrUpdateWebhookContactResponseBody {
	return s.Body
}

func (s *CreateOrUpdateWebhookContactResponse) SetHeaders(v map[string]*string) *CreateOrUpdateWebhookContactResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateWebhookContactResponse) SetStatusCode(v int32) *CreateOrUpdateWebhookContactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponse) SetBody(v *CreateOrUpdateWebhookContactResponseBody) *CreateOrUpdateWebhookContactResponse {
	s.Body = v
	return s
}

func (s *CreateOrUpdateWebhookContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
