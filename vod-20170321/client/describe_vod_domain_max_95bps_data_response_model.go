// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainMax95BpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainMax95BpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainMax95BpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainMax95BpsDataResponseBody) *DescribeVodDomainMax95BpsDataResponse
	GetBody() *DescribeVodDomainMax95BpsDataResponseBody
}

type DescribeVodDomainMax95BpsDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainMax95BpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainMax95BpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainMax95BpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainMax95BpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainMax95BpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainMax95BpsDataResponse) GetBody() *DescribeVodDomainMax95BpsDataResponseBody {
	return s.Body
}

func (s *DescribeVodDomainMax95BpsDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainMax95BpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainMax95BpsDataResponse) SetStatusCode(v int32) *DescribeVodDomainMax95BpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataResponse) SetBody(v *DescribeVodDomainMax95BpsDataResponseBody) *DescribeVodDomainMax95BpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainMax95BpsDataResponse) Validate() error {
	return dara.Validate(s)
}
