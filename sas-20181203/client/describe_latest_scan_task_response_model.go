// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLatestScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLatestScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLatestScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLatestScanTaskResponseBody) *DescribeLatestScanTaskResponse
	GetBody() *DescribeLatestScanTaskResponseBody
}

type DescribeLatestScanTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLatestScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLatestScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLatestScanTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeLatestScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLatestScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLatestScanTaskResponse) GetBody() *DescribeLatestScanTaskResponseBody {
	return s.Body
}

func (s *DescribeLatestScanTaskResponse) SetHeaders(v map[string]*string) *DescribeLatestScanTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeLatestScanTaskResponse) SetStatusCode(v int32) *DescribeLatestScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLatestScanTaskResponse) SetBody(v *DescribeLatestScanTaskResponseBody) *DescribeLatestScanTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeLatestScanTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
