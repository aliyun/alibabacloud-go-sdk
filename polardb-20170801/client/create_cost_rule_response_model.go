// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCostRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCostRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCostRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateCostRuleResponseBody) *CreateCostRuleResponse
	GetBody() *CreateCostRuleResponseBody
}

type CreateCostRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCostRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCostRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCostRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateCostRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCostRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCostRuleResponse) GetBody() *CreateCostRuleResponseBody {
	return s.Body
}

func (s *CreateCostRuleResponse) SetHeaders(v map[string]*string) *CreateCostRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateCostRuleResponse) SetStatusCode(v int32) *CreateCostRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCostRuleResponse) SetBody(v *CreateCostRuleResponseBody) *CreateCostRuleResponse {
	s.Body = v
	return s
}

func (s *CreateCostRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
