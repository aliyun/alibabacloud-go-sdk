// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadAllWebhookContactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadAllWebhookContactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadAllWebhookContactsResponse
	GetStatusCode() *int32
	SetBody(v *ReadAllWebhookContactsResponseBody) *ReadAllWebhookContactsResponse
	GetBody() *ReadAllWebhookContactsResponseBody
}

type ReadAllWebhookContactsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadAllWebhookContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadAllWebhookContactsResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadAllWebhookContactsResponse) GoString() string {
	return s.String()
}

func (s *ReadAllWebhookContactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadAllWebhookContactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadAllWebhookContactsResponse) GetBody() *ReadAllWebhookContactsResponseBody {
	return s.Body
}

func (s *ReadAllWebhookContactsResponse) SetHeaders(v map[string]*string) *ReadAllWebhookContactsResponse {
	s.Headers = v
	return s
}

func (s *ReadAllWebhookContactsResponse) SetStatusCode(v int32) *ReadAllWebhookContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadAllWebhookContactsResponse) SetBody(v *ReadAllWebhookContactsResponseBody) *ReadAllWebhookContactsResponse {
	s.Body = v
	return s
}

func (s *ReadAllWebhookContactsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
