// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageScanAuthCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageScanAuthCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageScanAuthCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageScanAuthCountResponseBody) *DescribeImageScanAuthCountResponse
	GetBody() *DescribeImageScanAuthCountResponseBody
}

type DescribeImageScanAuthCountResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageScanAuthCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageScanAuthCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageScanAuthCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageScanAuthCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageScanAuthCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageScanAuthCountResponse) GetBody() *DescribeImageScanAuthCountResponseBody {
	return s.Body
}

func (s *DescribeImageScanAuthCountResponse) SetHeaders(v map[string]*string) *DescribeImageScanAuthCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageScanAuthCountResponse) SetStatusCode(v int32) *DescribeImageScanAuthCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageScanAuthCountResponse) SetBody(v *DescribeImageScanAuthCountResponseBody) *DescribeImageScanAuthCountResponse {
	s.Body = v
	return s
}

func (s *DescribeImageScanAuthCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
