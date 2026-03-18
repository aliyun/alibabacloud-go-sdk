// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceList(v []*DescribeInstanceListResponseBodyInstanceList) *DescribeInstanceListResponseBody
	GetInstanceList() []*DescribeInstanceListResponseBodyInstanceList
	SetRequestId(v string) *DescribeInstanceListResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeInstanceListResponseBody
	GetTotal() *int64
}

type DescribeInstanceListResponseBody struct {
	// The details about the Anti-DDoS Origin instances.
	InstanceList []*DescribeInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The details about the Anti-DDoS Origin instance.
	//
	// example:
	//
	// 381D5D33-BB8F-395F-8EE4-AE3BB4B523C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about the Anti-DDoS Origin instances.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstanceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponseBody) GetInstanceList() []*DescribeInstanceListResponseBodyInstanceList {
	return s.InstanceList
}

func (s *DescribeInstanceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceListResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeInstanceListResponseBody) SetInstanceList(v []*DescribeInstanceListResponseBodyInstanceList) *DescribeInstanceListResponseBody {
	s.InstanceList = v
	return s
}

func (s *DescribeInstanceListResponseBody) SetRequestId(v string) *DescribeInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceListResponseBody) SetTotal(v int64) *DescribeInstanceListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeInstanceListResponseBody) Validate() error {
	if s.InstanceList != nil {
		for _, item := range s.InstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceListResponseBodyInstanceList struct {
	// The event that triggers automatic association. Valid values:
	//
	// 	- **any**: The instance is automatically associated with an object based on traffic scrubbing events or blackhole filtering events.
	//
	// 	- **clean**: The instance is automatically associated with an object based on traffic scrubbing events.
	//
	// 	- **blackhole**: The instance is automatically associated with an object based on blackhole filtering events.
	AutoProtectCondition *DescribeInstanceListResponseBodyInstanceListAutoProtectCondition `json:"AutoProtectCondition,omitempty" xml:"AutoProtectCondition,omitempty" type:"Struct"`
	// The time when the instance expires. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The type of the instance.
	//
	// 	- **ddos_ddosorigin_public_cn**: Anti-DDoS Origin 2.0 (Pay-as-you-go) on the China site (aliyun.com).
	//
	// 	- **ddos_ddosorigin_public_intl**: Anti-DDoS Origin 2.0 (Pay-as-you-go) on the International site (alibabacloud.com).
	//
	// example:
	//
	// 0
	BlackholdingCount *string `json:"BlackholdingCount,omitempty" xml:"BlackholdingCount,omitempty"`
	// The condition that triggers automatic association of the instance with an object.
	//
	// example:
	//
	// ddos_ddosorigin_public_cn
	CommodityType *string `json:"CommodityType,omitempty" xml:"CommodityType,omitempty"`
	// Indicates whether overdue payments exist. Valid values:
	//
	// 	- **0**: Overdue payments do not exist.
	//
	// 	- **1**: Overdue payments exist.
	//
	// example:
	//
	// 1
	CoverageType *int32 `json:"CoverageType,omitempty" xml:"CoverageType,omitempty"`
	// The events that trigger automatic association.
	//
	// example:
	//
	// 0
	DebtStatus *int64 `json:"DebtStatus,omitempty" xml:"DebtStatus,omitempty"`
	// The time when the instance was purchased. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1640275200000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The mitigation plan of the instance. Valid values:
	//
	// 	- **0**: the Professional mitigation plan
	//
	// 	- **1**: the Enterprise mitigation plan
	//
	// example:
	//
	// 1592886047000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The number of protected public IP addresses for which blackhole filtering is triggered.
	//
	// >  You can call the [DeleteBlackhole](https://help.aliyun.com/document_detail/118692.html) operation to deactivate blackhole filtering for a protected IP address.
	//
	// example:
	//
	// ddosbgp-cn-oew1pjrk****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The application scope of the instance.
	//
	// 	- **1**: The instance supports public IP addresses in all regions.
	//
	// 	- **2**: The instance supports public IP addresses in regions in the Chinese mainland.
	//
	// 	- **3**: The instance supports public IP addresses in regions outside the Chinese mainland.
	//
	// 	- **4**: The instance supports public IP addresses in a region in or outside the Chinese mainland.
	//
	// example:
	//
	// 1
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// IPv4
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// gamebox
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The type of the cloud service that is associated with the Anti-DDoS Origin instance By default, this parameter is not returned. If the Anti-DDoS Origin instance is created by using a different cloud service, the code of the cloud service is returned.
	//
	// Valid values:
	//
	// 	- **gamebox**: The Anti-DDoS Origin instance is created by using Game Security Box.
	//
	// 	- **eip**: The Anti-DDoS Origin instance is created by using an elastic IP address (EIP) for which Anti-DDoS (Enhanced Edition) is enabled.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek3ccjxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether auto-renewal is enabled for the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceListResponseBodyInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceListResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponseBodyInstanceList) GetAutoProtectCondition() *DescribeInstanceListResponseBodyInstanceListAutoProtectCondition {
	return s.AutoProtectCondition
}

func (s *DescribeInstanceListResponseBodyInstanceList) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *DescribeInstanceListResponseBodyInstanceList) GetBlackholdingCount() *string {
	return s.BlackholdingCount
}

func (s *DescribeInstanceListResponseBodyInstanceList) GetCommodityType() *string {
	return s.CommodityType
}

func (s *DescribeInstanceListResponseBodyInstanceList) GetCoverageType() *int32 {
	return s.CoverageType
}

func (s *DescribeInstanceListResponseBodyInstanceList) GetDebtStatus() *int64 {
	return s.DebtStatus
}

func (s *DescribeInstanceListResponseBodyInstanceList) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeInstanceListResponseBodyInstanceList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeInstanceListResponseBodyInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceListResponseBodyInstanceList) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceListResponseBodyInstanceList) GetIpType() *string {
	return s.IpType
}

