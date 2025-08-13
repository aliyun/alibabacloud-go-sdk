// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigRuleReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConfigRuleReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConfigRuleReportResponse
	GetStatusCode() *int32
	SetBody(v *GetConfigRuleReportResponseBody) *GetConfigRuleReportResponse
	GetBody() *GetConfigRuleReportResponseBody
}

type GetConfigRuleReportResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConfigRuleReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConfigRuleReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleReportResponse) GoString() string {
	return s.String()
}

func (s *GetConfigRuleReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConfigRuleReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConfigRuleReportResponse) GetBody() *GetConfigRuleReportResponseBody {
	return s.Body
}

func (s *GetConfigRuleReportResponse) SetHeaders(v map[string]*string) *GetConfigRuleReportResponse {
	s.Headers = v
	return s
}

func (s *GetConfigRuleReportResponse) SetStatusCode(v int32) *GetConfigRuleReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfigRuleReportResponse) SetBody(v *GetConfigRuleReportResponseBody) *GetConfigRuleReportResponse {
	s.Body = v
	return s
}

func (s *GetConfigRuleReportResponse) Validate() error {
	return dara.Validate(s)
}
