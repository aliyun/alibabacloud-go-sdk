// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadSQLLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDownloadSQLLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDownloadSQLLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDownloadSQLLogsResponseBody) *DescribeDownloadSQLLogsResponse
	GetBody() *DescribeDownloadSQLLogsResponseBody
}

type DescribeDownloadSQLLogsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDownloadSQLLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDownloadSQLLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadSQLLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadSQLLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDownloadSQLLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDownloadSQLLogsResponse) GetBody() *DescribeDownloadSQLLogsResponseBody {
	return s.Body
}

func (s *DescribeDownloadSQLLogsResponse) SetHeaders(v map[string]*string) *DescribeDownloadSQLLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadSQLLogsResponse) SetStatusCode(v int32) *DescribeDownloadSQLLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDownloadSQLLogsResponse) SetBody(v *DescribeDownloadSQLLogsResponseBody) *DescribeDownloadSQLLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeDownloadSQLLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