func (s *DescribeInstanceListResponseBodyInstanceList) GetProduct() *string {
	return s.Product
}

func (s *DescribeInstanceListResponseBodyInstanceList) GetRemark() *string {
	return s.Remark
}

func (s *DescribeInstanceListResponseBodyInstanceList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstanceListResponseBodyInstanceList) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetAutoProtectCondition(v *DescribeInstanceListResponseBodyInstanceListAutoProtectCondition) *DescribeInstanceListResponseBodyInstanceList {
	s.AutoProtectCondition = v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetAutoRenewal(v bool) *DescribeInstanceListResponseBodyInstanceList {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetBlackholdingCount(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.BlackholdingCount = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetCommodityType(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.CommodityType = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetCoverageType(v int32) *DescribeInstanceListResponseBodyInstanceList {
	s.CoverageType = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetDebtStatus(v int64) *DescribeInstanceListResponseBodyInstanceList {
	s.DebtStatus = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetExpireTime(v int64) *DescribeInstanceListResponseBodyInstanceList {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetGmtCreate(v int64) *DescribeInstanceListResponseBodyInstanceList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetInstanceId(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetInstanceType(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetIpType(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.IpType = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetProduct(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Product = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetRemark(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Remark = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetResourceGroupId(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetStatus(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Status = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) Validate() error {
	if s.AutoProtectCondition != nil {
		if err := s.AutoProtectCondition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceListResponseBodyInstanceListAutoProtectCondition struct {
	// Events which result in auto binding.
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
}

func (s DescribeInstanceListResponseBodyInstanceListAutoProtectCondition) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceListResponseBodyInstanceListAutoProtectCondition) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponseBodyInstanceListAutoProtectCondition) GetEvents() []*string {
	return s.Events
}

func (s *DescribeInstanceListResponseBodyInstanceListAutoProtectCondition) SetEvents(v []*string) *DescribeInstanceListResponseBodyInstanceListAutoProtectCondition {
	s.Events = v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceListAutoProtectCondition) Validate() error {
	return dara.Validate(s)
}
