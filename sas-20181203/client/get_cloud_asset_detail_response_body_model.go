// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAssetDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *GetCloudAssetDetailResponseBody
	GetCount() *int32
	SetInstances(v []*GetCloudAssetDetailResponseBodyInstances) *GetCloudAssetDetailResponseBody
	GetInstances() []*GetCloudAssetDetailResponseBodyInstances
	SetRequestId(v string) *GetCloudAssetDetailResponseBody
	GetRequestId() *string
}

type GetCloudAssetDetailResponseBody struct {
	// The number of instances in the list of cloud assets returned.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// An array that consists of the details of the cloud assets.
	Instances []*GetCloudAssetDetailResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// CB45CAED-31C3-517A-8619-10F632D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCloudAssetDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudAssetDetailResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *GetCloudAssetDetailResponseBody) GetInstances() []*GetCloudAssetDetailResponseBodyInstances {
	return s.Instances
}

func (s *GetCloudAssetDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCloudAssetDetailResponseBody) SetCount(v int32) *GetCloudAssetDetailResponseBody {
	s.Count = &v
	return s
}

func (s *GetCloudAssetDetailResponseBody) SetInstances(v []*GetCloudAssetDetailResponseBodyInstances) *GetCloudAssetDetailResponseBody {
	s.Instances = v
	return s
}

