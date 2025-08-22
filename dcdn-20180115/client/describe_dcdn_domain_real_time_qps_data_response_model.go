// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeQpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeQpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainRealTimeQpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainRealTimeQpsDataResponseBody) *DescribeDcdnDomainRealTimeQpsDataResponse
	GetBody() *DescribeDcdnDomainRealTimeQpsDataResponseBody
}

type DescribeDcdnDomainRealTimeQpsDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainRealTimeQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainRealTimeQpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponse) GetBody() *DescribeDcdnDomainRealTimeQpsDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainRealTimeQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainRealTimeQpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponse) SetBody(v *DescribeDcdnDomainRealTimeQpsDataResponseBody) *DescribeDcdnDomainRealTimeQpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponse) Validate() error {
	return dara.Validate(s)
}
