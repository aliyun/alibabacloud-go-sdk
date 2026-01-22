// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertWebhooksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlertWebhooksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlertWebhooksResponse
	GetStatusCode() *int32
	SetBody(v *ListAlertWebhooksResponseBody) *ListAlertWebhooksResponse
	GetBody() *ListAlertWebhooksResponseBody
}

type ListAlertWebhooksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertWebhooksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertWebhooksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlertWebhooksResponse) GoString() string {
	return s.String()
}

func (s *ListAlertWebhooksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlertWebhooksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlertWebhooksResponse) GetBody() *ListAlertWebhooksResponseBody {
	return s.Body
}

func (s *ListAlertWebhooksResponse) SetHeaders(v map[string]*string) *ListAlertWebhooksResponse {
	s.Headers = v
	return s
}

func (s *ListAlertWebhooksResponse) SetStatusCode(v int32) *ListAlertWebhooksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertWebhooksResponse) SetBody(v *ListAlertWebhooksResponseBody) *ListAlertWebhooksResponse {
	s.Body = v
	return s
}

func (s *ListAlertWebhooksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
