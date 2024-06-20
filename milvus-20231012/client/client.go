// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeAccessControlListRequest struct {
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeAccessControlListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessControlListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListRequest) SetInstanceId(v string) *DescribeAccessControlListRequest {
	s.InstanceId = &v
	return s
}

type DescribeAccessControlListResponseBody struct {
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxxx",     "EncodedDiagnosticMessage": "xxxx",     "AuthPrincipalType": "SubUser",     "AuthPrincipalDisplayName": "xxxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "milvus:xxxx" }
	AccessDeniedDetail *string                                    `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *DescribeAccessControlListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Failed to find instance c-123xxx
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// Instance.NotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ABCD-1234-5678-EFGH
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAccessControlListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessControlListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListResponseBody) SetAccessDeniedDetail(v string) *DescribeAccessControlListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeAccessControlListResponseBody) SetData(v *DescribeAccessControlListResponseBodyData) *DescribeAccessControlListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAccessControlListResponseBody) SetErrMessage(v string) *DescribeAccessControlListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeAccessControlListResponseBody) SetErrorCode(v string) *DescribeAccessControlListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeAccessControlListResponseBody) SetHttpStatusCode(v int32) *DescribeAccessControlListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAccessControlListResponseBody) SetRequestId(v string) *DescribeAccessControlListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessControlListResponseBody) SetSuccess(v bool) *DescribeAccessControlListResponseBody {
	s.Success = &v
	return s
}

type DescribeAccessControlListResponseBodyData struct {
	// example:
	//
	// acl-xxxx
	AclId *string   `json:"AclId,omitempty" xml:"AclId,omitempty"`
	Cidr  []*string `json:"Cidr,omitempty" xml:"Cidr,omitempty" type:"Repeated"`
}

func (s DescribeAccessControlListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessControlListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListResponseBodyData) SetAclId(v string) *DescribeAccessControlListResponseBodyData {
	s.AclId = &v
	return s
}

func (s *DescribeAccessControlListResponseBodyData) SetCidr(v []*string) *DescribeAccessControlListResponseBodyData {
	s.Cidr = v
	return s
}

type DescribeAccessControlListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccessControlListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccessControlListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccessControlListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListResponse) SetHeaders(v map[string]*string) *DescribeAccessControlListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessControlListResponse) SetStatusCode(v int32) *DescribeAccessControlListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessControlListResponse) SetBody(v *DescribeAccessControlListResponseBody) *DescribeAccessControlListResponse {
	s.Body = v
	return s
}

type DescribeInstanceConfigsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceConfigsRequest) SetInstanceId(v string) *DescribeInstanceConfigsRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceConfigsResponseBody struct {
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxxx",     "EncodedDiagnosticMessage": "xxxx",     "AuthPrincipalType": "SubUser",     "AuthPrincipalDisplayName": "xxxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "milvus:xxxx" }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// dataCoord:\\n  enableCompaction: true
	Data []byte `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Instance.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Failed to find instance c-123xxx.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ABCD-1234-5678-EFGH
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstanceConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceConfigsResponseBody) SetAccessDeniedDetail(v string) *DescribeInstanceConfigsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetData(v []byte) *DescribeInstanceConfigsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetErrCode(v string) *DescribeInstanceConfigsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetErrMessage(v string) *DescribeInstanceConfigsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetHttpStatusCode(v int32) *DescribeInstanceConfigsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetRequestId(v string) *DescribeInstanceConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceConfigsResponseBody) SetSuccess(v bool) *DescribeInstanceConfigsResponseBody {
	s.Success = &v
	return s
}

type DescribeInstanceConfigsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceConfigsResponse) SetHeaders(v map[string]*string) *DescribeInstanceConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceConfigsResponse) SetStatusCode(v int32) *DescribeInstanceConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceConfigsResponse) SetBody(v *DescribeInstanceConfigsResponseBody) *DescribeInstanceConfigsResponse {
	s.Body = v
	return s
}

type GetInstanceDetailRequest struct {
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstanceDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceDetailRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceDetailRequest) SetInstanceId(v string) *GetInstanceDetailRequest {
	s.InstanceId = &v
	return s
}

