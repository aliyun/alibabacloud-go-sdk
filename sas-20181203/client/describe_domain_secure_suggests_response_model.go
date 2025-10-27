// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecureSuggestsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainSecureSuggestsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainSecureSuggestsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainSecureSuggestsResponseBody) *DescribeDomainSecureSuggestsResponse
	GetBody() *DescribeDomainSecureSuggestsResponseBody
}

type DescribeDomainSecureSuggestsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainSecureSuggestsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainSecureSuggestsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureSuggestsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureSuggestsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainSecureSuggestsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainSecureSuggestsResponse) GetBody() *DescribeDomainSecureSuggestsResponseBody {
	return s.Body
}

func (s *DescribeDomainSecureSuggestsResponse) SetHeaders(v map[string]*string) *DescribeDomainSecureSuggestsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSecureSuggestsResponse) SetStatusCode(v int32) *DescribeDomainSecureSuggestsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainSecureSuggestsResponse) SetBody(v *DescribeDomainSecureSuggestsResponseBody) *DescribeDomainSecureSuggestsResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainSecureSuggestsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
