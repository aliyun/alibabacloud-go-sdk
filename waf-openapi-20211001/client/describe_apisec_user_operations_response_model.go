// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecUserOperationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecUserOperationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecUserOperationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecUserOperationsResponseBody) *DescribeApisecUserOperationsResponse
	GetBody() *DescribeApisecUserOperationsResponseBody
}

type DescribeApisecUserOperationsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecUserOperationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecUserOperationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecUserOperationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecUserOperationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecUserOperationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecUserOperationsResponse) GetBody() *DescribeApisecUserOperationsResponseBody {
	return s.Body
}

func (s *DescribeApisecUserOperationsResponse) SetHeaders(v map[string]*string) *DescribeApisecUserOperationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecUserOperationsResponse) SetStatusCode(v int32) *DescribeApisecUserOperationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecUserOperationsResponse) SetBody(v *DescribeApisecUserOperationsResponseBody) *DescribeApisecUserOperationsResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecUserOperationsResponse) Validate() error {
	return dara.Validate(s)
}
