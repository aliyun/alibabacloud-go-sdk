// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateConfigRulesReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateConfigRulesReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateConfigRulesReportResponse
	GetStatusCode() *int32
	SetBody(v *GenerateConfigRulesReportResponseBody) *GenerateConfigRulesReportResponse
	GetBody() *GenerateConfigRulesReportResponseBody
}

type GenerateConfigRulesReportResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateConfigRulesReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateConfigRulesReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateConfigRulesReportResponse) GoString() string {
	return s.String()
}

func (s *GenerateConfigRulesReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateConfigRulesReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateConfigRulesReportResponse) GetBody() *GenerateConfigRulesReportResponseBody {
	return s.Body
}

func (s *GenerateConfigRulesReportResponse) SetHeaders(v map[string]*string) *GenerateConfigRulesReportResponse {
	s.Headers = v
	return s
}

func (s *GenerateConfigRulesReportResponse) SetStatusCode(v int32) *GenerateConfigRulesReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateConfigRulesReportResponse) SetBody(v *GenerateConfigRulesReportResponseBody) *GenerateConfigRulesReportResponse {
	s.Body = v
	return s
}

func (s *GenerateConfigRulesReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
