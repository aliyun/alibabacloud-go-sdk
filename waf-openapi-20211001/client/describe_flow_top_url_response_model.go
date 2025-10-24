// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowTopUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFlowTopUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFlowTopUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFlowTopUrlResponseBody) *DescribeFlowTopUrlResponse
	GetBody() *DescribeFlowTopUrlResponseBody
}

type DescribeFlowTopUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFlowTopUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFlowTopUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowTopUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFlowTopUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFlowTopUrlResponse) GetBody() *DescribeFlowTopUrlResponseBody {
	return s.Body
}

func (s *DescribeFlowTopUrlResponse) SetHeaders(v map[string]*string) *DescribeFlowTopUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowTopUrlResponse) SetStatusCode(v int32) *DescribeFlowTopUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowTopUrlResponse) SetBody(v *DescribeFlowTopUrlResponseBody) *DescribeFlowTopUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeFlowTopUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
