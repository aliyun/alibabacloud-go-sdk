// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBudgetPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBudgetPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBudgetPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBudgetPoliciesResponseBody) *DescribeBudgetPoliciesResponse
	GetBody() *DescribeBudgetPoliciesResponseBody
}

type DescribeBudgetPoliciesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBudgetPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBudgetPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetPoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBudgetPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBudgetPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBudgetPoliciesResponse) GetBody() *DescribeBudgetPoliciesResponseBody {
	return s.Body
}

func (s *DescribeBudgetPoliciesResponse) SetHeaders(v map[string]*string) *DescribeBudgetPoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBudgetPoliciesResponse) SetStatusCode(v int32) *DescribeBudgetPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBudgetPoliciesResponse) SetBody(v *DescribeBudgetPoliciesResponseBody) *DescribeBudgetPoliciesResponse {
	s.Body = v
	return s
}

func (s *DescribeBudgetPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
