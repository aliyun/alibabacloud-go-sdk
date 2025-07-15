// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainDetailResponseBody) *DescribeLiveDomainDetailResponse
	GetBody() *DescribeLiveDomainDetailResponseBody
}

type DescribeLiveDomainDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainDetailResponse) GetBody() *DescribeLiveDomainDetailResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainDetailResponse) SetStatusCode(v int32) *DescribeLiveDomainDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainDetailResponse) SetBody(v *DescribeLiveDomainDetailResponseBody) *DescribeLiveDomainDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainDetailResponse) Validate() error {
	return dara.Validate(s)
}
