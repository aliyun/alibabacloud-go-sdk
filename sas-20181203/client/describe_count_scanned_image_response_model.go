// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCountScannedImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCountScannedImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCountScannedImageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCountScannedImageResponseBody) *DescribeCountScannedImageResponse
	GetBody() *DescribeCountScannedImageResponseBody
}

type DescribeCountScannedImageResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCountScannedImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCountScannedImageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCountScannedImageResponse) GoString() string {
	return s.String()
}

func (s *DescribeCountScannedImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCountScannedImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCountScannedImageResponse) GetBody() *DescribeCountScannedImageResponseBody {
	return s.Body
}

func (s *DescribeCountScannedImageResponse) SetHeaders(v map[string]*string) *DescribeCountScannedImageResponse {
	s.Headers = v
	return s
}

func (s *DescribeCountScannedImageResponse) SetStatusCode(v int32) *DescribeCountScannedImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCountScannedImageResponse) SetBody(v *DescribeCountScannedImageResponseBody) *DescribeCountScannedImageResponse {
	s.Body = v
	return s
}

func (s *DescribeCountScannedImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
