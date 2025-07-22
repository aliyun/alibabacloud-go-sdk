// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQualityDistributionStatDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeQualityDistributionStatDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeQualityDistributionStatDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeQualityDistributionStatDataResponseBody) *DescribeQualityDistributionStatDataResponse
	GetBody() *DescribeQualityDistributionStatDataResponseBody
}

type DescribeQualityDistributionStatDataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQualityDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQualityDistributionStatDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeQualityDistributionStatDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeQualityDistributionStatDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeQualityDistributionStatDataResponse) GetBody() *DescribeQualityDistributionStatDataResponseBody {
	return s.Body
}

func (s *DescribeQualityDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeQualityDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeQualityDistributionStatDataResponse) SetStatusCode(v int32) *DescribeQualityDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponse) SetBody(v *DescribeQualityDistributionStatDataResponseBody) *DescribeQualityDistributionStatDataResponse {
	s.Body = v
	return s
}

func (s *DescribeQualityDistributionStatDataResponse) Validate() error {
	return dara.Validate(s)
}
