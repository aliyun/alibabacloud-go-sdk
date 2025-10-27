// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainDetailResponseBody) *DescribeDomainDetailResponse
	GetBody() *DescribeDomainDetailResponseBody
}

type DescribeDomainDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainDetailResponse) GetBody() *DescribeDomainDetailResponseBody {
	return s.Body
}

func (s *DescribeDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainDetailResponse) SetStatusCode(v int32) *DescribeDomainDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainDetailResponse) SetBody(v *DescribeDomainDetailResponseBody) *DescribeDomainDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
