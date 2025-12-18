// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAggregateConfigRulesReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateAggregateConfigRulesReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateAggregateConfigRulesReportResponse
	GetStatusCode() *int32
	SetBody(v *GenerateAggregateConfigRulesReportResponseBody) *GenerateAggregateConfigRulesReportResponse
	GetBody() *GenerateAggregateConfigRulesReportResponseBody
}

type GenerateAggregateConfigRulesReportResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateAggregateConfigRulesReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateAggregateConfigRulesReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateAggregateConfigRulesReportResponse) GoString() string {
	return s.String()
}

func (s *GenerateAggregateConfigRulesReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateAggregateConfigRulesReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateAggregateConfigRulesReportResponse) GetBody() *GenerateAggregateConfigRulesReportResponseBody {
	return s.Body
}

func (s *GenerateAggregateConfigRulesReportResponse) SetHeaders(v map[string]*string) *GenerateAggregateConfigRulesReportResponse {
	s.Headers = v
	return s
}

func (s *GenerateAggregateConfigRulesReportResponse) SetStatusCode(v int32) *GenerateAggregateConfigRulesReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAggregateConfigRulesReportResponse) SetBody(v *GenerateAggregateConfigRulesReportResponseBody) *GenerateAggregateConfigRulesReportResponse {
	s.Body = v
	return s
}

func (s *GenerateAggregateConfigRulesReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
