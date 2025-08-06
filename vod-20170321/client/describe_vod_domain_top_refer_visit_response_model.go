// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainTopReferVisitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainTopReferVisitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainTopReferVisitResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainTopReferVisitResponseBody) *DescribeVodDomainTopReferVisitResponse
	GetBody() *DescribeVodDomainTopReferVisitResponseBody
}

type DescribeVodDomainTopReferVisitResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainTopReferVisitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainTopReferVisitResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainTopReferVisitResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTopReferVisitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainTopReferVisitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainTopReferVisitResponse) GetBody() *DescribeVodDomainTopReferVisitResponseBody {
	return s.Body
}

func (s *DescribeVodDomainTopReferVisitResponse) SetHeaders(v map[string]*string) *DescribeVodDomainTopReferVisitResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainTopReferVisitResponse) SetStatusCode(v int32) *DescribeVodDomainTopReferVisitResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainTopReferVisitResponse) SetBody(v *DescribeVodDomainTopReferVisitResponseBody) *DescribeVodDomainTopReferVisitResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainTopReferVisitResponse) Validate() error {
	return dara.Validate(s)
}
