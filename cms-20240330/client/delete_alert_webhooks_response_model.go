// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertWebhooksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlertWebhooksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlertWebhooksResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAlertWebhooksResponseBody) *DeleteAlertWebhooksResponse
	GetBody() *DeleteAlertWebhooksResponseBody
}

type DeleteAlertWebhooksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlertWebhooksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlertWebhooksResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertWebhooksResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertWebhooksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlertWebhooksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlertWebhooksResponse) GetBody() *DeleteAlertWebhooksResponseBody {
	return s.Body
}

func (s *DeleteAlertWebhooksResponse) SetHeaders(v map[string]*string) *DeleteAlertWebhooksResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertWebhooksResponse) SetStatusCode(v int32) *DeleteAlertWebhooksResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertWebhooksResponse) SetBody(v *DeleteAlertWebhooksResponseBody) *DeleteAlertWebhooksResponse {
	s.Body = v
	return s
}

func (s *DeleteAlertWebhooksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
