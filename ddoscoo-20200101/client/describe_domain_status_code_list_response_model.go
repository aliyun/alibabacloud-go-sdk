// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainStatusCodeListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainStatusCodeListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainStatusCodeListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainStatusCodeListResponseBody) *DescribeDomainStatusCodeListResponse
	GetBody() *DescribeDomainStatusCodeListResponseBody
}

type DescribeDomainStatusCodeListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainStatusCodeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainStatusCodeListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatusCodeListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainStatusCodeListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainStatusCodeListResponse) GetBody() *DescribeDomainStatusCodeListResponseBody {
	return s.Body
}

func (s *DescribeDomainStatusCodeListResponse) SetHeaders(v map[string]*string) *DescribeDomainStatusCodeListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainStatusCodeListResponse) SetStatusCode(v int32) *DescribeDomainStatusCodeListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponse) SetBody(v *DescribeDomainStatusCodeListResponseBody) *DescribeDomainStatusCodeListResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainStatusCodeListResponse) Validate() error {
	return dara.Validate(s)
}
