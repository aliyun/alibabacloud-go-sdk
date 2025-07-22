// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelAreaDistributionStatDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChannelAreaDistributionStatDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChannelAreaDistributionStatDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChannelAreaDistributionStatDataResponseBody) *DescribeChannelAreaDistributionStatDataResponse
	GetBody() *DescribeChannelAreaDistributionStatDataResponseBody
}

type DescribeChannelAreaDistributionStatDataResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelAreaDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelAreaDistributionStatDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelAreaDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelAreaDistributionStatDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChannelAreaDistributionStatDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChannelAreaDistributionStatDataResponse) GetBody() *DescribeChannelAreaDistributionStatDataResponseBody {
	return s.Body
}

func (s *DescribeChannelAreaDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeChannelAreaDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponse) SetStatusCode(v int32) *DescribeChannelAreaDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponse) SetBody(v *DescribeChannelAreaDistributionStatDataResponseBody) *DescribeChannelAreaDistributionStatDataResponse {
	s.Body = v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponse) Validate() error {
	return dara.Validate(s)
}
