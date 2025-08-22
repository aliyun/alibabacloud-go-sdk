// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainIpaBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainIpaBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainIpaBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainIpaBpsDataResponseBody) *DescribeDcdnDomainIpaBpsDataResponse
	GetBody() *DescribeDcdnDomainIpaBpsDataResponseBody
}

type DescribeDcdnDomainIpaBpsDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainIpaBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainIpaBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIpaBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainIpaBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainIpaBpsDataResponse) GetBody() *DescribeDcdnDomainIpaBpsDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainIpaBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainIpaBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainIpaBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponse) SetBody(v *DescribeDcdnDomainIpaBpsDataResponseBody) *DescribeDcdnDomainIpaBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponse) Validate() error {
	return dara.Validate(s)
}
