// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAssetInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*ListCloudAssetInstancesResponseBodyInstances) *ListCloudAssetInstancesResponseBody
	GetInstances() []*ListCloudAssetInstancesResponseBodyInstances
	SetPageInfo(v *ListCloudAssetInstancesResponseBodyPageInfo) *ListCloudAssetInstancesResponseBody
	GetPageInfo() *ListCloudAssetInstancesResponseBodyPageInfo
	SetRequestId(v string) *ListCloudAssetInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCloudAssetInstancesResponseBody
	GetSuccess() *bool
}

type ListCloudAssetInstancesResponseBody struct {
	// The details of the cloud assets.
	Instances []*ListCloudAssetInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListCloudAssetInstancesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 028CF634-5268-5660-9575-48C9ED6BF880
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCloudAssetInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAssetInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudAssetInstancesResponseBody) GetInstances() []*ListCloudAssetInstancesResponseBodyInstances {
	return s.Instances
}

func (s *ListCloudAssetInstancesResponseBody) GetPageInfo() *ListCloudAssetInstancesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListCloudAssetInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudAssetInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCloudAssetInstancesResponseBody) SetInstances(v []*ListCloudAssetInstancesResponseBodyInstances) *ListCloudAssetInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListCloudAssetInstancesResponseBody) SetPageInfo(v *ListCloudAssetInstancesResponseBodyPageInfo) *ListCloudAssetInstancesResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListCloudAssetInstancesResponseBody) SetRequestId(v string) *ListCloudAssetInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBody) SetSuccess(v bool) *ListCloudAssetInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudAssetInstancesResponseBodyInstances struct {
	// Indicates whether alerts are generated for the cloud asset. Valid values:
	//
	// 	- **YES**
	//
	// 	- **NO**
	//
	// example:
	//
	// NO
	AlarmStatus *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	// The subtype of the cloud service. The subtype of the cloud asset. Valid values:
	//
	// 	- **0**: ECS
	//
	//     	- **0**: instance
	//
	//     	- **1**: disk (storage)
	//
	//     	- **2**: security group
	//
	// 	- **1**: SLB
	//
	//     	- **0**: SLB
	//
	//     	- **1**: Application Load Balancer (ALB)
	//
	// 	- **3**: ApsaraDB RDS
	//
	//     	- **0**: instance
	//
	// 	- **4**: ApsaraDB for MongoDB
	//
	//     	- **0**: instance
	//
	// 	- **5**: ApsaraDB for Redis
	//
	//     	- **0**: instance
	//
	// 	- **6**: Container Registry
	//
	//     	- **1**: Enterprise Edition
	//
	//     	- **2**: Personal Edition
	//
	// 	- **8**: ACK
	//
	//     	- **0**: cluster
	//
	// 	- **9**: VPC
	//
	//     	- **0**: NAT gateway
	//
	//     	- **1**: EIP
	//
	//     	- **2**: VPN
	//
	//     	- **3**: FLOW_LOG
	//
	// 	- **11**: ActionTrail
	//
	//     	- **0**: trail
	//
	// 	- **12**: Alibaba Cloud CDN
	//
	//     	- **0**: instance
	//
	// 	- **13**: Certificate Management Service (formerly SSL Certificates Service)
	//
	//     	- **0**: certificate
	//
	// 	- **14**: Apsara Devops
	//
	//     	- **0**: organization
	//
	// 	- **16**: Anti-DDoS
	//
	//     	- **0**: instance
	//
	// 	- **17**: WAF
	//
	//     	- **0**: domain name
	//
	// 	- **18**: OSS
	//
	//     	- **0**: bucket
	//
	// 	- **19**: PolarDB
	//
	//     	- **0**: cluster
	//
	// 	- **20**: ApsaraDB RDS for PostgreSQL
	//
	//     	- **0**: instance
	//
	// 	- **21**: MSE
	//
	//     	- **0**: cluster
	//
	// 	- **22**: NAS
	//
	//     	- **0**: file system
	//
	// 	- **23**: DSC
	//
	//     	- **0**: instance
	//
	// 	- **24**: EIP
	//
	//     	- **0**: Anycast EIP
	//
	// 	- **25**: IDaaS EIAM
	//
	//     	- **0**: instance
	//
	// 	- **26**: PolarDB-X
	//
	//     	- **0**: instance
	//
	// 	- **27**: Elasticsearch
	//
	//     	- **0**: instance
	//
	// example:
	//
	// 0
	AssetSubType *string `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// The subtype name of the cloud asset.
	//
	// example:
	//
	// SECURITY_GROUP
	AssetSubTypeName *string `json:"AssetSubTypeName,omitempty" xml:"AssetSubTypeName,omitempty"`
	// The type of the cloud asset. Valid values:
	//
	// 	- **0**: Elastic Compute Service (ECS)
	//
	// 	- **1**: Server Load Balancer (SLB)
	//
	// 	- **3**: ApsaraDB RDS
	//
	// 	- **4**: ApsaraDB for MongoDB
	//
	// 	- **5**: ApsaraDB for Redis
	//
	// 	- **6**: Container Registry
	//
	// 	- **8**: Container Service for Kubernetes (ACK)
	//
	// 	- **9**: Virtual Private Cloud (VPC)
	//
	// 	- **11**: ActionTrail
	//
	// 	- **12**: Alibaba Cloud CDN
	//
	// 	- **13**: Certificate Management Service (formerly SSL Certificates Service)
	//
	// 	- **14**: Apsara Devops
	//
	// 	- **16**: Anti-DDoS
	//
	// 	- **17**: Web Application Firewall (WAF)
	//
	// 	- **18**: Object Storage Service (OSS)
	//
	// 	- **19**: PolarDB
	//
	// 	- **20**: ApsaraDB RDS for PostgreSQL
	//
	// 	- **21**: Microservices Engine (MSE)
	//
	// 	- **22**: File Storage NAS (NAS)
	//
	// 	- **23**: Data Security Center (DSC)
	//
	// 	- **24**: Elastic IP Address (EIP)
	//
	// 	- **25**: Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM)
	//
	// 	- **26**: PolarDB-X
	//
	// 	- **27**: Elasticsearch
	//
	// example:
	//
	// 0
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The type name of the cloud asset.
	//
	// example:
	//
	// ECS
	AssetTypeName *string `json:"AssetTypeName,omitempty" xml:"AssetTypeName,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 1607365213000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The instance ID of the cloud asset.
	//
	// example:
	//
	// d-uf60vevzkztnflx7cny5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name of the cloud asset.
	//
	// example:
	//
	// yztest-l***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the cloud asset.
	//
	// example:
	//
	// 1.2.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The ID of the region to which the cloud asset belongs.
	//
	// example:
	//
	// cn-hanghzou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Indicates whether risks are detected on the cloud asset. Valid values:
	//
	// 	- **YES**
	//
	// 	- **NO**
	//
	// example:
	//
	// NO
	RiskStatus *string `json:"RiskStatus,omitempty" xml:"RiskStatus,omitempty"`
	// The security information about the cloud asset.
	//
	// example:
	//
	// {"seriousNum":0,"appNum":0,"baselineMedium":0,"remindNum":0,"imageVulNntf":0,"cveNum":0,"vul":0,"uuid":"yuejia-test","emgNum":0,"weakPWNum":0,"imageMaliciousFileRemind":0,"imageBaselineMedium":0,"laterVulCount":0,"cmsNum":0,"imageMaliciousFileSerious":0,"agentlessMalicious":0,"suspNum":0,"imageBaselineHigh":0,"asapVulCount":0,"imageVulLater":0,"agentlessAll":0,"sysNum":0,"containerLater":0,"containerSuspicious":0,"imageBaselineNum":0,"newSuspicious":0,"nntfVulCount":0,"scaNum":0,"containerNntf":0,"health":0,"trojan":0,"suspicious":0,"imageMaliciousFileSuspicious":0,"containerRemind":0,"baselineLow":0,"imageVulAsap":0,"imageBaselineLow":0,"containerAsap":0,"agentlessBaseline":0,"agentlessVulSca":0,"agentlessVulCve":0,"containerSerious":0,"baselineHigh":0,"account":0,"baselineNum":5}
	SecurityInfo *string `json:"SecurityInfo,omitempty" xml:"SecurityInfo,omitempty"`
	// Tag list.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The service provider (SP) of the cloud asset. Valid values:
	//
	// 	- **0**: a cloud asset provided by Alibaba Cloud
	//
	// 	- **1**: a third-party cloud asset
	//
	// 	- **2**: a cloud asset in a data center
	//
	// 	- **3**, **4**, **5**, and **7**: other cloud asset
	//
	// 	- **8**: a lightweight cloud asset
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The account ID of the multi-cloud instance.
	//
	// example:
	//
	// 123xxx
	VendorUid *string `json:"VendorUid,omitempty" xml:"VendorUid,omitempty"`
	// The user name of the multi-cloud instance.
	//
	// example:
	//
	// testxxx
	VendorUserName *string `json:"VendorUserName,omitempty" xml:"VendorUserName,omitempty"`
}

