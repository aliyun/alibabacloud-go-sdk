// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetBillingBillSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterGetBillingBillSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterGetBillingBillSummaryResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterGetBillingBillSummaryResponseBody) *ModelRouterGetBillingBillSummaryResponse
	GetBody() *ModelRouterGetBillingBillSummaryResponseBody
}

type ModelRouterGetBillingBillSummaryResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterGetBillingBillSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterGetBillingBillSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetBillingBillSummaryResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterGetBillingBillSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterGetBillingBillSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterGetBillingBillSummaryResponse) GetBody() *ModelRouterGetBillingBillSummaryResponseBody {
	return s.Body
}

func (s *ModelRouterGetBillingBillSummaryResponse) SetHeaders(v map[string]*string) *ModelRouterGetBillingBillSummaryResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterGetBillingBillSummaryResponse) SetStatusCode(v int32) *ModelRouterGetBillingBillSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterGetBillingBillSummaryResponse) SetBody(v *ModelRouterGetBillingBillSummaryResponseBody) *ModelRouterGetBillingBillSummaryResponse {
	s.Body = v
	return s
}

func (s *ModelRouterGetBillingBillSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
