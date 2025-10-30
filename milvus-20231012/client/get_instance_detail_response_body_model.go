// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetInstanceDetailResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *GetInstanceDetailResponseBodyData) *GetInstanceDetailResponseBody
	GetData() *GetInstanceDetailResponseBodyData
	SetErrCode(v string) *GetInstanceDetailResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetInstanceDetailResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *GetInstanceDetailResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetInstanceDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceDetailResponseBody
	GetSuccess() *bool
}

type GetInstanceDetailResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxxx",     "EncodedDiagnosticMessage": "xxxx",     "AuthPrincipalType": "SubUser",     "AuthPrincipalDisplayName": "xxxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "milvus:xxxx" }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned result.
	Data *GetInstanceDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned.
	//
	// example:
	//
	// Instance.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Failed to find instance c-123xxx
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABCD-1234-5678-EFGH
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceDetailResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetInstanceDetailResponseBody) GetData() *GetInstanceDetailResponseBodyData {
	return s.Data
}

func (s *GetInstanceDetailResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetInstanceDetailResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetInstanceDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceDetailResponseBody) SetAccessDeniedDetail(v string) *GetInstanceDetailResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetData(v *GetInstanceDetailResponseBodyData) *GetInstanceDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceDetailResponseBody) SetErrCode(v string) *GetInstanceDetailResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetErrMessage(v string) *GetInstanceDetailResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetHttpStatusCode(v int32) *GetInstanceDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetRequestId(v string) *GetInstanceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetSuccess(v bool) *GetInstanceDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceDetailResponseBodyData struct {
	// AclId for Public Network Access Control.
	//
	// example:
	//
	// acl-123xxx
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1716863508000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The name of the bucket.
	//
	// example:
	//
	// milvus-test
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The address of the bucket.
	//
	// example:
	//
	// https://oss.console.aliyun.com/bucket/oss-cn-beijing/xxxx/object?spm=a2cug.25127996.0.0.577990370Ebsqt&path=milvus-c-123xxxx
	BucketPath *string `json:"BucketPath,omitempty" xml:"BucketPath,omitempty"`
	// The instance details.
	ClusterInfo *GetInstanceDetailResponseBodyDataClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Struct"`
	// The instance name.
	//
	// example:
	//
	// aliyun-test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// true
	EnableHa  *bool   `json:"EnableHa,omitempty" xml:"EnableHa,omitempty"`
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 1721664000000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance status. Valid values:
	//
	// 	- creating.
	//
	// 	- running.
	//
	// 	- updating. Cluster scaling (up/down), configuration changes, and enabling/disabling public network access.
	//
	// 	- disable. The cluster has expired and needs to be renewed for activation.
	//
	// 	- deleting.
	//
	// 	- deleted.
	//
	// example:
	//
	// running
	InstanceStatus *string                                         `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	KmsKeyId       *string                                         `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	MeasureConfig  *GetInstanceDetailResponseBodyDataMeasureConfig `json:"MeasureConfig,omitempty" xml:"MeasureConfig,omitempty" type:"Struct"`
	MultiZoneMode  *string                                         `json:"MultiZoneMode,omitempty" xml:"MultiZoneMode,omitempty"`
	// example:
	//
	// perf
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// Indicates whether Internet access is enabled.
	//
	// example:
	//
	// true
	OpenPublicNet *bool `json:"OpenPublicNet,omitempty" xml:"OpenPublicNet,omitempty"`
	// The specification details. Valid values:
	//
	// 	- trial.
	//
	// 	- standard.
	//
	// example:
	//
	// standard
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- 0: pay-as-you-go
	//
	// 	- 1: subscription
	//
	// example:
	//
	// 1
	PayType *int32 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The commodity code.
	//
	// example:
	//
	// milvus_milvuspre_public_cn
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-123xxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The runtime.
	//
	// example:
	//
	// 1743679
	RunningTime *int64 `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-123xxx
	SgId *string                                  `json:"SgId,omitempty" xml:"SgId,omitempty"`
	Tags []*GetInstanceDetailResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The version of the software stack.
	//
	// example:
	//
	// 2.4.1-1.0-0.0.1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// User-defined configuration.
	//
	// example:
	//
	// dataCoord:\\n  segment:\\n    maxSize: 1024
	UserConfig *string                                       `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
	VSwitches  []*GetInstanceDetailResponseBodyDataVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The kernel version.
	//
	// example:
	//
	// 2.4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-123xxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-123xxx
	VswId *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
	// The zone.
	//
	// example:
	//
	// cn-beijing-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetInstanceDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceDetailResponseBodyData) GetAclId() *string {
	return s.AclId
}

