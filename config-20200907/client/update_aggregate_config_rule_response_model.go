// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregateConfigRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAggregateConfigRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAggregateConfigRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAggregateConfigRuleResponseBody) *UpdateAggregateConfigRuleResponse
	GetBody() *UpdateAggregateConfigRuleResponseBody
}

type UpdateAggregateConfigRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAggregateConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAggregateConfigRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAggregateConfigRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAggregateConfigRuleResponse) GetBody() *UpdateAggregateConfigRuleResponseBody {
	return s.Body
}

func (s *UpdateAggregateConfigRuleResponse) SetHeaders(v map[string]*string) *UpdateAggregateConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateAggregateConfigRuleResponse) SetStatusCode(v int32) *UpdateAggregateConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAggregateConfigRuleResponse) SetBody(v *UpdateAggregateConfigRuleResponseBody) *UpdateAggregateConfigRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateAggregateConfigRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
