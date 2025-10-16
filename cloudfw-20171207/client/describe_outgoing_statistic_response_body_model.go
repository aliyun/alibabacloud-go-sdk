// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoreAssetCount(v int32) *DescribeOutgoingStatisticResponseBody
	GetIgnoreAssetCount() *int32
	SetIgnoreDomainCount(v int32) *DescribeOutgoingStatisticResponseBody
	GetIgnoreDomainCount() *int32
	SetIgnoreDstIPCount(v int32) *DescribeOutgoingStatisticResponseBody
	GetIgnoreDstIPCount() *int32
	SetPrivateRiskAssetCount(v int64) *DescribeOutgoingStatisticResponseBody
	GetPrivateRiskAssetCount() *int64
	SetPrivateTotalAssetCount(v int64) *DescribeOutgoingStatisticResponseBody
	GetPrivateTotalAssetCount() *int64
	SetRequestId(v string) *DescribeOutgoingStatisticResponseBody
	GetRequestId() *string
	SetRiskAssetCount(v int32) *DescribeOutgoingStatisticResponseBody
	GetRiskAssetCount() *int32
	SetRiskDomainCount(v int32) *DescribeOutgoingStatisticResponseBody
	GetRiskDomainCount() *int32
	SetRiskDstIPCount(v int32) *DescribeOutgoingStatisticResponseBody
	GetRiskDstIPCount() *int32
	SetSubscribeAssetCount(v int32) *DescribeOutgoingStatisticResponseBody
	GetSubscribeAssetCount() *int32
	SetSubscribeDomainCount(v int32) *DescribeOutgoingStatisticResponseBody
	GetSubscribeDomainCount() *int32
	SetSubscribeDstIPCount(v int32) *DescribeOutgoingStatisticResponseBody
	GetSubscribeDstIPCount() *int32
	SetTotalAssetCount(v int32) *DescribeOutgoingStatisticResponseBody
	GetTotalAssetCount() *int32
	SetTotalDomainCount(v int32) *DescribeOutgoingStatisticResponseBody
	GetTotalDomainCount() *int32
	SetTotalDstIPCount(v int32) *DescribeOutgoingStatisticResponseBody
	GetTotalDstIPCount() *int32
	SetTotalProtocolCount(v int32) *DescribeOutgoingStatisticResponseBody
	GetTotalProtocolCount() *int32
	SetUncoveredAclDomain(v int32) *DescribeOutgoingStatisticResponseBody
	GetUncoveredAclDomain() *int32
	SetUncoveredAclDstIP(v int32) *DescribeOutgoingStatisticResponseBody
	GetUncoveredAclDstIP() *int32
	SetUnknownProtocolRadio(v string) *DescribeOutgoingStatisticResponseBody
	GetUnknownProtocolRadio() *string
}

type DescribeOutgoingStatisticResponseBody struct {
	// example:
	//
	// 0
	IgnoreAssetCount *int32 `json:"IgnoreAssetCount,omitempty" xml:"IgnoreAssetCount,omitempty"`
	// example:
	//
	// 10
	IgnoreDomainCount *int32 `json:"IgnoreDomainCount,omitempty" xml:"IgnoreDomainCount,omitempty"`
	// example:
	//
	// 0
	IgnoreDstIPCount *int32 `json:"IgnoreDstIPCount,omitempty" xml:"IgnoreDstIPCount,omitempty"`
	// example:
	//
	// 0
	PrivateRiskAssetCount *int64 `json:"PrivateRiskAssetCount,omitempty" xml:"PrivateRiskAssetCount,omitempty"`
	// example:
	//
	// 0
	PrivateTotalAssetCount *int64 `json:"PrivateTotalAssetCount,omitempty" xml:"PrivateTotalAssetCount,omitempty"`
	// example:
	//
	// E2BD70F4-48BF-5EFD-B103-F0763E27*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	RiskAssetCount *int32 `json:"RiskAssetCount,omitempty" xml:"RiskAssetCount,omitempty"`
	// example:
	//
	// 1
	RiskDomainCount *int32 `json:"RiskDomainCount,omitempty" xml:"RiskDomainCount,omitempty"`
	// example:
	//
	// 1
	RiskDstIPCount *int32 `json:"RiskDstIPCount,omitempty" xml:"RiskDstIPCount,omitempty"`
	// example:
	//
	// 0
	SubscribeAssetCount *int32 `json:"SubscribeAssetCount,omitempty" xml:"SubscribeAssetCount,omitempty"`
	// example:
	//
	// 10
	SubscribeDomainCount *int32 `json:"SubscribeDomainCount,omitempty" xml:"SubscribeDomainCount,omitempty"`
	// example:
	//
	// 10
	SubscribeDstIPCount *int32 `json:"SubscribeDstIPCount,omitempty" xml:"SubscribeDstIPCount,omitempty"`
	// example:
	//
	// 13
	TotalAssetCount *int32 `json:"TotalAssetCount,omitempty" xml:"TotalAssetCount,omitempty"`
	// example:
	//
	// 10
	TotalDomainCount *int32 `json:"TotalDomainCount,omitempty" xml:"TotalDomainCount,omitempty"`
	// example:
	//
	// 107
	TotalDstIPCount *int32 `json:"TotalDstIPCount,omitempty" xml:"TotalDstIPCount,omitempty"`
	// example:
	//
	// 10
	TotalProtocolCount *int32 `json:"TotalProtocolCount,omitempty" xml:"TotalProtocolCount,omitempty"`
	// example:
	//
	// 0
	UncoveredAclDomain *int32 `json:"UncoveredAclDomain,omitempty" xml:"UncoveredAclDomain,omitempty"`
	// example:
	//
	// 10
	UncoveredAclDstIP *int32 `json:"UncoveredAclDstIP,omitempty" xml:"UncoveredAclDstIP,omitempty"`
	// example:
	//
	// 20.13
	UnknownProtocolRadio *string `json:"UnknownProtocolRadio,omitempty" xml:"UnknownProtocolRadio,omitempty"`
}

