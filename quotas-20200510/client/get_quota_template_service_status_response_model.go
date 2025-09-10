// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaTemplateServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQuotaTemplateServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQuotaTemplateServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetQuotaTemplateServiceStatusResponseBody) *GetQuotaTemplateServiceStatusResponse
	GetBody() *GetQuotaTemplateServiceStatusResponseBody
}

type GetQuotaTemplateServiceStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaTemplateServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaTemplateServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaTemplateServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaTemplateServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQuotaTemplateServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQuotaTemplateServiceStatusResponse) GetBody() *GetQuotaTemplateServiceStatusResponseBody {
	return s.Body
}

func (s *GetQuotaTemplateServiceStatusResponse) SetHeaders(v map[string]*string) *GetQuotaTemplateServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaTemplateServiceStatusResponse) SetStatusCode(v int32) *GetQuotaTemplateServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaTemplateServiceStatusResponse) SetBody(v *GetQuotaTemplateServiceStatusResponseBody) *GetQuotaTemplateServiceStatusResponse {
	s.Body = v
	return s
}

func (s *GetQuotaTemplateServiceStatusResponse) Validate() error {
	return dara.Validate(s)
}
