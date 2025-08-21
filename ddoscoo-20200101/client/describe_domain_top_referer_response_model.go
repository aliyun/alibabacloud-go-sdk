// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopRefererResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainTopRefererResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainTopRefererResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainTopRefererResponseBody) *DescribeDomainTopRefererResponse
	GetBody() *DescribeDomainTopRefererResponseBody
}

type DescribeDomainTopRefererResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainTopRefererResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainTopRefererResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopRefererResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopRefererResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainTopRefererResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainTopRefererResponse) GetBody() *DescribeDomainTopRefererResponseBody {
	return s.Body
}

func (s *DescribeDomainTopRefererResponse) SetHeaders(v map[string]*string) *DescribeDomainTopRefererResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainTopRefererResponse) SetStatusCode(v int32) *DescribeDomainTopRefererResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainTopRefererResponse) SetBody(v *DescribeDomainTopRefererResponseBody) *DescribeDomainTopRefererResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainTopRefererResponse) Validate() error {
	return dara.Validate(s)
}
