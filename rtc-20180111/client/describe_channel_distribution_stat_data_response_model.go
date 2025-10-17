// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelDistributionStatDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChannelDistributionStatDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChannelDistributionStatDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChannelDistributionStatDataResponseBody) *DescribeChannelDistributionStatDataResponse
	GetBody() *DescribeChannelDistributionStatDataResponseBody
}

type DescribeChannelDistributionStatDataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelDistributionStatDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelDistributionStatDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChannelDistributionStatDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChannelDistributionStatDataResponse) GetBody() *DescribeChannelDistributionStatDataResponseBody {
	return s.Body
}

func (s *DescribeChannelDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeChannelDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelDistributionStatDataResponse) SetStatusCode(v int32) *DescribeChannelDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelDistributionStatDataResponse) SetBody(v *DescribeChannelDistributionStatDataResponseBody) *DescribeChannelDistributionStatDataResponse {
	s.Body = v
	return s
}

func (s *DescribeChannelDistributionStatDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
