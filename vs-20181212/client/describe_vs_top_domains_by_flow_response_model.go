// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsTopDomainsByFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsTopDomainsByFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsTopDomainsByFlowResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsTopDomainsByFlowResponseBody) *DescribeVsTopDomainsByFlowResponse
	GetBody() *DescribeVsTopDomainsByFlowResponseBody
}

type DescribeVsTopDomainsByFlowResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsTopDomainsByFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsTopDomainsByFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsTopDomainsByFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsTopDomainsByFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsTopDomainsByFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsTopDomainsByFlowResponse) GetBody() *DescribeVsTopDomainsByFlowResponseBody {
	return s.Body
}

func (s *DescribeVsTopDomainsByFlowResponse) SetHeaders(v map[string]*string) *DescribeVsTopDomainsByFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponse) SetStatusCode(v int32) *DescribeVsTopDomainsByFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponse) SetBody(v *DescribeVsTopDomainsByFlowResponseBody) *DescribeVsTopDomainsByFlowResponse {
	s.Body = v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponse) Validate() error {
	return dara.Validate(s)
}
