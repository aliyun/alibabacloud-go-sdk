// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDownloadRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDownloadRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDownloadRecordsResponseBody) *DescribeDownloadRecordsResponse
	GetBody() *DescribeDownloadRecordsResponseBody
}

type DescribeDownloadRecordsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDownloadRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDownloadRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDownloadRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDownloadRecordsResponse) GetBody() *DescribeDownloadRecordsResponseBody {
	return s.Body
}

func (s *DescribeDownloadRecordsResponse) SetHeaders(v map[string]*string) *DescribeDownloadRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadRecordsResponse) SetStatusCode(v int32) *DescribeDownloadRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDownloadRecordsResponse) SetBody(v *DescribeDownloadRecordsResponseBody) *DescribeDownloadRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeDownloadRecordsResponse) Validate() error {
	return dara.Validate(s)
}
