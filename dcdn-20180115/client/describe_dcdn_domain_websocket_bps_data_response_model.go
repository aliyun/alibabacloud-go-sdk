// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainWebsocketBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainWebsocketBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainWebsocketBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainWebsocketBpsDataResponseBody) *DescribeDcdnDomainWebsocketBpsDataResponse
	GetBody() *DescribeDcdnDomainWebsocketBpsDataResponseBody
}

type DescribeDcdnDomainWebsocketBpsDataResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainWebsocketBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainWebsocketBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponse) GetBody() *DescribeDcdnDomainWebsocketBpsDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainWebsocketBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainWebsocketBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponse) SetBody(v *DescribeDcdnDomainWebsocketBpsDataResponseBody) *DescribeDcdnDomainWebsocketBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponse) Validate() error {
	return dara.Validate(s)
}
