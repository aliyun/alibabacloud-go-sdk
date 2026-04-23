// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateBillingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterCreateBillingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterCreateBillingRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterCreateBillingRuleResponseBody) *ModelRouterCreateBillingRuleResponse
	GetBody() *ModelRouterCreateBillingRuleResponseBody
}

type ModelRouterCreateBillingRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterCreateBillingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterCreateBillingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateBillingRuleResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateBillingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterCreateBillingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterCreateBillingRuleResponse) GetBody() *ModelRouterCreateBillingRuleResponseBody {
	return s.Body
}

func (s *ModelRouterCreateBillingRuleResponse) SetHeaders(v map[string]*string) *ModelRouterCreateBillingRuleResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterCreateBillingRuleResponse) SetStatusCode(v int32) *ModelRouterCreateBillingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterCreateBillingRuleResponse) SetBody(v *ModelRouterCreateBillingRuleResponseBody) *ModelRouterCreateBillingRuleResponse {
	s.Body = v
	return s
}

func (s *ModelRouterCreateBillingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
