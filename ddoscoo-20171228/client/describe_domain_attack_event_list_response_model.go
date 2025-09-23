// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainAttackEventListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainAttackEventListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainAttackEventListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainAttackEventListResponseBody) *DescribeDomainAttackEventListResponse
	GetBody() *DescribeDomainAttackEventListResponseBody
}

type DescribeDomainAttackEventListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainAttackEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainAttackEventListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAttackEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainAttackEventListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainAttackEventListResponse) GetBody() *DescribeDomainAttackEventListResponseBody {
	return s.Body
}

func (s *DescribeDomainAttackEventListResponse) SetHeaders(v map[string]*string) *DescribeDomainAttackEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainAttackEventListResponse) SetStatusCode(v int32) *DescribeDomainAttackEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainAttackEventListResponse) SetBody(v *DescribeDomainAttackEventListResponseBody) *DescribeDomainAttackEventListResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainAttackEventListResponse) Validate() error {
	return dara.Validate(s)
}
