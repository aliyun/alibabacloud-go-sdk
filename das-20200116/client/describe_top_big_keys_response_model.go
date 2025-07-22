// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTopBigKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTopBigKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTopBigKeysResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTopBigKeysResponseBody) *DescribeTopBigKeysResponse
	GetBody() *DescribeTopBigKeysResponseBody
}

type DescribeTopBigKeysResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTopBigKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTopBigKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopBigKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribeTopBigKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTopBigKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTopBigKeysResponse) GetBody() *DescribeTopBigKeysResponseBody {
	return s.Body
}

func (s *DescribeTopBigKeysResponse) SetHeaders(v map[string]*string) *DescribeTopBigKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribeTopBigKeysResponse) SetStatusCode(v int32) *DescribeTopBigKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTopBigKeysResponse) SetBody(v *DescribeTopBigKeysResponseBody) *DescribeTopBigKeysResponse {
	s.Body = v
	return s
}

func (s *DescribeTopBigKeysResponse) Validate() error {
	return dara.Validate(s)
}
