// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainResponseBody) *DescribeDomainResponse
	GetBody() *DescribeDomainResponseBody
}

type DescribeDomainResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainResponse) GetBody() *DescribeDomainResponseBody {
	return s.Body
}

func (s *DescribeDomainResponse) SetHeaders(v map[string]*string) *DescribeDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainResponse) SetStatusCode(v int32) *DescribeDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainResponse) SetBody(v *DescribeDomainResponseBody) *DescribeDomainResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainResponse) Validate() error {
	return dara.Validate(s)
}
