// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainTopUrlVisitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainTopUrlVisitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainTopUrlVisitResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainTopUrlVisitResponseBody) *DescribeDcdnDomainTopUrlVisitResponse
	GetBody() *DescribeDcdnDomainTopUrlVisitResponseBody
}

type DescribeDcdnDomainTopUrlVisitResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainTopUrlVisitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainTopUrlVisitResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTopUrlVisitResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTopUrlVisitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainTopUrlVisitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainTopUrlVisitResponse) GetBody() *DescribeDcdnDomainTopUrlVisitResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainTopUrlVisitResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainTopUrlVisitResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponse) SetStatusCode(v int32) *DescribeDcdnDomainTopUrlVisitResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponse) SetBody(v *DescribeDcdnDomainTopUrlVisitResponseBody) *DescribeDcdnDomainTopUrlVisitResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainTopUrlVisitResponse) Validate() error {
	return dara.Validate(s)
}
