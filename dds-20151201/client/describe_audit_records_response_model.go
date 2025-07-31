// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuditRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuditRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuditRecordsResponseBody) *DescribeAuditRecordsResponse
	GetBody() *DescribeAuditRecordsResponseBody
}

type DescribeAuditRecordsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuditRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuditRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuditRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuditRecordsResponse) GetBody() *DescribeAuditRecordsResponseBody {
	return s.Body
}

func (s *DescribeAuditRecordsResponse) SetHeaders(v map[string]*string) *DescribeAuditRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditRecordsResponse) SetStatusCode(v int32) *DescribeAuditRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditRecordsResponse) SetBody(v *DescribeAuditRecordsResponseBody) *DescribeAuditRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeAuditRecordsResponse) Validate() error {
	return dara.Validate(s)
}
