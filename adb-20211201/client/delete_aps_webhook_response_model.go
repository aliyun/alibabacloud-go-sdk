// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApsWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApsWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApsWebhookResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApsWebhookResponseBody) *DeleteApsWebhookResponse
	GetBody() *DeleteApsWebhookResponseBody
}

type DeleteApsWebhookResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApsWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApsWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApsWebhookResponse) GoString() string {
	return s.String()
}

func (s *DeleteApsWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApsWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApsWebhookResponse) GetBody() *DeleteApsWebhookResponseBody {
	return s.Body
}

func (s *DeleteApsWebhookResponse) SetHeaders(v map[string]*string) *DeleteApsWebhookResponse {
	s.Headers = v
	return s
}

func (s *DeleteApsWebhookResponse) SetStatusCode(v int32) *DeleteApsWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApsWebhookResponse) SetBody(v *DeleteApsWebhookResponseBody) *DeleteApsWebhookResponse {
	s.Body = v
	return s
}

func (s *DeleteApsWebhookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
