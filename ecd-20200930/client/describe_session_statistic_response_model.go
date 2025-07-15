// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSessionStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSessionStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSessionStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSessionStatisticResponseBody) *DescribeSessionStatisticResponse
	GetBody() *DescribeSessionStatisticResponseBody
}

type DescribeSessionStatisticResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSessionStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSessionStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSessionStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeSessionStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSessionStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSessionStatisticResponse) GetBody() *DescribeSessionStatisticResponseBody {
	return s.Body
}

func (s *DescribeSessionStatisticResponse) SetHeaders(v map[string]*string) *DescribeSessionStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeSessionStatisticResponse) SetStatusCode(v int32) *DescribeSessionStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSessionStatisticResponse) SetBody(v *DescribeSessionStatisticResponseBody) *DescribeSessionStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeSessionStatisticResponse) Validate() error {
	return dara.Validate(s)
}
