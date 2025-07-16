// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopClientIpVisitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainTopClientIpVisitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainTopClientIpVisitResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainTopClientIpVisitResponseBody) *DescribeDomainTopClientIpVisitResponse
	GetBody() *DescribeDomainTopClientIpVisitResponseBody
}

type DescribeDomainTopClientIpVisitResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainTopClientIpVisitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainTopClientIpVisitResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopClientIpVisitResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopClientIpVisitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainTopClientIpVisitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainTopClientIpVisitResponse) GetBody() *DescribeDomainTopClientIpVisitResponseBody {
	return s.Body
}

func (s *DescribeDomainTopClientIpVisitResponse) SetHeaders(v map[string]*string) *DescribeDomainTopClientIpVisitResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainTopClientIpVisitResponse) SetStatusCode(v int32) *DescribeDomainTopClientIpVisitResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitResponse) SetBody(v *DescribeDomainTopClientIpVisitResponseBody) *DescribeDomainTopClientIpVisitResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainTopClientIpVisitResponse) Validate() error {
	return dara.Validate(s)
}
