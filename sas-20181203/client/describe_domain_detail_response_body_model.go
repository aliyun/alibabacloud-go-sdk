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
	// The asset information related to the domain name.
	DomainDetailItems []*DescribeDomainDetailResponseBodyDomainDetailItems `json:"DomainDetailItems,omitempty" xml:"DomainDetailItems,omitempty" type:"Repeated"`
	// The request ID. The China value is a unique identifier that Alibaba Cloud generates for the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 3A85CFCF-05C8-451A-9E41-C0D5E96BA407
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The root domain name that corresponds to the domain name.
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
	if s.DomainDetailItems != nil {
		for _, item := range s.DomainDetailItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainDetailResponseBodyDomainDetailItems struct {
	// The Asset Type of the asset under the domain name. Valid values:
	//
	// - **0**: ECS
	//
	// - **1**: load balancing
	//
	// - **2**: NAT gateway
	//
	// - **3**: RDS database
	//
	// - **4**: MongoDB database
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
	// The UUID of the asset instance.
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
