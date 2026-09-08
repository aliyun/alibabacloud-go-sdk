// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestWebhookContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestWebhookContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestWebhookContactResponse
	GetStatusCode() *int32
	SetBody(v *TestWebhookContactResponseBody) *TestWebhookContactResponse
	GetBody() *TestWebhookContactResponseBody
}

type TestWebhookContactResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestWebhookContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestWebhookContactResponse) String() string {
	return dara.Prettify(s)
}

func (s TestWebhookContactResponse) GoString() string {
	return s.String()
}

func (s *TestWebhookContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestWebhookContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestWebhookContactResponse) GetBody() *TestWebhookContactResponseBody {
	return s.Body
}

func (s *TestWebhookContactResponse) SetHeaders(v map[string]*string) *TestWebhookContactResponse {
	s.Headers = v
	return s
}

func (s *TestWebhookContactResponse) SetStatusCode(v int32) *TestWebhookContactResponse {
	s.StatusCode = &v
	return s
}

func (s *TestWebhookContactResponse) SetBody(v *TestWebhookContactResponseBody) *TestWebhookContactResponse {
	s.Body = v
	return s
}

func (s *TestWebhookContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
