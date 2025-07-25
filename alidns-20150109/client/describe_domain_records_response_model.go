// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainRecordsResponseBody) *DescribeDomainRecordsResponse
	GetBody() *DescribeDomainRecordsResponseBody
}

type DescribeDomainRecordsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainRecordsResponse) GetBody() *DescribeDomainRecordsResponseBody {
	return s.Body
}

func (s *DescribeDomainRecordsResponse) SetHeaders(v map[string]*string) *DescribeDomainRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRecordsResponse) SetStatusCode(v int32) *DescribeDomainRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainRecordsResponse) SetBody(v *DescribeDomainRecordsResponseBody) *DescribeDomainRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainRecordsResponse) Validate() error {
	return dara.Validate(s)
}
