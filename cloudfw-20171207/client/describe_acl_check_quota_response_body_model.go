// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclCheckQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuota(v *DescribeAclCheckQuotaResponseBodyQuota) *DescribeAclCheckQuotaResponseBody
	GetQuota() *DescribeAclCheckQuotaResponseBodyQuota
	SetRequestId(v string) *DescribeAclCheckQuotaResponseBody
	GetRequestId() *string
}

type DescribeAclCheckQuotaResponseBody struct {
	Quota *DescribeAclCheckQuotaResponseBodyQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	// example:
	//
	// 7D5483BF-2262-586D-8706-BDDB8B42****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAclCheckQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclCheckQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAclCheckQuotaResponseBody) GetQuota() *DescribeAclCheckQuotaResponseBodyQuota {
	return s.Quota
}

func (s *DescribeAclCheckQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAclCheckQuotaResponseBody) SetQuota(v *DescribeAclCheckQuotaResponseBodyQuota) *DescribeAclCheckQuotaResponseBody {
	s.Quota = v
	return s
}

func (s *DescribeAclCheckQuotaResponseBody) SetRequestId(v string) *DescribeAclCheckQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAclCheckQuotaResponseBody) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAclCheckQuotaResponseBodyQuota struct {
	// example:
	//
	// 500
	AvailableQuota *int64 `json:"AvailableQuota,omitempty" xml:"AvailableQuota,omitempty"`
	// example:
	//
	// 1500
	ConsumedQuota *int64 `json:"ConsumedQuota,omitempty" xml:"ConsumedQuota,omitempty"`
	// example:
	//
	// 2000
	TotalQuota *int64 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// example:
	//
	// 1724982259
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeAclCheckQuotaResponseBodyQuota) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclCheckQuotaResponseBodyQuota) GoString() string {
	return s.String()
}

func (s *DescribeAclCheckQuotaResponseBodyQuota) GetAvailableQuota() *int64 {
	return s.AvailableQuota
}

func (s *DescribeAclCheckQuotaResponseBodyQuota) GetConsumedQuota() *int64 {
	return s.ConsumedQuota
}

func (s *DescribeAclCheckQuotaResponseBodyQuota) GetTotalQuota() *int64 {
	return s.TotalQuota
}

func (s *DescribeAclCheckQuotaResponseBodyQuota) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeAclCheckQuotaResponseBodyQuota) SetAvailableQuota(v int64) *DescribeAclCheckQuotaResponseBodyQuota {
	s.AvailableQuota = &v
	return s
}

func (s *DescribeAclCheckQuotaResponseBodyQuota) SetConsumedQuota(v int64) *DescribeAclCheckQuotaResponseBodyQuota {
	s.ConsumedQuota = &v
	return s
}

func (s *DescribeAclCheckQuotaResponseBodyQuota) SetTotalQuota(v int64) *DescribeAclCheckQuotaResponseBodyQuota {
	s.TotalQuota = &v
	return s
}

func (s *DescribeAclCheckQuotaResponseBodyQuota) SetUpdateTime(v string) *DescribeAclCheckQuotaResponseBodyQuota {
	s.UpdateTime = &v
	return s
}

func (s *DescribeAclCheckQuotaResponseBodyQuota) Validate() error {
	return dara.Validate(s)
}
