// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserResourcePackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCdnUserResourcePackageResponseBody
	GetRequestId() *string
	SetResourcePackageInfos(v *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos) *DescribeCdnUserResourcePackageResponseBody
	GetResourcePackageInfos() *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos
}

type DescribeCdnUserResourcePackageResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 84839536-2B7E-457D-9D8C-82E6C7D4E1A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The detailed information about resource plans. The returned information is displayed in an array of ResourcePackageInfo nodes.
	ResourcePackageInfos *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos `json:"ResourcePackageInfos,omitempty" xml:"ResourcePackageInfos,omitempty" type:"Struct"`
}

func (s DescribeCdnUserResourcePackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserResourcePackageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserResourcePackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnUserResourcePackageResponseBody) GetResourcePackageInfos() *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos {
	return s.ResourcePackageInfos
}

func (s *DescribeCdnUserResourcePackageResponseBody) SetRequestId(v string) *DescribeCdnUserResourcePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBody) SetResourcePackageInfos(v *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos) *DescribeCdnUserResourcePackageResponseBody {
	s.ResourcePackageInfos = v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos struct {
	ResourcePackageInfo []*DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo `json:"ResourcePackageInfo,omitempty" xml:"ResourcePackageInfo,omitempty" type:"Repeated"`
}

func (s DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos) GetResourcePackageInfo() []*DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	return s.ResourcePackageInfo
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos) SetResourcePackageInfo(v []*DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos {
	s.ResourcePackageInfo = v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo struct {
	// The ID of the resource plan.
	//
	// example:
	//
	// cdnflowbag
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The remaining quota of the resource plan.
	//
	// 	- For a data transfer plan, the quota is measured in bytes.
	//
	// 	- For a request resource plan, the quota is measured in the number of requests.
	//
	// example:
	//
	// 10995089554629
	CurrCapacity          *string `json:"CurrCapacity,omitempty" xml:"CurrCapacity,omitempty"`
	CurrCapacityBaseUnit  *string `json:"CurrCapacityBaseUnit,omitempty" xml:"CurrCapacityBaseUnit,omitempty"`
	CurrCapacityShowUnit  *string `json:"CurrCapacityShowUnit,omitempty" xml:"CurrCapacityShowUnit,omitempty"`
	CurrCapacityShowValue *string `json:"CurrCapacityShowValue,omitempty" xml:"CurrCapacityShowValue,omitempty"`
	// The name of the resource plan.
	//
	// example:
	//
	// CDN data transfer plan (Chinese mainland)
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2018-07-01T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The total quota of the resource plan.
	//
	// 	- For a data transfer plan, the quota is measured in bytes.
	//
	// 	- For a request resource plan, the quota is measured in the number of requests.
	//
	// example:
	//
	// 536870912000
	InitCapacity          *string `json:"InitCapacity,omitempty" xml:"InitCapacity,omitempty"`
	InitCapacityBaseUnit  *string `json:"InitCapacityBaseUnit,omitempty" xml:"InitCapacityBaseUnit,omitempty"`
	InitCapacityShowUnit  *string `json:"InitCapacityShowUnit,omitempty" xml:"InitCapacityShowUnit,omitempty"`
	InitCapacityShowValue *string `json:"InitCapacityShowValue,omitempty" xml:"InitCapacityShowValue,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// FP-ilttxc23a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The effective time.
	//
	// example:
	//
	// 2017-12-05T19:10:58Z
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
	// CDN data transfer plan
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetCurrCapacity() *string {
	return s.CurrCapacity
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetCurrCapacityBaseUnit() *string {
	return s.CurrCapacityBaseUnit
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetCurrCapacityShowUnit() *string {
	return s.CurrCapacityShowUnit
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetCurrCapacityShowValue() *string {
	return s.CurrCapacityShowValue
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetInitCapacity() *string {
	return s.InitCapacity
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetInitCapacityBaseUnit() *string {
	return s.InitCapacityBaseUnit
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetInitCapacityShowUnit() *string {
	return s.InitCapacityShowUnit
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetInitCapacityShowValue() *string {
	return s.InitCapacityShowValue
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetRegion() *string {
	return s.Region
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCommodityCode(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CommodityCode = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCurrCapacity(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CurrCapacity = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCurrCapacityBaseUnit(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CurrCapacityBaseUnit = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCurrCapacityShowUnit(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CurrCapacityShowUnit = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCurrCapacityShowValue(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CurrCapacityShowValue = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetDisplayName(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.DisplayName = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetEndTime(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInitCapacity(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InitCapacity = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInitCapacityBaseUnit(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InitCapacityBaseUnit = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInitCapacityShowUnit(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InitCapacityShowUnit = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInitCapacityShowValue(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InitCapacityShowValue = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInstanceId(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetRegion(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.Region = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetStartTime(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetStatus(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.Status = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetTemplateName(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.TemplateName = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) Validate() error {
	return dara.Validate(s)
}
