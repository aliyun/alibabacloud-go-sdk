// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainAttackEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainAttackEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainAttackEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainAttackEventsResponseBody) *DescribeDomainAttackEventsResponse
	GetBody() *DescribeDomainAttackEventsResponseBody
}

type DescribeDomainAttackEventsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainAttackEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainAttackEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAttackEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainAttackEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainAttackEventsResponse) GetBody() *DescribeDomainAttackEventsResponseBody {
	return s.Body
}

func (s *DescribeDomainAttackEventsResponse) SetHeaders(v map[string]*string) *DescribeDomainAttackEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainAttackEventsResponse) SetStatusCode(v int32) *DescribeDomainAttackEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainAttackEventsResponse) SetBody(v *DescribeDomainAttackEventsResponseBody) *DescribeDomainAttackEventsResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainAttackEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
