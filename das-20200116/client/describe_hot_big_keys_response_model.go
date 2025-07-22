// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHotBigKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHotBigKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHotBigKeysResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHotBigKeysResponseBody) *DescribeHotBigKeysResponse
	GetBody() *DescribeHotBigKeysResponseBody
}

type DescribeHotBigKeysResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHotBigKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHotBigKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotBigKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHotBigKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHotBigKeysResponse) GetBody() *DescribeHotBigKeysResponseBody {
	return s.Body
}

func (s *DescribeHotBigKeysResponse) SetHeaders(v map[string]*string) *DescribeHotBigKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribeHotBigKeysResponse) SetStatusCode(v int32) *DescribeHotBigKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHotBigKeysResponse) SetBody(v *DescribeHotBigKeysResponseBody) *DescribeHotBigKeysResponse {
	s.Body = v
	return s
}

func (s *DescribeHotBigKeysResponse) Validate() error {
	return dara.Validate(s)
}
