// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainViewSourceProvincesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainViewSourceProvincesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainViewSourceProvincesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainViewSourceProvincesResponseBody) *DescribeDomainViewSourceProvincesResponse
	GetBody() *DescribeDomainViewSourceProvincesResponseBody
}

type DescribeDomainViewSourceProvincesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainViewSourceProvincesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainViewSourceProvincesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainViewSourceProvincesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewSourceProvincesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainViewSourceProvincesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainViewSourceProvincesResponse) GetBody() *DescribeDomainViewSourceProvincesResponseBody {
	return s.Body
}

func (s *DescribeDomainViewSourceProvincesResponse) SetHeaders(v map[string]*string) *DescribeDomainViewSourceProvincesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainViewSourceProvincesResponse) SetStatusCode(v int32) *DescribeDomainViewSourceProvincesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainViewSourceProvincesResponse) SetBody(v *DescribeDomainViewSourceProvincesResponseBody) *DescribeDomainViewSourceProvincesResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainViewSourceProvincesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
