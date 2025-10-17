// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAuditLogCollectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterAuditLogCollectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterAuditLogCollectorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterAuditLogCollectorResponseBody) *DescribeDBClusterAuditLogCollectorResponse
	GetBody() *DescribeDBClusterAuditLogCollectorResponseBody
}

type DescribeDBClusterAuditLogCollectorResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterAuditLogCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterAuditLogCollectorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAuditLogCollectorResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAuditLogCollectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterAuditLogCollectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterAuditLogCollectorResponse) GetBody() *DescribeDBClusterAuditLogCollectorResponseBody {
	return s.Body
}

func (s *DescribeDBClusterAuditLogCollectorResponse) SetHeaders(v map[string]*string) *DescribeDBClusterAuditLogCollectorResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorResponse) SetStatusCode(v int32) *DescribeDBClusterAuditLogCollectorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorResponse) SetBody(v *DescribeDBClusterAuditLogCollectorResponseBody) *DescribeDBClusterAuditLogCollectorResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