func (s *GetCloudAssetDetailResponseBody) SetRequestId(v string) *GetCloudAssetDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudAssetDetailResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCloudAssetDetailResponseBodyInstances struct {
	// Indicates whether alerts are generated for the current cloud asset. Valid values:
	//
	// 	- **YES**
	//
	// 	- **NO**
	//
	// example:
	//
	// NO
	AlarmStatus *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	// The subtype of the cloud asset.
	//
	// example:
	//
	// 0
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// The name of the cloud asset subtype.
	//
	// example:
	//
	// INSTANCE
	AssetSubTypeName *string `json:"AssetSubTypeName,omitempty" xml:"AssetSubTypeName,omitempty"`
	// The type of the cloud asset. Valid values:
	//
	// 	- **0**: ECS.
	//
	// 	- **1**: SLB.
	//
	// 	- **3**: ApsaraDB RDS.
	//
	// 	- **4**: ApsaraDB for MongoDB.
	//
	// 	- **5**: ApsaraDB for Redis.
	//
	// 	- **6**: Container Registry.
	//
	// 	- **8**: Container Service for Kubernetes.
	//
	// 	- **9**: VPC.
	//
	// 	- **11**: ActionTrail.
	//
	// 	- **12**: CDN.
	//
	// 	- **13**: Certificate Management Service.
	//
	// 	- **14**: Apsara Devops.
	//
	// 	- **15**: RAM.
	//
	// 	- **16**: Anti-DDoS.
	//
	// 	- **17**: WAF.
	//
	// 	- **18**: OSS.
	//
	// 	- **19**: PolarDB.
	//
	// 	- **20**: ApsaraDB RDS for PostgreSQL.
	//
	// 	- **21**: MSE.
	//
	// 	- **22**: NAS.
	//
	// 	- **23**: DSC.
	//
	// 	- **24**: EIP.
	//
	// 	- **25**: IDaaS-EIAM.
	//
	// 	- **26**: PolarDB-X.
	//
	// 	- **27**: Elasticsearch.
	//
	// example:
	//
	// 3
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The name of the cloud asset type.
	//
	// example:
	//
	// RDS
	AssetTypeName *string `json:"AssetTypeName,omitempty" xml:"AssetTypeName,omitempty"`
	// The time when the instance was created. The value is a timestamp.
	//
	// example:
	//
	// 1607365213000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The detailed address of the cloud asset.
	DetailLink *string `json:"DetailLink,omitempty" xml:"DetailLink,omitempty"`
	// The instance ID of the cloud asset.
	//
	// example:
	//
	// rm-uf6t6u05n6g48****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name of the cloud asset.
	//
	// example:
	//
	// yztest-l***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the instance.
	//
	// example:
	//
	// 1.2.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Detailed asset information.
	//
	// example:
	//
	// {"owner":{"displayName":"123","id":"123"},"intranetEndpoint":"oss-cn-shanghai-internal.aliyuncs.com","extranetEndpoint":"oss-cn-shanghai.aliyuncs.com","storageClass":"Standard","name":"test","location":"oss-cn-shanghai","creationDate":1629882579000,"region":"cn-shanghai"}
	OriginalAssetInfo *string `json:"OriginalAssetInfo,omitempty" xml:"OriginalAssetInfo,omitempty"`
	// The region in which the cloud asset resides.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hanghzou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Indicates whether risks are detected on the current cloud asset. Valid values:
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
	// {\\"seriousNum\\":0,\\"appNum\\":0,\\"baselineMedium\\":0,\\"remindNum\\":0,\\"imageVulNntf\\":0,\\"cveNum\\":0,\\"vul\\":0,\\"uuid\\":\\"rm-uf6t6u05n6g485o70\\",\\"emgNum\\":0,\\"weakPWNum\\":0,\\"imageMaliciousFileRemind\\":0,\\"imageBaselineMedium\\":0,\\"laterVulCount\\":0,\\"cmsNum\\":0,\\"imageMaliciousFileSerious\\":0,\\"agentlessMalicious\\":0,\\"suspNum\\":0,\\"imageBaselineHigh\\":0,\\"asapVulCount\\":0,\\"imageVulLater\\":0,\\"agentlessAll\\":0,\\"sysNum\\":0,\\"containerLater\\":0,\\"containerSuspicious\\":0,\\"imageBaselineNum\\":0,\\"newSuspicious\\":0,\\"nntfVulCount\\":0,\\"scaNum\\":0,\\"containerNntf\\":0,\\"health\\":0,\\"trojan\\":0,\\"suspicious\\":0,\\"imageMaliciousFileSuspicious\\":0,\\"containerRemind\\":0,\\"baselineLow\\":0,\\"imageVulAsap\\":0,\\"imageBaselineLow\\":0,\\"containerAsap\\":0,\\"agentlessBaseline\\":0,\\"agentlessVulSca\\":0,\\"agentlessVulCve\\":0,\\"containerSerious\\":0,\\"baselineHigh\\":0,\\"account\\":0,\\"baselineNum\\":6}
	SecurityInfo *string `json:"SecurityInfo,omitempty" xml:"SecurityInfo,omitempty"`
	// The service provider of the cloud asset. Valid values:
	//
	// 	- **0**: Alibaba Cloud.
	//
	// 	- **1**: service provider that is unrecognized.
	//
	// 	- **2**: data center.
	//
	// 	- **3**, **4**, **5**, and **7**: third-party service provider.
	//
	// 	- **8**: simple application server.
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// Account id for multi-cloud instances.
	//
	// example:
	//
	// 123
	VendorUid *string `json:"VendorUid,omitempty" xml:"VendorUid,omitempty"`
	// The account name of the multi-cloud instance.
	//
	// example:
	//
	// test
	VendorUserName *string `json:"VendorUserName,omitempty" xml:"VendorUserName,omitempty"`
}

func (s GetCloudAssetDetailResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetDetailResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *GetCloudAssetDetailResponseBodyInstances) GetAlarmStatus() *string {
	return s.AlarmStatus
}

func (s *GetCloudAssetDetailResponseBodyInstances) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *GetCloudAssetDetailResponseBodyInstances) GetAssetSubTypeName() *string {
	return s.AssetSubTypeName
}

func (s *GetCloudAssetDetailResponseBodyInstances) GetAssetType() *int32 {
	return s.AssetType
}

func (s *GetCloudAssetDetailResponseBodyInstances) GetAssetTypeName() *string {
	return s.AssetTypeName
}

func (s *GetCloudAssetDetailResponseBodyInstances) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetCloudAssetDetailResponseBodyInstances) GetDetailLink() *string {
	return s.DetailLink
}

func (s *GetCloudAssetDetailResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCloudAssetDetailResponseBodyInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetCloudAssetDetailResponseBodyInstances) GetInternetIp() *string {
	return s.InternetIp
}