type GetInstanceDetailResponseBody struct {
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxxx",     "EncodedDiagnosticMessage": "xxxx",     "AuthPrincipalType": "SubUser",     "AuthPrincipalDisplayName": "xxxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "milvus:xxxx" }
	AccessDeniedDetail *string                            `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *GetInstanceDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Instance.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Failed to find instance c-123xxx
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ABCD-1234-5678-EFGH
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceDetailResponseBody) GoString() string {
	return s.String()
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

type GetInstanceDetailResponseBodyData struct {
	// example:
	//
	// acl-123xxx
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// example:
	//
	// 1716863508000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// milvus-test
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// https://oss.console.aliyun.com/bucket/oss-cn-beijing/xxxx/object?spm=a2cug.25127996.0.0.577990370Ebsqt&path=milvus-c-123xxxx
	BucketPath  *string                                       `json:"BucketPath,omitempty" xml:"BucketPath,omitempty"`
	ClusterInfo *GetInstanceDetailResponseBodyDataClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Struct"`
	// example:
	//
	// aliyun-test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// 1721664000000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// running
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// example:
	//
	// true
	OpenPublicNet *bool `json:"OpenPublicNet,omitempty" xml:"OpenPublicNet,omitempty"`
	// example:
	//
	// standard
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// example:
	//
	// 1
	PayType *int32 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1743679
	RunningTime *int64 `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	// example:
	//
	// sg-123xxx
	SgId *string `json:"SgId,omitempty" xml:"SgId,omitempty"`
	// example:
	//
	// 2.4.1-1.0-0.0.1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// example:
	//
	// dataCoord:\\n  segment:\\n    maxSize: 1024
	UserConfig *string `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
	// example:
	//
	// 2.4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// vpc-123xxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-123xxx
	VswId *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
	// example:
	//
	// cn-beijing-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetInstanceDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceDetailResponseBodyData) GoString() string {
	return s.String()
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

