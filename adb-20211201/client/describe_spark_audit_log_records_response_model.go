// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkAuditLogRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSparkAuditLogRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSparkAuditLogRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSparkAuditLogRecordsResponseBody) *DescribeSparkAuditLogRecordsResponse
	GetBody() *DescribeSparkAuditLogRecordsResponseBody
}

type DescribeSparkAuditLogRecordsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSparkAuditLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSparkAuditLogRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkAuditLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSparkAuditLogRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSparkAuditLogRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSparkAuditLogRecordsResponse) GetBody() *DescribeSparkAuditLogRecordsResponseBody {
	return s.Body
}

func (s *DescribeSparkAuditLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeSparkAuditLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponse) SetStatusCode(v int32) *DescribeSparkAuditLogRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponse) SetBody(v *DescribeSparkAuditLogRecordsResponseBody) *DescribeSparkAuditLogRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeSparkAuditLogRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
