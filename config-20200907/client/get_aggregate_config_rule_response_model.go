// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateConfigRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateConfigRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateConfigRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateConfigRuleResponseBody) *GetAggregateConfigRuleResponse
	GetBody() *GetAggregateConfigRuleResponseBody
}

type GetAggregateConfigRuleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateConfigRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateConfigRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateConfigRuleResponse) GetBody() *GetAggregateConfigRuleResponseBody {
	return s.Body
}

func (s *GetAggregateConfigRuleResponse) SetHeaders(v map[string]*string) *GetAggregateConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateConfigRuleResponse) SetStatusCode(v int32) *GetAggregateConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateConfigRuleResponse) SetBody(v *GetAggregateConfigRuleResponseBody) *GetAggregateConfigRuleResponse {
	s.Body = v
	return s
}

func (s *GetAggregateConfigRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
