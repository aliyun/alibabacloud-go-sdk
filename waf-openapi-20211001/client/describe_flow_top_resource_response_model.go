// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowTopResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFlowTopResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFlowTopResourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFlowTopResourceResponseBody) *DescribeFlowTopResourceResponse
	GetBody() *DescribeFlowTopResourceResponseBody
}

type DescribeFlowTopResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFlowTopResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFlowTopResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowTopResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFlowTopResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFlowTopResourceResponse) GetBody() *DescribeFlowTopResourceResponseBody {
	return s.Body
}

func (s *DescribeFlowTopResourceResponse) SetHeaders(v map[string]*string) *DescribeFlowTopResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowTopResourceResponse) SetStatusCode(v int32) *DescribeFlowTopResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowTopResourceResponse) SetBody(v *DescribeFlowTopResourceResponseBody) *DescribeFlowTopResourceResponse {
	s.Body = v
	return s
}

func (s *DescribeFlowTopResourceResponse) Validate() error {
	return dara.Validate(s)
}
