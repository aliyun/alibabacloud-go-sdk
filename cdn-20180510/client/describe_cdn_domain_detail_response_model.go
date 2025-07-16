// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnDomainDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnDomainDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnDomainDetailResponseBody) *DescribeCdnDomainDetailResponse
	GetBody() *DescribeCdnDomainDetailResponseBody
}

type DescribeCdnDomainDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnDomainDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnDomainDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnDomainDetailResponse) GetBody() *DescribeCdnDomainDetailResponseBody {
	return s.Body
}

func (s *DescribeCdnDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeCdnDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDomainDetailResponse) SetStatusCode(v int32) *DescribeCdnDomainDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnDomainDetailResponse) SetBody(v *DescribeCdnDomainDetailResponseBody) *DescribeCdnDomainDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnDomainDetailResponse) Validate() error {
	return dara.Validate(s)
}
