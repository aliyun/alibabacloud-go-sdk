// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainRealTimeBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainRealTimeBpsDataResponseBody) *DescribeVodDomainRealTimeBpsDataResponse
	GetBody() *DescribeVodDomainRealTimeBpsDataResponseBody
}

type DescribeVodDomainRealTimeBpsDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainRealTimeBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainRealTimeBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainRealTimeBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainRealTimeBpsDataResponse) GetBody() *DescribeVodDomainRealTimeBpsDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainRealTimeBpsDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainRealTimeBpsDataResponse) SetStatusCode(v int32) *DescribeVodDomainRealTimeBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainRealTimeBpsDataResponse) SetBody(v *DescribeVodDomainRealTimeBpsDataResponseBody) *DescribeVodDomainRealTimeBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainRealTimeBpsDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
