// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudCenterInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCloudCenterInstancesResponseBody
	GetCode() *string
	SetData(v *DescribeCloudCenterInstancesResponseBodyData) *DescribeCloudCenterInstancesResponseBody
	GetData() *DescribeCloudCenterInstancesResponseBodyData
	SetMessage(v string) *DescribeCloudCenterInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCloudCenterInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCloudCenterInstancesResponseBody
	GetSuccess() *bool
}

type DescribeCloudCenterInstancesResponseBody struct {
	// example:
	//
	// 200
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeCloudCenterInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1B4C9A14-94E6-5EEB-BF39-7DACCE9AC0D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DescribeCloudCenterInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCloudCenterInstancesResponseBody) GetData() *DescribeCloudCenterInstancesResponseBodyData {
	return s.Data
}

func (s *DescribeCloudCenterInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCloudCenterInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudCenterInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCloudCenterInstancesResponseBody) SetCode(v string) *DescribeCloudCenterInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBody) SetData(v *DescribeCloudCenterInstancesResponseBodyData) *DescribeCloudCenterInstancesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBody) SetMessage(v string) *DescribeCloudCenterInstancesResponseBody {
	s.Message = &v
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudCenterInstancesResponseBodyData struct {
	Body *DescribeCloudCenterInstancesResponseBodyDataBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
}

func (s DescribeCloudCenterInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudCenterInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudCenterInstancesResponseBodyData) GetBody() *DescribeCloudCenterInstancesResponseBodyDataBody {
	return s.Body
}

func (s *DescribeCloudCenterInstancesResponseBodyData) SetBody(v *DescribeCloudCenterInstancesResponseBodyDataBody) *DescribeCloudCenterInstancesResponseBodyData {
	s.Body = v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyData) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudCenterInstancesResponseBodyDataBody struct {
	Instances []*DescribeCloudCenterInstancesResponseBodyDataBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	PageInfo  *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo    `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// example:
	//
	// 20EBDE7B-AA36-5D60-9DCA-151C48EDB9F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudCenterInstancesResponseBodyDataBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudCenterInstancesResponseBodyDataBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBody) GetInstances() []*DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	return s.Instances
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBody) GetPageInfo() *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBody) SetInstances(v []*DescribeCloudCenterInstancesResponseBodyDataBodyInstances) *DescribeCloudCenterInstancesResponseBodyDataBody {
	s.Instances = v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBody) SetPageInfo(v *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo) *DescribeCloudCenterInstancesResponseBodyDataBody {
	s.PageInfo = v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBody) SetRequestId(v string) *DescribeCloudCenterInstancesResponseBodyDataBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBody) SetSuccess(v bool) *DescribeCloudCenterInstancesResponseBodyDataBody {
	s.Success = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBody) Validate() error {
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

type DescribeCloudCenterInstancesResponseBodyDataBodyInstances struct {
	// example:
	//
	// NO
	AlarmStatus *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	// example:
	//
	// FC2U0JVHWS49S2OT
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// guokent
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// example:
	//
	// xxxxxx
	AssetTypeName *string `json:"AssetTypeName,omitempty" xml:"AssetTypeName,omitempty"`
	// example:
	//
	// 1627974044000
	AuthModifyTime *int64 `json:"AuthModifyTime,omitempty" xml:"AuthModifyTime,omitempty"`
	// example:
	//
	// 5
	AuthVersion *int32 `json:"AuthVersion,omitempty" xml:"AuthVersion,omitempty"`
	// example:
	//
	// 免费版
	AuthVersionName *string `json:"AuthVersionName,omitempty" xml:"AuthVersionName,omitempty"`
	// example:
	//
	// true
	Bind *bool `json:"Bind,omitempty" xml:"Bind,omitempty"`
	// example:
	//
	// none
	BindFileProtectType *string `json:"BindFileProtectType,omitempty" xml:"BindFileProtectType,omitempty"`
	// example:
	//
	// online
	ClientStatus *string `json:"ClientStatus,omitempty" xml:"ClientStatus,omitempty"`
	// example:
	//
	// online
	ClientSubStatus *string `json:"ClientSubStatus,omitempty" xml:"ClientSubStatus,omitempty"`
	// example:
	//
	// cb703cb0ba6bd40d4a6d8de5bff050fb9
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// auto-cn-heyuan
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// 9
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// example:
	//
	// Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz
	CpuInfo *string `json:"CpuInfo,omitempty" xml:"CpuInfo,omitempty"`
	// example:
	//
	// 1607365213000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 0
	ExposedStatus *int32 `json:"ExposedStatus,omitempty" xml:"ExposedStatus,omitempty"`
	// example:
	//
	// 0,1,2
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// example:
	//
	// ALIYUN
	FlagName *string `json:"FlagName,omitempty" xml:"FlagName,omitempty"`
	// example:
	//
	// 86d30f8b0e124aadb7ef3197f9dbd1f5
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// default
	GroupTrace *string `json:"GroupTrace,omitempty" xml:"GroupTrace,omitempty"`
	// example:
	//
	// YES
	HasContainer *string `json:"HasContainer,omitempty" xml:"HasContainer,omitempty"`
	// example:
	//
	// YES
	HcStatus *string `json:"HcStatus,omitempty" xml:"HcStatus,omitempty"`
	// example:
	//
	// 1
	HealthCheckCount *int32 `json:"HealthCheckCount,omitempty" xml:"HealthCheckCount,omitempty"`
	// example:
	//
	// 2
	Importance *int32 `json:"Importance,omitempty" xml:"Importance,omitempty"`
	// example:
	//
	// ls-cn-tl32rf**008
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ra-supabase-22u1iv3hr**5v9
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 47.1**.52.125
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// example:
	//
	// 172.16.1**.245
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// example:
	//
	// 114.55.*4.*6
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// 172.31.XX.XX,172.171.XX.XX
	IpListString *string `json:"IpListString,omitempty" xml:"IpListString,omitempty"`
	// example:
	//
	// 3.10.0-1127.19.1.el7.x86_64
	Kernel *string `json:"Kernel,omitempty" xml:"Kernel,omitempty"`
	// example:
	//
	// 1637592907000
	LastLoginTimestamp *int64 `json:"LastLoginTimestamp,omitempty" xml:"LastLoginTimestamp,omitempty"`
	// example:
	//
	// 00:13:3e:31:13:39,02:12:67:b8:**:**
	MacListString *string `json:"MacListString,omitempty" xml:"MacListString,omitempty"`
	// example:
	//
	// 1024
	Mem *string `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// example:
	//
	// slsshpcorlsmetrics
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// AliOS7U2-x86-64
	OsName *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	// example:
	//
	// 8
	PodCount *int32 `json:"PodCount,omitempty" xml:"PodCount,omitempty"`
	// example:
	//
	// 1
	PostPaidFlag *int32 `json:"PostPaidFlag,omitempty" xml:"PostPaidFlag,omitempty"`
	// example:
	//
	// cn-hangzhouxxxx
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// example:
	//
	// {
	//
	//       "account": 0,
	//
	//       "appNum": 0,
	//
	//       "asapVulCount": 0,
	//
	//       "baselineHigh": 0,
	//
	//       "baselineLow": 0,
	//
	//       "baselineMedium": 0,
	//
	//       "baselineNum": 0,
	//
	//       "cmsNum": 0,
	//
	//       "containerAsap": 0,
	//
	//       "containerLater": 0,
	//
	//       "containerNntf": 0,
	//
	//       "containerRemind": 0,
	//
	//       "containerSerious": 0,
	//
	//       "containerSuspicious": 0,
	//
	//       "cveNum": 0,
	//
	//       "emgNum": 0,
	//
	//       "health": 0,
	//
	//       "imageBaselineHigh": 0,
	//
	//       "imageBaselineLow": 0,
	//
	//       "imageBaselineMedium": 0,
	//
	//       "imageBaselineNum": 0,
	//
	//       "imageMaliciousFileRemind": 0,
	//
	//       "imageMaliciousFileSerious": 0,
	//
	//       "imageMaliciousFileSuspicious": 0,
	//
	//       "imageVulAsap": 0,
	//
	//       "imageVulLater": 0,
	//
	//       "imageVulNntf": 0,
	//
	//       "laterVulCount": 0,
	//
	//       "newSuspicious": 0,
	//
	//       "nntfVulCount": 0,
	//
	//       "remindNum": 0,
	//
	//       "scaNum": 0,
	//
	//       "seriousNum": 0,
	//
	//       "suspNum": 0,
	//
	//       "suspicious": 0,
	//
	//       "sysNum": 0,
	//
	//       "trojan": 0,
	//
	//       "uuid": "inet-37316411-37fe-4b72-b245-346a2721****",
	//
	//       "vul": 0,
	//
	//       "weakPWNum": 0
	//
	// }
	RiskCount *string `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// example:
	//
	// NO
	RiskStatus *string `json:"RiskStatus,omitempty" xml:"RiskStatus,omitempty"`
	// example:
	//
	// 5
	SafeEventCount *string `json:"SafeEventCount,omitempty" xml:"SafeEventCount,omitempty"`
	// example:
	//
	// dsw-76jlywunsif09bp15p
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// pre_20250714_idpt_adjust
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// d8586ab8be4549e3815995858d277763
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// example:
	//
	// app:test,type:lingjun
	TagResources *string `json:"TagResources,omitempty" xml:"TagResources,omitempty"`
	// example:
	//
	// 1f0459ee-ed49-6484-8958-4f10f61e6362
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// example:
	//
	// IDC
	VendorName *string `json:"VendorName,omitempty" xml:"VendorName,omitempty"`
	// example:
	//
	// 123
	VendorUid *string `json:"VendorUid,omitempty" xml:"VendorUid,omitempty"`
	// example:
	//
	// VendorUserName
	VendorUserName *string `json:"VendorUserName,omitempty" xml:"VendorUserName,omitempty"`
	// example:
	//
	// vpc-2zek7v0z4r6lbp02xckei
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
	// example:
	//
	// 2
	VulCount *int32 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
	// example:
	//
	// YES
	VulStatus *string `json:"VulStatus,omitempty" xml:"VulStatus,omitempty"`
}

func (s DescribeCloudCenterInstancesResponseBodyDataBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetAlarmStatus() *string {
	return s.AlarmStatus
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetAppId() *string {
	return s.AppId
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetAppName() *string {
	return s.AppName
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetAssetType() *string {
	return s.AssetType
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetAssetTypeName() *string {
	return s.AssetTypeName
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetAuthModifyTime() *int64 {
	return s.AuthModifyTime
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetAuthVersion() *int32 {
	return s.AuthVersion
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetAuthVersionName() *string {
	return s.AuthVersionName
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetBind() *bool {
	return s.Bind
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetBindFileProtectType() *string {
	return s.BindFileProtectType
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetClientStatus() *string {
	return s.ClientStatus
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetClientSubStatus() *string {
	return s.ClientSubStatus
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetCores() *int32 {
	return s.Cores
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetCpuInfo() *string {
	return s.CpuInfo
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetExposedStatus() *int32 {
	return s.ExposedStatus
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetFlag() *int32 {
	return s.Flag
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetFlagName() *string {
	return s.FlagName
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetGroupTrace() *string {
	return s.GroupTrace
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetHasContainer() *string {
	return s.HasContainer
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetHcStatus() *string {
	return s.HcStatus
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetHealthCheckCount() *int32 {
	return s.HealthCheckCount
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetImportance() *int32 {
	return s.Importance
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetIp() *string {
	return s.Ip
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetIpListString() *string {
	return s.IpListString
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetKernel() *string {
	return s.Kernel
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetLastLoginTimestamp() *int64 {
	return s.LastLoginTimestamp
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetMacListString() *string {
	return s.MacListString
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetMem() *string {
	return s.Mem
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetOs() *string {
	return s.Os
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetOsName() *string {
	return s.OsName
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetPodCount() *int32 {
	return s.PodCount
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetPostPaidFlag() *int32 {
	return s.PostPaidFlag
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetRegion() *string {
	return s.Region
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetRiskCount() *string {
	return s.RiskCount
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetRiskStatus() *string {
	return s.RiskStatus
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetSafeEventCount() *string {
	return s.SafeEventCount
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetTag() *string {
	return s.Tag
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetTagId() *string {
	return s.TagId
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetTagResources() *string {
	return s.TagResources
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetVendor() *int32 {
	return s.Vendor
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetVendorName() *string {
	return s.VendorName
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetVendorUid() *string {
	return s.VendorUid
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetVendorUserName() *string {
	return s.VendorUserName
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetVulCount() *int32 {
	return s.VulCount
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) GetVulStatus() *string {
	return s.VulStatus
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetAlarmStatus(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.AlarmStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetAppId(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.AppId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetAppName(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.AppName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetAssetType(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.AssetType = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetAssetTypeName(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.AssetTypeName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetAuthModifyTime(v int64) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.AuthModifyTime = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetAuthVersion(v int32) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.AuthVersion = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetAuthVersionName(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.AuthVersionName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetBind(v bool) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.Bind = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetBindFileProtectType(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.BindFileProtectType = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetClientStatus(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.ClientStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetClientSubStatus(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.ClientSubStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetClusterId(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.ClusterId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetClusterName(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.ClusterName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetCores(v int32) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.Cores = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetCpuInfo(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.CpuInfo = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetCreatedTime(v int64) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.CreatedTime = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetExposedStatus(v int32) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.ExposedStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetFlag(v int32) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.Flag = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetFlagName(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.FlagName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetGroupId(v int64) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.GroupId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetGroupTrace(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.GroupTrace = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetHasContainer(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.HasContainer = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetHcStatus(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.HcStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetHealthCheckCount(v int32) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.HealthCheckCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetImportance(v int32) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.Importance = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetInstanceId(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetInstanceName(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetInternetIp(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.InternetIp = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetIntranetIp(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.IntranetIp = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetIp(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.Ip = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetIpListString(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.IpListString = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetKernel(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.Kernel = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetLastLoginTimestamp(v int64) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.LastLoginTimestamp = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetMacListString(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.MacListString = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetMem(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.Mem = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetNamespace(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.Namespace = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetOs(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.Os = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetOsName(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.OsName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetPodCount(v int32) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.PodCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetPostPaidFlag(v int32) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.PostPaidFlag = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetRegion(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.Region = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetRegionId(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetRegionName(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.RegionName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetRiskCount(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.RiskCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetRiskStatus(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.RiskStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetSafeEventCount(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.SafeEventCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetServiceId(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.ServiceId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetStatus(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.Status = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetTag(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.Tag = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetTagId(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.TagId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetTagResources(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.TagResources = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetUuid(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.Uuid = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetVendor(v int32) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.Vendor = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetVendorName(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.VendorName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetVendorUid(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.VendorUid = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetVendorUserName(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.VendorUserName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetVpcInstanceId(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetVulCount(v int32) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.VulCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) SetVulStatus(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyInstances {
	s.VulStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo struct {
	// example:
	//
	// 7
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// m1NGAAAAAABzLzIwMjQwMg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo) SetCount(v int32) *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo) SetCurrentPage(v int32) *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo) SetNextToken(v string) *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo) SetPageSize(v int32) *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo) SetTotalCount(v int32) *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyDataBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
