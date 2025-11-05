// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainsBySourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainsBySourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainsBySourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainsBySourceResponseBody) *DescribeDomainsBySourceResponse
	GetBody() *DescribeDomainsBySourceResponseBody
}

type DescribeDomainsBySourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainsBySourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainsBySourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsBySourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainsBySourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainsBySourceResponse) GetBody() *DescribeDomainsBySourceResponseBody {
	return s.Body
}

func (s *DescribeDomainsBySourceResponse) SetHeaders(v map[string]*string) *DescribeDomainsBySourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainsBySourceResponse) SetStatusCode(v int32) *DescribeDomainsBySourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainsBySourceResponse) SetBody(v *DescribeDomainsBySourceResponseBody) *DescribeDomainsBySourceResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainsBySourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
