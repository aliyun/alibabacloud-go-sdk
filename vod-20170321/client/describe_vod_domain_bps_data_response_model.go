// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainBpsDataResponseBody) *DescribeVodDomainBpsDataResponse
	GetBody() *DescribeVodDomainBpsDataResponseBody
}

type DescribeVodDomainBpsDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainBpsDataResponse) GetBody() *DescribeVodDomainBpsDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainBpsDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainBpsDataResponse) SetStatusCode(v int32) *DescribeVodDomainBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponse) SetBody(v *DescribeVodDomainBpsDataResponseBody) *DescribeVodDomainBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainBpsDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