func (s *GetInstanceDetailResponseBodyData) SetRegionId(v string) *GetInstanceDetailResponseBodyData {
	s.RegionId = &v
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

func (s *GetInstanceDetailResponseBodyData) SetTemplateVersion(v string) *GetInstanceDetailResponseBodyData {
	s.TemplateVersion = &v
	return s
}

func (s *GetInstanceDetailResponseBodyData) SetUserConfig(v string) *GetInstanceDetailResponseBodyData {
	s.UserConfig = &v
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

type GetInstanceDetailResponseBodyDataClusterInfo struct {
	// example:
	//
	// 3000
	AttuPort *int32 `json:"AttuPort,omitempty" xml:"AttuPort,omitempty"`
	// example:
	//
	// c-123xxx-.milvus.aliyuncs.com
	InternetUrl *string `json:"InternetUrl,omitempty" xml:"InternetUrl,omitempty"`
	// example:
	//
	// c-123xxx-internal.milvus.aliyuncs.com
	IntranetUrl            *string                                                               `json:"IntranetUrl,omitempty" xml:"IntranetUrl,omitempty"`
	MilvusResourceInfoList []*GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList `json:"MilvusResourceInfoList,omitempty" xml:"MilvusResourceInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 125.15 GB
	OssStorageSize *string `json:"OssStorageSize,omitempty" xml:"OssStorageSize,omitempty"`
	// example:
	//
	// 1718604000000
	OssStorageTimestamp *int64 `json:"OssStorageTimestamp,omitempty" xml:"OssStorageTimestamp,omitempty"`
	// example:
	//
	// 19530
	ProxyPort *int32 `json:"ProxyPort,omitempty" xml:"ProxyPort,omitempty"`
	// example:
	//
	// 24
	TotalCuNum *int32 `json:"TotalCuNum,omitempty" xml:"TotalCuNum,omitempty"`
	// example:
	//
	// 500
	TotalDiskSize *int32 `json:"TotalDiskSize,omitempty" xml:"TotalDiskSize,omitempty"`
}

func (s GetInstanceDetailResponseBodyDataClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceDetailResponseBodyDataClusterInfo) GoString() string {
	return s.String()
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

type GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList struct {
	// example:
	//
	// QueryNode
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// example:
	//
	// 8
	CuNum *int32 `json:"CuNum,omitempty" xml:"CuNum,omitempty"`
	// example:
	//
	// 100
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// example:
	//
	// essd
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// example:
	//
	// 2
	Replica *int32 `json:"Replica,omitempty" xml:"Replica,omitempty"`
}

func (s GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) GoString() string {
	return s.String()
}

func (s *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) SetComponentType(v string) *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList {
	s.ComponentType = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList) SetCuNum(v int32) *GetInstanceDetailResponseBodyDataClusterInfoMilvusResourceInfoList {
	s.CuNum = &v
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

type GetInstanceDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceDetailResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceDetailResponse) SetHeaders(v map[string]*string) *GetInstanceDetailResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceDetailResponse) SetStatusCode(v int32) *GetInstanceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceDetailResponse) SetBody(v *GetInstanceDetailResponseBody) *GetInstanceDetailResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// example:
	//
	// c-123xxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// milvus-test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetClusterId(v string) *ListInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInstancesRequest) SetClusterName(v string) *ListInstancesRequest {
	s.ClusterName = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetRegionId(v string) *ListInstancesRequest {
	s.RegionId = &v
	return s
}

type ListInstancesResponseBody struct {
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxxx",     "EncodedDiagnosticMessage": "xxxx",     "AuthPrincipalType": "SubUser",     "AuthPrincipalDisplayName": "xxxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "milvus:xxxx" }
	AccessDeniedDetail *string                          `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               []*ListInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Instance.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Failed to find instance c-123xxx
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ABCD-1234-5678-EFGH
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 15
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
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

type ListInstancesResponseBodyData struct {
	// example:
	//
	// 1718608505000
	BeginTime   *int64                                    `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	ClusterInfo *ListInstancesResponseBodyDataClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Struct"`
	// example:
	//
	// milvus-test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// 1721232000000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// running
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// example:
	//
	// true
	OpenPublicNet *bool `json:"OpenPublicNet,omitempty" xml:"OpenPublicNet,omitempty"`
	// example:
	//
	// trial
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// example:
	//
	// 1
	PayType *int32 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 536611
	RunningTime *int32 `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	// example:
	//
	// sg-123xxx
	SgId *string `json:"SgId,omitempty" xml:"SgId,omitempty"`
	// example:
	//
	// vpc-123xxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-123xxx
	VswId *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
	// example:
	//
	// cn-beijing-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyData) GoString() string {
	return s.String()
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

func (s *ListInstancesResponseBodyData) SetRegionId(v string) *ListInstancesResponseBodyData {
	s.RegionId = &v
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

type ListInstancesResponseBodyDataClusterInfo struct {
	// example:
	//
	// 3000
	AttuPort *int32 `json:"AttuPort,omitempty" xml:"AttuPort,omitempty"`
	// example:
	//
	// c-123xxx.milvus.aliyuncs.com
	InternetUrl *string `json:"InternetUrl,omitempty" xml:"InternetUrl,omitempty"`
	// example:
	//
	// c-123xxx-internal.milvus.aliyuncs.com
	IntranetUrl            *string                                                           `json:"IntranetUrl,omitempty" xml:"IntranetUrl,omitempty"`
	MilvusResourceInfoList []*ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList `json:"MilvusResourceInfoList,omitempty" xml:"MilvusResourceInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 19530
	ProxyPort *int32 `json:"ProxyPort,omitempty" xml:"ProxyPort,omitempty"`
	// example:
	//
	// 24
	TotalCuNum *int32 `json:"TotalCuNum,omitempty" xml:"TotalCuNum,omitempty"`
	// example:
	//
	// 500
	TotalDiskSize *int32 `json:"TotalDiskSize,omitempty" xml:"TotalDiskSize,omitempty"`
}

func (s ListInstancesResponseBodyDataClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyDataClusterInfo) GoString() string {
	return s.String()
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

type ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList struct {
	// example:
	//
	// standalone
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// example:
	//
	// 16
	CuNum *int32 `json:"CuNum,omitempty" xml:"CuNum,omitempty"`
	// example:
	//
	// 100
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// example:
	//
	// essd
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// example:
	//
	// 1
	Replica *int32 `json:"Replica,omitempty" xml:"Replica,omitempty"`
}

func (s ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyDataClusterInfoMilvusResourceInfoList) GoString() string {
	return s.String()
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

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ModifyInstanceConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// for test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// dataCoord:\\n  segment:\\n    maxSize: 1024
	UserConfig *string `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
}

func (s ModifyInstanceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigRequest) SetInstanceId(v string) *ModifyInstanceConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetReason(v string) *ModifyInstanceConfigRequest {
	s.Reason = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetUserConfig(v string) *ModifyInstanceConfigRequest {
	s.UserConfig = &v
	return s
}

type ModifyInstanceConfigResponseBody struct {
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxxx",     "EncodedDiagnosticMessage": "xxxx",     "AuthPrincipalType": "SubUser",     "AuthPrincipalDisplayName": "xxxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "milvus:xxxx" }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Instance.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Failed to find instance c-123xxx
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ABCD-1234-5678-EFGH
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyInstanceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigResponseBody) SetAccessDeniedDetail(v string) *ModifyInstanceConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyInstanceConfigResponseBody) SetData(v bool) *ModifyInstanceConfigResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyInstanceConfigResponseBody) SetErrCode(v string) *ModifyInstanceConfigResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyInstanceConfigResponseBody) SetErrMessage(v string) *ModifyInstanceConfigResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyInstanceConfigResponseBody) SetHttpStatusCode(v int32) *ModifyInstanceConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyInstanceConfigResponseBody) SetRequestId(v string) *ModifyInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceConfigResponseBody) SetSuccess(v bool) *ModifyInstanceConfigResponseBody {
	s.Success = &v
	return s
}

type ModifyInstanceConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigResponse) SetHeaders(v map[string]*string) *ModifyInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceConfigResponse) SetStatusCode(v int32) *ModifyInstanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceConfigResponse) SetBody(v *ModifyInstanceConfigResponseBody) *ModifyInstanceConfigResponse {
	s.Body = v
	return s
}

type UpdateAccessControlListRequest struct {
	// example:
	//
	// acl-123xxx
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// example:
	//
	// 192.168.1.0/24,172.16.0.0/16
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateAccessControlListRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessControlListRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccessControlListRequest) SetAclId(v string) *UpdateAccessControlListRequest {
	s.AclId = &v
	return s
}

func (s *UpdateAccessControlListRequest) SetCidr(v string) *UpdateAccessControlListRequest {
	s.Cidr = &v
	return s
}

func (s *UpdateAccessControlListRequest) SetInstanceId(v string) *UpdateAccessControlListRequest {
	s.InstanceId = &v
	return s
}

type UpdateAccessControlListResponseBody struct {
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxxx",     "EncodedDiagnosticMessage": "xxxx",     "AuthPrincipalType": "SubUser",     "AuthPrincipalDisplayName": "xxxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "milvus:xxxx" }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// The format of cidr is illegal.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// ILLEGAL.MilvusCidrFormat
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ABCD-1234-5678-EFGH
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAccessControlListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessControlListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccessControlListResponseBody) SetAccessDeniedDetail(v string) *UpdateAccessControlListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAccessControlListResponseBody) SetData(v bool) *UpdateAccessControlListResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAccessControlListResponseBody) SetErrMessage(v string) *UpdateAccessControlListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateAccessControlListResponseBody) SetErrorCode(v string) *UpdateAccessControlListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAccessControlListResponseBody) SetHttpStatusCode(v int32) *UpdateAccessControlListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAccessControlListResponseBody) SetRequestId(v string) *UpdateAccessControlListResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAccessControlListResponseBody) SetSuccess(v bool) *UpdateAccessControlListResponseBody {
	s.Success = &v
	return s
}

type UpdateAccessControlListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccessControlListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccessControlListResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessControlListResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccessControlListResponse) SetHeaders(v map[string]*string) *UpdateAccessControlListResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccessControlListResponse) SetStatusCode(v int32) *UpdateAccessControlListResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccessControlListResponse) SetBody(v *UpdateAccessControlListResponseBody) *UpdateAccessControlListResponse {
	s.Body = v
	return s
}

type UpdateInstanceNameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// milvus-test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateInstanceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameRequest) SetClusterName(v string) *UpdateInstanceNameRequest {
	s.ClusterName = &v
	return s
}

func (s *UpdateInstanceNameRequest) SetInstanceId(v string) *UpdateInstanceNameRequest {
	s.InstanceId = &v
	return s
}

type UpdateInstanceNameResponseBody struct {
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxxx",     "EncodedDiagnosticMessage": "xxxx",     "AuthPrincipalType": "SubUser",     "AuthPrincipalDisplayName": "xxxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "milvus:xxxx" }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Instance.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Failed to find instance c-123xxx
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ABCD-1234-5678-EFGH
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameResponseBody) SetAccessDeniedDetail(v string) *UpdateInstanceNameResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetData(v bool) *UpdateInstanceNameResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetErrCode(v string) *UpdateInstanceNameResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetErrMessage(v string) *UpdateInstanceNameResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetHttpStatusCode(v int32) *UpdateInstanceNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetRequestId(v string) *UpdateInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetSuccess(v bool) *UpdateInstanceNameResponseBody {
	s.Success = &v
	return s
}

type UpdateInstanceNameResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameResponse) SetHeaders(v map[string]*string) *UpdateInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceNameResponse) SetStatusCode(v int32) *UpdateInstanceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceNameResponse) SetBody(v *UpdateInstanceNameResponseBody) *UpdateInstanceNameResponse {
	s.Body = v
	return s
}

type UpdatePublicNetworkStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Proxy
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	PublicNetworkEnabled *bool `json:"PublicNetworkEnabled,omitempty" xml:"PublicNetworkEnabled,omitempty"`
}

func (s UpdatePublicNetworkStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicNetworkStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdatePublicNetworkStatusRequest) SetComponentType(v string) *UpdatePublicNetworkStatusRequest {
	s.ComponentType = &v
	return s
}

func (s *UpdatePublicNetworkStatusRequest) SetInstanceId(v string) *UpdatePublicNetworkStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdatePublicNetworkStatusRequest) SetPublicNetworkEnabled(v bool) *UpdatePublicNetworkStatusRequest {
	s.PublicNetworkEnabled = &v
	return s
}

type UpdatePublicNetworkStatusResponseBody struct {
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxxx",     "EncodedDiagnosticMessage": "xxxx",     "AuthPrincipalType": "SubUser",     "AuthPrincipalDisplayName": "xxxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "milvus:xxxx" }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Failed to find instance
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Failed to find instance c-123xxx
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ABCD-1234-5678-EFGH
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePublicNetworkStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicNetworkStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePublicNetworkStatusResponseBody) SetAccessDeniedDetail(v string) *UpdatePublicNetworkStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetData(v bool) *UpdatePublicNetworkStatusResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetErrCode(v string) *UpdatePublicNetworkStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetErrMessage(v string) *UpdatePublicNetworkStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetHttpStatusCode(v int32) *UpdatePublicNetworkStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetRequestId(v string) *UpdatePublicNetworkStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponseBody) SetSuccess(v bool) *UpdatePublicNetworkStatusResponseBody {
	s.Success = &v
	return s
}

type UpdatePublicNetworkStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePublicNetworkStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePublicNetworkStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicNetworkStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdatePublicNetworkStatusResponse) SetHeaders(v map[string]*string) *UpdatePublicNetworkStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdatePublicNetworkStatusResponse) SetStatusCode(v int32) *UpdatePublicNetworkStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePublicNetworkStatusResponse) SetBody(v *UpdatePublicNetworkStatusResponseBody) *UpdatePublicNetworkStatusResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("milvus"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Milvus公网访问ACL信息
//
// @param request - DescribeAccessControlListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessControlListResponse
func (client *Client) DescribeAccessControlListWithOptions(request *DescribeAccessControlListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAccessControlListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccessControlList"),
		Version:     tea.String("2023-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/milvus/describe_access_control_list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccessControlListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Milvus公网访问ACL信息
//
// @param request - DescribeAccessControlListRequest
//
// @return DescribeAccessControlListResponse
func (client *Client) DescribeAccessControlList(request *DescribeAccessControlListRequest) (_result *DescribeAccessControlListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAccessControlListResponse{}
	_body, _err := client.DescribeAccessControlListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例用户配置
//
// @param request - DescribeInstanceConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceConfigsResponse
func (client *Client) DescribeInstanceConfigsWithOptions(request *DescribeInstanceConfigsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeInstanceConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceConfigs"),
		Version:     tea.String("2023-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/config/describe_milvus_user_config"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例用户配置
//
// @param request - DescribeInstanceConfigsRequest
//
// @return DescribeInstanceConfigsResponse
func (client *Client) DescribeInstanceConfigs(request *DescribeInstanceConfigsRequest) (_result *DescribeInstanceConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeInstanceConfigsResponse{}
	_body, _err := client.DescribeInstanceConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据集群ID获取集群的详细信息
//
// @param request - GetInstanceDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceDetailResponse
func (client *Client) GetInstanceDetailWithOptions(request *GetInstanceDetailRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceDetail"),
		Version:     tea.String("2023-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/cluster/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据集群ID获取集群的详细信息
//
// @param request - GetInstanceDetailRequest
//
// @return GetInstanceDetailResponse
func (client *Client) GetInstanceDetail(request *GetInstanceDetailRequest) (_result *GetInstanceDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceDetailResponse{}
	_body, _err := client.GetInstanceDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据集群ID或者名称等信息过滤集群
//
// @param request - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2023-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/order/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据集群ID或者名称等信息过滤集群
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改实例配置
//
// @param request - ModifyInstanceConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceConfigResponse
func (client *Client) ModifyInstanceConfigWithOptions(request *ModifyInstanceConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyInstanceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		query["Reason"] = request.Reason
	}

	if !tea.BoolValue(util.IsUnset(request.UserConfig)) {
		query["UserConfig"] = request.UserConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceConfig"),
		Version:     tea.String("2023-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/config/modify_milvus_config"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改实例配置
//
// @param request - ModifyInstanceConfigRequest
//
// @return ModifyInstanceConfigResponse
func (client *Client) ModifyInstanceConfig(request *ModifyInstanceConfigRequest) (_result *ModifyInstanceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyInstanceConfigResponse{}
	_body, _err := client.ModifyInstanceConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新Milvus公网访问ACL信息
//
// @param request - UpdateAccessControlListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAccessControlListResponse
func (client *Client) UpdateAccessControlListWithOptions(request *UpdateAccessControlListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAccessControlListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.Cidr)) {
		query["Cidr"] = request.Cidr
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAccessControlList"),
		Version:     tea.String("2023-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/milvus/update_access_control_list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAccessControlListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Milvus公网访问ACL信息
//
// @param request - UpdateAccessControlListRequest
//
// @return UpdateAccessControlListResponse
func (client *Client) UpdateAccessControlList(request *UpdateAccessControlListRequest) (_result *UpdateAccessControlListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAccessControlListResponse{}
	_body, _err := client.UpdateAccessControlListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改集群名称
//
// @param request - UpdateInstanceNameRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceNameResponse
func (client *Client) UpdateInstanceNameWithOptions(request *UpdateInstanceNameRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceName"),
		Version:     tea.String("2023-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/cluster/update_name"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改集群名称
//
// @param request - UpdateInstanceNameRequest
//
// @return UpdateInstanceNameResponse
func (client *Client) UpdateInstanceName(request *UpdateInstanceNameRequest) (_result *UpdateInstanceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceNameResponse{}
	_body, _err := client.UpdateInstanceNameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 该接口用于开通/关闭 Proxy的公网SLB。
//
// @param request - UpdatePublicNetworkStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePublicNetworkStatusResponse
func (client *Client) UpdatePublicNetworkStatusWithOptions(request *UpdatePublicNetworkStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdatePublicNetworkStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComponentType)) {
		query["ComponentType"] = request.ComponentType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PublicNetworkEnabled)) {
		query["PublicNetworkEnabled"] = request.PublicNetworkEnabled
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePublicNetworkStatus"),
		Version:     tea.String("2023-10-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/network/updatePublicNetworkStatus"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePublicNetworkStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 该接口用于开通/关闭 Proxy的公网SLB。
//
// @param request - UpdatePublicNetworkStatusRequest
//
// @return UpdatePublicNetworkStatusResponse
func (client *Client) UpdatePublicNetworkStatus(request *UpdatePublicNetworkStatusRequest) (_result *UpdatePublicNetworkStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePublicNetworkStatusResponse{}
	_body, _err := client.UpdatePublicNetworkStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
