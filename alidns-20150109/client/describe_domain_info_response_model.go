// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainInfoResponseBody) *DescribeDomainInfoResponse
	GetBody() *DescribeDomainInfoResponseBody
}

type DescribeDomainInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainInfoResponse) GetBody() *DescribeDomainInfoResponseBody {
	return s.Body
}

func (s *DescribeDomainInfoResponse) SetHeaders(v map[string]*string) *DescribeDomainInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainInfoResponse) SetStatusCode(v int32) *DescribeDomainInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainInfoResponse) SetBody(v *DescribeDomainInfoResponseBody) *DescribeDomainInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainInfoResponse) Validate() error {
	return dara.Validate(s)
}
