// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopAttackListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainTopAttackListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainTopAttackListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainTopAttackListResponseBody) *DescribeDomainTopAttackListResponse
	GetBody() *DescribeDomainTopAttackListResponseBody
}

type DescribeDomainTopAttackListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainTopAttackListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainTopAttackListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopAttackListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopAttackListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainTopAttackListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainTopAttackListResponse) GetBody() *DescribeDomainTopAttackListResponseBody {
	return s.Body
}

func (s *DescribeDomainTopAttackListResponse) SetHeaders(v map[string]*string) *DescribeDomainTopAttackListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainTopAttackListResponse) SetStatusCode(v int32) *DescribeDomainTopAttackListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainTopAttackListResponse) SetBody(v *DescribeDomainTopAttackListResponseBody) *DescribeDomainTopAttackListResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainTopAttackListResponse) Validate() error {
	return dara.Validate(s)
}
