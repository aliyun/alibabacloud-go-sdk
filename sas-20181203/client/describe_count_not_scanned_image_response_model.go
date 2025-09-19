// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCountNotScannedImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCountNotScannedImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCountNotScannedImageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCountNotScannedImageResponseBody) *DescribeCountNotScannedImageResponse
	GetBody() *DescribeCountNotScannedImageResponseBody
}

type DescribeCountNotScannedImageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCountNotScannedImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCountNotScannedImageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCountNotScannedImageResponse) GoString() string {
	return s.String()
}

func (s *DescribeCountNotScannedImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCountNotScannedImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCountNotScannedImageResponse) GetBody() *DescribeCountNotScannedImageResponseBody {
	return s.Body
}

func (s *DescribeCountNotScannedImageResponse) SetHeaders(v map[string]*string) *DescribeCountNotScannedImageResponse {
	s.Headers = v
	return s
}

func (s *DescribeCountNotScannedImageResponse) SetStatusCode(v int32) *DescribeCountNotScannedImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCountNotScannedImageResponse) SetBody(v *DescribeCountNotScannedImageResponseBody) *DescribeCountNotScannedImageResponse {
	s.Body = v
	return s
}

func (s *DescribeCountNotScannedImageResponse) Validate() error {
	return dara.Validate(s)
}
