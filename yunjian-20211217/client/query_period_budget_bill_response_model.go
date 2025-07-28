// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPeriodBudgetBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPeriodBudgetBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPeriodBudgetBillResponse
	GetStatusCode() *int32
	SetBody(v *QueryPeriodBudgetBillResponseBody) *QueryPeriodBudgetBillResponse
	GetBody() *QueryPeriodBudgetBillResponseBody
}

type QueryPeriodBudgetBillResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPeriodBudgetBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPeriodBudgetBillResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPeriodBudgetBillResponse) GoString() string {
	return s.String()
}

func (s *QueryPeriodBudgetBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPeriodBudgetBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPeriodBudgetBillResponse) GetBody() *QueryPeriodBudgetBillResponseBody {
	return s.Body
}

func (s *QueryPeriodBudgetBillResponse) SetHeaders(v map[string]*string) *QueryPeriodBudgetBillResponse {
	s.Headers = v
	return s
}

func (s *QueryPeriodBudgetBillResponse) SetStatusCode(v int32) *QueryPeriodBudgetBillResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPeriodBudgetBillResponse) SetBody(v *QueryPeriodBudgetBillResponseBody) *QueryPeriodBudgetBillResponse {
	s.Body = v
	return s
}

func (s *QueryPeriodBudgetBillResponse) Validate() error {
	return dara.Validate(s)
}
