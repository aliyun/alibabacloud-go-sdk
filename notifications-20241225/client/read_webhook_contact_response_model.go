// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadWebhookContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadWebhookContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadWebhookContactResponse
	GetStatusCode() *int32
	SetBody(v *ReadWebhookContactResponseBody) *ReadWebhookContactResponse
	GetBody() *ReadWebhookContactResponseBody
}

type ReadWebhookContactResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadWebhookContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadWebhookContactResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadWebhookContactResponse) GoString() string {
	return s.String()
}

func (s *ReadWebhookContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadWebhookContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadWebhookContactResponse) GetBody() *ReadWebhookContactResponseBody {
	return s.Body
}

func (s *ReadWebhookContactResponse) SetHeaders(v map[string]*string) *ReadWebhookContactResponse {
	s.Headers = v
	return s
}

func (s *ReadWebhookContactResponse) SetStatusCode(v int32) *ReadWebhookContactResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadWebhookContactResponse) SetBody(v *ReadWebhookContactResponseBody) *ReadWebhookContactResponse {
	s.Body = v
	return s
}

func (s *ReadWebhookContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
