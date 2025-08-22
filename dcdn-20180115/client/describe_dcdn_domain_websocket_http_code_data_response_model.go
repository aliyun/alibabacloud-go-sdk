// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainWebsocketHttpCodeDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainWebsocketHttpCodeDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainWebsocketHttpCodeDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) *DescribeDcdnDomainWebsocketHttpCodeDataResponse
	GetBody() *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody
}

type DescribeDcdnDomainWebsocketHttpCodeDataResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponse) GetBody() *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainWebsocketHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainWebsocketHttpCodeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponse) SetBody(v *DescribeDcdnDomainWebsocketHttpCodeDataResponseBody) *DescribeDcdnDomainWebsocketHttpCodeDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainWebsocketHttpCodeDataResponse) Validate() error {
	return dara.Validate(s)
}
