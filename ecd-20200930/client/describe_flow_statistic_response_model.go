// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFlowStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFlowStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFlowStatisticResponseBody) *DescribeFlowStatisticResponse
	GetBody() *DescribeFlowStatisticResponseBody
}

type DescribeFlowStatisticResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFlowStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFlowStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFlowStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFlowStatisticResponse) GetBody() *DescribeFlowStatisticResponseBody {
	return s.Body
}

func (s *DescribeFlowStatisticResponse) SetHeaders(v map[string]*string) *DescribeFlowStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowStatisticResponse) SetStatusCode(v int32) *DescribeFlowStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowStatisticResponse) SetBody(v *DescribeFlowStatisticResponseBody) *DescribeFlowStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeFlowStatisticResponse) Validate() error {
	return dara.Validate(s)
}
