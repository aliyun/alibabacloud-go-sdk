// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetInstanceRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRequest) SetRegionId(v string) *GetInstanceRequest {
	s.RegionId = &v
	return s
}

type GetInstanceResponseBody struct {
	ErrorCode      *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode *string                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Instance       *GetInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetErrorCode(v string) *GetInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetErrorMessage(v string) *GetInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetInstanceResponseBody) SetHttpStatusCode(v string) *GetInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstance(v *GetInstanceResponseBodyInstance) *GetInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetSuccess(v bool) *GetInstanceResponseBody {
	s.Success = &v
	return s
}

type GetInstanceResponseBodyInstance struct {
	AutoRenewal        *string                                     `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	ColdStorage        *int64                                      `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	CommodityCode      *string                                     `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ComputeNodeCount   *int64                                      `json:"ComputeNodeCount,omitempty" xml:"ComputeNodeCount,omitempty"`
	Cpu                *int64                                      `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreationTime       *string                                     `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Disk               *string                                     `json:"Disk,omitempty" xml:"Disk,omitempty"`
	EnableHiveAccess   *string                                     `json:"EnableHiveAccess,omitempty" xml:"EnableHiveAccess,omitempty"`
	Endpoints          []*GetInstanceResponseBodyInstanceEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	ExpirationTime     *string                                     `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	InstanceChargeType *string                                     `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceId         *string                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName       *string                                     `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceOwner      *string                                     `json:"InstanceOwner,omitempty" xml:"InstanceOwner,omitempty"`
	InstanceStatus     *string                                     `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceType       *string                                     `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	LeaderInstanceId   *string                                     `json:"LeaderInstanceId,omitempty" xml:"LeaderInstanceId,omitempty"`
	Memory             *int64                                      `json:"Memory,omitempty" xml:"Memory,omitempty"`
	ResourceGroupId    *string                                     `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SuspendReason      *string                                     `json:"SuspendReason,omitempty" xml:"SuspendReason,omitempty"`
	Tags               []*GetInstanceResponseBodyInstanceTags      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Version            *string                                     `json:"Version,omitempty" xml:"Version,omitempty"`
	ZoneId             *string                                     `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetInstanceResponseBodyInstance) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstance) SetAutoRenewal(v string) *GetInstanceResponseBodyInstance {
	s.AutoRenewal = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetColdStorage(v int64) *GetInstanceResponseBodyInstance {
	s.ColdStorage = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetCommodityCode(v string) *GetInstanceResponseBodyInstance {
	s.CommodityCode = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetComputeNodeCount(v int64) *GetInstanceResponseBodyInstance {
	s.ComputeNodeCount = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetCpu(v int64) *GetInstanceResponseBodyInstance {
	s.Cpu = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetCreationTime(v string) *GetInstanceResponseBodyInstance {
	s.CreationTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDisk(v string) *GetInstanceResponseBodyInstance {
	s.Disk = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetEnableHiveAccess(v string) *GetInstanceResponseBodyInstance {
	s.EnableHiveAccess = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetEndpoints(v []*GetInstanceResponseBodyInstanceEndpoints) *GetInstanceResponseBodyInstance {
	s.Endpoints = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetExpirationTime(v string) *GetInstanceResponseBodyInstance {
	s.ExpirationTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceChargeType(v string) *GetInstanceResponseBodyInstance {
	s.InstanceChargeType = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceId(v string) *GetInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceName(v string) *GetInstanceResponseBodyInstance {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceOwner(v string) *GetInstanceResponseBodyInstance {
	s.InstanceOwner = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceStatus(v string) *GetInstanceResponseBodyInstance {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceType(v string) *GetInstanceResponseBodyInstance {
	s.InstanceType = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetLeaderInstanceId(v string) *GetInstanceResponseBodyInstance {
	s.LeaderInstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetMemory(v int64) *GetInstanceResponseBodyInstance {
	s.Memory = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetResourceGroupId(v string) *GetInstanceResponseBodyInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetSuspendReason(v string) *GetInstanceResponseBodyInstance {
	s.SuspendReason = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetTags(v []*GetInstanceResponseBodyInstanceTags) *GetInstanceResponseBodyInstance {
	s.Tags = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetVersion(v string) *GetInstanceResponseBodyInstance {
	s.Version = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetZoneId(v string) *GetInstanceResponseBodyInstance {
	s.ZoneId = &v
	return s
}

type GetInstanceResponseBodyInstanceEndpoints struct {
	AlternativeEndpoints *string `json:"AlternativeEndpoints,omitempty" xml:"AlternativeEndpoints,omitempty"`
	Enabled              *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Endpoint             *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Type                 *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s GetInstanceResponseBodyInstanceEndpoints) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceEndpoints) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceEndpoints) SetAlternativeEndpoints(v string) *GetInstanceResponseBodyInstanceEndpoints {
	s.AlternativeEndpoints = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceEndpoints) SetEnabled(v bool) *GetInstanceResponseBodyInstanceEndpoints {
	s.Enabled = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceEndpoints) SetEndpoint(v string) *GetInstanceResponseBodyInstanceEndpoints {
	s.Endpoint = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceEndpoints) SetType(v string) *GetInstanceResponseBodyInstanceEndpoints {
	s.Type = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceEndpoints) SetVSwitchId(v string) *GetInstanceResponseBodyInstanceEndpoints {
	s.VSwitchId = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceEndpoints) SetVpcId(v string) *GetInstanceResponseBodyInstanceEndpoints {
	s.VpcId = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceEndpoints) SetVpcInstanceId(v string) *GetInstanceResponseBodyInstanceEndpoints {
	s.VpcInstanceId = &v
	return s
}

type GetInstanceResponseBodyInstanceTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceResponseBodyInstanceTags) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceTags) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceTags) SetKey(v string) *GetInstanceResponseBodyInstanceTags {
	s.Key = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceTags) SetValue(v string) *GetInstanceResponseBodyInstanceTags {
	s.Value = &v
	return s
}

type GetInstanceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetStatusCode(v int32) *GetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	RegionId        *string                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                    `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Tag             []*ListInstancesRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetRegionId(v string) *ListInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstancesRequest) SetResourceGroupId(v string) *ListInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesRequest) SetTag(v []*ListInstancesRequestTag) *ListInstancesRequest {
	s.Tag = v
	return s
}

type ListInstancesRequestTag struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListInstancesRequestTag) SetKey(v string) *ListInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListInstancesRequestTag) SetValue(v string) *ListInstancesRequestTag {
	s.Value = &v
	return s
}

type ListInstancesResponseBody struct {
	ErrorCode      *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode *string                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceList   []*ListInstancesResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetErrorCode(v string) *ListInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetErrorMessage(v string) *ListInstancesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListInstancesResponseBody) SetHttpStatusCode(v string) *ListInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetInstanceList(v []*ListInstancesResponseBodyInstanceList) *ListInstancesResponseBody {
	s.InstanceList = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetSuccess(v string) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

type ListInstancesResponseBodyInstanceList struct {
	CommodityCode      *string                                           `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CreationTime       *string                                           `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	EnableHiveAccess   *string                                           `json:"EnableHiveAccess,omitempty" xml:"EnableHiveAccess,omitempty"`
	Endpoints          []*ListInstancesResponseBodyInstanceListEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	ExpirationTime     *string                                           `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	InstanceChargeType *string                                           `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceId         *string                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName       *string                                           `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceStatus     *string                                           `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceType       *string                                           `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	LeaderInstanceId   *string                                           `json:"LeaderInstanceId,omitempty" xml:"LeaderInstanceId,omitempty"`
	ResourceGroupId    *string                                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SuspendReason      *string                                           `json:"SuspendReason,omitempty" xml:"SuspendReason,omitempty"`
	Tags               []*ListInstancesResponseBodyInstanceListTags      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Version            *string                                           `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListInstancesResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceList) SetCommodityCode(v string) *ListInstancesResponseBodyInstanceList {
	s.CommodityCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetCreationTime(v string) *ListInstancesResponseBodyInstanceList {
	s.CreationTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetEnableHiveAccess(v string) *ListInstancesResponseBodyInstanceList {
	s.EnableHiveAccess = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetEndpoints(v []*ListInstancesResponseBodyInstanceListEndpoints) *ListInstancesResponseBodyInstanceList {
	s.Endpoints = v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetExpirationTime(v string) *ListInstancesResponseBodyInstanceList {
	s.ExpirationTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceChargeType(v string) *ListInstancesResponseBodyInstanceList {
	s.InstanceChargeType = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceId(v string) *ListInstancesResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceName(v string) *ListInstancesResponseBodyInstanceList {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceStatus(v string) *ListInstancesResponseBodyInstanceList {
	s.InstanceStatus = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceType(v string) *ListInstancesResponseBodyInstanceList {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetLeaderInstanceId(v string) *ListInstancesResponseBodyInstanceList {
	s.LeaderInstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetResourceGroupId(v string) *ListInstancesResponseBodyInstanceList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetSuspendReason(v string) *ListInstancesResponseBodyInstanceList {
	s.SuspendReason = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetTags(v []*ListInstancesResponseBodyInstanceListTags) *ListInstancesResponseBodyInstanceList {
	s.Tags = v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetVersion(v string) *ListInstancesResponseBodyInstanceList {
	s.Version = &v
	return s
}

type ListInstancesResponseBodyInstanceListEndpoints struct {
	Enabled   *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Endpoint  *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s ListInstancesResponseBodyInstanceListEndpoints) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceListEndpoints) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetEnabled(v bool) *ListInstancesResponseBodyInstanceListEndpoints {
	s.Enabled = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetEndpoint(v string) *ListInstancesResponseBodyInstanceListEndpoints {
	s.Endpoint = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetType(v string) *ListInstancesResponseBodyInstanceListEndpoints {
	s.Type = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetVSwitchId(v string) *ListInstancesResponseBodyInstanceListEndpoints {
	s.VSwitchId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetVpcId(v string) *ListInstancesResponseBodyInstanceListEndpoints {
	s.VpcId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetVpcInstanceId(v string) *ListInstancesResponseBodyInstanceListEndpoints {
	s.VpcInstanceId = &v
	return s
}

type ListInstancesResponseBodyInstanceListTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesResponseBodyInstanceListTags) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceListTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceListTags) SetKey(v string) *ListInstancesResponseBodyInstanceListTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListTags) SetValue(v string) *ListInstancesResponseBodyInstanceListTags {
	s.Value = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type RestartInstanceRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RestartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartInstanceRequest) SetRegionId(v string) *RestartInstanceRequest {
	s.RegionId = &v
	return s
}

type RestartInstanceResponseBody struct {
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBody) SetData(v bool) *RestartInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *RestartInstanceResponseBody) SetErrorCode(v string) *RestartInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RestartInstanceResponseBody) SetErrorMessage(v string) *RestartInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RestartInstanceResponseBody) SetHttpStatusCode(v string) *RestartInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RestartInstanceResponseBody) SetRequestId(v string) *RestartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartInstanceResponseBody) SetSuccess(v bool) *RestartInstanceResponseBody {
	s.Success = &v
	return s
}

type RestartInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponse) SetHeaders(v map[string]*string) *RestartInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartInstanceResponse) SetStatusCode(v int32) *RestartInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartInstanceResponse) SetBody(v *RestartInstanceResponseBody) *RestartInstanceResponse {
	s.Body = v
	return s
}

type ResumeInstanceRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResumeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeInstanceRequest) GoString() string {
	return s.String()
}

func (s *ResumeInstanceRequest) SetRegionId(v string) *ResumeInstanceRequest {
	s.RegionId = &v
	return s
}

type ResumeInstanceResponseBody struct {
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResumeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeInstanceResponseBody) SetData(v bool) *ResumeInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetErrorCode(v string) *ResumeInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetErrorMessage(v string) *ResumeInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetHttpStatusCode(v string) *ResumeInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetRequestId(v string) *ResumeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetSuccess(v bool) *ResumeInstanceResponseBody {
	s.Success = &v
	return s
}

type ResumeInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResumeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeInstanceResponse) GoString() string {
	return s.String()
}

func (s *ResumeInstanceResponse) SetHeaders(v map[string]*string) *ResumeInstanceResponse {
	s.Headers = v
	return s
}

func (s *ResumeInstanceResponse) SetStatusCode(v int32) *ResumeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeInstanceResponse) SetBody(v *ResumeInstanceResponseBody) *ResumeInstanceResponse {
	s.Body = v
	return s
}

type StopInstanceRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) SetRegionId(v string) *StopInstanceRequest {
	s.RegionId = &v
	return s
}

type StopInstanceResponseBody struct {
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceResponseBody) SetData(v bool) *StopInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *StopInstanceResponseBody) SetErrorCode(v string) *StopInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopInstanceResponseBody) SetErrorMessage(v string) *StopInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopInstanceResponseBody) SetHttpStatusCode(v string) *StopInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopInstanceResponseBody) SetRequestId(v string) *StopInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopInstanceResponseBody) SetSuccess(v bool) *StopInstanceResponseBody {
	s.Success = &v
	return s
}

type StopInstanceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopInstanceResponse) SetHeaders(v map[string]*string) *StopInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopInstanceResponse) SetStatusCode(v int32) *StopInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopInstanceResponse) SetBody(v *StopInstanceResponseBody) *StopInstanceResponse {
	s.Body = v
	return s
}

type UpdateInstanceNameRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
}

func (s UpdateInstanceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameRequest) SetRegionId(v string) *UpdateInstanceNameRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateInstanceNameRequest) SetInstanceName(v string) *UpdateInstanceNameRequest {
	s.InstanceName = &v
	return s
}

type UpdateInstanceNameResponseBody struct {
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameResponseBody) SetData(v bool) *UpdateInstanceNameResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetErrorCode(v string) *UpdateInstanceNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetErrorMessage(v string) *UpdateInstanceNameResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetHttpStatusCode(v string) *UpdateInstanceNameResponseBody {
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateInstanceNetworkTypeRequest struct {
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AnyTunnelToSingleTunnel *string `json:"anyTunnelToSingleTunnel,omitempty" xml:"anyTunnelToSingleTunnel,omitempty"`
	NetworkTypes            *string `json:"networkTypes,omitempty" xml:"networkTypes,omitempty"`
	VSwitchId               *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// VPC ID。
	VpcId      *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VpcOwnerId *string `json:"vpcOwnerId,omitempty" xml:"vpcOwnerId,omitempty"`
	// vpc regionId。
	VpcRegionId *string `json:"vpcRegionId,omitempty" xml:"vpcRegionId,omitempty"`
}

func (s UpdateInstanceNetworkTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNetworkTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkTypeRequest) SetRegionId(v string) *UpdateInstanceNetworkTypeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateInstanceNetworkTypeRequest) SetAnyTunnelToSingleTunnel(v string) *UpdateInstanceNetworkTypeRequest {
	s.AnyTunnelToSingleTunnel = &v
	return s
}

func (s *UpdateInstanceNetworkTypeRequest) SetNetworkTypes(v string) *UpdateInstanceNetworkTypeRequest {
	s.NetworkTypes = &v
	return s
}

func (s *UpdateInstanceNetworkTypeRequest) SetVSwitchId(v string) *UpdateInstanceNetworkTypeRequest {
	s.VSwitchId = &v
	return s
}

func (s *UpdateInstanceNetworkTypeRequest) SetVpcId(v string) *UpdateInstanceNetworkTypeRequest {
	s.VpcId = &v
	return s
}

func (s *UpdateInstanceNetworkTypeRequest) SetVpcOwnerId(v string) *UpdateInstanceNetworkTypeRequest {
	s.VpcOwnerId = &v
	return s
}

func (s *UpdateInstanceNetworkTypeRequest) SetVpcRegionId(v string) *UpdateInstanceNetworkTypeRequest {
	s.VpcRegionId = &v
	return s
}

type UpdateInstanceNetworkTypeResponseBody struct {
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceNetworkTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNetworkTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetData(v string) *UpdateInstanceNetworkTypeResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetErrorCode(v string) *UpdateInstanceNetworkTypeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetErrorMessage(v string) *UpdateInstanceNetworkTypeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetHttpStatusCode(v string) *UpdateInstanceNetworkTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetRequestId(v string) *UpdateInstanceNetworkTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetSuccess(v bool) *UpdateInstanceNetworkTypeResponseBody {
	s.Success = &v
	return s
}

type UpdateInstanceNetworkTypeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInstanceNetworkTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceNetworkTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNetworkTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkTypeResponse) SetHeaders(v map[string]*string) *UpdateInstanceNetworkTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceNetworkTypeResponse) SetStatusCode(v int32) *UpdateInstanceNetworkTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponse) SetBody(v *UpdateInstanceNetworkTypeResponseBody) *UpdateInstanceNetworkTypeResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("hologram"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetInstanceWithOptions(instanceId *string, request *GetInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstance(instanceId *string, request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances"),
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

func (client *Client) RestartInstanceWithOptions(instanceId *string, request *RestartInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartInstance"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/restart"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartInstance(instanceId *string, request *RestartInstanceRequest) (_result *RestartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartInstanceResponse{}
	_body, _err := client.RestartInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeInstanceWithOptions(instanceId *string, request *ResumeInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ResumeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeInstance"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/resume"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeInstance(instanceId *string, request *ResumeInstanceRequest) (_result *ResumeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeInstanceResponse{}
	_body, _err := client.ResumeInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopInstanceWithOptions(instanceId *string, request *StopInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopInstance"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopInstance(instanceId *string, request *StopInstanceRequest) (_result *StopInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopInstanceResponse{}
	_body, _err := client.StopInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceNameWithOptions(instanceId *string, request *UpdateInstanceNameRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["instanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceName"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/instanceName"),
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

func (client *Client) UpdateInstanceName(instanceId *string, request *UpdateInstanceNameRequest) (_result *UpdateInstanceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceNameResponse{}
	_body, _err := client.UpdateInstanceNameWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceNetworkTypeWithOptions(instanceId *string, request *UpdateInstanceNetworkTypeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceNetworkTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnyTunnelToSingleTunnel)) {
		body["anyTunnelToSingleTunnel"] = request.AnyTunnelToSingleTunnel
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkTypes)) {
		body["networkTypes"] = request.NetworkTypes
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		body["vSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["vpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcOwnerId)) {
		body["vpcOwnerId"] = request.VpcOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcRegionId)) {
		body["vpcRegionId"] = request.VpcRegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceNetworkType"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/network"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceNetworkTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstanceNetworkType(instanceId *string, request *UpdateInstanceNetworkTypeRequest) (_result *UpdateInstanceNetworkTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceNetworkTypeResponse{}
	_body, _err := client.UpdateInstanceNetworkTypeWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
