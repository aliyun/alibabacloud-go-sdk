// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebhookContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWebhookContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWebhookContactResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWebhookContactResponseBody) *DeleteWebhookContactResponse
	GetBody() *DeleteWebhookContactResponseBody
}

type DeleteWebhookContactResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWebhookContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWebhookContactResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebhookContactResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebhookContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWebhookContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWebhookContactResponse) GetBody() *DeleteWebhookContactResponseBody {
	return s.Body
}

func (s *DeleteWebhookContactResponse) SetHeaders(v map[string]*string) *DeleteWebhookContactResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebhookContactResponse) SetStatusCode(v int32) *DeleteWebhookContactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWebhookContactResponse) SetBody(v *DeleteWebhookContactResponseBody) *DeleteWebhookContactResponse {
	s.Body = v
	return s
}

func (s *DeleteWebhookContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
