// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIPv6TranslatorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6Translators(v *DescribeIPv6TranslatorsResponseBodyIpv6Translators) *DescribeIPv6TranslatorsResponseBody
	GetIpv6Translators() *DescribeIPv6TranslatorsResponseBodyIpv6Translators
	SetPageNumber(v int32) *DescribeIPv6TranslatorsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIPv6TranslatorsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeIPv6TranslatorsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeIPv6TranslatorsResponseBody
	GetTotalCount() *int32
}

type DescribeIPv6TranslatorsResponseBody struct {
	// The list of IPv6 Translation Service instances.
	Ipv6Translators *DescribeIPv6TranslatorsResponseBodyIpv6Translators `json:"Ipv6Translators,omitempty" xml:"Ipv6Translators,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeIPv6TranslatorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorsResponseBody) GetIpv6Translators() *DescribeIPv6TranslatorsResponseBodyIpv6Translators {
	return s.Ipv6Translators
}

func (s *DescribeIPv6TranslatorsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIPv6TranslatorsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIPv6TranslatorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIPv6TranslatorsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeIPv6TranslatorsResponseBody) SetIpv6Translators(v *DescribeIPv6TranslatorsResponseBodyIpv6Translators) *DescribeIPv6TranslatorsResponseBody {
	s.Ipv6Translators = v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBody) SetPageNumber(v int32) *DescribeIPv6TranslatorsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBody) SetPageSize(v int32) *DescribeIPv6TranslatorsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBody) SetRequestId(v string) *DescribeIPv6TranslatorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBody) SetTotalCount(v int32) *DescribeIPv6TranslatorsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBody) Validate() error {
	if s.Ipv6Translators != nil {
		if err := s.Ipv6Translators.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIPv6TranslatorsResponseBodyIpv6Translators struct {
	Ipv6Translator []*DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator `json:"Ipv6Translator,omitempty" xml:"Ipv6Translator,omitempty" type:"Repeated"`
}

func (s DescribeIPv6TranslatorsResponseBodyIpv6Translators) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorsResponseBodyIpv6Translators) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6Translators) GetIpv6Translator() []*DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator {
	return s.Ipv6Translator
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6Translators) SetIpv6Translator(v []*DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) *DescribeIPv6TranslatorsResponseBodyIpv6Translators {
	s.Ipv6Translator = v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6Translators) Validate() error {
	if s.Ipv6Translator != nil {
		for _, item := range s.Ipv6Translator {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator struct {
	// The IPv4 address allocated to the IPv6 Translation Service instance.
	//
	// example:
	//
	// 47.99.XX.XX
	AllocateIpv4Addr *string `json:"AllocateIpv4Addr,omitempty" xml:"AllocateIpv4Addr,omitempty"`
	// The IPv6 address allocated to the IPv6 Translation Service instance.
	//
	// example:
	//
	// 2400:3200:1600::XXXX
	AllocateIpv6Addr *string `json:"AllocateIpv6Addr,omitempty" xml:"AllocateIpv6Addr,omitempty"`
	// The bandwidth of the IPv6 Translation Service instance.
	//
	// example:
	//
	// 1
	AvailableBandwidth *string `json:"AvailableBandwidth,omitempty" xml:"AvailableBandwidth,omitempty"`
	// The bandwidth of the IPv6 Translation Service instance. Unit: Mbit/s.
	//
	// example:
	//
	// 1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The business status of the IPv6 Translation Service instance. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **FinancialLocked**
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The timestamp when the IPv6 Translation Service instance was created.
	//
	// example:
	//
	// 1537151540000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the IPv6 Translation Service instance.
	//
	// example:
	//
	// descriptionforinstance
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timestamp when IPv6 Translation Service instance expires.
	//
	// example:
	//
	// 1539792000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IDs of IPv6 mapping entries of the IPv6 Translation Service instance.
	Ipv6TranslatorEntryIds *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6TranslatorIpv6TranslatorEntryIds `json:"Ipv6TranslatorEntryIds,omitempty" xml:"Ipv6TranslatorEntryIds,omitempty" type:"Struct"`
	// The ID of the IPv6 Translation Service instance.
	//
	// example:
	//
	// ipv6trans-bp1858ys*****
	Ipv6TranslatorId *string `json:"Ipv6TranslatorId,omitempty" xml:"Ipv6TranslatorId,omitempty"`
	// The name of the IPv6 Translation Service instance.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The billing method of the IPv6 Translation Service instance.
	//
	// 	- **Prepay**: subscription
	//
	// 	- **Postpay**: pay-as-you-go
	//
	// example:
	//
	// Prepay
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region of the IPv6 Translation Service instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The specification of the IPv6 Translation Service instance.
	//
	// example:
	//
	// small
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The status of the IPv6 Translation Service instance.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) GetAllocateIpv4Addr() *string {
	return s.AllocateIpv4Addr
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) GetAllocateIpv6Addr() *string {
	return s.AllocateIpv6Addr
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) GetAvailableBandwidth() *string {
	return s.AvailableBandwidth
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) GetDescription() *string {
	return s.Description
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) GetIpv6TranslatorEntryIds() *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6TranslatorIpv6TranslatorEntryIds {
	return s.Ipv6TranslatorEntryIds
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) GetIpv6TranslatorId() *string {
	return s.Ipv6TranslatorId
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) GetName() *string {
	return s.Name
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) GetPayType() *string {
	return s.PayType
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) GetSpec() *string {
	return s.Spec
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) GetStatus() *string {
	return s.Status
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) SetAllocateIpv4Addr(v string) *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator {
	s.AllocateIpv4Addr = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) SetAllocateIpv6Addr(v string) *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator {
	s.AllocateIpv6Addr = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) SetAvailableBandwidth(v string) *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator {
	s.AvailableBandwidth = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) SetBandwidth(v int32) *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator {
	s.Bandwidth = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) SetBusinessStatus(v string) *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) SetCreateTime(v int64) *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator {
	s.CreateTime = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) SetDescription(v string) *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator {
	s.Description = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) SetEndTime(v int64) *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator {
	s.EndTime = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) SetIpv6TranslatorEntryIds(v *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6TranslatorIpv6TranslatorEntryIds) *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator {
	s.Ipv6TranslatorEntryIds = v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) SetIpv6TranslatorId(v string) *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator {
	s.Ipv6TranslatorId = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) SetName(v string) *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator {
	s.Name = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) SetPayType(v string) *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator {
	s.PayType = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) SetRegionId(v string) *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator {
	s.RegionId = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) SetSpec(v string) *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator {
	s.Spec = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) SetStatus(v string) *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator {
	s.Status = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6Translator) Validate() error {
	if s.Ipv6TranslatorEntryIds != nil {
		if err := s.Ipv6TranslatorEntryIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6TranslatorIpv6TranslatorEntryIds struct {
	Ipv6TranslatorEntryId []*string `json:"Ipv6TranslatorEntryId,omitempty" xml:"Ipv6TranslatorEntryId,omitempty" type:"Repeated"`
}

func (s DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6TranslatorIpv6TranslatorEntryIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6TranslatorIpv6TranslatorEntryIds) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6TranslatorIpv6TranslatorEntryIds) GetIpv6TranslatorEntryId() []*string {
	return s.Ipv6TranslatorEntryId
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6TranslatorIpv6TranslatorEntryIds) SetIpv6TranslatorEntryId(v []*string) *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6TranslatorIpv6TranslatorEntryIds {
	s.Ipv6TranslatorEntryId = v
	return s
}

func (s *DescribeIPv6TranslatorsResponseBodyIpv6TranslatorsIpv6TranslatorIpv6TranslatorEntryIds) Validate() error {
	return dara.Validate(s)
}
