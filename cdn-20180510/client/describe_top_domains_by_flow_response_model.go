// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTopDomainsByFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTopDomainsByFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTopDomainsByFlowResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTopDomainsByFlowResponseBody) *DescribeTopDomainsByFlowResponse
	GetBody() *DescribeTopDomainsByFlowResponseBody
}

type DescribeTopDomainsByFlowResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTopDomainsByFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTopDomainsByFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopDomainsByFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeTopDomainsByFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTopDomainsByFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTopDomainsByFlowResponse) GetBody() *DescribeTopDomainsByFlowResponseBody {
	return s.Body
}

func (s *DescribeTopDomainsByFlowResponse) SetHeaders(v map[string]*string) *DescribeTopDomainsByFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeTopDomainsByFlowResponse) SetStatusCode(v int32) *DescribeTopDomainsByFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponse) SetBody(v *DescribeTopDomainsByFlowResponseBody) *DescribeTopDomainsByFlowResponse {
	s.Body = v
	return s
}

func (s *DescribeTopDomainsByFlowResponse) Validate() error {
	return dara.Validate(s)
}
