// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepositoryWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepositoryWebhookResponse
	GetStatusCode() *int32
	SetBody(v *ListRepositoryWebhookResponseBody) *ListRepositoryWebhookResponse
	GetBody() *ListRepositoryWebhookResponseBody
}

type ListRepositoryWebhookResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepositoryWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepositoryWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryWebhookResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepositoryWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepositoryWebhookResponse) GetBody() *ListRepositoryWebhookResponseBody {
	return s.Body
}

func (s *ListRepositoryWebhookResponse) SetHeaders(v map[string]*string) *ListRepositoryWebhookResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryWebhookResponse) SetStatusCode(v int32) *ListRepositoryWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryWebhookResponse) SetBody(v *ListRepositoryWebhookResponseBody) *ListRepositoryWebhookResponse {
	s.Body = v
	return s
}

func (s *ListRepositoryWebhookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
