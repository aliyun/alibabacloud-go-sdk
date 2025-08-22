// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainCcActivityLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainCcActivityLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainCcActivityLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainCcActivityLogResponseBody) *DescribeDcdnDomainCcActivityLogResponse
	GetBody() *DescribeDcdnDomainCcActivityLogResponseBody
}

type DescribeDcdnDomainCcActivityLogResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainCcActivityLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainCcActivityLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainCcActivityLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCcActivityLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainCcActivityLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainCcActivityLogResponse) GetBody() *DescribeDcdnDomainCcActivityLogResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainCcActivityLogResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainCcActivityLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponse) SetStatusCode(v int32) *DescribeDcdnDomainCcActivityLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponse) SetBody(v *DescribeDcdnDomainCcActivityLogResponseBody) *DescribeDcdnDomainCcActivityLogResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainCcActivityLogResponse) Validate() error {
	return dara.Validate(s)
}
