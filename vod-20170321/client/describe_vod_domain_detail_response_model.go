// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainDetailResponseBody) *DescribeVodDomainDetailResponse
	GetBody() *DescribeVodDomainDetailResponseBody
}

type DescribeVodDomainDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainDetailResponse) GetBody() *DescribeVodDomainDetailResponseBody {
	return s.Body
}

func (s *DescribeVodDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeVodDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainDetailResponse) SetStatusCode(v int32) *DescribeVodDomainDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainDetailResponse) SetBody(v *DescribeVodDomainDetailResponseBody) *DescribeVodDomainDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
