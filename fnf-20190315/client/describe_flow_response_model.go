// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFlowResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFlowResponseBody) *DescribeFlowResponse
	GetBody() *DescribeFlowResponseBody
}

type DescribeFlowResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFlowResponse) GetBody() *DescribeFlowResponseBody {
	return s.Body
}

func (s *DescribeFlowResponse) SetHeaders(v map[string]*string) *DescribeFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowResponse) SetStatusCode(v int32) *DescribeFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowResponse) SetBody(v *DescribeFlowResponseBody) *DescribeFlowResponse {
	s.Body = v
	return s
}

func (s *DescribeFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
