// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainAttackMaxQpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainAttackMaxQpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainAttackMaxQpsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainAttackMaxQpsResponseBody) *DescribeDomainAttackMaxQpsResponse
	GetBody() *DescribeDomainAttackMaxQpsResponseBody
}

type DescribeDomainAttackMaxQpsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainAttackMaxQpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainAttackMaxQpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAttackMaxQpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackMaxQpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainAttackMaxQpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainAttackMaxQpsResponse) GetBody() *DescribeDomainAttackMaxQpsResponseBody {
	return s.Body
}

func (s *DescribeDomainAttackMaxQpsResponse) SetHeaders(v map[string]*string) *DescribeDomainAttackMaxQpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainAttackMaxQpsResponse) SetStatusCode(v int32) *DescribeDomainAttackMaxQpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainAttackMaxQpsResponse) SetBody(v *DescribeDomainAttackMaxQpsResponseBody) *DescribeDomainAttackMaxQpsResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainAttackMaxQpsResponse) Validate() error {
	return dara.Validate(s)
}