func (s DescribeOutgoingStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingStatisticResponseBody) GetIgnoreAssetCount() *int32 {
	return s.IgnoreAssetCount
}

func (s *DescribeOutgoingStatisticResponseBody) GetIgnoreDomainCount() *int32 {
	return s.IgnoreDomainCount
}

func (s *DescribeOutgoingStatisticResponseBody) GetIgnoreDstIPCount() *int32 {
	return s.IgnoreDstIPCount
}

func (s *DescribeOutgoingStatisticResponseBody) GetPrivateRiskAssetCount() *int64 {
	return s.PrivateRiskAssetCount
}

func (s *DescribeOutgoingStatisticResponseBody) GetPrivateTotalAssetCount() *int64 {
	return s.PrivateTotalAssetCount
}

func (s *DescribeOutgoingStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOutgoingStatisticResponseBody) GetRiskAssetCount() *int32 {
	return s.RiskAssetCount
}

func (s *DescribeOutgoingStatisticResponseBody) GetRiskDomainCount() *int32 {
	return s.RiskDomainCount
}

func (s *DescribeOutgoingStatisticResponseBody) GetRiskDstIPCount() *int32 {
	return s.RiskDstIPCount
}

func (s *DescribeOutgoingStatisticResponseBody) GetSubscribeAssetCount() *int32 {
	return s.SubscribeAssetCount
}

func (s *DescribeOutgoingStatisticResponseBody) GetSubscribeDomainCount() *int32 {
	return s.SubscribeDomainCount
}

func (s *DescribeOutgoingStatisticResponseBody) GetSubscribeDstIPCount() *int32 {
	return s.SubscribeDstIPCount
}

func (s *DescribeOutgoingStatisticResponseBody) GetTotalAssetCount() *int32 {
	return s.TotalAssetCount
}

func (s *DescribeOutgoingStatisticResponseBody) GetTotalDomainCount() *int32 {
	return s.TotalDomainCount
}

func (s *DescribeOutgoingStatisticResponseBody) GetTotalDstIPCount() *int32 {
	return s.TotalDstIPCount
}

func (s *DescribeOutgoingStatisticResponseBody) GetTotalProtocolCount() *int32 {
	return s.TotalProtocolCount
}

func (s *DescribeOutgoingStatisticResponseBody) GetUncoveredAclDomain() *int32 {
	return s.UncoveredAclDomain
}

func (s *DescribeOutgoingStatisticResponseBody) GetUncoveredAclDstIP() *int32 {
	return s.UncoveredAclDstIP
}

func (s *DescribeOutgoingStatisticResponseBody) GetUnknownProtocolRadio() *string {
	return s.UnknownProtocolRadio
}

func (s *DescribeOutgoingStatisticResponseBody) SetIgnoreAssetCount(v int32) *DescribeOutgoingStatisticResponseBody {
	s.IgnoreAssetCount = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetIgnoreDomainCount(v int32) *DescribeOutgoingStatisticResponseBody {
	s.IgnoreDomainCount = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetIgnoreDstIPCount(v int32) *DescribeOutgoingStatisticResponseBody {
	s.IgnoreDstIPCount = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetPrivateRiskAssetCount(v int64) *DescribeOutgoingStatisticResponseBody {
	s.PrivateRiskAssetCount = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetPrivateTotalAssetCount(v int64) *DescribeOutgoingStatisticResponseBody {
	s.PrivateTotalAssetCount = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetRequestId(v string) *DescribeOutgoingStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetRiskAssetCount(v int32) *DescribeOutgoingStatisticResponseBody {
	s.RiskAssetCount = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetRiskDomainCount(v int32) *DescribeOutgoingStatisticResponseBody {
	s.RiskDomainCount = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetRiskDstIPCount(v int32) *DescribeOutgoingStatisticResponseBody {
	s.RiskDstIPCount = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetSubscribeAssetCount(v int32) *DescribeOutgoingStatisticResponseBody {
	s.SubscribeAssetCount = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetSubscribeDomainCount(v int32) *DescribeOutgoingStatisticResponseBody {
	s.SubscribeDomainCount = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetSubscribeDstIPCount(v int32) *DescribeOutgoingStatisticResponseBody {
	s.SubscribeDstIPCount = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetTotalAssetCount(v int32) *DescribeOutgoingStatisticResponseBody {
	s.TotalAssetCount = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetTotalDomainCount(v int32) *DescribeOutgoingStatisticResponseBody {
	s.TotalDomainCount = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetTotalDstIPCount(v int32) *DescribeOutgoingStatisticResponseBody {
	s.TotalDstIPCount = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetTotalProtocolCount(v int32) *DescribeOutgoingStatisticResponseBody {
	s.TotalProtocolCount = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetUncoveredAclDomain(v int32) *DescribeOutgoingStatisticResponseBody {
	s.UncoveredAclDomain = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetUncoveredAclDstIP(v int32) *DescribeOutgoingStatisticResponseBody {
	s.UncoveredAclDstIP = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) SetUnknownProtocolRadio(v string) *DescribeOutgoingStatisticResponseBody {
	s.UnknownProtocolRadio = &v
	return s
}

func (s *DescribeOutgoingStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}
