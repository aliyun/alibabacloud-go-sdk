// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBudgetUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBudgetUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBudgetUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBudgetUsageResponseBody) *DescribeBudgetUsageResponse
	GetBody() *DescribeBudgetUsageResponseBody
}

type DescribeBudgetUsageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBudgetUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBudgetUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeBudgetUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBudgetUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBudgetUsageResponse) GetBody() *DescribeBudgetUsageResponseBody {
	return s.Body
}

func (s *DescribeBudgetUsageResponse) SetHeaders(v map[string]*string) *DescribeBudgetUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeBudgetUsageResponse) SetStatusCode(v int32) *DescribeBudgetUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBudgetUsageResponse) SetBody(v *DescribeBudgetUsageResponseBody) *DescribeBudgetUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeBudgetUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
