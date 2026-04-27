// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryBillingCostBreakdownResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryBillingCostBreakdownResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryBillingCostBreakdownResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryBillingCostBreakdownResponseBody) *ModelRouterQueryBillingCostBreakdownResponse
	GetBody() *ModelRouterQueryBillingCostBreakdownResponseBody
}

type ModelRouterQueryBillingCostBreakdownResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryBillingCostBreakdownResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryBillingCostBreakdownResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryBillingCostBreakdownResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryBillingCostBreakdownResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryBillingCostBreakdownResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryBillingCostBreakdownResponse) GetBody() *ModelRouterQueryBillingCostBreakdownResponseBody {
	return s.Body
}

func (s *ModelRouterQueryBillingCostBreakdownResponse) SetHeaders(v map[string]*string) *ModelRouterQueryBillingCostBreakdownResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownResponse) SetStatusCode(v int32) *ModelRouterQueryBillingCostBreakdownResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownResponse) SetBody(v *ModelRouterQueryBillingCostBreakdownResponseBody) *ModelRouterQueryBillingCostBreakdownResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryBillingCostBreakdownResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