func (s ListCloudAssetInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAssetInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListCloudAssetInstancesResponseBodyInstances) GetAlarmStatus() *string {
	return s.AlarmStatus
}

func (s *ListCloudAssetInstancesResponseBodyInstances) GetAssetSubType() *string {
	return s.AssetSubType
}

func (s *ListCloudAssetInstancesResponseBodyInstances) GetAssetSubTypeName() *string {
	return s.AssetSubTypeName
}

func (s *ListCloudAssetInstancesResponseBodyInstances) GetAssetType() *int32 {
	return s.AssetType
}

func (s *ListCloudAssetInstancesResponseBodyInstances) GetAssetTypeName() *string {
	return s.AssetTypeName
}

func (s *ListCloudAssetInstancesResponseBodyInstances) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListCloudAssetInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCloudAssetInstancesResponseBodyInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListCloudAssetInstancesResponseBodyInstances) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListCloudAssetInstancesResponseBodyInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCloudAssetInstancesResponseBodyInstances) GetRiskStatus() *string {
	return s.RiskStatus
}

func (s *ListCloudAssetInstancesResponseBodyInstances) GetSecurityInfo() *string {
	return s.SecurityInfo
}

func (s *ListCloudAssetInstancesResponseBodyInstances) GetTags() []*string {
	return s.Tags
}

