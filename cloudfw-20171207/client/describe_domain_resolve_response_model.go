// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainResolveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainResolveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainResolveResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainResolveResponseBody) *DescribeDomainResolveResponse
	GetBody() *DescribeDomainResolveResponseBody
}

type DescribeDomainResolveResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainResolveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainResolveResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainResolveResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainResolveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainResolveResponse) GetBody() *DescribeDomainResolveResponseBody {
	return s.Body
}

func (s *DescribeDomainResolveResponse) SetHeaders(v map[string]*string) *DescribeDomainResolveResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainResolveResponse) SetStatusCode(v int32) *DescribeDomainResolveResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainResolveResponse) SetBody(v *DescribeDomainResolveResponseBody) *DescribeDomainResolveResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainResolveResponse) Validate() error {
	return dara.Validate(s)
}