func (s *GetInstanceDetailResponseBodyData) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetInstanceDetailResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *GetInstanceDetailResponseBodyData) GetBucketPath() *string {
	return s.BucketPath
}

func (s *GetInstanceDetailResponseBodyData) GetClusterInfo() *GetInstanceDetailResponseBodyDataClusterInfo {
	return s.ClusterInfo
}

func (s *GetInstanceDetailResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetInstanceDetailResponseBodyData) GetEnableHa() *bool {
	return s.EnableHa
}

func (s *GetInstanceDetailResponseBodyData) GetEncrypted() *string {
	return s.Encrypted
}

func (s *GetInstanceDetailResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetInstanceDetailResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceDetailResponseBodyData) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *GetInstanceDetailResponseBodyData) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *GetInstanceDetailResponseBodyData) GetMeasureConfig() *GetInstanceDetailResponseBodyDataMeasureConfig {
	return s.MeasureConfig
}

func (s *GetInstanceDetailResponseBodyData) GetMultiZoneMode() *string {
	return s.MultiZoneMode
}

func (s *GetInstanceDetailResponseBodyData) GetNodeType() *string {
	return s.NodeType
}

func (s *GetInstanceDetailResponseBodyData) GetOpenPublicNet() *bool {
	return s.OpenPublicNet
}

func (s *GetInstanceDetailResponseBodyData) GetPackageType() *string {
	return s.PackageType
}

func (s *GetInstanceDetailResponseBodyData) GetPayType() *int32 {
	return s.PayType
}

func (s *GetInstanceDetailResponseBodyData) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetInstanceDetailResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceDetailResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetInstanceDetailResponseBodyData) GetRunningTime() *int64 {
	return s.RunningTime
}

func (s *GetInstanceDetailResponseBodyData) GetSgId() *string {
	return s.SgId
}

func (s *GetInstanceDetailResponseBodyData) GetTags() []*GetInstanceDetailResponseBodyDataTags {
	return s.Tags
}

func (s *GetInstanceDetailResponseBodyData) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetInstanceDetailResponseBodyData) GetUserConfig() *string {
	return s.UserConfig
}

func (s *GetInstanceDetailResponseBodyData) GetVSwitches() []*GetInstanceDetailResponseBodyDataVSwitches {
	return s.VSwitches
}

func (s *GetInstanceDetailResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetInstanceDetailResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *GetInstanceDetailResponseBodyData) GetVswId() *string {
	return s.VswId
}

