// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCollectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCollectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCollectorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCollectorResponseBody) *DescribeCollectorResponse
	GetBody() *DescribeCollectorResponseBody
}

type DescribeCollectorResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCollectorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCollectorResponse) GoString() string {
	return s.String()
}

func (s *DescribeCollectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCollectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCollectorResponse) GetBody() *DescribeCollectorResponseBody {
	return s.Body
}

func (s *DescribeCollectorResponse) SetHeaders(v map[string]*string) *DescribeCollectorResponse {
	s.Headers = v
	return s
}

func (s *DescribeCollectorResponse) SetStatusCode(v int32) *DescribeCollectorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCollectorResponse) SetBody(v *DescribeCollectorResponseBody) *DescribeCollectorResponse {
	s.Body = v
	return s
}

func (s *DescribeCollectorResponse) Validate() error {
	return dara.Validate(s)
}
