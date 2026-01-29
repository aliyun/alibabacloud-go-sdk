// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceBillResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceBillResponseBody) *DescribeInstanceBillResponse
	GetBody() *DescribeInstanceBillResponseBody
}

type DescribeInstanceBillResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceBillResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceBillResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceBillResponse) GetBody() *DescribeInstanceBillResponseBody {
	return s.Body
}

func (s *DescribeInstanceBillResponse) SetHeaders(v map[string]*string) *DescribeInstanceBillResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceBillResponse) SetStatusCode(v int32) *DescribeInstanceBillResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceBillResponse) SetBody(v *DescribeInstanceBillResponseBody) *DescribeInstanceBillResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceBillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