func (s *GetInstanceDetailResponseBodyData) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetInstanceDetailResponseBodyData) SetAclId(v string) *GetInstanceDetailResponseBodyData {
	s.AclId = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetBeginTime(v int64) *GetInstanceDetailResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetBucketName(v string) *GetInstanceDetailResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetBucketPath(v string) *GetInstanceDetailResponseBodyData {
	s.BucketPath = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetClusterInfo(v *GetInstanceDetailResponseBodyDataClusterInfo) *GetInstanceDetailResponseBodyData {
	s.ClusterInfo = v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetClusterName(v string) *GetInstanceDetailResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetEnableHa(v bool) *GetInstanceDetailResponseBodyData {
	s.EnableHa = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetEncrypted(v string) *GetInstanceDetailResponseBodyData {
	s.Encrypted = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetExpireTime(v int64) *GetInstanceDetailResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetInstanceId(v string) *GetInstanceDetailResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetInstanceStatus(v string) *GetInstanceDetailResponseBodyData {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetKmsKeyId(v string) *GetInstanceDetailResponseBodyData {
	s.KmsKeyId = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetMeasureConfig(v *GetInstanceDetailResponseBodyDataMeasureConfig) *GetInstanceDetailResponseBodyData {
	s.MeasureConfig = v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetMultiZoneMode(v string) *GetInstanceDetailResponseBodyData {
	s.MultiZoneMode = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetNodeType(v string) *GetInstanceDetailResponseBodyData {
	s.NodeType = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetOpenPublicNet(v bool) *GetInstanceDetailResponseBodyData {
	s.OpenPublicNet = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetPackageType(v string) *GetInstanceDetailResponseBodyData {
	s.PackageType = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetPayType(v int32) *GetInstanceDetailResponseBodyData {
	s.PayType = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetProductCode(v string) *GetInstanceDetailResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetRegionId(v string) *GetInstanceDetailResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetResourceGroupId(v string) *GetInstanceDetailResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetRunningTime(v int64) *GetInstanceDetailResponseBodyData {
	s.RunningTime = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetSgId(v string) *GetInstanceDetailResponseBodyData {
	s.SgId = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetTags(v []*GetInstanceDetailResponseBodyDataTags) *GetInstanceDetailResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetTemplateVersion(v string) *GetInstanceDetailResponseBodyData {
	s.TemplateVersion = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetUserConfig(v string) *GetInstanceDetailResponseBodyData {
	s.UserConfig = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetVSwitches(v []*GetInstanceDetailResponseBodyDataVSwitches) *GetInstanceDetailResponseBodyData {
	s.VSwitches = v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetVersion(v string) *GetInstanceDetailResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetVpcId(v string) *GetInstanceDetailResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetVswId(v string) *GetInstanceDetailResponseBodyData {
	s.VswId = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetZoneId(v string) *GetInstanceDetailResponseBodyData {
	s.ZoneId = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) Validate() error {
	if s.ClusterInfo != nil {
		if err := s.ClusterInfo.Validate(); err != nil {
			return err
		}
	}
	if s.MeasureConfig != nil {
		if err := s.MeasureConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VSwitches != nil {
		for _, item := range s.VSwitches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceDetailResponseBodyDataClusterInfo struct {
	// The port of the Attu component.
	//
	// example:
	//
	// 3000
	AttuPort *int32 `json:"AttuPort,omitempty" xml:"AttuPort,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// c-123xxx-.milvus.aliyuncs.com
	InternetUrl *string `json:"InternetUrl,omitempty" xml:"InternetUrl,omitempty"`
	// The internal IP address.
	//
	// example:
	//
	// c-123xxx-internal.milvus.aliyuncs.com
	IntranetUrl *string `json:"IntranetUrl,omitempty" xml:"IntranetUrl,omitempty"`
	// The resource details.
	MilvusResourceInfoList []*GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList `json:"MilvusResourceInfoList,omitempty" xml:"MilvusResourceInfoList,omitempty" type:"Repeated"`
	// The size of the data stored in OSS.
	//
	// example:
	//
	// 125.15 GB
	OssStorageSize *string `json:"OssStorageSize,omitempty" xml:"OssStorageSize,omitempty"`
	// The timestamp when the OSS metric is stored.
	//
	// example:
	//
	// 1718604000000
	OssStorageTimestamp *int64 `json:"OssStorageTimestamp,omitempty" xml:"OssStorageTimestamp,omitempty"`
	// The proxy port.
	//
	// example:
	//
	// 19530
	ProxyPort *int32 `json:"ProxyPort,omitempty" xml:"ProxyPort,omitempty"`
	// The total number of CUs.
	//
	// example:
	//
	// 24
	TotalCuNum *int32 `json:"TotalCuNum,omitempty" xml:"TotalCuNum,omitempty"`
	// The total number of disks.
	//
	// example:
	//
	// 500
	TotalDiskSize *int32 `json:"TotalDiskSize,omitempty" xml:"TotalDiskSize,omitempty"`
}

func (s GetInstanceDetailResponseBodyDataClusterInfo) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDetailResponseBodyDataClusterInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) GetAttuPort() *int32 {
	return s.AttuPort
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) GetInternetUrl() *string {
	return s.InternetUrl
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) GetIntranetUrl() *string {
	return s.IntranetUrl
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) GetMilvusResourceInfoList() []*GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList {
	return s.MilvusResourceInfoList
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) GetOssStorageSize() *string {
	return s.OssStorageSize
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) GetOssStorageTimestamp() *int64 {
	return s.OssStorageTimestamp
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) GetProxyPort() *int32 {
	return s.ProxyPort
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) GetTotalCuNum() *int32 {
	return s.TotalCuNum
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) GetTotalDiskSize() *int32 {
	return s.TotalDiskSize
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) SetAttuPort(v int32) *GetInstanceDetailResponseBodyDataClusterInfo {
	s.AttuPort = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) SetInternetUrl(v string) *GetInstanceDetailResponseBodyDataClusterInfo {
	s.InternetUrl = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) SetIntranetUrl(v string) *GetInstanceDetailResponseBodyDataClusterInfo {
	s.IntranetUrl = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) SetMilvusResourceInfoList(v []*GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) *GetInstanceDetailResponseBodyDataClusterInfo {
	s.MilvusResourceInfoList = v
	return s
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) SetOssStorageSize(v string) *GetInstanceDetailResponseBodyDataClusterInfo {
	s.OssStorageSize = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) SetOssStorageTimestamp(v int64) *GetInstanceDetailResponseBodyDataClusterInfo {
	s.OssStorageTimestamp = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) SetProxyPort(v int32) *GetInstanceDetailResponseBodyDataClusterInfo {
	s.ProxyPort = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) SetTotalCuNum(v int32) *GetInstanceDetailResponseBodyDataClusterInfo {
	s.TotalCuNum = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) SetTotalDiskSize(v int32) *GetInstanceDetailResponseBodyDataClusterInfo {
	s.TotalDiskSize = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataClusterInfo) Validate() error {
	if s.MilvusResourceInfoList != nil {
		for _, item := range s.MilvusResourceInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList struct {
	// The component type. Valid values:
	//
	// 	- standalone
	//
	// 	- proxy
	//
	// 	- mix_coordinator
	//
	// 	- query
	//
	// 	- index
	//
	// 	- data
	//
	// example:
	//
	// QueryNode
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The number of CUs.
	//
	// example:
	//
	// 8
	CuNum *int32 `json:"CuNum,omitempty" xml:"CuNum,omitempty"`
	// example:
	//
	// 4
	CuRatio *int32 `json:"CuRatio,omitempty" xml:"CuRatio,omitempty"`
	// The disk size.
	//
	// example:
	//
	// 100
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type.
	//
	// example:
	//
	// essd
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The number of replicas.
	//
	// example:
	//
	// 2
	Replica *int32 `json:"Replica,omitempty" xml:"Replica,omitempty"`
}

func (s GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) GoString() string {
	return s.String()
}

func (s *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) GetComponentType() *string {
	return s.ComponentType
}

func (s *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) GetCuNum() *int32 {
	return s.CuNum
}

func (s *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) GetCuRatio() *int32 {
	return s.CuRatio
}

func (s *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) GetDiskType() *string {
	return s.DiskType
}

func (s *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) GetReplica() *int32 {
	return s.Replica
}

func (s *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) SetComponentType(v string) *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList {
	s.ComponentType = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) SetCuNum(v int32) *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList {
	s.CuNum = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) SetCuRatio(v int32) *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList {
	s.CuRatio = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) SetDiskSize(v int32) *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList {
	s.DiskSize = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) SetDiskType(v string) *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList {
	s.DiskType = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) SetReplica(v int32) *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList {
	s.Replica = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) Validate() error {
	return dara.Validate(s)
}

type GetInstanceDetailResponseBodyDataMeasureConfig struct {
	// example:
	//
	// 4
	DataNodeCuNum *int32 `json:"DataNodeCuNum,omitempty" xml:"DataNodeCuNum,omitempty"`
	// example:
	//
	// 2
	DataNodeReplica *int32 `json:"DataNodeReplica,omitempty" xml:"DataNodeReplica,omitempty"`
	// example:
	//
	// 4
	IndexNodeCuNum *int32 `json:"IndexNodeCuNum,omitempty" xml:"IndexNodeCuNum,omitempty"`
	// example:
	//
	// 2
	IndexNodeReplica *int32 `json:"IndexNodeReplica,omitempty" xml:"IndexNodeReplica,omitempty"`
	// example:
	//
	// 4
	MixCoodinatorNodeCuNum *int32 `json:"MixCoodinatorNodeCuNum,omitempty" xml:"MixCoodinatorNodeCuNum,omitempty"`
	// example:
	//
	// 2
	MixCoodinatorNodeReplica *int32 `json:"MixCoodinatorNodeReplica,omitempty" xml:"MixCoodinatorNodeReplica,omitempty"`
	// example:
	//
	// 4
	ProxyNodeCuNum *int32 `json:"ProxyNodeCuNum,omitempty" xml:"ProxyNodeCuNum,omitempty"`
	// example:
	//
	// 2
	ProxyNodeReplica *int32 `json:"ProxyNodeReplica,omitempty" xml:"ProxyNodeReplica,omitempty"`
	// example:
	//
	// 4
	QueryNodeCuNum *int32 `json:"QueryNodeCuNum,omitempty" xml:"QueryNodeCuNum,omitempty"`
	// example:
	//
	// 2
	QueryNodeReplica *int32 `json:"QueryNodeReplica,omitempty" xml:"QueryNodeReplica,omitempty"`
}

func (s GetInstanceDetailResponseBodyDataMeasureConfig) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDetailResponseBodyDataMeasureConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) GetDataNodeCuNum() *int32 {
	return s.DataNodeCuNum
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) GetDataNodeReplica() *int32 {
	return s.DataNodeReplica
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) GetIndexNodeCuNum() *int32 {
	return s.IndexNodeCuNum
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) GetIndexNodeReplica() *int32 {
	return s.IndexNodeReplica
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) GetMixCoodinatorNodeCuNum() *int32 {
	return s.MixCoodinatorNodeCuNum
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) GetMixCoodinatorNodeReplica() *int32 {
	return s.MixCoodinatorNodeReplica
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) GetProxyNodeCuNum() *int32 {
	return s.ProxyNodeCuNum
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) GetProxyNodeReplica() *int32 {
	return s.ProxyNodeReplica
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) GetQueryNodeCuNum() *int32 {
	return s.QueryNodeCuNum
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) GetQueryNodeReplica() *int32 {
	return s.QueryNodeReplica
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) SetDataNodeCuNum(v int32) *GetInstanceDetailResponseBodyDataMeasureConfig {
	s.DataNodeCuNum = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) SetDataNodeReplica(v int32) *GetInstanceDetailResponseBodyDataMeasureConfig {
	s.DataNodeReplica = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) SetIndexNodeCuNum(v int32) *GetInstanceDetailResponseBodyDataMeasureConfig {
	s.IndexNodeCuNum = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) SetIndexNodeReplica(v int32) *GetInstanceDetailResponseBodyDataMeasureConfig {
	s.IndexNodeReplica = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) SetMixCoodinatorNodeCuNum(v int32) *GetInstanceDetailResponseBodyDataMeasureConfig {
	s.MixCoodinatorNodeCuNum = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) SetMixCoodinatorNodeReplica(v int32) *GetInstanceDetailResponseBodyDataMeasureConfig {
	s.MixCoodinatorNodeReplica = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) SetProxyNodeCuNum(v int32) *GetInstanceDetailResponseBodyDataMeasureConfig {
	s.ProxyNodeCuNum = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) SetProxyNodeReplica(v int32) *GetInstanceDetailResponseBodyDataMeasureConfig {
	s.ProxyNodeReplica = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) SetQueryNodeCuNum(v int32) *GetInstanceDetailResponseBodyDataMeasureConfig {
	s.QueryNodeCuNum = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) SetQueryNodeReplica(v int32) *GetInstanceDetailResponseBodyDataMeasureConfig {
	s.QueryNodeReplica = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataMeasureConfig) Validate() error {
	return dara.Validate(s)
}

type GetInstanceDetailResponseBodyDataTags struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceDetailResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDetailResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetInstanceDetailResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *GetInstanceDetailResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *GetInstanceDetailResponseBodyDataTags) SetKey(v string) *GetInstanceDetailResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataTags) SetValue(v string) *GetInstanceDetailResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}

type GetInstanceDetailResponseBodyDataVSwitches struct {
	VswId  *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetInstanceDetailResponseBodyDataVSwitches) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDetailResponseBodyDataVSwitches) GoString() string {
	return s.String()
}

func (s *GetInstanceDetailResponseBodyDataVSwitches) GetVswId() *string {
	return s.VswId
}

func (s *GetInstanceDetailResponseBodyDataVSwitches) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetInstanceDetailResponseBodyDataVSwitches) SetVswId(v string) *GetInstanceDetailResponseBodyDataVSwitches {
	s.VswId = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataVSwitches) SetZoneId(v string) *GetInstanceDetailResponseBodyDataVSwitches {
	s.ZoneId = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataVSwitches) Validate() error {
	return dara.Validate(s)
}
