// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuditLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuditLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuditLogsResponseBody) *DescribeAuditLogsResponse
	GetBody() *DescribeAuditLogsResponseBody
}

type DescribeAuditLogsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuditLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuditLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuditLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuditLogsResponse) GetBody() *DescribeAuditLogsResponseBody {
	return s.Body
}

func (s *DescribeAuditLogsResponse) SetHeaders(v map[string]*string) *DescribeAuditLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditLogsResponse) SetStatusCode(v int32) *DescribeAuditLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditLogsResponse) SetBody(v *DescribeAuditLogsResponseBody) *DescribeAuditLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeAuditLogsResponse) Validate() error {
	return dara.Validate(s)
}
