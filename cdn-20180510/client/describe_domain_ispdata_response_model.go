// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainISPDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainISPDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainISPDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainISPDataResponseBody) *DescribeDomainISPDataResponse
	GetBody() *DescribeDomainISPDataResponseBody
}

type DescribeDomainISPDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainISPDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainISPDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainISPDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainISPDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainISPDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainISPDataResponse) GetBody() *DescribeDomainISPDataResponseBody {
	return s.Body
}

func (s *DescribeDomainISPDataResponse) SetHeaders(v map[string]*string) *DescribeDomainISPDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainISPDataResponse) SetStatusCode(v int32) *DescribeDomainISPDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainISPDataResponse) SetBody(v *DescribeDomainISPDataResponseBody) *DescribeDomainISPDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainISPDataResponse) Validate() error {
	return dara.Validate(s)
}
