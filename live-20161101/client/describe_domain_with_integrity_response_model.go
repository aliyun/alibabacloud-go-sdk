// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainWithIntegrityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainWithIntegrityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainWithIntegrityResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainWithIntegrityResponseBody) *DescribeDomainWithIntegrityResponse
	GetBody() *DescribeDomainWithIntegrityResponseBody
}

type DescribeDomainWithIntegrityResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainWithIntegrityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainWithIntegrityResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainWithIntegrityResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainWithIntegrityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainWithIntegrityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainWithIntegrityResponse) GetBody() *DescribeDomainWithIntegrityResponseBody {
	return s.Body
}

func (s *DescribeDomainWithIntegrityResponse) SetHeaders(v map[string]*string) *DescribeDomainWithIntegrityResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainWithIntegrityResponse) SetStatusCode(v int32) *DescribeDomainWithIntegrityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainWithIntegrityResponse) SetBody(v *DescribeDomainWithIntegrityResponseBody) *DescribeDomainWithIntegrityResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainWithIntegrityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
