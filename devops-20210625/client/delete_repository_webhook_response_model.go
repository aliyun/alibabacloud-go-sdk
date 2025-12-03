// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepositoryWebhookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRepositoryWebhookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRepositoryWebhookResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRepositoryWebhookResponseBody) *DeleteRepositoryWebhookResponse
	GetBody() *DeleteRepositoryWebhookResponseBody
}

type DeleteRepositoryWebhookResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRepositoryWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRepositoryWebhookResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryWebhookResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryWebhookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRepositoryWebhookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRepositoryWebhookResponse) GetBody() *DeleteRepositoryWebhookResponseBody {
	return s.Body
}

func (s *DeleteRepositoryWebhookResponse) SetHeaders(v map[string]*string) *DeleteRepositoryWebhookResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepositoryWebhookResponse) SetStatusCode(v int32) *DeleteRepositoryWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepositoryWebhookResponse) SetBody(v *DeleteRepositoryWebhookResponseBody) *DeleteRepositoryWebhookResponse {
	s.Body = v
	return s
}

func (s *DeleteRepositoryWebhookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
