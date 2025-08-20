// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceGroupSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceGroupSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceGroupSpecResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceGroupSpecResponseBody) *DescribeResourceGroupSpecResponse
	GetBody() *DescribeResourceGroupSpecResponseBody
}

type DescribeResourceGroupSpecResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceGroupSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceGroupSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceGroupSpecResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceGroupSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceGroupSpecResponse) GetBody() *DescribeResourceGroupSpecResponseBody {
	return s.Body
}

func (s *DescribeResourceGroupSpecResponse) SetHeaders(v map[string]*string) *DescribeResourceGroupSpecResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceGroupSpecResponse) SetStatusCode(v int32) *DescribeResourceGroupSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceGroupSpecResponse) SetBody(v *DescribeResourceGroupSpecResponseBody) *DescribeResourceGroupSpecResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceGroupSpecResponse) Validate() error {
	return dara.Validate(s)
}
