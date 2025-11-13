// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterScannerListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterScannerListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterScannerListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterScannerListResponseBody) *DescribeClusterScannerListResponse
	GetBody() *DescribeClusterScannerListResponseBody
}

type DescribeClusterScannerListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterScannerListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterScannerListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterScannerListResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterScannerListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterScannerListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterScannerListResponse) GetBody() *DescribeClusterScannerListResponseBody {
	return s.Body
}

func (s *DescribeClusterScannerListResponse) SetHeaders(v map[string]*string) *DescribeClusterScannerListResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterScannerListResponse) SetStatusCode(v int32) *DescribeClusterScannerListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterScannerListResponse) SetBody(v *DescribeClusterScannerListResponseBody) *DescribeClusterScannerListResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterScannerListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
