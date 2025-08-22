// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserCertificateExpireCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExpireWithin30DaysCount(v int32) *DescribeDcdnUserCertificateExpireCountResponseBody
	GetExpireWithin30DaysCount() *int32
	SetExpiredCount(v int32) *DescribeDcdnUserCertificateExpireCountResponseBody
	GetExpiredCount() *int32
	SetRequestId(v string) *DescribeDcdnUserCertificateExpireCountResponseBody
	GetRequestId() *string
}

type DescribeDcdnUserCertificateExpireCountResponseBody struct {
	// The number of domain names whose SSL certificates are about to expire within 30 days.
	//
	// example:
	//
	// 0
	ExpireWithin30DaysCount *int32 `json:"ExpireWithin30DaysCount,omitempty" xml:"ExpireWithin30DaysCount,omitempty"`
	// The number of domain names whose SSL certificates have already expired.
	//
	// example:
	//
	// 6
	ExpiredCount *int32 `json:"ExpiredCount,omitempty" xml:"ExpiredCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F5E8DF64-7175-4186-9B06-F002C0BBD0C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnUserCertificateExpireCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserCertificateExpireCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserCertificateExpireCountResponseBody) GetExpireWithin30DaysCount() *int32 {
	return s.ExpireWithin30DaysCount
}

func (s *DescribeDcdnUserCertificateExpireCountResponseBody) GetExpiredCount() *int32 {
	return s.ExpiredCount
}

func (s *DescribeDcdnUserCertificateExpireCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnUserCertificateExpireCountResponseBody) SetExpireWithin30DaysCount(v int32) *DescribeDcdnUserCertificateExpireCountResponseBody {
	s.ExpireWithin30DaysCount = &v
	return s
}

func (s *DescribeDcdnUserCertificateExpireCountResponseBody) SetExpiredCount(v int32) *DescribeDcdnUserCertificateExpireCountResponseBody {
	s.ExpiredCount = &v
	return s
}

func (s *DescribeDcdnUserCertificateExpireCountResponseBody) SetRequestId(v string) *DescribeDcdnUserCertificateExpireCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserCertificateExpireCountResponseBody) Validate() error {
	return dara.Validate(s)
}
