// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainQpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainQpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainQpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainQpsDataResponseBody) *DescribeDcdnDomainQpsDataResponse
	GetBody() *DescribeDcdnDomainQpsDataResponseBody
}

type DescribeDcdnDomainQpsDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainQpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainQpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainQpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainQpsDataResponse) GetBody() *DescribeDcdnDomainQpsDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainQpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainQpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponse) SetBody(v *DescribeDcdnDomainQpsDataResponseBody) *DescribeDcdnDomainQpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponse) Validate() error {
	return dara.Validate(s)
}
