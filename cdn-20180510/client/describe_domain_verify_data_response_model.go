// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainVerifyDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainVerifyDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainVerifyDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainVerifyDataResponseBody) *DescribeDomainVerifyDataResponse
	GetBody() *DescribeDomainVerifyDataResponseBody
}

type DescribeDomainVerifyDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainVerifyDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainVerifyDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainVerifyDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainVerifyDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainVerifyDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainVerifyDataResponse) GetBody() *DescribeDomainVerifyDataResponseBody {
	return s.Body
}

func (s *DescribeDomainVerifyDataResponse) SetHeaders(v map[string]*string) *DescribeDomainVerifyDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainVerifyDataResponse) SetStatusCode(v int32) *DescribeDomainVerifyDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainVerifyDataResponse) SetBody(v *DescribeDomainVerifyDataResponseBody) *DescribeDomainVerifyDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainVerifyDataResponse) Validate() error {
	return dara.Validate(s)
}
