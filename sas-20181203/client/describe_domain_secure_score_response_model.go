// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecureScoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainSecureScoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainSecureScoreResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainSecureScoreResponseBody) *DescribeDomainSecureScoreResponse
	GetBody() *DescribeDomainSecureScoreResponseBody
}

type DescribeDomainSecureScoreResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainSecureScoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainSecureScoreResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureScoreResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureScoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainSecureScoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainSecureScoreResponse) GetBody() *DescribeDomainSecureScoreResponseBody {
	return s.Body
}

func (s *DescribeDomainSecureScoreResponse) SetHeaders(v map[string]*string) *DescribeDomainSecureScoreResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSecureScoreResponse) SetStatusCode(v int32) *DescribeDomainSecureScoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainSecureScoreResponse) SetBody(v *DescribeDomainSecureScoreResponseBody) *DescribeDomainSecureScoreResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainSecureScoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
