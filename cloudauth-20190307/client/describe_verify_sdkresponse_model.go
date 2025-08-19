// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifySDKResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVerifySDKResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVerifySDKResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVerifySDKResponseBody) *DescribeVerifySDKResponse
	GetBody() *DescribeVerifySDKResponseBody
}

type DescribeVerifySDKResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVerifySDKResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVerifySDKResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifySDKResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifySDKResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVerifySDKResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVerifySDKResponse) GetBody() *DescribeVerifySDKResponseBody {
	return s.Body
}

func (s *DescribeVerifySDKResponse) SetHeaders(v map[string]*string) *DescribeVerifySDKResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifySDKResponse) SetStatusCode(v int32) *DescribeVerifySDKResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifySDKResponse) SetBody(v *DescribeVerifySDKResponseBody) *DescribeVerifySDKResponse {
	s.Body = v
	return s
}

func (s *DescribeVerifySDKResponse) Validate() error {
	return dara.Validate(s)
}
