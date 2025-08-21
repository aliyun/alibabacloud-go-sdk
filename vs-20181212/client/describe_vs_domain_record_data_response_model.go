// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainRecordDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsDomainRecordDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsDomainRecordDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsDomainRecordDataResponseBody) *DescribeVsDomainRecordDataResponse
	GetBody() *DescribeVsDomainRecordDataResponseBody
}

type DescribeVsDomainRecordDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsDomainRecordDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsDomainRecordDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainRecordDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRecordDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsDomainRecordDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsDomainRecordDataResponse) GetBody() *DescribeVsDomainRecordDataResponseBody {
	return s.Body
}

func (s *DescribeVsDomainRecordDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainRecordDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainRecordDataResponse) SetStatusCode(v int32) *DescribeVsDomainRecordDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsDomainRecordDataResponse) SetBody(v *DescribeVsDomainRecordDataResponseBody) *DescribeVsDomainRecordDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVsDomainRecordDataResponse) Validate() error {
	return dara.Validate(s)
}
