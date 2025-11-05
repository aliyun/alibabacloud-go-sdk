// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopReferVisitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainTopReferVisitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainTopReferVisitResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainTopReferVisitResponseBody) *DescribeDomainTopReferVisitResponse
	GetBody() *DescribeDomainTopReferVisitResponseBody
}

type DescribeDomainTopReferVisitResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainTopReferVisitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainTopReferVisitResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopReferVisitResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopReferVisitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainTopReferVisitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainTopReferVisitResponse) GetBody() *DescribeDomainTopReferVisitResponseBody {
	return s.Body
}

func (s *DescribeDomainTopReferVisitResponse) SetHeaders(v map[string]*string) *DescribeDomainTopReferVisitResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainTopReferVisitResponse) SetStatusCode(v int32) *DescribeDomainTopReferVisitResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainTopReferVisitResponse) SetBody(v *DescribeDomainTopReferVisitResponseBody) *DescribeDomainTopReferVisitResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainTopReferVisitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
