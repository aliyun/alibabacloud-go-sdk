// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserResourcePackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDcdnUserResourcePackageResponseBody
	GetRequestId() *string
	SetResourcePackageInfos(v *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos) *DescribeDcdnUserResourcePackageResponseBody
	GetResourcePackageInfos() *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos
}

type DescribeDcdnUserResourcePackageResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 96ED3127-EC7A-57C5-AFA6-A689B24B2530
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The detailed information about resource plans. The returned information is displayed in the format that is specified by the ResourcePackageInfo parameter.
	ResourcePackageInfos *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos `json:"ResourcePackageInfos,omitempty" xml:"ResourcePackageInfos,omitempty" type:"Struct"`
}

func (s DescribeDcdnUserResourcePackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserResourcePackageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserResourcePackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnUserResourcePackageResponseBody) GetResourcePackageInfos() *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos {
	return s.ResourcePackageInfos
}

func (s *DescribeDcdnUserResourcePackageResponseBody) SetRequestId(v string) *DescribeDcdnUserResourcePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBody) SetResourcePackageInfos(v *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos) *DescribeDcdnUserResourcePackageResponseBody {
	s.ResourcePackageInfos = v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos struct {
	ResourcePackageInfo []*DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo `json:"ResourcePackageInfo,omitempty" xml:"ResourcePackageInfo,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos) GetResourcePackageInfo() []*DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	return s.ResourcePackageInfo
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos) SetResourcePackageInfo(v []*DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos {
	s.ResourcePackageInfo = v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo struct {
	// The commodity code of the resource plan.
	//
	// example:
	//
	// dcdnpaybag
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The remaining quota of the resource plan.
	//
	// 	- The unit for traffic: bytes.
	//
	// 	- The unit for requests: count.
	//
	// example:
	//
	// 10000000
	CurrCapacity          *string `json:"CurrCapacity,omitempty" xml:"CurrCapacity,omitempty"`
	CurrCapacityBaseUnit  *string `json:"CurrCapacityBaseUnit,omitempty" xml:"CurrCapacityBaseUnit,omitempty"`
	CurrCapacityShowUnit  *string `json:"CurrCapacityShowUnit,omitempty" xml:"CurrCapacityShowUnit,omitempty"`
	CurrCapacityShowValue *string `json:"CurrCapacityShowValue,omitempty" xml:"CurrCapacityShowValue,omitempty"`
	// The name of the resource plan.
	//
	// example:
	//
	// HTTPS requests for static content
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The expiration time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-08-24T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The total quota of the resource plan.
	//
	// 	- The unit for traffic: bytes.
	//
	// 	- The unit for requests: count.
	//
	// example:
	//
	// 10000000
	InitCapacity          *string `json:"InitCapacity,omitempty" xml:"InitCapacity,omitempty"`
	InitCapacityBaseUnit  *string `json:"InitCapacityBaseUnit,omitempty" xml:"InitCapacityBaseUnit,omitempty"`
	InitCapacityShowUnit  *string `json:"InitCapacityShowUnit,omitempty" xml:"InitCapacityShowUnit,omitempty"`
	InitCapacityShowValue *string `json:"InitCapacityShowValue,omitempty" xml:"InitCapacityShowValue,omitempty"`
	// The ID of the resource plan.
	//
	// example:
	//
	// CDNFLOWBAG-cn-7pp2bihrb01ii0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The validation time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2021-08-24T04:09:22Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the resource plan. Valid values:
	//
	// 	- **valid**: valid
	//
	// 	- **closed**: expired
	//
	// example:
	//
	// valid
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// FPT_dcdnpaybag_deadlineAcc_1541151058
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetCurrCapacity() *string {
	return s.CurrCapacity
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetCurrCapacityBaseUnit() *string {
	return s.CurrCapacityBaseUnit
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetCurrCapacityShowUnit() *string {
	return s.CurrCapacityShowUnit
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetCurrCapacityShowValue() *string {
	return s.CurrCapacityShowValue
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetInitCapacity() *string {
	return s.InitCapacity
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetInitCapacityBaseUnit() *string {
	return s.InitCapacityBaseUnit
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetInitCapacityShowUnit() *string {
	return s.InitCapacityShowUnit
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetInitCapacityShowValue() *string {
	return s.InitCapacityShowValue
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetRegion() *string {
	return s.Region
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCommodityCode(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCurrCapacity(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CurrCapacity = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCurrCapacityBaseUnit(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CurrCapacityBaseUnit = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCurrCapacityShowUnit(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CurrCapacityShowUnit = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCurrCapacityShowValue(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CurrCapacityShowValue = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetDisplayName(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.DisplayName = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetEndTime(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInitCapacity(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InitCapacity = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInitCapacityBaseUnit(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InitCapacityBaseUnit = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInitCapacityShowUnit(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InitCapacityShowUnit = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInitCapacityShowValue(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InitCapacityShowValue = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInstanceId(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetRegion(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.Region = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetStartTime(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetStatus(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.Status = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetTemplateName(v string) *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.TemplateName = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) Validate() error {
	return dara.Validate(s)
}
