// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainTopUrlVisitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainTopUrlVisitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainTopUrlVisitResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainTopUrlVisitResponseBody) *DescribeVodDomainTopUrlVisitResponse
	GetBody() *DescribeVodDomainTopUrlVisitResponseBody
}

type DescribeVodDomainTopUrlVisitResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainTopUrlVisitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainTopUrlVisitResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopUrlVisitResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopUrlVisitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainTopUrlVisitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainTopUrlVisitResponse) GetBody() *DescribeVodDomainTopUrlVisitResponseBody {
	return s.Body
}

func (s *DescribeVodDomainTopUrlVisitResponse) SetHeaders(v map[string]*string) *DescribeVodDomainTopUrlVisitResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponse) SetStatusCode(v int32) *DescribeVodDomainTopUrlVisitResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponse) SetBody(v *DescribeVodDomainTopUrlVisitResponseBody) *DescribeVodDomainTopUrlVisitResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainTopUrlVisitResponse) Validate() error {
	return dara.Validate(s)
}
