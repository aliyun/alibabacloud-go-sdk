// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientRuleStrategyStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClientRuleStrategyStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClientRuleStrategyStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetClientRuleStrategyStatisticsResponseBody) *GetClientRuleStrategyStatisticsResponse
	GetBody() *GetClientRuleStrategyStatisticsResponseBody
}

type GetClientRuleStrategyStatisticsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClientRuleStrategyStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClientRuleStrategyStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClientRuleStrategyStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetClientRuleStrategyStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClientRuleStrategyStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClientRuleStrategyStatisticsResponse) GetBody() *GetClientRuleStrategyStatisticsResponseBody {
	return s.Body
}

func (s *GetClientRuleStrategyStatisticsResponse) SetHeaders(v map[string]*string) *GetClientRuleStrategyStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetClientRuleStrategyStatisticsResponse) SetStatusCode(v int32) *GetClientRuleStrategyStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClientRuleStrategyStatisticsResponse) SetBody(v *GetClientRuleStrategyStatisticsResponseBody) *GetClientRuleStrategyStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetClientRuleStrategyStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
