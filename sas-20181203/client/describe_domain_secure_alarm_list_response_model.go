// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecureAlarmListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainSecureAlarmListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainSecureAlarmListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainSecureAlarmListResponseBody) *DescribeDomainSecureAlarmListResponse
	GetBody() *DescribeDomainSecureAlarmListResponseBody
}

type DescribeDomainSecureAlarmListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainSecureAlarmListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainSecureAlarmListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureAlarmListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureAlarmListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainSecureAlarmListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainSecureAlarmListResponse) GetBody() *DescribeDomainSecureAlarmListResponseBody {
	return s.Body
}

func (s *DescribeDomainSecureAlarmListResponse) SetHeaders(v map[string]*string) *DescribeDomainSecureAlarmListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSecureAlarmListResponse) SetStatusCode(v int32) *DescribeDomainSecureAlarmListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponse) SetBody(v *DescribeDomainSecureAlarmListResponseBody) *DescribeDomainSecureAlarmListResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainSecureAlarmListResponse) Validate() error {
	return dara.Validate(s)
}
