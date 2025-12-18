// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateConfigRulesReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateConfigRulesReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateConfigRulesReportResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateConfigRulesReportResponseBody) *GetAggregateConfigRulesReportResponse
	GetBody() *GetAggregateConfigRulesReportResponseBody
}

type GetAggregateConfigRulesReportResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateConfigRulesReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateConfigRulesReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRulesReportResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRulesReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateConfigRulesReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateConfigRulesReportResponse) GetBody() *GetAggregateConfigRulesReportResponseBody {
	return s.Body
}

func (s *GetAggregateConfigRulesReportResponse) SetHeaders(v map[string]*string) *GetAggregateConfigRulesReportResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateConfigRulesReportResponse) SetStatusCode(v int32) *GetAggregateConfigRulesReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateConfigRulesReportResponse) SetBody(v *GetAggregateConfigRulesReportResponseBody) *GetAggregateConfigRulesReportResponse {
	s.Body = v
	return s
}

func (s *GetAggregateConfigRulesReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
