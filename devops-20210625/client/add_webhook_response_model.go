// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddWebhookResponse
	GetStatusCode() *int32
	SetBody(v *AddWebhookResponseBody) *AddWebhookResponse
	GetBody() *AddWebhookResponseBody
}

type AddWebhookResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s AddWebhookResponse) GoString() string {
	return s.String()
}

func (s *AddWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddWebhookResponse) GetBody() *AddWebhookResponseBody {
	return s.Body
}

func (s *AddWebhookResponse) SetHeaders(v map[string]*string) *AddWebhookResponse {
	s.Headers = v
	return s
}

func (s *AddWebhookResponse) SetStatusCode(v int32) *AddWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWebhookResponse) SetBody(v *AddWebhookResponseBody) *AddWebhookResponse {
	s.Body = v
	return s
}

func (s *AddWebhookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
