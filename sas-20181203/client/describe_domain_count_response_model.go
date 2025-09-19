// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainCountResponseBody) *DescribeDomainCountResponse
	GetBody() *DescribeDomainCountResponseBody
}

type DescribeDomainCountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainCountResponse) GetBody() *DescribeDomainCountResponseBody {
	return s.Body
}

func (s *DescribeDomainCountResponse) SetHeaders(v map[string]*string) *DescribeDomainCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainCountResponse) SetStatusCode(v int32) *DescribeDomainCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainCountResponse) SetBody(v *DescribeDomainCountResponseBody) *DescribeDomainCountResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainCountResponse) Validate() error {
	return dara.Validate(s)
}
