// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigRuleSummaryByRiskLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConfigRuleSummaryByRiskLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConfigRuleSummaryByRiskLevelResponse
	GetStatusCode() *int32
	SetBody(v *GetConfigRuleSummaryByRiskLevelResponseBody) *GetConfigRuleSummaryByRiskLevelResponse
	GetBody() *GetConfigRuleSummaryByRiskLevelResponseBody
}

type GetConfigRuleSummaryByRiskLevelResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConfigRuleSummaryByRiskLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConfigRuleSummaryByRiskLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleSummaryByRiskLevelResponse) GoString() string {
	return s.String()
}

func (s *GetConfigRuleSummaryByRiskLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConfigRuleSummaryByRiskLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConfigRuleSummaryByRiskLevelResponse) GetBody() *GetConfigRuleSummaryByRiskLevelResponseBody {
	return s.Body
}

func (s *GetConfigRuleSummaryByRiskLevelResponse) SetHeaders(v map[string]*string) *GetConfigRuleSummaryByRiskLevelResponse {
	s.Headers = v
	return s
}

func (s *GetConfigRuleSummaryByRiskLevelResponse) SetStatusCode(v int32) *GetConfigRuleSummaryByRiskLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfigRuleSummaryByRiskLevelResponse) SetBody(v *GetConfigRuleSummaryByRiskLevelResponseBody) *GetConfigRuleSummaryByRiskLevelResponse {
	s.Body = v
	return s
}

func (s *GetConfigRuleSummaryByRiskLevelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
