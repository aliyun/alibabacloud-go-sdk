// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppDomainDnsRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppDomainDnsRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppDomainDnsRecordResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppDomainDnsRecordResponseBody) *DescribeAppDomainDnsRecordResponse
	GetBody() *DescribeAppDomainDnsRecordResponseBody
}

type DescribeAppDomainDnsRecordResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppDomainDnsRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppDomainDnsRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppDomainDnsRecordResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppDomainDnsRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppDomainDnsRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppDomainDnsRecordResponse) GetBody() *DescribeAppDomainDnsRecordResponseBody {
	return s.Body
}

func (s *DescribeAppDomainDnsRecordResponse) SetHeaders(v map[string]*string) *DescribeAppDomainDnsRecordResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppDomainDnsRecordResponse) SetStatusCode(v int32) *DescribeAppDomainDnsRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppDomainDnsRecordResponse) SetBody(v *DescribeAppDomainDnsRecordResponseBody) *DescribeAppDomainDnsRecordResponse {
	s.Body = v
	return s
}

func (s *DescribeAppDomainDnsRecordResponse) Validate() error {
	return dara.Validate(s)
}
