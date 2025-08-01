// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListClustersResponseBodyData) *ListClustersResponseBody
	GetData() []*ListClustersResponseBodyData
	SetErrorCode(v string) *ListClustersResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *ListClustersResponseBody
	GetHttpCode() *string
	SetMessage(v string) *ListClustersResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListClustersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListClustersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListClustersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListClustersResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListClustersResponseBody
	GetTotalCount() *int32
}

type ListClustersResponseBody struct {
	// The details of the data.
	Data []*ListClustersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 202
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 69AD2AA7-DB47-449B-941B-B14409DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned instances.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) GetData() []*ListClustersResponseBodyData {
	return s.Data
}

func (s *ListClustersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListClustersResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListClustersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListClustersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListClustersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClustersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListClustersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListClustersResponseBody) SetData(v []*ListClustersResponseBodyData) *ListClustersResponseBody {
	s.Data = v
	return s
}

func (s *ListClustersResponseBody) SetErrorCode(v string) *ListClustersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListClustersResponseBody) SetHttpCode(v string) *ListClustersResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListClustersResponseBody) SetMessage(v string) *ListClustersResponseBody {
	s.Message = &v
	return s
}

func (s *ListClustersResponseBody) SetPageNumber(v int32) *ListClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClustersResponseBody) SetPageSize(v int32) *ListClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetSuccess(v bool) *ListClustersResponseBody {
	s.Success = &v
	return s
}

