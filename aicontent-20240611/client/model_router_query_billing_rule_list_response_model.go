// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryBillingRuleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryBillingRuleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryBillingRuleListResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryBillingRuleListResponseBody) *ModelRouterQueryBillingRuleListResponse
	GetBody() *ModelRouterQueryBillingRuleListResponseBody
}

type ModelRouterQueryBillingRuleListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryBillingRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryBillingRuleListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryBillingRuleListResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryBillingRuleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryBillingRuleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryBillingRuleListResponse) GetBody() *ModelRouterQueryBillingRuleListResponseBody {
	return s.Body
}

func (s *ModelRouterQueryBillingRuleListResponse) SetHeaders(v map[string]*string) *ModelRouterQueryBillingRuleListResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponse) SetStatusCode(v int32) *ModelRouterQueryBillingRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponse) SetBody(v *ModelRouterQueryBillingRuleListResponseBody) *ModelRouterQueryBillingRuleListResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryBillingRuleListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
