// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappTemplateMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatappTemplateMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatappTemplateMetricResponse
	GetStatusCode() *int32
	SetBody(v *GetChatappTemplateMetricResponseBody) *GetChatappTemplateMetricResponse
	GetBody() *GetChatappTemplateMetricResponseBody
}

type GetChatappTemplateMetricResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatappTemplateMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatappTemplateMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatappTemplateMetricResponse) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatappTemplateMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatappTemplateMetricResponse) GetBody() *GetChatappTemplateMetricResponseBody {
	return s.Body
}

func (s *GetChatappTemplateMetricResponse) SetHeaders(v map[string]*string) *GetChatappTemplateMetricResponse {
	s.Headers = v
	return s
}

func (s *GetChatappTemplateMetricResponse) SetStatusCode(v int32) *GetChatappTemplateMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatappTemplateMetricResponse) SetBody(v *GetChatappTemplateMetricResponseBody) *GetChatappTemplateMetricResponse {
	s.Body = v
	return s
}

func (s *GetChatappTemplateMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
