// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecurityProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainSecurityProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainSecurityProfileResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainSecurityProfileResponseBody) *DescribeDomainSecurityProfileResponse
	GetBody() *DescribeDomainSecurityProfileResponseBody
}

type DescribeDomainSecurityProfileResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainSecurityProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainSecurityProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecurityProfileResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecurityProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainSecurityProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainSecurityProfileResponse) GetBody() *DescribeDomainSecurityProfileResponseBody {
	return s.Body
}

func (s *DescribeDomainSecurityProfileResponse) SetHeaders(v map[string]*string) *DescribeDomainSecurityProfileResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSecurityProfileResponse) SetStatusCode(v int32) *DescribeDomainSecurityProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainSecurityProfileResponse) SetBody(v *DescribeDomainSecurityProfileResponseBody) *DescribeDomainSecurityProfileResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainSecurityProfileResponse) Validate() error {
	return dara.Validate(s)
}
