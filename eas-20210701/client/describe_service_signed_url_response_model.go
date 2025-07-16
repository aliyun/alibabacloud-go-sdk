// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceSignedUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceSignedUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceSignedUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceSignedUrlResponseBody) *DescribeServiceSignedUrlResponse
	GetBody() *DescribeServiceSignedUrlResponseBody
}

type DescribeServiceSignedUrlResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceSignedUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceSignedUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceSignedUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceSignedUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceSignedUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceSignedUrlResponse) GetBody() *DescribeServiceSignedUrlResponseBody {
	return s.Body
}

func (s *DescribeServiceSignedUrlResponse) SetHeaders(v map[string]*string) *DescribeServiceSignedUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceSignedUrlResponse) SetStatusCode(v int32) *DescribeServiceSignedUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceSignedUrlResponse) SetBody(v *DescribeServiceSignedUrlResponseBody) *DescribeServiceSignedUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceSignedUrlResponse) Validate() error {
	return dara.Validate(s)
}
