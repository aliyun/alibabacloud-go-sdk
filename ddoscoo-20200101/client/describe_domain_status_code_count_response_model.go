// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainStatusCodeCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainStatusCodeCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainStatusCodeCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainStatusCodeCountResponseBody) *DescribeDomainStatusCodeCountResponse
	GetBody() *DescribeDomainStatusCodeCountResponseBody
}

type DescribeDomainStatusCodeCountResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainStatusCodeCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainStatusCodeCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatusCodeCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainStatusCodeCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainStatusCodeCountResponse) GetBody() *DescribeDomainStatusCodeCountResponseBody {
	return s.Body
}

func (s *DescribeDomainStatusCodeCountResponse) SetHeaders(v map[string]*string) *DescribeDomainStatusCodeCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainStatusCodeCountResponse) SetStatusCode(v int32) *DescribeDomainStatusCodeCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponse) SetBody(v *DescribeDomainStatusCodeCountResponseBody) *DescribeDomainStatusCodeCountResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainStatusCodeCountResponse) Validate() error {
	return dara.Validate(s)
}
