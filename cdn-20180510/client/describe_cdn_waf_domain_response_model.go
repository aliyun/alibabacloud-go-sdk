// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnWafDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnWafDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnWafDomainResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnWafDomainResponseBody) *DescribeCdnWafDomainResponse
	GetBody() *DescribeCdnWafDomainResponseBody
}

type DescribeCdnWafDomainResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnWafDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnWafDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnWafDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnWafDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnWafDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnWafDomainResponse) GetBody() *DescribeCdnWafDomainResponseBody {
	return s.Body
}

func (s *DescribeCdnWafDomainResponse) SetHeaders(v map[string]*string) *DescribeCdnWafDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnWafDomainResponse) SetStatusCode(v int32) *DescribeCdnWafDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnWafDomainResponse) SetBody(v *DescribeCdnWafDomainResponseBody) *DescribeCdnWafDomainResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnWafDomainResponse) Validate() error {
	return dara.Validate(s)
}
