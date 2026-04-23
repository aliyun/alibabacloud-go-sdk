// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateBillingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterUpdateBillingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterUpdateBillingRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterUpdateBillingRuleResponseBody) *ModelRouterUpdateBillingRuleResponse
	GetBody() *ModelRouterUpdateBillingRuleResponseBody
}

type ModelRouterUpdateBillingRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterUpdateBillingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterUpdateBillingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateBillingRuleResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateBillingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterUpdateBillingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterUpdateBillingRuleResponse) GetBody() *ModelRouterUpdateBillingRuleResponseBody {
	return s.Body
}

func (s *ModelRouterUpdateBillingRuleResponse) SetHeaders(v map[string]*string) *ModelRouterUpdateBillingRuleResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponse) SetStatusCode(v int32) *ModelRouterUpdateBillingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponse) SetBody(v *ModelRouterUpdateBillingRuleResponseBody) *ModelRouterUpdateBillingRuleResponse {
	s.Body = v
	return s
}

func (s *ModelRouterUpdateBillingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
