// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuditLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuditLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuditLogConfigResponseBody) *DescribeAuditLogConfigResponse
	GetBody() *DescribeAuditLogConfigResponseBody
}

type DescribeAuditLogConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuditLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuditLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditLogConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuditLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuditLogConfigResponse) GetBody() *DescribeAuditLogConfigResponseBody {
	return s.Body
}

func (s *DescribeAuditLogConfigResponse) SetHeaders(v map[string]*string) *DescribeAuditLogConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditLogConfigResponse) SetStatusCode(v int32) *DescribeAuditLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditLogConfigResponse) SetBody(v *DescribeAuditLogConfigResponseBody) *DescribeAuditLogConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeAuditLogConfigResponse) Validate() error {
	return dara.Validate(s)
}
