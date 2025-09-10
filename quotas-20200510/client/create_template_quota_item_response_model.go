// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateQuotaItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTemplateQuotaItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTemplateQuotaItemResponse
	GetStatusCode() *int32
	SetBody(v *CreateTemplateQuotaItemResponseBody) *CreateTemplateQuotaItemResponse
	GetBody() *CreateTemplateQuotaItemResponseBody
}

type CreateTemplateQuotaItemResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTemplateQuotaItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTemplateQuotaItemResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateQuotaItemResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateQuotaItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTemplateQuotaItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTemplateQuotaItemResponse) GetBody() *CreateTemplateQuotaItemResponseBody {
	return s.Body
}

func (s *CreateTemplateQuotaItemResponse) SetHeaders(v map[string]*string) *CreateTemplateQuotaItemResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateQuotaItemResponse) SetStatusCode(v int32) *CreateTemplateQuotaItemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTemplateQuotaItemResponse) SetBody(v *CreateTemplateQuotaItemResponseBody) *CreateTemplateQuotaItemResponse {
	s.Body = v
	return s
}

func (s *CreateTemplateQuotaItemResponse) Validate() error {
	return dara.Validate(s)
}
