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
	// The details of the Anti-DDoS Origin instances.
	InstanceList []*DescribeInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 381D5D33-BB8F-395F-8EE4-AE3BB4B523C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of Anti-DDoS Origin instances returned.
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
	// The automatic binding condition.
	AutoProtectCondition *DescribeInstanceListResponseBodyInstanceListAutoProtectCondition `json:"AutoProtectCondition,omitempty" xml:"AutoProtectCondition,omitempty" type:"Struct"`
	// Indicates whether auto-renewal is enabled for the instance. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The number of assets that are assigned public IP addresses protected by the instance that are in blackhole filtering status.
	//
	// > You can invoke [DeleteBlackhole](https://help.aliyun.com/document_detail/118692.html) to deactivate blackhole filtering for a single protected IP address.
	//
	// example:
	//
	// 0
	BlackholdingCount *string `json:"BlackholdingCount,omitempty" xml:"BlackholdingCount,omitempty"`
	// The commodity type of the instance.
	//
	// - **ddos_ddosorigin_public_cn**: Anti-DDoS Origin 2.0 (Pay-as-you-go) China site.
	//
	// - **ddos_ddosorigin_public_intl**: Anti-DDoS Origin 2.0 (Pay-as-you-go) International site.
	//
	// example:
	//
	// ddos_ddosorigin_public_cn
	CommodityType *string `json:"CommodityType,omitempty" xml:"CommodityType,omitempty"`
	// The asset overwrite type of the instance.
	//
	// - **1**: Supports assets that are assigned public IP addresses in multiple regions worldwide.
	//
	// - **2**: Supports assets that are assigned public IP addresses in multiple regions in the Chinese mainland.
	//
	// - **3**: Supports assets that are assigned public IP addresses in multiple regions outside the Chinese mainland.
	//
	// - **4**: Supports assets that are assigned public IP addresses in a single region worldwide.
	//
	// example:
	//
	// 1
	CoverageType *int32 `json:"CoverageType,omitempty" xml:"CoverageType,omitempty"`
	// The overdue payment status. Valid values:
	//
	// - **0**: No overdue payment.
	//
	// - **1**: Overdue payment.
	//
	// example:
	//
	// 0
	DebtStatus *int64 `json:"DebtStatus,omitempty" xml:"DebtStatus,omitempty"`
	// The expiration time of the instance. The value is a timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1640275200000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The purchase time of the instance. The value is a timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1592886047000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ddosbgp-cn-oew1pjrk****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The mitigation plan type of the instance. Valid values:
	//
	// - **0**: Professional.
	//
	// - **1**: Enterprise.
	//
	// example:
	//
	// 1
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The protocol type of the IP assets protected by the instance. Valid values:
	//
	// - **IPv4**: IPv4 protocol.
	//
	// - **IPv6**: IPv6 protocol.
	//
	// example:
	//
	// IPv4
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	LogExt *string `json:"LogExt,omitempty" xml:"LogExt,omitempty"`
	// The type of the cloud service associated with the instance. This parameter is not returned by default. It is returned only when the Anti-DDoS Origin instance is created by another cloud service, with the corresponding cloud service code.
	//
	// Valid values:
	//
	// - **gamebox**: The Anti-DDoS Origin instance is created by Game Security Box.
	//
	// - **eip**: The Anti-DDoS Origin instance is created by an EIP with Anti-DDoS (Enhanced) enabled.
	//
	// example:
	//
	// gamebox
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The remark of the instance.
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
	// The status of the instance. Valid values:
	//
	// - **1**: Normal.
	//
	// - **2**: Expired.
	//
	// - **3**: Released.
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

func (s *DescribeInstanceListResponseBodyInstanceList) GetLogExt() *string {
	return s.LogExt
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

func (s *DescribeInstanceListResponseBodyInstanceList) SetLogExt(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.LogExt = &v
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
	// The events on which automatic binding is based.
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
