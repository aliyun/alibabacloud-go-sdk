// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostBudgetsSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCostBudgetsSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCostBudgetsSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCostBudgetsSummaryResponseBody) *DescribeCostBudgetsSummaryResponse
	GetBody() *DescribeCostBudgetsSummaryResponseBody
}

type DescribeCostBudgetsSummaryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCostBudgetsSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCostBudgetsSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostBudgetsSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeCostBudgetsSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCostBudgetsSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCostBudgetsSummaryResponse) GetBody() *DescribeCostBudgetsSummaryResponseBody {
	return s.Body
}

func (s *DescribeCostBudgetsSummaryResponse) SetHeaders(v map[string]*string) *DescribeCostBudgetsSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeCostBudgetsSummaryResponse) SetStatusCode(v int32) *DescribeCostBudgetsSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCostBudgetsSummaryResponse) SetBody(v *DescribeCostBudgetsSummaryResponseBody) *DescribeCostBudgetsSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeCostBudgetsSummaryResponse) Validate() error {
	return dara.Validate(s)
}
