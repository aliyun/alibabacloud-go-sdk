// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainReqHitRateDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainReqHitRateDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainReqHitRateDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainReqHitRateDataResponseBody) *DescribeVodDomainReqHitRateDataResponse
	GetBody() *DescribeVodDomainReqHitRateDataResponseBody
}

type DescribeVodDomainReqHitRateDataResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainReqHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainReqHitRateDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainReqHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainReqHitRateDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainReqHitRateDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainReqHitRateDataResponse) GetBody() *DescribeVodDomainReqHitRateDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainReqHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainReqHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainReqHitRateDataResponse) SetStatusCode(v int32) *DescribeVodDomainReqHitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainReqHitRateDataResponse) SetBody(v *DescribeVodDomainReqHitRateDataResponseBody) *DescribeVodDomainReqHitRateDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainReqHitRateDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
