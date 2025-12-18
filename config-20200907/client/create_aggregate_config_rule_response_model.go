// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregateConfigRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAggregateConfigRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAggregateConfigRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateAggregateConfigRuleResponseBody) *CreateAggregateConfigRuleResponse
	GetBody() *CreateAggregateConfigRuleResponseBody
}

type CreateAggregateConfigRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAggregateConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAggregateConfigRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAggregateConfigRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAggregateConfigRuleResponse) GetBody() *CreateAggregateConfigRuleResponseBody {
	return s.Body
}

func (s *CreateAggregateConfigRuleResponse) SetHeaders(v map[string]*string) *CreateAggregateConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAggregateConfigRuleResponse) SetStatusCode(v int32) *CreateAggregateConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAggregateConfigRuleResponse) SetBody(v *CreateAggregateConfigRuleResponseBody) *CreateAggregateConfigRuleResponse {
	s.Body = v
	return s
}

func (s *CreateAggregateConfigRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
