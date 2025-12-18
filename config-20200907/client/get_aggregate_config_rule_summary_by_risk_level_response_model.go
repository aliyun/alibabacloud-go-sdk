// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateConfigRuleSummaryByRiskLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateConfigRuleSummaryByRiskLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateConfigRuleSummaryByRiskLevelResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateConfigRuleSummaryByRiskLevelResponseBody) *GetAggregateConfigRuleSummaryByRiskLevelResponse
	GetBody() *GetAggregateConfigRuleSummaryByRiskLevelResponseBody
}

type GetAggregateConfigRuleSummaryByRiskLevelResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateConfigRuleSummaryByRiskLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateConfigRuleSummaryByRiskLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleSummaryByRiskLevelResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponse) GetBody() *GetAggregateConfigRuleSummaryByRiskLevelResponseBody {
	return s.Body
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponse) SetHeaders(v map[string]*string) *GetAggregateConfigRuleSummaryByRiskLevelResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponse) SetStatusCode(v int32) *GetAggregateConfigRuleSummaryByRiskLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponse) SetBody(v *GetAggregateConfigRuleSummaryByRiskLevelResponseBody) *GetAggregateConfigRuleSummaryByRiskLevelResponse {
	s.Body = v
	return s
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
