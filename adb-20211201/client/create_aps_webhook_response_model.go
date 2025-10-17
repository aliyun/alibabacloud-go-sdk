// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApsWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApsWebhookResponse
	GetStatusCode() *int32
	SetBody(v *CreateApsWebhookResponseBody) *CreateApsWebhookResponse
	GetBody() *CreateApsWebhookResponseBody
}

type CreateApsWebhookResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApsWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApsWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApsWebhookResponse) GoString() string {
	return s.String()
}

func (s *CreateApsWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApsWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApsWebhookResponse) GetBody() *CreateApsWebhookResponseBody {
	return s.Body
}

func (s *CreateApsWebhookResponse) SetHeaders(v map[string]*string) *CreateApsWebhookResponse {
	s.Headers = v
	return s
}

func (s *CreateApsWebhookResponse) SetStatusCode(v int32) *CreateApsWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApsWebhookResponse) SetBody(v *CreateApsWebhookResponseBody) *CreateApsWebhookResponse {
	s.Body = v
	return s
}

func (s *CreateApsWebhookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
