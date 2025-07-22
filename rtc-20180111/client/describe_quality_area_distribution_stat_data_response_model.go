// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQualityAreaDistributionStatDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeQualityAreaDistributionStatDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeQualityAreaDistributionStatDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeQualityAreaDistributionStatDataResponseBody) *DescribeQualityAreaDistributionStatDataResponse
	GetBody() *DescribeQualityAreaDistributionStatDataResponseBody
}

type DescribeQualityAreaDistributionStatDataResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQualityAreaDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQualityAreaDistributionStatDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityAreaDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeQualityAreaDistributionStatDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeQualityAreaDistributionStatDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeQualityAreaDistributionStatDataResponse) GetBody() *DescribeQualityAreaDistributionStatDataResponseBody {
	return s.Body
}

func (s *DescribeQualityAreaDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeQualityAreaDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponse) SetStatusCode(v int32) *DescribeQualityAreaDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponse) SetBody(v *DescribeQualityAreaDistributionStatDataResponseBody) *DescribeQualityAreaDistributionStatDataResponse {
	s.Body = v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponse) Validate() error {
	return dara.Validate(s)
}
