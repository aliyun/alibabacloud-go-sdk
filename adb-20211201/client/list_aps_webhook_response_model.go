// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApsWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApsWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApsWebhookResponse
	GetStatusCode() *int32
	SetBody(v *ListApsWebhookResponseBody) *ListApsWebhookResponse
	GetBody() *ListApsWebhookResponseBody
}

type ListApsWebhookResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApsWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApsWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApsWebhookResponse) GoString() string {
	return s.String()
}

func (s *ListApsWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApsWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApsWebhookResponse) GetBody() *ListApsWebhookResponseBody {
	return s.Body
}

func (s *ListApsWebhookResponse) SetHeaders(v map[string]*string) *ListApsWebhookResponse {
	s.Headers = v
	return s
}

func (s *ListApsWebhookResponse) SetStatusCode(v int32) *ListApsWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApsWebhookResponse) SetBody(v *ListApsWebhookResponseBody) *ListApsWebhookResponse {
	s.Body = v
	return s
}

func (s *ListApsWebhookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
