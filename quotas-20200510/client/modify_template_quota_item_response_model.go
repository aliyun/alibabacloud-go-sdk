// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateQuotaItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTemplateQuotaItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTemplateQuotaItemResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTemplateQuotaItemResponseBody) *ModifyTemplateQuotaItemResponse
	GetBody() *ModifyTemplateQuotaItemResponseBody
}

type ModifyTemplateQuotaItemResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTemplateQuotaItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTemplateQuotaItemResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateQuotaItemResponse) GoString() string {
	return s.String()
}

func (s *ModifyTemplateQuotaItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTemplateQuotaItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTemplateQuotaItemResponse) GetBody() *ModifyTemplateQuotaItemResponseBody {
	return s.Body
}

func (s *ModifyTemplateQuotaItemResponse) SetHeaders(v map[string]*string) *ModifyTemplateQuotaItemResponse {
	s.Headers = v
	return s
}

func (s *ModifyTemplateQuotaItemResponse) SetStatusCode(v int32) *ModifyTemplateQuotaItemResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTemplateQuotaItemResponse) SetBody(v *ModifyTemplateQuotaItemResponseBody) *ModifyTemplateQuotaItemResponse {
	s.Body = v
	return s
}

func (s *ModifyTemplateQuotaItemResponse) Validate() error {
	return dara.Validate(s)
}
