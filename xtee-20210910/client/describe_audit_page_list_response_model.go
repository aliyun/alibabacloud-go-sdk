// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditPageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuditPageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuditPageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuditPageListResponseBody) *DescribeAuditPageListResponse
	GetBody() *DescribeAuditPageListResponseBody
}

type DescribeAuditPageListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuditPageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuditPageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditPageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditPageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuditPageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuditPageListResponse) GetBody() *DescribeAuditPageListResponseBody {
	return s.Body
}

func (s *DescribeAuditPageListResponse) SetHeaders(v map[string]*string) *DescribeAuditPageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditPageListResponse) SetStatusCode(v int32) *DescribeAuditPageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditPageListResponse) SetBody(v *DescribeAuditPageListResponseBody) *DescribeAuditPageListResponse {
	s.Body = v
	return s
}

func (s *DescribeAuditPageListResponse) Validate() error {
	return dara.Validate(s)
}
