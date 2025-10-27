// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSecurityScanCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageSecurityScanCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageSecurityScanCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageSecurityScanCountResponseBody) *DescribeImageSecurityScanCountResponse
	GetBody() *DescribeImageSecurityScanCountResponseBody
}

type DescribeImageSecurityScanCountResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageSecurityScanCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageSecurityScanCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSecurityScanCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageSecurityScanCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageSecurityScanCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageSecurityScanCountResponse) GetBody() *DescribeImageSecurityScanCountResponseBody {
	return s.Body
}

func (s *DescribeImageSecurityScanCountResponse) SetHeaders(v map[string]*string) *DescribeImageSecurityScanCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageSecurityScanCountResponse) SetStatusCode(v int32) *DescribeImageSecurityScanCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageSecurityScanCountResponse) SetBody(v *DescribeImageSecurityScanCountResponseBody) *DescribeImageSecurityScanCountResponse {
	s.Body = v
	return s
}

func (s *DescribeImageSecurityScanCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