func (s *ListClustersResponseBody) SetTotalCount(v int32) *ListClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListClustersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListClustersResponseBodyData struct {
	// The application version.
	//
	// example:
	//
	// 1.9.3
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// Indicates whether the instance can be upgraded.
	//
	// example:
	//
	// true
	CanUpdate *bool `json:"CanUpdate,omitempty" xml:"CanUpdate,omitempty"`
	// The billing method, such as subscription or pay-as-you-go.
	//
	// example:
	//
	// Pay-as-you-go
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The alias of the cluster.
	//
	// example:
	//
	// mse-7413****
	ClusterAliasName *string `json:"ClusterAliasName,omitempty" xml:"ClusterAliasName,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// mse-cn-st21ri2****
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The type of the cluster. Valid values: ZooKeeper, Nacos-Ans, and Eureka.
	//
	// example:
	//
	// Eureka
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2020-07-31 11:36:08
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the cluster expires.
	//
	// example:
	//
	// 2021-08-01 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The initialization status of the instance.
	//
	// example:
	//
	// RESTART_SUCCESS
	InitStatus *string `json:"InitStatus,omitempty" xml:"InitStatus,omitempty"`
	// The number of clusters.
	//
	// example:
	//
	// 2
	InstanceCount *int64 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 47.98.XX.XX
	InternetAddress *string `json:"InternetAddress,omitempty" xml:"InternetAddress,omitempty"`
	// The public endpoint.
	//
	// example:
	//
	// mse-7413****-p.eureka.mse.aliyuncs.com
	InternetDomain *string `json:"InternetDomain,omitempty" xml:"InternetDomain,omitempty"`
	// The internal IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	IntranetAddress *string `json:"IntranetAddress,omitempty" xml:"IntranetAddress,omitempty"`
	// The internal endpoint.
	//
	// example:
	//
	// mse-7413****-eureka.mse.aliyuncs.com
	IntranetDomain    *string                                        `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	MaintenancePeriod *ListClustersResponseBodyDataMaintenancePeriod `json:"MaintenancePeriod,omitempty" xml:"MaintenancePeriod,omitempty" type:"Struct"`
	// The edition of the cluster.
	//
	// example:
	//
	// mse_pro
	MseVersion *string `json:"MseVersion,omitempty" xml:"MseVersion,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmv7jiavm4uxa
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags that are attached to the instance.
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The version information.
	//
	// example:
	//
	// EUREKA_1_9_3
	VersionCode      *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionLifecycle *string `json:"VersionLifecycle,omitempty" xml:"VersionLifecycle,omitempty"`
	// example:
	//
	// vpc-bp1hcg467ekqsv0zr****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListClustersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyData) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListClustersResponseBodyData) GetCanUpdate() *bool {
	return s.CanUpdate
}

func (s *ListClustersResponseBodyData) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListClustersResponseBodyData) GetClusterAliasName() *string {
	return s.ClusterAliasName
}

func (s *ListClustersResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListClustersResponseBodyData) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListClustersResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListClustersResponseBodyData) GetEndDate() *string {
	return s.EndDate
}

func (s *ListClustersResponseBodyData) GetInitStatus() *string {
	return s.InitStatus
}

func (s *ListClustersResponseBodyData) GetInstanceCount() *int64 {
	return s.InstanceCount
}

func (s *ListClustersResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListClustersResponseBodyData) GetInternetAddress() *string {
	return s.InternetAddress
}

func (s *ListClustersResponseBodyData) GetInternetDomain() *string {
	return s.InternetDomain
}

func (s *ListClustersResponseBodyData) GetIntranetAddress() *string {
	return s.IntranetAddress
}

func (s *ListClustersResponseBodyData) GetIntranetDomain() *string {
	return s.IntranetDomain
}

func (s *ListClustersResponseBodyData) GetMaintenancePeriod() *ListClustersResponseBodyDataMaintenancePeriod {
	return s.MaintenancePeriod
}

func (s *ListClustersResponseBodyData) GetMseVersion() *string {
	return s.MseVersion
}

func (s *ListClustersResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListClustersResponseBodyData) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListClustersResponseBodyData) GetVersionCode() *string {
	return s.VersionCode
}

func (s *ListClustersResponseBodyData) GetVersionLifecycle() *string {
	return s.VersionLifecycle
}

func (s *ListClustersResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *ListClustersResponseBodyData) SetAppVersion(v string) *ListClustersResponseBodyData {
	s.AppVersion = &v
	return s
}

func (s *ListClustersResponseBodyData) SetCanUpdate(v bool) *ListClustersResponseBodyData {
	s.CanUpdate = &v
	return s
}

func (s *ListClustersResponseBodyData) SetChargeType(v string) *ListClustersResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *ListClustersResponseBodyData) SetClusterAliasName(v string) *ListClustersResponseBodyData {
	s.ClusterAliasName = &v
	return s
}

func (s *ListClustersResponseBodyData) SetClusterName(v string) *ListClustersResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *ListClustersResponseBodyData) SetClusterType(v string) *ListClustersResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *ListClustersResponseBodyData) SetCreateTime(v string) *ListClustersResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListClustersResponseBodyData) SetEndDate(v string) *ListClustersResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *ListClustersResponseBodyData) SetInitStatus(v string) *ListClustersResponseBodyData {
	s.InitStatus = &v
	return s
}

func (s *ListClustersResponseBodyData) SetInstanceCount(v int64) *ListClustersResponseBodyData {
	s.InstanceCount = &v
	return s
}

func (s *ListClustersResponseBodyData) SetInstanceId(v string) *ListClustersResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListClustersResponseBodyData) SetInternetAddress(v string) *ListClustersResponseBodyData {
	s.InternetAddress = &v
	return s
}

func (s *ListClustersResponseBodyData) SetInternetDomain(v string) *ListClustersResponseBodyData {
	s.InternetDomain = &v
	return s
}

func (s *ListClustersResponseBodyData) SetIntranetAddress(v string) *ListClustersResponseBodyData {
	s.IntranetAddress = &v
	return s
}

func (s *ListClustersResponseBodyData) SetIntranetDomain(v string) *ListClustersResponseBodyData {
	s.IntranetDomain = &v
	return s
}

func (s *ListClustersResponseBodyData) SetMaintenancePeriod(v *ListClustersResponseBodyDataMaintenancePeriod) *ListClustersResponseBodyData {
	s.MaintenancePeriod = v
	return s
}

func (s *ListClustersResponseBodyData) SetMseVersion(v string) *ListClustersResponseBodyData {
	s.MseVersion = &v
	return s
}

func (s *ListClustersResponseBodyData) SetResourceGroupId(v string) *ListClustersResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClustersResponseBodyData) SetTags(v map[string]interface{}) *ListClustersResponseBodyData {
	s.Tags = v
	return s
}

func (s *ListClustersResponseBodyData) SetVersionCode(v string) *ListClustersResponseBodyData {
	s.VersionCode = &v
	return s
}

func (s *ListClustersResponseBodyData) SetVersionLifecycle(v string) *ListClustersResponseBodyData {
	s.VersionLifecycle = &v
	return s
}

func (s *ListClustersResponseBodyData) SetVpcId(v string) *ListClustersResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *ListClustersResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListClustersResponseBodyDataMaintenancePeriod struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListClustersResponseBodyDataMaintenancePeriod) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyDataMaintenancePeriod) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyDataMaintenancePeriod) GetEndTime() *string {
	return s.EndTime
}

func (s *ListClustersResponseBodyDataMaintenancePeriod) GetStartTime() *string {
	return s.StartTime
}

func (s *ListClustersResponseBodyDataMaintenancePeriod) SetEndTime(v string) *ListClustersResponseBodyDataMaintenancePeriod {
	s.EndTime = &v
	return s
}

func (s *ListClustersResponseBodyDataMaintenancePeriod) SetStartTime(v string) *ListClustersResponseBodyDataMaintenancePeriod {
	s.StartTime = &v
	return s
}

func (s *ListClustersResponseBodyDataMaintenancePeriod) Validate() error {
	return dara.Validate(s)
}
