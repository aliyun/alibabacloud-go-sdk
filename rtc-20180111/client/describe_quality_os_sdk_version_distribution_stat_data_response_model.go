// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQualityOsSdkVersionDistributionStatDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeQualityOsSdkVersionDistributionStatDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeQualityOsSdkVersionDistributionStatDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeQualityOsSdkVersionDistributionStatDataResponseBody) *DescribeQualityOsSdkVersionDistributionStatDataResponse
	GetBody() *DescribeQualityOsSdkVersionDistributionStatDataResponseBody
}

type DescribeQualityOsSdkVersionDistributionStatDataResponse struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQualityOsSdkVersionDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponse) GetBody() *DescribeQualityOsSdkVersionDistributionStatDataResponseBody {
	return s.Body
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeQualityOsSdkVersionDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponse) SetStatusCode(v int32) *DescribeQualityOsSdkVersionDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponse) SetBody(v *DescribeQualityOsSdkVersionDistributionStatDataResponseBody) *DescribeQualityOsSdkVersionDistributionStatDataResponse {
	s.Body = v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponse) Validate() error {
	return dara.Validate(s)
}
