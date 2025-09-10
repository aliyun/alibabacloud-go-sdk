// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaApplicationsForTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQuotaApplicationsForTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQuotaApplicationsForTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateQuotaApplicationsForTemplateResponseBody) *CreateQuotaApplicationsForTemplateResponse
	GetBody() *CreateQuotaApplicationsForTemplateResponseBody
}

type CreateQuotaApplicationsForTemplateResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQuotaApplicationsForTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQuotaApplicationsForTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaApplicationsForTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateQuotaApplicationsForTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQuotaApplicationsForTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQuotaApplicationsForTemplateResponse) GetBody() *CreateQuotaApplicationsForTemplateResponseBody {
	return s.Body
}

func (s *CreateQuotaApplicationsForTemplateResponse) SetHeaders(v map[string]*string) *CreateQuotaApplicationsForTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateQuotaApplicationsForTemplateResponse) SetStatusCode(v int32) *CreateQuotaApplicationsForTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQuotaApplicationsForTemplateResponse) SetBody(v *CreateQuotaApplicationsForTemplateResponseBody) *CreateQuotaApplicationsForTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateQuotaApplicationsForTemplateResponse) Validate() error {
	return dara.Validate(s)
}
