// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBillingCostTabsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterBillingCostTabsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterBillingCostTabsResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterBillingCostTabsResponseBody) *ModelRouterBillingCostTabsResponse
	GetBody() *ModelRouterBillingCostTabsResponseBody
}

type ModelRouterBillingCostTabsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterBillingCostTabsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterBillingCostTabsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBillingCostTabsResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterBillingCostTabsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterBillingCostTabsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterBillingCostTabsResponse) GetBody() *ModelRouterBillingCostTabsResponseBody {
	return s.Body
}

func (s *ModelRouterBillingCostTabsResponse) SetHeaders(v map[string]*string) *ModelRouterBillingCostTabsResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterBillingCostTabsResponse) SetStatusCode(v int32) *ModelRouterBillingCostTabsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterBillingCostTabsResponse) SetBody(v *ModelRouterBillingCostTabsResponseBody) *ModelRouterBillingCostTabsResponse {
	s.Body = v
	return s
}

func (s *ModelRouterBillingCostTabsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
