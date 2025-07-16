// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSrcTopUrlVisitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainSrcTopUrlVisitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainSrcTopUrlVisitResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainSrcTopUrlVisitResponseBody) *DescribeDomainSrcTopUrlVisitResponse
	GetBody() *DescribeDomainSrcTopUrlVisitResponseBody
}

type DescribeDomainSrcTopUrlVisitResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainSrcTopUrlVisitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainSrcTopUrlVisitResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainSrcTopUrlVisitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainSrcTopUrlVisitResponse) GetBody() *DescribeDomainSrcTopUrlVisitResponseBody {
	return s.Body
}

func (s *DescribeDomainSrcTopUrlVisitResponse) SetHeaders(v map[string]*string) *DescribeDomainSrcTopUrlVisitResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponse) SetStatusCode(v int32) *DescribeDomainSrcTopUrlVisitResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponse) SetBody(v *DescribeDomainSrcTopUrlVisitResponseBody) *DescribeDomainSrcTopUrlVisitResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponse) Validate() error {
	return dara.Validate(s)
}
