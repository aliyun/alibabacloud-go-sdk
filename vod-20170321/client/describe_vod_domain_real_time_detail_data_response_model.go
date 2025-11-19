// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeDetailDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeDetailDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainRealTimeDetailDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainRealTimeDetailDataResponseBody) *DescribeVodDomainRealTimeDetailDataResponse
	GetBody() *DescribeVodDomainRealTimeDetailDataResponseBody
}

type DescribeVodDomainRealTimeDetailDataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainRealTimeDetailDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainRealTimeDetailDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeDetailDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeDetailDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainRealTimeDetailDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainRealTimeDetailDataResponse) GetBody() *DescribeVodDomainRealTimeDetailDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainRealTimeDetailDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeDetailDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainRealTimeDetailDataResponse) SetStatusCode(v int32) *DescribeVodDomainRealTimeDetailDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainRealTimeDetailDataResponse) SetBody(v *DescribeVodDomainRealTimeDetailDataResponseBody) *DescribeVodDomainRealTimeDetailDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainRealTimeDetailDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
