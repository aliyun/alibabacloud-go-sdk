// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubDomainRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSubDomainRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSubDomainRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSubDomainRecordsResponseBody) *DescribeSubDomainRecordsResponse
	GetBody() *DescribeSubDomainRecordsResponseBody
}

type DescribeSubDomainRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSubDomainRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSubDomainRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubDomainRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSubDomainRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSubDomainRecordsResponse) GetBody() *DescribeSubDomainRecordsResponseBody {
	return s.Body
}

func (s *DescribeSubDomainRecordsResponse) SetHeaders(v map[string]*string) *DescribeSubDomainRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubDomainRecordsResponse) SetStatusCode(v int32) *DescribeSubDomainRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubDomainRecordsResponse) SetBody(v *DescribeSubDomainRecordsResponseBody) *DescribeSubDomainRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeSubDomainRecordsResponse) Validate() error {
	return dara.Validate(s)
}
