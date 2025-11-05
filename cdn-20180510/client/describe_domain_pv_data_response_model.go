// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainPvDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainPvDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainPvDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainPvDataResponseBody) *DescribeDomainPvDataResponse
	GetBody() *DescribeDomainPvDataResponseBody
}

type DescribeDomainPvDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainPvDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainPvDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainPvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainPvDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainPvDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainPvDataResponse) GetBody() *DescribeDomainPvDataResponseBody {
	return s.Body
}

func (s *DescribeDomainPvDataResponse) SetHeaders(v map[string]*string) *DescribeDomainPvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainPvDataResponse) SetStatusCode(v int32) *DescribeDomainPvDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainPvDataResponse) SetBody(v *DescribeDomainPvDataResponseBody) *DescribeDomainPvDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainPvDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
