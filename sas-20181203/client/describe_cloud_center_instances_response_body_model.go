// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudCenterInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*DescribeCloudCenterInstancesResponseBodyInstances) *DescribeCloudCenterInstancesResponseBody
	GetInstances() []*DescribeCloudCenterInstancesResponseBodyInstances
	SetPageInfo(v *DescribeCloudCenterInstancesResponseBodyPageInfo) *DescribeCloudCenterInstancesResponseBody
	GetPageInfo() *DescribeCloudCenterInstancesResponseBodyPageInfo
	SetRequestId(v string) *DescribeCloudCenterInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCloudCenterInstancesResponseBody
	GetSuccess() *bool
}

type DescribeCloudCenterInstancesResponseBody struct {
	// The details about the assets.
	Instances []*DescribeCloudCenterInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeCloudCenterInstancesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 32A73759-4C0F-4801-BE98-901223ACEE9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- **true**: The call is successful.
	//
	// 	- **false**: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudCenterInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudCenterInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudCenterInstancesResponseBody) GetInstances() []*DescribeCloudCenterInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeCloudCenterInstancesResponseBody) GetPageInfo() *DescribeCloudCenterInstancesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeCloudCenterInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudCenterInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCloudCenterInstancesResponseBody) SetInstances(v []*DescribeCloudCenterInstancesResponseBodyInstances) *DescribeCloudCenterInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBody) SetPageInfo(v *DescribeCloudCenterInstancesResponseBodyPageInfo) *DescribeCloudCenterInstancesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBody) SetRequestId(v string) *DescribeCloudCenterInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBody) SetSuccess(v bool) *DescribeCloudCenterInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBody) Validate() error {
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

type DescribeCloudCenterInstancesResponseBodyInstances struct {
	// Indicates whether alerts are generated on the asset. Valid values:
	//
	// 	- **YES**
	//
	// 	- **NO**
	//
	// example:
	//
	// NO
	AlarmStatus *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	// The ID of the application.
	//
	// >  This parameter is available only when the **Vendor*	- parameter is set to 9.
	//
	// example:
	//
	// test
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// >  This parameter is available only when the **Vendor*	- parameter is set to 9.
	//
	// example:
	//
	// testAppName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- **0**: an ECS instance
	//
	// 	- **1**: a Server Load Balancer (SLB) instance
	//
	// 	- **2**: a Network Address Translation (NAT) gateway
	//
	// 	- **3**: an ApsaraDB RDS instance
	//
	// 	- **4**: an ApsaraDB for MongoDB instance
	//
	// 	- **5**: an ApsaraDB for Redis instance
	//
	// 	- **6**: a container image
	//
	// 	- **7**: a container
	//
	// example:
	//
	// ecs
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The name of the asset type.
	//
	// example:
	//
	// Elastic Compute Service
	AssetTypeName *string `json:"AssetTypeName,omitempty" xml:"AssetTypeName,omitempty"`
	// The timestamp when Security Center is authorized to scan the asset.
	//
	// example:
	//
	// 1627974044000
	AuthModifyTime *int64 `json:"AuthModifyTime,omitempty" xml:"AuthModifyTime,omitempty"`
	// The edition of Security Center that is authorized to scan the asset. Valid values:
	//
	// 	- **1**: Basic edition
	//
	// 	- **6**: Anti-virus edition
	//
	// 	- **5**: Advanced edition
	//
	// 	- **3**: Enterprise edition
	//
	// 	- **7**: Ultimate edition
	//
	// 	- **10**: Value-added Plan edition
	//
	// example:
	//
	// 3
	AuthVersion *int32 `json:"AuthVersion,omitempty" xml:"AuthVersion,omitempty"`
	// The name of the Security Center edition that is authorized to protect the asset. Valid values:
	//
	// 	- Basic edition
	//
	// 	- Anti-virus edition
	//
	// 	- Advanced edition
	//
	// 	- Enterprise edition
	//
	// 	- Ultimate edition
	//
	// example:
	//
	// Ultimate Edition
	AuthVersionName *string `json:"AuthVersionName,omitempty" xml:"AuthVersionName,omitempty"`
	// Indicates whether Security Center is authorized to scan the asset. Valid values:
	//
	// 	- **true**: Security Center is authorized to scan the asset.
	//
	// 	- **false**: Security Center is not authorized to scan the asset.
	//
	// example:
	//
	// true
	Bind *bool `json:"Bind,omitempty" xml:"Bind,omitempty"`
	// Whether to bind tamper-proof authorization. Values:
	//
	// - **block**: Yes
	//
	// - **none**: No
	//
	// example:
	//
	// block
	BindFileProtectType *string `json:"BindFileProtectType,omitempty" xml:"BindFileProtectType,omitempty"`
	// The status of the Security Center agent installed on the asset. Valid values:
	//
	// 	- **online**: The Security Center agent is **enabled**.
	//
	// 	- **offline**: The Security Center agent is **disabled**.
	//
	// 	- **pause**: The Security Center agent is **suspended**.
	//
	// example:
	//
	// online
	ClientStatus *string `json:"ClientStatus,omitempty" xml:"ClientStatus,omitempty"`
	// The sub-status of the Security Center agent installed on the asset. Valid values:
	//
	// 	- **online**: The Security Center agent is **enabled**.
	//
	// 	- **offline**: The Security Center agent is **disabled**.
	//
	// 	- **pause**: The Security Center agent is **suspended**.
	//
	// 	- **uninstalled**: The Security Center agent is **uninstalled**.
	//
	// 	- **stopped**: The Security Center agent is **stopped**.
	//
	// example:
	//
	// online
	ClientSubStatus *string `json:"ClientSubStatus,omitempty" xml:"ClientSubStatus,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// c690a0789419f4284a4e0a29e12fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// cluster1
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The number of the CPU cores used by the asset.
	//
	// example:
	//
	// 4
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The CPU information about the asset.
	//
	// example:
	//
	// Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz
	CpuInfo *string `json:"CpuInfo,omitempty" xml:"CpuInfo,omitempty"`
	// The timestamp when the cluster was created. Unit: milliseconds.
	//
	// example:
	//
	// 1607365213000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// Indicates whether the asset is exposed. Valid values:
	//
	// 	- **0**: The asset is not exposed.
	//
	// 	- **1**: The asset is exposed.
	//
	// example:
	//
	// 0
	ExposedStatus *int32 `json:"ExposedStatus,omitempty" xml:"ExposedStatus,omitempty"`
	// Indicates whether the asset is an Alibaba Cloud asset. Valid values:
	//
	// 	- **0**: The asset is an Alibaba Cloud asset.
	//
	// 	- **1**: The asset is not an Alibaba Cloud asset.
	//
	// example:
	//
	// 0
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// Asset vendor. Values:
	//
	// - **ALIYUN*	-
	//
	// - **OUT**
	//
	// - **IDC*	-
	//
	// - **Tencent*	-
	//
	// - **HUAWEICLOUD*	-
	//
	// - **Azure*	-
	//
	// - **AWS*	-
	//
	// - **ASK*	-
	//
	// - **TRIPARTITE*	-
	//
	// - **SAE*	-
	//
	// - **PAI*	-
	//
	// - **google*	-
	//
	// - **VOLCENGINE**
	//
	// example:
	//
	// ASK
	FlagName *string `json:"FlagName,omitempty" xml:"FlagName,omitempty"`
	// The ID of the asset group to which the asset belongs.
	//
	// example:
	//
	// 4120080
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the group to which the asset belongs.
	//
	// example:
	//
	// default
	GroupTrace *string `json:"GroupTrace,omitempty" xml:"GroupTrace,omitempty"`
	// Indicates if containers are included. Valid values:
	//
	// 	- **YES**: yes.
	//
	// 	- **NO**: no.
	//
	// example:
	//
	// YES
	HasContainer *string `json:"HasContainer,omitempty" xml:"HasContainer,omitempty"`
	// Indicates whether baseline risks are detected on the asset. Valid values:
	//
	// 	- **YES**
	//
	// 	- **NO**
	//
	// example:
	//
	// YES
	HcStatus *string `json:"HcStatus,omitempty" xml:"HcStatus,omitempty"`
	// The number of baseline risks that are detected on the asset.
	//
	// example:
	//
	// 1
	HealthCheckCount *int32 `json:"HealthCheckCount,omitempty" xml:"HealthCheckCount,omitempty"`
	// The importance of the asset. Valid values:
	//
	// 	- **2**: an important asset
	//
	// 	- **1**: a common asset
	//
	// 	- **0**: a test asset
	//
	// example:
	//
	// 2
	Importance *int32 `json:"Importance,omitempty" xml:"Importance,omitempty"`
	// The ID of the asset.
	//
	// example:
	//
	// i-m5***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// yztest-l***
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
	// The public IP address of the asset.
	//
	// example:
	//
	// 1.2.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The IP addresses of the system.
	//
	// example:
	//
	// 172.31.XX.XX,172.171.XX.XX
	IpListString *string `json:"IpListString,omitempty" xml:"IpListString,omitempty"`
	// The version of the kernel.
	//
	// example:
	//
	// 3.10.0-1127.19.1.el7.x86_64
	Kernel *string `json:"Kernel,omitempty" xml:"Kernel,omitempty"`
	// The timestamp when the Security Center agent was last online. Unit: milliseconds.
	//
	// example:
	//
	// 1637592907000
	LastLoginTimestamp *int64 `json:"LastLoginTimestamp,omitempty" xml:"LastLoginTimestamp,omitempty"`
	// The MAC addresses of the system.
	//
	// example:
	//
	// 00:13:3e:31:13:39,02:12:67:b8:**:**
	MacListString *string `json:"MacListString,omitempty" xml:"MacListString,omitempty"`
	// The size of the memory. Unit: MB.
	//
	// example:
	//
	// 1024
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The namespace.
	//
	// example:
	//
	// crm-test
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The operating system of the asset.
	//
	// example:
	//
	// Linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The kernel version of the asset.
	//
	// example:
	//
	// -
	OsName *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	// The number of pods.
	//
	// example:
	//
	// 1
	PodCount *int32 `json:"PodCount,omitempty" xml:"PodCount,omitempty"`
	// The billing method of the protection version currently bound to the asset. Values: - **0**: Subscription - **1**: Pay-as-you-go
	//
	// example:
	//
	// 0
	PostPaidFlag *int32 `json:"PostPaidFlag,omitempty" xml:"PostPaidFlag,omitempty"`
	// The region ID of the asset.
	//
	// example:
	//
	// cn-hangzhou-cm***-***
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the region in which the asset resides.
	//
	// example:
	//
	// cn-hanghzou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region in which the asset resides.
	//
	// example:
	//
	// China (Hangzhou)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The total number of baseline risks that are detected on the asset. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **account**: the number of accounts that are used to log on from unapproved logon locations and whose passwords are cracked
	//
	// 	- **appNum**: the number of scanners
	//
	// 	- **asapVulCount**: the total number of high-severity vulnerabilities
	//
	// 	- **baselineHigh**: the number of high-risk baseline risks
	//
	// 	- **baselineLow**: the number of low-risk baseline risks
	//
	// 	- **baselineMedium**: the number of medium-risk baseline risks
	//
	// 	- **baselineNum**: the total number of baseline risks
	//
	// 	- **cmsNum**: the number of Web-CMS vulnerabilities
	//
	// 	- **containerAsap**: the number of high-severity vulnerabilities that are detected on containers
	//
	// 	- **containerLater**: the number of medium-severity vulnerabilities that are detected on containers
	//
	// 	- **containerNntf**: the number of low-severity vulnerabilities that are detected on containers
	//
	// 	- **containerRemind**: the number of alerts whose Emergency level is Reminder on containers
	//
	// 	- **containerSerious**: the number of alerts Emergency level is Urgent on containers
	//
	// 	- **containerSuspicious**: the number of alerts whose Emergency level is Suspicious on containers
	//
	// 	- **cveNum**: the number of Linux software vulnerabilities
	//
	// 	- **emgNum**: the number of urgent vulnerabilities
	//
	// 	- **health**: the number of baseline alerts that are unhandled
	//
	// 	- **imageBaselineHigh**: the number of high-risk baseline risks that are detected on images
	//
	// 	- **imageBaselineLow**: the number of low-risk baseline risks that are detected on images
	//
	// 	- **imageBaselineMedium**: the number of medium-risk baseline risks that are detected on images
	//
	// 	- **imageBaselineNum**: the total number of baseline risks that are detected on images
	//
	// 	- **imageMaliciousFileRemind**: the number of malicious files that are detected on images and have the Emergency level of Reminder
	//
	// 	- **imageMaliciousFileSerious**: the number of malicious files that are detected on images and have the Emergency level of Urgent
	//
	// 	- **imageMaliciousFileSuspicious**: the number of malicious files that are detected on images and have the Emergency level of Suspicious
	//
	// 	- **imageVulAsap**: the number of high-severity vulnerabilities that are detected on images
	//
	// 	- **imageVulLater**: the number of medium-severity vulnerabilities that are detected on an image
	//
	// 	- **imageVulNntf**: the number of low-severity vulnerabilities that are detected on an image
	//
	// 	- **laterVulCount**: the number of medium-severity vulnerabilities
	//
	// 	- **newSuspicious**: the number of alerts
	//
	// 	- **nntfVulCount**: the number of low-severity vulnerabilities.
	//
	// 	- **remindNum**: the number of alerts whose Emergency level is Reminder
	//
	// 	- **scaNum**: the number of vulnerabilities that are detected based on software component analysis
	//
	// 	- **seriousNum**: the number of alerts whose Emergency level is Urgent
	//
	// 	- **suspNum**: the number of alerts whose Emergency level is Suspicious
	//
	// 	- **suspicious**: the total number of alerts
	//
	// 	- **sysNum**: the number of Windows system vulnerabilities
	//
	// 	- **trojan**: the number of trojans
	//
	// 	- **uuid**: the UUIDs of assets
	//
	// 	- **vul**: the number of vulnerabilities
	//
	// 	- **weakPWNum**: the number of weak passwords
	//
	// example:
	//
	// {"account":0,"appNum":0,"asapVulCount":0,"baselineHigh":0,"baselineLow":0,"baselineMedium":0,"baselineNum":0,"cmsNum":0,"containerAsap":0,"containerLater":0,"containerNntf":0,"containerRemind":0,"containerSerious":0,"containerSuspicious":0,"cveNum":0,"emgNum":0,"health":0,"imageBaselineHigh":0,"imageBaselineLow":0,"imageBaselineMedium":0,"imageBaselineNum":0,"imageMaliciousFileRemind":0,"imageMaliciousFileSerious":0,"imageMaliciousFileSuspicious":0,"imageVulAsap":0,"imageVulLater":0,"imageVulNntf":0,"laterVulCount":0,"newSuspicious":0,"nntfVulCount":0,"remindNum":0,"scaNum":0,"seriousNum":0,"suspNum":0,"suspicious":0,"sysNum":0,"trojan":0,"uuid":"inet-37316411-37fe-4b72-b245-346a2721d4b6","vul":0,"weakPWNum":0}
	RiskCount *string `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// Indicates whether risks are detected on the asset. Valid values:
	//
	// 	- **YES**
	//
	// 	- **NO**
	//
	// example:
	//
	// NO
	RiskStatus *string `json:"RiskStatus,omitempty" xml:"RiskStatus,omitempty"`
	// The number of alerts that are generated on the asset.
	//
	// example:
	//
	// 5
	SafeEventCount *int32 `json:"SafeEventCount,omitempty" xml:"SafeEventCount,omitempty"`
	// Service ID. Only available for PAI instances.
	//
	// example:
	//
	// dsw-76jlywunsif09bp15p
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The status of the asset. Valid values:
	//
	// 	- **Running**: running
	//
	// 	- **notRunning**: stopped
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the asset tag.
	//
	// example:
	//
	// InternetIp,test
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The ID of the asset tag.
	//
	// example:
	//
	// 121313,41412
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The custom tag added to the Lingjun node. This parameter is returned only for LINGJUN GPU-accelerated instances.
	//
	// example:
	//
	// app:test,type:lingjun
	TagResources *string `json:"TagResources,omitempty" xml:"TagResources,omitempty"`
	// The UUID of the asset.
	//
	// example:
	//
	// c9107c04-942f-40c1-981a-f1c1***
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// Asset vendor. Values:
	//
	// - **0**: an asset provided by Alibaba Cloud
	//
	// - **1**: an asset outside Alibaba Cloud
	//
	// - **2**: an asset in a data center
	//
	// - **3**, **4**, **5**, **7**, **14**, **16**: an asset from a third-party cloud service provider
	//
	// - **8**: a lightweight asset
	//
	// - **9**: a Serverless App Engine (SAE) instance
	//
	// - **10**: an instance in Platform for AI (PAI)
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The name of the service provider for the asset.
	//
	// Values:
	//
	//  - **ALIYUN**: Alibaba Cloud
	//
	// - **OUT**: a third-party service provider
	//
	// - **IDC**: an asset in a data center
	//
	// - **TENCENT**: Tencent Cloud
	//
	// - **HUAWEICLOUD**: Huawei Cloud
	//
	// - **Microsoft**: Microsoft Azure
	//
	// - **AWS**: Amazon Web Services (AWS)
	//
	// - **TRIPARTITE**: a lightweight server
	//
	// - **SAE**: a Serverless App Engine (SAE) instance
	//
	// - **PAI**: an instance in Platform for AI (PAI)
	//
	// - **VOLCENGINE**: VOLCENGINE Cloud
	//
	// - **google**: GOOGLE Cloud
	//
	// example:
	//
	// Tencent
	VendorName *string `json:"VendorName,omitempty" xml:"VendorName,omitempty"`
	// Account ID of the multi-cloud instance.
	//
	// example:
	//
	// 123
	VendorUid *string `json:"VendorUid,omitempty" xml:"VendorUid,omitempty"`
	// Account name of the multi-cloud instance.
	//
	// example:
	//
	// VendorUserName
	VendorUserName *string `json:"VendorUserName,omitempty" xml:"VendorUserName,omitempty"`
	// The ID of the VPC to which the asset belongs.
	//
	// example:
	//
	// vpc-uf60agqq65bs98zoo****
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
	// The number of vulnerabilities that are detected on the asset.
	//
	// example:
	//
	// 2
	VulCount *int32 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
	// Indicates whether vulnerabilities are detected on the asset. Valid values:
	//
	// 	- **YES**
	//
	// 	- **NO**
	//
	// example:
	//
	// YES
	VulStatus *string `json:"VulStatus,omitempty" xml:"VulStatus,omitempty"`
}

func (s DescribeCloudCenterInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudCenterInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetAlarmStatus() *string {
	return s.AlarmStatus
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetAppId() *string {
	return s.AppId
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetAppName() *string {
	return s.AppName
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetAssetType() *string {
	return s.AssetType
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetAssetTypeName() *string {
	return s.AssetTypeName
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetAuthModifyTime() *int64 {
	return s.AuthModifyTime
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetAuthVersion() *int32 {
	return s.AuthVersion
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetAuthVersionName() *string {
	return s.AuthVersionName
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetBind() *bool {
	return s.Bind
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetBindFileProtectType() *string {
	return s.BindFileProtectType
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetClientStatus() *string {
	return s.ClientStatus
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetClientSubStatus() *string {
	return s.ClientSubStatus
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetCores() *int32 {
	return s.Cores
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetCpuInfo() *string {
	return s.CpuInfo
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetExposedStatus() *int32 {
	return s.ExposedStatus
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetFlag() *int32 {
	return s.Flag
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetFlagName() *string {
	return s.FlagName
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetGroupTrace() *string {
	return s.GroupTrace
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetHasContainer() *string {
	return s.HasContainer
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetHcStatus() *string {
	return s.HcStatus
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetHealthCheckCount() *int32 {
	return s.HealthCheckCount
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetImportance() *int32 {
	return s.Importance
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetIp() *string {
	return s.Ip
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetIpListString() *string {
	return s.IpListString
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetKernel() *string {
	return s.Kernel
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetLastLoginTimestamp() *int64 {
	return s.LastLoginTimestamp
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetMacListString() *string {
	return s.MacListString
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetMem() *int32 {
	return s.Mem
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetOs() *string {
	return s.Os
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetOsName() *string {
	return s.OsName
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetPodCount() *int32 {
	return s.PodCount
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetPostPaidFlag() *int32 {
	return s.PostPaidFlag
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetRegion() *string {
	return s.Region
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetRiskCount() *string {
	return s.RiskCount
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetRiskStatus() *string {
	return s.RiskStatus
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetSafeEventCount() *int32 {
	return s.SafeEventCount
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetTag() *string {
	return s.Tag
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetTagId() *string {
	return s.TagId
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetTagResources() *string {
	return s.TagResources
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetVendor() *int32 {
	return s.Vendor
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetVendorName() *string {
	return s.VendorName
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetVendorUid() *string {
	return s.VendorUid
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetVendorUserName() *string {
	return s.VendorUserName
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetVulCount() *int32 {
	return s.VulCount
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) GetVulStatus() *string {
	return s.VulStatus
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetAlarmStatus(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.AlarmStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetAppId(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.AppId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetAppName(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.AppName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetAssetType(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.AssetType = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetAssetTypeName(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.AssetTypeName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetAuthModifyTime(v int64) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.AuthModifyTime = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetAuthVersion(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.AuthVersion = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetAuthVersionName(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.AuthVersionName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetBind(v bool) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Bind = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetBindFileProtectType(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.BindFileProtectType = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetClientStatus(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.ClientStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetClientSubStatus(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.ClientSubStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetClusterId(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.ClusterId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetClusterName(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.ClusterName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetCores(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Cores = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetCpuInfo(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.CpuInfo = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetCreatedTime(v int64) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.CreatedTime = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetExposedStatus(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.ExposedStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetFlag(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Flag = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetFlagName(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.FlagName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetGroupId(v int64) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.GroupId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetGroupTrace(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.GroupTrace = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetHasContainer(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.HasContainer = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetHcStatus(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.HcStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetHealthCheckCount(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.HealthCheckCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetImportance(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Importance = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetInstanceName(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetInternetIp(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.InternetIp = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetIntranetIp(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.IntranetIp = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetIp(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Ip = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetIpListString(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.IpListString = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetKernel(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Kernel = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetLastLoginTimestamp(v int64) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.LastLoginTimestamp = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetMacListString(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.MacListString = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetMem(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Mem = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetNamespace(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Namespace = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetOs(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Os = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetOsName(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.OsName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetPodCount(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.PodCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetPostPaidFlag(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.PostPaidFlag = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetRegion(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Region = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetRegionId(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetRegionName(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.RegionName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetRiskCount(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.RiskCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetRiskStatus(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.RiskStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetSafeEventCount(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.SafeEventCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetServiceId(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.ServiceId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetStatus(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetTag(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Tag = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetTagId(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.TagId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetTagResources(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.TagResources = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetUuid(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Uuid = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetVendor(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Vendor = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetVendorName(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.VendorName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetVendorUid(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.VendorUid = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetVendorUserName(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.VendorUserName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetVpcInstanceId(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetVulCount(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.VulCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetVulStatus(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.VulStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudCenterInstancesResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The value of NextToken that is returned when the NextToken method is used.
	//
	// example:
	//
	// B604532DEF982B875E8360A6EFA3B***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries returned per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudCenterInstancesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudCenterInstancesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCloudCenterInstancesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeCloudCenterInstancesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCloudCenterInstancesResponseBodyPageInfo) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCloudCenterInstancesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudCenterInstancesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCloudCenterInstancesResponseBodyPageInfo) SetCount(v int32) *DescribeCloudCenterInstancesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeCloudCenterInstancesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyPageInfo) SetNextToken(v string) *DescribeCloudCenterInstancesResponseBodyPageInfo {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyPageInfo) SetPageSize(v int32) *DescribeCloudCenterInstancesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeCloudCenterInstancesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
