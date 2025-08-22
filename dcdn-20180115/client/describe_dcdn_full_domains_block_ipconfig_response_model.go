// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnFullDomainsBlockIPConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnFullDomainsBlockIPConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnFullDomainsBlockIPConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnFullDomainsBlockIPConfigResponseBody) *DescribeDcdnFullDomainsBlockIPConfigResponse
	GetBody() *DescribeDcdnFullDomainsBlockIPConfigResponseBody
}

type DescribeDcdnFullDomainsBlockIPConfigResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnFullDomainsBlockIPConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnFullDomainsBlockIPConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnFullDomainsBlockIPConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnFullDomainsBlockIPConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnFullDomainsBlockIPConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnFullDomainsBlockIPConfigResponse) GetBody() *DescribeDcdnFullDomainsBlockIPConfigResponseBody {
	return s.Body
}

func (s *DescribeDcdnFullDomainsBlockIPConfigResponse) SetHeaders(v map[string]*string) *DescribeDcdnFullDomainsBlockIPConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPConfigResponse) SetStatusCode(v int32) *DescribeDcdnFullDomainsBlockIPConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPConfigResponse) SetBody(v *DescribeDcdnFullDomainsBlockIPConfigResponseBody) *DescribeDcdnFullDomainsBlockIPConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPConfigResponse) Validate() error {
	return dara.Validate(s)
}
