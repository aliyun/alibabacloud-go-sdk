// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainLogResponseBody) *DescribeVodDomainLogResponse
	GetBody() *DescribeVodDomainLogResponseBody
}

type DescribeVodDomainLogResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainLogResponse) GetBody() *DescribeVodDomainLogResponseBody {
	return s.Body
}

func (s *DescribeVodDomainLogResponse) SetHeaders(v map[string]*string) *DescribeVodDomainLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainLogResponse) SetStatusCode(v int32) *DescribeVodDomainLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainLogResponse) SetBody(v *DescribeVodDomainLogResponseBody) *DescribeVodDomainLogResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainLogResponse) Validate() error {
	return dara.Validate(s)
}
