// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainDNSRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainDNSRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainDNSRecordResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainDNSRecordResponseBody) *DescribeDomainDNSRecordResponse
	GetBody() *DescribeDomainDNSRecordResponseBody
}

type DescribeDomainDNSRecordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainDNSRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainDNSRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDNSRecordResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainDNSRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainDNSRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainDNSRecordResponse) GetBody() *DescribeDomainDNSRecordResponseBody {
	return s.Body
}

func (s *DescribeDomainDNSRecordResponse) SetHeaders(v map[string]*string) *DescribeDomainDNSRecordResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainDNSRecordResponse) SetStatusCode(v int32) *DescribeDomainDNSRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainDNSRecordResponse) SetBody(v *DescribeDomainDNSRecordResponseBody) *DescribeDomainDNSRecordResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainDNSRecordResponse) Validate() error {
	return dara.Validate(s)
}
