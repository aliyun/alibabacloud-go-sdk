// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainPathDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainPathDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainPathDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainPathDataResponseBody) *DescribeDomainPathDataResponse
	GetBody() *DescribeDomainPathDataResponseBody
}

type DescribeDomainPathDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainPathDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainPathDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainPathDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainPathDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainPathDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainPathDataResponse) GetBody() *DescribeDomainPathDataResponseBody {
	return s.Body
}

func (s *DescribeDomainPathDataResponse) SetHeaders(v map[string]*string) *DescribeDomainPathDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainPathDataResponse) SetStatusCode(v int32) *DescribeDomainPathDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainPathDataResponse) SetBody(v *DescribeDomainPathDataResponseBody) *DescribeDomainPathDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainPathDataResponse) Validate() error {
	return dara.Validate(s)
}
