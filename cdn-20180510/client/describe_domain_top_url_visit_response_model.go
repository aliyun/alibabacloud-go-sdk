// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopUrlVisitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainTopUrlVisitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainTopUrlVisitResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainTopUrlVisitResponseBody) *DescribeDomainTopUrlVisitResponse
	GetBody() *DescribeDomainTopUrlVisitResponseBody
}

type DescribeDomainTopUrlVisitResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainTopUrlVisitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainTopUrlVisitResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainTopUrlVisitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainTopUrlVisitResponse) GetBody() *DescribeDomainTopUrlVisitResponseBody {
	return s.Body
}

func (s *DescribeDomainTopUrlVisitResponse) SetHeaders(v map[string]*string) *DescribeDomainTopUrlVisitResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainTopUrlVisitResponse) SetStatusCode(v int32) *DescribeDomainTopUrlVisitResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponse) SetBody(v *DescribeDomainTopUrlVisitResponseBody) *DescribeDomainTopUrlVisitResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainTopUrlVisitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
