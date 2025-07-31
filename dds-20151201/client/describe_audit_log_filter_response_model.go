// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditLogFilterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuditLogFilterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuditLogFilterResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuditLogFilterResponseBody) *DescribeAuditLogFilterResponse
	GetBody() *DescribeAuditLogFilterResponseBody
}

type DescribeAuditLogFilterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuditLogFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuditLogFilterResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditLogFilterResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogFilterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuditLogFilterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuditLogFilterResponse) GetBody() *DescribeAuditLogFilterResponseBody {
	return s.Body
}

func (s *DescribeAuditLogFilterResponse) SetHeaders(v map[string]*string) *DescribeAuditLogFilterResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditLogFilterResponse) SetStatusCode(v int32) *DescribeAuditLogFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditLogFilterResponse) SetBody(v *DescribeAuditLogFilterResponseBody) *DescribeAuditLogFilterResponse {
	s.Body = v
	return s
}

func (s *DescribeAuditLogFilterResponse) Validate() error {
	return dara.Validate(s)
}
