// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainLogResponseBody) *DescribeDcdnDomainLogResponse
	GetBody() *DescribeDcdnDomainLogResponseBody
}

type DescribeDcdnDomainLogResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainLogResponse) GetBody() *DescribeDcdnDomainLogResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainLogResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainLogResponse) SetStatusCode(v int32) *DescribeDcdnDomainLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainLogResponse) SetBody(v *DescribeDcdnDomainLogResponseBody) *DescribeDcdnDomainLogResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainLogResponse) Validate() error {
	return dara.Validate(s)
}
