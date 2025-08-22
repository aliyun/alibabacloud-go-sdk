// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnTopDomainsByFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnTopDomainsByFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnTopDomainsByFlowResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnTopDomainsByFlowResponseBody) *DescribeDcdnTopDomainsByFlowResponse
	GetBody() *DescribeDcdnTopDomainsByFlowResponseBody
}

type DescribeDcdnTopDomainsByFlowResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnTopDomainsByFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnTopDomainsByFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnTopDomainsByFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTopDomainsByFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnTopDomainsByFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnTopDomainsByFlowResponse) GetBody() *DescribeDcdnTopDomainsByFlowResponseBody {
	return s.Body
}

func (s *DescribeDcdnTopDomainsByFlowResponse) SetHeaders(v map[string]*string) *DescribeDcdnTopDomainsByFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponse) SetStatusCode(v int32) *DescribeDcdnTopDomainsByFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponse) SetBody(v *DescribeDcdnTopDomainsByFlowResponseBody) *DescribeDcdnTopDomainsByFlowResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponse) Validate() error {
	return dara.Validate(s)
}
