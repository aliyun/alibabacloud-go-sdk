// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageLatestScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageLatestScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageLatestScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageLatestScanTaskResponseBody) *DescribeImageLatestScanTaskResponse
	GetBody() *DescribeImageLatestScanTaskResponseBody
}

type DescribeImageLatestScanTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageLatestScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageLatestScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageLatestScanTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageLatestScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageLatestScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageLatestScanTaskResponse) GetBody() *DescribeImageLatestScanTaskResponseBody {
	return s.Body
}

func (s *DescribeImageLatestScanTaskResponse) SetHeaders(v map[string]*string) *DescribeImageLatestScanTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageLatestScanTaskResponse) SetStatusCode(v int32) *DescribeImageLatestScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageLatestScanTaskResponse) SetBody(v *DescribeImageLatestScanTaskResponseBody) *DescribeImageLatestScanTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeImageLatestScanTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
