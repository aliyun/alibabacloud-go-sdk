// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateConfigRuleReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateConfigRuleReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateConfigRuleReportResponse
	GetStatusCode() *int32
	SetBody(v *GenerateConfigRuleReportResponseBody) *GenerateConfigRuleReportResponse
	GetBody() *GenerateConfigRuleReportResponseBody
}

type GenerateConfigRuleReportResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateConfigRuleReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateConfigRuleReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateConfigRuleReportResponse) GoString() string {
	return s.String()
}

func (s *GenerateConfigRuleReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateConfigRuleReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateConfigRuleReportResponse) GetBody() *GenerateConfigRuleReportResponseBody {
	return s.Body
}

func (s *GenerateConfigRuleReportResponse) SetHeaders(v map[string]*string) *GenerateConfigRuleReportResponse {
	s.Headers = v
	return s
}

func (s *GenerateConfigRuleReportResponse) SetStatusCode(v int32) *GenerateConfigRuleReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateConfigRuleReportResponse) SetBody(v *GenerateConfigRuleReportResponseBody) *GenerateConfigRuleReportResponse {
	s.Body = v
	return s
}

func (s *GenerateConfigRuleReportResponse) Validate() error {
	return dara.Validate(s)
}
