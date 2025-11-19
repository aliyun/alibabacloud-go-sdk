// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainHitRateDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainHitRateDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainHitRateDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainHitRateDataResponseBody) *DescribeVodDomainHitRateDataResponse
	GetBody() *DescribeVodDomainHitRateDataResponseBody
}

type DescribeVodDomainHitRateDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainHitRateDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainHitRateDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainHitRateDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainHitRateDataResponse) GetBody() *DescribeVodDomainHitRateDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainHitRateDataResponse) SetStatusCode(v int32) *DescribeVodDomainHitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainHitRateDataResponse) SetBody(v *DescribeVodDomainHitRateDataResponseBody) *DescribeVodDomainHitRateDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainHitRateDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
