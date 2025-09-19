// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmCount(v int32) *DescribeDomainDetailResponseBody
	GetAlarmCount() *int32
	SetDomain(v string) *DescribeDomainDetailResponseBody
	GetDomain() *string
	SetDomainDetailItems(v []*DescribeDomainDetailResponseBodyDomainDetailItems) *DescribeDomainDetailResponseBody
	GetDomainDetailItems() []*DescribeDomainDetailResponseBodyDomainDetailItems
	SetRequestId(v string) *DescribeDomainDetailResponseBody
	GetRequestId() *string
	SetRootDomain(v string) *DescribeDomainDetailResponseBody
	GetRootDomain() *string
	SetVulCount(v int32) *DescribeDomainDetailResponseBody
	GetVulCount() *int32
}

type DescribeDomainDetailResponseBody struct {
	// The total number of alerts in your website assets.
	//
	// example:
	//
	// 2
	AlarmCount *int32 `json:"AlarmCount,omitempty" xml:"AlarmCount,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// An array that consists of the details about the domain asset.
	DomainDetailItems []*DescribeDomainDetailResponseBodyDomainDetailItems `json:"DomainDetailItems,omitempty" xml:"DomainDetailItems,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3A85CFCF-05C8-451A-9E41-C0D5E96BA407
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the root domain that corresponds to the domain.
	//
	// example:
	//
	// example.com
	RootDomain *string `json:"RootDomain,omitempty" xml:"RootDomain,omitempty"`
	// The total number of vulnerabilities in your website assets.
	//
	// example:
	//
	// 2
	VulCount *int32 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
}

func (s DescribeDomainDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBody) GetAlarmCount() *int32 {
	return s.AlarmCount
}

func (s *DescribeDomainDetailResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainDetailResponseBody) GetDomainDetailItems() []*DescribeDomainDetailResponseBodyDomainDetailItems {
	return s.DomainDetailItems
}

func (s *DescribeDomainDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainDetailResponseBody) GetRootDomain() *string {
	return s.RootDomain
}

func (s *DescribeDomainDetailResponseBody) GetVulCount() *int32 {
	return s.VulCount
}

func (s *DescribeDomainDetailResponseBody) SetAlarmCount(v int32) *DescribeDomainDetailResponseBody {
	s.AlarmCount = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetDomain(v string) *DescribeDomainDetailResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetDomainDetailItems(v []*DescribeDomainDetailResponseBodyDomainDetailItems) *DescribeDomainDetailResponseBody {
	s.DomainDetailItems = v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetRequestId(v string) *DescribeDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetRootDomain(v string) *DescribeDomainDetailResponseBody {
	s.RootDomain = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetVulCount(v int32) *DescribeDomainDetailResponseBody {
	s.VulCount = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainDetailResponseBodyDomainDetailItems struct {
	// The type of the domain asset. Valid values:
	//
	// 	- **0**: an Elastic Compute Service (ECS) instance
	//
	// 	- **1**: a Server Load Balancer (SLB) instance
	//
	// 	- **2**: a Network Address Translation (NAT) gateway
	//
	// 	- **3**: an ApsaraDB RDS instance
	//
	// 	- **4**: an ApsaraDB for MongoDB instance
	//
	// example:
	//
	// 0
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The instance ID of the asset.
	//
	// example:
	//
	// i-m5e6w7dzsktt6mz4***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// iZm5e6w7dzsktt6mz4yimeZ-6****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the asset.
	//
	// example:
	//
	// 1.2.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the asset.
	//
	// example:
	//
	// 1.2.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The instance UUID of the domain asset.
	//
	// example:
	//
	// lb-bp1g9dohoyin9cjhn6****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeDomainDetailResponseBodyDomainDetailItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyDomainDetailItems) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) GetAssetType() *string {
	return s.AssetType
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) SetAssetType(v string) *DescribeDomainDetailResponseBodyDomainDetailItems {
	s.AssetType = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) SetInstanceId(v string) *DescribeDomainDetailResponseBodyDomainDetailItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) SetInstanceName(v string) *DescribeDomainDetailResponseBodyDomainDetailItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) SetInternetIp(v string) *DescribeDomainDetailResponseBodyDomainDetailItems {
	s.InternetIp = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) SetIntranetIp(v string) *DescribeDomainDetailResponseBodyDomainDetailItems {
	s.IntranetIp = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) SetUuid(v string) *DescribeDomainDetailResponseBodyDomainDetailItems {
	s.Uuid = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) Validate() error {
	return dara.Validate(s)
}
