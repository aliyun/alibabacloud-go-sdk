// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodTopDomainsByFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodTopDomainsByFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodTopDomainsByFlowResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodTopDomainsByFlowResponseBody) *DescribeVodTopDomainsByFlowResponse
	GetBody() *DescribeVodTopDomainsByFlowResponseBody
}

type DescribeVodTopDomainsByFlowResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodTopDomainsByFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodTopDomainsByFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTopDomainsByFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodTopDomainsByFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodTopDomainsByFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodTopDomainsByFlowResponse) GetBody() *DescribeVodTopDomainsByFlowResponseBody {
	return s.Body
}

func (s *DescribeVodTopDomainsByFlowResponse) SetHeaders(v map[string]*string) *DescribeVodTopDomainsByFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodTopDomainsByFlowResponse) SetStatusCode(v int32) *DescribeVodTopDomainsByFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowResponse) SetBody(v *DescribeVodTopDomainsByFlowResponseBody) *DescribeVodTopDomainsByFlowResponse {
	s.Body = v
	return s
}

func (s *DescribeVodTopDomainsByFlowResponse) Validate() error {
	return dara.Validate(s)
}