func (s *GetCloudAssetDetailResponseBodyInstances) GetOriginalAssetInfo() *string {
	return s.OriginalAssetInfo
}

func (s *GetCloudAssetDetailResponseBodyInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCloudAssetDetailResponseBodyInstances) GetRiskStatus() *string {
	return s.RiskStatus
}

func (s *GetCloudAssetDetailResponseBodyInstances) GetSecurityInfo() *string {
	return s.SecurityInfo
}

func (s *GetCloudAssetDetailResponseBodyInstances) GetVendor() *int32 {
	return s.Vendor
}

func (s *GetCloudAssetDetailResponseBodyInstances) GetVendorUid() *string {
	return s.VendorUid
}

func (s *GetCloudAssetDetailResponseBodyInstances) GetVendorUserName() *string {
	return s.VendorUserName
}

func (s *GetCloudAssetDetailResponseBodyInstances) SetAlarmStatus(v string) *GetCloudAssetDetailResponseBodyInstances {
	s.AlarmStatus = &v
	return s
}

func (s *GetCloudAssetDetailResponseBodyInstances) SetAssetSubType(v int32) *GetCloudAssetDetailResponseBodyInstances {
	s.AssetSubType = &v
	return s
}

func (s *GetCloudAssetDetailResponseBodyInstances) SetAssetSubTypeName(v string) *GetCloudAssetDetailResponseBodyInstances {
	s.AssetSubTypeName = &v
	return s
}

func (s *GetCloudAssetDetailResponseBodyInstances) SetAssetType(v int32) *GetCloudAssetDetailResponseBodyInstances {
	s.AssetType = &v
	return s
}

func (s *GetCloudAssetDetailResponseBodyInstances) SetAssetTypeName(v string) *GetCloudAssetDetailResponseBodyInstances {
	s.AssetTypeName = &v
	return s
}

func (s *GetCloudAssetDetailResponseBodyInstances) SetCreatedTime(v int64) *GetCloudAssetDetailResponseBodyInstances {
	s.CreatedTime = &v
	return s
}

func (s *GetCloudAssetDetailResponseBodyInstances) SetDetailLink(v string) *GetCloudAssetDetailResponseBodyInstances {
	s.DetailLink = &v
	return s
}

func (s *GetCloudAssetDetailResponseBodyInstances) SetInstanceId(v string) *GetCloudAssetDetailResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *GetCloudAssetDetailResponseBodyInstances) SetInstanceName(v string) *GetCloudAssetDetailResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *GetCloudAssetDetailResponseBodyInstances) SetInternetIp(v string) *GetCloudAssetDetailResponseBodyInstances {
	s.InternetIp = &v
	return s
}

func (s *GetCloudAssetDetailResponseBodyInstances) SetOriginalAssetInfo(v string) *GetCloudAssetDetailResponseBodyInstances {
	s.OriginalAssetInfo = &v
	return s
}

func (s *GetCloudAssetDetailResponseBodyInstances) SetRegionId(v string) *GetCloudAssetDetailResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *GetCloudAssetDetailResponseBodyInstances) SetRiskStatus(v string) *GetCloudAssetDetailResponseBodyInstances {
	s.RiskStatus = &v
	return s
}

func (s *GetCloudAssetDetailResponseBodyInstances) SetSecurityInfo(v string) *GetCloudAssetDetailResponseBodyInstances {
	s.SecurityInfo = &v
	return s
}

func (s *GetCloudAssetDetailResponseBodyInstances) SetVendor(v int32) *GetCloudAssetDetailResponseBodyInstances {
	s.Vendor = &v
	return s
}

func (s *GetCloudAssetDetailResponseBodyInstances) SetVendorUid(v string) *GetCloudAssetDetailResponseBodyInstances {
	s.VendorUid = &v
	return s
}

func (s *GetCloudAssetDetailResponseBodyInstances) SetVendorUserName(v string) *GetCloudAssetDetailResponseBodyInstances {
	s.VendorUserName = &v
	return s
}

func (s *GetCloudAssetDetailResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}
