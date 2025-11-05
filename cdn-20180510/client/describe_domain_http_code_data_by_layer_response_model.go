// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainHttpCodeDataByLayerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainHttpCodeDataByLayerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainHttpCodeDataByLayerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainHttpCodeDataByLayerResponseBody) *DescribeDomainHttpCodeDataByLayerResponse
	GetBody() *DescribeDomainHttpCodeDataByLayerResponseBody
}

type DescribeDomainHttpCodeDataByLayerResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainHttpCodeDataByLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainHttpCodeDataByLayerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainHttpCodeDataByLayerResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataByLayerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainHttpCodeDataByLayerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainHttpCodeDataByLayerResponse) GetBody() *DescribeDomainHttpCodeDataByLayerResponseBody {
	return s.Body
}

func (s *DescribeDomainHttpCodeDataByLayerResponse) SetHeaders(v map[string]*string) *DescribeDomainHttpCodeDataByLayerResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerResponse) SetStatusCode(v int32) *DescribeDomainHttpCodeDataByLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerResponse) SetBody(v *DescribeDomainHttpCodeDataByLayerResponseBody) *DescribeDomainHttpCodeDataByLayerResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
