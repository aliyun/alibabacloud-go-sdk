// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainWebsocketTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainWebsocketTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainWebsocketTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainWebsocketTrafficDataResponseBody) *DescribeDcdnDomainWebsocketTrafficDataResponse
	GetBody() *DescribeDcdnDomainWebsocketTrafficDataResponseBody
}

type DescribeDcdnDomainWebsocketTrafficDataResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainWebsocketTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainWebsocketTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponse) GetBody() *DescribeDcdnDomainWebsocketTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainWebsocketTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainWebsocketTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponse) SetBody(v *DescribeDcdnDomainWebsocketTrafficDataResponseBody) *DescribeDcdnDomainWebsocketTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponse) Validate() error {
	return dara.Validate(s)
}
