// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSpecificationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceSpecificationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceSpecificationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceSpecificationsResponseBody) *DescribeInstanceSpecificationsResponse
	GetBody() *DescribeInstanceSpecificationsResponseBody
}

type DescribeInstanceSpecificationsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceSpecificationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceSpecificationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSpecificationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecificationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceSpecificationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceSpecificationsResponse) GetBody() *DescribeInstanceSpecificationsResponseBody {
	return s.Body
}

func (s *DescribeInstanceSpecificationsResponse) SetHeaders(v map[string]*string) *DescribeInstanceSpecificationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSpecificationsResponse) SetStatusCode(v int32) *DescribeInstanceSpecificationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponse) SetBody(v *DescribeInstanceSpecificationsResponseBody) *DescribeInstanceSpecificationsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceSpecificationsResponse) Validate() error {
	return dara.Validate(s)
}
