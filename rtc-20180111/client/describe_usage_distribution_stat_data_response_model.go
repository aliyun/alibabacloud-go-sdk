// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsageDistributionStatDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUsageDistributionStatDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUsageDistributionStatDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUsageDistributionStatDataResponseBody) *DescribeUsageDistributionStatDataResponse
	GetBody() *DescribeUsageDistributionStatDataResponseBody
}

type DescribeUsageDistributionStatDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsageDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsageDistributionStatDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsageDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsageDistributionStatDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUsageDistributionStatDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUsageDistributionStatDataResponse) GetBody() *DescribeUsageDistributionStatDataResponseBody {
	return s.Body
}

func (s *DescribeUsageDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeUsageDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsageDistributionStatDataResponse) SetStatusCode(v int32) *DescribeUsageDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponse) SetBody(v *DescribeUsageDistributionStatDataResponseBody) *DescribeUsageDistributionStatDataResponse {
	s.Body = v
	return s
}

func (s *DescribeUsageDistributionStatDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
