// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainNsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainNsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainNsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainNsResponseBody) *DescribeDomainNsResponse
	GetBody() *DescribeDomainNsResponseBody
}

type DescribeDomainNsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainNsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainNsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainNsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainNsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainNsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainNsResponse) GetBody() *DescribeDomainNsResponseBody {
	return s.Body
}

func (s *DescribeDomainNsResponse) SetHeaders(v map[string]*string) *DescribeDomainNsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainNsResponse) SetStatusCode(v int32) *DescribeDomainNsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainNsResponse) SetBody(v *DescribeDomainNsResponseBody) *DescribeDomainNsResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainNsResponse) Validate() error {
	return dara.Validate(s)
}
