// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQuotaTemplateServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyQuotaTemplateServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyQuotaTemplateServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyQuotaTemplateServiceStatusResponseBody) *ModifyQuotaTemplateServiceStatusResponse
	GetBody() *ModifyQuotaTemplateServiceStatusResponseBody
}

type ModifyQuotaTemplateServiceStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyQuotaTemplateServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyQuotaTemplateServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyQuotaTemplateServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyQuotaTemplateServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyQuotaTemplateServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyQuotaTemplateServiceStatusResponse) GetBody() *ModifyQuotaTemplateServiceStatusResponseBody {
	return s.Body
}

func (s *ModifyQuotaTemplateServiceStatusResponse) SetHeaders(v map[string]*string) *ModifyQuotaTemplateServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyQuotaTemplateServiceStatusResponse) SetStatusCode(v int32) *ModifyQuotaTemplateServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyQuotaTemplateServiceStatusResponse) SetBody(v *ModifyQuotaTemplateServiceStatusResponseBody) *ModifyQuotaTemplateServiceStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyQuotaTemplateServiceStatusResponse) Validate() error {
	return dara.Validate(s)
}
