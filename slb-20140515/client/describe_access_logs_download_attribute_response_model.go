// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessLogsDownloadAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccessLogsDownloadAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccessLogsDownloadAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccessLogsDownloadAttributeResponseBody) *DescribeAccessLogsDownloadAttributeResponse
	GetBody() *DescribeAccessLogsDownloadAttributeResponseBody
}

type DescribeAccessLogsDownloadAttributeResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccessLogsDownloadAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccessLogsDownloadAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessLogsDownloadAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessLogsDownloadAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccessLogsDownloadAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccessLogsDownloadAttributeResponse) GetBody() *DescribeAccessLogsDownloadAttributeResponseBody {
	return s.Body
}

func (s *DescribeAccessLogsDownloadAttributeResponse) SetHeaders(v map[string]*string) *DescribeAccessLogsDownloadAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponse) SetStatusCode(v int32) *DescribeAccessLogsDownloadAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponse) SetBody(v *DescribeAccessLogsDownloadAttributeResponseBody) *DescribeAccessLogsDownloadAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeAccessLogsDownloadAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
