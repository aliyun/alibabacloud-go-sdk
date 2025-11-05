// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainUvDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainUvDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainUvDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainUvDataResponseBody) *DescribeDomainUvDataResponse
	GetBody() *DescribeDomainUvDataResponseBody
}

type DescribeDomainUvDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainUvDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainUvDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainUvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainUvDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainUvDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainUvDataResponse) GetBody() *DescribeDomainUvDataResponseBody {
	return s.Body
}

func (s *DescribeDomainUvDataResponse) SetHeaders(v map[string]*string) *DescribeDomainUvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainUvDataResponse) SetStatusCode(v int32) *DescribeDomainUvDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainUvDataResponse) SetBody(v *DescribeDomainUvDataResponseBody) *DescribeDomainUvDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainUvDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
