// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserCertificateExpireCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExpireWithin30DaysCount(v int32) *DescribeUserCertificateExpireCountResponseBody
	GetExpireWithin30DaysCount() *int32
	SetExpiredCount(v int32) *DescribeUserCertificateExpireCountResponseBody
	GetExpiredCount() *int32
	SetRequestId(v string) *DescribeUserCertificateExpireCountResponseBody
	GetRequestId() *string
}

type DescribeUserCertificateExpireCountResponseBody struct {
	// The number of domain names whose SSL certificates are about to expires within 30 days.
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

func (s DescribeUserCertificateExpireCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserCertificateExpireCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserCertificateExpireCountResponseBody) GetExpireWithin30DaysCount() *int32 {
	return s.ExpireWithin30DaysCount
}

func (s *DescribeUserCertificateExpireCountResponseBody) GetExpiredCount() *int32 {
	return s.ExpiredCount
}

func (s *DescribeUserCertificateExpireCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserCertificateExpireCountResponseBody) SetExpireWithin30DaysCount(v int32) *DescribeUserCertificateExpireCountResponseBody {
	s.ExpireWithin30DaysCount = &v
	return s
}

func (s *DescribeUserCertificateExpireCountResponseBody) SetExpiredCount(v int32) *DescribeUserCertificateExpireCountResponseBody {
	s.ExpiredCount = &v
	return s
}

func (s *DescribeUserCertificateExpireCountResponseBody) SetRequestId(v string) *DescribeUserCertificateExpireCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserCertificateExpireCountResponseBody) Validate() error {
	return dara.Validate(s)
}
