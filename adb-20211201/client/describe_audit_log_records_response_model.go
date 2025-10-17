// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditLogRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuditLogRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuditLogRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuditLogRecordsResponseBody) *DescribeAuditLogRecordsResponse
	GetBody() *DescribeAuditLogRecordsResponseBody
}

type DescribeAuditLogRecordsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuditLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuditLogRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuditLogRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuditLogRecordsResponse) GetBody() *DescribeAuditLogRecordsResponseBody {
	return s.Body
}

func (s *DescribeAuditLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeAuditLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditLogRecordsResponse) SetStatusCode(v int32) *DescribeAuditLogRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditLogRecordsResponse) SetBody(v *DescribeAuditLogRecordsResponseBody) *DescribeAuditLogRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeAuditLogRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
