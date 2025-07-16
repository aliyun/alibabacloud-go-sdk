// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnDomainConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnDomainConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnDomainConfigsResponseBody) *DescribeCdnDomainConfigsResponse
	GetBody() *DescribeCdnDomainConfigsResponseBody
}

type DescribeCdnDomainConfigsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnDomainConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnDomainConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnDomainConfigsResponse) GetBody() *DescribeCdnDomainConfigsResponseBody {
	return s.Body
}

func (s *DescribeCdnDomainConfigsResponse) SetHeaders(v map[string]*string) *DescribeCdnDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDomainConfigsResponse) SetStatusCode(v int32) *DescribeCdnDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponse) SetBody(v *DescribeCdnDomainConfigsResponseBody) *DescribeCdnDomainConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnDomainConfigsResponse) Validate() error {
	return dara.Validate(s)
}
