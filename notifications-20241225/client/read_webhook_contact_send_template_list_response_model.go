// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadWebhookContactSendTemplateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadWebhookContactSendTemplateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadWebhookContactSendTemplateListResponse
	GetStatusCode() *int32
	SetBody(v *ReadWebhookContactSendTemplateListResponseBody) *ReadWebhookContactSendTemplateListResponse
	GetBody() *ReadWebhookContactSendTemplateListResponseBody
}

type ReadWebhookContactSendTemplateListResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadWebhookContactSendTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadWebhookContactSendTemplateListResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadWebhookContactSendTemplateListResponse) GoString() string {
	return s.String()
}

func (s *ReadWebhookContactSendTemplateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadWebhookContactSendTemplateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadWebhookContactSendTemplateListResponse) GetBody() *ReadWebhookContactSendTemplateListResponseBody {
	return s.Body
}

func (s *ReadWebhookContactSendTemplateListResponse) SetHeaders(v map[string]*string) *ReadWebhookContactSendTemplateListResponse {
	s.Headers = v
	return s
}

func (s *ReadWebhookContactSendTemplateListResponse) SetStatusCode(v int32) *ReadWebhookContactSendTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadWebhookContactSendTemplateListResponse) SetBody(v *ReadWebhookContactSendTemplateListResponseBody) *ReadWebhookContactSendTemplateListResponse {
	s.Body = v
	return s
}

func (s *ReadWebhookContactSendTemplateListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