func (s *ListCloudAssetInstancesResponseBodyInstances) GetVendor() *int32 {
	return s.Vendor
}

func (s *ListCloudAssetInstancesResponseBodyInstances) GetVendorUid() *string {
	return s.VendorUid
}

func (s *ListCloudAssetInstancesResponseBodyInstances) GetVendorUserName() *string {
	return s.VendorUserName
}

func (s *ListCloudAssetInstancesResponseBodyInstances) SetAlarmStatus(v string) *ListCloudAssetInstancesResponseBodyInstances {
	s.AlarmStatus = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyInstances) SetAssetSubType(v string) *ListCloudAssetInstancesResponseBodyInstances {
	s.AssetSubType = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyInstances) SetAssetSubTypeName(v string) *ListCloudAssetInstancesResponseBodyInstances {
	s.AssetSubTypeName = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyInstances) SetAssetType(v int32) *ListCloudAssetInstancesResponseBodyInstances {
	s.AssetType = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyInstances) SetAssetTypeName(v string) *ListCloudAssetInstancesResponseBodyInstances {
	s.AssetTypeName = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyInstances) SetCreatedTime(v int64) *ListCloudAssetInstancesResponseBodyInstances {
	s.CreatedTime = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyInstances) SetInstanceId(v string) *ListCloudAssetInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyInstances) SetInstanceName(v string) *ListCloudAssetInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyInstances) SetInternetIp(v string) *ListCloudAssetInstancesResponseBodyInstances {
	s.InternetIp = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyInstances) SetRegionId(v string) *ListCloudAssetInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyInstances) SetRiskStatus(v string) *ListCloudAssetInstancesResponseBodyInstances {
	s.RiskStatus = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyInstances) SetSecurityInfo(v string) *ListCloudAssetInstancesResponseBodyInstances {
	s.SecurityInfo = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyInstances) SetTags(v []*string) *ListCloudAssetInstancesResponseBodyInstances {
	s.Tags = v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyInstances) SetVendor(v int32) *ListCloudAssetInstancesResponseBodyInstances {
	s.Vendor = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyInstances) SetVendorUid(v string) *ListCloudAssetInstancesResponseBodyInstances {
	s.VendorUid = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyInstances) SetVendorUserName(v string) *ListCloudAssetInstancesResponseBodyInstances {
	s.VendorUserName = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}

type ListCloudAssetInstancesResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 2
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of cloud assets.
	//
	// example:
	//
	// 69
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudAssetInstancesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAssetInstancesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListCloudAssetInstancesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListCloudAssetInstancesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCloudAssetInstancesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudAssetInstancesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCloudAssetInstancesResponseBodyPageInfo) SetCount(v int32) *ListCloudAssetInstancesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyPageInfo) SetCurrentPage(v int32) *ListCloudAssetInstancesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyPageInfo) SetPageSize(v int32) *ListCloudAssetInstancesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyPageInfo) SetTotalCount(v int32) *ListCloudAssetInstancesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCloudAssetInstancesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
