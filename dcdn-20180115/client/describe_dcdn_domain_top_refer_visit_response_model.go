// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainTopReferVisitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainTopReferVisitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainTopReferVisitResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainTopReferVisitResponseBody) *DescribeDcdnDomainTopReferVisitResponse
	GetBody() *DescribeDcdnDomainTopReferVisitResponseBody
}

type DescribeDcdnDomainTopReferVisitResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainTopReferVisitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainTopReferVisitResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopReferVisitResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopReferVisitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainTopReferVisitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainTopReferVisitResponse) GetBody() *DescribeDcdnDomainTopReferVisitResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainTopReferVisitResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainTopReferVisitResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponse) SetStatusCode(v int32) *DescribeDcdnDomainTopReferVisitResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponse) SetBody(v *DescribeDcdnDomainTopReferVisitResponseBody) *DescribeDcdnDomainTopReferVisitResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainTopReferVisitResponse) Validate() error {
	return dara.Validate(s)
}
