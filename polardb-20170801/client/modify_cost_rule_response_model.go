// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCostRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCostRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCostRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCostRuleResponseBody) *ModifyCostRuleResponse
	GetBody() *ModifyCostRuleResponseBody
}

type ModifyCostRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCostRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCostRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyCostRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCostRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCostRuleResponse) GetBody() *ModifyCostRuleResponseBody {
	return s.Body
}

func (s *ModifyCostRuleResponse) SetHeaders(v map[string]*string) *ModifyCostRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyCostRuleResponse) SetStatusCode(v int32) *ModifyCostRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCostRuleResponse) SetBody(v *ModifyCostRuleResponseBody) *ModifyCostRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyCostRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
