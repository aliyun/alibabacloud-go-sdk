// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainMax95BpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainMax95BpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainMax95BpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainMax95BpsDataResponseBody) *DescribeDomainMax95BpsDataResponse
	GetBody() *DescribeDomainMax95BpsDataResponseBody
}

type DescribeDomainMax95BpsDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainMax95BpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainMax95BpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainMax95BpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainMax95BpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainMax95BpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainMax95BpsDataResponse) GetBody() *DescribeDomainMax95BpsDataResponseBody {
	return s.Body
}

func (s *DescribeDomainMax95BpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainMax95BpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainMax95BpsDataResponse) SetStatusCode(v int32) *DescribeDomainMax95BpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponse) SetBody(v *DescribeDomainMax95BpsDataResponseBody) *DescribeDomainMax95BpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainMax95BpsDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
