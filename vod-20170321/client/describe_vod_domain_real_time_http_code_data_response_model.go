// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeHttpCodeDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeHttpCodeDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainRealTimeHttpCodeDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainRealTimeHttpCodeDataResponseBody) *DescribeVodDomainRealTimeHttpCodeDataResponse
	GetBody() *DescribeVodDomainRealTimeHttpCodeDataResponseBody
}

type DescribeVodDomainRealTimeHttpCodeDataResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainRealTimeHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainRealTimeHttpCodeDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeHttpCodeDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainRealTimeHttpCodeDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainRealTimeHttpCodeDataResponse) GetBody() *DescribeVodDomainRealTimeHttpCodeDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainRealTimeHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainRealTimeHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainRealTimeHttpCodeDataResponse) SetStatusCode(v int32) *DescribeVodDomainRealTimeHttpCodeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainRealTimeHttpCodeDataResponse) SetBody(v *DescribeVodDomainRealTimeHttpCodeDataResponseBody) *DescribeVodDomainRealTimeHttpCodeDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainRealTimeHttpCodeDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
