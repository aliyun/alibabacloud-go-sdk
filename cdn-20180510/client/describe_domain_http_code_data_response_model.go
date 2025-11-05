// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainHttpCodeDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainHttpCodeDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainHttpCodeDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainHttpCodeDataResponseBody) *DescribeDomainHttpCodeDataResponse
	GetBody() *DescribeDomainHttpCodeDataResponseBody
}

type DescribeDomainHttpCodeDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainHttpCodeDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainHttpCodeDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainHttpCodeDataResponse) GetBody() *DescribeDomainHttpCodeDataResponseBody {
	return s.Body
}

func (s *DescribeDomainHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeDomainHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainHttpCodeDataResponse) SetStatusCode(v int32) *DescribeDomainHttpCodeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponse) SetBody(v *DescribeDomainHttpCodeDataResponseBody) *DescribeDomainHttpCodeDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainHttpCodeDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
