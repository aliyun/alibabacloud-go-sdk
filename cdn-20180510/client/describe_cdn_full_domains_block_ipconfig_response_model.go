// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnFullDomainsBlockIPConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnFullDomainsBlockIPConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnFullDomainsBlockIPConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnFullDomainsBlockIPConfigResponseBody) *DescribeCdnFullDomainsBlockIPConfigResponse
	GetBody() *DescribeCdnFullDomainsBlockIPConfigResponseBody
}

type DescribeCdnFullDomainsBlockIPConfigResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnFullDomainsBlockIPConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnFullDomainsBlockIPConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnFullDomainsBlockIPConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnFullDomainsBlockIPConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnFullDomainsBlockIPConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnFullDomainsBlockIPConfigResponse) GetBody() *DescribeCdnFullDomainsBlockIPConfigResponseBody {
	return s.Body
}

func (s *DescribeCdnFullDomainsBlockIPConfigResponse) SetHeaders(v map[string]*string) *DescribeCdnFullDomainsBlockIPConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPConfigResponse) SetStatusCode(v int32) *DescribeCdnFullDomainsBlockIPConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPConfigResponse) SetBody(v *DescribeCdnFullDomainsBlockIPConfigResponseBody) *DescribeCdnFullDomainsBlockIPConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPConfigResponse) Validate() error {
	return dara.Validate(s)
}
