// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainListResponseBody) *DescribeDomainListResponse
	GetBody() *DescribeDomainListResponseBody
}

type DescribeDomainListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainListResponse) GetBody() *DescribeDomainListResponseBody {
	return s.Body
}

func (s *DescribeDomainListResponse) SetHeaders(v map[string]*string) *DescribeDomainListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainListResponse) SetStatusCode(v int32) *DescribeDomainListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainListResponse) SetBody(v *DescribeDomainListResponseBody) *DescribeDomainListResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
