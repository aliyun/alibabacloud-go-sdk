// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateConfigRuleComplianceByPackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateConfigRuleComplianceByPackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateConfigRuleComplianceByPackResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateConfigRuleComplianceByPackResponseBody) *GetAggregateConfigRuleComplianceByPackResponse
	GetBody() *GetAggregateConfigRuleComplianceByPackResponseBody
}

type GetAggregateConfigRuleComplianceByPackResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateConfigRuleComplianceByPackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateConfigRuleComplianceByPackResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleComplianceByPackResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleComplianceByPackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateConfigRuleComplianceByPackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateConfigRuleComplianceByPackResponse) GetBody() *GetAggregateConfigRuleComplianceByPackResponseBody {
	return s.Body
}

func (s *GetAggregateConfigRuleComplianceByPackResponse) SetHeaders(v map[string]*string) *GetAggregateConfigRuleComplianceByPackResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponse) SetStatusCode(v int32) *GetAggregateConfigRuleComplianceByPackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponse) SetBody(v *GetAggregateConfigRuleComplianceByPackResponseBody) *GetAggregateConfigRuleComplianceByPackResponse {
	s.Body = v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
