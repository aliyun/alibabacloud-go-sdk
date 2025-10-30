// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourceDataDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataSourceDataDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataSourceDataDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataSourceDataDownloadUrlResponseBody) *DescribeDataSourceDataDownloadUrlResponse
	GetBody() *DescribeDataSourceDataDownloadUrlResponseBody
}

type DescribeDataSourceDataDownloadUrlResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataSourceDataDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataSourceDataDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceDataDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceDataDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataSourceDataDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataSourceDataDownloadUrlResponse) GetBody() *DescribeDataSourceDataDownloadUrlResponseBody {
	return s.Body
}

func (s *DescribeDataSourceDataDownloadUrlResponse) SetHeaders(v map[string]*string) *DescribeDataSourceDataDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataSourceDataDownloadUrlResponse) SetStatusCode(v int32) *DescribeDataSourceDataDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataSourceDataDownloadUrlResponse) SetBody(v *DescribeDataSourceDataDownloadUrlResponseBody) *DescribeDataSourceDataDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeDataSourceDataDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
