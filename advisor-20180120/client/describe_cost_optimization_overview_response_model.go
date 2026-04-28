// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostOptimizationOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCostOptimizationOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCostOptimizationOverviewResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCostOptimizationOverviewResponseBody) *DescribeCostOptimizationOverviewResponse
	GetBody() *DescribeCostOptimizationOverviewResponseBody
}

type DescribeCostOptimizationOverviewResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCostOptimizationOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCostOptimizationOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostOptimizationOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeCostOptimizationOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCostOptimizationOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCostOptimizationOverviewResponse) GetBody() *DescribeCostOptimizationOverviewResponseBody {
	return s.Body
}

func (s *DescribeCostOptimizationOverviewResponse) SetHeaders(v map[string]*string) *DescribeCostOptimizationOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribeCostOptimizationOverviewResponse) SetStatusCode(v int32) *DescribeCostOptimizationOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponse) SetBody(v *DescribeCostOptimizationOverviewResponseBody) *DescribeCostOptimizationOverviewResponse {
	s.Body = v
	return s
}

func (s *DescribeCostOptimizationOverviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
