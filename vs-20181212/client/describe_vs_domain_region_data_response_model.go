// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainRegionDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsDomainRegionDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsDomainRegionDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsDomainRegionDataResponseBody) *DescribeVsDomainRegionDataResponse
	GetBody() *DescribeVsDomainRegionDataResponseBody
}

type DescribeVsDomainRegionDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsDomainRegionDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsDomainRegionDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainRegionDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRegionDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsDomainRegionDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsDomainRegionDataResponse) GetBody() *DescribeVsDomainRegionDataResponseBody {
	return s.Body
}

func (s *DescribeVsDomainRegionDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainRegionDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainRegionDataResponse) SetStatusCode(v int32) *DescribeVsDomainRegionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponse) SetBody(v *DescribeVsDomainRegionDataResponseBody) *DescribeVsDomainRegionDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVsDomainRegionDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
