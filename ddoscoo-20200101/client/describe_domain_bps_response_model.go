// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainBpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainBpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainBpsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainBpsResponseBody) *DescribeDomainBpsResponse
	GetBody() *DescribeDomainBpsResponseBody
}

type DescribeDomainBpsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainBpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainBpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainBpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainBpsResponse) GetBody() *DescribeDomainBpsResponseBody {
	return s.Body
}

func (s *DescribeDomainBpsResponse) SetHeaders(v map[string]*string) *DescribeDomainBpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainBpsResponse) SetStatusCode(v int32) *DescribeDomainBpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainBpsResponse) SetBody(v *DescribeDomainBpsResponseBody) *DescribeDomainBpsResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainBpsResponse) Validate() error {
	return dara.Validate(s)
}
