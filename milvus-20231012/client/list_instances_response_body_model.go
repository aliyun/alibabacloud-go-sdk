// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListInstancesResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*ListInstancesResponseBodyData) *ListInstancesResponseBody
	GetData() []*ListInstancesResponseBodyData
	SetErrCode(v string) *ListInstancesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListInstancesResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ListInstancesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstancesResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListInstancesResponseBody
	GetTotal() *int32
}

type ListInstancesResponseBody struct {
	// The detailed information about the failed permission verification.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxxx",     "EncodedDiagnosticMessage": "xxxx",     "AuthPrincipalType": "SubUser",     "AuthPrincipalDisplayName": "xxxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "milvus:xxxx" }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned result.
	Data []*ListInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// The total number.
	//
	// example:
	//
	// 15
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListInstancesResponseBody) GetData() []*ListInstancesResponseBodyData {
	return s.Data
}

func (s *ListInstancesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListInstancesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListInstancesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstancesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListInstancesResponseBody) SetAccessDeniedDetail(v string) *ListInstancesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListInstancesResponseBody) SetData(v []*ListInstancesResponseBodyData) *ListInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListInstancesResponseBody) SetErrCode(v string) *ListInstancesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetErrMessage(v string) *ListInstancesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListInstancesResponseBody) SetHttpStatusCode(v int32) *ListInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetSuccess(v bool) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotal(v int32) *ListInstancesResponseBody {
	s.Total = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyData struct {
	// example:
	//
	// true
	AutoBackup *bool `json:"AutoBackup,omitempty" xml:"AutoBackup,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1718608505000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The instance details.
	ClusterInfo *ListInstancesResponseBodyDataClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Struct"`
	// The instance name.
	//
	// example:
	//
	// milvus-test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 1721232000000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the bastion host. Valid values:
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
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
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
	// trial
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
	// 536611
	RunningTime *int32 `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-123xxx
	SgId *string                              `json:"SgId,omitempty" xml:"SgId,omitempty"`
	Tags []*ListInstancesResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s ListInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyData) GetAutoBackup() *bool {
	return s.AutoBackup
}

func (s *ListInstancesResponseBodyData) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *ListInstancesResponseBodyData) GetClusterInfo() *ListInstancesResponseBodyDataClusterInfo {
	return s.ClusterInfo
}

func (s *ListInstancesResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListInstancesResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *ListInstancesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyData) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *ListInstancesResponseBodyData) GetNodeType() *string {
	return s.NodeType
}

func (s *ListInstancesResponseBodyData) GetOpenPublicNet() *bool {
	return s.OpenPublicNet
}

func (s *ListInstancesResponseBodyData) GetPackageType() *string {
	return s.PackageType
}

func (s *ListInstancesResponseBodyData) GetPayType() *int32 {
	return s.PayType
}

func (s *ListInstancesResponseBodyData) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListInstancesResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesResponseBodyData) GetRunningTime() *int32 {
	return s.RunningTime
}

func (s *ListInstancesResponseBodyData) GetSgId() *string {
	return s.SgId
}

func (s *ListInstancesResponseBodyData) GetTags() []*ListInstancesResponseBodyDataTags {
	return s.Tags
}

func (s *ListInstancesResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *ListInstancesResponseBodyData) GetVswId() *string {
	return s.VswId
}

func (s *ListInstancesResponseBodyData) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListInstancesResponseBodyData) SetAutoBackup(v bool) *ListInstancesResponseBodyData {
	s.AutoBackup = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetBeginTime(v int64) *ListInstancesResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetClusterInfo(v *ListInstancesResponseBodyDataClusterInfo) *ListInstancesResponseBodyData {
	s.ClusterInfo = v
	return s
}

func (s *ListInstancesResponseBodyData) SetClusterName(v string) *ListInstancesResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetExpireTime(v int64) *ListInstancesResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetInstanceId(v string) *ListInstancesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetInstanceStatus(v string) *ListInstancesResponseBodyData {
	s.InstanceStatus = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetNodeType(v string) *ListInstancesResponseBodyData {
	s.NodeType = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetOpenPublicNet(v bool) *ListInstancesResponseBodyData {
	s.OpenPublicNet = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetPackageType(v string) *ListInstancesResponseBodyData {
	s.PackageType = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetPayType(v int32) *ListInstancesResponseBodyData {
	s.PayType = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetProductCode(v string) *ListInstancesResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetRegionId(v string) *ListInstancesResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetResourceGroupId(v string) *ListInstancesResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetRunningTime(v int32) *ListInstancesResponseBodyData {
	s.RunningTime = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetSgId(v string) *ListInstancesResponseBodyData {
	s.SgId = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetTags(v []*ListInstancesResponseBodyDataTags) *ListInstancesResponseBodyData {
	s.Tags = v
	return s
}

func (s *ListInstancesResponseBodyData) SetVpcId(v string) *ListInstancesResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetVswId(v string) *ListInstancesResponseBodyData {
	s.VswId = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetZoneId(v string) *ListInstancesResponseBodyData {
	s.ZoneId = &v
	return s
}

func (s *ListInstancesResponseBodyData) Validate() error {
	if s.ClusterInfo != nil {
		if err := s.ClusterInfo.Validate(); err != nil {
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
	return nil
}

type ListInstancesResponseBodyDataClusterInfo struct {
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
	// c-123xxx.milvus.aliyuncs.com
	InternetUrl *string `json:"InternetUrl,omitempty" xml:"InternetUrl,omitempty"`
	// The internal endpoint.
	//
	// example:
	//
	// c-123xxx-internal.milvus.aliyuncs.com
	IntranetUrl *string `json:"IntranetUrl,omitempty" xml:"IntranetUrl,omitempty"`
	// The resource details.
	MilvusResourceInfoList []*ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList `json:"MilvusResourceInfoList,omitempty" xml:"MilvusResourceInfoList,omitempty" type:"Repeated"`
	// The proxy port.
	//
	// example:
	//
	// 19530
	ProxyPort *int32 `json:"ProxyPort,omitempty" xml:"ProxyPort,omitempty"`
	// The number of CUs.
	//
	// example:
	//
	// 24
	TotalCuNum *int32 `json:"TotalCuNum,omitempty" xml:"TotalCuNum,omitempty"`
	// The total capacity of the disk.
	//
	// example:
	//
	// 500
	TotalDiskSize *int32 `json:"TotalDiskSize,omitempty" xml:"TotalDiskSize,omitempty"`
}

func (s ListInstancesResponseBodyDataClusterInfo) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyDataClusterInfo) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataClusterInfo) GetAttuPort() *int32 {
	return s.AttuPort
}

func (s *ListInstancesResponseBodyDataClusterInfo) GetInternetUrl() *string {
	return s.InternetUrl
}

func (s *ListInstancesResponseBodyDataClusterInfo) GetIntranetUrl() *string {
	return s.IntranetUrl
}

func (s *ListInstancesResponseBodyDataClusterInfo) GetMilvusResourceInfoList() []*ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList {
	return s.MilvusResourceInfoList
}

func (s *ListInstancesResponseBodyDataClusterInfo) GetProxyPort() *int32 {
	return s.ProxyPort
}

func (s *ListInstancesResponseBodyDataClusterInfo) GetTotalCuNum() *int32 {
	return s.TotalCuNum
}

func (s *ListInstancesResponseBodyDataClusterInfo) GetTotalDiskSize() *int32 {
	return s.TotalDiskSize
}

func (s *ListInstancesResponseBodyDataClusterInfo) SetAttuPort(v int32) *ListInstancesResponseBodyDataClusterInfo {
	s.AttuPort = &v
	return s
}

func (s *ListInstancesResponseBodyDataClusterInfo) SetInternetUrl(v string) *ListInstancesResponseBodyDataClusterInfo {
	s.InternetUrl = &v
	return s
}

func (s *ListInstancesResponseBodyDataClusterInfo) SetIntranetUrl(v string) *ListInstancesResponseBodyDataClusterInfo {
	s.IntranetUrl = &v
	return s
}

func (s *ListInstancesResponseBodyDataClusterInfo) SetMilvusResourceInfoList(v []*ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList) *ListInstancesResponseBodyDataClusterInfo {
	s.MilvusResourceInfoList = v
	return s
}

func (s *ListInstancesResponseBodyDataClusterInfo) SetProxyPort(v int32) *ListInstancesResponseBodyDataClusterInfo {
	s.ProxyPort = &v
	return s
}

func (s *ListInstancesResponseBodyDataClusterInfo) SetTotalCuNum(v int32) *ListInstancesResponseBodyDataClusterInfo {
	s.TotalCuNum = &v
	return s
}

func (s *ListInstancesResponseBodyDataClusterInfo) SetTotalDiskSize(v int32) *ListInstancesResponseBodyDataClusterInfo {
	s.TotalDiskSize = &v
	return s
}

func (s *ListInstancesResponseBodyDataClusterInfo) Validate() error {
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

type ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList struct {
	// The type of the component. Valid values:
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
	// standalone
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The number of CUs.
	//
	// example:
	//
	// 16
	CuNum *int32 `json:"CuNum,omitempty" xml:"CuNum,omitempty"`
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
	// 1
	Replica *int32 `json:"Replica,omitempty" xml:"Replica,omitempty"`
}

func (s ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList) GetComponentType() *string {
	return s.ComponentType
}

func (s *ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList) GetCuNum() *int32 {
	return s.CuNum
}

func (s *ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList) GetDiskType() *string {
	return s.DiskType
}

func (s *ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList) GetReplica() *int32 {
	return s.Replica
}

func (s *ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList) SetComponentType(v string) *ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList {
	s.ComponentType = &v
	return s
}

func (s *ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList) SetCuNum(v int32) *ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList {
	s.CuNum = &v
	return s
}

func (s *ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList) SetDiskSize(v int32) *ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList {
	s.DiskSize = &v
	return s
}

func (s *ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList) SetDiskType(v string) *ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList {
	s.DiskType = &v
	return s
}

func (s *ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList) SetReplica(v int32) *ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList {
	s.Replica = &v
	return s
}

func (s *ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyDataTags struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *ListInstancesResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *ListInstancesResponseBodyDataTags) SetKey(v string) *ListInstancesResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyDataTags) SetValue(v string) *ListInstancesResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *ListInstancesResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}
