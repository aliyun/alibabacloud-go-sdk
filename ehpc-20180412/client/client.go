// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddContainerAppRequest struct {
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageTag      *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Repository    *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
}

func (s AddContainerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s AddContainerAppRequest) GoString() string {
	return s.String()
}

func (s *AddContainerAppRequest) SetContainerType(v string) *AddContainerAppRequest {
	s.ContainerType = &v
	return s
}

func (s *AddContainerAppRequest) SetDescription(v string) *AddContainerAppRequest {
	s.Description = &v
	return s
}

func (s *AddContainerAppRequest) SetImageTag(v string) *AddContainerAppRequest {
	s.ImageTag = &v
	return s
}

func (s *AddContainerAppRequest) SetName(v string) *AddContainerAppRequest {
	s.Name = &v
	return s
}

func (s *AddContainerAppRequest) SetRepository(v string) *AddContainerAppRequest {
	s.Repository = &v
	return s
}

type AddContainerAppResponseBody struct {
	ContainerId *AddContainerAppResponseBodyContainerId `json:"ContainerId,omitempty" xml:"ContainerId,omitempty" type:"Struct"`
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddContainerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddContainerAppResponseBody) GoString() string {
	return s.String()
}

func (s *AddContainerAppResponseBody) SetContainerId(v *AddContainerAppResponseBodyContainerId) *AddContainerAppResponseBody {
	s.ContainerId = v
	return s
}

func (s *AddContainerAppResponseBody) SetRequestId(v string) *AddContainerAppResponseBody {
	s.RequestId = &v
	return s
}

type AddContainerAppResponseBodyContainerId struct {
	ContainerId []*string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty" type:"Repeated"`
}

func (s AddContainerAppResponseBodyContainerId) String() string {
	return tea.Prettify(s)
}

func (s AddContainerAppResponseBodyContainerId) GoString() string {
	return s.String()
}

func (s *AddContainerAppResponseBodyContainerId) SetContainerId(v []*string) *AddContainerAppResponseBodyContainerId {
	s.ContainerId = v
	return s
}

type AddContainerAppResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddContainerAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddContainerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s AddContainerAppResponse) GoString() string {
	return s.String()
}

func (s *AddContainerAppResponse) SetHeaders(v map[string]*string) *AddContainerAppResponse {
	s.Headers = v
	return s
}

func (s *AddContainerAppResponse) SetBody(v *AddContainerAppResponseBody) *AddContainerAppResponse {
	s.Body = v
	return s
}

type AddExistedNodesRequest struct {
	ClientToken     *string                           `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterId       *string                           `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ImageId         *string                           `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageOwnerAlias *string                           `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	Instance        []*AddExistedNodesRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
	JobQueue        *string                           `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
}

func (s AddExistedNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddExistedNodesRequest) GoString() string {
	return s.String()
}

func (s *AddExistedNodesRequest) SetClientToken(v string) *AddExistedNodesRequest {
	s.ClientToken = &v
	return s
}

func (s *AddExistedNodesRequest) SetClusterId(v string) *AddExistedNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *AddExistedNodesRequest) SetImageId(v string) *AddExistedNodesRequest {
	s.ImageId = &v
	return s
}

func (s *AddExistedNodesRequest) SetImageOwnerAlias(v string) *AddExistedNodesRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *AddExistedNodesRequest) SetInstance(v []*AddExistedNodesRequestInstance) *AddExistedNodesRequest {
	s.Instance = v
	return s
}

func (s *AddExistedNodesRequest) SetJobQueue(v string) *AddExistedNodesRequest {
	s.JobQueue = &v
	return s
}

type AddExistedNodesRequestInstance struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AddExistedNodesRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s AddExistedNodesRequestInstance) GoString() string {
	return s.String()
}

func (s *AddExistedNodesRequestInstance) SetId(v string) *AddExistedNodesRequestInstance {
	s.Id = &v
	return s
}

type AddExistedNodesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddExistedNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddExistedNodesResponseBody) GoString() string {
	return s.String()
}

func (s *AddExistedNodesResponseBody) SetRequestId(v string) *AddExistedNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddExistedNodesResponseBody) SetTaskId(v string) *AddExistedNodesResponseBody {
	s.TaskId = &v
	return s
}

type AddExistedNodesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddExistedNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddExistedNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddExistedNodesResponse) GoString() string {
	return s.String()
}

func (s *AddExistedNodesResponse) SetHeaders(v map[string]*string) *AddExistedNodesResponse {
	s.Headers = v
	return s
}

func (s *AddExistedNodesResponse) SetBody(v *AddExistedNodesResponseBody) *AddExistedNodesResponse {
	s.Body = v
	return s
}

type AddLocalNodesRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Nodes     *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s AddLocalNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLocalNodesRequest) GoString() string {
	return s.String()
}

func (s *AddLocalNodesRequest) SetClusterId(v string) *AddLocalNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *AddLocalNodesRequest) SetNodes(v string) *AddLocalNodesRequest {
	s.Nodes = &v
	return s
}

type AddLocalNodesResponseBody struct {
	InstanceIds *AddLocalNodesResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLocalNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddLocalNodesResponseBody) GoString() string {
	return s.String()
}

func (s *AddLocalNodesResponseBody) SetInstanceIds(v *AddLocalNodesResponseBodyInstanceIds) *AddLocalNodesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *AddLocalNodesResponseBody) SetRequestId(v string) *AddLocalNodesResponseBody {
	s.RequestId = &v
	return s
}

type AddLocalNodesResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s AddLocalNodesResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s AddLocalNodesResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *AddLocalNodesResponseBodyInstanceIds) SetInstanceId(v []*string) *AddLocalNodesResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

type AddLocalNodesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddLocalNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddLocalNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLocalNodesResponse) GoString() string {
	return s.String()
}

func (s *AddLocalNodesResponse) SetHeaders(v map[string]*string) *AddLocalNodesResponse {
	s.Headers = v
	return s
}

func (s *AddLocalNodesResponse) SetBody(v *AddLocalNodesResponseBody) *AddLocalNodesResponse {
	s.Body = v
	return s
}

type AddNodesRequest struct {
	AllocatePublicAddress   *bool                       `json:"AllocatePublicAddress,omitempty" xml:"AllocatePublicAddress,omitempty"`
	AutoRenew               *string                     `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod         *int32                      `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	ClientToken             *string                     `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterId               *string                     `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ComputeEnableHt         *bool                       `json:"ComputeEnableHt,omitempty" xml:"ComputeEnableHt,omitempty"`
	ComputeSpotPriceLimit   *string                     `json:"ComputeSpotPriceLimit,omitempty" xml:"ComputeSpotPriceLimit,omitempty"`
	ComputeSpotStrategy     *string                     `json:"ComputeSpotStrategy,omitempty" xml:"ComputeSpotStrategy,omitempty"`
	Count                   *int32                      `json:"Count,omitempty" xml:"Count,omitempty"`
	CreateMode              *string                     `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	DataDisks               []*AddNodesRequestDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	EcsChargeType           *string                     `json:"EcsChargeType,omitempty" xml:"EcsChargeType,omitempty"`
	HostNamePrefix          *string                     `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	HostNameSuffix          *string                     `json:"HostNameSuffix,omitempty" xml:"HostNameSuffix,omitempty"`
	ImageId                 *string                     `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageOwnerAlias         *string                     `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	InstanceType            *string                     `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetChargeType      *string                     `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandWidthIn  *int32                      `json:"InternetMaxBandWidthIn,omitempty" xml:"InternetMaxBandWidthIn,omitempty"`
	InternetMaxBandWidthOut *int32                      `json:"InternetMaxBandWidthOut,omitempty" xml:"InternetMaxBandWidthOut,omitempty"`
	JobQueue                *string                     `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	MinCount                *int32                      `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	Period                  *int32                      `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit              *string                     `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	Sync                    *bool                       `json:"Sync,omitempty" xml:"Sync,omitempty"`
	SystemDiskLevel         *string                     `json:"SystemDiskLevel,omitempty" xml:"SystemDiskLevel,omitempty"`
	SystemDiskSize          *int32                      `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	SystemDiskType          *string                     `json:"SystemDiskType,omitempty" xml:"SystemDiskType,omitempty"`
	VSwitchId               *string                     `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                  *string                     `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AddNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddNodesRequest) GoString() string {
	return s.String()
}

func (s *AddNodesRequest) SetAllocatePublicAddress(v bool) *AddNodesRequest {
	s.AllocatePublicAddress = &v
	return s
}

func (s *AddNodesRequest) SetAutoRenew(v string) *AddNodesRequest {
	s.AutoRenew = &v
	return s
}

func (s *AddNodesRequest) SetAutoRenewPeriod(v int32) *AddNodesRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *AddNodesRequest) SetClientToken(v string) *AddNodesRequest {
	s.ClientToken = &v
	return s
}

func (s *AddNodesRequest) SetClusterId(v string) *AddNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *AddNodesRequest) SetComputeEnableHt(v bool) *AddNodesRequest {
	s.ComputeEnableHt = &v
	return s
}

func (s *AddNodesRequest) SetComputeSpotPriceLimit(v string) *AddNodesRequest {
	s.ComputeSpotPriceLimit = &v
	return s
}

func (s *AddNodesRequest) SetComputeSpotStrategy(v string) *AddNodesRequest {
	s.ComputeSpotStrategy = &v
	return s
}

func (s *AddNodesRequest) SetCount(v int32) *AddNodesRequest {
	s.Count = &v
	return s
}

func (s *AddNodesRequest) SetCreateMode(v string) *AddNodesRequest {
	s.CreateMode = &v
	return s
}

func (s *AddNodesRequest) SetDataDisks(v []*AddNodesRequestDataDisks) *AddNodesRequest {
	s.DataDisks = v
	return s
}

func (s *AddNodesRequest) SetEcsChargeType(v string) *AddNodesRequest {
	s.EcsChargeType = &v
	return s
}

func (s *AddNodesRequest) SetHostNamePrefix(v string) *AddNodesRequest {
	s.HostNamePrefix = &v
	return s
}

func (s *AddNodesRequest) SetHostNameSuffix(v string) *AddNodesRequest {
	s.HostNameSuffix = &v
	return s
}

func (s *AddNodesRequest) SetImageId(v string) *AddNodesRequest {
	s.ImageId = &v
	return s
}

func (s *AddNodesRequest) SetImageOwnerAlias(v string) *AddNodesRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *AddNodesRequest) SetInstanceType(v string) *AddNodesRequest {
	s.InstanceType = &v
	return s
}

func (s *AddNodesRequest) SetInternetChargeType(v string) *AddNodesRequest {
	s.InternetChargeType = &v
	return s
}

func (s *AddNodesRequest) SetInternetMaxBandWidthIn(v int32) *AddNodesRequest {
	s.InternetMaxBandWidthIn = &v
	return s
}

func (s *AddNodesRequest) SetInternetMaxBandWidthOut(v int32) *AddNodesRequest {
	s.InternetMaxBandWidthOut = &v
	return s
}

func (s *AddNodesRequest) SetJobQueue(v string) *AddNodesRequest {
	s.JobQueue = &v
	return s
}

func (s *AddNodesRequest) SetMinCount(v int32) *AddNodesRequest {
	s.MinCount = &v
	return s
}

func (s *AddNodesRequest) SetPeriod(v int32) *AddNodesRequest {
	s.Period = &v
	return s
}

func (s *AddNodesRequest) SetPeriodUnit(v string) *AddNodesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *AddNodesRequest) SetSync(v bool) *AddNodesRequest {
	s.Sync = &v
	return s
}

func (s *AddNodesRequest) SetSystemDiskLevel(v string) *AddNodesRequest {
	s.SystemDiskLevel = &v
	return s
}

func (s *AddNodesRequest) SetSystemDiskSize(v int32) *AddNodesRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *AddNodesRequest) SetSystemDiskType(v string) *AddNodesRequest {
	s.SystemDiskType = &v
	return s
}

func (s *AddNodesRequest) SetVSwitchId(v string) *AddNodesRequest {
	s.VSwitchId = &v
	return s
}

func (s *AddNodesRequest) SetZoneId(v string) *AddNodesRequest {
	s.ZoneId = &v
	return s
}

type AddNodesRequestDataDisks struct {
	DataDiskCategory           *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	DataDiskDeleteWithInstance *bool   `json:"DataDiskDeleteWithInstance,omitempty" xml:"DataDiskDeleteWithInstance,omitempty"`
	DataDiskEncrypted          *bool   `json:"DataDiskEncrypted,omitempty" xml:"DataDiskEncrypted,omitempty"`
	DataDiskKMSKeyId           *string `json:"DataDiskKMSKeyId,omitempty" xml:"DataDiskKMSKeyId,omitempty"`
	DataDiskPerformanceLevel   *string `json:"DataDiskPerformanceLevel,omitempty" xml:"DataDiskPerformanceLevel,omitempty"`
	DataDiskSize               *int32  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
}

func (s AddNodesRequestDataDisks) String() string {
	return tea.Prettify(s)
}

func (s AddNodesRequestDataDisks) GoString() string {
	return s.String()
}

func (s *AddNodesRequestDataDisks) SetDataDiskCategory(v string) *AddNodesRequestDataDisks {
	s.DataDiskCategory = &v
	return s
}

func (s *AddNodesRequestDataDisks) SetDataDiskDeleteWithInstance(v bool) *AddNodesRequestDataDisks {
	s.DataDiskDeleteWithInstance = &v
	return s
}

func (s *AddNodesRequestDataDisks) SetDataDiskEncrypted(v bool) *AddNodesRequestDataDisks {
	s.DataDiskEncrypted = &v
	return s
}

func (s *AddNodesRequestDataDisks) SetDataDiskKMSKeyId(v string) *AddNodesRequestDataDisks {
	s.DataDiskKMSKeyId = &v
	return s
}

func (s *AddNodesRequestDataDisks) SetDataDiskPerformanceLevel(v string) *AddNodesRequestDataDisks {
	s.DataDiskPerformanceLevel = &v
	return s
}

func (s *AddNodesRequestDataDisks) SetDataDiskSize(v int32) *AddNodesRequestDataDisks {
	s.DataDiskSize = &v
	return s
}

type AddNodesResponseBody struct {
	InstanceIds *AddNodesResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	RequestId   *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId      *string                          `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddNodesResponseBody) GoString() string {
	return s.String()
}

func (s *AddNodesResponseBody) SetInstanceIds(v *AddNodesResponseBodyInstanceIds) *AddNodesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *AddNodesResponseBody) SetRequestId(v string) *AddNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddNodesResponseBody) SetTaskId(v string) *AddNodesResponseBody {
	s.TaskId = &v
	return s
}

type AddNodesResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s AddNodesResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s AddNodesResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *AddNodesResponseBodyInstanceIds) SetInstanceId(v []*string) *AddNodesResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

type AddNodesResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddNodesResponse) GoString() string {
	return s.String()
}

func (s *AddNodesResponse) SetHeaders(v map[string]*string) *AddNodesResponse {
	s.Headers = v
	return s
}

func (s *AddNodesResponse) SetBody(v *AddNodesResponseBody) *AddNodesResponse {
	s.Body = v
	return s
}

type AddQueueRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s AddQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s AddQueueRequest) GoString() string {
	return s.String()
}

func (s *AddQueueRequest) SetClusterId(v string) *AddQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *AddQueueRequest) SetQueueName(v string) *AddQueueRequest {
	s.QueueName = &v
	return s
}

type AddQueueResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddQueueResponseBody) GoString() string {
	return s.String()
}

func (s *AddQueueResponseBody) SetRequestId(v string) *AddQueueResponseBody {
	s.RequestId = &v
	return s
}

type AddQueueResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s AddQueueResponse) GoString() string {
	return s.String()
}

func (s *AddQueueResponse) SetHeaders(v map[string]*string) *AddQueueResponse {
	s.Headers = v
	return s
}

func (s *AddQueueResponse) SetBody(v *AddQueueResponseBody) *AddQueueResponse {
	s.Body = v
	return s
}

type AddSecurityGroupRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s AddSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *AddSecurityGroupRequest) SetClientToken(v string) *AddSecurityGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AddSecurityGroupRequest) SetClusterId(v string) *AddSecurityGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *AddSecurityGroupRequest) SetSecurityGroupId(v string) *AddSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

type AddSecurityGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSecurityGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddSecurityGroupResponseBody) SetRequestId(v string) *AddSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

type AddSecurityGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *AddSecurityGroupResponse) SetHeaders(v map[string]*string) *AddSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *AddSecurityGroupResponse) SetBody(v *AddSecurityGroupResponseBody) *AddSecurityGroupResponse {
	s.Body = v
	return s
}

type AddUsersRequest struct {
	ClusterId *string                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	User      []*AddUsersRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s AddUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUsersRequest) GoString() string {
	return s.String()
}

func (s *AddUsersRequest) SetClusterId(v string) *AddUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *AddUsersRequest) SetUser(v []*AddUsersRequestUser) *AddUsersRequest {
	s.User = v
	return s
}

type AddUsersRequestUser struct {
	Group    *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s AddUsersRequestUser) String() string {
	return tea.Prettify(s)
}

func (s AddUsersRequestUser) GoString() string {
	return s.String()
}

func (s *AddUsersRequestUser) SetGroup(v string) *AddUsersRequestUser {
	s.Group = &v
	return s
}

func (s *AddUsersRequestUser) SetName(v string) *AddUsersRequestUser {
	s.Name = &v
	return s
}

func (s *AddUsersRequestUser) SetPassword(v string) *AddUsersRequestUser {
	s.Password = &v
	return s
}

type AddUsersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUsersResponseBody) GoString() string {
	return s.String()
}

func (s *AddUsersResponseBody) SetRequestId(v string) *AddUsersResponseBody {
	s.RequestId = &v
	return s
}

type AddUsersResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUsersResponse) GoString() string {
	return s.String()
}

func (s *AddUsersResponse) SetHeaders(v map[string]*string) *AddUsersResponse {
	s.Headers = v
	return s
}

func (s *AddUsersResponse) SetBody(v *AddUsersResponseBody) *AddUsersResponse {
	s.Body = v
	return s
}

type ApplyNodesRequest struct {
	AllocatePublicAddress         *bool                                 `json:"AllocatePublicAddress,omitempty" xml:"AllocatePublicAddress,omitempty"`
	ClusterId                     *string                               `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ComputeSpotPriceLimit         *float32                              `json:"ComputeSpotPriceLimit,omitempty" xml:"ComputeSpotPriceLimit,omitempty"`
	ComputeSpotStrategy           *string                               `json:"ComputeSpotStrategy,omitempty" xml:"ComputeSpotStrategy,omitempty"`
	Cores                         *int32                                `json:"Cores,omitempty" xml:"Cores,omitempty"`
	HostNamePrefix                *string                               `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	HostNameSuffix                *string                               `json:"HostNameSuffix,omitempty" xml:"HostNameSuffix,omitempty"`
	ImageId                       *string                               `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceFamilyLevel           *string                               `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	InstanceTypeModel             []*ApplyNodesRequestInstanceTypeModel `json:"InstanceTypeModel,omitempty" xml:"InstanceTypeModel,omitempty" type:"Repeated"`
	InternetChargeType            *string                               `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandWidthIn        *int32                                `json:"InternetMaxBandWidthIn,omitempty" xml:"InternetMaxBandWidthIn,omitempty"`
	InternetMaxBandWidthOut       *int32                                `json:"InternetMaxBandWidthOut,omitempty" xml:"InternetMaxBandWidthOut,omitempty"`
	Interval                      *int32                                `json:"Interval,omitempty" xml:"Interval,omitempty"`
	JobQueue                      *string                               `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	Memory                        *int32                                `json:"Memory,omitempty" xml:"Memory,omitempty"`
	PriorityStrategy              *string                               `json:"PriorityStrategy,omitempty" xml:"PriorityStrategy,omitempty"`
	ResourceAmountType            *string                               `json:"ResourceAmountType,omitempty" xml:"ResourceAmountType,omitempty"`
	Round                         *int32                                `json:"Round,omitempty" xml:"Round,omitempty"`
	StrictResourceProvision       *bool                                 `json:"StrictResourceProvision,omitempty" xml:"StrictResourceProvision,omitempty"`
	StrictSatisfiedTargetCapacity *bool                                 `json:"StrictSatisfiedTargetCapacity,omitempty" xml:"StrictSatisfiedTargetCapacity,omitempty"`
	SystemDiskLevel               *string                               `json:"SystemDiskLevel,omitempty" xml:"SystemDiskLevel,omitempty"`
	SystemDiskSize                *int32                                `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	SystemDiskType                *string                               `json:"SystemDiskType,omitempty" xml:"SystemDiskType,omitempty"`
	Tag                           []*ApplyNodesRequestTag               `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TargetCapacity                *int32                                `json:"TargetCapacity,omitempty" xml:"TargetCapacity,omitempty"`
	ZoneInfos                     []*ApplyNodesRequestZoneInfos         `json:"ZoneInfos,omitempty" xml:"ZoneInfos,omitempty" type:"Repeated"`
}

func (s ApplyNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesRequest) GoString() string {
	return s.String()
}

func (s *ApplyNodesRequest) SetAllocatePublicAddress(v bool) *ApplyNodesRequest {
	s.AllocatePublicAddress = &v
	return s
}

func (s *ApplyNodesRequest) SetClusterId(v string) *ApplyNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ApplyNodesRequest) SetComputeSpotPriceLimit(v float32) *ApplyNodesRequest {
	s.ComputeSpotPriceLimit = &v
	return s
}

func (s *ApplyNodesRequest) SetComputeSpotStrategy(v string) *ApplyNodesRequest {
	s.ComputeSpotStrategy = &v
	return s
}

func (s *ApplyNodesRequest) SetCores(v int32) *ApplyNodesRequest {
	s.Cores = &v
	return s
}

func (s *ApplyNodesRequest) SetHostNamePrefix(v string) *ApplyNodesRequest {
	s.HostNamePrefix = &v
	return s
}

func (s *ApplyNodesRequest) SetHostNameSuffix(v string) *ApplyNodesRequest {
	s.HostNameSuffix = &v
	return s
}

func (s *ApplyNodesRequest) SetImageId(v string) *ApplyNodesRequest {
	s.ImageId = &v
	return s
}

func (s *ApplyNodesRequest) SetInstanceFamilyLevel(v string) *ApplyNodesRequest {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *ApplyNodesRequest) SetInstanceTypeModel(v []*ApplyNodesRequestInstanceTypeModel) *ApplyNodesRequest {
	s.InstanceTypeModel = v
	return s
}

func (s *ApplyNodesRequest) SetInternetChargeType(v string) *ApplyNodesRequest {
	s.InternetChargeType = &v
	return s
}

func (s *ApplyNodesRequest) SetInternetMaxBandWidthIn(v int32) *ApplyNodesRequest {
	s.InternetMaxBandWidthIn = &v
	return s
}

func (s *ApplyNodesRequest) SetInternetMaxBandWidthOut(v int32) *ApplyNodesRequest {
	s.InternetMaxBandWidthOut = &v
	return s
}

func (s *ApplyNodesRequest) SetInterval(v int32) *ApplyNodesRequest {
	s.Interval = &v
	return s
}

func (s *ApplyNodesRequest) SetJobQueue(v string) *ApplyNodesRequest {
	s.JobQueue = &v
	return s
}

func (s *ApplyNodesRequest) SetMemory(v int32) *ApplyNodesRequest {
	s.Memory = &v
	return s
}

func (s *ApplyNodesRequest) SetPriorityStrategy(v string) *ApplyNodesRequest {
	s.PriorityStrategy = &v
	return s
}

func (s *ApplyNodesRequest) SetResourceAmountType(v string) *ApplyNodesRequest {
	s.ResourceAmountType = &v
	return s
}

func (s *ApplyNodesRequest) SetRound(v int32) *ApplyNodesRequest {
	s.Round = &v
	return s
}

func (s *ApplyNodesRequest) SetStrictResourceProvision(v bool) *ApplyNodesRequest {
	s.StrictResourceProvision = &v
	return s
}

func (s *ApplyNodesRequest) SetStrictSatisfiedTargetCapacity(v bool) *ApplyNodesRequest {
	s.StrictSatisfiedTargetCapacity = &v
	return s
}

func (s *ApplyNodesRequest) SetSystemDiskLevel(v string) *ApplyNodesRequest {
	s.SystemDiskLevel = &v
	return s
}

func (s *ApplyNodesRequest) SetSystemDiskSize(v int32) *ApplyNodesRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *ApplyNodesRequest) SetSystemDiskType(v string) *ApplyNodesRequest {
	s.SystemDiskType = &v
	return s
}

func (s *ApplyNodesRequest) SetTag(v []*ApplyNodesRequestTag) *ApplyNodesRequest {
	s.Tag = v
	return s
}

func (s *ApplyNodesRequest) SetTargetCapacity(v int32) *ApplyNodesRequest {
	s.TargetCapacity = &v
	return s
}

func (s *ApplyNodesRequest) SetZoneInfos(v []*ApplyNodesRequestZoneInfos) *ApplyNodesRequest {
	s.ZoneInfos = v
	return s
}

type ApplyNodesRequestInstanceTypeModel struct {
	InstanceType  *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MaxPrice      *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	TargetImageId *string  `json:"TargetImageId,omitempty" xml:"TargetImageId,omitempty"`
}

func (s ApplyNodesRequestInstanceTypeModel) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesRequestInstanceTypeModel) GoString() string {
	return s.String()
}

func (s *ApplyNodesRequestInstanceTypeModel) SetInstanceType(v string) *ApplyNodesRequestInstanceTypeModel {
	s.InstanceType = &v
	return s
}

func (s *ApplyNodesRequestInstanceTypeModel) SetMaxPrice(v float32) *ApplyNodesRequestInstanceTypeModel {
	s.MaxPrice = &v
	return s
}

func (s *ApplyNodesRequestInstanceTypeModel) SetTargetImageId(v string) *ApplyNodesRequestInstanceTypeModel {
	s.TargetImageId = &v
	return s
}

type ApplyNodesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ApplyNodesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesRequestTag) GoString() string {
	return s.String()
}

func (s *ApplyNodesRequestTag) SetKey(v string) *ApplyNodesRequestTag {
	s.Key = &v
	return s
}

func (s *ApplyNodesRequestTag) SetValue(v string) *ApplyNodesRequestTag {
	s.Value = &v
	return s
}

type ApplyNodesRequestZoneInfos struct {
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ApplyNodesRequestZoneInfos) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesRequestZoneInfos) GoString() string {
	return s.String()
}

func (s *ApplyNodesRequestZoneInfos) SetVSwitchId(v string) *ApplyNodesRequestZoneInfos {
	s.VSwitchId = &v
	return s
}

func (s *ApplyNodesRequestZoneInfos) SetZoneId(v string) *ApplyNodesRequestZoneInfos {
	s.ZoneId = &v
	return s
}

type ApplyNodesResponseBody struct {
	Detail          *string                            `json:"Detail,omitempty" xml:"Detail,omitempty"`
	InstanceIds     *ApplyNodesResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	RequestId       *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SatisfiedAmount *int32                             `json:"SatisfiedAmount,omitempty" xml:"SatisfiedAmount,omitempty"`
	TaskId          *string                            `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ApplyNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyNodesResponseBody) SetDetail(v string) *ApplyNodesResponseBody {
	s.Detail = &v
	return s
}

func (s *ApplyNodesResponseBody) SetInstanceIds(v *ApplyNodesResponseBodyInstanceIds) *ApplyNodesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *ApplyNodesResponseBody) SetRequestId(v string) *ApplyNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyNodesResponseBody) SetSatisfiedAmount(v int32) *ApplyNodesResponseBody {
	s.SatisfiedAmount = &v
	return s
}

func (s *ApplyNodesResponseBody) SetTaskId(v string) *ApplyNodesResponseBody {
	s.TaskId = &v
	return s
}

type ApplyNodesResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s ApplyNodesResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *ApplyNodesResponseBodyInstanceIds) SetInstanceId(v []*string) *ApplyNodesResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

type ApplyNodesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplyNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesResponse) GoString() string {
	return s.String()
}

func (s *ApplyNodesResponse) SetHeaders(v map[string]*string) *ApplyNodesResponse {
	s.Headers = v
	return s
}

func (s *ApplyNodesResponse) SetBody(v *ApplyNodesResponseBody) *ApplyNodesResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	EcsOrder              *CreateClusterRequestEcsOrder            `json:"EcsOrder,omitempty" xml:"EcsOrder,omitempty" type:"Struct"`
	AccountType           *string                                  `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AdditionalVolumes     []*CreateClusterRequestAdditionalVolumes `json:"AdditionalVolumes,omitempty" xml:"AdditionalVolumes,omitempty" type:"Repeated"`
	Application           []*CreateClusterRequestApplication       `json:"Application,omitempty" xml:"Application,omitempty" type:"Repeated"`
	AutoRenew             *string                                  `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod       *int32                                   `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	ClientToken           *string                                  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClientVersion         *string                                  `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ClusterVersion        *string                                  `json:"ClusterVersion,omitempty" xml:"ClusterVersion,omitempty"`
	ComputeEnableHt       *bool                                    `json:"ComputeEnableHt,omitempty" xml:"ComputeEnableHt,omitempty"`
	ComputeSpotPriceLimit *string                                  `json:"ComputeSpotPriceLimit,omitempty" xml:"ComputeSpotPriceLimit,omitempty"`
	ComputeSpotStrategy   *string                                  `json:"ComputeSpotStrategy,omitempty" xml:"ComputeSpotStrategy,omitempty"`
	DeployMode            *string                                  `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	Description           *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	EcsChargeType         *string                                  `json:"EcsChargeType,omitempty" xml:"EcsChargeType,omitempty"`
	EhpcVersion           *string                                  `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	HaEnable              *bool                                    `json:"HaEnable,omitempty" xml:"HaEnable,omitempty"`
	ImageId               *string                                  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageOwnerAlias       *string                                  `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	InputFileUrl          *string                                  `json:"InputFileUrl,omitempty" xml:"InputFileUrl,omitempty"`
	IsComputeEss          *bool                                    `json:"IsComputeEss,omitempty" xml:"IsComputeEss,omitempty"`
	JobQueue              *string                                  `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	KeyPairName           *string                                  `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	Name                  *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	OsTag                 *string                                  `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	Password              *string                                  `json:"Password,omitempty" xml:"Password,omitempty"`
	Period                *int32                                   `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit            *string                                  `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	Plugin                *string                                  `json:"Plugin,omitempty" xml:"Plugin,omitempty"`
	PostInstallScript     []*CreateClusterRequestPostInstallScript `json:"PostInstallScript,omitempty" xml:"PostInstallScript,omitempty" type:"Repeated"`
	RamNodeTypes          []*string                                `json:"RamNodeTypes,omitempty" xml:"RamNodeTypes,omitempty" type:"Repeated"`
	RamRoleName           *string                                  `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	RemoteDirectory       *string                                  `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	RemoteVisEnable       *string                                  `json:"RemoteVisEnable,omitempty" xml:"RemoteVisEnable,omitempty"`
	ResourceGroupId       *string                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SccClusterId          *string                                  `json:"SccClusterId,omitempty" xml:"SccClusterId,omitempty"`
	SchedulerType         *string                                  `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	SecurityGroupId       *string                                  `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupName     *string                                  `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	SystemDiskLevel       *string                                  `json:"SystemDiskLevel,omitempty" xml:"SystemDiskLevel,omitempty"`
	SystemDiskSize        *int32                                   `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	SystemDiskType        *string                                  `json:"SystemDiskType,omitempty" xml:"SystemDiskType,omitempty"`
	VSwitchId             *string                                  `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VolumeId              *string                                  `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	VolumeMountpoint      *string                                  `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	VolumeProtocol        *string                                  `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	VolumeType            *string                                  `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
	VpcId                 *string                                  `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	WithoutAgent          *bool                                    `json:"WithoutAgent,omitempty" xml:"WithoutAgent,omitempty"`
	WithoutElasticIp      *bool                                    `json:"WithoutElasticIp,omitempty" xml:"WithoutElasticIp,omitempty"`
	ZoneId                *string                                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetEcsOrder(v *CreateClusterRequestEcsOrder) *CreateClusterRequest {
	s.EcsOrder = v
	return s
}

func (s *CreateClusterRequest) SetAccountType(v string) *CreateClusterRequest {
	s.AccountType = &v
	return s
}

func (s *CreateClusterRequest) SetAdditionalVolumes(v []*CreateClusterRequestAdditionalVolumes) *CreateClusterRequest {
	s.AdditionalVolumes = v
	return s
}

func (s *CreateClusterRequest) SetApplication(v []*CreateClusterRequestApplication) *CreateClusterRequest {
	s.Application = v
	return s
}

func (s *CreateClusterRequest) SetAutoRenew(v string) *CreateClusterRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateClusterRequest) SetAutoRenewPeriod(v int32) *CreateClusterRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateClusterRequest) SetClientToken(v string) *CreateClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateClusterRequest) SetClientVersion(v string) *CreateClusterRequest {
	s.ClientVersion = &v
	return s
}

func (s *CreateClusterRequest) SetClusterVersion(v string) *CreateClusterRequest {
	s.ClusterVersion = &v
	return s
}

func (s *CreateClusterRequest) SetComputeEnableHt(v bool) *CreateClusterRequest {
	s.ComputeEnableHt = &v
	return s
}

func (s *CreateClusterRequest) SetComputeSpotPriceLimit(v string) *CreateClusterRequest {
	s.ComputeSpotPriceLimit = &v
	return s
}

func (s *CreateClusterRequest) SetComputeSpotStrategy(v string) *CreateClusterRequest {
	s.ComputeSpotStrategy = &v
	return s
}

func (s *CreateClusterRequest) SetDeployMode(v string) *CreateClusterRequest {
	s.DeployMode = &v
	return s
}

func (s *CreateClusterRequest) SetDescription(v string) *CreateClusterRequest {
	s.Description = &v
	return s
}

func (s *CreateClusterRequest) SetEcsChargeType(v string) *CreateClusterRequest {
	s.EcsChargeType = &v
	return s
}

func (s *CreateClusterRequest) SetEhpcVersion(v string) *CreateClusterRequest {
	s.EhpcVersion = &v
	return s
}

func (s *CreateClusterRequest) SetHaEnable(v bool) *CreateClusterRequest {
	s.HaEnable = &v
	return s
}

func (s *CreateClusterRequest) SetImageId(v string) *CreateClusterRequest {
	s.ImageId = &v
	return s
}

func (s *CreateClusterRequest) SetImageOwnerAlias(v string) *CreateClusterRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *CreateClusterRequest) SetInputFileUrl(v string) *CreateClusterRequest {
	s.InputFileUrl = &v
	return s
}

func (s *CreateClusterRequest) SetIsComputeEss(v bool) *CreateClusterRequest {
	s.IsComputeEss = &v
	return s
}

func (s *CreateClusterRequest) SetJobQueue(v string) *CreateClusterRequest {
	s.JobQueue = &v
	return s
}

func (s *CreateClusterRequest) SetKeyPairName(v string) *CreateClusterRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateClusterRequest) SetName(v string) *CreateClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateClusterRequest) SetOsTag(v string) *CreateClusterRequest {
	s.OsTag = &v
	return s
}

func (s *CreateClusterRequest) SetPassword(v string) *CreateClusterRequest {
	s.Password = &v
	return s
}

func (s *CreateClusterRequest) SetPeriod(v int32) *CreateClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateClusterRequest) SetPeriodUnit(v string) *CreateClusterRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateClusterRequest) SetPlugin(v string) *CreateClusterRequest {
	s.Plugin = &v
	return s
}

func (s *CreateClusterRequest) SetPostInstallScript(v []*CreateClusterRequestPostInstallScript) *CreateClusterRequest {
	s.PostInstallScript = v
	return s
}

func (s *CreateClusterRequest) SetRamNodeTypes(v []*string) *CreateClusterRequest {
	s.RamNodeTypes = v
	return s
}

func (s *CreateClusterRequest) SetRamRoleName(v string) *CreateClusterRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateClusterRequest) SetRemoteDirectory(v string) *CreateClusterRequest {
	s.RemoteDirectory = &v
	return s
}

func (s *CreateClusterRequest) SetRemoteVisEnable(v string) *CreateClusterRequest {
	s.RemoteVisEnable = &v
	return s
}

func (s *CreateClusterRequest) SetResourceGroupId(v string) *CreateClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetSccClusterId(v string) *CreateClusterRequest {
	s.SccClusterId = &v
	return s
}

func (s *CreateClusterRequest) SetSchedulerType(v string) *CreateClusterRequest {
	s.SchedulerType = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityGroupId(v string) *CreateClusterRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityGroupName(v string) *CreateClusterRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *CreateClusterRequest) SetSystemDiskLevel(v string) *CreateClusterRequest {
	s.SystemDiskLevel = &v
	return s
}

func (s *CreateClusterRequest) SetSystemDiskSize(v int32) *CreateClusterRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateClusterRequest) SetSystemDiskType(v string) *CreateClusterRequest {
	s.SystemDiskType = &v
	return s
}

func (s *CreateClusterRequest) SetVSwitchId(v string) *CreateClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequest) SetVolumeId(v string) *CreateClusterRequest {
	s.VolumeId = &v
	return s
}

func (s *CreateClusterRequest) SetVolumeMountpoint(v string) *CreateClusterRequest {
	s.VolumeMountpoint = &v
	return s
}

func (s *CreateClusterRequest) SetVolumeProtocol(v string) *CreateClusterRequest {
	s.VolumeProtocol = &v
	return s
}

func (s *CreateClusterRequest) SetVolumeType(v string) *CreateClusterRequest {
	s.VolumeType = &v
	return s
}

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequest) SetWithoutAgent(v bool) *CreateClusterRequest {
	s.WithoutAgent = &v
	return s
}

func (s *CreateClusterRequest) SetWithoutElasticIp(v bool) *CreateClusterRequest {
	s.WithoutElasticIp = &v
	return s
}

func (s *CreateClusterRequest) SetZoneId(v string) *CreateClusterRequest {
	s.ZoneId = &v
	return s
}

type CreateClusterRequestEcsOrder struct {
	Compute *CreateClusterRequestEcsOrderCompute `json:"Compute,omitempty" xml:"Compute,omitempty" require:"true" type:"Struct"`
	Login   *CreateClusterRequestEcsOrderLogin   `json:"Login,omitempty" xml:"Login,omitempty" require:"true" type:"Struct"`
	Manager *CreateClusterRequestEcsOrderManager `json:"Manager,omitempty" xml:"Manager,omitempty" require:"true" type:"Struct"`
}

func (s CreateClusterRequestEcsOrder) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestEcsOrder) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestEcsOrder) SetCompute(v *CreateClusterRequestEcsOrderCompute) *CreateClusterRequestEcsOrder {
	s.Compute = v
	return s
}

func (s *CreateClusterRequestEcsOrder) SetLogin(v *CreateClusterRequestEcsOrderLogin) *CreateClusterRequestEcsOrder {
	s.Login = v
	return s
}

func (s *CreateClusterRequestEcsOrder) SetManager(v *CreateClusterRequestEcsOrderManager) *CreateClusterRequestEcsOrder {
	s.Manager = v
	return s
}

type CreateClusterRequestEcsOrderCompute struct {
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateClusterRequestEcsOrderCompute) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestEcsOrderCompute) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestEcsOrderCompute) SetCount(v int32) *CreateClusterRequestEcsOrderCompute {
	s.Count = &v
	return s
}

func (s *CreateClusterRequestEcsOrderCompute) SetInstanceType(v string) *CreateClusterRequestEcsOrderCompute {
	s.InstanceType = &v
	return s
}

type CreateClusterRequestEcsOrderLogin struct {
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateClusterRequestEcsOrderLogin) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestEcsOrderLogin) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestEcsOrderLogin) SetCount(v int32) *CreateClusterRequestEcsOrderLogin {
	s.Count = &v
	return s
}

func (s *CreateClusterRequestEcsOrderLogin) SetInstanceType(v string) *CreateClusterRequestEcsOrderLogin {
	s.InstanceType = &v
	return s
}

type CreateClusterRequestEcsOrderManager struct {
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateClusterRequestEcsOrderManager) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestEcsOrderManager) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestEcsOrderManager) SetCount(v int32) *CreateClusterRequestEcsOrderManager {
	s.Count = &v
	return s
}

func (s *CreateClusterRequestEcsOrderManager) SetInstanceType(v string) *CreateClusterRequestEcsOrderManager {
	s.InstanceType = &v
	return s
}

type CreateClusterRequestAdditionalVolumes struct {
	JobQueue         *string                                       `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	LocalDirectory   *string                                       `json:"LocalDirectory,omitempty" xml:"LocalDirectory,omitempty"`
	Location         *string                                       `json:"Location,omitempty" xml:"Location,omitempty"`
	RemoteDirectory  *string                                       `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	Roles            []*CreateClusterRequestAdditionalVolumesRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	VolumeId         *string                                       `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	VolumeMountpoint *string                                       `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	VolumeProtocol   *string                                       `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	VolumeType       *string                                       `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
}

func (s CreateClusterRequestAdditionalVolumes) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestAdditionalVolumes) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestAdditionalVolumes) SetJobQueue(v string) *CreateClusterRequestAdditionalVolumes {
	s.JobQueue = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetLocalDirectory(v string) *CreateClusterRequestAdditionalVolumes {
	s.LocalDirectory = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetLocation(v string) *CreateClusterRequestAdditionalVolumes {
	s.Location = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetRemoteDirectory(v string) *CreateClusterRequestAdditionalVolumes {
	s.RemoteDirectory = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetRoles(v []*CreateClusterRequestAdditionalVolumesRoles) *CreateClusterRequestAdditionalVolumes {
	s.Roles = v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetVolumeId(v string) *CreateClusterRequestAdditionalVolumes {
	s.VolumeId = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetVolumeMountpoint(v string) *CreateClusterRequestAdditionalVolumes {
	s.VolumeMountpoint = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetVolumeProtocol(v string) *CreateClusterRequestAdditionalVolumes {
	s.VolumeProtocol = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetVolumeType(v string) *CreateClusterRequestAdditionalVolumes {
	s.VolumeType = &v
	return s
}

type CreateClusterRequestAdditionalVolumesRoles struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateClusterRequestAdditionalVolumesRoles) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestAdditionalVolumesRoles) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestAdditionalVolumesRoles) SetName(v string) *CreateClusterRequestAdditionalVolumesRoles {
	s.Name = &v
	return s
}

type CreateClusterRequestApplication struct {
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateClusterRequestApplication) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestApplication) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestApplication) SetTag(v string) *CreateClusterRequestApplication {
	s.Tag = &v
	return s
}

type CreateClusterRequestPostInstallScript struct {
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateClusterRequestPostInstallScript) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestPostInstallScript) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestPostInstallScript) SetArgs(v string) *CreateClusterRequestPostInstallScript {
	s.Args = &v
	return s
}

func (s *CreateClusterRequestPostInstallScript) SetUrl(v string) *CreateClusterRequestPostInstallScript {
	s.Url = &v
	return s
}

type CreateClusterResponseBody struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetTaskId(v string) *CreateClusterResponseBody {
	s.TaskId = &v
	return s
}

type CreateClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateGWSClusterRequest struct {
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateGWSClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateGWSClusterRequest) SetClusterType(v string) *CreateGWSClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateGWSClusterRequest) SetName(v string) *CreateGWSClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateGWSClusterRequest) SetVSwitchId(v string) *CreateGWSClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateGWSClusterRequest) SetVpcId(v string) *CreateGWSClusterRequest {
	s.VpcId = &v
	return s
}

type CreateGWSClusterResponseBody struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGWSClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGWSClusterResponseBody) SetClusterId(v string) *CreateGWSClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateGWSClusterResponseBody) SetRequestId(v string) *CreateGWSClusterResponseBody {
	s.RequestId = &v
	return s
}

type CreateGWSClusterResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGWSClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGWSClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateGWSClusterResponse) SetHeaders(v map[string]*string) *CreateGWSClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateGWSClusterResponse) SetBody(v *CreateGWSClusterResponseBody) *CreateGWSClusterResponse {
	s.Body = v
	return s
}

type CreateGWSImageRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateGWSImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSImageRequest) GoString() string {
	return s.String()
}

func (s *CreateGWSImageRequest) SetInstanceId(v string) *CreateGWSImageRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateGWSImageRequest) SetName(v string) *CreateGWSImageRequest {
	s.Name = &v
	return s
}

type CreateGWSImageResponseBody struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGWSImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGWSImageResponseBody) SetImageId(v string) *CreateGWSImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateGWSImageResponseBody) SetRequestId(v string) *CreateGWSImageResponseBody {
	s.RequestId = &v
	return s
}

type CreateGWSImageResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGWSImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGWSImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSImageResponse) GoString() string {
	return s.String()
}

func (s *CreateGWSImageResponse) SetHeaders(v map[string]*string) *CreateGWSImageResponse {
	s.Headers = v
	return s
}

func (s *CreateGWSImageResponse) SetBody(v *CreateGWSImageResponseBody) *CreateGWSImageResponse {
	s.Body = v
	return s
}

type CreateGWSInstanceRequest struct {
	AllocatePublicAddress   *bool   `json:"AllocatePublicAddress,omitempty" xml:"AllocatePublicAddress,omitempty"`
	AppList                 *string `json:"AppList,omitempty" xml:"AppList,omitempty"`
	AutoRenew               *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ImageId                 *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceChargeType      *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceType            *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetChargeType      *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthIn  *int32  `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut *int32  `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Period                  *string `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit              *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	SystemDiskCategory      *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	SystemDiskSize          *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	VSwitchId               *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	WorkMode                *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s CreateGWSInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateGWSInstanceRequest) SetAllocatePublicAddress(v bool) *CreateGWSInstanceRequest {
	s.AllocatePublicAddress = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetAppList(v string) *CreateGWSInstanceRequest {
	s.AppList = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetAutoRenew(v bool) *CreateGWSInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetClusterId(v string) *CreateGWSInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetImageId(v string) *CreateGWSInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetInstanceChargeType(v string) *CreateGWSInstanceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetInstanceType(v string) *CreateGWSInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetInternetChargeType(v string) *CreateGWSInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetInternetMaxBandwidthIn(v int32) *CreateGWSInstanceRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetInternetMaxBandwidthOut(v int32) *CreateGWSInstanceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetName(v string) *CreateGWSInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetPeriod(v string) *CreateGWSInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetPeriodUnit(v string) *CreateGWSInstanceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetSystemDiskCategory(v string) *CreateGWSInstanceRequest {
	s.SystemDiskCategory = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetSystemDiskSize(v int32) *CreateGWSInstanceRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetVSwitchId(v string) *CreateGWSInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetWorkMode(v string) *CreateGWSInstanceRequest {
	s.WorkMode = &v
	return s
}

type CreateGWSInstanceResponseBody struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGWSInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGWSInstanceResponseBody) SetInstanceId(v string) *CreateGWSInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateGWSInstanceResponseBody) SetRequestId(v string) *CreateGWSInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateGWSInstanceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGWSInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGWSInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateGWSInstanceResponse) SetHeaders(v map[string]*string) *CreateGWSInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateGWSInstanceResponse) SetBody(v *CreateGWSInstanceResponseBody) *CreateGWSInstanceResponse {
	s.Body = v
	return s
}

type CreateHybridClusterRequest struct {
	EcsOrder                  *CreateHybridClusterRequestEcsOrder            `json:"EcsOrder,omitempty" xml:"EcsOrder,omitempty" type:"Struct"`
	Application               []*CreateHybridClusterRequestApplication       `json:"Application,omitempty" xml:"Application,omitempty" type:"Repeated"`
	ClientToken               *string                                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClientVersion             *string                                        `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ComputeSpotPriceLimit     *float32                                       `json:"ComputeSpotPriceLimit,omitempty" xml:"ComputeSpotPriceLimit,omitempty"`
	ComputeSpotStrategy       *string                                        `json:"ComputeSpotStrategy,omitempty" xml:"ComputeSpotStrategy,omitempty"`
	Description               *string                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	Domain                    *string                                        `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EhpcVersion               *string                                        `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	ImageId                   *string                                        `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageOwnerAlias           *string                                        `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	JobQueue                  *string                                        `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	KeyPairName               *string                                        `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	Location                  *string                                        `json:"Location,omitempty" xml:"Location,omitempty"`
	MultiOs                   *bool                                          `json:"MultiOs,omitempty" xml:"MultiOs,omitempty"`
	Name                      *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Nodes                     []*CreateHybridClusterRequestNodes             `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	OnPremiseVolumeLocalPath  *string                                        `json:"OnPremiseVolumeLocalPath,omitempty" xml:"OnPremiseVolumeLocalPath,omitempty"`
	OnPremiseVolumeMountPoint *string                                        `json:"OnPremiseVolumeMountPoint,omitempty" xml:"OnPremiseVolumeMountPoint,omitempty"`
	OnPremiseVolumeProtocol   *string                                        `json:"OnPremiseVolumeProtocol,omitempty" xml:"OnPremiseVolumeProtocol,omitempty"`
	OnPremiseVolumeRemotePath *string                                        `json:"OnPremiseVolumeRemotePath,omitempty" xml:"OnPremiseVolumeRemotePath,omitempty"`
	OsTag                     *string                                        `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	Password                  *string                                        `json:"Password,omitempty" xml:"Password,omitempty"`
	PostInstallScript         []*CreateHybridClusterRequestPostInstallScript `json:"PostInstallScript,omitempty" xml:"PostInstallScript,omitempty" type:"Repeated"`
	RemoteDirectory           *string                                        `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	ResourceGroupId           *string                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SchedulerPreInstall       *bool                                          `json:"SchedulerPreInstall,omitempty" xml:"SchedulerPreInstall,omitempty"`
	SecurityGroupId           *string                                        `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupName         *string                                        `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	VSwitchId                 *string                                        `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VolumeId                  *string                                        `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	VolumeMountpoint          *string                                        `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	VolumeProtocol            *string                                        `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	VolumeType                *string                                        `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
	VpcId                     *string                                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                    *string                                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateHybridClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequest) SetEcsOrder(v *CreateHybridClusterRequestEcsOrder) *CreateHybridClusterRequest {
	s.EcsOrder = v
	return s
}

func (s *CreateHybridClusterRequest) SetApplication(v []*CreateHybridClusterRequestApplication) *CreateHybridClusterRequest {
	s.Application = v
	return s
}

func (s *CreateHybridClusterRequest) SetClientToken(v string) *CreateHybridClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateHybridClusterRequest) SetClientVersion(v string) *CreateHybridClusterRequest {
	s.ClientVersion = &v
	return s
}

func (s *CreateHybridClusterRequest) SetComputeSpotPriceLimit(v float32) *CreateHybridClusterRequest {
	s.ComputeSpotPriceLimit = &v
	return s
}

func (s *CreateHybridClusterRequest) SetComputeSpotStrategy(v string) *CreateHybridClusterRequest {
	s.ComputeSpotStrategy = &v
	return s
}

func (s *CreateHybridClusterRequest) SetDescription(v string) *CreateHybridClusterRequest {
	s.Description = &v
	return s
}

func (s *CreateHybridClusterRequest) SetDomain(v string) *CreateHybridClusterRequest {
	s.Domain = &v
	return s
}

func (s *CreateHybridClusterRequest) SetEhpcVersion(v string) *CreateHybridClusterRequest {
	s.EhpcVersion = &v
	return s
}

func (s *CreateHybridClusterRequest) SetImageId(v string) *CreateHybridClusterRequest {
	s.ImageId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetImageOwnerAlias(v string) *CreateHybridClusterRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *CreateHybridClusterRequest) SetJobQueue(v string) *CreateHybridClusterRequest {
	s.JobQueue = &v
	return s
}

func (s *CreateHybridClusterRequest) SetKeyPairName(v string) *CreateHybridClusterRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateHybridClusterRequest) SetLocation(v string) *CreateHybridClusterRequest {
	s.Location = &v
	return s
}

func (s *CreateHybridClusterRequest) SetMultiOs(v bool) *CreateHybridClusterRequest {
	s.MultiOs = &v
	return s
}

func (s *CreateHybridClusterRequest) SetName(v string) *CreateHybridClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateHybridClusterRequest) SetNodes(v []*CreateHybridClusterRequestNodes) *CreateHybridClusterRequest {
	s.Nodes = v
	return s
}

func (s *CreateHybridClusterRequest) SetOnPremiseVolumeLocalPath(v string) *CreateHybridClusterRequest {
	s.OnPremiseVolumeLocalPath = &v
	return s
}

func (s *CreateHybridClusterRequest) SetOnPremiseVolumeMountPoint(v string) *CreateHybridClusterRequest {
	s.OnPremiseVolumeMountPoint = &v
	return s
}

func (s *CreateHybridClusterRequest) SetOnPremiseVolumeProtocol(v string) *CreateHybridClusterRequest {
	s.OnPremiseVolumeProtocol = &v
	return s
}

func (s *CreateHybridClusterRequest) SetOnPremiseVolumeRemotePath(v string) *CreateHybridClusterRequest {
	s.OnPremiseVolumeRemotePath = &v
	return s
}

func (s *CreateHybridClusterRequest) SetOsTag(v string) *CreateHybridClusterRequest {
	s.OsTag = &v
	return s
}

func (s *CreateHybridClusterRequest) SetPassword(v string) *CreateHybridClusterRequest {
	s.Password = &v
	return s
}

func (s *CreateHybridClusterRequest) SetPostInstallScript(v []*CreateHybridClusterRequestPostInstallScript) *CreateHybridClusterRequest {
	s.PostInstallScript = v
	return s
}

func (s *CreateHybridClusterRequest) SetRemoteDirectory(v string) *CreateHybridClusterRequest {
	s.RemoteDirectory = &v
	return s
}

func (s *CreateHybridClusterRequest) SetResourceGroupId(v string) *CreateHybridClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetSchedulerPreInstall(v bool) *CreateHybridClusterRequest {
	s.SchedulerPreInstall = &v
	return s
}

func (s *CreateHybridClusterRequest) SetSecurityGroupId(v string) *CreateHybridClusterRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetSecurityGroupName(v string) *CreateHybridClusterRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVSwitchId(v string) *CreateHybridClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVolumeId(v string) *CreateHybridClusterRequest {
	s.VolumeId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVolumeMountpoint(v string) *CreateHybridClusterRequest {
	s.VolumeMountpoint = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVolumeProtocol(v string) *CreateHybridClusterRequest {
	s.VolumeProtocol = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVolumeType(v string) *CreateHybridClusterRequest {
	s.VolumeType = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVpcId(v string) *CreateHybridClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetZoneId(v string) *CreateHybridClusterRequest {
	s.ZoneId = &v
	return s
}

type CreateHybridClusterRequestEcsOrder struct {
	Compute *CreateHybridClusterRequestEcsOrderCompute `json:"Compute,omitempty" xml:"Compute,omitempty" require:"true" type:"Struct"`
}

func (s CreateHybridClusterRequestEcsOrder) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestEcsOrder) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestEcsOrder) SetCompute(v *CreateHybridClusterRequestEcsOrderCompute) *CreateHybridClusterRequestEcsOrder {
	s.Compute = v
	return s
}

type CreateHybridClusterRequestEcsOrderCompute struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateHybridClusterRequestEcsOrderCompute) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestEcsOrderCompute) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestEcsOrderCompute) SetInstanceType(v string) *CreateHybridClusterRequestEcsOrderCompute {
	s.InstanceType = &v
	return s
}

type CreateHybridClusterRequestApplication struct {
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateHybridClusterRequestApplication) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestApplication) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestApplication) SetTag(v string) *CreateHybridClusterRequestApplication {
	s.Tag = &v
	return s
}

type CreateHybridClusterRequestNodes struct {
	AccountType   *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	HostName      *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	IpAddress     *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Role          *string `json:"Role,omitempty" xml:"Role,omitempty"`
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
}

func (s CreateHybridClusterRequestNodes) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestNodes) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestNodes) SetAccountType(v string) *CreateHybridClusterRequestNodes {
	s.AccountType = &v
	return s
}

func (s *CreateHybridClusterRequestNodes) SetHostName(v string) *CreateHybridClusterRequestNodes {
	s.HostName = &v
	return s
}

func (s *CreateHybridClusterRequestNodes) SetIpAddress(v string) *CreateHybridClusterRequestNodes {
	s.IpAddress = &v
	return s
}

func (s *CreateHybridClusterRequestNodes) SetRole(v string) *CreateHybridClusterRequestNodes {
	s.Role = &v
	return s
}

func (s *CreateHybridClusterRequestNodes) SetSchedulerType(v string) *CreateHybridClusterRequestNodes {
	s.SchedulerType = &v
	return s
}

type CreateHybridClusterRequestPostInstallScript struct {
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateHybridClusterRequestPostInstallScript) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestPostInstallScript) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestPostInstallScript) SetArgs(v string) *CreateHybridClusterRequestPostInstallScript {
	s.Args = &v
	return s
}

func (s *CreateHybridClusterRequestPostInstallScript) SetUrl(v string) *CreateHybridClusterRequestPostInstallScript {
	s.Url = &v
	return s
}

type CreateHybridClusterResponseBody struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateHybridClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterResponseBody) SetClusterId(v string) *CreateHybridClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateHybridClusterResponseBody) SetRequestId(v string) *CreateHybridClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHybridClusterResponseBody) SetTaskId(v string) *CreateHybridClusterResponseBody {
	s.TaskId = &v
	return s
}

type CreateHybridClusterResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateHybridClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHybridClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterResponse) SetHeaders(v map[string]*string) *CreateHybridClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateHybridClusterResponse) SetBody(v *CreateHybridClusterResponseBody) *CreateHybridClusterResponse {
	s.Body = v
	return s
}

type CreateJobFileRequest struct {
	ClusterId         *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Content           *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RunasUser         *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	RunasUserPassword *string `json:"RunasUserPassword,omitempty" xml:"RunasUserPassword,omitempty"`
	TargetFile        *string `json:"TargetFile,omitempty" xml:"TargetFile,omitempty"`
}

func (s CreateJobFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobFileRequest) GoString() string {
	return s.String()
}

func (s *CreateJobFileRequest) SetClusterId(v string) *CreateJobFileRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateJobFileRequest) SetContent(v string) *CreateJobFileRequest {
	s.Content = &v
	return s
}

func (s *CreateJobFileRequest) SetRunasUser(v string) *CreateJobFileRequest {
	s.RunasUser = &v
	return s
}

func (s *CreateJobFileRequest) SetRunasUserPassword(v string) *CreateJobFileRequest {
	s.RunasUserPassword = &v
	return s
}

func (s *CreateJobFileRequest) SetTargetFile(v string) *CreateJobFileRequest {
	s.TargetFile = &v
	return s
}

type CreateJobFileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateJobFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobFileResponseBody) SetRequestId(v string) *CreateJobFileResponseBody {
	s.RequestId = &v
	return s
}

type CreateJobFileResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateJobFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateJobFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobFileResponse) GoString() string {
	return s.String()
}

func (s *CreateJobFileResponse) SetHeaders(v map[string]*string) *CreateJobFileResponse {
	s.Headers = v
	return s
}

func (s *CreateJobFileResponse) SetBody(v *CreateJobFileResponseBody) *CreateJobFileResponse {
	s.Body = v
	return s
}

type CreateJobTemplateRequest struct {
	ArrayRequest       *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	ClockTime          *string `json:"ClockTime,omitempty" xml:"ClockTime,omitempty"`
	CommandLine        *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	Gpu                *int32  `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	InputFileUrl       *string `json:"InputFileUrl,omitempty" xml:"InputFileUrl,omitempty"`
	Mem                *string `json:"Mem,omitempty" xml:"Mem,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Node               *int32  `json:"Node,omitempty" xml:"Node,omitempty"`
	PackagePath        *string `json:"PackagePath,omitempty" xml:"PackagePath,omitempty"`
	Priority           *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Queue              *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	ReRunable          *bool   `json:"ReRunable,omitempty" xml:"ReRunable,omitempty"`
	RunasUser          *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" xml:"StderrRedirectPath,omitempty"`
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" xml:"StdoutRedirectPath,omitempty"`
	Task               *int32  `json:"Task,omitempty" xml:"Task,omitempty"`
	Thread             *int32  `json:"Thread,omitempty" xml:"Thread,omitempty"`
	UnzipCmd           *string `json:"UnzipCmd,omitempty" xml:"UnzipCmd,omitempty"`
	Variables          *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
	WithUnzipCmd       *bool   `json:"WithUnzipCmd,omitempty" xml:"WithUnzipCmd,omitempty"`
}

func (s CreateJobTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequest) SetArrayRequest(v string) *CreateJobTemplateRequest {
	s.ArrayRequest = &v
	return s
}

func (s *CreateJobTemplateRequest) SetClockTime(v string) *CreateJobTemplateRequest {
	s.ClockTime = &v
	return s
}

func (s *CreateJobTemplateRequest) SetCommandLine(v string) *CreateJobTemplateRequest {
	s.CommandLine = &v
	return s
}

func (s *CreateJobTemplateRequest) SetGpu(v int32) *CreateJobTemplateRequest {
	s.Gpu = &v
	return s
}

func (s *CreateJobTemplateRequest) SetInputFileUrl(v string) *CreateJobTemplateRequest {
	s.InputFileUrl = &v
	return s
}

func (s *CreateJobTemplateRequest) SetMem(v string) *CreateJobTemplateRequest {
	s.Mem = &v
	return s
}

func (s *CreateJobTemplateRequest) SetName(v string) *CreateJobTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateJobTemplateRequest) SetNode(v int32) *CreateJobTemplateRequest {
	s.Node = &v
	return s
}

func (s *CreateJobTemplateRequest) SetPackagePath(v string) *CreateJobTemplateRequest {
	s.PackagePath = &v
	return s
}

func (s *CreateJobTemplateRequest) SetPriority(v int32) *CreateJobTemplateRequest {
	s.Priority = &v
	return s
}

func (s *CreateJobTemplateRequest) SetQueue(v string) *CreateJobTemplateRequest {
	s.Queue = &v
	return s
}

func (s *CreateJobTemplateRequest) SetReRunable(v bool) *CreateJobTemplateRequest {
	s.ReRunable = &v
	return s
}

func (s *CreateJobTemplateRequest) SetRunasUser(v string) *CreateJobTemplateRequest {
	s.RunasUser = &v
	return s
}

func (s *CreateJobTemplateRequest) SetStderrRedirectPath(v string) *CreateJobTemplateRequest {
	s.StderrRedirectPath = &v
	return s
}

func (s *CreateJobTemplateRequest) SetStdoutRedirectPath(v string) *CreateJobTemplateRequest {
	s.StdoutRedirectPath = &v
	return s
}

func (s *CreateJobTemplateRequest) SetTask(v int32) *CreateJobTemplateRequest {
	s.Task = &v
	return s
}

func (s *CreateJobTemplateRequest) SetThread(v int32) *CreateJobTemplateRequest {
	s.Thread = &v
	return s
}

func (s *CreateJobTemplateRequest) SetUnzipCmd(v string) *CreateJobTemplateRequest {
	s.UnzipCmd = &v
	return s
}

func (s *CreateJobTemplateRequest) SetVariables(v string) *CreateJobTemplateRequest {
	s.Variables = &v
	return s
}

func (s *CreateJobTemplateRequest) SetWithUnzipCmd(v bool) *CreateJobTemplateRequest {
	s.WithUnzipCmd = &v
	return s
}

type CreateJobTemplateResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateJobTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateResponseBody) SetRequestId(v string) *CreateJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobTemplateResponseBody) SetTemplateId(v string) *CreateJobTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type CreateJobTemplateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateJobTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateJobTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateResponse) SetHeaders(v map[string]*string) *CreateJobTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateJobTemplateResponse) SetBody(v *CreateJobTemplateResponseBody) *CreateJobTemplateResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ReleaseInstance *string `json:"ReleaseInstance,omitempty" xml:"ReleaseInstance,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterRequest) SetReleaseInstance(v string) *DeleteClusterRequest {
	s.ReleaseInstance = &v
	return s
}

type DeleteClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClusterResponseBody) SetTaskId(v string) *DeleteClusterResponseBody {
	s.TaskId = &v
	return s
}

type DeleteClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type DeleteContainerAppsRequest struct {
	ContainerApp []*DeleteContainerAppsRequestContainerApp `json:"ContainerApp,omitempty" xml:"ContainerApp,omitempty" type:"Repeated"`
}

func (s DeleteContainerAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerAppsRequest) GoString() string {
	return s.String()
}

func (s *DeleteContainerAppsRequest) SetContainerApp(v []*DeleteContainerAppsRequestContainerApp) *DeleteContainerAppsRequest {
	s.ContainerApp = v
	return s
}

type DeleteContainerAppsRequestContainerApp struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteContainerAppsRequestContainerApp) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerAppsRequestContainerApp) GoString() string {
	return s.String()
}

func (s *DeleteContainerAppsRequestContainerApp) SetId(v string) *DeleteContainerAppsRequestContainerApp {
	s.Id = &v
	return s
}

type DeleteContainerAppsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContainerAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContainerAppsResponseBody) SetRequestId(v string) *DeleteContainerAppsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteContainerAppsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteContainerAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteContainerAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerAppsResponse) GoString() string {
	return s.String()
}

func (s *DeleteContainerAppsResponse) SetHeaders(v map[string]*string) *DeleteContainerAppsResponse {
	s.Headers = v
	return s
}

func (s *DeleteContainerAppsResponse) SetBody(v *DeleteContainerAppsResponseBody) *DeleteContainerAppsResponse {
	s.Body = v
	return s
}

type DeleteGWSClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DeleteGWSClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteGWSClusterRequest) SetClusterId(v string) *DeleteGWSClusterRequest {
	s.ClusterId = &v
	return s
}

type DeleteGWSClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGWSClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGWSClusterResponseBody) SetRequestId(v string) *DeleteGWSClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGWSClusterResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGWSClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGWSClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteGWSClusterResponse) SetHeaders(v map[string]*string) *DeleteGWSClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteGWSClusterResponse) SetBody(v *DeleteGWSClusterResponseBody) *DeleteGWSClusterResponse {
	s.Body = v
	return s
}

type DeleteGWSInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteGWSInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGWSInstanceRequest) SetInstanceId(v string) *DeleteGWSInstanceRequest {
	s.InstanceId = &v
	return s
}

type DeleteGWSInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGWSInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGWSInstanceResponseBody) SetRequestId(v string) *DeleteGWSInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGWSInstanceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGWSInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGWSInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGWSInstanceResponse) SetHeaders(v map[string]*string) *DeleteGWSInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGWSInstanceResponse) SetBody(v *DeleteGWSInstanceResponseBody) *DeleteGWSInstanceResponse {
	s.Body = v
	return s
}

type DeleteImageRequest struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
	ImageTag      *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	Repository    *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
}

func (s DeleteImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) SetClusterId(v string) *DeleteImageRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteImageRequest) SetContainerType(v string) *DeleteImageRequest {
	s.ContainerType = &v
	return s
}

func (s *DeleteImageRequest) SetImageTag(v string) *DeleteImageRequest {
	s.ImageTag = &v
	return s
}

func (s *DeleteImageRequest) SetRepository(v string) *DeleteImageRequest {
	s.Repository = &v
	return s
}

type DeleteImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBody) SetRequestId(v string) *DeleteImageResponseBody {
	s.RequestId = &v
	return s
}

type DeleteImageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageResponse) SetHeaders(v map[string]*string) *DeleteImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageResponse) SetBody(v *DeleteImageResponseBody) *DeleteImageResponse {
	s.Body = v
	return s
}

type DeleteJobTemplatesRequest struct {
	Templates *string `json:"Templates,omitempty" xml:"Templates,omitempty"`
}

func (s DeleteJobTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobTemplatesRequest) SetTemplates(v string) *DeleteJobTemplatesRequest {
	s.Templates = &v
	return s
}

type DeleteJobTemplatesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobTemplatesResponseBody) SetRequestId(v string) *DeleteJobTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteJobTemplatesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteJobTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteJobTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobTemplatesResponse) SetHeaders(v map[string]*string) *DeleteJobTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobTemplatesResponse) SetBody(v *DeleteJobTemplatesResponseBody) *DeleteJobTemplatesResponse {
	s.Body = v
	return s
}

type DeleteJobsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Jobs      *string `json:"Jobs,omitempty" xml:"Jobs,omitempty"`
}

func (s DeleteJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobsRequest) SetClusterId(v string) *DeleteJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteJobsRequest) SetJobs(v string) *DeleteJobsRequest {
	s.Jobs = &v
	return s
}

type DeleteJobsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobsResponseBody) SetRequestId(v string) *DeleteJobsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteJobsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobsResponse) SetHeaders(v map[string]*string) *DeleteJobsResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobsResponse) SetBody(v *DeleteJobsResponseBody) *DeleteJobsResponse {
	s.Body = v
	return s
}

type DeleteNodesRequest struct {
	ClusterId       *string                       `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Instance        []*DeleteNodesRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
	ReleaseInstance *bool                         `json:"ReleaseInstance,omitempty" xml:"ReleaseInstance,omitempty"`
	Sync            *bool                         `json:"Sync,omitempty" xml:"Sync,omitempty"`
}

func (s DeleteNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodesRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodesRequest) SetClusterId(v string) *DeleteNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteNodesRequest) SetInstance(v []*DeleteNodesRequestInstance) *DeleteNodesRequest {
	s.Instance = v
	return s
}

func (s *DeleteNodesRequest) SetReleaseInstance(v bool) *DeleteNodesRequest {
	s.ReleaseInstance = &v
	return s
}

func (s *DeleteNodesRequest) SetSync(v bool) *DeleteNodesRequest {
	s.Sync = &v
	return s
}

type DeleteNodesRequestInstance struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteNodesRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodesRequestInstance) GoString() string {
	return s.String()
}

func (s *DeleteNodesRequestInstance) SetId(v string) *DeleteNodesRequestInstance {
	s.Id = &v
	return s
}

type DeleteNodesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodesResponseBody) SetRequestId(v string) *DeleteNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNodesResponseBody) SetTaskId(v string) *DeleteNodesResponseBody {
	s.TaskId = &v
	return s
}

type DeleteNodesResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodesResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodesResponse) SetHeaders(v map[string]*string) *DeleteNodesResponse {
	s.Headers = v
	return s
}

func (s *DeleteNodesResponse) SetBody(v *DeleteNodesResponseBody) *DeleteNodesResponse {
	s.Body = v
	return s
}

type DeleteQueueRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s DeleteQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueRequest) GoString() string {
	return s.String()
}

func (s *DeleteQueueRequest) SetClusterId(v string) *DeleteQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteQueueRequest) SetQueueName(v string) *DeleteQueueRequest {
	s.QueueName = &v
	return s
}

type DeleteQueueResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponseBody) SetRequestId(v string) *DeleteQueueResponseBody {
	s.RequestId = &v
	return s
}

type DeleteQueueResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueResponse) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponse) SetHeaders(v map[string]*string) *DeleteQueueResponse {
	s.Headers = v
	return s
}

func (s *DeleteQueueResponse) SetBody(v *DeleteQueueResponseBody) *DeleteQueueResponse {
	s.Body = v
	return s
}

type DeleteSecurityGroupRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DeleteSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupRequest) SetClusterId(v string) *DeleteSecurityGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteSecurityGroupRequest) SetSecurityGroupId(v string) *DeleteSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

type DeleteSecurityGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecurityGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupResponseBody) SetRequestId(v string) *DeleteSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSecurityGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupResponse) SetHeaders(v map[string]*string) *DeleteSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityGroupResponse) SetBody(v *DeleteSecurityGroupResponseBody) *DeleteSecurityGroupResponse {
	s.Body = v
	return s
}

type DeleteUsersRequest struct {
	ClusterId *string                   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	User      []*DeleteUsersRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s DeleteUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersRequest) GoString() string {
	return s.String()
}

func (s *DeleteUsersRequest) SetClusterId(v string) *DeleteUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteUsersRequest) SetUser(v []*DeleteUsersRequestUser) *DeleteUsersRequest {
	s.User = v
	return s
}

type DeleteUsersRequestUser struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteUsersRequestUser) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersRequestUser) GoString() string {
	return s.String()
}

func (s *DeleteUsersRequestUser) SetName(v string) *DeleteUsersRequestUser {
	s.Name = &v
	return s
}

type DeleteUsersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUsersResponseBody) SetRequestId(v string) *DeleteUsersResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUsersResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersResponse) GoString() string {
	return s.String()
}

func (s *DeleteUsersResponse) SetHeaders(v map[string]*string) *DeleteUsersResponse {
	s.Headers = v
	return s
}

func (s *DeleteUsersResponse) SetBody(v *DeleteUsersResponseBody) *DeleteUsersResponse {
	s.Body = v
	return s
}

type DescribeAutoScaleConfigRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeAutoScaleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoScaleConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoScaleConfigRequest) SetClusterId(v string) *DescribeAutoScaleConfigRequest {
	s.ClusterId = &v
	return s
}

type DescribeAutoScaleConfigResponseBody struct {
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterType             *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	EnableAutoGrow          *bool   `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	EnableAutoShrink        *bool   `json:"EnableAutoShrink,omitempty" xml:"EnableAutoShrink,omitempty"`
	ExcludeNodes            *string `json:"ExcludeNodes,omitempty" xml:"ExcludeNodes,omitempty"`
	ExtraNodesGrowRatio     *int32  `json:"ExtraNodesGrowRatio,omitempty" xml:"ExtraNodesGrowRatio,omitempty"`
	GrowIntervalInMinutes   *int32  `json:"GrowIntervalInMinutes,omitempty" xml:"GrowIntervalInMinutes,omitempty"`
	GrowRatio               *int32  `json:"GrowRatio,omitempty" xml:"GrowRatio,omitempty"`
	GrowTimeoutInMinutes    *int32  `json:"GrowTimeoutInMinutes,omitempty" xml:"GrowTimeoutInMinutes,omitempty"`
	MaxNodesInCluster       *int32  `json:"MaxNodesInCluster,omitempty" xml:"MaxNodesInCluster,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShrinkIdleTimes         *int32  `json:"ShrinkIdleTimes,omitempty" xml:"ShrinkIdleTimes,omitempty"`
	ShrinkIntervalInMinutes *int32  `json:"ShrinkIntervalInMinutes,omitempty" xml:"ShrinkIntervalInMinutes,omitempty"`
	SpotPriceLimit          *string `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy            *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	Uid                     *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeAutoScaleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoScaleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoScaleConfigResponseBody) SetClusterId(v string) *DescribeAutoScaleConfigResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetClusterType(v string) *DescribeAutoScaleConfigResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetEnableAutoGrow(v bool) *DescribeAutoScaleConfigResponseBody {
	s.EnableAutoGrow = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetEnableAutoShrink(v bool) *DescribeAutoScaleConfigResponseBody {
	s.EnableAutoShrink = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetExcludeNodes(v string) *DescribeAutoScaleConfigResponseBody {
	s.ExcludeNodes = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetExtraNodesGrowRatio(v int32) *DescribeAutoScaleConfigResponseBody {
	s.ExtraNodesGrowRatio = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetGrowIntervalInMinutes(v int32) *DescribeAutoScaleConfigResponseBody {
	s.GrowIntervalInMinutes = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetGrowRatio(v int32) *DescribeAutoScaleConfigResponseBody {
	s.GrowRatio = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetGrowTimeoutInMinutes(v int32) *DescribeAutoScaleConfigResponseBody {
	s.GrowTimeoutInMinutes = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetMaxNodesInCluster(v int32) *DescribeAutoScaleConfigResponseBody {
	s.MaxNodesInCluster = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetRequestId(v string) *DescribeAutoScaleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetShrinkIdleTimes(v int32) *DescribeAutoScaleConfigResponseBody {
	s.ShrinkIdleTimes = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetShrinkIntervalInMinutes(v int32) *DescribeAutoScaleConfigResponseBody {
	s.ShrinkIntervalInMinutes = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetSpotPriceLimit(v string) *DescribeAutoScaleConfigResponseBody {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetSpotStrategy(v string) *DescribeAutoScaleConfigResponseBody {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetUid(v string) *DescribeAutoScaleConfigResponseBody {
	s.Uid = &v
	return s
}

type DescribeAutoScaleConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAutoScaleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutoScaleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoScaleConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoScaleConfigResponse) SetHeaders(v map[string]*string) *DescribeAutoScaleConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoScaleConfigResponse) SetBody(v *DescribeAutoScaleConfigResponseBody) *DescribeAutoScaleConfigResponse {
	s.Body = v
	return s
}

type DescribeClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterRequest) SetClusterId(v string) *DescribeClusterRequest {
	s.ClusterId = &v
	return s
}

type DescribeClusterResponseBody struct {
	ClusterInfo *DescribeClusterResponseBodyClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Struct"`
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBody) SetClusterInfo(v *DescribeClusterResponseBodyClusterInfo) *DescribeClusterResponseBody {
	s.ClusterInfo = v
	return s
}

func (s *DescribeClusterResponseBody) SetRequestId(v string) *DescribeClusterResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClusterResponseBodyClusterInfo struct {
	AccountType        *string                                                   `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Applications       *DescribeClusterResponseBodyClusterInfoApplications       `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	BaseOsTag          *string                                                   `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	ClientVersion      *string                                                   `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	CreateTime         *string                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeployMode         *string                                                   `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	Description        *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	EcsChargeType      *string                                                   `json:"EcsChargeType,omitempty" xml:"EcsChargeType,omitempty"`
	EcsInfo            *DescribeClusterResponseBodyClusterInfoEcsInfo            `json:"EcsInfo,omitempty" xml:"EcsInfo,omitempty" type:"Struct"`
	HaEnable           *bool                                                     `json:"HaEnable,omitempty" xml:"HaEnable,omitempty"`
	Id                 *string                                                   `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageId            *string                                                   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName          *string                                                   `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageOwnerAlias    *string                                                   `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	KeyPairName        *string                                                   `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	Location           *string                                                   `json:"Location,omitempty" xml:"Location,omitempty"`
	Name               *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	OnPremiseInfo      *DescribeClusterResponseBodyClusterInfoOnPremiseInfo      `json:"OnPremiseInfo,omitempty" xml:"OnPremiseInfo,omitempty" type:"Struct"`
	OsTag              *string                                                   `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	PostInstallScripts *DescribeClusterResponseBodyClusterInfoPostInstallScripts `json:"PostInstallScripts,omitempty" xml:"PostInstallScripts,omitempty" type:"Struct"`
	RegionId           *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RemoteDirectory    *string                                                   `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	SccClusterId       *string                                                   `json:"SccClusterId,omitempty" xml:"SccClusterId,omitempty"`
	SchedulerType      *string                                                   `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	SecurityGroupId    *string                                                   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Status             *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchId          *string                                                   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VolumeId           *string                                                   `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	VolumeMountpoint   *string                                                   `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	VolumeProtocol     *string                                                   `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	VolumeType         *string                                                   `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
	VpcId              *string                                                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfo) SetAccountType(v string) *DescribeClusterResponseBodyClusterInfo {
	s.AccountType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetApplications(v *DescribeClusterResponseBodyClusterInfoApplications) *DescribeClusterResponseBodyClusterInfo {
	s.Applications = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetBaseOsTag(v string) *DescribeClusterResponseBodyClusterInfo {
	s.BaseOsTag = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetClientVersion(v string) *DescribeClusterResponseBodyClusterInfo {
	s.ClientVersion = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetCreateTime(v string) *DescribeClusterResponseBodyClusterInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetDeployMode(v string) *DescribeClusterResponseBodyClusterInfo {
	s.DeployMode = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetDescription(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Description = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetEcsChargeType(v string) *DescribeClusterResponseBodyClusterInfo {
	s.EcsChargeType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetEcsInfo(v *DescribeClusterResponseBodyClusterInfoEcsInfo) *DescribeClusterResponseBodyClusterInfo {
	s.EcsInfo = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetHaEnable(v bool) *DescribeClusterResponseBodyClusterInfo {
	s.HaEnable = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Id = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetImageId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.ImageId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetImageName(v string) *DescribeClusterResponseBodyClusterInfo {
	s.ImageName = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetImageOwnerAlias(v string) *DescribeClusterResponseBodyClusterInfo {
	s.ImageOwnerAlias = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetKeyPairName(v string) *DescribeClusterResponseBodyClusterInfo {
	s.KeyPairName = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetLocation(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Location = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetName(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Name = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetOnPremiseInfo(v *DescribeClusterResponseBodyClusterInfoOnPremiseInfo) *DescribeClusterResponseBodyClusterInfo {
	s.OnPremiseInfo = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetOsTag(v string) *DescribeClusterResponseBodyClusterInfo {
	s.OsTag = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetPostInstallScripts(v *DescribeClusterResponseBodyClusterInfoPostInstallScripts) *DescribeClusterResponseBodyClusterInfo {
	s.PostInstallScripts = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetRegionId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetRemoteDirectory(v string) *DescribeClusterResponseBodyClusterInfo {
	s.RemoteDirectory = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetSccClusterId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.SccClusterId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetSchedulerType(v string) *DescribeClusterResponseBodyClusterInfo {
	s.SchedulerType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetSecurityGroupId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetStatus(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Status = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVSwitchId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVolumeId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VolumeId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVolumeMountpoint(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VolumeMountpoint = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVolumeProtocol(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VolumeProtocol = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVolumeType(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VolumeType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVpcId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VpcId = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoApplications struct {
	ApplicationInfo []*DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo `json:"ApplicationInfo,omitempty" xml:"ApplicationInfo,omitempty" type:"Repeated"`
}

func (s DescribeClusterResponseBodyClusterInfoApplications) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoApplications) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoApplications) SetApplicationInfo(v []*DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) *DescribeClusterResponseBodyClusterInfoApplications {
	s.ApplicationInfo = v
	return s
}

type DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Tag     *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) SetName(v string) *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo {
	s.Name = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) SetTag(v string) *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo {
	s.Tag = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) SetVersion(v string) *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo {
	s.Version = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfo struct {
	Compute  *DescribeClusterResponseBodyClusterInfoEcsInfoCompute  `json:"Compute,omitempty" xml:"Compute,omitempty" type:"Struct"`
	Login    *DescribeClusterResponseBodyClusterInfoEcsInfoLogin    `json:"Login,omitempty" xml:"Login,omitempty" type:"Struct"`
	Manager  *DescribeClusterResponseBodyClusterInfoEcsInfoManager  `json:"Manager,omitempty" xml:"Manager,omitempty" type:"Struct"`
	ProxyMgr *DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr `json:"ProxyMgr,omitempty" xml:"ProxyMgr,omitempty" type:"Struct"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfo) SetCompute(v *DescribeClusterResponseBodyClusterInfoEcsInfoCompute) *DescribeClusterResponseBodyClusterInfoEcsInfo {
	s.Compute = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfo) SetLogin(v *DescribeClusterResponseBodyClusterInfoEcsInfoLogin) *DescribeClusterResponseBodyClusterInfoEcsInfo {
	s.Login = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfo) SetManager(v *DescribeClusterResponseBodyClusterInfoEcsInfoManager) *DescribeClusterResponseBodyClusterInfoEcsInfo {
	s.Manager = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfo) SetProxyMgr(v *DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr) *DescribeClusterResponseBodyClusterInfoEcsInfo {
	s.ProxyMgr = v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfoCompute struct {
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoCompute) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoCompute) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoCompute) SetCount(v int32) *DescribeClusterResponseBodyClusterInfoEcsInfoCompute {
	s.Count = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoCompute) SetInstanceType(v string) *DescribeClusterResponseBodyClusterInfoEcsInfoCompute {
	s.InstanceType = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfoLogin struct {
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoLogin) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoLogin) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoLogin) SetCount(v int32) *DescribeClusterResponseBodyClusterInfoEcsInfoLogin {
	s.Count = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoLogin) SetInstanceType(v string) *DescribeClusterResponseBodyClusterInfoEcsInfoLogin {
	s.InstanceType = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfoManager struct {
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoManager) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoManager) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoManager) SetCount(v int32) *DescribeClusterResponseBodyClusterInfoEcsInfoManager {
	s.Count = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoManager) SetInstanceType(v string) *DescribeClusterResponseBodyClusterInfoEcsInfoManager {
	s.InstanceType = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr struct {
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr) SetCount(v int32) *DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr {
	s.Count = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr) SetInstanceType(v string) *DescribeClusterResponseBodyClusterInfoEcsInfoProxyMgr {
	s.InstanceType = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoOnPremiseInfo struct {
	OnPremiseInfo []*DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo `json:"OnPremiseInfo,omitempty" xml:"OnPremiseInfo,omitempty" type:"Repeated"`
}

func (s DescribeClusterResponseBodyClusterInfoOnPremiseInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoOnPremiseInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoOnPremiseInfo) SetOnPremiseInfo(v []*DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo) *DescribeClusterResponseBodyClusterInfoOnPremiseInfo {
	s.OnPremiseInfo = v
	return s
}

type DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo struct {
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	IP       *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo) SetHostName(v string) *DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo {
	s.HostName = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo) SetIP(v string) *DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo {
	s.IP = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo) SetType(v string) *DescribeClusterResponseBodyClusterInfoOnPremiseInfoOnPremiseInfo {
	s.Type = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoPostInstallScripts struct {
	PostInstallScriptInfo []*DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo `json:"PostInstallScriptInfo,omitempty" xml:"PostInstallScriptInfo,omitempty" type:"Repeated"`
}

func (s DescribeClusterResponseBodyClusterInfoPostInstallScripts) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoPostInstallScripts) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoPostInstallScripts) SetPostInstallScriptInfo(v []*DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo) *DescribeClusterResponseBodyClusterInfoPostInstallScripts {
	s.PostInstallScriptInfo = v
	return s
}

type DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo struct {
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo) SetArgs(v string) *DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo {
	s.Args = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo) SetUrl(v string) *DescribeClusterResponseBodyClusterInfoPostInstallScriptsPostInstallScriptInfo {
	s.Url = &v
	return s
}

type DescribeClusterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponse) SetHeaders(v map[string]*string) *DescribeClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterResponse) SetBody(v *DescribeClusterResponseBody) *DescribeClusterResponse {
	s.Body = v
	return s
}

type DescribeContainerAppRequest struct {
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
}

func (s DescribeContainerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerAppRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerAppRequest) SetContainerId(v string) *DescribeContainerAppRequest {
	s.ContainerId = &v
	return s
}

type DescribeContainerAppResponseBody struct {
	ContainerAppInfo *DescribeContainerAppResponseBodyContainerAppInfo `json:"ContainerAppInfo,omitempty" xml:"ContainerAppInfo,omitempty" type:"Struct"`
	RequestId        *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerAppResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerAppResponseBody) SetContainerAppInfo(v *DescribeContainerAppResponseBodyContainerAppInfo) *DescribeContainerAppResponseBody {
	s.ContainerAppInfo = v
	return s
}

func (s *DescribeContainerAppResponseBody) SetRequestId(v string) *DescribeContainerAppResponseBody {
	s.RequestId = &v
	return s
}

type DescribeContainerAppResponseBodyContainerAppInfo struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageTag    *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Repository  *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeContainerAppResponseBodyContainerAppInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerAppResponseBodyContainerAppInfo) GoString() string {
	return s.String()
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetCreateTime(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetDescription(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.Description = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetId(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.Id = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetImageTag(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.ImageTag = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetName(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.Name = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetRepository(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.Repository = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetType(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.Type = &v
	return s
}

type DescribeContainerAppResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeContainerAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerAppResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerAppResponse) SetHeaders(v map[string]*string) *DescribeContainerAppResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerAppResponse) SetBody(v *DescribeContainerAppResponseBody) *DescribeContainerAppResponse {
	s.Body = v
	return s
}

type DescribeGWSClusterPolicyRequest struct {
	AsyncMode *bool   `json:"AsyncMode,omitempty" xml:"AsyncMode,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeGWSClusterPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClusterPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeGWSClusterPolicyRequest) SetAsyncMode(v bool) *DescribeGWSClusterPolicyRequest {
	s.AsyncMode = &v
	return s
}

func (s *DescribeGWSClusterPolicyRequest) SetClusterId(v string) *DescribeGWSClusterPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeGWSClusterPolicyRequest) SetTaskId(v string) *DescribeGWSClusterPolicyRequest {
	s.TaskId = &v
	return s
}

type DescribeGWSClusterPolicyResponseBody struct {
	Clipboard   *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	LocalDrive  *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	Watermark   *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s DescribeGWSClusterPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClusterPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGWSClusterPolicyResponseBody) SetClipboard(v string) *DescribeGWSClusterPolicyResponseBody {
	s.Clipboard = &v
	return s
}

func (s *DescribeGWSClusterPolicyResponseBody) SetLocalDrive(v string) *DescribeGWSClusterPolicyResponseBody {
	s.LocalDrive = &v
	return s
}

func (s *DescribeGWSClusterPolicyResponseBody) SetRequestId(v string) *DescribeGWSClusterPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGWSClusterPolicyResponseBody) SetUsbRedirect(v string) *DescribeGWSClusterPolicyResponseBody {
	s.UsbRedirect = &v
	return s
}

func (s *DescribeGWSClusterPolicyResponseBody) SetWatermark(v string) *DescribeGWSClusterPolicyResponseBody {
	s.Watermark = &v
	return s
}

type DescribeGWSClusterPolicyResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGWSClusterPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGWSClusterPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClusterPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeGWSClusterPolicyResponse) SetHeaders(v map[string]*string) *DescribeGWSClusterPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeGWSClusterPolicyResponse) SetBody(v *DescribeGWSClusterPolicyResponseBody) *DescribeGWSClusterPolicyResponse {
	s.Body = v
	return s
}

type DescribeGWSClustersRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGWSClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeGWSClustersRequest) SetClusterId(v string) *DescribeGWSClustersRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeGWSClustersRequest) SetPageNumber(v int32) *DescribeGWSClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGWSClustersRequest) SetPageSize(v int32) *DescribeGWSClustersRequest {
	s.PageSize = &v
	return s
}

type DescribeGWSClustersResponseBody struct {
	CallerType *string                                  `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
	Clusters   *DescribeGWSClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Struct"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGWSClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGWSClustersResponseBody) SetCallerType(v string) *DescribeGWSClustersResponseBody {
	s.CallerType = &v
	return s
}

func (s *DescribeGWSClustersResponseBody) SetClusters(v *DescribeGWSClustersResponseBodyClusters) *DescribeGWSClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeGWSClustersResponseBody) SetPageNumber(v int32) *DescribeGWSClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGWSClustersResponseBody) SetPageSize(v int32) *DescribeGWSClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGWSClustersResponseBody) SetRequestId(v string) *DescribeGWSClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGWSClustersResponseBody) SetTotalCount(v int32) *DescribeGWSClustersResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeGWSClustersResponseBodyClusters struct {
	ClusterInfo []*DescribeGWSClustersResponseBodyClustersClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Repeated"`
}

func (s DescribeGWSClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeGWSClustersResponseBodyClusters) SetClusterInfo(v []*DescribeGWSClustersResponseBodyClustersClusterInfo) *DescribeGWSClustersResponseBodyClusters {
	s.ClusterInfo = v
	return s
}

type DescribeGWSClustersResponseBodyClustersClusterInfo struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceCount *int32  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeGWSClustersResponseBodyClustersClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClustersResponseBodyClustersClusterInfo) GoString() string {
	return s.String()
}

func (s *DescribeGWSClustersResponseBodyClustersClusterInfo) SetClusterId(v string) *DescribeGWSClustersResponseBodyClustersClusterInfo {
	s.ClusterId = &v
	return s
}

func (s *DescribeGWSClustersResponseBodyClustersClusterInfo) SetCreateTime(v string) *DescribeGWSClustersResponseBodyClustersClusterInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeGWSClustersResponseBodyClustersClusterInfo) SetInstanceCount(v int32) *DescribeGWSClustersResponseBodyClustersClusterInfo {
	s.InstanceCount = &v
	return s
}

func (s *DescribeGWSClustersResponseBodyClustersClusterInfo) SetStatus(v string) *DescribeGWSClustersResponseBodyClustersClusterInfo {
	s.Status = &v
	return s
}

func (s *DescribeGWSClustersResponseBodyClustersClusterInfo) SetVpcId(v string) *DescribeGWSClustersResponseBodyClustersClusterInfo {
	s.VpcId = &v
	return s
}

type DescribeGWSClustersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGWSClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGWSClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeGWSClustersResponse) SetHeaders(v map[string]*string) *DescribeGWSClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeGWSClustersResponse) SetBody(v *DescribeGWSClustersResponseBody) *DescribeGWSClustersResponse {
	s.Body = v
	return s
}

type DescribeGWSImagesRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGWSImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGWSImagesRequest) SetPageNumber(v int32) *DescribeGWSImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGWSImagesRequest) SetPageSize(v int32) *DescribeGWSImagesRequest {
	s.PageSize = &v
	return s
}

type DescribeGWSImagesResponseBody struct {
	Images     *DescribeGWSImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGWSImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGWSImagesResponseBody) SetImages(v *DescribeGWSImagesResponseBodyImages) *DescribeGWSImagesResponseBody {
	s.Images = v
	return s
}

func (s *DescribeGWSImagesResponseBody) SetPageNumber(v int32) *DescribeGWSImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGWSImagesResponseBody) SetPageSize(v int32) *DescribeGWSImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGWSImagesResponseBody) SetRequestId(v string) *DescribeGWSImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGWSImagesResponseBody) SetTotalCount(v int32) *DescribeGWSImagesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeGWSImagesResponseBodyImages struct {
	ImageInfo []*DescribeGWSImagesResponseBodyImagesImageInfo `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" type:"Repeated"`
}

func (s DescribeGWSImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeGWSImagesResponseBodyImages) SetImageInfo(v []*DescribeGWSImagesResponseBodyImagesImageInfo) *DescribeGWSImagesResponseBodyImages {
	s.ImageInfo = v
	return s
}

type DescribeGWSImagesResponseBodyImagesImageInfo struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ImageId    *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageType  *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Progress   *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGWSImagesResponseBodyImagesImageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSImagesResponseBodyImagesImageInfo) GoString() string {
	return s.String()
}

func (s *DescribeGWSImagesResponseBodyImagesImageInfo) SetCreateTime(v string) *DescribeGWSImagesResponseBodyImagesImageInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImagesImageInfo) SetImageId(v string) *DescribeGWSImagesResponseBodyImagesImageInfo {
	s.ImageId = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImagesImageInfo) SetImageType(v string) *DescribeGWSImagesResponseBodyImagesImageInfo {
	s.ImageType = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImagesImageInfo) SetName(v string) *DescribeGWSImagesResponseBodyImagesImageInfo {
	s.Name = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImagesImageInfo) SetProgress(v string) *DescribeGWSImagesResponseBodyImagesImageInfo {
	s.Progress = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImagesImageInfo) SetSize(v int32) *DescribeGWSImagesResponseBodyImagesImageInfo {
	s.Size = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImagesImageInfo) SetStatus(v string) *DescribeGWSImagesResponseBodyImagesImageInfo {
	s.Status = &v
	return s
}

type DescribeGWSImagesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGWSImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGWSImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGWSImagesResponse) SetHeaders(v map[string]*string) *DescribeGWSImagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGWSImagesResponse) SetBody(v *DescribeGWSImagesResponseBody) *DescribeGWSImagesResponse {
	s.Body = v
	return s
}

type DescribeGWSInstancesRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserUid    *int64  `json:"UserUid,omitempty" xml:"UserUid,omitempty"`
}

func (s DescribeGWSInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesRequest) SetClusterId(v string) *DescribeGWSInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeGWSInstancesRequest) SetInstanceId(v string) *DescribeGWSInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGWSInstancesRequest) SetPageNumber(v int32) *DescribeGWSInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGWSInstancesRequest) SetPageSize(v int32) *DescribeGWSInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGWSInstancesRequest) SetUserName(v string) *DescribeGWSInstancesRequest {
	s.UserName = &v
	return s
}

func (s *DescribeGWSInstancesRequest) SetUserUid(v int64) *DescribeGWSInstancesRequest {
	s.UserUid = &v
	return s
}

type DescribeGWSInstancesResponseBody struct {
	Instances  *DescribeGWSInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGWSInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesResponseBody) SetInstances(v *DescribeGWSInstancesResponseBodyInstances) *DescribeGWSInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeGWSInstancesResponseBody) SetPageNumber(v int32) *DescribeGWSInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGWSInstancesResponseBody) SetPageSize(v int32) *DescribeGWSInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGWSInstancesResponseBody) SetRequestId(v string) *DescribeGWSInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGWSInstancesResponseBody) SetTotalCount(v int32) *DescribeGWSInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeGWSInstancesResponseBodyInstances struct {
	InstanceInfo []*DescribeGWSInstancesResponseBodyInstancesInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Repeated"`
}

func (s DescribeGWSInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesResponseBodyInstances) SetInstanceInfo(v []*DescribeGWSInstancesResponseBodyInstancesInstanceInfo) *DescribeGWSInstancesResponseBodyInstances {
	s.InstanceInfo = v
	return s
}

type DescribeGWSInstancesResponseBodyInstancesInstanceInfo struct {
	AppList      *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Struct"`
	ClusterId    *string                                                       `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateTime   *string                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime   *string                                                       `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId   *string                                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string                                                       `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Name         *string                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	Status       *string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	UserName     *string                                                       `json:"UserName,omitempty" xml:"UserName,omitempty"`
	WorkMode     *string                                                       `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s DescribeGWSInstancesResponseBodyInstancesInstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesResponseBodyInstancesInstanceInfo) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetAppList(v *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppList) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.AppList = v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetClusterId(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.ClusterId = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetCreateTime(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetExpireTime(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.ExpireTime = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetInstanceId(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetInstanceType(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.InstanceType = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetName(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.Name = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetStatus(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.Status = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetUserName(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.UserName = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfo) SetWorkMode(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfo {
	s.WorkMode = &v
	return s
}

type DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppList struct {
	AppInfo []*DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo `json:"AppInfo,omitempty" xml:"AppInfo,omitempty" type:"Repeated"`
}

func (s DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppList) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppList) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppList) SetAppInfo(v []*DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo) *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppList {
	s.AppInfo = v
	return s
}

type DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo struct {
	AppArgs *string `json:"AppArgs,omitempty" xml:"AppArgs,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppPath *string `json:"AppPath,omitempty" xml:"AppPath,omitempty"`
}

func (s DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo) SetAppArgs(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo {
	s.AppArgs = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo) SetAppName(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo {
	s.AppName = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo) SetAppPath(v string) *DescribeGWSInstancesResponseBodyInstancesInstanceInfoAppListAppInfo {
	s.AppPath = &v
	return s
}

type DescribeGWSInstancesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGWSInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGWSInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesResponse) SetHeaders(v map[string]*string) *DescribeGWSInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGWSInstancesResponse) SetBody(v *DescribeGWSInstancesResponseBody) *DescribeGWSInstancesResponse {
	s.Body = v
	return s
}

type DescribeImageRequest struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
	ImageTag      *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	Repository    *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
}

func (s DescribeImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageRequest) SetClusterId(v string) *DescribeImageRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeImageRequest) SetContainerType(v string) *DescribeImageRequest {
	s.ContainerType = &v
	return s
}

func (s *DescribeImageRequest) SetImageTag(v string) *DescribeImageRequest {
	s.ImageTag = &v
	return s
}

func (s *DescribeImageRequest) SetRepository(v string) *DescribeImageRequest {
	s.Repository = &v
	return s
}

type DescribeImageResponseBody struct {
	ImageInfo *DescribeImageResponseBodyImageInfo `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageResponseBody) SetImageInfo(v *DescribeImageResponseBodyImageInfo) *DescribeImageResponseBody {
	s.ImageInfo = v
	return s
}

func (s *DescribeImageResponseBody) SetRequestId(v string) *DescribeImageResponseBody {
	s.RequestId = &v
	return s
}

type DescribeImageResponseBodyImageInfo struct {
	ImageId        *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Repository     *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	System         *string `json:"System,omitempty" xml:"System,omitempty"`
	Tag            *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateDateTime *string `json:"UpdateDateTime,omitempty" xml:"UpdateDateTime,omitempty"`
}

func (s DescribeImageResponseBodyImageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResponseBodyImageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageResponseBodyImageInfo) SetImageId(v string) *DescribeImageResponseBodyImageInfo {
	s.ImageId = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetRepository(v string) *DescribeImageResponseBodyImageInfo {
	s.Repository = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetStatus(v string) *DescribeImageResponseBodyImageInfo {
	s.Status = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetSystem(v string) *DescribeImageResponseBodyImageInfo {
	s.System = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetTag(v string) *DescribeImageResponseBodyImageInfo {
	s.Tag = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetType(v string) *DescribeImageResponseBodyImageInfo {
	s.Type = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetUpdateDateTime(v string) *DescribeImageResponseBodyImageInfo {
	s.UpdateDateTime = &v
	return s
}

type DescribeImageResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageResponse) SetHeaders(v map[string]*string) *DescribeImageResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageResponse) SetBody(v *DescribeImageResponseBody) *DescribeImageResponse {
	s.Body = v
	return s
}

type DescribeImageGatewayConfigRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeImageGatewayConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGatewayConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageGatewayConfigRequest) SetClusterId(v string) *DescribeImageGatewayConfigRequest {
	s.ClusterId = &v
	return s
}

type DescribeImageGatewayConfigResponseBody struct {
	Imagegw   *DescribeImageGatewayConfigResponseBodyImagegw `json:"Imagegw,omitempty" xml:"Imagegw,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageGatewayConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGatewayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageGatewayConfigResponseBody) SetImagegw(v *DescribeImageGatewayConfigResponseBodyImagegw) *DescribeImageGatewayConfigResponseBody {
	s.Imagegw = v
	return s
}

func (s *DescribeImageGatewayConfigResponseBody) SetRequestId(v string) *DescribeImageGatewayConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeImageGatewayConfigResponseBodyImagegw struct {
	DefaultImageLocation   *string                                                 `json:"DefaultImageLocation,omitempty" xml:"DefaultImageLocation,omitempty"`
	ImageExpirationTimeout *string                                                 `json:"ImageExpirationTimeout,omitempty" xml:"ImageExpirationTimeout,omitempty"`
	Locations              *DescribeImageGatewayConfigResponseBodyImagegwLocations `json:"Locations,omitempty" xml:"Locations,omitempty" type:"Struct"`
	MongoDBURI             *string                                                 `json:"MongoDBURI,omitempty" xml:"MongoDBURI,omitempty"`
	PullUpdateTimeout      *int64                                                  `json:"PullUpdateTimeout,omitempty" xml:"PullUpdateTimeout,omitempty"`
	UpdateDateTime         *string                                                 `json:"UpdateDateTime,omitempty" xml:"UpdateDateTime,omitempty"`
}

func (s DescribeImageGatewayConfigResponseBodyImagegw) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGatewayConfigResponseBodyImagegw) GoString() string {
	return s.String()
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetDefaultImageLocation(v string) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.DefaultImageLocation = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetImageExpirationTimeout(v string) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.ImageExpirationTimeout = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetLocations(v *DescribeImageGatewayConfigResponseBodyImagegwLocations) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.Locations = v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetMongoDBURI(v string) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.MongoDBURI = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetPullUpdateTimeout(v int64) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.PullUpdateTimeout = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetUpdateDateTime(v string) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.UpdateDateTime = &v
	return s
}

type DescribeImageGatewayConfigResponseBodyImagegwLocations struct {
	LocationInfo []*DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo `json:"LocationInfo,omitempty" xml:"LocationInfo,omitempty" type:"Repeated"`
}

func (s DescribeImageGatewayConfigResponseBodyImagegwLocations) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGatewayConfigResponseBodyImagegwLocations) GoString() string {
	return s.String()
}

func (s *DescribeImageGatewayConfigResponseBodyImagegwLocations) SetLocationInfo(v []*DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo) *DescribeImageGatewayConfigResponseBodyImagegwLocations {
	s.LocationInfo = v
	return s
}

type DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo struct {
	Authentication *string `json:"Authentication,omitempty" xml:"Authentication,omitempty"`
	Location       *string `json:"Location,omitempty" xml:"Location,omitempty"`
	RemoteType     *string `json:"RemoteType,omitempty" xml:"RemoteType,omitempty"`
	URL            *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo) SetAuthentication(v string) *DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo {
	s.Authentication = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo) SetLocation(v string) *DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo {
	s.Location = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo) SetRemoteType(v string) *DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo {
	s.RemoteType = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo) SetURL(v string) *DescribeImageGatewayConfigResponseBodyImagegwLocationsLocationInfo {
	s.URL = &v
	return s
}

type DescribeImageGatewayConfigResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeImageGatewayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImageGatewayConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGatewayConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageGatewayConfigResponse) SetHeaders(v map[string]*string) *DescribeImageGatewayConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageGatewayConfigResponse) SetBody(v *DescribeImageGatewayConfigResponseBody) *DescribeImageGatewayConfigResponse {
	s.Body = v
	return s
}

type DescribeImagePriceRequest struct {
	Amount    *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	Period    *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	SkuCode   *string `json:"SkuCode,omitempty" xml:"SkuCode,omitempty"`
}

func (s DescribeImagePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeImagePriceRequest) SetAmount(v int32) *DescribeImagePriceRequest {
	s.Amount = &v
	return s
}

func (s *DescribeImagePriceRequest) SetImageId(v string) *DescribeImagePriceRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeImagePriceRequest) SetOrderType(v string) *DescribeImagePriceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeImagePriceRequest) SetPeriod(v int32) *DescribeImagePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribeImagePriceRequest) SetPriceUnit(v string) *DescribeImagePriceRequest {
	s.PriceUnit = &v
	return s
}

func (s *DescribeImagePriceRequest) SetSkuCode(v string) *DescribeImagePriceRequest {
	s.SkuCode = &v
	return s
}

type DescribeImagePriceResponseBody struct {
	Amount        *int32   `json:"Amount,omitempty" xml:"Amount,omitempty"`
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	ImageId       *string  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	RequestId     *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TradePrice    *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeImagePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImagePriceResponseBody) SetAmount(v int32) *DescribeImagePriceResponseBody {
	s.Amount = &v
	return s
}

func (s *DescribeImagePriceResponseBody) SetDiscountPrice(v float32) *DescribeImagePriceResponseBody {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeImagePriceResponseBody) SetImageId(v string) *DescribeImagePriceResponseBody {
	s.ImageId = &v
	return s
}

func (s *DescribeImagePriceResponseBody) SetOriginalPrice(v float32) *DescribeImagePriceResponseBody {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeImagePriceResponseBody) SetRequestId(v string) *DescribeImagePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImagePriceResponseBody) SetTradePrice(v float32) *DescribeImagePriceResponseBody {
	s.TradePrice = &v
	return s
}

type DescribeImagePriceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeImagePriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImagePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeImagePriceResponse) SetHeaders(v map[string]*string) *DescribeImagePriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeImagePriceResponse) SetBody(v *DescribeImagePriceResponseBody) *DescribeImagePriceResponse {
	s.Body = v
	return s
}

type DescribeJobRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DescribeJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobRequest) SetClusterId(v string) *DescribeJobRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeJobRequest) SetJobId(v string) *DescribeJobRequest {
	s.JobId = &v
	return s
}

type DescribeJobResponseBody struct {
	Message   *DescribeJobResponseBodyMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBody) SetMessage(v *DescribeJobResponseBodyMessage) *DescribeJobResponseBody {
	s.Message = v
	return s
}

func (s *DescribeJobResponseBody) SetRequestId(v string) *DescribeJobResponseBody {
	s.RequestId = &v
	return s
}

type DescribeJobResponseBodyMessage struct {
	JobInfo *string `json:"JobInfo,omitempty" xml:"JobInfo,omitempty"`
}

func (s DescribeJobResponseBodyMessage) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobResponseBodyMessage) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBodyMessage) SetJobInfo(v string) *DescribeJobResponseBodyMessage {
	s.JobInfo = &v
	return s
}

type DescribeJobResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobResponse) SetHeaders(v map[string]*string) *DescribeJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobResponse) SetBody(v *DescribeJobResponseBody) *DescribeJobResponse {
	s.Body = v
	return s
}

type DescribeNFSClientStatusRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeNFSClientStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNFSClientStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeNFSClientStatusRequest) SetInstanceId(v string) *DescribeNFSClientStatusRequest {
	s.InstanceId = &v
	return s
}

type DescribeNFSClientStatusResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeNFSClientStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Status    *string                                    `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeNFSClientStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNFSClientStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNFSClientStatusResponseBody) SetRequestId(v string) *DescribeNFSClientStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNFSClientStatusResponseBody) SetResult(v *DescribeNFSClientStatusResponseBodyResult) *DescribeNFSClientStatusResponseBody {
	s.Result = v
	return s
}

func (s *DescribeNFSClientStatusResponseBody) SetStatus(v string) *DescribeNFSClientStatusResponseBody {
	s.Status = &v
	return s
}

type DescribeNFSClientStatusResponseBodyResult struct {
	ExitCode           *int32  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	Output             *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s DescribeNFSClientStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeNFSClientStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeNFSClientStatusResponseBodyResult) SetExitCode(v int32) *DescribeNFSClientStatusResponseBodyResult {
	s.ExitCode = &v
	return s
}

func (s *DescribeNFSClientStatusResponseBodyResult) SetInvokeRecordStatus(v string) *DescribeNFSClientStatusResponseBodyResult {
	s.InvokeRecordStatus = &v
	return s
}

func (s *DescribeNFSClientStatusResponseBodyResult) SetOutput(v string) *DescribeNFSClientStatusResponseBodyResult {
	s.Output = &v
	return s
}

type DescribeNFSClientStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNFSClientStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNFSClientStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNFSClientStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeNFSClientStatusResponse) SetHeaders(v map[string]*string) *DescribeNFSClientStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeNFSClientStatusResponse) SetBody(v *DescribeNFSClientStatusResponseBody) *DescribeNFSClientStatusResponse {
	s.Body = v
	return s
}

type DescribePriceRequest struct {
	ChargeType  *string                            `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Commodities []*DescribePriceRequestCommodities `json:"Commodities,omitempty" xml:"Commodities,omitempty" type:"Repeated"`
	OrderType   *string                            `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PriceUnit   *string                            `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
}

func (s DescribePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) SetChargeType(v string) *DescribePriceRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribePriceRequest) SetCommodities(v []*DescribePriceRequestCommodities) *DescribePriceRequest {
	s.Commodities = v
	return s
}

func (s *DescribePriceRequest) SetOrderType(v string) *DescribePriceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribePriceRequest) SetPriceUnit(v string) *DescribePriceRequest {
	s.PriceUnit = &v
	return s
}

type DescribePriceRequestCommodities struct {
	Amount                     *int32                                      `json:"Amount,omitempty" xml:"Amount,omitempty"`
	DataDisks                  []*DescribePriceRequestCommoditiesDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	InstanceType               *string                                     `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetChargeType         *string                                     `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandWidthOut    *int32                                      `json:"InternetMaxBandWidthOut,omitempty" xml:"InternetMaxBandWidthOut,omitempty"`
	NetworkType                *string                                     `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	NodeType                   *string                                     `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Period                     *int32                                      `json:"Period,omitempty" xml:"Period,omitempty"`
	SystemDiskCategory         *string                                     `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	SystemDiskPerformanceLevel *string                                     `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	SystemDiskSize             *int32                                      `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribePriceRequestCommodities) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequestCommodities) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestCommodities) SetAmount(v int32) *DescribePriceRequestCommodities {
	s.Amount = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetDataDisks(v []*DescribePriceRequestCommoditiesDataDisks) *DescribePriceRequestCommodities {
	s.DataDisks = v
	return s
}

func (s *DescribePriceRequestCommodities) SetInstanceType(v string) *DescribePriceRequestCommodities {
	s.InstanceType = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetInternetChargeType(v string) *DescribePriceRequestCommodities {
	s.InternetChargeType = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetInternetMaxBandWidthOut(v int32) *DescribePriceRequestCommodities {
	s.InternetMaxBandWidthOut = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetNetworkType(v string) *DescribePriceRequestCommodities {
	s.NetworkType = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetNodeType(v string) *DescribePriceRequestCommodities {
	s.NodeType = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetPeriod(v int32) *DescribePriceRequestCommodities {
	s.Period = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetSystemDiskCategory(v string) *DescribePriceRequestCommodities {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetSystemDiskPerformanceLevel(v string) *DescribePriceRequestCommodities {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetSystemDiskSize(v int32) *DescribePriceRequestCommodities {
	s.SystemDiskSize = &v
	return s
}

type DescribePriceRequestCommoditiesDataDisks struct {
	Category           *string `json:"category,omitempty" xml:"category,omitempty"`
	DeleteWithInstance *bool   `json:"deleteWithInstance,omitempty" xml:"deleteWithInstance,omitempty"`
	Encrypted          *bool   `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	PerformanceLevel   *string `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
	Size               *int32  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s DescribePriceRequestCommoditiesDataDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequestCommoditiesDataDisks) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestCommoditiesDataDisks) SetCategory(v string) *DescribePriceRequestCommoditiesDataDisks {
	s.Category = &v
	return s
}

func (s *DescribePriceRequestCommoditiesDataDisks) SetDeleteWithInstance(v bool) *DescribePriceRequestCommoditiesDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *DescribePriceRequestCommoditiesDataDisks) SetEncrypted(v bool) *DescribePriceRequestCommoditiesDataDisks {
	s.Encrypted = &v
	return s
}

func (s *DescribePriceRequestCommoditiesDataDisks) SetPerformanceLevel(v string) *DescribePriceRequestCommoditiesDataDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribePriceRequestCommoditiesDataDisks) SetSize(v int32) *DescribePriceRequestCommoditiesDataDisks {
	s.Size = &v
	return s
}

type DescribePriceResponseBody struct {
	Prices          *DescribePriceResponseBodyPrices `json:"Prices,omitempty" xml:"Prices,omitempty" type:"Struct"`
	RequestId       *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalTradePrice *float32                         `json:"TotalTradePrice,omitempty" xml:"TotalTradePrice,omitempty"`
}

func (s DescribePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBody) SetPrices(v *DescribePriceResponseBodyPrices) *DescribePriceResponseBody {
	s.Prices = v
	return s
}

func (s *DescribePriceResponseBody) SetRequestId(v string) *DescribePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePriceResponseBody) SetTotalTradePrice(v float32) *DescribePriceResponseBody {
	s.TotalTradePrice = &v
	return s
}

type DescribePriceResponseBodyPrices struct {
	PriceInfo []*DescribePriceResponseBodyPricesPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyPrices) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyPrices) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPrices) SetPriceInfo(v []*DescribePriceResponseBodyPricesPriceInfo) *DescribePriceResponseBodyPrices {
	s.PriceInfo = v
	return s
}

type DescribePriceResponseBodyPricesPriceInfo struct {
	Currency      *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	NodeType      *string  `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	TradePrice    *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribePriceResponseBodyPricesPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyPricesPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPricesPriceInfo) SetCurrency(v string) *DescribePriceResponseBodyPricesPriceInfo {
	s.Currency = &v
	return s
}

func (s *DescribePriceResponseBodyPricesPriceInfo) SetNodeType(v string) *DescribePriceResponseBodyPricesPriceInfo {
	s.NodeType = &v
	return s
}

func (s *DescribePriceResponseBodyPricesPriceInfo) SetOriginalPrice(v float32) *DescribePriceResponseBodyPricesPriceInfo {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPricesPriceInfo) SetTradePrice(v float32) *DescribePriceResponseBodyPricesPriceInfo {
	s.TradePrice = &v
	return s
}

type DescribePriceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribePriceResponse) SetHeaders(v map[string]*string) *DescribePriceResponse {
	s.Headers = v
	return s
}

func (s *DescribePriceResponse) SetBody(v *DescribePriceResponseBody) *DescribePriceResponse {
	s.Body = v
	return s
}

type EditJobTemplateRequest struct {
	ArrayRequest       *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	ClockTime          *string `json:"ClockTime,omitempty" xml:"ClockTime,omitempty"`
	CommandLine        *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	Gpu                *int32  `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	InputFileUrl       *string `json:"InputFileUrl,omitempty" xml:"InputFileUrl,omitempty"`
	Mem                *string `json:"Mem,omitempty" xml:"Mem,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Node               *int32  `json:"Node,omitempty" xml:"Node,omitempty"`
	PackagePath        *string `json:"PackagePath,omitempty" xml:"PackagePath,omitempty"`
	Priority           *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Queue              *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	ReRunable          *bool   `json:"ReRunable,omitempty" xml:"ReRunable,omitempty"`
	RunasUser          *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" xml:"StderrRedirectPath,omitempty"`
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" xml:"StdoutRedirectPath,omitempty"`
	Task               *int32  `json:"Task,omitempty" xml:"Task,omitempty"`
	TemplateId         *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Thread             *int32  `json:"Thread,omitempty" xml:"Thread,omitempty"`
	UnzipCmd           *string `json:"UnzipCmd,omitempty" xml:"UnzipCmd,omitempty"`
	Variables          *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
	WithUnzipCmd       *bool   `json:"WithUnzipCmd,omitempty" xml:"WithUnzipCmd,omitempty"`
}

func (s EditJobTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s EditJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *EditJobTemplateRequest) SetArrayRequest(v string) *EditJobTemplateRequest {
	s.ArrayRequest = &v
	return s
}

func (s *EditJobTemplateRequest) SetClockTime(v string) *EditJobTemplateRequest {
	s.ClockTime = &v
	return s
}

func (s *EditJobTemplateRequest) SetCommandLine(v string) *EditJobTemplateRequest {
	s.CommandLine = &v
	return s
}

func (s *EditJobTemplateRequest) SetGpu(v int32) *EditJobTemplateRequest {
	s.Gpu = &v
	return s
}

func (s *EditJobTemplateRequest) SetInputFileUrl(v string) *EditJobTemplateRequest {
	s.InputFileUrl = &v
	return s
}

func (s *EditJobTemplateRequest) SetMem(v string) *EditJobTemplateRequest {
	s.Mem = &v
	return s
}

func (s *EditJobTemplateRequest) SetName(v string) *EditJobTemplateRequest {
	s.Name = &v
	return s
}

func (s *EditJobTemplateRequest) SetNode(v int32) *EditJobTemplateRequest {
	s.Node = &v
	return s
}

func (s *EditJobTemplateRequest) SetPackagePath(v string) *EditJobTemplateRequest {
	s.PackagePath = &v
	return s
}

func (s *EditJobTemplateRequest) SetPriority(v int32) *EditJobTemplateRequest {
	s.Priority = &v
	return s
}

func (s *EditJobTemplateRequest) SetQueue(v string) *EditJobTemplateRequest {
	s.Queue = &v
	return s
}

func (s *EditJobTemplateRequest) SetReRunable(v bool) *EditJobTemplateRequest {
	s.ReRunable = &v
	return s
}

func (s *EditJobTemplateRequest) SetRunasUser(v string) *EditJobTemplateRequest {
	s.RunasUser = &v
	return s
}

func (s *EditJobTemplateRequest) SetStderrRedirectPath(v string) *EditJobTemplateRequest {
	s.StderrRedirectPath = &v
	return s
}

func (s *EditJobTemplateRequest) SetStdoutRedirectPath(v string) *EditJobTemplateRequest {
	s.StdoutRedirectPath = &v
	return s
}

func (s *EditJobTemplateRequest) SetTask(v int32) *EditJobTemplateRequest {
	s.Task = &v
	return s
}

func (s *EditJobTemplateRequest) SetTemplateId(v string) *EditJobTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *EditJobTemplateRequest) SetThread(v int32) *EditJobTemplateRequest {
	s.Thread = &v
	return s
}

func (s *EditJobTemplateRequest) SetUnzipCmd(v string) *EditJobTemplateRequest {
	s.UnzipCmd = &v
	return s
}

func (s *EditJobTemplateRequest) SetVariables(v string) *EditJobTemplateRequest {
	s.Variables = &v
	return s
}

func (s *EditJobTemplateRequest) SetWithUnzipCmd(v bool) *EditJobTemplateRequest {
	s.WithUnzipCmd = &v
	return s
}

type EditJobTemplateResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s EditJobTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *EditJobTemplateResponseBody) SetRequestId(v string) *EditJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *EditJobTemplateResponseBody) SetTemplateId(v string) *EditJobTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type EditJobTemplateResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditJobTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditJobTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s EditJobTemplateResponse) GoString() string {
	return s.String()
}

func (s *EditJobTemplateResponse) SetHeaders(v map[string]*string) *EditJobTemplateResponse {
	s.Headers = v
	return s
}

func (s *EditJobTemplateResponse) SetBody(v *EditJobTemplateResponseBody) *EditJobTemplateResponse {
	s.Body = v
	return s
}

type GetAccountingReportRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Dim         *string `json:"Dim,omitempty" xml:"Dim,omitempty"`
	EndTime     *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	JobId       *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ReportType  *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	StartTime   *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAccountingReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountingReportRequest) GoString() string {
	return s.String()
}

func (s *GetAccountingReportRequest) SetClusterId(v string) *GetAccountingReportRequest {
	s.ClusterId = &v
	return s
}

func (s *GetAccountingReportRequest) SetDim(v string) *GetAccountingReportRequest {
	s.Dim = &v
	return s
}

func (s *GetAccountingReportRequest) SetEndTime(v int32) *GetAccountingReportRequest {
	s.EndTime = &v
	return s
}

func (s *GetAccountingReportRequest) SetFilterValue(v string) *GetAccountingReportRequest {
	s.FilterValue = &v
	return s
}

func (s *GetAccountingReportRequest) SetJobId(v string) *GetAccountingReportRequest {
	s.JobId = &v
	return s
}

func (s *GetAccountingReportRequest) SetPageNumber(v int32) *GetAccountingReportRequest {
	s.PageNumber = &v
	return s
}

func (s *GetAccountingReportRequest) SetPageSize(v int32) *GetAccountingReportRequest {
	s.PageSize = &v
	return s
}

func (s *GetAccountingReportRequest) SetReportType(v string) *GetAccountingReportRequest {
	s.ReportType = &v
	return s
}

func (s *GetAccountingReportRequest) SetStartTime(v int32) *GetAccountingReportRequest {
	s.StartTime = &v
	return s
}

type GetAccountingReportResponseBody struct {
	Data          *GetAccountingReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Metrics       *string                              `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	PageNumber    *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCoreTime *int32                               `json:"TotalCoreTime,omitempty" xml:"TotalCoreTime,omitempty"`
	TotalCount    *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAccountingReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountingReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountingReportResponseBody) SetData(v *GetAccountingReportResponseBodyData) *GetAccountingReportResponseBody {
	s.Data = v
	return s
}

func (s *GetAccountingReportResponseBody) SetMetrics(v string) *GetAccountingReportResponseBody {
	s.Metrics = &v
	return s
}

func (s *GetAccountingReportResponseBody) SetPageNumber(v int32) *GetAccountingReportResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetAccountingReportResponseBody) SetPageSize(v int32) *GetAccountingReportResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetAccountingReportResponseBody) SetRequestId(v string) *GetAccountingReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountingReportResponseBody) SetTotalCoreTime(v int32) *GetAccountingReportResponseBody {
	s.TotalCoreTime = &v
	return s
}

func (s *GetAccountingReportResponseBody) SetTotalCount(v int32) *GetAccountingReportResponseBody {
	s.TotalCount = &v
	return s
}

type GetAccountingReportResponseBodyData struct {
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetAccountingReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAccountingReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAccountingReportResponseBodyData) SetData(v []*string) *GetAccountingReportResponseBodyData {
	s.Data = v
	return s
}

type GetAccountingReportResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAccountingReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccountingReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountingReportResponse) GoString() string {
	return s.String()
}

func (s *GetAccountingReportResponse) SetHeaders(v map[string]*string) *GetAccountingReportResponse {
	s.Headers = v
	return s
}

func (s *GetAccountingReportResponse) SetBody(v *GetAccountingReportResponseBody) *GetAccountingReportResponse {
	s.Body = v
	return s
}

type GetAutoScaleConfigRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetAutoScaleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigRequest) SetClusterId(v string) *GetAutoScaleConfigRequest {
	s.ClusterId = &v
	return s
}

type GetAutoScaleConfigResponseBody struct {
	ClusterId               *string                               `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterType             *string                               `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	EnableAutoGrow          *bool                                 `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	EnableAutoShrink        *bool                                 `json:"EnableAutoShrink,omitempty" xml:"EnableAutoShrink,omitempty"`
	ExcludeNodes            *string                               `json:"ExcludeNodes,omitempty" xml:"ExcludeNodes,omitempty"`
	ExtraNodesGrowRatio     *int32                                `json:"ExtraNodesGrowRatio,omitempty" xml:"ExtraNodesGrowRatio,omitempty"`
	GrowIntervalInMinutes   *int32                                `json:"GrowIntervalInMinutes,omitempty" xml:"GrowIntervalInMinutes,omitempty"`
	GrowRatio               *int32                                `json:"GrowRatio,omitempty" xml:"GrowRatio,omitempty"`
	GrowTimeoutInMinutes    *int32                                `json:"GrowTimeoutInMinutes,omitempty" xml:"GrowTimeoutInMinutes,omitempty"`
	ImageId                 *string                               `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	MaxNodesInCluster       *int32                                `json:"MaxNodesInCluster,omitempty" xml:"MaxNodesInCluster,omitempty"`
	Queues                  *GetAutoScaleConfigResponseBodyQueues `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Struct"`
	RequestId               *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShrinkIdleTimes         *int32                                `json:"ShrinkIdleTimes,omitempty" xml:"ShrinkIdleTimes,omitempty"`
	ShrinkIntervalInMinutes *int32                                `json:"ShrinkIntervalInMinutes,omitempty" xml:"ShrinkIntervalInMinutes,omitempty"`
	SpotPriceLimit          *float32                              `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy            *string                               `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	Uid                     *string                               `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetAutoScaleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBody) SetClusterId(v string) *GetAutoScaleConfigResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetClusterType(v string) *GetAutoScaleConfigResponseBody {
	s.ClusterType = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetEnableAutoGrow(v bool) *GetAutoScaleConfigResponseBody {
	s.EnableAutoGrow = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetEnableAutoShrink(v bool) *GetAutoScaleConfigResponseBody {
	s.EnableAutoShrink = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetExcludeNodes(v string) *GetAutoScaleConfigResponseBody {
	s.ExcludeNodes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetExtraNodesGrowRatio(v int32) *GetAutoScaleConfigResponseBody {
	s.ExtraNodesGrowRatio = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetGrowIntervalInMinutes(v int32) *GetAutoScaleConfigResponseBody {
	s.GrowIntervalInMinutes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetGrowRatio(v int32) *GetAutoScaleConfigResponseBody {
	s.GrowRatio = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetGrowTimeoutInMinutes(v int32) *GetAutoScaleConfigResponseBody {
	s.GrowTimeoutInMinutes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetImageId(v string) *GetAutoScaleConfigResponseBody {
	s.ImageId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetMaxNodesInCluster(v int32) *GetAutoScaleConfigResponseBody {
	s.MaxNodesInCluster = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetQueues(v *GetAutoScaleConfigResponseBodyQueues) *GetAutoScaleConfigResponseBody {
	s.Queues = v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetRequestId(v string) *GetAutoScaleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetShrinkIdleTimes(v int32) *GetAutoScaleConfigResponseBody {
	s.ShrinkIdleTimes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetShrinkIntervalInMinutes(v int32) *GetAutoScaleConfigResponseBody {
	s.ShrinkIntervalInMinutes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetSpotPriceLimit(v float32) *GetAutoScaleConfigResponseBody {
	s.SpotPriceLimit = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetSpotStrategy(v string) *GetAutoScaleConfigResponseBody {
	s.SpotStrategy = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetUid(v string) *GetAutoScaleConfigResponseBody {
	s.Uid = &v
	return s
}

type GetAutoScaleConfigResponseBodyQueues struct {
	QueueInfo []*GetAutoScaleConfigResponseBodyQueuesQueueInfo `json:"QueueInfo,omitempty" xml:"QueueInfo,omitempty" type:"Repeated"`
}

func (s GetAutoScaleConfigResponseBodyQueues) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBodyQueues) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBodyQueues) SetQueueInfo(v []*GetAutoScaleConfigResponseBodyQueuesQueueInfo) *GetAutoScaleConfigResponseBodyQueues {
	s.QueueInfo = v
	return s
}

type GetAutoScaleConfigResponseBodyQueuesQueueInfo struct {
	DataDisks          *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisks     `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Struct"`
	EnableAutoGrow     *bool                                                       `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	EnableAutoShrink   *bool                                                       `json:"EnableAutoShrink,omitempty" xml:"EnableAutoShrink,omitempty"`
	HostNamePrefix     *string                                                     `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	HostNameSuffix     *string                                                     `json:"HostNameSuffix,omitempty" xml:"HostNameSuffix,omitempty"`
	InstanceType       *string                                                     `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceTypes      *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Struct"`
	MaxNodesInQueue    *int32                                                      `json:"MaxNodesInQueue,omitempty" xml:"MaxNodesInQueue,omitempty"`
	MinNodesInQueue    *int32                                                      `json:"MinNodesInQueue,omitempty" xml:"MinNodesInQueue,omitempty"`
	QueueImageId       *string                                                     `json:"QueueImageId,omitempty" xml:"QueueImageId,omitempty"`
	QueueName          *string                                                     `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	ResourceGroupId    *string                                                     `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SpotPriceLimit     *float32                                                    `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy       *string                                                     `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDiskCategory *string                                                     `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	SystemDiskLevel    *string                                                     `json:"SystemDiskLevel,omitempty" xml:"SystemDiskLevel,omitempty"`
	SystemDiskSize     *int32                                                      `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfo) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetDataDisks(v *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisks) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.DataDisks = v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetEnableAutoGrow(v bool) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.EnableAutoGrow = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetEnableAutoShrink(v bool) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.EnableAutoShrink = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetHostNamePrefix(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.HostNamePrefix = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetHostNameSuffix(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.HostNameSuffix = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetInstanceType(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.InstanceType = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetInstanceTypes(v *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypes) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.InstanceTypes = v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetMaxNodesInQueue(v int32) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.MaxNodesInQueue = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetMinNodesInQueue(v int32) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.MinNodesInQueue = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetQueueImageId(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.QueueImageId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetQueueName(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.QueueName = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetResourceGroupId(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetSpotPriceLimit(v float32) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.SpotPriceLimit = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetSpotStrategy(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.SpotStrategy = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetSystemDiskCategory(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.SystemDiskCategory = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetSystemDiskLevel(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.SystemDiskLevel = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfo) SetSystemDiskSize(v int32) *GetAutoScaleConfigResponseBodyQueuesQueueInfo {
	s.SystemDiskSize = &v
	return s
}

type GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisks struct {
	DataDisksInfo []*GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo `json:"DataDisksInfo,omitempty" xml:"DataDisksInfo,omitempty" type:"Repeated"`
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisks) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisks) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisks) SetDataDisksInfo(v []*GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisks {
	s.DataDisksInfo = v
	return s
}

type GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo struct {
	DataDiskCategory           *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	DataDiskDeleteWithInstance *bool   `json:"DataDiskDeleteWithInstance,omitempty" xml:"DataDiskDeleteWithInstance,omitempty"`
	DataDiskEncrypted          *bool   `json:"DataDiskEncrypted,omitempty" xml:"DataDiskEncrypted,omitempty"`
	DataDiskKMSKeyId           *string `json:"DataDiskKMSKeyId,omitempty" xml:"DataDiskKMSKeyId,omitempty"`
	DataDiskPerformanceLevel   *string `json:"DataDiskPerformanceLevel,omitempty" xml:"DataDiskPerformanceLevel,omitempty"`
	DataDiskSize               *int32  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) SetDataDiskCategory(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo {
	s.DataDiskCategory = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) SetDataDiskDeleteWithInstance(v bool) *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo {
	s.DataDiskDeleteWithInstance = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) SetDataDiskEncrypted(v bool) *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo {
	s.DataDiskEncrypted = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) SetDataDiskKMSKeyId(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo {
	s.DataDiskKMSKeyId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) SetDataDiskPerformanceLevel(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo {
	s.DataDiskPerformanceLevel = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo) SetDataDiskSize(v int32) *GetAutoScaleConfigResponseBodyQueuesQueueInfoDataDisksDataDisksInfo {
	s.DataDiskSize = &v
	return s
}

type GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypes struct {
	InstanceTypeInfo []*GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo `json:"InstanceTypeInfo,omitempty" xml:"InstanceTypeInfo,omitempty" type:"Repeated"`
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypes) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypes) SetInstanceTypeInfo(v []*GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypes {
	s.InstanceTypeInfo = v
	return s
}

type GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo struct {
	HostNamePrefix *string  `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	InstanceType   *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy   *string  `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	VSwitchId      *string  `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId         *string  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) SetHostNamePrefix(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo {
	s.HostNamePrefix = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) SetInstanceType(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo {
	s.InstanceType = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) SetSpotPriceLimit(v float32) *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo {
	s.SpotPriceLimit = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) SetSpotStrategy(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo {
	s.SpotStrategy = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) SetVSwitchId(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo {
	s.VSwitchId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo) SetZoneId(v string) *GetAutoScaleConfigResponseBodyQueuesQueueInfoInstanceTypesInstanceTypeInfo {
	s.ZoneId = &v
	return s
}

type GetAutoScaleConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAutoScaleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAutoScaleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponse) SetHeaders(v map[string]*string) *GetAutoScaleConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAutoScaleConfigResponse) SetBody(v *GetAutoScaleConfigResponseBody) *GetAutoScaleConfigResponse {
	s.Body = v
	return s
}

type GetCloudMetricLogsRequest struct {
	AggregationInterval *int32  `json:"AggregationInterval,omitempty" xml:"AggregationInterval,omitempty"`
	AggregationType     *string `json:"AggregationType,omitempty" xml:"AggregationType,omitempty"`
	ClusterId           *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Filter              *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	From                *int32  `json:"From,omitempty" xml:"From,omitempty"`
	MetricCategories    *string `json:"MetricCategories,omitempty" xml:"MetricCategories,omitempty"`
	MetricScope         *string `json:"MetricScope,omitempty" xml:"MetricScope,omitempty"`
	Reverse             *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	To                  *int32  `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetCloudMetricLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricLogsRequest) GoString() string {
	return s.String()
}

func (s *GetCloudMetricLogsRequest) SetAggregationInterval(v int32) *GetCloudMetricLogsRequest {
	s.AggregationInterval = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetAggregationType(v string) *GetCloudMetricLogsRequest {
	s.AggregationType = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetClusterId(v string) *GetCloudMetricLogsRequest {
	s.ClusterId = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetFilter(v string) *GetCloudMetricLogsRequest {
	s.Filter = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetFrom(v int32) *GetCloudMetricLogsRequest {
	s.From = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetMetricCategories(v string) *GetCloudMetricLogsRequest {
	s.MetricCategories = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetMetricScope(v string) *GetCloudMetricLogsRequest {
	s.MetricScope = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetReverse(v bool) *GetCloudMetricLogsRequest {
	s.Reverse = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetTo(v int32) *GetCloudMetricLogsRequest {
	s.To = &v
	return s
}

type GetCloudMetricLogsResponseBody struct {
	MetricLogs *GetCloudMetricLogsResponseBodyMetricLogs `json:"MetricLogs,omitempty" xml:"MetricLogs,omitempty" type:"Struct"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCloudMetricLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricLogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudMetricLogsResponseBody) SetMetricLogs(v *GetCloudMetricLogsResponseBodyMetricLogs) *GetCloudMetricLogsResponseBody {
	s.MetricLogs = v
	return s
}

func (s *GetCloudMetricLogsResponseBody) SetRequestId(v string) *GetCloudMetricLogsResponseBody {
	s.RequestId = &v
	return s
}

type GetCloudMetricLogsResponseBodyMetricLogs struct {
	MetricLog []*GetCloudMetricLogsResponseBodyMetricLogsMetricLog `json:"MetricLog,omitempty" xml:"MetricLog,omitempty" type:"Repeated"`
}

func (s GetCloudMetricLogsResponseBodyMetricLogs) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricLogsResponseBodyMetricLogs) GoString() string {
	return s.String()
}

func (s *GetCloudMetricLogsResponseBodyMetricLogs) SetMetricLog(v []*GetCloudMetricLogsResponseBodyMetricLogsMetricLog) *GetCloudMetricLogsResponseBodyMetricLogs {
	s.MetricLog = v
	return s
}

type GetCloudMetricLogsResponseBodyMetricLogsMetricLog struct {
	DiskDevice       *string `json:"DiskDevice,omitempty" xml:"DiskDevice,omitempty"`
	Hostname         *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MetricData       *string `json:"MetricData,omitempty" xml:"MetricData,omitempty"`
	NetworkInterface *string `json:"NetworkInterface,omitempty" xml:"NetworkInterface,omitempty"`
	Time             *int32  `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetCloudMetricLogsResponseBodyMetricLogsMetricLog) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricLogsResponseBodyMetricLogsMetricLog) GoString() string {
	return s.String()
}

func (s *GetCloudMetricLogsResponseBodyMetricLogsMetricLog) SetDiskDevice(v string) *GetCloudMetricLogsResponseBodyMetricLogsMetricLog {
	s.DiskDevice = &v
	return s
}

func (s *GetCloudMetricLogsResponseBodyMetricLogsMetricLog) SetHostname(v string) *GetCloudMetricLogsResponseBodyMetricLogsMetricLog {
	s.Hostname = &v
	return s
}

func (s *GetCloudMetricLogsResponseBodyMetricLogsMetricLog) SetInstanceId(v string) *GetCloudMetricLogsResponseBodyMetricLogsMetricLog {
	s.InstanceId = &v
	return s
}

func (s *GetCloudMetricLogsResponseBodyMetricLogsMetricLog) SetMetricData(v string) *GetCloudMetricLogsResponseBodyMetricLogsMetricLog {
	s.MetricData = &v
	return s
}

func (s *GetCloudMetricLogsResponseBodyMetricLogsMetricLog) SetNetworkInterface(v string) *GetCloudMetricLogsResponseBodyMetricLogsMetricLog {
	s.NetworkInterface = &v
	return s
}

func (s *GetCloudMetricLogsResponseBodyMetricLogsMetricLog) SetTime(v int32) *GetCloudMetricLogsResponseBodyMetricLogsMetricLog {
	s.Time = &v
	return s
}

type GetCloudMetricLogsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCloudMetricLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCloudMetricLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricLogsResponse) GoString() string {
	return s.String()
}

func (s *GetCloudMetricLogsResponse) SetHeaders(v map[string]*string) *GetCloudMetricLogsResponse {
	s.Headers = v
	return s
}

func (s *GetCloudMetricLogsResponse) SetBody(v *GetCloudMetricLogsResponseBody) *GetCloudMetricLogsResponse {
	s.Body = v
	return s
}

type GetCloudMetricProfilingRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ProfilingId *string `json:"ProfilingId,omitempty" xml:"ProfilingId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetCloudMetricProfilingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricProfilingRequest) GoString() string {
	return s.String()
}

func (s *GetCloudMetricProfilingRequest) SetClusterId(v string) *GetCloudMetricProfilingRequest {
	s.ClusterId = &v
	return s
}

func (s *GetCloudMetricProfilingRequest) SetProfilingId(v string) *GetCloudMetricProfilingRequest {
	s.ProfilingId = &v
	return s
}

func (s *GetCloudMetricProfilingRequest) SetRegionId(v string) *GetCloudMetricProfilingRequest {
	s.RegionId = &v
	return s
}

type GetCloudMetricProfilingResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SvgUrls   *GetCloudMetricProfilingResponseBodySvgUrls `json:"SvgUrls,omitempty" xml:"SvgUrls,omitempty" type:"Struct"`
}

func (s GetCloudMetricProfilingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricProfilingResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudMetricProfilingResponseBody) SetRequestId(v string) *GetCloudMetricProfilingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudMetricProfilingResponseBody) SetSvgUrls(v *GetCloudMetricProfilingResponseBodySvgUrls) *GetCloudMetricProfilingResponseBody {
	s.SvgUrls = v
	return s
}

type GetCloudMetricProfilingResponseBodySvgUrls struct {
	SvgInfo []*GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo `json:"SvgInfo,omitempty" xml:"SvgInfo,omitempty" type:"Repeated"`
}

func (s GetCloudMetricProfilingResponseBodySvgUrls) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricProfilingResponseBodySvgUrls) GoString() string {
	return s.String()
}

func (s *GetCloudMetricProfilingResponseBodySvgUrls) SetSvgInfo(v []*GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo) *GetCloudMetricProfilingResponseBodySvgUrls {
	s.SvgInfo = v
	return s
}

type GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Size *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo) GoString() string {
	return s.String()
}

func (s *GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo) SetName(v string) *GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo {
	s.Name = &v
	return s
}

func (s *GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo) SetSize(v int32) *GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo {
	s.Size = &v
	return s
}

func (s *GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo) SetType(v string) *GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo {
	s.Type = &v
	return s
}

func (s *GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo) SetUrl(v string) *GetCloudMetricProfilingResponseBodySvgUrlsSvgInfo {
	s.Url = &v
	return s
}

type GetCloudMetricProfilingResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCloudMetricProfilingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCloudMetricProfilingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricProfilingResponse) GoString() string {
	return s.String()
}

func (s *GetCloudMetricProfilingResponse) SetHeaders(v map[string]*string) *GetCloudMetricProfilingResponse {
	s.Headers = v
	return s
}

func (s *GetCloudMetricProfilingResponse) SetBody(v *GetCloudMetricProfilingResponseBody) *GetCloudMetricProfilingResponse {
	s.Body = v
	return s
}

type GetClusterVolumesRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetClusterVolumesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesRequest) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesRequest) SetClusterId(v string) *GetClusterVolumesRequest {
	s.ClusterId = &v
	return s
}

type GetClusterVolumesResponseBody struct {
	RegionId  *string                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Volumes   *GetClusterVolumesResponseBodyVolumes `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Struct"`
}

func (s GetClusterVolumesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesResponseBody) SetRegionId(v string) *GetClusterVolumesResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetClusterVolumesResponseBody) SetRequestId(v string) *GetClusterVolumesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterVolumesResponseBody) SetVolumes(v *GetClusterVolumesResponseBodyVolumes) *GetClusterVolumesResponseBody {
	s.Volumes = v
	return s
}

type GetClusterVolumesResponseBodyVolumes struct {
	VolumeInfo []*GetClusterVolumesResponseBodyVolumesVolumeInfo `json:"VolumeInfo,omitempty" xml:"VolumeInfo,omitempty" type:"Repeated"`
}

func (s GetClusterVolumesResponseBodyVolumes) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesResponseBodyVolumes) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesResponseBodyVolumes) SetVolumeInfo(v []*GetClusterVolumesResponseBodyVolumesVolumeInfo) *GetClusterVolumesResponseBodyVolumes {
	s.VolumeInfo = v
	return s
}

type GetClusterVolumesResponseBodyVolumesVolumeInfo struct {
	JobQueue         *string                                              `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	LocalDirectory   *string                                              `json:"LocalDirectory,omitempty" xml:"LocalDirectory,omitempty"`
	Location         *string                                              `json:"Location,omitempty" xml:"Location,omitempty"`
	MustKeep         *bool                                                `json:"MustKeep,omitempty" xml:"MustKeep,omitempty"`
	RemoteDirectory  *string                                              `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	Roles            *GetClusterVolumesResponseBodyVolumesVolumeInfoRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
	VolumeId         *string                                              `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	VolumeMountpoint *string                                              `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	VolumeProtocol   *string                                              `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	VolumeType       *string                                              `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
}

func (s GetClusterVolumesResponseBodyVolumesVolumeInfo) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesResponseBodyVolumesVolumeInfo) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetJobQueue(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.JobQueue = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetLocalDirectory(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.LocalDirectory = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetLocation(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.Location = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetMustKeep(v bool) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.MustKeep = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetRemoteDirectory(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.RemoteDirectory = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetRoles(v *GetClusterVolumesResponseBodyVolumesVolumeInfoRoles) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.Roles = v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetVolumeId(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeId = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetVolumeMountpoint(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeMountpoint = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetVolumeProtocol(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeProtocol = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfo) SetVolumeType(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeType = &v
	return s
}

type GetClusterVolumesResponseBodyVolumesVolumeInfoRoles struct {
	RoleInfo []*GetClusterVolumesResponseBodyVolumesVolumeInfoRolesRoleInfo `json:"RoleInfo,omitempty" xml:"RoleInfo,omitempty" type:"Repeated"`
}

func (s GetClusterVolumesResponseBodyVolumesVolumeInfoRoles) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesResponseBodyVolumesVolumeInfoRoles) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfoRoles) SetRoleInfo(v []*GetClusterVolumesResponseBodyVolumesVolumeInfoRolesRoleInfo) *GetClusterVolumesResponseBodyVolumesVolumeInfoRoles {
	s.RoleInfo = v
	return s
}

type GetClusterVolumesResponseBodyVolumesVolumeInfoRolesRoleInfo struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetClusterVolumesResponseBodyVolumesVolumeInfoRolesRoleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesResponseBodyVolumesVolumeInfoRolesRoleInfo) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesResponseBodyVolumesVolumeInfoRolesRoleInfo) SetName(v string) *GetClusterVolumesResponseBodyVolumesVolumeInfoRolesRoleInfo {
	s.Name = &v
	return s
}

type GetClusterVolumesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetClusterVolumesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClusterVolumesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesResponse) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesResponse) SetHeaders(v map[string]*string) *GetClusterVolumesResponse {
	s.Headers = v
	return s
}

func (s *GetClusterVolumesResponse) SetBody(v *GetClusterVolumesResponseBody) *GetClusterVolumesResponse {
	s.Body = v
	return s
}

type GetGWSConnectTicketRequest struct {
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetGWSConnectTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGWSConnectTicketRequest) GoString() string {
	return s.String()
}

func (s *GetGWSConnectTicketRequest) SetAppName(v string) *GetGWSConnectTicketRequest {
	s.AppName = &v
	return s
}

func (s *GetGWSConnectTicketRequest) SetInstanceId(v string) *GetGWSConnectTicketRequest {
	s.InstanceId = &v
	return s
}

type GetGWSConnectTicketResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Ticket    *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s GetGWSConnectTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGWSConnectTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetGWSConnectTicketResponseBody) SetRequestId(v string) *GetGWSConnectTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGWSConnectTicketResponseBody) SetTicket(v string) *GetGWSConnectTicketResponseBody {
	s.Ticket = &v
	return s
}

type GetGWSConnectTicketResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGWSConnectTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGWSConnectTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGWSConnectTicketResponse) GoString() string {
	return s.String()
}

func (s *GetGWSConnectTicketResponse) SetHeaders(v map[string]*string) *GetGWSConnectTicketResponse {
	s.Headers = v
	return s
}

func (s *GetGWSConnectTicketResponse) SetBody(v *GetGWSConnectTicketResponseBody) *GetGWSConnectTicketResponse {
	s.Body = v
	return s
}

type GetHybridClusterConfigRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Node      *string `json:"Node,omitempty" xml:"Node,omitempty"`
}

func (s GetHybridClusterConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHybridClusterConfigRequest) GoString() string {
	return s.String()
}

func (s *GetHybridClusterConfigRequest) SetClusterId(v string) *GetHybridClusterConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *GetHybridClusterConfigRequest) SetNode(v string) *GetHybridClusterConfigRequest {
	s.Node = &v
	return s
}

type GetHybridClusterConfigResponseBody struct {
	ClusterConfig *string `json:"ClusterConfig,omitempty" xml:"ClusterConfig,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHybridClusterConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHybridClusterConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetHybridClusterConfigResponseBody) SetClusterConfig(v string) *GetHybridClusterConfigResponseBody {
	s.ClusterConfig = &v
	return s
}

func (s *GetHybridClusterConfigResponseBody) SetRequestId(v string) *GetHybridClusterConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetHybridClusterConfigResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHybridClusterConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHybridClusterConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHybridClusterConfigResponse) GoString() string {
	return s.String()
}

func (s *GetHybridClusterConfigResponse) SetHeaders(v map[string]*string) *GetHybridClusterConfigResponse {
	s.Headers = v
	return s
}

func (s *GetHybridClusterConfigResponse) SetBody(v *GetHybridClusterConfigResponseBody) *GetHybridClusterConfigResponse {
	s.Body = v
	return s
}

type GetIfEcsTypeSupportHtConfigRequest struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s GetIfEcsTypeSupportHtConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIfEcsTypeSupportHtConfigRequest) GoString() string {
	return s.String()
}

func (s *GetIfEcsTypeSupportHtConfigRequest) SetInstanceType(v string) *GetIfEcsTypeSupportHtConfigRequest {
	s.InstanceType = &v
	return s
}

type GetIfEcsTypeSupportHtConfigResponseBody struct {
	DefaultHtEnabled *bool   `json:"DefaultHtEnabled,omitempty" xml:"DefaultHtEnabled,omitempty"`
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SupportHtConfig  *bool   `json:"SupportHtConfig,omitempty" xml:"SupportHtConfig,omitempty"`
}

func (s GetIfEcsTypeSupportHtConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIfEcsTypeSupportHtConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetIfEcsTypeSupportHtConfigResponseBody) SetDefaultHtEnabled(v bool) *GetIfEcsTypeSupportHtConfigResponseBody {
	s.DefaultHtEnabled = &v
	return s
}

func (s *GetIfEcsTypeSupportHtConfigResponseBody) SetInstanceType(v string) *GetIfEcsTypeSupportHtConfigResponseBody {
	s.InstanceType = &v
	return s
}

func (s *GetIfEcsTypeSupportHtConfigResponseBody) SetRequestId(v string) *GetIfEcsTypeSupportHtConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIfEcsTypeSupportHtConfigResponseBody) SetSupportHtConfig(v bool) *GetIfEcsTypeSupportHtConfigResponseBody {
	s.SupportHtConfig = &v
	return s
}

type GetIfEcsTypeSupportHtConfigResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIfEcsTypeSupportHtConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIfEcsTypeSupportHtConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIfEcsTypeSupportHtConfigResponse) GoString() string {
	return s.String()
}

func (s *GetIfEcsTypeSupportHtConfigResponse) SetHeaders(v map[string]*string) *GetIfEcsTypeSupportHtConfigResponse {
	s.Headers = v
	return s
}

func (s *GetIfEcsTypeSupportHtConfigResponse) SetBody(v *GetIfEcsTypeSupportHtConfigResponseBody) *GetIfEcsTypeSupportHtConfigResponse {
	s.Body = v
	return s
}

type GetSchedulerInfoRequest struct {
	ClusterId *string                             `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Scheduler []*GetSchedulerInfoRequestScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Repeated"`
}

func (s GetSchedulerInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSchedulerInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSchedulerInfoRequest) SetClusterId(v string) *GetSchedulerInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *GetSchedulerInfoRequest) SetRegionId(v string) *GetSchedulerInfoRequest {
	s.RegionId = &v
	return s
}

func (s *GetSchedulerInfoRequest) SetScheduler(v []*GetSchedulerInfoRequestScheduler) *GetSchedulerInfoRequest {
	s.Scheduler = v
	return s
}

type GetSchedulerInfoRequestScheduler struct {
	SchedName *string `json:"SchedName,omitempty" xml:"SchedName,omitempty"`
}

func (s GetSchedulerInfoRequestScheduler) String() string {
	return tea.Prettify(s)
}

func (s GetSchedulerInfoRequestScheduler) GoString() string {
	return s.String()
}

func (s *GetSchedulerInfoRequestScheduler) SetSchedName(v string) *GetSchedulerInfoRequestScheduler {
	s.SchedName = &v
	return s
}

type GetSchedulerInfoResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SchedInfo []*GetSchedulerInfoResponseBodySchedInfo `json:"SchedInfo,omitempty" xml:"SchedInfo,omitempty" type:"Repeated"`
}

func (s GetSchedulerInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSchedulerInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSchedulerInfoResponseBody) SetRequestId(v string) *GetSchedulerInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSchedulerInfoResponseBody) SetSchedInfo(v []*GetSchedulerInfoResponseBodySchedInfo) *GetSchedulerInfoResponseBody {
	s.SchedInfo = v
	return s
}

type GetSchedulerInfoResponseBodySchedInfo struct {
	Configuration *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	SchedName     *string `json:"SchedName,omitempty" xml:"SchedName,omitempty"`
}

func (s GetSchedulerInfoResponseBodySchedInfo) String() string {
	return tea.Prettify(s)
}

func (s GetSchedulerInfoResponseBodySchedInfo) GoString() string {
	return s.String()
}

func (s *GetSchedulerInfoResponseBodySchedInfo) SetConfiguration(v string) *GetSchedulerInfoResponseBodySchedInfo {
	s.Configuration = &v
	return s
}

func (s *GetSchedulerInfoResponseBodySchedInfo) SetSchedName(v string) *GetSchedulerInfoResponseBodySchedInfo {
	s.SchedName = &v
	return s
}

type GetSchedulerInfoResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSchedulerInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSchedulerInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSchedulerInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSchedulerInfoResponse) SetHeaders(v map[string]*string) *GetSchedulerInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSchedulerInfoResponse) SetBody(v *GetSchedulerInfoResponseBody) *GetSchedulerInfoResponse {
	s.Body = v
	return s
}

type GetVisualServiceStatusRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetVisualServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVisualServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetVisualServiceStatusRequest) SetClusterId(v string) *GetVisualServiceStatusRequest {
	s.ClusterId = &v
	return s
}

type GetVisualServiceStatusResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVisualServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVisualServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetVisualServiceStatusResponseBody) SetMessage(v string) *GetVisualServiceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetVisualServiceStatusResponseBody) SetRequestId(v string) *GetVisualServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetVisualServiceStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVisualServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVisualServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVisualServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetVisualServiceStatusResponse) SetHeaders(v map[string]*string) *GetVisualServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetVisualServiceStatusResponse) SetBody(v *GetVisualServiceStatusResponseBody) *GetVisualServiceStatusResponse {
	s.Body = v
	return s
}

type InitializeEHPCRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InitializeEHPCRequest) String() string {
	return tea.Prettify(s)
}

func (s InitializeEHPCRequest) GoString() string {
	return s.String()
}

func (s *InitializeEHPCRequest) SetRegionId(v string) *InitializeEHPCRequest {
	s.RegionId = &v
	return s
}

type InitializeEHPCResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitializeEHPCResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitializeEHPCResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeEHPCResponseBody) SetRequestId(v string) *InitializeEHPCResponseBody {
	s.RequestId = &v
	return s
}

type InitializeEHPCResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InitializeEHPCResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitializeEHPCResponse) String() string {
	return tea.Prettify(s)
}

func (s InitializeEHPCResponse) GoString() string {
	return s.String()
}

func (s *InitializeEHPCResponse) SetHeaders(v map[string]*string) *InitializeEHPCResponse {
	s.Headers = v
	return s
}

func (s *InitializeEHPCResponse) SetBody(v *InitializeEHPCResponseBody) *InitializeEHPCResponse {
	s.Body = v
	return s
}

type InstallSoftwareRequest struct {
	Application *string `json:"Application,omitempty" xml:"Application,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s InstallSoftwareRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallSoftwareRequest) GoString() string {
	return s.String()
}

func (s *InstallSoftwareRequest) SetApplication(v string) *InstallSoftwareRequest {
	s.Application = &v
	return s
}

func (s *InstallSoftwareRequest) SetClusterId(v string) *InstallSoftwareRequest {
	s.ClusterId = &v
	return s
}

type InstallSoftwareResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallSoftwareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallSoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *InstallSoftwareResponseBody) SetRequestId(v string) *InstallSoftwareResponseBody {
	s.RequestId = &v
	return s
}

type InstallSoftwareResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InstallSoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallSoftwareResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallSoftwareResponse) GoString() string {
	return s.String()
}

func (s *InstallSoftwareResponse) SetHeaders(v map[string]*string) *InstallSoftwareResponse {
	s.Headers = v
	return s
}

func (s *InstallSoftwareResponse) SetBody(v *InstallSoftwareResponseBody) *InstallSoftwareResponse {
	s.Body = v
	return s
}

type InvokeShellCommandRequest struct {
	ClusterId  *string                              `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Command    *string                              `json:"Command,omitempty" xml:"Command,omitempty"`
	Instance   []*InvokeShellCommandRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
	Timeout    *int32                               `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	WorkingDir *string                              `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s InvokeShellCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeShellCommandRequest) GoString() string {
	return s.String()
}

func (s *InvokeShellCommandRequest) SetClusterId(v string) *InvokeShellCommandRequest {
	s.ClusterId = &v
	return s
}

func (s *InvokeShellCommandRequest) SetCommand(v string) *InvokeShellCommandRequest {
	s.Command = &v
	return s
}

func (s *InvokeShellCommandRequest) SetInstance(v []*InvokeShellCommandRequestInstance) *InvokeShellCommandRequest {
	s.Instance = v
	return s
}

func (s *InvokeShellCommandRequest) SetTimeout(v int32) *InvokeShellCommandRequest {
	s.Timeout = &v
	return s
}

func (s *InvokeShellCommandRequest) SetWorkingDir(v string) *InvokeShellCommandRequest {
	s.WorkingDir = &v
	return s
}

type InvokeShellCommandRequestInstance struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s InvokeShellCommandRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s InvokeShellCommandRequestInstance) GoString() string {
	return s.String()
}

func (s *InvokeShellCommandRequestInstance) SetId(v string) *InvokeShellCommandRequestInstance {
	s.Id = &v
	return s
}

type InvokeShellCommandResponseBody struct {
	CommandId   *string                                    `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	InstanceIds *InvokeShellCommandResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	RequestId   *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InvokeShellCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeShellCommandResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeShellCommandResponseBody) SetCommandId(v string) *InvokeShellCommandResponseBody {
	s.CommandId = &v
	return s
}

func (s *InvokeShellCommandResponseBody) SetInstanceIds(v *InvokeShellCommandResponseBodyInstanceIds) *InvokeShellCommandResponseBody {
	s.InstanceIds = v
	return s
}

func (s *InvokeShellCommandResponseBody) SetRequestId(v string) *InvokeShellCommandResponseBody {
	s.RequestId = &v
	return s
}

type InvokeShellCommandResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s InvokeShellCommandResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s InvokeShellCommandResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *InvokeShellCommandResponseBodyInstanceIds) SetInstanceId(v []*string) *InvokeShellCommandResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

type InvokeShellCommandResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InvokeShellCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvokeShellCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeShellCommandResponse) GoString() string {
	return s.String()
}

func (s *InvokeShellCommandResponse) SetHeaders(v map[string]*string) *InvokeShellCommandResponse {
	s.Headers = v
	return s
}

func (s *InvokeShellCommandResponse) SetBody(v *InvokeShellCommandResponseBody) *InvokeShellCommandResponse {
	s.Body = v
	return s
}

type ListAvailableEcsTypesRequest struct {
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	ShowSoldOut        *bool   `json:"ShowSoldOut,omitempty" xml:"ShowSoldOut,omitempty"`
	SpotStrategy       *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	ZoneId             *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListAvailableEcsTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesRequest) SetInstanceChargeType(v string) *ListAvailableEcsTypesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *ListAvailableEcsTypesRequest) SetShowSoldOut(v bool) *ListAvailableEcsTypesRequest {
	s.ShowSoldOut = &v
	return s
}

func (s *ListAvailableEcsTypesRequest) SetSpotStrategy(v string) *ListAvailableEcsTypesRequest {
	s.SpotStrategy = &v
	return s
}

func (s *ListAvailableEcsTypesRequest) SetZoneId(v string) *ListAvailableEcsTypesRequest {
	s.ZoneId = &v
	return s
}

type ListAvailableEcsTypesResponseBody struct {
	InstanceTypeFamilies *ListAvailableEcsTypesResponseBodyInstanceTypeFamilies `json:"InstanceTypeFamilies,omitempty" xml:"InstanceTypeFamilies,omitempty" type:"Struct"`
	RequestId            *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SupportSpotInstance  *bool                                                  `json:"SupportSpotInstance,omitempty" xml:"SupportSpotInstance,omitempty"`
}

func (s ListAvailableEcsTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponseBody) SetInstanceTypeFamilies(v *ListAvailableEcsTypesResponseBodyInstanceTypeFamilies) *ListAvailableEcsTypesResponseBody {
	s.InstanceTypeFamilies = v
	return s
}

func (s *ListAvailableEcsTypesResponseBody) SetRequestId(v string) *ListAvailableEcsTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBody) SetSupportSpotInstance(v bool) *ListAvailableEcsTypesResponseBody {
	s.SupportSpotInstance = &v
	return s
}

type ListAvailableEcsTypesResponseBodyInstanceTypeFamilies struct {
	InstanceTypeFamilyInfo []*ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo `json:"InstanceTypeFamilyInfo,omitempty" xml:"InstanceTypeFamilyInfo,omitempty" type:"Repeated"`
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamilies) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamilies) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamilies) SetInstanceTypeFamilyInfo(v []*ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo) *ListAvailableEcsTypesResponseBodyInstanceTypeFamilies {
	s.InstanceTypeFamilyInfo = v
	return s
}

type ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo struct {
	Generation           *string                                                                           `json:"Generation,omitempty" xml:"Generation,omitempty"`
	InstanceTypeFamilyId *string                                                                           `json:"InstanceTypeFamilyId,omitempty" xml:"InstanceTypeFamilyId,omitempty"`
	Types                *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypes `json:"Types,omitempty" xml:"Types,omitempty" type:"Struct"`
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo) SetGeneration(v string) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo {
	s.Generation = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo) SetInstanceTypeFamilyId(v string) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo {
	s.InstanceTypeFamilyId = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo) SetTypes(v *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypes) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfo {
	s.Types = v
	return s
}

type ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypes struct {
	TypesInfo []*ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo `json:"TypesInfo,omitempty" xml:"TypesInfo,omitempty" type:"Repeated"`
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypes) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypes) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypes) SetTypesInfo(v []*ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypes {
	s.TypesInfo = v
	return s
}

type ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo struct {
	CpuCoreCount        *int32                                                                                            `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty"`
	EniQuantity         *int32                                                                                            `json:"EniQuantity,omitempty" xml:"EniQuantity,omitempty"`
	GPUAmount           *int32                                                                                            `json:"GPUAmount,omitempty" xml:"GPUAmount,omitempty"`
	GPUSpec             *string                                                                                           `json:"GPUSpec,omitempty" xml:"GPUSpec,omitempty"`
	InstanceBandwidthRx *int32                                                                                            `json:"InstanceBandwidthRx,omitempty" xml:"InstanceBandwidthRx,omitempty"`
	InstanceBandwidthTx *int32                                                                                            `json:"InstanceBandwidthTx,omitempty" xml:"InstanceBandwidthTx,omitempty"`
	InstancePpsRx       *int32                                                                                            `json:"InstancePpsRx,omitempty" xml:"InstancePpsRx,omitempty"`
	InstancePpsTx       *int32                                                                                            `json:"InstancePpsTx,omitempty" xml:"InstancePpsTx,omitempty"`
	InstanceTypeId      *string                                                                                           `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty"`
	MemorySize          *int32                                                                                            `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	Status              *string                                                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	ZoneIds             *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfoZoneIds `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty" type:"Struct"`
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetCpuCoreCount(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.CpuCoreCount = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetEniQuantity(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.EniQuantity = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetGPUAmount(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.GPUAmount = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetGPUSpec(v string) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.GPUSpec = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetInstanceBandwidthRx(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.InstanceBandwidthRx = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetInstanceBandwidthTx(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.InstanceBandwidthTx = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetInstancePpsRx(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.InstancePpsRx = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetInstancePpsTx(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.InstancePpsTx = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetInstanceTypeId(v string) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.InstanceTypeId = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetMemorySize(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.MemorySize = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetStatus(v string) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.Status = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo) SetZoneIds(v *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfoZoneIds) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfo {
	s.ZoneIds = v
	return s
}

type ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfoZoneIds struct {
	ZoneId []*string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" type:"Repeated"`
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfoZoneIds) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfoZoneIds) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfoZoneIds) SetZoneId(v []*string) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesInstanceTypeFamilyInfoTypesTypesInfoZoneIds {
	s.ZoneId = v
	return s
}

type ListAvailableEcsTypesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAvailableEcsTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAvailableEcsTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponse) SetHeaders(v map[string]*string) *ListAvailableEcsTypesResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableEcsTypesResponse) SetBody(v *ListAvailableEcsTypesResponseBody) *ListAvailableEcsTypesResponse {
	s.Body = v
	return s
}

type ListCloudMetricProfilingsRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCloudMetricProfilingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCloudMetricProfilingsRequest) GoString() string {
	return s.String()
}

func (s *ListCloudMetricProfilingsRequest) SetClusterId(v string) *ListCloudMetricProfilingsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCloudMetricProfilingsRequest) SetPageNumber(v int32) *ListCloudMetricProfilingsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCloudMetricProfilingsRequest) SetPageSize(v int32) *ListCloudMetricProfilingsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudMetricProfilingsRequest) SetRegionId(v string) *ListCloudMetricProfilingsRequest {
	s.RegionId = &v
	return s
}

type ListCloudMetricProfilingsResponseBody struct {
	PageNumber *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Profilings *ListCloudMetricProfilingsResponseBodyProfilings `json:"Profilings,omitempty" xml:"Profilings,omitempty" type:"Struct"`
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudMetricProfilingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCloudMetricProfilingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudMetricProfilingsResponseBody) SetPageNumber(v int32) *ListCloudMetricProfilingsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBody) SetPageSize(v int32) *ListCloudMetricProfilingsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBody) SetProfilings(v *ListCloudMetricProfilingsResponseBodyProfilings) *ListCloudMetricProfilingsResponseBody {
	s.Profilings = v
	return s
}

func (s *ListCloudMetricProfilingsResponseBody) SetRequestId(v string) *ListCloudMetricProfilingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBody) SetTotalCount(v int32) *ListCloudMetricProfilingsResponseBody {
	s.TotalCount = &v
	return s
}

type ListCloudMetricProfilingsResponseBodyProfilings struct {
	ProfilingInfo []*ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo `json:"ProfilingInfo,omitempty" xml:"ProfilingInfo,omitempty" type:"Repeated"`
}

func (s ListCloudMetricProfilingsResponseBodyProfilings) String() string {
	return tea.Prettify(s)
}

func (s ListCloudMetricProfilingsResponseBodyProfilings) GoString() string {
	return s.String()
}

func (s *ListCloudMetricProfilingsResponseBodyProfilings) SetProfilingInfo(v []*ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) *ListCloudMetricProfilingsResponseBodyProfilings {
	s.ProfilingInfo = v
	return s
}

type ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo struct {
	Duration    *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Freq        *int32  `json:"Freq,omitempty" xml:"Freq,omitempty"`
	HostName    *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Pid         *int32  `json:"Pid,omitempty" xml:"Pid,omitempty"`
	ProfilingId *string `json:"ProfilingId,omitempty" xml:"ProfilingId,omitempty"`
	TriggerTime *string `json:"TriggerTime,omitempty" xml:"TriggerTime,omitempty"`
}

func (s ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) GoString() string {
	return s.String()
}

func (s *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) SetDuration(v int32) *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo {
	s.Duration = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) SetFreq(v int32) *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo {
	s.Freq = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) SetHostName(v string) *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo {
	s.HostName = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) SetInstanceId(v string) *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo {
	s.InstanceId = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) SetPid(v int32) *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo {
	s.Pid = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) SetProfilingId(v string) *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo {
	s.ProfilingId = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo) SetTriggerTime(v string) *ListCloudMetricProfilingsResponseBodyProfilingsProfilingInfo {
	s.TriggerTime = &v
	return s
}

type ListCloudMetricProfilingsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCloudMetricProfilingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCloudMetricProfilingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCloudMetricProfilingsResponse) GoString() string {
	return s.String()
}

func (s *ListCloudMetricProfilingsResponse) SetHeaders(v map[string]*string) *ListCloudMetricProfilingsResponse {
	s.Headers = v
	return s
}

func (s *ListCloudMetricProfilingsResponse) SetBody(v *ListCloudMetricProfilingsResponseBody) *ListCloudMetricProfilingsResponse {
	s.Body = v
	return s
}

type ListClusterLogsRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClusterLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsRequest) GoString() string {
	return s.String()
}

func (s *ListClusterLogsRequest) SetClusterId(v string) *ListClusterLogsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterLogsRequest) SetPageNumber(v int32) *ListClusterLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClusterLogsRequest) SetPageSize(v int32) *ListClusterLogsRequest {
	s.PageSize = &v
	return s
}

type ListClusterLogsResponseBody struct {
	ClusterId  *string                          `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Logs       *ListClusterLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Struct"`
	PageNumber *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClusterLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterLogsResponseBody) SetClusterId(v string) *ListClusterLogsResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ListClusterLogsResponseBody) SetLogs(v *ListClusterLogsResponseBodyLogs) *ListClusterLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListClusterLogsResponseBody) SetPageNumber(v int32) *ListClusterLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClusterLogsResponseBody) SetPageSize(v int32) *ListClusterLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClusterLogsResponseBody) SetRequestId(v string) *ListClusterLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterLogsResponseBody) SetTotalCount(v int32) *ListClusterLogsResponseBody {
	s.TotalCount = &v
	return s
}

type ListClusterLogsResponseBodyLogs struct {
	LogInfo []*ListClusterLogsResponseBodyLogsLogInfo `json:"LogInfo,omitempty" xml:"LogInfo,omitempty" type:"Repeated"`
}

func (s ListClusterLogsResponseBodyLogs) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListClusterLogsResponseBodyLogs) SetLogInfo(v []*ListClusterLogsResponseBodyLogsLogInfo) *ListClusterLogsResponseBodyLogs {
	s.LogInfo = v
	return s
}

type ListClusterLogsResponseBodyLogsLogInfo struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Level      *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Operation  *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s ListClusterLogsResponseBodyLogsLogInfo) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsResponseBodyLogsLogInfo) GoString() string {
	return s.String()
}

func (s *ListClusterLogsResponseBodyLogsLogInfo) SetCreateTime(v string) *ListClusterLogsResponseBodyLogsLogInfo {
	s.CreateTime = &v
	return s
}

func (s *ListClusterLogsResponseBodyLogsLogInfo) SetLevel(v string) *ListClusterLogsResponseBodyLogsLogInfo {
	s.Level = &v
	return s
}

func (s *ListClusterLogsResponseBodyLogsLogInfo) SetMessage(v string) *ListClusterLogsResponseBodyLogsLogInfo {
	s.Message = &v
	return s
}

func (s *ListClusterLogsResponseBodyLogsLogInfo) SetOperation(v string) *ListClusterLogsResponseBodyLogsLogInfo {
	s.Operation = &v
	return s
}

type ListClusterLogsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsResponse) GoString() string {
	return s.String()
}

func (s *ListClusterLogsResponse) SetHeaders(v map[string]*string) *ListClusterLogsResponse {
	s.Headers = v
	return s
}

func (s *ListClusterLogsResponse) SetBody(v *ListClusterLogsResponseBody) *ListClusterLogsResponse {
	s.Body = v
	return s
}

type ListClustersRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) SetPageNumber(v int32) *ListClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClustersRequest) SetPageSize(v int32) *ListClustersRequest {
	s.PageSize = &v
	return s
}

type ListClustersResponseBody struct {
	Clusters   *ListClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Struct"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) SetClusters(v *ListClustersResponseBodyClusters) *ListClustersResponseBody {
	s.Clusters = v
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

func (s *ListClustersResponseBody) SetTotalCount(v int32) *ListClustersResponseBody {
	s.TotalCount = &v
	return s
}

type ListClustersResponseBodyClusters struct {
	ClusterInfoSimple []*ListClustersResponseBodyClustersClusterInfoSimple `json:"ClusterInfoSimple,omitempty" xml:"ClusterInfoSimple,omitempty" type:"Repeated"`
}

func (s ListClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClusters) SetClusterInfoSimple(v []*ListClustersResponseBodyClustersClusterInfoSimple) *ListClustersResponseBodyClusters {
	s.ClusterInfoSimple = v
	return s
}

type ListClustersResponseBodyClustersClusterInfoSimple struct {
	AccountType           *string                                                          `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	BaseOsTag             *string                                                          `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	ClientVersion         *string                                                          `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ComputeSpotPriceLimit *float32                                                         `json:"ComputeSpotPriceLimit,omitempty" xml:"ComputeSpotPriceLimit,omitempty"`
	ComputeSpotStrategy   *string                                                          `json:"ComputeSpotStrategy,omitempty" xml:"ComputeSpotStrategy,omitempty"`
	Computes              *ListClustersResponseBodyClustersClusterInfoSimpleComputes       `json:"Computes,omitempty" xml:"Computes,omitempty" type:"Struct"`
	Count                 *int32                                                           `json:"Count,omitempty" xml:"Count,omitempty"`
	CreateTime            *string                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeployMode            *string                                                          `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	Description           *string                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	EhpcVersion           *string                                                          `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	HasPlugin             *bool                                                            `json:"HasPlugin,omitempty" xml:"HasPlugin,omitempty"`
	Id                    *string                                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageId               *string                                                          `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageOwnerAlias       *string                                                          `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	InstanceChargeType    *string                                                          `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceType          *string                                                          `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IsComputeEss          *bool                                                            `json:"IsComputeEss,omitempty" xml:"IsComputeEss,omitempty"`
	Location              *string                                                          `json:"Location,omitempty" xml:"Location,omitempty"`
	LoginNodes            *string                                                          `json:"LoginNodes,omitempty" xml:"LoginNodes,omitempty"`
	Managers              *ListClustersResponseBodyClustersClusterInfoSimpleManagers       `json:"Managers,omitempty" xml:"Managers,omitempty" type:"Struct"`
	Name                  *string                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	NodePrefix            *string                                                          `json:"NodePrefix,omitempty" xml:"NodePrefix,omitempty"`
	NodeSuffix            *string                                                          `json:"NodeSuffix,omitempty" xml:"NodeSuffix,omitempty"`
	OsTag                 *string                                                          `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	RegionId              *string                                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SchedulerType         *string                                                          `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	Status                *string                                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalResources        *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty" type:"Struct"`
	UsedResources         *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources  `json:"UsedResources,omitempty" xml:"UsedResources,omitempty" type:"Struct"`
	VSwitchId             *string                                                          `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                 *string                                                          `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                *string                                                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoSimple) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoSimple) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetAccountType(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.AccountType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetBaseOsTag(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.BaseOsTag = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetClientVersion(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.ClientVersion = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetComputeSpotPriceLimit(v float32) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.ComputeSpotPriceLimit = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetComputeSpotStrategy(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.ComputeSpotStrategy = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetComputes(v *ListClustersResponseBodyClustersClusterInfoSimpleComputes) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Computes = v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Count = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetCreateTime(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.CreateTime = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetDeployMode(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.DeployMode = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetDescription(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Description = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetEhpcVersion(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.EhpcVersion = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetHasPlugin(v bool) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.HasPlugin = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetId(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Id = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetImageId(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.ImageId = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetImageOwnerAlias(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetInstanceChargeType(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.InstanceChargeType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetInstanceType(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.InstanceType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetIsComputeEss(v bool) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.IsComputeEss = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetLocation(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Location = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetLoginNodes(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.LoginNodes = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetManagers(v *ListClustersResponseBodyClustersClusterInfoSimpleManagers) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Managers = v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetName(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetNodePrefix(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.NodePrefix = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetNodeSuffix(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.NodeSuffix = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetOsTag(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.OsTag = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetRegionId(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.RegionId = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetSchedulerType(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.SchedulerType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetStatus(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Status = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetTotalResources(v *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.TotalResources = v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetUsedResources(v *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.UsedResources = v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetVSwitchId(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.VSwitchId = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetVpcId(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.VpcId = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetZoneId(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.ZoneId = &v
	return s
}

type ListClustersResponseBodyClustersClusterInfoSimpleComputes struct {
	ExceptionCount *int32 `json:"ExceptionCount,omitempty" xml:"ExceptionCount,omitempty"`
	NormalCount    *int32 `json:"NormalCount,omitempty" xml:"NormalCount,omitempty"`
	OperatingCount *int32 `json:"OperatingCount,omitempty" xml:"OperatingCount,omitempty"`
	StoppedCount   *int32 `json:"StoppedCount,omitempty" xml:"StoppedCount,omitempty"`
	Total          *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleComputes) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleComputes) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleComputes) SetExceptionCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleComputes {
	s.ExceptionCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleComputes) SetNormalCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleComputes {
	s.NormalCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleComputes) SetOperatingCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleComputes {
	s.OperatingCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleComputes) SetStoppedCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleComputes {
	s.StoppedCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleComputes) SetTotal(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleComputes {
	s.Total = &v
	return s
}

type ListClustersResponseBodyClustersClusterInfoSimpleManagers struct {
	ExceptionCount *int32 `json:"ExceptionCount,omitempty" xml:"ExceptionCount,omitempty"`
	NormalCount    *int32 `json:"NormalCount,omitempty" xml:"NormalCount,omitempty"`
	OperatingCount *int32 `json:"OperatingCount,omitempty" xml:"OperatingCount,omitempty"`
	StoppedCount   *int32 `json:"StoppedCount,omitempty" xml:"StoppedCount,omitempty"`
	Total          *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleManagers) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleManagers) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleManagers) SetExceptionCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleManagers {
	s.ExceptionCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleManagers) SetNormalCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleManagers {
	s.NormalCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleManagers) SetOperatingCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleManagers {
	s.OperatingCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleManagers) SetStoppedCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleManagers {
	s.StoppedCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleManagers) SetTotal(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleManagers {
	s.Total = &v
	return s
}

type ListClustersResponseBodyClustersClusterInfoSimpleTotalResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) SetCpu(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources {
	s.Cpu = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) SetGpu(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources {
	s.Gpu = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) SetMemory(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources {
	s.Memory = &v
	return s
}

type ListClustersResponseBodyClustersClusterInfoSimpleUsedResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) SetCpu(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources {
	s.Cpu = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) SetGpu(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources {
	s.Gpu = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) SetMemory(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources {
	s.Memory = &v
	return s
}

type ListClustersResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponse) GoString() string {
	return s.String()
}

func (s *ListClustersResponse) SetHeaders(v map[string]*string) *ListClustersResponse {
	s.Headers = v
	return s
}

func (s *ListClustersResponse) SetBody(v *ListClustersResponseBody) *ListClustersResponse {
	s.Body = v
	return s
}

type ListClustersMetaRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClustersMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersMetaRequest) GoString() string {
	return s.String()
}

func (s *ListClustersMetaRequest) SetPageNumber(v int32) *ListClustersMetaRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClustersMetaRequest) SetPageSize(v int32) *ListClustersMetaRequest {
	s.PageSize = &v
	return s
}

type ListClustersMetaResponseBody struct {
	Clusters   *ListClustersMetaResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Struct"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClustersMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersMetaResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersMetaResponseBody) SetClusters(v *ListClustersMetaResponseBodyClusters) *ListClustersMetaResponseBody {
	s.Clusters = v
	return s
}

func (s *ListClustersMetaResponseBody) SetPageNumber(v int32) *ListClustersMetaResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClustersMetaResponseBody) SetPageSize(v int32) *ListClustersMetaResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClustersMetaResponseBody) SetRequestId(v string) *ListClustersMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersMetaResponseBody) SetTotalCount(v int32) *ListClustersMetaResponseBody {
	s.TotalCount = &v
	return s
}

type ListClustersMetaResponseBodyClusters struct {
	ClusterInfoSimple []*ListClustersMetaResponseBodyClustersClusterInfoSimple `json:"ClusterInfoSimple,omitempty" xml:"ClusterInfoSimple,omitempty" type:"Repeated"`
}

func (s ListClustersMetaResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListClustersMetaResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersMetaResponseBodyClusters) SetClusterInfoSimple(v []*ListClustersMetaResponseBodyClustersClusterInfoSimple) *ListClustersMetaResponseBodyClusters {
	s.ClusterInfoSimple = v
	return s
}

type ListClustersMetaResponseBodyClustersClusterInfoSimple struct {
	AccountType   *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	DeployMode    *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HasPlugin     *bool   `json:"HasPlugin,omitempty" xml:"HasPlugin,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IsComputeEss  *bool   `json:"IsComputeEss,omitempty" xml:"IsComputeEss,omitempty"`
	Location      *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OsTag         *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListClustersMetaResponseBodyClustersClusterInfoSimple) String() string {
	return tea.Prettify(s)
}

func (s ListClustersMetaResponseBodyClustersClusterInfoSimple) GoString() string {
	return s.String()
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetAccountType(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.AccountType = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetClientVersion(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.ClientVersion = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetDeployMode(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.DeployMode = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetDescription(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.Description = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetHasPlugin(v bool) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.HasPlugin = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetId(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.Id = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetIsComputeEss(v bool) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.IsComputeEss = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetLocation(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.Location = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetName(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.Name = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetOsTag(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.OsTag = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetSchedulerType(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.SchedulerType = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetStatus(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.Status = &v
	return s
}

func (s *ListClustersMetaResponseBodyClustersClusterInfoSimple) SetVpcId(v string) *ListClustersMetaResponseBodyClustersClusterInfoSimple {
	s.VpcId = &v
	return s
}

type ListClustersMetaResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClustersMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClustersMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersMetaResponse) GoString() string {
	return s.String()
}

func (s *ListClustersMetaResponse) SetHeaders(v map[string]*string) *ListClustersMetaResponse {
	s.Headers = v
	return s
}

func (s *ListClustersMetaResponse) SetBody(v *ListClustersMetaResponseBody) *ListClustersMetaResponse {
	s.Body = v
	return s
}

type ListCommandsRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CommandId  *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCommandsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCommandsRequest) GoString() string {
	return s.String()
}

func (s *ListCommandsRequest) SetClusterId(v string) *ListCommandsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCommandsRequest) SetCommandId(v string) *ListCommandsRequest {
	s.CommandId = &v
	return s
}

func (s *ListCommandsRequest) SetPageNumber(v int32) *ListCommandsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCommandsRequest) SetPageSize(v int32) *ListCommandsRequest {
	s.PageSize = &v
	return s
}

type ListCommandsResponseBody struct {
	Commands   *ListCommandsResponseBodyCommands `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Struct"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCommandsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCommandsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommandsResponseBody) SetCommands(v *ListCommandsResponseBodyCommands) *ListCommandsResponseBody {
	s.Commands = v
	return s
}

func (s *ListCommandsResponseBody) SetPageNumber(v int32) *ListCommandsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCommandsResponseBody) SetPageSize(v int32) *ListCommandsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCommandsResponseBody) SetRequestId(v string) *ListCommandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommandsResponseBody) SetTotalCount(v int32) *ListCommandsResponseBody {
	s.TotalCount = &v
	return s
}

type ListCommandsResponseBodyCommands struct {
	Command []*ListCommandsResponseBodyCommandsCommand `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s ListCommandsResponseBodyCommands) String() string {
	return tea.Prettify(s)
}

func (s ListCommandsResponseBodyCommands) GoString() string {
	return s.String()
}

func (s *ListCommandsResponseBodyCommands) SetCommand(v []*ListCommandsResponseBodyCommandsCommand) *ListCommandsResponseBodyCommands {
	s.Command = v
	return s
}

type ListCommandsResponseBodyCommandsCommand struct {
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	CommandId      *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	Timeout        *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	WorkingDir     *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s ListCommandsResponseBodyCommandsCommand) String() string {
	return tea.Prettify(s)
}

func (s ListCommandsResponseBodyCommandsCommand) GoString() string {
	return s.String()
}

func (s *ListCommandsResponseBodyCommandsCommand) SetCommandContent(v string) *ListCommandsResponseBodyCommandsCommand {
	s.CommandContent = &v
	return s
}

func (s *ListCommandsResponseBodyCommandsCommand) SetCommandId(v string) *ListCommandsResponseBodyCommandsCommand {
	s.CommandId = &v
	return s
}

func (s *ListCommandsResponseBodyCommandsCommand) SetTimeout(v string) *ListCommandsResponseBodyCommandsCommand {
	s.Timeout = &v
	return s
}

func (s *ListCommandsResponseBodyCommandsCommand) SetWorkingDir(v string) *ListCommandsResponseBodyCommandsCommand {
	s.WorkingDir = &v
	return s
}

type ListCommandsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCommandsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCommandsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCommandsResponse) GoString() string {
	return s.String()
}

func (s *ListCommandsResponse) SetHeaders(v map[string]*string) *ListCommandsResponse {
	s.Headers = v
	return s
}

func (s *ListCommandsResponse) SetBody(v *ListCommandsResponseBody) *ListCommandsResponse {
	s.Body = v
	return s
}

type ListContainerAppsRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListContainerAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListContainerAppsRequest) GoString() string {
	return s.String()
}

func (s *ListContainerAppsRequest) SetPageNumber(v int32) *ListContainerAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListContainerAppsRequest) SetPageSize(v int32) *ListContainerAppsRequest {
	s.PageSize = &v
	return s
}

type ListContainerAppsResponseBody struct {
	ContainerApps *ListContainerAppsResponseBodyContainerApps `json:"ContainerApps,omitempty" xml:"ContainerApps,omitempty" type:"Struct"`
	PageNumber    *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListContainerAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListContainerAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContainerAppsResponseBody) SetContainerApps(v *ListContainerAppsResponseBodyContainerApps) *ListContainerAppsResponseBody {
	s.ContainerApps = v
	return s
}

func (s *ListContainerAppsResponseBody) SetPageNumber(v int32) *ListContainerAppsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListContainerAppsResponseBody) SetPageSize(v int32) *ListContainerAppsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListContainerAppsResponseBody) SetRequestId(v string) *ListContainerAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContainerAppsResponseBody) SetTotalCount(v int32) *ListContainerAppsResponseBody {
	s.TotalCount = &v
	return s
}

type ListContainerAppsResponseBodyContainerApps struct {
	ContainerApps []*ListContainerAppsResponseBodyContainerAppsContainerApps `json:"ContainerApps,omitempty" xml:"ContainerApps,omitempty" type:"Repeated"`
}

func (s ListContainerAppsResponseBodyContainerApps) String() string {
	return tea.Prettify(s)
}

func (s ListContainerAppsResponseBodyContainerApps) GoString() string {
	return s.String()
}

func (s *ListContainerAppsResponseBodyContainerApps) SetContainerApps(v []*ListContainerAppsResponseBodyContainerAppsContainerApps) *ListContainerAppsResponseBodyContainerApps {
	s.ContainerApps = v
	return s
}

type ListContainerAppsResponseBodyContainerAppsContainerApps struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageTag    *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Repository  *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListContainerAppsResponseBodyContainerAppsContainerApps) String() string {
	return tea.Prettify(s)
}

func (s ListContainerAppsResponseBodyContainerAppsContainerApps) GoString() string {
	return s.String()
}

func (s *ListContainerAppsResponseBodyContainerAppsContainerApps) SetCreateTime(v string) *ListContainerAppsResponseBodyContainerAppsContainerApps {
	s.CreateTime = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerAppsContainerApps) SetDescription(v string) *ListContainerAppsResponseBodyContainerAppsContainerApps {
	s.Description = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerAppsContainerApps) SetId(v string) *ListContainerAppsResponseBodyContainerAppsContainerApps {
	s.Id = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerAppsContainerApps) SetImageTag(v string) *ListContainerAppsResponseBodyContainerAppsContainerApps {
	s.ImageTag = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerAppsContainerApps) SetName(v string) *ListContainerAppsResponseBodyContainerAppsContainerApps {
	s.Name = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerAppsContainerApps) SetRepository(v string) *ListContainerAppsResponseBodyContainerAppsContainerApps {
	s.Repository = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerAppsContainerApps) SetType(v string) *ListContainerAppsResponseBodyContainerAppsContainerApps {
	s.Type = &v
	return s
}

type ListContainerAppsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListContainerAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListContainerAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListContainerAppsResponse) GoString() string {
	return s.String()
}

func (s *ListContainerAppsResponse) SetHeaders(v map[string]*string) *ListContainerAppsResponse {
	s.Headers = v
	return s
}

func (s *ListContainerAppsResponse) SetBody(v *ListContainerAppsResponseBody) *ListContainerAppsResponse {
	s.Body = v
	return s
}

type ListContainerImagesRequest struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListContainerImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesRequest) GoString() string {
	return s.String()
}

func (s *ListContainerImagesRequest) SetClusterId(v string) *ListContainerImagesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListContainerImagesRequest) SetContainerType(v string) *ListContainerImagesRequest {
	s.ContainerType = &v
	return s
}

func (s *ListContainerImagesRequest) SetPageNumber(v int32) *ListContainerImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListContainerImagesRequest) SetPageSize(v int32) *ListContainerImagesRequest {
	s.PageSize = &v
	return s
}

type ListContainerImagesResponseBody struct {
	DBInfo     *string                                `json:"DBInfo,omitempty" xml:"DBInfo,omitempty"`
	Images     *ListContainerImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListContainerImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListContainerImagesResponseBody) SetDBInfo(v string) *ListContainerImagesResponseBody {
	s.DBInfo = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetImages(v *ListContainerImagesResponseBodyImages) *ListContainerImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListContainerImagesResponseBody) SetPageNumber(v int32) *ListContainerImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetPageSize(v int32) *ListContainerImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetRequestId(v string) *ListContainerImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetTotalCount(v int32) *ListContainerImagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListContainerImagesResponseBodyImages struct {
	Images []*ListContainerImagesResponseBodyImagesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
}

func (s ListContainerImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListContainerImagesResponseBodyImages) SetImages(v []*ListContainerImagesResponseBodyImagesImages) *ListContainerImagesResponseBodyImages {
	s.Images = v
	return s
}

type ListContainerImagesResponseBodyImagesImages struct {
	ImageId        *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Repository     *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	System         *string `json:"System,omitempty" xml:"System,omitempty"`
	Tag            *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateDateTime *string `json:"UpdateDateTime,omitempty" xml:"UpdateDateTime,omitempty"`
}

func (s ListContainerImagesResponseBodyImagesImages) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesResponseBodyImagesImages) GoString() string {
	return s.String()
}

func (s *ListContainerImagesResponseBodyImagesImages) SetImageId(v string) *ListContainerImagesResponseBodyImagesImages {
	s.ImageId = &v
	return s
}

func (s *ListContainerImagesResponseBodyImagesImages) SetRepository(v string) *ListContainerImagesResponseBodyImagesImages {
	s.Repository = &v
	return s
}

func (s *ListContainerImagesResponseBodyImagesImages) SetStatus(v string) *ListContainerImagesResponseBodyImagesImages {
	s.Status = &v
	return s
}

func (s *ListContainerImagesResponseBodyImagesImages) SetSystem(v string) *ListContainerImagesResponseBodyImagesImages {
	s.System = &v
	return s
}

func (s *ListContainerImagesResponseBodyImagesImages) SetTag(v string) *ListContainerImagesResponseBodyImagesImages {
	s.Tag = &v
	return s
}

func (s *ListContainerImagesResponseBodyImagesImages) SetType(v string) *ListContainerImagesResponseBodyImagesImages {
	s.Type = &v
	return s
}

func (s *ListContainerImagesResponseBodyImagesImages) SetUpdateDateTime(v string) *ListContainerImagesResponseBodyImagesImages {
	s.UpdateDateTime = &v
	return s
}

type ListContainerImagesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListContainerImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListContainerImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesResponse) GoString() string {
	return s.String()
}

func (s *ListContainerImagesResponse) SetHeaders(v map[string]*string) *ListContainerImagesResponse {
	s.Headers = v
	return s
}

func (s *ListContainerImagesResponse) SetBody(v *ListContainerImagesResponseBody) *ListContainerImagesResponse {
	s.Body = v
	return s
}

type ListCpfsFileSystemsRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCpfsFileSystemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsRequest) SetFileSystemId(v string) *ListCpfsFileSystemsRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListCpfsFileSystemsRequest) SetPageNumber(v int32) *ListCpfsFileSystemsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCpfsFileSystemsRequest) SetPageSize(v int32) *ListCpfsFileSystemsRequest {
	s.PageSize = &v
	return s
}

type ListCpfsFileSystemsResponseBody struct {
	FileSystemList *ListCpfsFileSystemsResponseBodyFileSystemList `json:"FileSystemList,omitempty" xml:"FileSystemList,omitempty" type:"Struct"`
	PageNumber     *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCpfsFileSystemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsResponseBody) SetFileSystemList(v *ListCpfsFileSystemsResponseBodyFileSystemList) *ListCpfsFileSystemsResponseBody {
	s.FileSystemList = v
	return s
}

func (s *ListCpfsFileSystemsResponseBody) SetPageNumber(v int32) *ListCpfsFileSystemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBody) SetPageSize(v int32) *ListCpfsFileSystemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBody) SetRequestId(v string) *ListCpfsFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBody) SetTotalCount(v int32) *ListCpfsFileSystemsResponseBody {
	s.TotalCount = &v
	return s
}

type ListCpfsFileSystemsResponseBodyFileSystemList struct {
	FileSystems []*ListCpfsFileSystemsResponseBodyFileSystemListFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Repeated"`
}

func (s ListCpfsFileSystemsResponseBodyFileSystemList) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsResponseBodyFileSystemList) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemList) SetFileSystems(v []*ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) *ListCpfsFileSystemsResponseBodyFileSystemList {
	s.FileSystems = v
	return s
}

type ListCpfsFileSystemsResponseBodyFileSystemListFileSystems struct {
	Capacity        *string                                                                  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	CreateTime      *string                                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Destription     *string                                                                  `json:"Destription,omitempty" xml:"Destription,omitempty"`
	FileSystemId    *string                                                                  `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountTargetList *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetList `json:"MountTargetList,omitempty" xml:"MountTargetList,omitempty" type:"Struct"`
	ProtocolType    *string                                                                  `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RegionId        *string                                                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId          *string                                                                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) SetCapacity(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems {
	s.Capacity = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) SetCreateTime(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems {
	s.CreateTime = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) SetDestription(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems {
	s.Destription = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) SetFileSystemId(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems {
	s.FileSystemId = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) SetMountTargetList(v *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetList) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems {
	s.MountTargetList = v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) SetProtocolType(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems {
	s.ProtocolType = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) SetRegionId(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems {
	s.RegionId = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems) SetZoneId(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystems {
	s.ZoneId = &v
	return s
}

type ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetList struct {
	MountTargets []*ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets `json:"MountTargets,omitempty" xml:"MountTargets,omitempty" type:"Repeated"`
}

func (s ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetList) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetList) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetList) SetMountTargets(v []*ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetList {
	s.MountTargets = v
	return s
}

type ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets struct {
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	NetworkType       *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId             *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswId             *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
}

func (s ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetMountTargetDomain(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.MountTargetDomain = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetNetworkType(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.NetworkType = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetStatus(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.Status = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetVpcId(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.VpcId = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetVswId(v string) *ListCpfsFileSystemsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.VswId = &v
	return s
}

type ListCpfsFileSystemsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCpfsFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCpfsFileSystemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsResponse) SetHeaders(v map[string]*string) *ListCpfsFileSystemsResponse {
	s.Headers = v
	return s
}

func (s *ListCpfsFileSystemsResponse) SetBody(v *ListCpfsFileSystemsResponseBody) *ListCpfsFileSystemsResponse {
	s.Body = v
	return s
}

type ListCurrentClientVersionResponseBody struct {
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCurrentClientVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCurrentClientVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ListCurrentClientVersionResponseBody) SetClientVersion(v string) *ListCurrentClientVersionResponseBody {
	s.ClientVersion = &v
	return s
}

func (s *ListCurrentClientVersionResponseBody) SetRequestId(v string) *ListCurrentClientVersionResponseBody {
	s.RequestId = &v
	return s
}

type ListCurrentClientVersionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCurrentClientVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCurrentClientVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCurrentClientVersionResponse) GoString() string {
	return s.String()
}

func (s *ListCurrentClientVersionResponse) SetHeaders(v map[string]*string) *ListCurrentClientVersionResponse {
	s.Headers = v
	return s
}

func (s *ListCurrentClientVersionResponse) SetBody(v *ListCurrentClientVersionResponseBody) *ListCurrentClientVersionResponse {
	s.Body = v
	return s
}

type ListCustomImagesRequest struct {
	BaseOsTag       *string `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ListCustomImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesRequest) GoString() string {
	return s.String()
}

func (s *ListCustomImagesRequest) SetBaseOsTag(v string) *ListCustomImagesRequest {
	s.BaseOsTag = &v
	return s
}

func (s *ListCustomImagesRequest) SetClusterId(v string) *ListCustomImagesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCustomImagesRequest) SetImageOwnerAlias(v string) *ListCustomImagesRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListCustomImagesRequest) SetInstanceType(v string) *ListCustomImagesRequest {
	s.InstanceType = &v
	return s
}

type ListCustomImagesResponseBody struct {
	Images    *ListCustomImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCustomImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBody) SetImages(v *ListCustomImagesResponseBodyImages) *ListCustomImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListCustomImagesResponseBody) SetRequestId(v string) *ListCustomImagesResponseBody {
	s.RequestId = &v
	return s
}

type ListCustomImagesResponseBodyImages struct {
	ImageInfo []*ListCustomImagesResponseBodyImagesImageInfo `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" type:"Repeated"`
}

func (s ListCustomImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyImages) SetImageInfo(v []*ListCustomImagesResponseBodyImagesImageInfo) *ListCustomImagesResponseBodyImages {
	s.ImageInfo = v
	return s
}

type ListCustomImagesResponseBodyImagesImageInfo struct {
	BaseOsTag         *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty" type:"Struct"`
	Description       *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageId           *string                                               `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName         *string                                               `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageOwnerAlias   *string                                               `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	OsTag             *ListCustomImagesResponseBodyImagesImageInfoOsTag     `json:"OsTag,omitempty" xml:"OsTag,omitempty" type:"Struct"`
	PostInstallScript *string                                               `json:"PostInstallScript,omitempty" xml:"PostInstallScript,omitempty"`
	PricingCycle      *string                                               `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	ProductCode       *string                                               `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Size              *int32                                                `json:"Size,omitempty" xml:"Size,omitempty"`
	SkuCode           *string                                               `json:"SkuCode,omitempty" xml:"SkuCode,omitempty"`
	Status            *string                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	Uid               *string                                               `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListCustomImagesResponseBodyImagesImageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBodyImagesImageInfo) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetBaseOsTag(v *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) *ListCustomImagesResponseBodyImagesImageInfo {
	s.BaseOsTag = v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetDescription(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.Description = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetImageId(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.ImageId = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetImageName(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.ImageName = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetImageOwnerAlias(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetOsTag(v *ListCustomImagesResponseBodyImagesImageInfoOsTag) *ListCustomImagesResponseBodyImagesImageInfo {
	s.OsTag = v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetPostInstallScript(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.PostInstallScript = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetPricingCycle(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.PricingCycle = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetProductCode(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.ProductCode = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetSize(v int32) *ListCustomImagesResponseBodyImagesImageInfo {
	s.Size = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetSkuCode(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.SkuCode = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetStatus(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.Status = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetUid(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.Uid = &v
	return s
}

type ListCustomImagesResponseBodyImagesImageInfoBaseOsTag struct {
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	OsTag        *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) SetArchitecture(v string) *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag {
	s.Architecture = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) SetOsTag(v string) *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag {
	s.OsTag = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) SetPlatform(v string) *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag {
	s.Platform = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) SetVersion(v string) *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag {
	s.Version = &v
	return s
}

type ListCustomImagesResponseBodyImagesImageInfoOsTag struct {
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	BaseOsTag    *string `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	OsTag        *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListCustomImagesResponseBodyImagesImageInfoOsTag) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBodyImagesImageInfoOsTag) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyImagesImageInfoOsTag) SetArchitecture(v string) *ListCustomImagesResponseBodyImagesImageInfoOsTag {
	s.Architecture = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfoOsTag) SetBaseOsTag(v string) *ListCustomImagesResponseBodyImagesImageInfoOsTag {
	s.BaseOsTag = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfoOsTag) SetOsTag(v string) *ListCustomImagesResponseBodyImagesImageInfoOsTag {
	s.OsTag = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfoOsTag) SetPlatform(v string) *ListCustomImagesResponseBodyImagesImageInfoOsTag {
	s.Platform = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfoOsTag) SetVersion(v string) *ListCustomImagesResponseBodyImagesImageInfoOsTag {
	s.Version = &v
	return s
}

type ListCustomImagesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCustomImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCustomImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponse) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponse) SetHeaders(v map[string]*string) *ListCustomImagesResponse {
	s.Headers = v
	return s
}

func (s *ListCustomImagesResponse) SetBody(v *ListCustomImagesResponseBody) *ListCustomImagesResponse {
	s.Body = v
	return s
}

type ListFileSystemWithMountTargetsRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFileSystemWithMountTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsRequest) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsRequest) SetPageNumber(v int32) *ListFileSystemWithMountTargetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFileSystemWithMountTargetsRequest) SetPageSize(v int32) *ListFileSystemWithMountTargetsRequest {
	s.PageSize = &v
	return s
}

type ListFileSystemWithMountTargetsResponseBody struct {
	FileSystemList *ListFileSystemWithMountTargetsResponseBodyFileSystemList `json:"FileSystemList,omitempty" xml:"FileSystemList,omitempty" type:"Struct"`
	PageNumber     *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFileSystemWithMountTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBody) SetFileSystemList(v *ListFileSystemWithMountTargetsResponseBodyFileSystemList) *ListFileSystemWithMountTargetsResponseBody {
	s.FileSystemList = v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBody) SetPageNumber(v int32) *ListFileSystemWithMountTargetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBody) SetPageSize(v int32) *ListFileSystemWithMountTargetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBody) SetRequestId(v string) *ListFileSystemWithMountTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBody) SetTotalCount(v int32) *ListFileSystemWithMountTargetsResponseBody {
	s.TotalCount = &v
	return s
}

type ListFileSystemWithMountTargetsResponseBodyFileSystemList struct {
	FileSystems []*ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Repeated"`
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemList) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemList) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemList) SetFileSystems(v []*ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) *ListFileSystemWithMountTargetsResponseBodyFileSystemList {
	s.FileSystems = v
	return s
}

type ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems struct {
	BandWidth       *int32                                                                              `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	Capacity        *int32                                                                              `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	CreateTime      *string                                                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Destription     *string                                                                             `json:"Destription,omitempty" xml:"Destription,omitempty"`
	EncryptType     *int32                                                                              `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	FileSystemId    *string                                                                             `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemType  *string                                                                             `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	MeteredSize     *int32                                                                              `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	MountTargetList *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetList `json:"MountTargetList,omitempty" xml:"MountTargetList,omitempty" type:"Struct"`
	PackageList     *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageList     `json:"PackageList,omitempty" xml:"PackageList,omitempty" type:"Struct"`
	ProtocolType    *string                                                                             `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RegionId        *string                                                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status          *string                                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageType     *string                                                                             `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetBandWidth(v int32) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.BandWidth = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetCapacity(v int32) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.Capacity = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetCreateTime(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.CreateTime = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetDestription(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.Destription = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetEncryptType(v int32) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.EncryptType = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetFileSystemId(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.FileSystemId = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetFileSystemType(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.FileSystemType = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetMeteredSize(v int32) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.MeteredSize = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetMountTargetList(v *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetList) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.MountTargetList = v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetPackageList(v *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageList) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.PackageList = v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetProtocolType(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.ProtocolType = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetRegionId(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.RegionId = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetStatus(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.Status = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems) SetStorageType(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystems {
	s.StorageType = &v
	return s
}

type ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetList struct {
	MountTargets []*ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets `json:"MountTargets,omitempty" xml:"MountTargets,omitempty" type:"Repeated"`
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetList) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetList) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetList) SetMountTargets(v []*ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetList {
	s.MountTargets = v
	return s
}

type ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets struct {
	AccessGroup       *string `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	NetworkType       *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId             *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswId             *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetAccessGroup(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.AccessGroup = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetMountTargetDomain(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.MountTargetDomain = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetNetworkType(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.NetworkType = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetStatus(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.Status = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetVpcId(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.VpcId = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets) SetVswId(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsMountTargetListMountTargets {
	s.VswId = &v
	return s
}

type ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageList struct {
	Packages []*ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageListPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Repeated"`
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageList) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageList) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageList) SetPackages(v []*ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageListPackages) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageList {
	s.Packages = v
	return s
}

type ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageListPackages struct {
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageListPackages) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageListPackages) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageListPackages) SetPackageId(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListFileSystemsPackageListPackages {
	s.PackageId = &v
	return s
}

type ListFileSystemWithMountTargetsResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFileSystemWithMountTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFileSystemWithMountTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponse) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponse) SetHeaders(v map[string]*string) *ListFileSystemWithMountTargetsResponse {
	s.Headers = v
	return s
}

func (s *ListFileSystemWithMountTargetsResponse) SetBody(v *ListFileSystemWithMountTargetsResponseBody) *ListFileSystemWithMountTargetsResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	BaseOsTag    *string `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetBaseOsTag(v string) *ListImagesRequest {
	s.BaseOsTag = &v
	return s
}

func (s *ListImagesRequest) SetInstanceType(v string) *ListImagesRequest {
	s.InstanceType = &v
	return s
}

type ListImagesResponseBody struct {
	OsTags    *ListImagesResponseBodyOsTags `json:"OsTags,omitempty" xml:"OsTags,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetOsTags(v *ListImagesResponseBodyOsTags) *ListImagesResponseBody {
	s.OsTags = v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

type ListImagesResponseBodyOsTags struct {
	OsInfo []*ListImagesResponseBodyOsTagsOsInfo `json:"OsInfo,omitempty" xml:"OsInfo,omitempty" type:"Repeated"`
}

func (s ListImagesResponseBodyOsTags) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyOsTags) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyOsTags) SetOsInfo(v []*ListImagesResponseBodyOsTagsOsInfo) *ListImagesResponseBodyOsTags {
	s.OsInfo = v
	return s
}

type ListImagesResponseBodyOsTagsOsInfo struct {
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	BaseOsTag    *string `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	OsTag        *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListImagesResponseBodyOsTagsOsInfo) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyOsTagsOsInfo) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetArchitecture(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.Architecture = &v
	return s
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetBaseOsTag(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.BaseOsTag = &v
	return s
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetImageId(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetOsTag(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.OsTag = &v
	return s
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetPlatform(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.Platform = &v
	return s
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetVersion(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.Version = &v
	return s
}

type ListImagesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type ListInstalledSoftwareRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListInstalledSoftwareRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwareRequest) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwareRequest) SetClusterId(v string) *ListInstalledSoftwareRequest {
	s.ClusterId = &v
	return s
}

type ListInstalledSoftwareResponseBody struct {
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SoftwareList *ListInstalledSoftwareResponseBodySoftwareList `json:"SoftwareList,omitempty" xml:"SoftwareList,omitempty" type:"Struct"`
}

func (s ListInstalledSoftwareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwareResponseBody) SetRequestId(v string) *ListInstalledSoftwareResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstalledSoftwareResponseBody) SetSoftwareList(v *ListInstalledSoftwareResponseBodySoftwareList) *ListInstalledSoftwareResponseBody {
	s.SoftwareList = v
	return s
}

type ListInstalledSoftwareResponseBodySoftwareList struct {
	SoftwareList []*ListInstalledSoftwareResponseBodySoftwareListSoftwareList `json:"SoftwareList,omitempty" xml:"SoftwareList,omitempty" type:"Repeated"`
}

func (s ListInstalledSoftwareResponseBodySoftwareList) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwareResponseBodySoftwareList) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwareResponseBodySoftwareList) SetSoftwareList(v []*ListInstalledSoftwareResponseBodySoftwareListSoftwareList) *ListInstalledSoftwareResponseBodySoftwareList {
	s.SoftwareList = v
	return s
}

type ListInstalledSoftwareResponseBodySoftwareListSoftwareList struct {
	SoftwareId      *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	SoftwareName    *string `json:"SoftwareName,omitempty" xml:"SoftwareName,omitempty"`
	SoftwareStatus  *string `json:"SoftwareStatus,omitempty" xml:"SoftwareStatus,omitempty"`
	SoftwareVersion *string `json:"SoftwareVersion,omitempty" xml:"SoftwareVersion,omitempty"`
}

func (s ListInstalledSoftwareResponseBodySoftwareListSoftwareList) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwareResponseBodySoftwareListSoftwareList) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwareResponseBodySoftwareListSoftwareList) SetSoftwareId(v string) *ListInstalledSoftwareResponseBodySoftwareListSoftwareList {
	s.SoftwareId = &v
	return s
}

func (s *ListInstalledSoftwareResponseBodySoftwareListSoftwareList) SetSoftwareName(v string) *ListInstalledSoftwareResponseBodySoftwareListSoftwareList {
	s.SoftwareName = &v
	return s
}

func (s *ListInstalledSoftwareResponseBodySoftwareListSoftwareList) SetSoftwareStatus(v string) *ListInstalledSoftwareResponseBodySoftwareListSoftwareList {
	s.SoftwareStatus = &v
	return s
}

func (s *ListInstalledSoftwareResponseBodySoftwareListSoftwareList) SetSoftwareVersion(v string) *ListInstalledSoftwareResponseBodySoftwareListSoftwareList {
	s.SoftwareVersion = &v
	return s
}

type ListInstalledSoftwareResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstalledSoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstalledSoftwareResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwareResponse) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwareResponse) SetHeaders(v map[string]*string) *ListInstalledSoftwareResponse {
	s.Headers = v
	return s
}

func (s *ListInstalledSoftwareResponse) SetBody(v *ListInstalledSoftwareResponseBody) *ListInstalledSoftwareResponse {
	s.Body = v
	return s
}

type ListInvocationResultsRequest struct {
	ClusterId          *string                                 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CommandId          *string                                 `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	Instance           []*ListInvocationResultsRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
	InvokeRecordStatus *string                                 `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	PageNumber         *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInvocationResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationResultsRequest) GoString() string {
	return s.String()
}

func (s *ListInvocationResultsRequest) SetClusterId(v string) *ListInvocationResultsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInvocationResultsRequest) SetCommandId(v string) *ListInvocationResultsRequest {
	s.CommandId = &v
	return s
}

func (s *ListInvocationResultsRequest) SetInstance(v []*ListInvocationResultsRequestInstance) *ListInvocationResultsRequest {
	s.Instance = v
	return s
}

func (s *ListInvocationResultsRequest) SetInvokeRecordStatus(v string) *ListInvocationResultsRequest {
	s.InvokeRecordStatus = &v
	return s
}

func (s *ListInvocationResultsRequest) SetPageNumber(v int32) *ListInvocationResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInvocationResultsRequest) SetPageSize(v int32) *ListInvocationResultsRequest {
	s.PageSize = &v
	return s
}

type ListInvocationResultsRequestInstance struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListInvocationResultsRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationResultsRequestInstance) GoString() string {
	return s.String()
}

func (s *ListInvocationResultsRequestInstance) SetId(v string) *ListInvocationResultsRequestInstance {
	s.Id = &v
	return s
}

type ListInvocationResultsResponseBody struct {
	InvocationResults *ListInvocationResultsResponseBodyInvocationResults `json:"InvocationResults,omitempty" xml:"InvocationResults,omitempty" type:"Struct"`
	PageNumber        *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount        *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInvocationResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInvocationResultsResponseBody) SetInvocationResults(v *ListInvocationResultsResponseBodyInvocationResults) *ListInvocationResultsResponseBody {
	s.InvocationResults = v
	return s
}

func (s *ListInvocationResultsResponseBody) SetPageNumber(v int32) *ListInvocationResultsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInvocationResultsResponseBody) SetPageSize(v int32) *ListInvocationResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInvocationResultsResponseBody) SetRequestId(v string) *ListInvocationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInvocationResultsResponseBody) SetTotalCount(v int32) *ListInvocationResultsResponseBody {
	s.TotalCount = &v
	return s
}

type ListInvocationResultsResponseBodyInvocationResults struct {
	InvocationResult []*ListInvocationResultsResponseBodyInvocationResultsInvocationResult `json:"InvocationResult,omitempty" xml:"InvocationResult,omitempty" type:"Repeated"`
}

func (s ListInvocationResultsResponseBodyInvocationResults) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationResultsResponseBodyInvocationResults) GoString() string {
	return s.String()
}

func (s *ListInvocationResultsResponseBodyInvocationResults) SetInvocationResult(v []*ListInvocationResultsResponseBodyInvocationResultsInvocationResult) *ListInvocationResultsResponseBodyInvocationResults {
	s.InvocationResult = v
	return s
}

type ListInvocationResultsResponseBodyInvocationResultsInvocationResult struct {
	CommandId          *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	ExitCode           *int32  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	FinishedTime       *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInvocationResultsResponseBodyInvocationResultsInvocationResult) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationResultsResponseBodyInvocationResultsInvocationResult) GoString() string {
	return s.String()
}

func (s *ListInvocationResultsResponseBodyInvocationResultsInvocationResult) SetCommandId(v string) *ListInvocationResultsResponseBodyInvocationResultsInvocationResult {
	s.CommandId = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResultsInvocationResult) SetExitCode(v int32) *ListInvocationResultsResponseBodyInvocationResultsInvocationResult {
	s.ExitCode = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResultsInvocationResult) SetFinishedTime(v string) *ListInvocationResultsResponseBodyInvocationResultsInvocationResult {
	s.FinishedTime = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResultsInvocationResult) SetInstanceId(v string) *ListInvocationResultsResponseBodyInvocationResultsInvocationResult {
	s.InstanceId = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResultsInvocationResult) SetInvokeRecordStatus(v string) *ListInvocationResultsResponseBodyInvocationResultsInvocationResult {
	s.InvokeRecordStatus = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResultsInvocationResult) SetMessage(v string) *ListInvocationResultsResponseBodyInvocationResultsInvocationResult {
	s.Message = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResultsInvocationResult) SetSuccess(v bool) *ListInvocationResultsResponseBodyInvocationResultsInvocationResult {
	s.Success = &v
	return s
}

type ListInvocationResultsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInvocationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInvocationResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationResultsResponse) GoString() string {
	return s.String()
}

func (s *ListInvocationResultsResponse) SetHeaders(v map[string]*string) *ListInvocationResultsResponse {
	s.Headers = v
	return s
}

func (s *ListInvocationResultsResponse) SetBody(v *ListInvocationResultsResponseBody) *ListInvocationResultsResponse {
	s.Body = v
	return s
}

type ListInvocationStatusRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
}

func (s ListInvocationStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationStatusRequest) GoString() string {
	return s.String()
}

func (s *ListInvocationStatusRequest) SetClusterId(v string) *ListInvocationStatusRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInvocationStatusRequest) SetCommandId(v string) *ListInvocationStatusRequest {
	s.CommandId = &v
	return s
}

type ListInvocationStatusResponseBody struct {
	CommandId       *string                                          `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	InvokeInstances *ListInvocationStatusResponseBodyInvokeInstances `json:"InvokeInstances,omitempty" xml:"InvokeInstances,omitempty" type:"Struct"`
	InvokeStatus    *string                                          `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	RequestId       *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInvocationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListInvocationStatusResponseBody) SetCommandId(v string) *ListInvocationStatusResponseBody {
	s.CommandId = &v
	return s
}

func (s *ListInvocationStatusResponseBody) SetInvokeInstances(v *ListInvocationStatusResponseBodyInvokeInstances) *ListInvocationStatusResponseBody {
	s.InvokeInstances = v
	return s
}

func (s *ListInvocationStatusResponseBody) SetInvokeStatus(v string) *ListInvocationStatusResponseBody {
	s.InvokeStatus = &v
	return s
}

func (s *ListInvocationStatusResponseBody) SetRequestId(v string) *ListInvocationStatusResponseBody {
	s.RequestId = &v
	return s
}

type ListInvocationStatusResponseBodyInvokeInstances struct {
	InvokeInstance []*ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance `json:"InvokeInstance,omitempty" xml:"InvokeInstance,omitempty" type:"Repeated"`
}

func (s ListInvocationStatusResponseBodyInvokeInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationStatusResponseBodyInvokeInstances) GoString() string {
	return s.String()
}

func (s *ListInvocationStatusResponseBodyInvokeInstances) SetInvokeInstance(v []*ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance) *ListInvocationStatusResponseBodyInvokeInstances {
	s.InvokeInstance = v
	return s
}

type ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance struct {
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceInvokeStatus *string `json:"InstanceInvokeStatus,omitempty" xml:"InstanceInvokeStatus,omitempty"`
}

func (s ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance) GoString() string {
	return s.String()
}

func (s *ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance) SetInstanceId(v string) *ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance {
	s.InstanceId = &v
	return s
}

func (s *ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance) SetInstanceInvokeStatus(v string) *ListInvocationStatusResponseBodyInvokeInstancesInvokeInstance {
	s.InstanceInvokeStatus = &v
	return s
}

type ListInvocationStatusResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInvocationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInvocationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationStatusResponse) GoString() string {
	return s.String()
}

func (s *ListInvocationStatusResponse) SetHeaders(v map[string]*string) *ListInvocationStatusResponse {
	s.Headers = v
	return s
}

func (s *ListInvocationStatusResponse) SetBody(v *ListInvocationStatusResponseBody) *ListInvocationStatusResponse {
	s.Body = v
	return s
}

type ListJobTemplatesRequest struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListJobTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesRequest) SetName(v string) *ListJobTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListJobTemplatesRequest) SetPageNumber(v int32) *ListJobTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobTemplatesRequest) SetPageSize(v int32) *ListJobTemplatesRequest {
	s.PageSize = &v
	return s
}

type ListJobTemplatesResponseBody struct {
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates  *ListJobTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Struct"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponseBody) SetPageNumber(v int32) *ListJobTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobTemplatesResponseBody) SetPageSize(v int32) *ListJobTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobTemplatesResponseBody) SetRequestId(v string) *ListJobTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobTemplatesResponseBody) SetTemplates(v *ListJobTemplatesResponseBodyTemplates) *ListJobTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *ListJobTemplatesResponseBody) SetTotalCount(v int32) *ListJobTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

type ListJobTemplatesResponseBodyTemplates struct {
	JobTemplates []*ListJobTemplatesResponseBodyTemplatesJobTemplates `json:"JobTemplates,omitempty" xml:"JobTemplates,omitempty" type:"Repeated"`
}

func (s ListJobTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponseBodyTemplates) SetJobTemplates(v []*ListJobTemplatesResponseBodyTemplatesJobTemplates) *ListJobTemplatesResponseBodyTemplates {
	s.JobTemplates = v
	return s
}

type ListJobTemplatesResponseBodyTemplatesJobTemplates struct {
	ArrayRequest       *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	ClockTime          *string `json:"ClockTime,omitempty" xml:"ClockTime,omitempty"`
	CommandLine        *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	Gpu                *int32  `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Id                 *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InputFileUrl       *string `json:"InputFileUrl,omitempty" xml:"InputFileUrl,omitempty"`
	Mem                *string `json:"Mem,omitempty" xml:"Mem,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Node               *int32  `json:"Node,omitempty" xml:"Node,omitempty"`
	PackagePath        *string `json:"PackagePath,omitempty" xml:"PackagePath,omitempty"`
	Priority           *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Queue              *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	ReRunable          *bool   `json:"ReRunable,omitempty" xml:"ReRunable,omitempty"`
	RunasUser          *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" xml:"StderrRedirectPath,omitempty"`
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" xml:"StdoutRedirectPath,omitempty"`
	Task               *int32  `json:"Task,omitempty" xml:"Task,omitempty"`
	Thread             *int32  `json:"Thread,omitempty" xml:"Thread,omitempty"`
	UnzipCmd           *string `json:"UnzipCmd,omitempty" xml:"UnzipCmd,omitempty"`
	Variables          *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
	WithUnzipCmd       *bool   `json:"WithUnzipCmd,omitempty" xml:"WithUnzipCmd,omitempty"`
}

func (s ListJobTemplatesResponseBodyTemplatesJobTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesResponseBodyTemplatesJobTemplates) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetArrayRequest(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.ArrayRequest = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetClockTime(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.ClockTime = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetCommandLine(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.CommandLine = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetGpu(v int32) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Gpu = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetId(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Id = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetInputFileUrl(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.InputFileUrl = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetMem(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Mem = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetName(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Name = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetNode(v int32) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Node = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetPackagePath(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.PackagePath = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetPriority(v int32) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Priority = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetQueue(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Queue = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetReRunable(v bool) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.ReRunable = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetRunasUser(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.RunasUser = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetStderrRedirectPath(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.StderrRedirectPath = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetStdoutRedirectPath(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.StdoutRedirectPath = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetTask(v int32) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Task = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetThread(v int32) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Thread = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetUnzipCmd(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.UnzipCmd = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetVariables(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Variables = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetWithUnzipCmd(v bool) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.WithUnzipCmd = &v
	return s
}

type ListJobTemplatesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListJobTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponse) SetHeaders(v map[string]*string) *ListJobTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListJobTemplatesResponse) SetBody(v *ListJobTemplatesResponseBody) *ListJobTemplatesResponse {
	s.Body = v
	return s
}

type ListJobsRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Owner      *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Rerunable  *string `json:"Rerunable,omitempty" xml:"Rerunable,omitempty"`
	State      *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) SetClusterId(v string) *ListJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListJobsRequest) SetOwner(v string) *ListJobsRequest {
	s.Owner = &v
	return s
}

func (s *ListJobsRequest) SetPageNumber(v int32) *ListJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) SetRerunable(v string) *ListJobsRequest {
	s.Rerunable = &v
	return s
}

func (s *ListJobsRequest) SetState(v string) *ListJobsRequest {
	s.State = &v
	return s
}

type ListJobsResponseBody struct {
	Jobs       *ListJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Struct"`
	PageNumber *int32                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) SetJobs(v *ListJobsResponseBodyJobs) *ListJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListJobsResponseBody) SetPageNumber(v int32) *ListJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobsResponseBody) SetPageSize(v int32) *ListJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalCount(v int32) *ListJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListJobsResponseBodyJobs struct {
	JobInfo []*ListJobsResponseBodyJobsJobInfo `json:"JobInfo,omitempty" xml:"JobInfo,omitempty" type:"Repeated"`
}

func (s ListJobsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobs) SetJobInfo(v []*ListJobsResponseBodyJobsJobInfo) *ListJobsResponseBodyJobs {
	s.JobInfo = v
	return s
}

type ListJobsResponseBodyJobsJobInfo struct {
	ArrayRequest   *string                                   `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	Comment        *string                                   `json:"Comment,omitempty" xml:"Comment,omitempty"`
	Id             *string                                   `json:"Id,omitempty" xml:"Id,omitempty"`
	LastModifyTime *string                                   `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	Name           *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	NodeList       *string                                   `json:"NodeList,omitempty" xml:"NodeList,omitempty"`
	Owner          *string                                   `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Priority       *string                                   `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Resources      *ListJobsResponseBodyJobsJobInfoResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	ShellPath      *string                                   `json:"ShellPath,omitempty" xml:"ShellPath,omitempty"`
	StartTime      *string                                   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State          *string                                   `json:"State,omitempty" xml:"State,omitempty"`
	Stderr         *string                                   `json:"Stderr,omitempty" xml:"Stderr,omitempty"`
	Stdout         *string                                   `json:"Stdout,omitempty" xml:"Stdout,omitempty"`
	SubmitTime     *string                                   `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s ListJobsResponseBodyJobsJobInfo) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobsJobInfo) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsJobInfo) SetArrayRequest(v string) *ListJobsResponseBodyJobsJobInfo {
	s.ArrayRequest = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetComment(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Comment = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetId(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Id = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetLastModifyTime(v string) *ListJobsResponseBodyJobsJobInfo {
	s.LastModifyTime = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetName(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Name = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetNodeList(v string) *ListJobsResponseBodyJobsJobInfo {
	s.NodeList = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetOwner(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Owner = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetPriority(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Priority = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetResources(v *ListJobsResponseBodyJobsJobInfoResources) *ListJobsResponseBodyJobsJobInfo {
	s.Resources = v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetShellPath(v string) *ListJobsResponseBodyJobsJobInfo {
	s.ShellPath = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetStartTime(v string) *ListJobsResponseBodyJobsJobInfo {
	s.StartTime = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetState(v string) *ListJobsResponseBodyJobsJobInfo {
	s.State = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetStderr(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Stderr = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetStdout(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Stdout = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetSubmitTime(v string) *ListJobsResponseBodyJobsJobInfo {
	s.SubmitTime = &v
	return s
}

type ListJobsResponseBodyJobsJobInfoResources struct {
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	Nodes *int32 `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s ListJobsResponseBodyJobsJobInfoResources) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobsJobInfoResources) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsJobInfoResources) SetCores(v int32) *ListJobsResponseBodyJobsJobInfoResources {
	s.Cores = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfoResources) SetNodes(v int32) *ListJobsResponseBodyJobsJobInfoResources {
	s.Nodes = &v
	return s
}

type ListJobsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponse) GoString() string {
	return s.String()
}

func (s *ListJobsResponse) SetHeaders(v map[string]*string) *ListJobsResponse {
	s.Headers = v
	return s
}

func (s *ListJobsResponse) SetBody(v *ListJobsResponseBody) *ListJobsResponse {
	s.Body = v
	return s
}

type ListNodesRequest struct {
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Filter           *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	HostName         *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HostNamePrefix   *string `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	HostNameSuffix   *string `json:"HostNameSuffix,omitempty" xml:"HostNameSuffix,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	Role             *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Sequence         *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	SortBy           *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) SetClusterId(v string) *ListNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodesRequest) SetFilter(v string) *ListNodesRequest {
	s.Filter = &v
	return s
}

func (s *ListNodesRequest) SetHostName(v string) *ListNodesRequest {
	s.HostName = &v
	return s
}

func (s *ListNodesRequest) SetHostNamePrefix(v string) *ListNodesRequest {
	s.HostNamePrefix = &v
	return s
}

func (s *ListNodesRequest) SetHostNameSuffix(v string) *ListNodesRequest {
	s.HostNameSuffix = &v
	return s
}

func (s *ListNodesRequest) SetPageNumber(v int32) *ListNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesRequest) SetPageSize(v int32) *ListNodesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesRequest) SetPrivateIpAddress(v string) *ListNodesRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *ListNodesRequest) SetRole(v string) *ListNodesRequest {
	s.Role = &v
	return s
}

func (s *ListNodesRequest) SetSequence(v string) *ListNodesRequest {
	s.Sequence = &v
	return s
}

func (s *ListNodesRequest) SetSortBy(v string) *ListNodesRequest {
	s.SortBy = &v
	return s
}

type ListNodesResponseBody struct {
	Nodes      *ListNodesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
	PageNumber *int32                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) SetNodes(v *ListNodesResponseBodyNodes) *ListNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *ListNodesResponseBody) SetPageNumber(v int32) *ListNodesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNodesResponseBody) SetPageSize(v int32) *ListNodesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesResponseBody) SetTotalCount(v int32) *ListNodesResponseBody {
	s.TotalCount = &v
	return s
}

type ListNodesResponseBodyNodes struct {
	NodeInfo []*ListNodesResponseBodyNodesNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodes) SetNodeInfo(v []*ListNodesResponseBodyNodesNodeInfo) *ListNodesResponseBodyNodes {
	s.NodeInfo = v
	return s
}

type ListNodesResponseBodyNodesNodeInfo struct {
	AddTime         *string                                           `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	CreateMode      *string                                           `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreatedByEhpc   *bool                                             `json:"CreatedByEhpc,omitempty" xml:"CreatedByEhpc,omitempty"`
	Expired         *bool                                             `json:"Expired,omitempty" xml:"Expired,omitempty"`
	ExpiredTime     *string                                           `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	HostName        *string                                           `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HtEnabled       *bool                                             `json:"HtEnabled,omitempty" xml:"HtEnabled,omitempty"`
	Id              *string                                           `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageId         *string                                           `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageOwnerAlias *string                                           `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	InstanceType    *string                                           `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IpAddress       *string                                           `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Location        *string                                           `json:"Location,omitempty" xml:"Location,omitempty"`
	LockReason      *string                                           `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	PublicIpAddress *string                                           `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	RegionId        *string                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Roles           *ListNodesResponseBodyNodesNodeInfoRoles          `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
	SpotStrategy    *string                                           `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	StateInSched    *string                                           `json:"StateInSched,omitempty" xml:"StateInSched,omitempty"`
	Status          *string                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalResources  *ListNodesResponseBodyNodesNodeInfoTotalResources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty" type:"Struct"`
	UsedResources   *ListNodesResponseBodyNodesNodeInfoUsedResources  `json:"UsedResources,omitempty" xml:"UsedResources,omitempty" type:"Struct"`
	VSwitchId       *string                                           `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Version         *string                                           `json:"Version,omitempty" xml:"Version,omitempty"`
	VpcId           *string                                           `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId          *string                                           `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListNodesResponseBodyNodesNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodesNodeInfo) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetAddTime(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.AddTime = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetCreateMode(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.CreateMode = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetCreatedByEhpc(v bool) *ListNodesResponseBodyNodesNodeInfo {
	s.CreatedByEhpc = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetExpired(v bool) *ListNodesResponseBodyNodesNodeInfo {
	s.Expired = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetExpiredTime(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.ExpiredTime = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetHostName(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.HostName = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetHtEnabled(v bool) *ListNodesResponseBodyNodesNodeInfo {
	s.HtEnabled = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetId(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetImageId(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.ImageId = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetImageOwnerAlias(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetInstanceType(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.InstanceType = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetIpAddress(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.IpAddress = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetLocation(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.Location = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetLockReason(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.LockReason = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetPublicIpAddress(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.PublicIpAddress = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetRegionId(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.RegionId = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetRoles(v *ListNodesResponseBodyNodesNodeInfoRoles) *ListNodesResponseBodyNodesNodeInfo {
	s.Roles = v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetSpotStrategy(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.SpotStrategy = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetStateInSched(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.StateInSched = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetStatus(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.Status = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetTotalResources(v *ListNodesResponseBodyNodesNodeInfoTotalResources) *ListNodesResponseBodyNodesNodeInfo {
	s.TotalResources = v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetUsedResources(v *ListNodesResponseBodyNodesNodeInfoUsedResources) *ListNodesResponseBodyNodesNodeInfo {
	s.UsedResources = v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetVSwitchId(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.VSwitchId = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetVersion(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.Version = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetVpcId(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.VpcId = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetZoneId(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.ZoneId = &v
	return s
}

type ListNodesResponseBodyNodesNodeInfoRoles struct {
	Role []*string `json:"Role,omitempty" xml:"Role,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBodyNodesNodeInfoRoles) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodesNodeInfoRoles) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodesNodeInfoRoles) SetRole(v []*string) *ListNodesResponseBodyNodesNodeInfoRoles {
	s.Role = v
	return s
}

type ListNodesResponseBodyNodesNodeInfoTotalResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesResponseBodyNodesNodeInfoTotalResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodesNodeInfoTotalResources) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodesNodeInfoTotalResources) SetCpu(v int32) *ListNodesResponseBodyNodesNodeInfoTotalResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfoTotalResources) SetGpu(v int32) *ListNodesResponseBodyNodesNodeInfoTotalResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfoTotalResources) SetMemory(v int32) *ListNodesResponseBodyNodesNodeInfoTotalResources {
	s.Memory = &v
	return s
}

type ListNodesResponseBodyNodesNodeInfoUsedResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesResponseBodyNodesNodeInfoUsedResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodesNodeInfoUsedResources) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodesNodeInfoUsedResources) SetCpu(v int32) *ListNodesResponseBodyNodesNodeInfoUsedResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfoUsedResources) SetGpu(v int32) *ListNodesResponseBodyNodesNodeInfoUsedResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfoUsedResources) SetMemory(v int32) *ListNodesResponseBodyNodesNodeInfoUsedResources {
	s.Memory = &v
	return s
}

type ListNodesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponse) GoString() string {
	return s.String()
}

func (s *ListNodesResponse) SetHeaders(v map[string]*string) *ListNodesResponse {
	s.Headers = v
	return s
}

func (s *ListNodesResponse) SetBody(v *ListNodesResponseBody) *ListNodesResponse {
	s.Body = v
	return s
}

type ListNodesByQueueRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueueName  *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s ListNodesByQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueRequest) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueRequest) SetClusterId(v string) *ListNodesByQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodesByQueueRequest) SetPageNumber(v int32) *ListNodesByQueueRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesByQueueRequest) SetPageSize(v int32) *ListNodesByQueueRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesByQueueRequest) SetQueueName(v string) *ListNodesByQueueRequest {
	s.QueueName = &v
	return s
}

type ListNodesByQueueResponseBody struct {
	Nodes      *ListNodesByQueueResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesByQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueResponseBody) SetNodes(v *ListNodesByQueueResponseBodyNodes) *ListNodesByQueueResponseBody {
	s.Nodes = v
	return s
}

func (s *ListNodesByQueueResponseBody) SetPageNumber(v int32) *ListNodesByQueueResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNodesByQueueResponseBody) SetPageSize(v int32) *ListNodesByQueueResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNodesByQueueResponseBody) SetRequestId(v string) *ListNodesByQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesByQueueResponseBody) SetTotalCount(v int32) *ListNodesByQueueResponseBody {
	s.TotalCount = &v
	return s
}

type ListNodesByQueueResponseBodyNodes struct {
	NodeInfo []*ListNodesByQueueResponseBodyNodesNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
}

func (s ListNodesByQueueResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueResponseBodyNodes) SetNodeInfo(v []*ListNodesByQueueResponseBodyNodesNodeInfo) *ListNodesByQueueResponseBodyNodes {
	s.NodeInfo = v
	return s
}

type ListNodesByQueueResponseBodyNodesNodeInfo struct {
	AddTime         *string                                                  `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	CreateMode      *string                                                  `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	CreatedByEhpc   *bool                                                    `json:"CreatedByEhpc,omitempty" xml:"CreatedByEhpc,omitempty"`
	Expired         *bool                                                    `json:"Expired,omitempty" xml:"Expired,omitempty"`
	ExpiredTime     *string                                                  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	HostName        *string                                                  `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HtEnabled       *bool                                                    `json:"HtEnabled,omitempty" xml:"HtEnabled,omitempty"`
	Id              *string                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageId         *string                                                  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageOwnerAlias *string                                                  `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	IpAddress       *string                                                  `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Location        *string                                                  `json:"Location,omitempty" xml:"Location,omitempty"`
	LockReason      *string                                                  `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	PublicIpAddress *string                                                  `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	RegionId        *string                                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SpotStrategy    *string                                                  `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	StateInSched    *string                                                  `json:"StateInSched,omitempty" xml:"StateInSched,omitempty"`
	Status          *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalResources  *ListNodesByQueueResponseBodyNodesNodeInfoTotalResources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty" type:"Struct"`
	UsedResources   *ListNodesByQueueResponseBodyNodesNodeInfoUsedResources  `json:"UsedResources,omitempty" xml:"UsedResources,omitempty" type:"Struct"`
	VSwitchId       *string                                                  `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Version         *string                                                  `json:"Version,omitempty" xml:"Version,omitempty"`
	VpcId           *string                                                  `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId          *string                                                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListNodesByQueueResponseBodyNodesNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueResponseBodyNodesNodeInfo) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetAddTime(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.AddTime = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetCreateMode(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.CreateMode = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetCreatedByEhpc(v bool) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.CreatedByEhpc = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetExpired(v bool) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.Expired = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetExpiredTime(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.ExpiredTime = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetHostName(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.HostName = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetHtEnabled(v bool) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.HtEnabled = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetId(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.Id = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetImageId(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.ImageId = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetImageOwnerAlias(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetIpAddress(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.IpAddress = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetLocation(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.Location = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetLockReason(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.LockReason = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetPublicIpAddress(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.PublicIpAddress = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetRegionId(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.RegionId = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetSpotStrategy(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.SpotStrategy = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetStateInSched(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.StateInSched = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetStatus(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.Status = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetTotalResources(v *ListNodesByQueueResponseBodyNodesNodeInfoTotalResources) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.TotalResources = v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetUsedResources(v *ListNodesByQueueResponseBodyNodesNodeInfoUsedResources) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.UsedResources = v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetVSwitchId(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.VSwitchId = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetVersion(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.Version = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetVpcId(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.VpcId = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfo) SetZoneId(v string) *ListNodesByQueueResponseBodyNodesNodeInfo {
	s.ZoneId = &v
	return s
}

type ListNodesByQueueResponseBodyNodesNodeInfoTotalResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesByQueueResponseBodyNodesNodeInfoTotalResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueResponseBodyNodesNodeInfoTotalResources) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfoTotalResources) SetCpu(v int32) *ListNodesByQueueResponseBodyNodesNodeInfoTotalResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfoTotalResources) SetGpu(v int32) *ListNodesByQueueResponseBodyNodesNodeInfoTotalResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfoTotalResources) SetMemory(v int32) *ListNodesByQueueResponseBodyNodesNodeInfoTotalResources {
	s.Memory = &v
	return s
}

type ListNodesByQueueResponseBodyNodesNodeInfoUsedResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesByQueueResponseBodyNodesNodeInfoUsedResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueResponseBodyNodesNodeInfoUsedResources) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfoUsedResources) SetCpu(v int32) *ListNodesByQueueResponseBodyNodesNodeInfoUsedResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfoUsedResources) SetGpu(v int32) *ListNodesByQueueResponseBodyNodesNodeInfoUsedResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesNodeInfoUsedResources) SetMemory(v int32) *ListNodesByQueueResponseBodyNodesNodeInfoUsedResources {
	s.Memory = &v
	return s
}

type ListNodesByQueueResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNodesByQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodesByQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueResponse) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueResponse) SetHeaders(v map[string]*string) *ListNodesByQueueResponse {
	s.Headers = v
	return s
}

func (s *ListNodesByQueueResponse) SetBody(v *ListNodesByQueueResponseBody) *ListNodesByQueueResponse {
	s.Body = v
	return s
}

type ListNodesNoPagingRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostName  *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Role      *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Sequence  *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListNodesNoPagingRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingRequest) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingRequest) SetClusterId(v string) *ListNodesNoPagingRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodesNoPagingRequest) SetHostName(v string) *ListNodesNoPagingRequest {
	s.HostName = &v
	return s
}

func (s *ListNodesNoPagingRequest) SetRole(v string) *ListNodesNoPagingRequest {
	s.Role = &v
	return s
}

func (s *ListNodesNoPagingRequest) SetSequence(v string) *ListNodesNoPagingRequest {
	s.Sequence = &v
	return s
}

type ListNodesNoPagingResponseBody struct {
	Nodes     *ListNodesNoPagingResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodesNoPagingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponseBody) SetNodes(v *ListNodesNoPagingResponseBodyNodes) *ListNodesNoPagingResponseBody {
	s.Nodes = v
	return s
}

func (s *ListNodesNoPagingResponseBody) SetRequestId(v string) *ListNodesNoPagingResponseBody {
	s.RequestId = &v
	return s
}

type ListNodesNoPagingResponseBodyNodes struct {
	NodeInfo []*ListNodesNoPagingResponseBodyNodesNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
}

func (s ListNodesNoPagingResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponseBodyNodes) SetNodeInfo(v []*ListNodesNoPagingResponseBodyNodesNodeInfo) *ListNodesNoPagingResponseBodyNodes {
	s.NodeInfo = v
	return s
}

type ListNodesNoPagingResponseBodyNodesNodeInfo struct {
	HostName     *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListNodesNoPagingResponseBodyNodesNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponseBodyNodesNodeInfo) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetHostName(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.HostName = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetId(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.Id = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetImageId(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.ImageId = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetInstanceType(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.InstanceType = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetStatus(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.Status = &v
	return s
}

type ListNodesNoPagingResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNodesNoPagingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodesNoPagingResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponse) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponse) SetHeaders(v map[string]*string) *ListNodesNoPagingResponse {
	s.Headers = v
	return s
}

func (s *ListNodesNoPagingResponse) SetBody(v *ListNodesNoPagingResponseBody) *ListNodesNoPagingResponse {
	s.Body = v
	return s
}

type ListPreferredEcsTypesRequest struct {
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	SpotStrategy       *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	ZoneId             *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListPreferredEcsTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesRequest) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesRequest) SetInstanceChargeType(v string) *ListPreferredEcsTypesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *ListPreferredEcsTypesRequest) SetSpotStrategy(v string) *ListPreferredEcsTypesRequest {
	s.SpotStrategy = &v
	return s
}

func (s *ListPreferredEcsTypesRequest) SetZoneId(v string) *ListPreferredEcsTypesRequest {
	s.ZoneId = &v
	return s
}

type ListPreferredEcsTypesResponseBody struct {
	RequestId           *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Series              *ListPreferredEcsTypesResponseBodySeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Struct"`
	SupportSpotInstance *bool                                    `json:"SupportSpotInstance,omitempty" xml:"SupportSpotInstance,omitempty"`
}

func (s ListPreferredEcsTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBody) SetRequestId(v string) *ListPreferredEcsTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPreferredEcsTypesResponseBody) SetSeries(v *ListPreferredEcsTypesResponseBodySeries) *ListPreferredEcsTypesResponseBody {
	s.Series = v
	return s
}

func (s *ListPreferredEcsTypesResponseBody) SetSupportSpotInstance(v bool) *ListPreferredEcsTypesResponseBody {
	s.SupportSpotInstance = &v
	return s
}

type ListPreferredEcsTypesResponseBodySeries struct {
	SeriesInfo []*ListPreferredEcsTypesResponseBodySeriesSeriesInfo `json:"SeriesInfo,omitempty" xml:"SeriesInfo,omitempty" type:"Repeated"`
}

func (s ListPreferredEcsTypesResponseBodySeries) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeries) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeries) SetSeriesInfo(v []*ListPreferredEcsTypesResponseBodySeriesSeriesInfo) *ListPreferredEcsTypesResponseBodySeries {
	s.SeriesInfo = v
	return s
}

type ListPreferredEcsTypesResponseBodySeriesSeriesInfo struct {
	Roles      *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
	SeriesId   *string                                                 `json:"SeriesId,omitempty" xml:"SeriesId,omitempty"`
	SeriesName *string                                                 `json:"SeriesName,omitempty" xml:"SeriesName,omitempty"`
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfo) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfo) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfo) SetRoles(v *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) *ListPreferredEcsTypesResponseBodySeriesSeriesInfo {
	s.Roles = v
	return s
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfo) SetSeriesId(v string) *ListPreferredEcsTypesResponseBodySeriesSeriesInfo {
	s.SeriesId = &v
	return s
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfo) SetSeriesName(v string) *ListPreferredEcsTypesResponseBodySeriesSeriesInfo {
	s.SeriesName = &v
	return s
}

type ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles struct {
	Compute *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute `json:"Compute,omitempty" xml:"Compute,omitempty" type:"Struct"`
	Login   *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin   `json:"Login,omitempty" xml:"Login,omitempty" type:"Struct"`
	Manager *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager `json:"Manager,omitempty" xml:"Manager,omitempty" type:"Struct"`
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) SetCompute(v *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles {
	s.Compute = v
	return s
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) SetLogin(v *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles {
	s.Login = v
	return s
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) SetManager(v *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles {
	s.Manager = v
	return s
}

type ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute struct {
	InstanceTypeId []*string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty" type:"Repeated"`
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute) SetInstanceTypeId(v []*string) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute {
	s.InstanceTypeId = v
	return s
}

type ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin struct {
	InstanceTypeId []*string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty" type:"Repeated"`
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin) SetInstanceTypeId(v []*string) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin {
	s.InstanceTypeId = v
	return s
}

type ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager struct {
	InstanceTypeId []*string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty" type:"Repeated"`
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager) SetInstanceTypeId(v []*string) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager {
	s.InstanceTypeId = v
	return s
}

type ListPreferredEcsTypesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPreferredEcsTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPreferredEcsTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponse) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponse) SetHeaders(v map[string]*string) *ListPreferredEcsTypesResponse {
	s.Headers = v
	return s
}

func (s *ListPreferredEcsTypesResponse) SetBody(v *ListPreferredEcsTypesResponseBody) *ListPreferredEcsTypesResponse {
	s.Body = v
	return s
}

type ListQueuesRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListQueuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesRequest) GoString() string {
	return s.String()
}

func (s *ListQueuesRequest) SetClusterId(v string) *ListQueuesRequest {
	s.ClusterId = &v
	return s
}

type ListQueuesResponseBody struct {
	Queues    *ListQueuesResponseBodyQueues `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListQueuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBody) SetQueues(v *ListQueuesResponseBodyQueues) *ListQueuesResponseBody {
	s.Queues = v
	return s
}

func (s *ListQueuesResponseBody) SetRequestId(v string) *ListQueuesResponseBody {
	s.RequestId = &v
	return s
}

type ListQueuesResponseBodyQueues struct {
	QueueInfo []*ListQueuesResponseBodyQueuesQueueInfo `json:"QueueInfo,omitempty" xml:"QueueInfo,omitempty" type:"Repeated"`
}

func (s ListQueuesResponseBodyQueues) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyQueues) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyQueues) SetQueueInfo(v []*ListQueuesResponseBodyQueuesQueueInfo) *ListQueuesResponseBodyQueues {
	s.QueueInfo = v
	return s
}

type ListQueuesResponseBodyQueuesQueueInfo struct {
	ComputeInstanceType *ListQueuesResponseBodyQueuesQueueInfoComputeInstanceType `json:"ComputeInstanceType,omitempty" xml:"ComputeInstanceType,omitempty" type:"Struct"`
	EnableAutoGrow      *bool                                                     `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	HostNamePrefix      *string                                                   `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	HostNameSuffix      *string                                                   `json:"HostNameSuffix,omitempty" xml:"HostNameSuffix,omitempty"`
	ImageId             *string                                                   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	QueueName           *string                                                   `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	ResourceGroupId     *string                                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SpotInstanceTypes   *ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypes   `json:"SpotInstanceTypes,omitempty" xml:"SpotInstanceTypes,omitempty" type:"Struct"`
	SpotStrategy        *string                                                   `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	Type                *string                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListQueuesResponseBodyQueuesQueueInfo) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyQueuesQueueInfo) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetComputeInstanceType(v *ListQueuesResponseBodyQueuesQueueInfoComputeInstanceType) *ListQueuesResponseBodyQueuesQueueInfo {
	s.ComputeInstanceType = v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetEnableAutoGrow(v bool) *ListQueuesResponseBodyQueuesQueueInfo {
	s.EnableAutoGrow = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetHostNamePrefix(v string) *ListQueuesResponseBodyQueuesQueueInfo {
	s.HostNamePrefix = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetHostNameSuffix(v string) *ListQueuesResponseBodyQueuesQueueInfo {
	s.HostNameSuffix = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetImageId(v string) *ListQueuesResponseBodyQueuesQueueInfo {
	s.ImageId = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetQueueName(v string) *ListQueuesResponseBodyQueuesQueueInfo {
	s.QueueName = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetResourceGroupId(v string) *ListQueuesResponseBodyQueuesQueueInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetSpotInstanceTypes(v *ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypes) *ListQueuesResponseBodyQueuesQueueInfo {
	s.SpotInstanceTypes = v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetSpotStrategy(v string) *ListQueuesResponseBodyQueuesQueueInfo {
	s.SpotStrategy = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfo) SetType(v string) *ListQueuesResponseBodyQueuesQueueInfo {
	s.Type = &v
	return s
}

type ListQueuesResponseBodyQueuesQueueInfoComputeInstanceType struct {
	InstanceType []*string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Repeated"`
}

func (s ListQueuesResponseBodyQueuesQueueInfoComputeInstanceType) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyQueuesQueueInfoComputeInstanceType) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyQueuesQueueInfoComputeInstanceType) SetInstanceType(v []*string) *ListQueuesResponseBodyQueuesQueueInfoComputeInstanceType {
	s.InstanceType = v
	return s
}

type ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypes struct {
	Instance []*ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypes) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypes) SetInstance(v []*ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance) *ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypes {
	s.Instance = v
	return s
}

type ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance struct {
	InstanceType   *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
}

func (s ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance) SetInstanceType(v string) *ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance {
	s.InstanceType = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance) SetSpotPriceLimit(v float32) *ListQueuesResponseBodyQueuesQueueInfoSpotInstanceTypesInstance {
	s.SpotPriceLimit = &v
	return s
}

type ListQueuesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListQueuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQueuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponse) GoString() string {
	return s.String()
}

func (s *ListQueuesResponse) SetHeaders(v map[string]*string) *ListQueuesResponse {
	s.Headers = v
	return s
}

func (s *ListQueuesResponse) SetBody(v *ListQueuesResponseBody) *ListQueuesResponse {
	s.Body = v
	return s
}

type ListRegionsResponseBody struct {
	Regions   *ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetRegions(v *ListRegionsResponseBodyRegions) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListRegionsResponseBodyRegions struct {
	RegionInfo []*ListRegionsResponseBodyRegionsRegionInfo `json:"RegionInfo,omitempty" xml:"RegionInfo,omitempty" type:"Repeated"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) SetRegionInfo(v []*ListRegionsResponseBodyRegionsRegionInfo) *ListRegionsResponseBodyRegions {
	s.RegionInfo = v
	return s
}

type ListRegionsResponseBodyRegionsRegionInfo struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsResponseBodyRegionsRegionInfo) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegionsRegionInfo) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegionsRegionInfo) SetLocalName(v string) *ListRegionsResponseBodyRegionsRegionInfo {
	s.LocalName = &v
	return s
}

func (s *ListRegionsResponseBodyRegionsRegionInfo) SetRegionId(v string) *ListRegionsResponseBodyRegionsRegionInfo {
	s.RegionId = &v
	return s
}

type ListRegionsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListSecurityGroupsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListSecurityGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupsRequest) SetClusterId(v string) *ListSecurityGroupsRequest {
	s.ClusterId = &v
	return s
}

type ListSecurityGroupsResponseBody struct {
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityGroups *ListSecurityGroupsResponseBodySecurityGroups `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Struct"`
	TotalCount     *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSecurityGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupsResponseBody) SetRequestId(v string) *ListSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecurityGroupsResponseBody) SetSecurityGroups(v *ListSecurityGroupsResponseBodySecurityGroups) *ListSecurityGroupsResponseBody {
	s.SecurityGroups = v
	return s
}

func (s *ListSecurityGroupsResponseBody) SetTotalCount(v int32) *ListSecurityGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSecurityGroupsResponseBodySecurityGroups struct {
	SecurityGroup []*string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty" type:"Repeated"`
}

func (s ListSecurityGroupsResponseBodySecurityGroups) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityGroupsResponseBodySecurityGroups) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupsResponseBodySecurityGroups) SetSecurityGroup(v []*string) *ListSecurityGroupsResponseBodySecurityGroups {
	s.SecurityGroup = v
	return s
}

type ListSecurityGroupsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSecurityGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupsResponse) SetHeaders(v map[string]*string) *ListSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityGroupsResponse) SetBody(v *ListSecurityGroupsResponseBody) *ListSecurityGroupsResponse {
	s.Body = v
	return s
}

type ListSoftwaresRequest struct {
	EhpcVersion *string `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	OsTag       *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
}

func (s ListSoftwaresRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresRequest) GoString() string {
	return s.String()
}

func (s *ListSoftwaresRequest) SetEhpcVersion(v string) *ListSoftwaresRequest {
	s.EhpcVersion = &v
	return s
}

func (s *ListSoftwaresRequest) SetOsTag(v string) *ListSoftwaresRequest {
	s.OsTag = &v
	return s
}

type ListSoftwaresResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Softwares *ListSoftwaresResponseBodySoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Struct"`
}

func (s ListSoftwaresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBody) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBody) SetRequestId(v string) *ListSoftwaresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSoftwaresResponseBody) SetSoftwares(v *ListSoftwaresResponseBodySoftwares) *ListSoftwaresResponseBody {
	s.Softwares = v
	return s
}

type ListSoftwaresResponseBodySoftwares struct {
	SoftwareInfo []*ListSoftwaresResponseBodySoftwaresSoftwareInfo `json:"SoftwareInfo,omitempty" xml:"SoftwareInfo,omitempty" type:"Repeated"`
}

func (s ListSoftwaresResponseBodySoftwares) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodySoftwares) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodySoftwares) SetSoftwareInfo(v []*ListSoftwaresResponseBodySoftwaresSoftwareInfo) *ListSoftwaresResponseBodySoftwares {
	s.SoftwareInfo = v
	return s
}

type ListSoftwaresResponseBodySoftwaresSoftwareInfo struct {
	AccountType      *string                                                     `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AccountVersion   *string                                                     `json:"AccountVersion,omitempty" xml:"AccountVersion,omitempty"`
	Applications     *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	EhpcVersion      *string                                                     `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	OsTag            *string                                                     `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	SchedulerType    *string                                                     `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	SchedulerVersion *string                                                     `json:"SchedulerVersion,omitempty" xml:"SchedulerVersion,omitempty"`
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfo) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetAccountType(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.AccountType = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetAccountVersion(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.AccountVersion = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetApplications(v *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.Applications = v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetEhpcVersion(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.EhpcVersion = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetOsTag(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.OsTag = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetSchedulerType(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.SchedulerType = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetSchedulerVersion(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.SchedulerVersion = &v
	return s
}

type ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications struct {
	ApplicationInfo []*ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo `json:"ApplicationInfo,omitempty" xml:"ApplicationInfo,omitempty" type:"Repeated"`
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications) SetApplicationInfo(v []*ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications {
	s.ApplicationInfo = v
	return s
}

type ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Required *bool   `json:"Required,omitempty" xml:"Required,omitempty"`
	Tag      *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Version  *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) SetName(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo {
	s.Name = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) SetRequired(v bool) *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo {
	s.Required = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) SetTag(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo {
	s.Tag = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) SetVersion(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo {
	s.Version = &v
	return s
}

type ListSoftwaresResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSoftwaresResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSoftwaresResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponse) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponse) SetHeaders(v map[string]*string) *ListSoftwaresResponse {
	s.Headers = v
	return s
}

func (s *ListSoftwaresResponse) SetBody(v *ListSoftwaresResponseBody) *ListSoftwaresResponse {
	s.Body = v
	return s
}

type ListTasksRequest struct {
	Archived   *bool   `json:"Archived,omitempty" xml:"Archived,omitempty"`
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) SetArchived(v bool) *ListTasksRequest {
	s.Archived = &v
	return s
}

func (s *ListTasksRequest) SetClusterId(v string) *ListTasksRequest {
	s.ClusterId = &v
	return s
}

func (s *ListTasksRequest) SetPageNumber(v int32) *ListTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTasksRequest) SetPageSize(v int32) *ListTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksRequest) SetTaskId(v string) *ListTasksRequest {
	s.TaskId = &v
	return s
}

type ListTasksResponseBody struct {
	PageNumber *int32                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks      []*ListTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) SetPageNumber(v int32) *ListTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTasksResponseBody) SetPageSize(v int32) *ListTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetTasks(v []*ListTasksResponseBodyTasks) *ListTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListTasksResponseBody) SetTotalCount(v int32) *ListTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListTasksResponseBodyTasks struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CurrentStep *int32  `json:"CurrentStep,omitempty" xml:"CurrentStep,omitempty"`
	Errors      *string `json:"Errors,omitempty" xml:"Errors,omitempty"`
	Request     *string `json:"Request,omitempty" xml:"Request,omitempty"`
	Result      *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType    *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TotalSteps  *int32  `json:"TotalSteps,omitempty" xml:"TotalSteps,omitempty"`
}

func (s ListTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTasks) SetClusterId(v string) *ListTasksResponseBodyTasks {
	s.ClusterId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetCurrentStep(v int32) *ListTasksResponseBodyTasks {
	s.CurrentStep = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetErrors(v string) *ListTasksResponseBodyTasks {
	s.Errors = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetRequest(v string) *ListTasksResponseBodyTasks {
	s.Request = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetResult(v string) *ListTasksResponseBodyTasks {
	s.Result = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetStatus(v string) *ListTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTaskId(v string) *ListTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTaskType(v string) *ListTasksResponseBodyTasks {
	s.TaskType = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTotalSteps(v int32) *ListTasksResponseBodyTasks {
	s.TotalSteps = &v
	return s
}

type ListTasksResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTasksResponse) SetHeaders(v map[string]*string) *ListTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTasksResponse) SetBody(v *ListTasksResponseBody) *ListTasksResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetClusterId(v string) *ListUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *ListUsersRequest) SetPageNumber(v int32) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

type ListUsersResponseBody struct {
	PageNumber *int32                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Users      *ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetPageNumber(v int32) *ListUsersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBody) SetPageSize(v int32) *ListUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int32) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v *ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

type ListUsersResponseBodyUsers struct {
	UserInfo []*ListUsersResponseBodyUsersUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) SetUserInfo(v []*ListUsersResponseBodyUsersUserInfo) *ListUsersResponseBodyUsers {
	s.UserInfo = v
	return s
}

type ListUsersResponseBodyUsersUserInfo struct {
	AddTime *string `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	Group   *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListUsersResponseBodyUsersUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsersUserInfo) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsersUserInfo) SetAddTime(v string) *ListUsersResponseBodyUsersUserInfo {
	s.AddTime = &v
	return s
}

func (s *ListUsersResponseBodyUsersUserInfo) SetGroup(v string) *ListUsersResponseBodyUsersUserInfo {
	s.Group = &v
	return s
}

func (s *ListUsersResponseBodyUsersUserInfo) SetName(v string) *ListUsersResponseBodyUsersUserInfo {
	s.Name = &v
	return s
}

type ListUsersResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ListVolumesRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVolumesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesRequest) GoString() string {
	return s.String()
}

func (s *ListVolumesRequest) SetPageNumber(v int32) *ListVolumesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVolumesRequest) SetPageSize(v int32) *ListVolumesRequest {
	s.PageSize = &v
	return s
}

type ListVolumesResponseBody struct {
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Volumes    *ListVolumesResponseBodyVolumes `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Struct"`
}

func (s ListVolumesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVolumesResponseBody) SetPageNumber(v int32) *ListVolumesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListVolumesResponseBody) SetPageSize(v int32) *ListVolumesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVolumesResponseBody) SetRequestId(v string) *ListVolumesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVolumesResponseBody) SetTotalCount(v int32) *ListVolumesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVolumesResponseBody) SetVolumes(v *ListVolumesResponseBodyVolumes) *ListVolumesResponseBody {
	s.Volumes = v
	return s
}

type ListVolumesResponseBodyVolumes struct {
	VolumeInfo []*ListVolumesResponseBodyVolumesVolumeInfo `json:"VolumeInfo,omitempty" xml:"VolumeInfo,omitempty" type:"Repeated"`
}

func (s ListVolumesResponseBodyVolumes) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponseBodyVolumes) GoString() string {
	return s.String()
}

func (s *ListVolumesResponseBodyVolumes) SetVolumeInfo(v []*ListVolumesResponseBodyVolumesVolumeInfo) *ListVolumesResponseBodyVolumes {
	s.VolumeInfo = v
	return s
}

type ListVolumesResponseBodyVolumesVolumeInfo struct {
	AdditionalVolumes *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumes `json:"AdditionalVolumes,omitempty" xml:"AdditionalVolumes,omitempty" type:"Struct"`
	ClusterId         *string                                                    `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName       *string                                                    `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	RegionId          *string                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RemoteDirectory   *string                                                    `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	VolumeId          *string                                                    `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	VolumeMountpoint  *string                                                    `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	VolumeProtocol    *string                                                    `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	VolumeType        *string                                                    `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
}

func (s ListVolumesResponseBodyVolumesVolumeInfo) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponseBodyVolumesVolumeInfo) GoString() string {
	return s.String()
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetAdditionalVolumes(v *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumes) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.AdditionalVolumes = v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetClusterId(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.ClusterId = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetClusterName(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.ClusterName = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetRegionId(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.RegionId = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetRemoteDirectory(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.RemoteDirectory = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetVolumeId(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeId = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetVolumeMountpoint(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeMountpoint = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetVolumeProtocol(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeProtocol = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetVolumeType(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeType = &v
	return s
}

type ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumes struct {
	VolumeInfo []*ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo `json:"VolumeInfo,omitempty" xml:"VolumeInfo,omitempty" type:"Repeated"`
}

func (s ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumes) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumes) GoString() string {
	return s.String()
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumes) SetVolumeInfo(v []*ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumes {
	s.VolumeInfo = v
	return s
}

type ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo struct {
	JobQueue         *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	LocalDirectory   *string `json:"LocalDirectory,omitempty" xml:"LocalDirectory,omitempty"`
	Location         *string `json:"Location,omitempty" xml:"Location,omitempty"`
	RemoteDirectory  *string `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	Role             *string `json:"Role,omitempty" xml:"Role,omitempty"`
	VolumeId         *string `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	VolumeMountpoint *string `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	VolumeProtocol   *string `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	VolumeType       *string `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
}

func (s ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) GoString() string {
	return s.String()
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetJobQueue(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.JobQueue = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetLocalDirectory(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.LocalDirectory = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetLocation(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.Location = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetRemoteDirectory(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.RemoteDirectory = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetRole(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.Role = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetVolumeId(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.VolumeId = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetVolumeMountpoint(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.VolumeMountpoint = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetVolumeProtocol(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.VolumeProtocol = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo) SetVolumeType(v string) *ListVolumesResponseBodyVolumesVolumeInfoAdditionalVolumesVolumeInfo {
	s.VolumeType = &v
	return s
}

type ListVolumesResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVolumesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVolumesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponse) GoString() string {
	return s.String()
}

func (s *ListVolumesResponse) SetHeaders(v map[string]*string) *ListVolumesResponse {
	s.Headers = v
	return s
}

func (s *ListVolumesResponse) SetBody(v *ListVolumesResponseBody) *ListVolumesResponse {
	s.Body = v
	return s
}

type ModifyClusterAttributesRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyClusterAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterAttributesRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterAttributesRequest) SetClusterId(v string) *ModifyClusterAttributesRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterAttributesRequest) SetDescription(v string) *ModifyClusterAttributesRequest {
	s.Description = &v
	return s
}

func (s *ModifyClusterAttributesRequest) SetImageId(v string) *ModifyClusterAttributesRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyClusterAttributesRequest) SetImageOwnerAlias(v string) *ModifyClusterAttributesRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ModifyClusterAttributesRequest) SetName(v string) *ModifyClusterAttributesRequest {
	s.Name = &v
	return s
}

type ModifyClusterAttributesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterAttributesResponseBody) SetRequestId(v string) *ModifyClusterAttributesResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterAttributesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyClusterAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyClusterAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterAttributesResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterAttributesResponse) SetHeaders(v map[string]*string) *ModifyClusterAttributesResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterAttributesResponse) SetBody(v *ModifyClusterAttributesResponseBody) *ModifyClusterAttributesResponse {
	s.Body = v
	return s
}

type ModifyContainerAppAttributesRequest struct {
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyContainerAppAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyContainerAppAttributesRequest) GoString() string {
	return s.String()
}

func (s *ModifyContainerAppAttributesRequest) SetContainerId(v string) *ModifyContainerAppAttributesRequest {
	s.ContainerId = &v
	return s
}

func (s *ModifyContainerAppAttributesRequest) SetDescription(v string) *ModifyContainerAppAttributesRequest {
	s.Description = &v
	return s
}

type ModifyContainerAppAttributesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyContainerAppAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyContainerAppAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyContainerAppAttributesResponseBody) SetRequestId(v string) *ModifyContainerAppAttributesResponseBody {
	s.RequestId = &v
	return s
}

type ModifyContainerAppAttributesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyContainerAppAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyContainerAppAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyContainerAppAttributesResponse) GoString() string {
	return s.String()
}

func (s *ModifyContainerAppAttributesResponse) SetHeaders(v map[string]*string) *ModifyContainerAppAttributesResponse {
	s.Headers = v
	return s
}

func (s *ModifyContainerAppAttributesResponse) SetBody(v *ModifyContainerAppAttributesResponseBody) *ModifyContainerAppAttributesResponse {
	s.Body = v
	return s
}

type ModifyImageGatewayConfigRequest struct {
	ClusterId              *string                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DBPassword             *string                                `json:"DBPassword,omitempty" xml:"DBPassword,omitempty"`
	DBServerInfo           *string                                `json:"DBServerInfo,omitempty" xml:"DBServerInfo,omitempty"`
	DBType                 *string                                `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBUsername             *string                                `json:"DBUsername,omitempty" xml:"DBUsername,omitempty"`
	DefaultRepoLocation    *string                                `json:"DefaultRepoLocation,omitempty" xml:"DefaultRepoLocation,omitempty"`
	ImageExpirationTimeout *string                                `json:"ImageExpirationTimeout,omitempty" xml:"ImageExpirationTimeout,omitempty"`
	PullUpdateTimeout      *int32                                 `json:"PullUpdateTimeout,omitempty" xml:"PullUpdateTimeout,omitempty"`
	Repo                   []*ModifyImageGatewayConfigRequestRepo `json:"Repo,omitempty" xml:"Repo,omitempty" type:"Repeated"`
}

func (s ModifyImageGatewayConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageGatewayConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageGatewayConfigRequest) SetClusterId(v string) *ModifyImageGatewayConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetDBPassword(v string) *ModifyImageGatewayConfigRequest {
	s.DBPassword = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetDBServerInfo(v string) *ModifyImageGatewayConfigRequest {
	s.DBServerInfo = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetDBType(v string) *ModifyImageGatewayConfigRequest {
	s.DBType = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetDBUsername(v string) *ModifyImageGatewayConfigRequest {
	s.DBUsername = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetDefaultRepoLocation(v string) *ModifyImageGatewayConfigRequest {
	s.DefaultRepoLocation = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetImageExpirationTimeout(v string) *ModifyImageGatewayConfigRequest {
	s.ImageExpirationTimeout = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetPullUpdateTimeout(v int32) *ModifyImageGatewayConfigRequest {
	s.PullUpdateTimeout = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetRepo(v []*ModifyImageGatewayConfigRequestRepo) *ModifyImageGatewayConfigRequest {
	s.Repo = v
	return s
}

type ModifyImageGatewayConfigRequestRepo struct {
	Auth     *string `json:"Auth,omitempty" xml:"Auth,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	URL      *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s ModifyImageGatewayConfigRequestRepo) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageGatewayConfigRequestRepo) GoString() string {
	return s.String()
}

func (s *ModifyImageGatewayConfigRequestRepo) SetAuth(v string) *ModifyImageGatewayConfigRequestRepo {
	s.Auth = &v
	return s
}

func (s *ModifyImageGatewayConfigRequestRepo) SetLocation(v string) *ModifyImageGatewayConfigRequestRepo {
	s.Location = &v
	return s
}

func (s *ModifyImageGatewayConfigRequestRepo) SetURL(v string) *ModifyImageGatewayConfigRequestRepo {
	s.URL = &v
	return s
}

type ModifyImageGatewayConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyImageGatewayConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageGatewayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImageGatewayConfigResponseBody) SetRequestId(v string) *ModifyImageGatewayConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyImageGatewayConfigResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyImageGatewayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyImageGatewayConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageGatewayConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageGatewayConfigResponse) SetHeaders(v map[string]*string) *ModifyImageGatewayConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyImageGatewayConfigResponse) SetBody(v *ModifyImageGatewayConfigResponseBody) *ModifyImageGatewayConfigResponse {
	s.Body = v
	return s
}

type ModifyUserGroupsRequest struct {
	ClusterId *string                        `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	User      []*ModifyUserGroupsRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ModifyUserGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupsRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupsRequest) SetClusterId(v string) *ModifyUserGroupsRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyUserGroupsRequest) SetUser(v []*ModifyUserGroupsRequestUser) *ModifyUserGroupsRequest {
	s.User = v
	return s
}

type ModifyUserGroupsRequestUser struct {
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyUserGroupsRequestUser) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupsRequestUser) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupsRequestUser) SetGroup(v string) *ModifyUserGroupsRequestUser {
	s.Group = &v
	return s
}

func (s *ModifyUserGroupsRequestUser) SetName(v string) *ModifyUserGroupsRequestUser {
	s.Name = &v
	return s
}

type ModifyUserGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupsResponseBody) SetRequestId(v string) *ModifyUserGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserGroupsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyUserGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUserGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupsResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupsResponse) SetHeaders(v map[string]*string) *ModifyUserGroupsResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserGroupsResponse) SetBody(v *ModifyUserGroupsResponseBody) *ModifyUserGroupsResponse {
	s.Body = v
	return s
}

type ModifyUserPasswordsRequest struct {
	ClusterId *string                           `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	User      []*ModifyUserPasswordsRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ModifyUserPasswordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserPasswordsRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordsRequest) SetClusterId(v string) *ModifyUserPasswordsRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyUserPasswordsRequest) SetUser(v []*ModifyUserPasswordsRequestUser) *ModifyUserPasswordsRequest {
	s.User = v
	return s
}

type ModifyUserPasswordsRequestUser struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s ModifyUserPasswordsRequestUser) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserPasswordsRequestUser) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordsRequestUser) SetName(v string) *ModifyUserPasswordsRequestUser {
	s.Name = &v
	return s
}

func (s *ModifyUserPasswordsRequestUser) SetPassword(v string) *ModifyUserPasswordsRequestUser {
	s.Password = &v
	return s
}

type ModifyUserPasswordsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserPasswordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserPasswordsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordsResponseBody) SetRequestId(v string) *ModifyUserPasswordsResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserPasswordsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyUserPasswordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUserPasswordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserPasswordsResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordsResponse) SetHeaders(v map[string]*string) *ModifyUserPasswordsResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserPasswordsResponse) SetBody(v *ModifyUserPasswordsResponseBody) *ModifyUserPasswordsResponse {
	s.Body = v
	return s
}

type ModifyVisualServicePasswdRequest struct {
	ClusterId         *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Passwd            *string `json:"Passwd,omitempty" xml:"Passwd,omitempty"`
	RunasUser         *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	RunasUserPassword *string `json:"RunasUserPassword,omitempty" xml:"RunasUserPassword,omitempty"`
}

func (s ModifyVisualServicePasswdRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVisualServicePasswdRequest) GoString() string {
	return s.String()
}

func (s *ModifyVisualServicePasswdRequest) SetClusterId(v string) *ModifyVisualServicePasswdRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyVisualServicePasswdRequest) SetPasswd(v string) *ModifyVisualServicePasswdRequest {
	s.Passwd = &v
	return s
}

func (s *ModifyVisualServicePasswdRequest) SetRunasUser(v string) *ModifyVisualServicePasswdRequest {
	s.RunasUser = &v
	return s
}

func (s *ModifyVisualServicePasswdRequest) SetRunasUserPassword(v string) *ModifyVisualServicePasswdRequest {
	s.RunasUserPassword = &v
	return s
}

type ModifyVisualServicePasswdResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVisualServicePasswdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVisualServicePasswdResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVisualServicePasswdResponseBody) SetMessage(v string) *ModifyVisualServicePasswdResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyVisualServicePasswdResponseBody) SetRequestId(v string) *ModifyVisualServicePasswdResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVisualServicePasswdResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyVisualServicePasswdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyVisualServicePasswdResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVisualServicePasswdResponse) GoString() string {
	return s.String()
}

func (s *ModifyVisualServicePasswdResponse) SetHeaders(v map[string]*string) *ModifyVisualServicePasswdResponse {
	s.Headers = v
	return s
}

func (s *ModifyVisualServicePasswdResponse) SetBody(v *ModifyVisualServicePasswdResponseBody) *ModifyVisualServicePasswdResponse {
	s.Body = v
	return s
}

type MountNFSRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MountDir     *string `json:"MountDir,omitempty" xml:"MountDir,omitempty"`
	NfsDir       *string `json:"NfsDir,omitempty" xml:"NfsDir,omitempty"`
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RemoteDir    *string `json:"RemoteDir,omitempty" xml:"RemoteDir,omitempty"`
}

func (s MountNFSRequest) String() string {
	return tea.Prettify(s)
}

func (s MountNFSRequest) GoString() string {
	return s.String()
}

func (s *MountNFSRequest) SetInstanceId(v string) *MountNFSRequest {
	s.InstanceId = &v
	return s
}

func (s *MountNFSRequest) SetMountDir(v string) *MountNFSRequest {
	s.MountDir = &v
	return s
}

func (s *MountNFSRequest) SetNfsDir(v string) *MountNFSRequest {
	s.NfsDir = &v
	return s
}

func (s *MountNFSRequest) SetProtocolType(v string) *MountNFSRequest {
	s.ProtocolType = &v
	return s
}

func (s *MountNFSRequest) SetRemoteDir(v string) *MountNFSRequest {
	s.RemoteDir = &v
	return s
}

type MountNFSResponseBody struct {
	InvokeId  *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MountNFSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MountNFSResponseBody) GoString() string {
	return s.String()
}

func (s *MountNFSResponseBody) SetInvokeId(v string) *MountNFSResponseBody {
	s.InvokeId = &v
	return s
}

func (s *MountNFSResponseBody) SetRequestId(v string) *MountNFSResponseBody {
	s.RequestId = &v
	return s
}

type MountNFSResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MountNFSResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MountNFSResponse) String() string {
	return tea.Prettify(s)
}

func (s MountNFSResponse) GoString() string {
	return s.String()
}

func (s *MountNFSResponse) SetHeaders(v map[string]*string) *MountNFSResponse {
	s.Headers = v
	return s
}

func (s *MountNFSResponse) SetBody(v *MountNFSResponseBody) *MountNFSResponse {
	s.Body = v
	return s
}

type PullImageRequest struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
	ImageTag      *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	Repository    *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
}

func (s PullImageRequest) String() string {
	return tea.Prettify(s)
}

func (s PullImageRequest) GoString() string {
	return s.String()
}

func (s *PullImageRequest) SetClusterId(v string) *PullImageRequest {
	s.ClusterId = &v
	return s
}

func (s *PullImageRequest) SetContainerType(v string) *PullImageRequest {
	s.ContainerType = &v
	return s
}

func (s *PullImageRequest) SetImageTag(v string) *PullImageRequest {
	s.ImageTag = &v
	return s
}

func (s *PullImageRequest) SetRepository(v string) *PullImageRequest {
	s.Repository = &v
	return s
}

type PullImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PullImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PullImageResponseBody) GoString() string {
	return s.String()
}

func (s *PullImageResponseBody) SetRequestId(v string) *PullImageResponseBody {
	s.RequestId = &v
	return s
}

type PullImageResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PullImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PullImageResponse) String() string {
	return tea.Prettify(s)
}

func (s PullImageResponse) GoString() string {
	return s.String()
}

func (s *PullImageResponse) SetHeaders(v map[string]*string) *PullImageResponse {
	s.Headers = v
	return s
}

func (s *PullImageResponse) SetBody(v *PullImageResponseBody) *PullImageResponse {
	s.Body = v
	return s
}

type QueryServicePackAndPriceResponseBody struct {
	ChargeAmount   *int32                                           `json:"ChargeAmount,omitempty" xml:"ChargeAmount,omitempty"`
	Currency       *string                                          `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice  *float32                                         `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalAmount *int32                                           `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	OriginalPrice  *float32                                         `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	RegionId       *string                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServicePack    *QueryServicePackAndPriceResponseBodyServicePack `json:"ServicePack,omitempty" xml:"ServicePack,omitempty" type:"Struct"`
	TradePrice     *float32                                         `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s QueryServicePackAndPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryServicePackAndPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryServicePackAndPriceResponseBody) SetChargeAmount(v int32) *QueryServicePackAndPriceResponseBody {
	s.ChargeAmount = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetCurrency(v string) *QueryServicePackAndPriceResponseBody {
	s.Currency = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetDiscountPrice(v float32) *QueryServicePackAndPriceResponseBody {
	s.DiscountPrice = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetOriginalAmount(v int32) *QueryServicePackAndPriceResponseBody {
	s.OriginalAmount = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetOriginalPrice(v float32) *QueryServicePackAndPriceResponseBody {
	s.OriginalPrice = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetRegionId(v string) *QueryServicePackAndPriceResponseBody {
	s.RegionId = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetRequestId(v string) *QueryServicePackAndPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetServicePack(v *QueryServicePackAndPriceResponseBodyServicePack) *QueryServicePackAndPriceResponseBody {
	s.ServicePack = v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetTradePrice(v float32) *QueryServicePackAndPriceResponseBody {
	s.TradePrice = &v
	return s
}

type QueryServicePackAndPriceResponseBodyServicePack struct {
	ServicePackInfo []*QueryServicePackAndPriceResponseBodyServicePackServicePackInfo `json:"ServicePackInfo,omitempty" xml:"ServicePackInfo,omitempty" type:"Repeated"`
}

func (s QueryServicePackAndPriceResponseBodyServicePack) String() string {
	return tea.Prettify(s)
}

func (s QueryServicePackAndPriceResponseBodyServicePack) GoString() string {
	return s.String()
}

func (s *QueryServicePackAndPriceResponseBodyServicePack) SetServicePackInfo(v []*QueryServicePackAndPriceResponseBodyServicePackServicePackInfo) *QueryServicePackAndPriceResponseBodyServicePack {
	s.ServicePackInfo = v
	return s
}

type QueryServicePackAndPriceResponseBodyServicePackServicePackInfo struct {
	Capacity     *int32  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	EndTime      *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	StartTime    *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryServicePackAndPriceResponseBodyServicePackServicePackInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryServicePackAndPriceResponseBodyServicePackServicePackInfo) GoString() string {
	return s.String()
}

func (s *QueryServicePackAndPriceResponseBodyServicePackServicePackInfo) SetCapacity(v int32) *QueryServicePackAndPriceResponseBodyServicePackServicePackInfo {
	s.Capacity = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBodyServicePackServicePackInfo) SetEndTime(v int32) *QueryServicePackAndPriceResponseBodyServicePackServicePackInfo {
	s.EndTime = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBodyServicePackServicePackInfo) SetInstanceName(v string) *QueryServicePackAndPriceResponseBodyServicePackServicePackInfo {
	s.InstanceName = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBodyServicePackServicePackInfo) SetStartTime(v int32) *QueryServicePackAndPriceResponseBodyServicePackServicePackInfo {
	s.StartTime = &v
	return s
}

type QueryServicePackAndPriceResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryServicePackAndPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryServicePackAndPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryServicePackAndPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryServicePackAndPriceResponse) SetHeaders(v map[string]*string) *QueryServicePackAndPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryServicePackAndPriceResponse) SetBody(v *QueryServicePackAndPriceResponseBody) *QueryServicePackAndPriceResponse {
	s.Body = v
	return s
}

type RecoverClusterRequest struct {
	AccountType     *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	ClientVersion   *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	OsTag           *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	SchedulerType   *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
}

func (s RecoverClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s RecoverClusterRequest) GoString() string {
	return s.String()
}

func (s *RecoverClusterRequest) SetAccountType(v string) *RecoverClusterRequest {
	s.AccountType = &v
	return s
}

func (s *RecoverClusterRequest) SetClientVersion(v string) *RecoverClusterRequest {
	s.ClientVersion = &v
	return s
}

func (s *RecoverClusterRequest) SetClusterId(v string) *RecoverClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *RecoverClusterRequest) SetImageId(v string) *RecoverClusterRequest {
	s.ImageId = &v
	return s
}

func (s *RecoverClusterRequest) SetImageOwnerAlias(v string) *RecoverClusterRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *RecoverClusterRequest) SetOsTag(v string) *RecoverClusterRequest {
	s.OsTag = &v
	return s
}

func (s *RecoverClusterRequest) SetSchedulerType(v string) *RecoverClusterRequest {
	s.SchedulerType = &v
	return s
}

type RecoverClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RecoverClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecoverClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverClusterResponseBody) SetRequestId(v string) *RecoverClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoverClusterResponseBody) SetTaskId(v string) *RecoverClusterResponseBody {
	s.TaskId = &v
	return s
}

type RecoverClusterResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecoverClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecoverClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s RecoverClusterResponse) GoString() string {
	return s.String()
}

func (s *RecoverClusterResponse) SetHeaders(v map[string]*string) *RecoverClusterResponse {
	s.Headers = v
	return s
}

func (s *RecoverClusterResponse) SetBody(v *RecoverClusterResponseBody) *RecoverClusterResponse {
	s.Body = v
	return s
}

type RerunJobsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Jobs      *string `json:"Jobs,omitempty" xml:"Jobs,omitempty"`
}

func (s RerunJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s RerunJobsRequest) GoString() string {
	return s.String()
}

func (s *RerunJobsRequest) SetClusterId(v string) *RerunJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *RerunJobsRequest) SetJobs(v string) *RerunJobsRequest {
	s.Jobs = &v
	return s
}

type RerunJobsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RerunJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RerunJobsResponseBody) GoString() string {
	return s.String()
}

func (s *RerunJobsResponseBody) SetRequestId(v string) *RerunJobsResponseBody {
	s.RequestId = &v
	return s
}

type RerunJobsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RerunJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RerunJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s RerunJobsResponse) GoString() string {
	return s.String()
}

func (s *RerunJobsResponse) SetHeaders(v map[string]*string) *RerunJobsResponse {
	s.Headers = v
	return s
}

func (s *RerunJobsResponse) SetBody(v *RerunJobsResponseBody) *RerunJobsResponse {
	s.Body = v
	return s
}

type ResetNodesRequest struct {
	ClusterId *string                      `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Instance  []*ResetNodesRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s ResetNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetNodesRequest) GoString() string {
	return s.String()
}

func (s *ResetNodesRequest) SetClusterId(v string) *ResetNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ResetNodesRequest) SetInstance(v []*ResetNodesRequestInstance) *ResetNodesRequest {
	s.Instance = v
	return s
}

type ResetNodesRequestInstance struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ResetNodesRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s ResetNodesRequestInstance) GoString() string {
	return s.String()
}

func (s *ResetNodesRequestInstance) SetId(v string) *ResetNodesRequestInstance {
	s.Id = &v
	return s
}

type ResetNodesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ResetNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ResetNodesResponseBody) SetRequestId(v string) *ResetNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetNodesResponseBody) SetTaskId(v string) *ResetNodesResponseBody {
	s.TaskId = &v
	return s
}

type ResetNodesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetNodesResponse) GoString() string {
	return s.String()
}

func (s *ResetNodesResponse) SetHeaders(v map[string]*string) *ResetNodesResponse {
	s.Headers = v
	return s
}

func (s *ResetNodesResponse) SetBody(v *ResetNodesResponseBody) *ResetNodesResponse {
	s.Body = v
	return s
}

type RunCloudMetricProfilingRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Duration  *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Freq      *int32  `json:"Freq,omitempty" xml:"Freq,omitempty"`
	HostName  *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ProcessId *int32  `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RunCloudMetricProfilingRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCloudMetricProfilingRequest) GoString() string {
	return s.String()
}

func (s *RunCloudMetricProfilingRequest) SetClusterId(v string) *RunCloudMetricProfilingRequest {
	s.ClusterId = &v
	return s
}

func (s *RunCloudMetricProfilingRequest) SetDuration(v int32) *RunCloudMetricProfilingRequest {
	s.Duration = &v
	return s
}

func (s *RunCloudMetricProfilingRequest) SetFreq(v int32) *RunCloudMetricProfilingRequest {
	s.Freq = &v
	return s
}

func (s *RunCloudMetricProfilingRequest) SetHostName(v string) *RunCloudMetricProfilingRequest {
	s.HostName = &v
	return s
}

func (s *RunCloudMetricProfilingRequest) SetProcessId(v int32) *RunCloudMetricProfilingRequest {
	s.ProcessId = &v
	return s
}

func (s *RunCloudMetricProfilingRequest) SetRegionId(v string) *RunCloudMetricProfilingRequest {
	s.RegionId = &v
	return s
}

type RunCloudMetricProfilingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunCloudMetricProfilingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCloudMetricProfilingResponseBody) GoString() string {
	return s.String()
}

func (s *RunCloudMetricProfilingResponseBody) SetRequestId(v string) *RunCloudMetricProfilingResponseBody {
	s.RequestId = &v
	return s
}

type RunCloudMetricProfilingResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunCloudMetricProfilingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunCloudMetricProfilingResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCloudMetricProfilingResponse) GoString() string {
	return s.String()
}

func (s *RunCloudMetricProfilingResponse) SetHeaders(v map[string]*string) *RunCloudMetricProfilingResponse {
	s.Headers = v
	return s
}

func (s *RunCloudMetricProfilingResponse) SetBody(v *RunCloudMetricProfilingResponseBody) *RunCloudMetricProfilingResponse {
	s.Body = v
	return s
}

type SetAutoScaleConfigRequest struct {
	ClusterId               *string                            `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	EnableAutoGrow          *bool                              `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	EnableAutoShrink        *bool                              `json:"EnableAutoShrink,omitempty" xml:"EnableAutoShrink,omitempty"`
	ExcludeNodes            *string                            `json:"ExcludeNodes,omitempty" xml:"ExcludeNodes,omitempty"`
	ExtraNodesGrowRatio     *int32                             `json:"ExtraNodesGrowRatio,omitempty" xml:"ExtraNodesGrowRatio,omitempty"`
	GrowIntervalInMinutes   *int32                             `json:"GrowIntervalInMinutes,omitempty" xml:"GrowIntervalInMinutes,omitempty"`
	GrowRatio               *int32                             `json:"GrowRatio,omitempty" xml:"GrowRatio,omitempty"`
	GrowTimeoutInMinutes    *int32                             `json:"GrowTimeoutInMinutes,omitempty" xml:"GrowTimeoutInMinutes,omitempty"`
	ImageId                 *string                            `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	MaxNodesInCluster       *int32                             `json:"MaxNodesInCluster,omitempty" xml:"MaxNodesInCluster,omitempty"`
	Queues                  []*SetAutoScaleConfigRequestQueues `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
	ShrinkIdleTimes         *int32                             `json:"ShrinkIdleTimes,omitempty" xml:"ShrinkIdleTimes,omitempty"`
	ShrinkIntervalInMinutes *int32                             `json:"ShrinkIntervalInMinutes,omitempty" xml:"ShrinkIntervalInMinutes,omitempty"`
	SpotPriceLimit          *float32                           `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy            *string                            `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s SetAutoScaleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigRequest) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigRequest) SetClusterId(v string) *SetAutoScaleConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetEnableAutoGrow(v bool) *SetAutoScaleConfigRequest {
	s.EnableAutoGrow = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetEnableAutoShrink(v bool) *SetAutoScaleConfigRequest {
	s.EnableAutoShrink = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetExcludeNodes(v string) *SetAutoScaleConfigRequest {
	s.ExcludeNodes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetExtraNodesGrowRatio(v int32) *SetAutoScaleConfigRequest {
	s.ExtraNodesGrowRatio = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetGrowIntervalInMinutes(v int32) *SetAutoScaleConfigRequest {
	s.GrowIntervalInMinutes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetGrowRatio(v int32) *SetAutoScaleConfigRequest {
	s.GrowRatio = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetGrowTimeoutInMinutes(v int32) *SetAutoScaleConfigRequest {
	s.GrowTimeoutInMinutes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetImageId(v string) *SetAutoScaleConfigRequest {
	s.ImageId = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetMaxNodesInCluster(v int32) *SetAutoScaleConfigRequest {
	s.MaxNodesInCluster = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetQueues(v []*SetAutoScaleConfigRequestQueues) *SetAutoScaleConfigRequest {
	s.Queues = v
	return s
}

func (s *SetAutoScaleConfigRequest) SetShrinkIdleTimes(v int32) *SetAutoScaleConfigRequest {
	s.ShrinkIdleTimes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetShrinkIntervalInMinutes(v int32) *SetAutoScaleConfigRequest {
	s.ShrinkIntervalInMinutes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetSpotPriceLimit(v float32) *SetAutoScaleConfigRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetSpotStrategy(v string) *SetAutoScaleConfigRequest {
	s.SpotStrategy = &v
	return s
}

type SetAutoScaleConfigRequestQueues struct {
	DataDisks          []*SetAutoScaleConfigRequestQueuesDataDisks     `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	EnableAutoGrow     *bool                                           `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	EnableAutoShrink   *bool                                           `json:"EnableAutoShrink,omitempty" xml:"EnableAutoShrink,omitempty"`
	HostNamePrefix     *string                                         `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	HostNameSuffix     *string                                         `json:"HostNameSuffix,omitempty" xml:"HostNameSuffix,omitempty"`
	InstanceType       *string                                         `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceTypes      []*SetAutoScaleConfigRequestQueuesInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	MaxNodesInQueue    *int32                                          `json:"MaxNodesInQueue,omitempty" xml:"MaxNodesInQueue,omitempty"`
	MinNodesInQueue    *int32                                          `json:"MinNodesInQueue,omitempty" xml:"MinNodesInQueue,omitempty"`
	QueueImageId       *string                                         `json:"QueueImageId,omitempty" xml:"QueueImageId,omitempty"`
	QueueName          *string                                         `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	SpotPriceLimit     *float32                                        `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy       *string                                         `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDiskCategory *string                                         `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	SystemDiskLevel    *string                                         `json:"SystemDiskLevel,omitempty" xml:"SystemDiskLevel,omitempty"`
	SystemDiskSize     *int32                                          `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s SetAutoScaleConfigRequestQueues) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigRequestQueues) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigRequestQueues) SetDataDisks(v []*SetAutoScaleConfigRequestQueuesDataDisks) *SetAutoScaleConfigRequestQueues {
	s.DataDisks = v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetEnableAutoGrow(v bool) *SetAutoScaleConfigRequestQueues {
	s.EnableAutoGrow = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetEnableAutoShrink(v bool) *SetAutoScaleConfigRequestQueues {
	s.EnableAutoShrink = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetHostNamePrefix(v string) *SetAutoScaleConfigRequestQueues {
	s.HostNamePrefix = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetHostNameSuffix(v string) *SetAutoScaleConfigRequestQueues {
	s.HostNameSuffix = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetInstanceType(v string) *SetAutoScaleConfigRequestQueues {
	s.InstanceType = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetInstanceTypes(v []*SetAutoScaleConfigRequestQueuesInstanceTypes) *SetAutoScaleConfigRequestQueues {
	s.InstanceTypes = v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetMaxNodesInQueue(v int32) *SetAutoScaleConfigRequestQueues {
	s.MaxNodesInQueue = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetMinNodesInQueue(v int32) *SetAutoScaleConfigRequestQueues {
	s.MinNodesInQueue = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetQueueImageId(v string) *SetAutoScaleConfigRequestQueues {
	s.QueueImageId = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetQueueName(v string) *SetAutoScaleConfigRequestQueues {
	s.QueueName = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetSpotPriceLimit(v float32) *SetAutoScaleConfigRequestQueues {
	s.SpotPriceLimit = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetSpotStrategy(v string) *SetAutoScaleConfigRequestQueues {
	s.SpotStrategy = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetSystemDiskCategory(v string) *SetAutoScaleConfigRequestQueues {
	s.SystemDiskCategory = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetSystemDiskLevel(v string) *SetAutoScaleConfigRequestQueues {
	s.SystemDiskLevel = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetSystemDiskSize(v int32) *SetAutoScaleConfigRequestQueues {
	s.SystemDiskSize = &v
	return s
}

type SetAutoScaleConfigRequestQueuesDataDisks struct {
	DataDiskCategory           *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	DataDiskDeleteWithInstance *bool   `json:"DataDiskDeleteWithInstance,omitempty" xml:"DataDiskDeleteWithInstance,omitempty"`
	DataDiskEncrypted          *bool   `json:"DataDiskEncrypted,omitempty" xml:"DataDiskEncrypted,omitempty"`
	DataDiskKMSKeyId           *string `json:"DataDiskKMSKeyId,omitempty" xml:"DataDiskKMSKeyId,omitempty"`
	DataDiskPerformanceLevel   *string `json:"DataDiskPerformanceLevel,omitempty" xml:"DataDiskPerformanceLevel,omitempty"`
	DataDiskSize               *int32  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
}

func (s SetAutoScaleConfigRequestQueuesDataDisks) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigRequestQueuesDataDisks) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigRequestQueuesDataDisks) SetDataDiskCategory(v string) *SetAutoScaleConfigRequestQueuesDataDisks {
	s.DataDiskCategory = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesDataDisks) SetDataDiskDeleteWithInstance(v bool) *SetAutoScaleConfigRequestQueuesDataDisks {
	s.DataDiskDeleteWithInstance = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesDataDisks) SetDataDiskEncrypted(v bool) *SetAutoScaleConfigRequestQueuesDataDisks {
	s.DataDiskEncrypted = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesDataDisks) SetDataDiskKMSKeyId(v string) *SetAutoScaleConfigRequestQueuesDataDisks {
	s.DataDiskKMSKeyId = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesDataDisks) SetDataDiskPerformanceLevel(v string) *SetAutoScaleConfigRequestQueuesDataDisks {
	s.DataDiskPerformanceLevel = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesDataDisks) SetDataDiskSize(v int32) *SetAutoScaleConfigRequestQueuesDataDisks {
	s.DataDiskSize = &v
	return s
}

type SetAutoScaleConfigRequestQueuesInstanceTypes struct {
	InstanceType   *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy   *string  `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	VSwitchId      *string  `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId         *string  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SetAutoScaleConfigRequestQueuesInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigRequestQueuesInstanceTypes) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetInstanceType(v string) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.InstanceType = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetSpotPriceLimit(v float32) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.SpotPriceLimit = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetSpotStrategy(v string) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.SpotStrategy = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetVSwitchId(v string) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.VSwitchId = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetZoneId(v string) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.ZoneId = &v
	return s
}

type SetAutoScaleConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAutoScaleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigResponseBody) SetRequestId(v string) *SetAutoScaleConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetAutoScaleConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetAutoScaleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAutoScaleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigResponse) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigResponse) SetHeaders(v map[string]*string) *SetAutoScaleConfigResponse {
	s.Headers = v
	return s
}

func (s *SetAutoScaleConfigResponse) SetBody(v *SetAutoScaleConfigResponseBody) *SetAutoScaleConfigResponse {
	s.Body = v
	return s
}

type SetGWSClusterPolicyRequest struct {
	AsyncMode   *bool   `json:"AsyncMode,omitempty" xml:"AsyncMode,omitempty"`
	Clipboard   *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	LocalDrive  *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	UdpPort     *string `json:"UdpPort,omitempty" xml:"UdpPort,omitempty"`
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	Watermark   *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s SetGWSClusterPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGWSClusterPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetGWSClusterPolicyRequest) SetAsyncMode(v bool) *SetGWSClusterPolicyRequest {
	s.AsyncMode = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetClipboard(v string) *SetGWSClusterPolicyRequest {
	s.Clipboard = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetClusterId(v string) *SetGWSClusterPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetLocalDrive(v string) *SetGWSClusterPolicyRequest {
	s.LocalDrive = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetUdpPort(v string) *SetGWSClusterPolicyRequest {
	s.UdpPort = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetUsbRedirect(v string) *SetGWSClusterPolicyRequest {
	s.UsbRedirect = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetWatermark(v string) *SetGWSClusterPolicyRequest {
	s.Watermark = &v
	return s
}

type SetGWSClusterPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetGWSClusterPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGWSClusterPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetGWSClusterPolicyResponseBody) SetRequestId(v string) *SetGWSClusterPolicyResponseBody {
	s.RequestId = &v
	return s
}

type SetGWSClusterPolicyResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetGWSClusterPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGWSClusterPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGWSClusterPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetGWSClusterPolicyResponse) SetHeaders(v map[string]*string) *SetGWSClusterPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetGWSClusterPolicyResponse) SetBody(v *SetGWSClusterPolicyResponseBody) *SetGWSClusterPolicyResponse {
	s.Body = v
	return s
}

type SetGWSInstanceNameRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SetGWSInstanceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceNameRequest) SetInstanceId(v string) *SetGWSInstanceNameRequest {
	s.InstanceId = &v
	return s
}

func (s *SetGWSInstanceNameRequest) SetName(v string) *SetGWSInstanceNameRequest {
	s.Name = &v
	return s
}

type SetGWSInstanceNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetGWSInstanceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceNameResponseBody) SetRequestId(v string) *SetGWSInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

type SetGWSInstanceNameResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetGWSInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGWSInstanceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceNameResponse) SetHeaders(v map[string]*string) *SetGWSInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *SetGWSInstanceNameResponse) SetBody(v *SetGWSInstanceNameResponseBody) *SetGWSInstanceNameResponse {
	s.Body = v
	return s
}

type SetGWSInstanceUserRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserUid    *string `json:"UserUid,omitempty" xml:"UserUid,omitempty"`
}

func (s SetGWSInstanceUserRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceUserRequest) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceUserRequest) SetInstanceId(v string) *SetGWSInstanceUserRequest {
	s.InstanceId = &v
	return s
}

func (s *SetGWSInstanceUserRequest) SetUserName(v string) *SetGWSInstanceUserRequest {
	s.UserName = &v
	return s
}

func (s *SetGWSInstanceUserRequest) SetUserUid(v string) *SetGWSInstanceUserRequest {
	s.UserUid = &v
	return s
}

type SetGWSInstanceUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetGWSInstanceUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceUserResponseBody) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceUserResponseBody) SetRequestId(v string) *SetGWSInstanceUserResponseBody {
	s.RequestId = &v
	return s
}

type SetGWSInstanceUserResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetGWSInstanceUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGWSInstanceUserResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceUserResponse) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceUserResponse) SetHeaders(v map[string]*string) *SetGWSInstanceUserResponse {
	s.Headers = v
	return s
}

func (s *SetGWSInstanceUserResponse) SetBody(v *SetGWSInstanceUserResponseBody) *SetGWSInstanceUserResponse {
	s.Body = v
	return s
}

type SetQueueRequest struct {
	ClusterId *string                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Node      []*SetQueueRequestNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Repeated"`
	QueueName *string                `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s SetQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s SetQueueRequest) GoString() string {
	return s.String()
}

func (s *SetQueueRequest) SetClusterId(v string) *SetQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *SetQueueRequest) SetNode(v []*SetQueueRequestNode) *SetQueueRequest {
	s.Node = v
	return s
}

func (s *SetQueueRequest) SetQueueName(v string) *SetQueueRequest {
	s.QueueName = &v
	return s
}

type SetQueueRequestNode struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SetQueueRequestNode) String() string {
	return tea.Prettify(s)
}

func (s SetQueueRequestNode) GoString() string {
	return s.String()
}

func (s *SetQueueRequestNode) SetName(v string) *SetQueueRequestNode {
	s.Name = &v
	return s
}

type SetQueueResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetQueueResponseBody) GoString() string {
	return s.String()
}

func (s *SetQueueResponseBody) SetRequestId(v string) *SetQueueResponseBody {
	s.RequestId = &v
	return s
}

type SetQueueResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s SetQueueResponse) GoString() string {
	return s.String()
}

func (s *SetQueueResponse) SetHeaders(v map[string]*string) *SetQueueResponse {
	s.Headers = v
	return s
}

func (s *SetQueueResponse) SetBody(v *SetQueueResponseBody) *SetQueueResponse {
	s.Body = v
	return s
}

type SetSchedulerInfoRequest struct {
	ClusterId *string                             `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PbsInfo   []*SetSchedulerInfoRequestPbsInfo   `json:"PbsInfo,omitempty" xml:"PbsInfo,omitempty" type:"Repeated"`
	RegionId  *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Scheduler []*SetSchedulerInfoRequestScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Repeated"`
	SlurmInfo []*SetSchedulerInfoRequestSlurmInfo `json:"SlurmInfo,omitempty" xml:"SlurmInfo,omitempty" type:"Repeated"`
}

func (s SetSchedulerInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSchedulerInfoRequest) GoString() string {
	return s.String()
}

func (s *SetSchedulerInfoRequest) SetClusterId(v string) *SetSchedulerInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *SetSchedulerInfoRequest) SetPbsInfo(v []*SetSchedulerInfoRequestPbsInfo) *SetSchedulerInfoRequest {
	s.PbsInfo = v
	return s
}

func (s *SetSchedulerInfoRequest) SetRegionId(v string) *SetSchedulerInfoRequest {
	s.RegionId = &v
	return s
}

func (s *SetSchedulerInfoRequest) SetScheduler(v []*SetSchedulerInfoRequestScheduler) *SetSchedulerInfoRequest {
	s.Scheduler = v
	return s
}

func (s *SetSchedulerInfoRequest) SetSlurmInfo(v []*SetSchedulerInfoRequestSlurmInfo) *SetSchedulerInfoRequest {
	s.SlurmInfo = v
	return s
}

type SetSchedulerInfoRequestPbsInfo struct {
	AclLimit           []*SetSchedulerInfoRequestPbsInfoAclLimit      `json:"AclLimit,omitempty" xml:"AclLimit,omitempty" type:"Repeated"`
	JobHistoryDuration *int32                                         `json:"JobHistoryDuration,omitempty" xml:"JobHistoryDuration,omitempty"`
	ResourceLimit      []*SetSchedulerInfoRequestPbsInfoResourceLimit `json:"ResourceLimit,omitempty" xml:"ResourceLimit,omitempty" type:"Repeated"`
	SchedInterval      *int32                                         `json:"SchedInterval,omitempty" xml:"SchedInterval,omitempty"`
}

func (s SetSchedulerInfoRequestPbsInfo) String() string {
	return tea.Prettify(s)
}

func (s SetSchedulerInfoRequestPbsInfo) GoString() string {
	return s.String()
}

func (s *SetSchedulerInfoRequestPbsInfo) SetAclLimit(v []*SetSchedulerInfoRequestPbsInfoAclLimit) *SetSchedulerInfoRequestPbsInfo {
	s.AclLimit = v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfo) SetJobHistoryDuration(v int32) *SetSchedulerInfoRequestPbsInfo {
	s.JobHistoryDuration = &v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfo) SetResourceLimit(v []*SetSchedulerInfoRequestPbsInfoResourceLimit) *SetSchedulerInfoRequestPbsInfo {
	s.ResourceLimit = v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfo) SetSchedInterval(v int32) *SetSchedulerInfoRequestPbsInfo {
	s.SchedInterval = &v
	return s
}

type SetSchedulerInfoRequestPbsInfoAclLimit struct {
	AclUsers *string `json:"AclUsers,omitempty" xml:"AclUsers,omitempty"`
	Queue    *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
}

func (s SetSchedulerInfoRequestPbsInfoAclLimit) String() string {
	return tea.Prettify(s)
}

func (s SetSchedulerInfoRequestPbsInfoAclLimit) GoString() string {
	return s.String()
}

func (s *SetSchedulerInfoRequestPbsInfoAclLimit) SetAclUsers(v string) *SetSchedulerInfoRequestPbsInfoAclLimit {
	s.AclUsers = &v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfoAclLimit) SetQueue(v string) *SetSchedulerInfoRequestPbsInfoAclLimit {
	s.Queue = &v
	return s
}

type SetSchedulerInfoRequestPbsInfoResourceLimit struct {
	Cpus  *int32  `json:"Cpus,omitempty" xml:"Cpus,omitempty"`
	Mem   *string `json:"Mem,omitempty" xml:"Mem,omitempty"`
	Nodes *int32  `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	Queue *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	User  *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s SetSchedulerInfoRequestPbsInfoResourceLimit) String() string {
	return tea.Prettify(s)
}

func (s SetSchedulerInfoRequestPbsInfoResourceLimit) GoString() string {
	return s.String()
}

func (s *SetSchedulerInfoRequestPbsInfoResourceLimit) SetCpus(v int32) *SetSchedulerInfoRequestPbsInfoResourceLimit {
	s.Cpus = &v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfoResourceLimit) SetMem(v string) *SetSchedulerInfoRequestPbsInfoResourceLimit {
	s.Mem = &v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfoResourceLimit) SetNodes(v int32) *SetSchedulerInfoRequestPbsInfoResourceLimit {
	s.Nodes = &v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfoResourceLimit) SetQueue(v string) *SetSchedulerInfoRequestPbsInfoResourceLimit {
	s.Queue = &v
	return s
}

func (s *SetSchedulerInfoRequestPbsInfoResourceLimit) SetUser(v string) *SetSchedulerInfoRequestPbsInfoResourceLimit {
	s.User = &v
	return s
}

type SetSchedulerInfoRequestScheduler struct {
	SchedName *string `json:"SchedName,omitempty" xml:"SchedName,omitempty"`
}

func (s SetSchedulerInfoRequestScheduler) String() string {
	return tea.Prettify(s)
}

func (s SetSchedulerInfoRequestScheduler) GoString() string {
	return s.String()
}

func (s *SetSchedulerInfoRequestScheduler) SetSchedName(v string) *SetSchedulerInfoRequestScheduler {
	s.SchedName = &v
	return s
}

type SetSchedulerInfoRequestSlurmInfo struct {
	BackfillInterval *int32 `json:"BackfillInterval,omitempty" xml:"BackfillInterval,omitempty"`
	SchedInterval    *int32 `json:"SchedInterval,omitempty" xml:"SchedInterval,omitempty"`
}

func (s SetSchedulerInfoRequestSlurmInfo) String() string {
	return tea.Prettify(s)
}

func (s SetSchedulerInfoRequestSlurmInfo) GoString() string {
	return s.String()
}

func (s *SetSchedulerInfoRequestSlurmInfo) SetBackfillInterval(v int32) *SetSchedulerInfoRequestSlurmInfo {
	s.BackfillInterval = &v
	return s
}

func (s *SetSchedulerInfoRequestSlurmInfo) SetSchedInterval(v int32) *SetSchedulerInfoRequestSlurmInfo {
	s.SchedInterval = &v
	return s
}

type SetSchedulerInfoResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetSchedulerInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetSchedulerInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SetSchedulerInfoResponseBody) SetMessage(v string) *SetSchedulerInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SetSchedulerInfoResponseBody) SetRequestId(v string) *SetSchedulerInfoResponseBody {
	s.RequestId = &v
	return s
}

type SetSchedulerInfoResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetSchedulerInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetSchedulerInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SetSchedulerInfoResponse) GoString() string {
	return s.String()
}

func (s *SetSchedulerInfoResponse) SetHeaders(v map[string]*string) *SetSchedulerInfoResponse {
	s.Headers = v
	return s
}

func (s *SetSchedulerInfoResponse) SetBody(v *SetSchedulerInfoResponseBody) *SetSchedulerInfoResponse {
	s.Body = v
	return s
}

type StartClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s StartClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s StartClusterRequest) GoString() string {
	return s.String()
}

func (s *StartClusterRequest) SetClusterId(v string) *StartClusterRequest {
	s.ClusterId = &v
	return s
}

type StartClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StartClusterResponseBody) SetRequestId(v string) *StartClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartClusterResponseBody) SetTaskId(v string) *StartClusterResponseBody {
	s.TaskId = &v
	return s
}

type StartClusterResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s StartClusterResponse) GoString() string {
	return s.String()
}

func (s *StartClusterResponse) SetHeaders(v map[string]*string) *StartClusterResponse {
	s.Headers = v
	return s
}

func (s *StartClusterResponse) SetBody(v *StartClusterResponseBody) *StartClusterResponse {
	s.Body = v
	return s
}

type StartGWSInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartGWSInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartGWSInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartGWSInstanceRequest) SetInstanceId(v string) *StartGWSInstanceRequest {
	s.InstanceId = &v
	return s
}

type StartGWSInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartGWSInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartGWSInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartGWSInstanceResponseBody) SetRequestId(v string) *StartGWSInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartGWSInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartGWSInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartGWSInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartGWSInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartGWSInstanceResponse) SetHeaders(v map[string]*string) *StartGWSInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartGWSInstanceResponse) SetBody(v *StartGWSInstanceResponseBody) *StartGWSInstanceResponse {
	s.Body = v
	return s
}

type StartNodesRequest struct {
	ClusterId *string                      `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Instance  []*StartNodesRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
	Role      *string                      `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s StartNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s StartNodesRequest) GoString() string {
	return s.String()
}

func (s *StartNodesRequest) SetClusterId(v string) *StartNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *StartNodesRequest) SetInstance(v []*StartNodesRequestInstance) *StartNodesRequest {
	s.Instance = v
	return s
}

func (s *StartNodesRequest) SetRole(v string) *StartNodesRequest {
	s.Role = &v
	return s
}

type StartNodesRequestInstance struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s StartNodesRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s StartNodesRequestInstance) GoString() string {
	return s.String()
}

func (s *StartNodesRequestInstance) SetId(v string) *StartNodesRequestInstance {
	s.Id = &v
	return s
}

type StartNodesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartNodesResponseBody) GoString() string {
	return s.String()
}

func (s *StartNodesResponseBody) SetRequestId(v string) *StartNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartNodesResponseBody) SetTaskId(v string) *StartNodesResponseBody {
	s.TaskId = &v
	return s
}

type StartNodesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s StartNodesResponse) GoString() string {
	return s.String()
}

func (s *StartNodesResponse) SetHeaders(v map[string]*string) *StartNodesResponse {
	s.Headers = v
	return s
}

func (s *StartNodesResponse) SetBody(v *StartNodesResponseBody) *StartNodesResponse {
	s.Body = v
	return s
}

type StartVisualServiceRequest struct {
	CidrIp    *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Port      *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s StartVisualServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartVisualServiceRequest) GoString() string {
	return s.String()
}

func (s *StartVisualServiceRequest) SetCidrIp(v string) *StartVisualServiceRequest {
	s.CidrIp = &v
	return s
}

func (s *StartVisualServiceRequest) SetClusterId(v string) *StartVisualServiceRequest {
	s.ClusterId = &v
	return s
}

func (s *StartVisualServiceRequest) SetPort(v int32) *StartVisualServiceRequest {
	s.Port = &v
	return s
}

type StartVisualServiceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartVisualServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartVisualServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StartVisualServiceResponseBody) SetMessage(v string) *StartVisualServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StartVisualServiceResponseBody) SetRequestId(v string) *StartVisualServiceResponseBody {
	s.RequestId = &v
	return s
}

type StartVisualServiceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartVisualServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartVisualServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartVisualServiceResponse) GoString() string {
	return s.String()
}

func (s *StartVisualServiceResponse) SetHeaders(v map[string]*string) *StartVisualServiceResponse {
	s.Headers = v
	return s
}

func (s *StartVisualServiceResponse) SetBody(v *StartVisualServiceResponseBody) *StartVisualServiceResponse {
	s.Body = v
	return s
}

type StopClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s StopClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s StopClusterRequest) GoString() string {
	return s.String()
}

func (s *StopClusterRequest) SetClusterId(v string) *StopClusterRequest {
	s.ClusterId = &v
	return s
}

type StopClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StopClusterResponseBody) SetRequestId(v string) *StopClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopClusterResponseBody) SetTaskId(v string) *StopClusterResponseBody {
	s.TaskId = &v
	return s
}

type StopClusterResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s StopClusterResponse) GoString() string {
	return s.String()
}

func (s *StopClusterResponse) SetHeaders(v map[string]*string) *StopClusterResponse {
	s.Headers = v
	return s
}

func (s *StopClusterResponse) SetBody(v *StopClusterResponseBody) *StopClusterResponse {
	s.Body = v
	return s
}

type StopGWSInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopGWSInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopGWSInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopGWSInstanceRequest) SetInstanceId(v string) *StopGWSInstanceRequest {
	s.InstanceId = &v
	return s
}

type StopGWSInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopGWSInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopGWSInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopGWSInstanceResponseBody) SetRequestId(v string) *StopGWSInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StopGWSInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopGWSInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopGWSInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopGWSInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopGWSInstanceResponse) SetHeaders(v map[string]*string) *StopGWSInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopGWSInstanceResponse) SetBody(v *StopGWSInstanceResponseBody) *StopGWSInstanceResponse {
	s.Body = v
	return s
}

type StopJobsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Jobs      *string `json:"Jobs,omitempty" xml:"Jobs,omitempty"`
}

func (s StopJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s StopJobsRequest) GoString() string {
	return s.String()
}

func (s *StopJobsRequest) SetClusterId(v string) *StopJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *StopJobsRequest) SetJobs(v string) *StopJobsRequest {
	s.Jobs = &v
	return s
}

type StopJobsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopJobsResponseBody) GoString() string {
	return s.String()
}

func (s *StopJobsResponseBody) SetRequestId(v string) *StopJobsResponseBody {
	s.RequestId = &v
	return s
}

type StopJobsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s StopJobsResponse) GoString() string {
	return s.String()
}

func (s *StopJobsResponse) SetHeaders(v map[string]*string) *StopJobsResponse {
	s.Headers = v
	return s
}

func (s *StopJobsResponse) SetBody(v *StopJobsResponseBody) *StopJobsResponse {
	s.Body = v
	return s
}

type StopNodesRequest struct {
	ClusterId *string                     `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Instance  []*StopNodesRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
	Role      *string                     `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s StopNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s StopNodesRequest) GoString() string {
	return s.String()
}

func (s *StopNodesRequest) SetClusterId(v string) *StopNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *StopNodesRequest) SetInstance(v []*StopNodesRequestInstance) *StopNodesRequest {
	s.Instance = v
	return s
}

func (s *StopNodesRequest) SetRole(v string) *StopNodesRequest {
	s.Role = &v
	return s
}

type StopNodesRequestInstance struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s StopNodesRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s StopNodesRequestInstance) GoString() string {
	return s.String()
}

func (s *StopNodesRequestInstance) SetId(v string) *StopNodesRequestInstance {
	s.Id = &v
	return s
}

type StopNodesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopNodesResponseBody) GoString() string {
	return s.String()
}

func (s *StopNodesResponseBody) SetRequestId(v string) *StopNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopNodesResponseBody) SetTaskId(v string) *StopNodesResponseBody {
	s.TaskId = &v
	return s
}

type StopNodesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s StopNodesResponse) GoString() string {
	return s.String()
}

func (s *StopNodesResponse) SetHeaders(v map[string]*string) *StopNodesResponse {
	s.Headers = v
	return s
}

func (s *StopNodesResponse) SetBody(v *StopNodesResponseBody) *StopNodesResponse {
	s.Body = v
	return s
}

type StopVisualServiceRequest struct {
	CidrIp    *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Port      *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s StopVisualServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopVisualServiceRequest) GoString() string {
	return s.String()
}

func (s *StopVisualServiceRequest) SetCidrIp(v string) *StopVisualServiceRequest {
	s.CidrIp = &v
	return s
}

func (s *StopVisualServiceRequest) SetClusterId(v string) *StopVisualServiceRequest {
	s.ClusterId = &v
	return s
}

func (s *StopVisualServiceRequest) SetPort(v int32) *StopVisualServiceRequest {
	s.Port = &v
	return s
}

type StopVisualServiceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopVisualServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopVisualServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StopVisualServiceResponseBody) SetMessage(v string) *StopVisualServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StopVisualServiceResponseBody) SetRequestId(v string) *StopVisualServiceResponseBody {
	s.RequestId = &v
	return s
}

type StopVisualServiceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopVisualServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopVisualServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopVisualServiceResponse) GoString() string {
	return s.String()
}

func (s *StopVisualServiceResponse) SetHeaders(v map[string]*string) *StopVisualServiceResponse {
	s.Headers = v
	return s
}

func (s *StopVisualServiceResponse) SetBody(v *StopVisualServiceResponseBody) *StopVisualServiceResponse {
	s.Body = v
	return s
}

type SubmitJobRequest struct {
	ArrayRequest       *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	ClockTime          *string `json:"ClockTime,omitempty" xml:"ClockTime,omitempty"`
	ClusterId          *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CommandLine        *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	ContainerId        *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	Gpu                *int32  `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	InputFileUrl       *string `json:"InputFileUrl,omitempty" xml:"InputFileUrl,omitempty"`
	JobQueue           *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	Mem                *string `json:"Mem,omitempty" xml:"Mem,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Node               *int32  `json:"Node,omitempty" xml:"Node,omitempty"`
	PackagePath        *string `json:"PackagePath,omitempty" xml:"PackagePath,omitempty"`
	PostCmdLine        *string `json:"PostCmdLine,omitempty" xml:"PostCmdLine,omitempty"`
	Priority           *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ReRunable          *bool   `json:"ReRunable,omitempty" xml:"ReRunable,omitempty"`
	RunasUser          *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	RunasUserPassword  *string `json:"RunasUserPassword,omitempty" xml:"RunasUserPassword,omitempty"`
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" xml:"StderrRedirectPath,omitempty"`
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" xml:"StdoutRedirectPath,omitempty"`
	Task               *int32  `json:"Task,omitempty" xml:"Task,omitempty"`
	Thread             *int32  `json:"Thread,omitempty" xml:"Thread,omitempty"`
	UnzipCmd           *string `json:"UnzipCmd,omitempty" xml:"UnzipCmd,omitempty"`
	Variables          *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s SubmitJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitJobRequest) SetArrayRequest(v string) *SubmitJobRequest {
	s.ArrayRequest = &v
	return s
}

func (s *SubmitJobRequest) SetClockTime(v string) *SubmitJobRequest {
	s.ClockTime = &v
	return s
}

func (s *SubmitJobRequest) SetClusterId(v string) *SubmitJobRequest {
	s.ClusterId = &v
	return s
}

func (s *SubmitJobRequest) SetCommandLine(v string) *SubmitJobRequest {
	s.CommandLine = &v
	return s
}

func (s *SubmitJobRequest) SetContainerId(v string) *SubmitJobRequest {
	s.ContainerId = &v
	return s
}

func (s *SubmitJobRequest) SetGpu(v int32) *SubmitJobRequest {
	s.Gpu = &v
	return s
}

func (s *SubmitJobRequest) SetInputFileUrl(v string) *SubmitJobRequest {
	s.InputFileUrl = &v
	return s
}

func (s *SubmitJobRequest) SetJobQueue(v string) *SubmitJobRequest {
	s.JobQueue = &v
	return s
}

func (s *SubmitJobRequest) SetMem(v string) *SubmitJobRequest {
	s.Mem = &v
	return s
}

func (s *SubmitJobRequest) SetName(v string) *SubmitJobRequest {
	s.Name = &v
	return s
}

func (s *SubmitJobRequest) SetNode(v int32) *SubmitJobRequest {
	s.Node = &v
	return s
}

func (s *SubmitJobRequest) SetPackagePath(v string) *SubmitJobRequest {
	s.PackagePath = &v
	return s
}

func (s *SubmitJobRequest) SetPostCmdLine(v string) *SubmitJobRequest {
	s.PostCmdLine = &v
	return s
}

func (s *SubmitJobRequest) SetPriority(v int32) *SubmitJobRequest {
	s.Priority = &v
	return s
}

func (s *SubmitJobRequest) SetReRunable(v bool) *SubmitJobRequest {
	s.ReRunable = &v
	return s
}

func (s *SubmitJobRequest) SetRunasUser(v string) *SubmitJobRequest {
	s.RunasUser = &v
	return s
}

func (s *SubmitJobRequest) SetRunasUserPassword(v string) *SubmitJobRequest {
	s.RunasUserPassword = &v
	return s
}

func (s *SubmitJobRequest) SetStderrRedirectPath(v string) *SubmitJobRequest {
	s.StderrRedirectPath = &v
	return s
}

func (s *SubmitJobRequest) SetStdoutRedirectPath(v string) *SubmitJobRequest {
	s.StdoutRedirectPath = &v
	return s
}

func (s *SubmitJobRequest) SetTask(v int32) *SubmitJobRequest {
	s.Task = &v
	return s
}

func (s *SubmitJobRequest) SetThread(v int32) *SubmitJobRequest {
	s.Thread = &v
	return s
}

func (s *SubmitJobRequest) SetUnzipCmd(v string) *SubmitJobRequest {
	s.UnzipCmd = &v
	return s
}

func (s *SubmitJobRequest) SetVariables(v string) *SubmitJobRequest {
	s.Variables = &v
	return s
}

type SubmitJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitJobResponseBody) SetJobId(v string) *SubmitJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitJobResponseBody) SetRequestId(v string) *SubmitJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitJobResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitJobResponse) SetHeaders(v map[string]*string) *SubmitJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitJobResponse) SetBody(v *SubmitJobResponseBody) *SubmitJobResponse {
	s.Body = v
	return s
}

type SyncUsersRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SyncUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncUsersRequest) GoString() string {
	return s.String()
}

func (s *SyncUsersRequest) SetClusterId(v string) *SyncUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *SyncUsersRequest) SetRegionId(v string) *SyncUsersRequest {
	s.RegionId = &v
	return s
}

type SyncUsersResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncUsersResponseBody) GoString() string {
	return s.String()
}

func (s *SyncUsersResponseBody) SetRequestId(v string) *SyncUsersResponseBody {
	s.RequestId = &v
	return s
}

type SyncUsersResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncUsersResponse) GoString() string {
	return s.String()
}

func (s *SyncUsersResponse) SetHeaders(v map[string]*string) *SyncUsersResponse {
	s.Headers = v
	return s
}

func (s *SyncUsersResponse) SetBody(v *SyncUsersResponseBody) *SyncUsersResponse {
	s.Body = v
	return s
}

type UninstallSoftwareRequest struct {
	Application *string `json:"Application,omitempty" xml:"Application,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UninstallSoftwareRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallSoftwareRequest) GoString() string {
	return s.String()
}

func (s *UninstallSoftwareRequest) SetApplication(v string) *UninstallSoftwareRequest {
	s.Application = &v
	return s
}

func (s *UninstallSoftwareRequest) SetClusterId(v string) *UninstallSoftwareRequest {
	s.ClusterId = &v
	return s
}

type UninstallSoftwareResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UninstallSoftwareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallSoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallSoftwareResponseBody) SetRequestId(v string) *UninstallSoftwareResponseBody {
	s.RequestId = &v
	return s
}

type UninstallSoftwareResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UninstallSoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UninstallSoftwareResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallSoftwareResponse) GoString() string {
	return s.String()
}

func (s *UninstallSoftwareResponse) SetHeaders(v map[string]*string) *UninstallSoftwareResponse {
	s.Headers = v
	return s
}

func (s *UninstallSoftwareResponse) SetBody(v *UninstallSoftwareResponseBody) *UninstallSoftwareResponse {
	s.Body = v
	return s
}

type UpdateClusterVolumesRequest struct {
	AdditionalVolumes []*UpdateClusterVolumesRequestAdditionalVolumes `json:"AdditionalVolumes,omitempty" xml:"AdditionalVolumes,omitempty" type:"Repeated"`
	ClusterId         *string                                         `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UpdateClusterVolumesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterVolumesRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterVolumesRequest) SetAdditionalVolumes(v []*UpdateClusterVolumesRequestAdditionalVolumes) *UpdateClusterVolumesRequest {
	s.AdditionalVolumes = v
	return s
}

func (s *UpdateClusterVolumesRequest) SetClusterId(v string) *UpdateClusterVolumesRequest {
	s.ClusterId = &v
	return s
}

type UpdateClusterVolumesRequestAdditionalVolumes struct {
	JobQueue         *string                                              `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	LocalDirectory   *string                                              `json:"LocalDirectory,omitempty" xml:"LocalDirectory,omitempty"`
	Location         *string                                              `json:"Location,omitempty" xml:"Location,omitempty"`
	RemoteDirectory  *string                                              `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	Roles            []*UpdateClusterVolumesRequestAdditionalVolumesRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	VolumeId         *string                                              `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	VolumeMountpoint *string                                              `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	VolumeProtocol   *string                                              `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	VolumeType       *string                                              `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
}

func (s UpdateClusterVolumesRequestAdditionalVolumes) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterVolumesRequestAdditionalVolumes) GoString() string {
	return s.String()
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetJobQueue(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.JobQueue = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetLocalDirectory(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.LocalDirectory = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetLocation(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.Location = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetRemoteDirectory(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.RemoteDirectory = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetRoles(v []*UpdateClusterVolumesRequestAdditionalVolumesRoles) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.Roles = v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetVolumeId(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.VolumeId = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetVolumeMountpoint(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.VolumeMountpoint = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetVolumeProtocol(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.VolumeProtocol = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetVolumeType(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.VolumeType = &v
	return s
}

type UpdateClusterVolumesRequestAdditionalVolumesRoles struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateClusterVolumesRequestAdditionalVolumesRoles) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterVolumesRequestAdditionalVolumesRoles) GoString() string {
	return s.String()
}

func (s *UpdateClusterVolumesRequestAdditionalVolumesRoles) SetName(v string) *UpdateClusterVolumesRequestAdditionalVolumesRoles {
	s.Name = &v
	return s
}

type UpdateClusterVolumesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateClusterVolumesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterVolumesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterVolumesResponseBody) SetRequestId(v string) *UpdateClusterVolumesResponseBody {
	s.RequestId = &v
	return s
}

type UpdateClusterVolumesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateClusterVolumesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateClusterVolumesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterVolumesResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterVolumesResponse) SetHeaders(v map[string]*string) *UpdateClusterVolumesResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterVolumesResponse) SetBody(v *UpdateClusterVolumesResponseBody) *UpdateClusterVolumesResponse {
	s.Body = v
	return s
}

type UpdateQueueConfigRequest struct {
	ClusterId           *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ComputeInstanceType *string `json:"ComputeInstanceType,omitempty" xml:"ComputeInstanceType,omitempty"`
	QueueName           *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	ResourceGroupId     *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UpdateQueueConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQueueConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateQueueConfigRequest) SetClusterId(v string) *UpdateQueueConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateQueueConfigRequest) SetComputeInstanceType(v string) *UpdateQueueConfigRequest {
	s.ComputeInstanceType = &v
	return s
}

func (s *UpdateQueueConfigRequest) SetQueueName(v string) *UpdateQueueConfigRequest {
	s.QueueName = &v
	return s
}

func (s *UpdateQueueConfigRequest) SetResourceGroupId(v string) *UpdateQueueConfigRequest {
	s.ResourceGroupId = &v
	return s
}

type UpdateQueueConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateQueueConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateQueueConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQueueConfigResponseBody) SetRequestId(v string) *UpdateQueueConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateQueueConfigResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateQueueConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateQueueConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateQueueConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateQueueConfigResponse) SetHeaders(v map[string]*string) *UpdateQueueConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateQueueConfigResponse) SetBody(v *UpdateQueueConfigResponseBody) *UpdateQueueConfigResponse {
	s.Body = v
	return s
}

type UpgradeClientRequest struct {
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UpgradeClientRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClientRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClientRequest) SetClientVersion(v string) *UpgradeClientRequest {
	s.ClientVersion = &v
	return s
}

func (s *UpgradeClientRequest) SetClusterId(v string) *UpgradeClientRequest {
	s.ClusterId = &v
	return s
}

type UpgradeClientResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeClientResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClientResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeClientResponseBody) SetRequestId(v string) *UpgradeClientResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeClientResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeClientResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeClientResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClientResponse) GoString() string {
	return s.String()
}

func (s *UpgradeClientResponse) SetHeaders(v map[string]*string) *UpgradeClientResponse {
	s.Headers = v
	return s
}

func (s *UpgradeClientResponse) SetBody(v *UpgradeClientResponseBody) *UpgradeClientResponse {
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
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ehpc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddContainerAppWithOptions(request *AddContainerAppRequest, runtime *util.RuntimeOptions) (_result *AddContainerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddContainerApp"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddContainerAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddContainerApp(request *AddContainerAppRequest) (_result *AddContainerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddContainerAppResponse{}
	_body, _err := client.AddContainerAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddExistedNodesWithOptions(request *AddExistedNodesRequest, runtime *util.RuntimeOptions) (_result *AddExistedNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddExistedNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddExistedNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddExistedNodes(request *AddExistedNodesRequest) (_result *AddExistedNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddExistedNodesResponse{}
	_body, _err := client.AddExistedNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLocalNodesWithOptions(request *AddLocalNodesRequest, runtime *util.RuntimeOptions) (_result *AddLocalNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddLocalNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddLocalNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLocalNodes(request *AddLocalNodesRequest) (_result *AddLocalNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLocalNodesResponse{}
	_body, _err := client.AddLocalNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddNodesWithOptions(request *AddNodesRequest, runtime *util.RuntimeOptions) (_result *AddNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddNodes(request *AddNodesRequest) (_result *AddNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddNodesResponse{}
	_body, _err := client.AddNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddQueueWithOptions(request *AddQueueRequest, runtime *util.RuntimeOptions) (_result *AddQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddQueue"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddQueue(request *AddQueueRequest) (_result *AddQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddQueueResponse{}
	_body, _err := client.AddQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSecurityGroupWithOptions(request *AddSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *AddSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddSecurityGroup"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddSecurityGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSecurityGroup(request *AddSecurityGroupRequest) (_result *AddSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSecurityGroupResponse{}
	_body, _err := client.AddSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUsersWithOptions(request *AddUsersRequest, runtime *util.RuntimeOptions) (_result *AddUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUsers"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUsers(request *AddUsersRequest) (_result *AddUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUsersResponse{}
	_body, _err := client.AddUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyNodesWithOptions(request *ApplyNodesRequest, runtime *util.RuntimeOptions) (_result *ApplyNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyNodes(request *ApplyNodesRequest) (_result *ApplyNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyNodesResponse{}
	_body, _err := client.ApplyNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGWSClusterWithOptions(request *CreateGWSClusterRequest, runtime *util.RuntimeOptions) (_result *CreateGWSClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGWSCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGWSClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGWSCluster(request *CreateGWSClusterRequest) (_result *CreateGWSClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGWSClusterResponse{}
	_body, _err := client.CreateGWSClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGWSImageWithOptions(request *CreateGWSImageRequest, runtime *util.RuntimeOptions) (_result *CreateGWSImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGWSImage"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGWSImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGWSImage(request *CreateGWSImageRequest) (_result *CreateGWSImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGWSImageResponse{}
	_body, _err := client.CreateGWSImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGWSInstanceWithOptions(request *CreateGWSInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateGWSInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGWSInstance"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGWSInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGWSInstance(request *CreateGWSInstanceRequest) (_result *CreateGWSInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGWSInstanceResponse{}
	_body, _err := client.CreateGWSInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHybridClusterWithOptions(request *CreateHybridClusterRequest, runtime *util.RuntimeOptions) (_result *CreateHybridClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHybridCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHybridClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHybridCluster(request *CreateHybridClusterRequest) (_result *CreateHybridClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHybridClusterResponse{}
	_body, _err := client.CreateHybridClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateJobFileWithOptions(request *CreateJobFileRequest, runtime *util.RuntimeOptions) (_result *CreateJobFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateJobFile"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateJobFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJobFile(request *CreateJobFileRequest) (_result *CreateJobFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateJobFileResponse{}
	_body, _err := client.CreateJobFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateJobTemplateWithOptions(request *CreateJobTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateJobTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateJobTemplate"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateJobTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJobTemplate(request *CreateJobTemplateRequest) (_result *CreateJobTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateJobTemplateResponse{}
	_body, _err := client.CreateJobTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteContainerAppsWithOptions(request *DeleteContainerAppsRequest, runtime *util.RuntimeOptions) (_result *DeleteContainerAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteContainerApps"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteContainerAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteContainerApps(request *DeleteContainerAppsRequest) (_result *DeleteContainerAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteContainerAppsResponse{}
	_body, _err := client.DeleteContainerAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGWSClusterWithOptions(request *DeleteGWSClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteGWSClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGWSCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGWSClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGWSCluster(request *DeleteGWSClusterRequest) (_result *DeleteGWSClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGWSClusterResponse{}
	_body, _err := client.DeleteGWSClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGWSInstanceWithOptions(request *DeleteGWSInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteGWSInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGWSInstance"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGWSInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGWSInstance(request *DeleteGWSInstanceRequest) (_result *DeleteGWSInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGWSInstanceResponse{}
	_body, _err := client.DeleteGWSInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageWithOptions(request *DeleteImageRequest, runtime *util.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteImage"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImage(request *DeleteImageRequest) (_result *DeleteImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImageResponse{}
	_body, _err := client.DeleteImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobTemplatesWithOptions(request *DeleteJobTemplatesRequest, runtime *util.RuntimeOptions) (_result *DeleteJobTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteJobTemplates"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteJobTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJobTemplates(request *DeleteJobTemplatesRequest) (_result *DeleteJobTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteJobTemplatesResponse{}
	_body, _err := client.DeleteJobTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobsWithOptions(request *DeleteJobsRequest, runtime *util.RuntimeOptions) (_result *DeleteJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteJobs"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJobs(request *DeleteJobsRequest) (_result *DeleteJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteJobsResponse{}
	_body, _err := client.DeleteJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNodesWithOptions(request *DeleteNodesRequest, runtime *util.RuntimeOptions) (_result *DeleteNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNodes(request *DeleteNodesRequest) (_result *DeleteNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNodesResponse{}
	_body, _err := client.DeleteNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteQueueWithOptions(request *DeleteQueueRequest, runtime *util.RuntimeOptions) (_result *DeleteQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteQueue"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteQueue(request *DeleteQueueRequest) (_result *DeleteQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteQueueResponse{}
	_body, _err := client.DeleteQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSecurityGroupWithOptions(request *DeleteSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSecurityGroup"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSecurityGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) (_result *DeleteSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecurityGroupResponse{}
	_body, _err := client.DeleteSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUsersWithOptions(request *DeleteUsersRequest, runtime *util.RuntimeOptions) (_result *DeleteUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUsers"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUsers(request *DeleteUsersRequest) (_result *DeleteUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUsersResponse{}
	_body, _err := client.DeleteUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutoScaleConfigWithOptions(request *DescribeAutoScaleConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeAutoScaleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAutoScaleConfig"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAutoScaleConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutoScaleConfig(request *DescribeAutoScaleConfigRequest) (_result *DescribeAutoScaleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutoScaleConfigResponse{}
	_body, _err := client.DescribeAutoScaleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterWithOptions(request *DescribeClusterRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCluster(request *DescribeClusterRequest) (_result *DescribeClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterResponse{}
	_body, _err := client.DescribeClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerAppWithOptions(request *DescribeContainerAppRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContainerApp"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContainerAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerApp(request *DescribeContainerAppRequest) (_result *DescribeContainerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerAppResponse{}
	_body, _err := client.DescribeContainerAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGWSClusterPolicyWithOptions(request *DescribeGWSClusterPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeGWSClusterPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AsyncMode)) {
		query["AsyncMode"] = request.AsyncMode
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGWSClusterPolicy"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGWSClusterPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGWSClusterPolicy(request *DescribeGWSClusterPolicyRequest) (_result *DescribeGWSClusterPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGWSClusterPolicyResponse{}
	_body, _err := client.DescribeGWSClusterPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGWSClustersWithOptions(request *DescribeGWSClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeGWSClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGWSClusters"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGWSClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGWSClusters(request *DescribeGWSClustersRequest) (_result *DescribeGWSClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGWSClustersResponse{}
	_body, _err := client.DescribeGWSClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGWSImagesWithOptions(request *DescribeGWSImagesRequest, runtime *util.RuntimeOptions) (_result *DescribeGWSImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGWSImages"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGWSImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGWSImages(request *DescribeGWSImagesRequest) (_result *DescribeGWSImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGWSImagesResponse{}
	_body, _err := client.DescribeGWSImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGWSInstancesWithOptions(request *DescribeGWSInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeGWSInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGWSInstances"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGWSInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGWSInstances(request *DescribeGWSInstancesRequest) (_result *DescribeGWSInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGWSInstancesResponse{}
	_body, _err := client.DescribeGWSInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageWithOptions(request *DescribeImageRequest, runtime *util.RuntimeOptions) (_result *DescribeImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImage"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImage(request *DescribeImageRequest) (_result *DescribeImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageResponse{}
	_body, _err := client.DescribeImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageGatewayConfigWithOptions(request *DescribeImageGatewayConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeImageGatewayConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImageGatewayConfig"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeImageGatewayConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImageGatewayConfig(request *DescribeImageGatewayConfigRequest) (_result *DescribeImageGatewayConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageGatewayConfigResponse{}
	_body, _err := client.DescribeImageGatewayConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImagePriceWithOptions(request *DescribeImagePriceRequest, runtime *util.RuntimeOptions) (_result *DescribeImagePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImagePrice"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeImagePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImagePrice(request *DescribeImagePriceRequest) (_result *DescribeImagePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImagePriceResponse{}
	_body, _err := client.DescribeImagePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeJobWithOptions(request *DescribeJobRequest, runtime *util.RuntimeOptions) (_result *DescribeJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeJob"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeJob(request *DescribeJobRequest) (_result *DescribeJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeJobResponse{}
	_body, _err := client.DescribeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNFSClientStatusWithOptions(request *DescribeNFSClientStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeNFSClientStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNFSClientStatus"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNFSClientStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNFSClientStatus(request *DescribeNFSClientStatusRequest) (_result *DescribeNFSClientStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNFSClientStatusResponse{}
	_body, _err := client.DescribeNFSClientStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePriceWithOptions(request *DescribePriceRequest, runtime *util.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePrice"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePrice(request *DescribePriceRequest) (_result *DescribePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePriceResponse{}
	_body, _err := client.DescribePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditJobTemplateWithOptions(request *EditJobTemplateRequest, runtime *util.RuntimeOptions) (_result *EditJobTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EditJobTemplate"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EditJobTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditJobTemplate(request *EditJobTemplateRequest) (_result *EditJobTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditJobTemplateResponse{}
	_body, _err := client.EditJobTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccountingReportWithOptions(request *GetAccountingReportRequest, runtime *util.RuntimeOptions) (_result *GetAccountingReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccountingReport"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountingReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccountingReport(request *GetAccountingReportRequest) (_result *GetAccountingReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountingReportResponse{}
	_body, _err := client.GetAccountingReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAutoScaleConfigWithOptions(request *GetAutoScaleConfigRequest, runtime *util.RuntimeOptions) (_result *GetAutoScaleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAutoScaleConfig"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAutoScaleConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAutoScaleConfig(request *GetAutoScaleConfigRequest) (_result *GetAutoScaleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAutoScaleConfigResponse{}
	_body, _err := client.GetAutoScaleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCloudMetricLogsWithOptions(request *GetCloudMetricLogsRequest, runtime *util.RuntimeOptions) (_result *GetCloudMetricLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCloudMetricLogs"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCloudMetricLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCloudMetricLogs(request *GetCloudMetricLogsRequest) (_result *GetCloudMetricLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCloudMetricLogsResponse{}
	_body, _err := client.GetCloudMetricLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCloudMetricProfilingWithOptions(request *GetCloudMetricProfilingRequest, runtime *util.RuntimeOptions) (_result *GetCloudMetricProfilingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCloudMetricProfiling"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCloudMetricProfilingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCloudMetricProfiling(request *GetCloudMetricProfilingRequest) (_result *GetCloudMetricProfilingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCloudMetricProfilingResponse{}
	_body, _err := client.GetCloudMetricProfilingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClusterVolumesWithOptions(request *GetClusterVolumesRequest, runtime *util.RuntimeOptions) (_result *GetClusterVolumesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetClusterVolumes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClusterVolumesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClusterVolumes(request *GetClusterVolumesRequest) (_result *GetClusterVolumesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClusterVolumesResponse{}
	_body, _err := client.GetClusterVolumesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGWSConnectTicketWithOptions(request *GetGWSConnectTicketRequest, runtime *util.RuntimeOptions) (_result *GetGWSConnectTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGWSConnectTicket"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGWSConnectTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGWSConnectTicket(request *GetGWSConnectTicketRequest) (_result *GetGWSConnectTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGWSConnectTicketResponse{}
	_body, _err := client.GetGWSConnectTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHybridClusterConfigWithOptions(request *GetHybridClusterConfigRequest, runtime *util.RuntimeOptions) (_result *GetHybridClusterConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHybridClusterConfig"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHybridClusterConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHybridClusterConfig(request *GetHybridClusterConfigRequest) (_result *GetHybridClusterConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHybridClusterConfigResponse{}
	_body, _err := client.GetHybridClusterConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIfEcsTypeSupportHtConfigWithOptions(request *GetIfEcsTypeSupportHtConfigRequest, runtime *util.RuntimeOptions) (_result *GetIfEcsTypeSupportHtConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIfEcsTypeSupportHtConfig"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIfEcsTypeSupportHtConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIfEcsTypeSupportHtConfig(request *GetIfEcsTypeSupportHtConfigRequest) (_result *GetIfEcsTypeSupportHtConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIfEcsTypeSupportHtConfigResponse{}
	_body, _err := client.GetIfEcsTypeSupportHtConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSchedulerInfoWithOptions(request *GetSchedulerInfoRequest, runtime *util.RuntimeOptions) (_result *GetSchedulerInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSchedulerInfo"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSchedulerInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSchedulerInfo(request *GetSchedulerInfoRequest) (_result *GetSchedulerInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSchedulerInfoResponse{}
	_body, _err := client.GetSchedulerInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVisualServiceStatusWithOptions(request *GetVisualServiceStatusRequest, runtime *util.RuntimeOptions) (_result *GetVisualServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVisualServiceStatus"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVisualServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVisualServiceStatus(request *GetVisualServiceStatusRequest) (_result *GetVisualServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVisualServiceStatusResponse{}
	_body, _err := client.GetVisualServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitializeEHPCWithOptions(request *InitializeEHPCRequest, runtime *util.RuntimeOptions) (_result *InitializeEHPCResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InitializeEHPC"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitializeEHPCResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitializeEHPC(request *InitializeEHPCRequest) (_result *InitializeEHPCResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitializeEHPCResponse{}
	_body, _err := client.InitializeEHPCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallSoftwareWithOptions(request *InstallSoftwareRequest, runtime *util.RuntimeOptions) (_result *InstallSoftwareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallSoftware"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallSoftwareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallSoftware(request *InstallSoftwareRequest) (_result *InstallSoftwareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallSoftwareResponse{}
	_body, _err := client.InstallSoftwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvokeShellCommandWithOptions(request *InvokeShellCommandRequest, runtime *util.RuntimeOptions) (_result *InvokeShellCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InvokeShellCommand"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InvokeShellCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvokeShellCommand(request *InvokeShellCommandRequest) (_result *InvokeShellCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InvokeShellCommandResponse{}
	_body, _err := client.InvokeShellCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAvailableEcsTypesWithOptions(request *ListAvailableEcsTypesRequest, runtime *util.RuntimeOptions) (_result *ListAvailableEcsTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAvailableEcsTypes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAvailableEcsTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAvailableEcsTypes(request *ListAvailableEcsTypesRequest) (_result *ListAvailableEcsTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAvailableEcsTypesResponse{}
	_body, _err := client.ListAvailableEcsTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCloudMetricProfilingsWithOptions(request *ListCloudMetricProfilingsRequest, runtime *util.RuntimeOptions) (_result *ListCloudMetricProfilingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCloudMetricProfilings"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCloudMetricProfilingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCloudMetricProfilings(request *ListCloudMetricProfilingsRequest) (_result *ListCloudMetricProfilingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCloudMetricProfilingsResponse{}
	_body, _err := client.ListCloudMetricProfilingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterLogsWithOptions(request *ListClusterLogsRequest, runtime *util.RuntimeOptions) (_result *ListClusterLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterLogs"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterLogs(request *ListClusterLogsRequest) (_result *ListClusterLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterLogsResponse{}
	_body, _err := client.ListClusterLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusters"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClustersMetaWithOptions(request *ListClustersMetaRequest, runtime *util.RuntimeOptions) (_result *ListClustersMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClustersMeta"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClustersMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClustersMeta(request *ListClustersMetaRequest) (_result *ListClustersMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClustersMetaResponse{}
	_body, _err := client.ListClustersMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCommandsWithOptions(request *ListCommandsRequest, runtime *util.RuntimeOptions) (_result *ListCommandsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCommands"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCommandsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCommands(request *ListCommandsRequest) (_result *ListCommandsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCommandsResponse{}
	_body, _err := client.ListCommandsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListContainerAppsWithOptions(request *ListContainerAppsRequest, runtime *util.RuntimeOptions) (_result *ListContainerAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListContainerApps"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListContainerAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListContainerApps(request *ListContainerAppsRequest) (_result *ListContainerAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListContainerAppsResponse{}
	_body, _err := client.ListContainerAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListContainerImagesWithOptions(request *ListContainerImagesRequest, runtime *util.RuntimeOptions) (_result *ListContainerImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListContainerImages"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListContainerImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListContainerImages(request *ListContainerImagesRequest) (_result *ListContainerImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListContainerImagesResponse{}
	_body, _err := client.ListContainerImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCpfsFileSystemsWithOptions(request *ListCpfsFileSystemsRequest, runtime *util.RuntimeOptions) (_result *ListCpfsFileSystemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCpfsFileSystems"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCpfsFileSystemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCpfsFileSystems(request *ListCpfsFileSystemsRequest) (_result *ListCpfsFileSystemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCpfsFileSystemsResponse{}
	_body, _err := client.ListCpfsFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCurrentClientVersionWithOptions(runtime *util.RuntimeOptions) (_result *ListCurrentClientVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListCurrentClientVersion"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCurrentClientVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCurrentClientVersion() (_result *ListCurrentClientVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCurrentClientVersionResponse{}
	_body, _err := client.ListCurrentClientVersionWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCustomImagesWithOptions(request *ListCustomImagesRequest, runtime *util.RuntimeOptions) (_result *ListCustomImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomImages"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCustomImages(request *ListCustomImagesRequest) (_result *ListCustomImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCustomImagesResponse{}
	_body, _err := client.ListCustomImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFileSystemWithMountTargetsWithOptions(request *ListFileSystemWithMountTargetsRequest, runtime *util.RuntimeOptions) (_result *ListFileSystemWithMountTargetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFileSystemWithMountTargets"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFileSystemWithMountTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFileSystemWithMountTargets(request *ListFileSystemWithMountTargetsRequest) (_result *ListFileSystemWithMountTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFileSystemWithMountTargetsResponse{}
	_body, _err := client.ListFileSystemWithMountTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImagesWithOptions(request *ListImagesRequest, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImages"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstalledSoftwareWithOptions(request *ListInstalledSoftwareRequest, runtime *util.RuntimeOptions) (_result *ListInstalledSoftwareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstalledSoftware"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstalledSoftwareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstalledSoftware(request *ListInstalledSoftwareRequest) (_result *ListInstalledSoftwareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstalledSoftwareResponse{}
	_body, _err := client.ListInstalledSoftwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInvocationResultsWithOptions(request *ListInvocationResultsRequest, runtime *util.RuntimeOptions) (_result *ListInvocationResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInvocationResults"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInvocationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInvocationResults(request *ListInvocationResultsRequest) (_result *ListInvocationResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInvocationResultsResponse{}
	_body, _err := client.ListInvocationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInvocationStatusWithOptions(request *ListInvocationStatusRequest, runtime *util.RuntimeOptions) (_result *ListInvocationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInvocationStatus"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInvocationStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInvocationStatus(request *ListInvocationStatusRequest) (_result *ListInvocationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInvocationStatusResponse{}
	_body, _err := client.ListInvocationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobTemplatesWithOptions(request *ListJobTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListJobTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobTemplates"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobTemplates(request *ListJobTemplatesRequest) (_result *ListJobTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJobTemplatesResponse{}
	_body, _err := client.ListJobTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobsWithOptions(request *ListJobsRequest, runtime *util.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobs"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobs(request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodesWithOptions(request *ListNodesRequest, runtime *util.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodesByQueueWithOptions(request *ListNodesByQueueRequest, runtime *util.RuntimeOptions) (_result *ListNodesByQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodesByQueue"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesByQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodesByQueue(request *ListNodesByQueueRequest) (_result *ListNodesByQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesByQueueResponse{}
	_body, _err := client.ListNodesByQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodesNoPagingWithOptions(request *ListNodesNoPagingRequest, runtime *util.RuntimeOptions) (_result *ListNodesNoPagingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodesNoPaging"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesNoPagingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodesNoPaging(request *ListNodesNoPagingRequest) (_result *ListNodesNoPagingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesNoPagingResponse{}
	_body, _err := client.ListNodesNoPagingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPreferredEcsTypesWithOptions(request *ListPreferredEcsTypesRequest, runtime *util.RuntimeOptions) (_result *ListPreferredEcsTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPreferredEcsTypes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPreferredEcsTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPreferredEcsTypes(request *ListPreferredEcsTypesRequest) (_result *ListPreferredEcsTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPreferredEcsTypesResponse{}
	_body, _err := client.ListPreferredEcsTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQueuesWithOptions(request *ListQueuesRequest, runtime *util.RuntimeOptions) (_result *ListQueuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueues"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQueues(request *ListQueuesRequest) (_result *ListQueuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQueuesResponse{}
	_body, _err := client.ListQueuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegionsWithOptions(runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSecurityGroupsWithOptions(request *ListSecurityGroupsRequest, runtime *util.RuntimeOptions) (_result *ListSecurityGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSecurityGroups"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSecurityGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSecurityGroups(request *ListSecurityGroupsRequest) (_result *ListSecurityGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSecurityGroupsResponse{}
	_body, _err := client.ListSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSoftwaresWithOptions(request *ListSoftwaresRequest, runtime *util.RuntimeOptions) (_result *ListSoftwaresResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSoftwares"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSoftwaresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSoftwares(request *ListSoftwaresRequest) (_result *ListSoftwaresResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSoftwaresResponse{}
	_body, _err := client.ListSoftwaresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTasksWithOptions(request *ListTasksRequest, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTasks"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTasks(request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVolumesWithOptions(request *ListVolumesRequest, runtime *util.RuntimeOptions) (_result *ListVolumesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVolumes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVolumesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVolumes(request *ListVolumesRequest) (_result *ListVolumesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVolumesResponse{}
	_body, _err := client.ListVolumesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterAttributesWithOptions(request *ModifyClusterAttributesRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterAttributes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterAttributes(request *ModifyClusterAttributesRequest) (_result *ModifyClusterAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterAttributesResponse{}
	_body, _err := client.ModifyClusterAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyContainerAppAttributesWithOptions(request *ModifyContainerAppAttributesRequest, runtime *util.RuntimeOptions) (_result *ModifyContainerAppAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyContainerAppAttributes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyContainerAppAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyContainerAppAttributes(request *ModifyContainerAppAttributesRequest) (_result *ModifyContainerAppAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyContainerAppAttributesResponse{}
	_body, _err := client.ModifyContainerAppAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyImageGatewayConfigWithOptions(request *ModifyImageGatewayConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyImageGatewayConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyImageGatewayConfig"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyImageGatewayConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyImageGatewayConfig(request *ModifyImageGatewayConfigRequest) (_result *ModifyImageGatewayConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyImageGatewayConfigResponse{}
	_body, _err := client.ModifyImageGatewayConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserGroupsWithOptions(request *ModifyUserGroupsRequest, runtime *util.RuntimeOptions) (_result *ModifyUserGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUserGroups"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUserGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserGroups(request *ModifyUserGroupsRequest) (_result *ModifyUserGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserGroupsResponse{}
	_body, _err := client.ModifyUserGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserPasswordsWithOptions(request *ModifyUserPasswordsRequest, runtime *util.RuntimeOptions) (_result *ModifyUserPasswordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUserPasswords"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUserPasswordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserPasswords(request *ModifyUserPasswordsRequest) (_result *ModifyUserPasswordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserPasswordsResponse{}
	_body, _err := client.ModifyUserPasswordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyVisualServicePasswdWithOptions(request *ModifyVisualServicePasswdRequest, runtime *util.RuntimeOptions) (_result *ModifyVisualServicePasswdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVisualServicePasswd"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVisualServicePasswdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyVisualServicePasswd(request *ModifyVisualServicePasswdRequest) (_result *ModifyVisualServicePasswdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVisualServicePasswdResponse{}
	_body, _err := client.ModifyVisualServicePasswdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MountNFSWithOptions(request *MountNFSRequest, runtime *util.RuntimeOptions) (_result *MountNFSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MountNFS"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MountNFSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MountNFS(request *MountNFSRequest) (_result *MountNFSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MountNFSResponse{}
	_body, _err := client.MountNFSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PullImageWithOptions(request *PullImageRequest, runtime *util.RuntimeOptions) (_result *PullImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PullImage"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PullImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PullImage(request *PullImageRequest) (_result *PullImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PullImageResponse{}
	_body, _err := client.PullImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryServicePackAndPriceWithOptions(runtime *util.RuntimeOptions) (_result *QueryServicePackAndPriceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QueryServicePackAndPrice"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryServicePackAndPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryServicePackAndPrice() (_result *QueryServicePackAndPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryServicePackAndPriceResponse{}
	_body, _err := client.QueryServicePackAndPriceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecoverClusterWithOptions(request *RecoverClusterRequest, runtime *util.RuntimeOptions) (_result *RecoverClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecoverCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecoverClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecoverCluster(request *RecoverClusterRequest) (_result *RecoverClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecoverClusterResponse{}
	_body, _err := client.RecoverClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RerunJobsWithOptions(request *RerunJobsRequest, runtime *util.RuntimeOptions) (_result *RerunJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RerunJobs"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RerunJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RerunJobs(request *RerunJobsRequest) (_result *RerunJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RerunJobsResponse{}
	_body, _err := client.RerunJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetNodesWithOptions(request *ResetNodesRequest, runtime *util.RuntimeOptions) (_result *ResetNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetNodes(request *ResetNodesRequest) (_result *ResetNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetNodesResponse{}
	_body, _err := client.ResetNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunCloudMetricProfilingWithOptions(request *RunCloudMetricProfilingRequest, runtime *util.RuntimeOptions) (_result *RunCloudMetricProfilingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunCloudMetricProfiling"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunCloudMetricProfilingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunCloudMetricProfiling(request *RunCloudMetricProfilingRequest) (_result *RunCloudMetricProfilingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCloudMetricProfilingResponse{}
	_body, _err := client.RunCloudMetricProfilingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAutoScaleConfigWithOptions(request *SetAutoScaleConfigRequest, runtime *util.RuntimeOptions) (_result *SetAutoScaleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetAutoScaleConfig"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetAutoScaleConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAutoScaleConfig(request *SetAutoScaleConfigRequest) (_result *SetAutoScaleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAutoScaleConfigResponse{}
	_body, _err := client.SetAutoScaleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGWSClusterPolicyWithOptions(request *SetGWSClusterPolicyRequest, runtime *util.RuntimeOptions) (_result *SetGWSClusterPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AsyncMode)) {
		query["AsyncMode"] = request.AsyncMode
	}

	if !tea.BoolValue(util.IsUnset(request.Clipboard)) {
		query["Clipboard"] = request.Clipboard
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.LocalDrive)) {
		query["LocalDrive"] = request.LocalDrive
	}

	if !tea.BoolValue(util.IsUnset(request.UdpPort)) {
		query["UdpPort"] = request.UdpPort
	}

	if !tea.BoolValue(util.IsUnset(request.UsbRedirect)) {
		query["UsbRedirect"] = request.UsbRedirect
	}

	if !tea.BoolValue(util.IsUnset(request.Watermark)) {
		query["Watermark"] = request.Watermark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetGWSClusterPolicy"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetGWSClusterPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGWSClusterPolicy(request *SetGWSClusterPolicyRequest) (_result *SetGWSClusterPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGWSClusterPolicyResponse{}
	_body, _err := client.SetGWSClusterPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGWSInstanceNameWithOptions(request *SetGWSInstanceNameRequest, runtime *util.RuntimeOptions) (_result *SetGWSInstanceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetGWSInstanceName"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetGWSInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGWSInstanceName(request *SetGWSInstanceNameRequest) (_result *SetGWSInstanceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGWSInstanceNameResponse{}
	_body, _err := client.SetGWSInstanceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGWSInstanceUserWithOptions(request *SetGWSInstanceUserRequest, runtime *util.RuntimeOptions) (_result *SetGWSInstanceUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetGWSInstanceUser"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetGWSInstanceUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGWSInstanceUser(request *SetGWSInstanceUserRequest) (_result *SetGWSInstanceUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGWSInstanceUserResponse{}
	_body, _err := client.SetGWSInstanceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetQueueWithOptions(request *SetQueueRequest, runtime *util.RuntimeOptions) (_result *SetQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetQueue"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetQueue(request *SetQueueRequest) (_result *SetQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetQueueResponse{}
	_body, _err := client.SetQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetSchedulerInfoWithOptions(request *SetSchedulerInfoRequest, runtime *util.RuntimeOptions) (_result *SetSchedulerInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetSchedulerInfo"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetSchedulerInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetSchedulerInfo(request *SetSchedulerInfoRequest) (_result *SetSchedulerInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetSchedulerInfoResponse{}
	_body, _err := client.SetSchedulerInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartClusterWithOptions(request *StartClusterRequest, runtime *util.RuntimeOptions) (_result *StartClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCluster(request *StartClusterRequest) (_result *StartClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartClusterResponse{}
	_body, _err := client.StartClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartGWSInstanceWithOptions(request *StartGWSInstanceRequest, runtime *util.RuntimeOptions) (_result *StartGWSInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartGWSInstance"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartGWSInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartGWSInstance(request *StartGWSInstanceRequest) (_result *StartGWSInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartGWSInstanceResponse{}
	_body, _err := client.StartGWSInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartNodesWithOptions(request *StartNodesRequest, runtime *util.RuntimeOptions) (_result *StartNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartNodes(request *StartNodesRequest) (_result *StartNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartNodesResponse{}
	_body, _err := client.StartNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartVisualServiceWithOptions(request *StartVisualServiceRequest, runtime *util.RuntimeOptions) (_result *StartVisualServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartVisualService"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartVisualServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartVisualService(request *StartVisualServiceRequest) (_result *StartVisualServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartVisualServiceResponse{}
	_body, _err := client.StartVisualServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopClusterWithOptions(request *StopClusterRequest, runtime *util.RuntimeOptions) (_result *StopClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopCluster"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopCluster(request *StopClusterRequest) (_result *StopClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopClusterResponse{}
	_body, _err := client.StopClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopGWSInstanceWithOptions(request *StopGWSInstanceRequest, runtime *util.RuntimeOptions) (_result *StopGWSInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopGWSInstance"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopGWSInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopGWSInstance(request *StopGWSInstanceRequest) (_result *StopGWSInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopGWSInstanceResponse{}
	_body, _err := client.StopGWSInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopJobsWithOptions(request *StopJobsRequest, runtime *util.RuntimeOptions) (_result *StopJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopJobs"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopJobs(request *StopJobsRequest) (_result *StopJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopJobsResponse{}
	_body, _err := client.StopJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopNodesWithOptions(request *StopNodesRequest, runtime *util.RuntimeOptions) (_result *StopNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopNodes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopNodes(request *StopNodesRequest) (_result *StopNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopNodesResponse{}
	_body, _err := client.StopNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopVisualServiceWithOptions(request *StopVisualServiceRequest, runtime *util.RuntimeOptions) (_result *StopVisualServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopVisualService"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopVisualServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopVisualService(request *StopVisualServiceRequest) (_result *StopVisualServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopVisualServiceResponse{}
	_body, _err := client.StopVisualServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitJobWithOptions(request *SubmitJobRequest, runtime *util.RuntimeOptions) (_result *SubmitJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitJob"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitJob(request *SubmitJobRequest) (_result *SubmitJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitJobResponse{}
	_body, _err := client.SubmitJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncUsersWithOptions(request *SyncUsersRequest, runtime *util.RuntimeOptions) (_result *SyncUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncUsers"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncUsers(request *SyncUsersRequest) (_result *SyncUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncUsersResponse{}
	_body, _err := client.SyncUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UninstallSoftwareWithOptions(request *UninstallSoftwareRequest, runtime *util.RuntimeOptions) (_result *UninstallSoftwareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UninstallSoftware"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UninstallSoftwareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UninstallSoftware(request *UninstallSoftwareRequest) (_result *UninstallSoftwareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UninstallSoftwareResponse{}
	_body, _err := client.UninstallSoftwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateClusterVolumesWithOptions(request *UpdateClusterVolumesRequest, runtime *util.RuntimeOptions) (_result *UpdateClusterVolumesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateClusterVolumes"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateClusterVolumesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateClusterVolumes(request *UpdateClusterVolumesRequest) (_result *UpdateClusterVolumesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateClusterVolumesResponse{}
	_body, _err := client.UpdateClusterVolumesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateQueueConfigWithOptions(request *UpdateQueueConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateQueueConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateQueueConfig"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateQueueConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateQueueConfig(request *UpdateQueueConfigRequest) (_result *UpdateQueueConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateQueueConfigResponse{}
	_body, _err := client.UpdateQueueConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeClientWithOptions(request *UpgradeClientRequest, runtime *util.RuntimeOptions) (_result *UpgradeClientResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeClient"),
		Version:     tea.String("2018-04-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeClient(request *UpgradeClientRequest) (_result *UpgradeClientResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeClientResponse{}
	_body, _err := client.UpgradeClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
