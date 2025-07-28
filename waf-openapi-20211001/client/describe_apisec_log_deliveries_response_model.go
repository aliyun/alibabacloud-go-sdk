// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecLogDeliveriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecLogDeliveriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecLogDeliveriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecLogDeliveriesResponseBody) *DescribeApisecLogDeliveriesResponse
	GetBody() *DescribeApisecLogDeliveriesResponseBody
}

type DescribeApisecLogDeliveriesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecLogDeliveriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecLogDeliveriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecLogDeliveriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecLogDeliveriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecLogDeliveriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecLogDeliveriesResponse) GetBody() *DescribeApisecLogDeliveriesResponseBody {
	return s.Body
}

func (s *DescribeApisecLogDeliveriesResponse) SetHeaders(v map[string]*string) *DescribeApisecLogDeliveriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecLogDeliveriesResponse) SetStatusCode(v int32) *DescribeApisecLogDeliveriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecLogDeliveriesResponse) SetBody(v *DescribeApisecLogDeliveriesResponseBody) *DescribeApisecLogDeliveriesResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecLogDeliveriesResponse) Validate() error {
	return dara.Validate(s)
}
