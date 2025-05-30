// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddonNodeTemplate struct {
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// 1
	AutoRenewPeriod *int32                        `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	DataDisks       []*AddonNodeTemplateDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// true
	EnableHT *bool `json:"EnableHT,omitempty" xml:"EnableHT,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20221102.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ecs.c7.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CentOS  7.6 64 位
	OsName *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CentOS  7.6 64 bit
	OsNameEN *string `json:"OsNameEN,omitempty" xml:"OsNameEN,omitempty"`
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// Month
	PeriodUnit       *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// example:
	//
	// 0.97
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// example:
	//
	// NoSpot
	SpotStrategy *string                      `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDisk   *AddonNodeTemplateSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
}

func (s AddonNodeTemplate) String() string {
	return tea.Prettify(s)
}

func (s AddonNodeTemplate) GoString() string {
	return s.String()
}

func (s *AddonNodeTemplate) SetAutoRenew(v bool) *AddonNodeTemplate {
	s.AutoRenew = &v
	return s
}

func (s *AddonNodeTemplate) SetAutoRenewPeriod(v int32) *AddonNodeTemplate {
	s.AutoRenewPeriod = &v
	return s
}

func (s *AddonNodeTemplate) SetDataDisks(v []*AddonNodeTemplateDataDisks) *AddonNodeTemplate {
	s.DataDisks = v
	return s
}

func (s *AddonNodeTemplate) SetDuration(v int32) *AddonNodeTemplate {
	s.Duration = &v
	return s
}

func (s *AddonNodeTemplate) SetEnableHT(v bool) *AddonNodeTemplate {
	s.EnableHT = &v
	return s
}

func (s *AddonNodeTemplate) SetImageId(v string) *AddonNodeTemplate {
	s.ImageId = &v
	return s
}

func (s *AddonNodeTemplate) SetInstanceChargeType(v string) *AddonNodeTemplate {
	s.InstanceChargeType = &v
	return s
}

func (s *AddonNodeTemplate) SetInstanceId(v string) *AddonNodeTemplate {
	s.InstanceId = &v
	return s
}

func (s *AddonNodeTemplate) SetInstanceType(v string) *AddonNodeTemplate {
	s.InstanceType = &v
	return s
}

func (s *AddonNodeTemplate) SetOsName(v string) *AddonNodeTemplate {
	s.OsName = &v
	return s
}

func (s *AddonNodeTemplate) SetOsNameEN(v string) *AddonNodeTemplate {
	s.OsNameEN = &v
	return s
}

func (s *AddonNodeTemplate) SetPeriod(v int32) *AddonNodeTemplate {
	s.Period = &v
	return s
}

func (s *AddonNodeTemplate) SetPeriodUnit(v string) *AddonNodeTemplate {
	s.PeriodUnit = &v
	return s
}

func (s *AddonNodeTemplate) SetPrivateIpAddress(v string) *AddonNodeTemplate {
	s.PrivateIpAddress = &v
	return s
}

func (s *AddonNodeTemplate) SetSpotPriceLimit(v float32) *AddonNodeTemplate {
	s.SpotPriceLimit = &v
	return s
}

func (s *AddonNodeTemplate) SetSpotStrategy(v string) *AddonNodeTemplate {
	s.SpotStrategy = &v
	return s
}

func (s *AddonNodeTemplate) SetSystemDisk(v *AddonNodeTemplateSystemDisk) *AddonNodeTemplate {
	s.SystemDisk = v
	return s
}

type AddonNodeTemplateDataDisks struct {
	// example:
	//
	// cloud_auto
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// false
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// example:
	//
	// PL0
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s AddonNodeTemplateDataDisks) String() string {
	return tea.Prettify(s)
}

func (s AddonNodeTemplateDataDisks) GoString() string {
	return s.String()
}

func (s *AddonNodeTemplateDataDisks) SetCategory(v string) *AddonNodeTemplateDataDisks {
	s.Category = &v
	return s
}

func (s *AddonNodeTemplateDataDisks) SetDeleteWithInstance(v bool) *AddonNodeTemplateDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *AddonNodeTemplateDataDisks) SetLevel(v string) *AddonNodeTemplateDataDisks {
	s.Level = &v
	return s
}

func (s *AddonNodeTemplateDataDisks) SetSize(v int32) *AddonNodeTemplateDataDisks {
	s.Size = &v
	return s
}

type AddonNodeTemplateSystemDisk struct {
	// example:
	//
	// cloud_auto
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// PL0
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s AddonNodeTemplateSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s AddonNodeTemplateSystemDisk) GoString() string {
	return s.String()
}

func (s *AddonNodeTemplateSystemDisk) SetCategory(v string) *AddonNodeTemplateSystemDisk {
	s.Category = &v
	return s
}

func (s *AddonNodeTemplateSystemDisk) SetLevel(v string) *AddonNodeTemplateSystemDisk {
	s.Level = &v
	return s
}

func (s *AddonNodeTemplateSystemDisk) SetSize(v int32) *AddonNodeTemplateSystemDisk {
	s.Size = &v
	return s
}

type NodeTemplate struct {
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// 1
	AutoRenewPeriod *int32                   `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	DataDisks       []*NodeTemplateDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// true
	EnableHT *bool `json:"EnableHT,omitempty" xml:"EnableHT,omitempty"`
	// example:
	//
	// aliyun_3_x64_20G_alibase_20221102.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// example:
	//
	// ecs.c7.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// 0.97
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// example:
	//
	// NoSpot
	SpotStrategy *string                 `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDisk   *NodeTemplateSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
}

func (s NodeTemplate) String() string {
	return tea.Prettify(s)
}

func (s NodeTemplate) GoString() string {
	return s.String()
}

func (s *NodeTemplate) SetAutoRenew(v bool) *NodeTemplate {
	s.AutoRenew = &v
	return s
}

func (s *NodeTemplate) SetAutoRenewPeriod(v int32) *NodeTemplate {
	s.AutoRenewPeriod = &v
	return s
}

func (s *NodeTemplate) SetDataDisks(v []*NodeTemplateDataDisks) *NodeTemplate {
	s.DataDisks = v
	return s
}

func (s *NodeTemplate) SetDuration(v int32) *NodeTemplate {
	s.Duration = &v
	return s
}

func (s *NodeTemplate) SetEnableHT(v bool) *NodeTemplate {
	s.EnableHT = &v
	return s
}

func (s *NodeTemplate) SetImageId(v string) *NodeTemplate {
	s.ImageId = &v
	return s
}

func (s *NodeTemplate) SetInstanceChargeType(v string) *NodeTemplate {
	s.InstanceChargeType = &v
	return s
}

func (s *NodeTemplate) SetInstanceType(v string) *NodeTemplate {
	s.InstanceType = &v
	return s
}

func (s *NodeTemplate) SetPeriod(v int32) *NodeTemplate {
	s.Period = &v
	return s
}

func (s *NodeTemplate) SetPeriodUnit(v string) *NodeTemplate {
	s.PeriodUnit = &v
	return s
}

func (s *NodeTemplate) SetSpotPriceLimit(v float32) *NodeTemplate {
	s.SpotPriceLimit = &v
	return s
}

func (s *NodeTemplate) SetSpotStrategy(v string) *NodeTemplate {
	s.SpotStrategy = &v
	return s
}

func (s *NodeTemplate) SetSystemDisk(v *NodeTemplateSystemDisk) *NodeTemplate {
	s.SystemDisk = v
	return s
}

type NodeTemplateDataDisks struct {
	// example:
	//
	// cloud_auto
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// false
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// example:
	//
	// PL0
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// /data1
	MountDir *string `json:"MountDir,omitempty" xml:"MountDir,omitempty"`
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// s-bp1ei2b44ripxuo46hym
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s NodeTemplateDataDisks) String() string {
	return tea.Prettify(s)
}

func (s NodeTemplateDataDisks) GoString() string {
	return s.String()
}

func (s *NodeTemplateDataDisks) SetCategory(v string) *NodeTemplateDataDisks {
	s.Category = &v
	return s
}

func (s *NodeTemplateDataDisks) SetDeleteWithInstance(v bool) *NodeTemplateDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *NodeTemplateDataDisks) SetDevice(v string) *NodeTemplateDataDisks {
	s.Device = &v
	return s
}

func (s *NodeTemplateDataDisks) SetLevel(v string) *NodeTemplateDataDisks {
	s.Level = &v
	return s
}

func (s *NodeTemplateDataDisks) SetMountDir(v string) *NodeTemplateDataDisks {
	s.MountDir = &v
	return s
}

func (s *NodeTemplateDataDisks) SetSize(v int32) *NodeTemplateDataDisks {
	s.Size = &v
	return s
}

func (s *NodeTemplateDataDisks) SetSnapshotId(v string) *NodeTemplateDataDisks {
	s.SnapshotId = &v
	return s
}

type NodeTemplateSystemDisk struct {
	// example:
	//
	// cloud_auto
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// PL0
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s NodeTemplateSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s NodeTemplateSystemDisk) GoString() string {
	return s.String()
}

func (s *NodeTemplateSystemDisk) SetCategory(v string) *NodeTemplateSystemDisk {
	s.Category = &v
	return s
}

func (s *NodeTemplateSystemDisk) SetLevel(v string) *NodeTemplateSystemDisk {
	s.Level = &v
	return s
}

func (s *NodeTemplateSystemDisk) SetSize(v int32) *NodeTemplateSystemDisk {
	s.Size = &v
	return s
}

type QueueTemplate struct {
	// example:
	//
	// PriorityInstanceType
	AllocationStrategy *string         `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	ComputeNodes       []*NodeTemplate `json:"ComputeNodes,omitempty" xml:"ComputeNodes,omitempty" type:"Repeated"`
	// example:
	//
	// true
	EnableScaleIn *bool `json:"EnableScaleIn,omitempty" xml:"EnableScaleIn,omitempty"`
	// example:
	//
	// true
	EnableScaleOut *bool `json:"EnableScaleOut,omitempty" xml:"EnableScaleOut,omitempty"`
	// example:
	//
	// compute
	HostnamePrefix *string `json:"HostnamePrefix,omitempty" xml:"HostnamePrefix,omitempty"`
	HostnameSuffix *string `json:"HostnameSuffix,omitempty" xml:"HostnameSuffix,omitempty"`
	// example:
	//
	// 0
	InitialCount *int32 `json:"InitialCount,omitempty" xml:"InitialCount,omitempty"`
	// example:
	//
	// erdma
	InterConnect   *string   `json:"InterConnect,omitempty" xml:"InterConnect,omitempty"`
	KeepAliveNodes []*string `json:"KeepAliveNodes,omitempty" xml:"KeepAliveNodes,omitempty" type:"Repeated"`
	// example:
	//
	// 1000
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// example:
	//
	// 99
	MaxCountPerCycle *int64 `json:"MaxCountPerCycle,omitempty" xml:"MaxCountPerCycle,omitempty"`
	// example:
	//
	// 0
	MinCount *int32 `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// comp
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// example:
	//
	// AliyunECSInstanceForEHPCRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// example:
	//
	// rnp-756vlp7a
	ReservedNodePoolId *string   `json:"ReservedNodePoolId,omitempty" xml:"ReservedNodePoolId,omitempty"`
	VSwitchIds         []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s QueueTemplate) String() string {
	return tea.Prettify(s)
}

func (s QueueTemplate) GoString() string {
	return s.String()
}

func (s *QueueTemplate) SetAllocationStrategy(v string) *QueueTemplate {
	s.AllocationStrategy = &v
	return s
}

func (s *QueueTemplate) SetComputeNodes(v []*NodeTemplate) *QueueTemplate {
	s.ComputeNodes = v
	return s
}

func (s *QueueTemplate) SetEnableScaleIn(v bool) *QueueTemplate {
	s.EnableScaleIn = &v
	return s
}

func (s *QueueTemplate) SetEnableScaleOut(v bool) *QueueTemplate {
	s.EnableScaleOut = &v
	return s
}

func (s *QueueTemplate) SetHostnamePrefix(v string) *QueueTemplate {
	s.HostnamePrefix = &v
	return s
}

func (s *QueueTemplate) SetHostnameSuffix(v string) *QueueTemplate {
	s.HostnameSuffix = &v
	return s
}

func (s *QueueTemplate) SetInitialCount(v int32) *QueueTemplate {
	s.InitialCount = &v
	return s
}

func (s *QueueTemplate) SetInterConnect(v string) *QueueTemplate {
	s.InterConnect = &v
	return s
}

func (s *QueueTemplate) SetKeepAliveNodes(v []*string) *QueueTemplate {
	s.KeepAliveNodes = v
	return s
}

func (s *QueueTemplate) SetMaxCount(v int32) *QueueTemplate {
	s.MaxCount = &v
	return s
}

func (s *QueueTemplate) SetMaxCountPerCycle(v int64) *QueueTemplate {
	s.MaxCountPerCycle = &v
	return s
}

func (s *QueueTemplate) SetMinCount(v int32) *QueueTemplate {
	s.MinCount = &v
	return s
}

func (s *QueueTemplate) SetQueueName(v string) *QueueTemplate {
	s.QueueName = &v
	return s
}

func (s *QueueTemplate) SetRamRole(v string) *QueueTemplate {
	s.RamRole = &v
	return s
}

func (s *QueueTemplate) SetReservedNodePoolId(v string) *QueueTemplate {
	s.ReservedNodePoolId = &v
	return s
}

func (s *QueueTemplate) SetVSwitchIds(v []*string) *QueueTemplate {
	s.VSwitchIds = v
	return s
}

type SharedStorageTemplate struct {
	// example:
	//
	// 008b63****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// /home
	MountDirectory *string `json:"MountDirectory,omitempty" xml:"MountDirectory,omitempty"`
	// example:
	//
	// -t nfs -o vers=3,nolock,noresvport
	MountOptions *string `json:"MountOptions,omitempty" xml:"MountOptions,omitempty"`
	// example:
	//
	// 008b****-sihc.cn-hangzhou.extreme.nas.aliyuncs.com
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// example:
	//
	// /
	NASDirectory *string `json:"NASDirectory,omitempty" xml:"NASDirectory,omitempty"`
	// example:
	//
	// NFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
}

func (s SharedStorageTemplate) String() string {
	return tea.Prettify(s)
}

func (s SharedStorageTemplate) GoString() string {
	return s.String()
}

func (s *SharedStorageTemplate) SetFileSystemId(v string) *SharedStorageTemplate {
	s.FileSystemId = &v
	return s
}

func (s *SharedStorageTemplate) SetMountDirectory(v string) *SharedStorageTemplate {
	s.MountDirectory = &v
	return s
}

func (s *SharedStorageTemplate) SetMountOptions(v string) *SharedStorageTemplate {
	s.MountOptions = &v
	return s
}

func (s *SharedStorageTemplate) SetMountTargetDomain(v string) *SharedStorageTemplate {
	s.MountTargetDomain = &v
	return s
}

func (s *SharedStorageTemplate) SetNASDirectory(v string) *SharedStorageTemplate {
	s.NASDirectory = &v
	return s
}

func (s *SharedStorageTemplate) SetProtocolType(v string) *SharedStorageTemplate {
	s.ProtocolType = &v
	return s
}

type AttachNodesRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-xxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The compute node information.
	//
	// This parameter is required.
	ComputeNode *AttachNodesRequestComputeNode `json:"ComputeNode,omitempty" xml:"ComputeNode,omitempty" type:"Struct"`
	// The name of the queue to which the instance is to be added.
	//
	// example:
	//
	// comp
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s AttachNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachNodesRequest) GoString() string {
	return s.String()
}

func (s *AttachNodesRequest) SetClusterId(v string) *AttachNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *AttachNodesRequest) SetComputeNode(v *AttachNodesRequestComputeNode) *AttachNodesRequest {
	s.ComputeNode = v
	return s
}

func (s *AttachNodesRequest) SetQueueName(v string) *AttachNodesRequest {
	s.QueueName = &v
	return s
}

type AttachNodesRequestComputeNode struct {
	// The image ID. This image will be used to replace the original system disk image.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-bp1h765oyqyxxxxxxxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance IDs.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s AttachNodesRequestComputeNode) String() string {
	return tea.Prettify(s)
}

func (s AttachNodesRequestComputeNode) GoString() string {
	return s.String()
}

func (s *AttachNodesRequestComputeNode) SetImageId(v string) *AttachNodesRequestComputeNode {
	s.ImageId = &v
	return s
}

func (s *AttachNodesRequestComputeNode) SetInstanceIds(v []*string) *AttachNodesRequestComputeNode {
	s.InstanceIds = v
	return s
}

type AttachNodesShrinkRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-xxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The compute node information.
	//
	// This parameter is required.
	ComputeNodeShrink *string `json:"ComputeNode,omitempty" xml:"ComputeNode,omitempty"`
	// The name of the queue to which the instance is to be added.
	//
	// example:
	//
	// comp
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s AttachNodesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachNodesShrinkRequest) SetClusterId(v string) *AttachNodesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *AttachNodesShrinkRequest) SetComputeNodeShrink(v string) *AttachNodesShrinkRequest {
	s.ComputeNodeShrink = &v
	return s
}

func (s *AttachNodesShrinkRequest) SetQueueName(v string) *AttachNodesShrinkRequest {
	s.QueueName = &v
	return s
}

type AttachNodesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2263XXXX-XXXX-XXXX-XXXX-XXXX2448XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachNodesResponseBody) GoString() string {
	return s.String()
}

func (s *AttachNodesResponseBody) SetRequestId(v string) *AttachNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachNodesResponseBody) SetSuccess(v bool) *AttachNodesResponseBody {
	s.Success = &v
	return s
}

type AttachNodesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachNodesResponse) GoString() string {
	return s.String()
}

func (s *AttachNodesResponse) SetHeaders(v map[string]*string) *AttachNodesResponse {
	s.Headers = v
	return s
}

func (s *AttachNodesResponse) SetStatusCode(v int32) *AttachNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachNodesResponse) SetBody(v *AttachNodesResponseBody) *AttachNodesResponse {
	s.Body = v
	return s
}

type AttachSharedStoragesRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about the shared storage resources that you want to attach to the cluster.
	//
	// This parameter is required.
	SharedStorages []*AttachSharedStoragesRequestSharedStorages `json:"SharedStorages,omitempty" xml:"SharedStorages,omitempty" type:"Repeated"`
}

func (s AttachSharedStoragesRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachSharedStoragesRequest) GoString() string {
	return s.String()
}

func (s *AttachSharedStoragesRequest) SetClusterId(v string) *AttachSharedStoragesRequest {
	s.ClusterId = &v
	return s
}

func (s *AttachSharedStoragesRequest) SetSharedStorages(v []*AttachSharedStoragesRequestSharedStorages) *AttachSharedStoragesRequest {
	s.SharedStorages = v
	return s
}

type AttachSharedStoragesRequestSharedStorages struct {
	// The ID of the file system to be attached.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0bd504b0**
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The storage location of the file system to be attached. Valid values:
	//
	// 	- OnPremise: The file system is deployed on a hybrid cloud.
	//
	// 	- PublicCloud: The file system is deployed on a public cloud.
	//
	// example:
	//
	// PublicCloud
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The local mount directory of the file system that you want to attach.
	//
	// This parameter is required.
	//
	// example:
	//
	// /test
	MountDirectory *string `json:"MountDirectory,omitempty" xml:"MountDirectory,omitempty"`
	// The attaching options of the file system to be attached. Valid values:
	//
	// 	- \\-t nfs -o vers=3,nolock,proto=tcp,noresvport
	//
	// 	- \\-t nfs -o vers=4.0,noresvport
	//
	// Default value:-t nfs -o vers=3,nolock,proto=tcp,noresvport
	//
	// >  The v3 version is recommended for higher performance if multiple Elastic Compute Service (ECS) instances do not edit the same file at the same time.
	//
	// example:
	//
	// -t nfs -o vers=3,nolock,proto=tcp,noresvport
	MountOptions *string `json:"MountOptions,omitempty" xml:"MountOptions,omitempty"`
	// The address of the mount point of the file system to be attached.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0bd504b***-ngq26.cn-hangzhou.nas.aliyuncs.com
	MountTarget *string `json:"MountTarget,omitempty" xml:"MountTarget,omitempty"`
	// The protocol type of the file system to be attached. Valid values:
	//
	// 	- NFS
	//
	// 	- SMB
	//
	// This parameter is required.
	//
	// example:
	//
	// NFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The storage directory of the file system. You can mount any directory in the file system to the specified cluster directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// /testehpc
	StorageDirectory *string `json:"StorageDirectory,omitempty" xml:"StorageDirectory,omitempty"`
	// The type of the file system to be attached. Valid values:
	//
	// 	- nas
	//
	// 	- cpfs
	//
	// This parameter is required.
	//
	// example:
	//
	// nas
	VolumeType *string `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
}

func (s AttachSharedStoragesRequestSharedStorages) String() string {
	return tea.Prettify(s)
}

func (s AttachSharedStoragesRequestSharedStorages) GoString() string {
	return s.String()
}

func (s *AttachSharedStoragesRequestSharedStorages) SetFileSystemId(v string) *AttachSharedStoragesRequestSharedStorages {
	s.FileSystemId = &v
	return s
}

func (s *AttachSharedStoragesRequestSharedStorages) SetLocation(v string) *AttachSharedStoragesRequestSharedStorages {
	s.Location = &v
	return s
}

func (s *AttachSharedStoragesRequestSharedStorages) SetMountDirectory(v string) *AttachSharedStoragesRequestSharedStorages {
	s.MountDirectory = &v
	return s
}

func (s *AttachSharedStoragesRequestSharedStorages) SetMountOptions(v string) *AttachSharedStoragesRequestSharedStorages {
	s.MountOptions = &v
	return s
}

func (s *AttachSharedStoragesRequestSharedStorages) SetMountTarget(v string) *AttachSharedStoragesRequestSharedStorages {
	s.MountTarget = &v
	return s
}

func (s *AttachSharedStoragesRequestSharedStorages) SetProtocolType(v string) *AttachSharedStoragesRequestSharedStorages {
	s.ProtocolType = &v
	return s
}

func (s *AttachSharedStoragesRequestSharedStorages) SetStorageDirectory(v string) *AttachSharedStoragesRequestSharedStorages {
	s.StorageDirectory = &v
	return s
}

func (s *AttachSharedStoragesRequestSharedStorages) SetVolumeType(v string) *AttachSharedStoragesRequestSharedStorages {
	s.VolumeType = &v
	return s
}

type AttachSharedStoragesShrinkRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about the shared storage resources that you want to attach to the cluster.
	//
	// This parameter is required.
	SharedStoragesShrink *string `json:"SharedStorages,omitempty" xml:"SharedStorages,omitempty"`
}

func (s AttachSharedStoragesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachSharedStoragesShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachSharedStoragesShrinkRequest) SetClusterId(v string) *AttachSharedStoragesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *AttachSharedStoragesShrinkRequest) SetSharedStoragesShrink(v string) *AttachSharedStoragesShrinkRequest {
	s.SharedStoragesShrink = &v
	return s
}

type AttachSharedStoragesResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F9B7BEF8-E42E-5090-9880-55FB7872****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachSharedStoragesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachSharedStoragesResponseBody) GoString() string {
	return s.String()
}

func (s *AttachSharedStoragesResponseBody) SetClusterId(v string) *AttachSharedStoragesResponseBody {
	s.ClusterId = &v
	return s
}

func (s *AttachSharedStoragesResponseBody) SetRequestId(v string) *AttachSharedStoragesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachSharedStoragesResponseBody) SetSuccess(v string) *AttachSharedStoragesResponseBody {
	s.Success = &v
	return s
}

type AttachSharedStoragesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachSharedStoragesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachSharedStoragesResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachSharedStoragesResponse) GoString() string {
	return s.String()
}

func (s *AttachSharedStoragesResponse) SetHeaders(v map[string]*string) *AttachSharedStoragesResponse {
	s.Headers = v
	return s
}

func (s *AttachSharedStoragesResponse) SetStatusCode(v int32) *AttachSharedStoragesResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachSharedStoragesResponse) SetBody(v *AttachSharedStoragesResponseBody) *AttachSharedStoragesResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	// The list of software that you want to install in the cluster. Valid values of N: 0 to 10.
	AdditionalPackages []*CreateClusterRequestAdditionalPackages `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty" type:"Repeated"`
	// The configurations of the custom addons in the cluster. Only one addon is supported.
	Addons []*CreateClusterRequestAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
	// The client version. By default, the latest version is used.
	//
	// example:
	//
	// 2.1.0
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The cluster type. Valid values:
	//
	// 	- Standard
	//
	// 	- Serverless
	//
	// example:
	//
	// Standard
	ClusterCategory *string `json:"ClusterCategory,omitempty" xml:"ClusterCategory,omitempty"`
	// The access credentials of the cluster.
	ClusterCredentials *CreateClusterRequestClusterCredentials `json:"ClusterCredentials,omitempty" xml:"ClusterCredentials,omitempty" type:"Struct"`
	// The post-processing script of the cluster.
	ClusterCustomConfiguration *CreateClusterRequestClusterCustomConfiguration `json:"ClusterCustomConfiguration,omitempty" xml:"ClusterCustomConfiguration,omitempty" type:"Struct"`
	// The cluster description. The description must be 1 to 128 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// slurm22.05.8-cluster-20240718
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// The deployment mode of the cluster. Valid values:
	//
	// 	- Integrated
	//
	// 	- Hybrid
	//
	// 	- Custom
	//
	// example:
	//
	// Integrated
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// The cluster name. The name must be 1 to 128 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// slurm22.05.8-cluster-20240718
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The ID of the vSwitch that you want the cluster to use. The vSwitch must reside in the VPC that is specified by the `ClusterVpcId` parameter.
	//
	// You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/448581.html) operation to query information about the created VPCs and vSwitches.
	//
	// example:
	//
	// vsw-f8za5p0mwzgdu3wgx****
	ClusterVSwitchId *string `json:"ClusterVSwitchId,omitempty" xml:"ClusterVSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the cluster resides.
	//
	// example:
	//
	// vpc-m5efjevmclc0xdmys****
	ClusterVpcId *string `json:"ClusterVpcId,omitempty" xml:"ClusterVpcId,omitempty"`
	// Specifies whether to enable deletion protection for the cluster. Deletion protection decides whether the cluster can be deleted in the console or by calling the [DeleteCluster](https://help.aliyun.com/document_detail/424406.html) operation. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// Specifies whether to use an advanced security group. Valid values:
	//
	// 	- true: automatically creates and uses an advanced security group.
	//
	// 	- false: automatically creates and uses a basic security group.
	//
	// For more information, see [Basic security groups and advanced security groups](https://help.aliyun.com/document_detail/605897.html).
	//
	// example:
	//
	// false
	IsEnterpriseSecurityGroup *bool `json:"IsEnterpriseSecurityGroup,omitempty" xml:"IsEnterpriseSecurityGroup,omitempty"`
	// The configurations of the cluster management node.
	Manager *CreateClusterRequestManager `json:"Manager,omitempty" xml:"Manager,omitempty" type:"Struct"`
	// The maximum number of vCPUs that can be used by compute nodes in the cluster. Valid values: 0 to 100,000.
	//
	// example:
	//
	// 10000
	MaxCoreCount *int32 `json:"MaxCoreCount,omitempty" xml:"MaxCoreCount,omitempty"`
	// The maximum number of compute nodes that the cluster can manage. Valid values: 0 to 5,000.
	//
	// example:
	//
	// 500
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The queues in the cluster. The number of queues can be 0 to 8.
	Queues []*QueueTemplate `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
	// The ID of the resource group to which the cluster belongs.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to obtain the IDs of the resource groups.
	//
	// example:
	//
	// rg-acfmxazb4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the security group to which the cluster belongs.
	//
	// You can call the [DescribeSecurityGroups](https://help.aliyun.com/document_detail/25556.html) operation to query available security groups in the current region.
	//
	// example:
	//
	// sg-bp13n61xsydodfyg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The shared storage resources of the cluster.
	SharedStorages []*SharedStorageTemplate `json:"SharedStorages,omitempty" xml:"SharedStorages,omitempty" type:"Repeated"`
	// The tags of the cluster.
	Tags []*CreateClusterRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetAdditionalPackages(v []*CreateClusterRequestAdditionalPackages) *CreateClusterRequest {
	s.AdditionalPackages = v
	return s
}

func (s *CreateClusterRequest) SetAddons(v []*CreateClusterRequestAddons) *CreateClusterRequest {
	s.Addons = v
	return s
}

func (s *CreateClusterRequest) SetClientVersion(v string) *CreateClusterRequest {
	s.ClientVersion = &v
	return s
}

func (s *CreateClusterRequest) SetClusterCategory(v string) *CreateClusterRequest {
	s.ClusterCategory = &v
	return s
}

func (s *CreateClusterRequest) SetClusterCredentials(v *CreateClusterRequestClusterCredentials) *CreateClusterRequest {
	s.ClusterCredentials = v
	return s
}

func (s *CreateClusterRequest) SetClusterCustomConfiguration(v *CreateClusterRequestClusterCustomConfiguration) *CreateClusterRequest {
	s.ClusterCustomConfiguration = v
	return s
}

func (s *CreateClusterRequest) SetClusterDescription(v string) *CreateClusterRequest {
	s.ClusterDescription = &v
	return s
}

func (s *CreateClusterRequest) SetClusterMode(v string) *CreateClusterRequest {
	s.ClusterMode = &v
	return s
}

func (s *CreateClusterRequest) SetClusterName(v string) *CreateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterRequest) SetClusterVSwitchId(v string) *CreateClusterRequest {
	s.ClusterVSwitchId = &v
	return s
}

func (s *CreateClusterRequest) SetClusterVpcId(v string) *CreateClusterRequest {
	s.ClusterVpcId = &v
	return s
}

func (s *CreateClusterRequest) SetDeletionProtection(v bool) *CreateClusterRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateClusterRequest) SetIsEnterpriseSecurityGroup(v bool) *CreateClusterRequest {
	s.IsEnterpriseSecurityGroup = &v
	return s
}

func (s *CreateClusterRequest) SetManager(v *CreateClusterRequestManager) *CreateClusterRequest {
	s.Manager = v
	return s
}

func (s *CreateClusterRequest) SetMaxCoreCount(v int32) *CreateClusterRequest {
	s.MaxCoreCount = &v
	return s
}

func (s *CreateClusterRequest) SetMaxCount(v int32) *CreateClusterRequest {
	s.MaxCount = &v
	return s
}

func (s *CreateClusterRequest) SetQueues(v []*QueueTemplate) *CreateClusterRequest {
	s.Queues = v
	return s
}

func (s *CreateClusterRequest) SetResourceGroupId(v string) *CreateClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityGroupId(v string) *CreateClusterRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetSharedStorages(v []*SharedStorageTemplate) *CreateClusterRequest {
	s.SharedStorages = v
	return s
}

func (s *CreateClusterRequest) SetTags(v []*CreateClusterRequestTags) *CreateClusterRequest {
	s.Tags = v
	return s
}

type CreateClusterRequestAdditionalPackages struct {
	// The name of the software that you want to install in the cluster.
	//
	// example:
	//
	// mpich
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the software that you want to install in the cluster.
	//
	// example:
	//
	// 4.0.3
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateClusterRequestAdditionalPackages) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestAdditionalPackages) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestAdditionalPackages) SetName(v string) *CreateClusterRequestAdditionalPackages {
	s.Name = &v
	return s
}

func (s *CreateClusterRequestAdditionalPackages) SetVersion(v string) *CreateClusterRequestAdditionalPackages {
	s.Version = &v
	return s
}

type CreateClusterRequestAddons struct {
	// The addon name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource configurations of the addon.
	//
	// example:
	//
	// "{\\\\"EipResource\\\\": {\\\\"AutoCreate\\\\": true}, \\\\"EcsResources\\\\": [{\\\\"InstanceType\\\\": \\\\"ecs.c7.xlarge\\\\", \\\\"ImageId\\\\": \\\\"centos_7_6_x64_20G_alibase_20211130.vhd\\\\", \\\\"SystemDisk\\\\": {\\\\"Category\\\\": \\\\"cloud_essd\\\\", \\\\"Size\\\\": 40, \\\\"Level\\\\": \\\\"PL0\\\\"}, \\\\"EnableHT\\\\": true, \\\\"InstanceChargeType\\\\": \\\\"PostPaid\\\\", \\\\"SpotStrategy\\\\": \\\\"NoSpot\\\\"}]}"
	ResourcesSpec *string `json:"ResourcesSpec,omitempty" xml:"ResourcesSpec,omitempty"`
	// The service configurations of the addon.
	//
	// example:
	//
	// "[{\\\\"ServiceName\\\\": \\\\"SSH\\\\", \\\\"ServiceAccessType\\\\": null, \\\\"ServiceAccessUrl\\\\": null, \\\\"NetworkACL\\\\": [{\\\\"IpProtocol\\\\": \\\\"TCP\\\\", \\\\"Port\\\\": 22, \\\\"SourceCidrIp\\\\": \\\\"0.0.0.0/0\\\\"}]}, {\\\\"ServiceName\\\\": \\\\"VNC\\\\", \\\\"ServiceAccessType\\\\": null, \\\\"ServiceAccessUrl\\\\": null, \\\\"NetworkACL\\\\": [{\\\\"IpProtocol\\\\": \\\\"TCP\\\\", \\\\"Port\\\\": 12016, \\\\"SourceCidrIp\\\\": \\\\"0.0.0.0/0\\\\"}]}, {\\\\"ServiceName\\\\": \\\\"CLIENT\\\\", \\\\"ServiceAccessType\\\\": \\\\"URL\\\\", \\\\"ServiceAccessUrl\\\\": \\\\"\\\\", \\\\"NetworkACL\\\\": [{\\\\"IpProtocol\\\\": \\\\"TCP\\\\", \\\\"Port\\\\": 12011, \\\\"SourceCidrIp\\\\": \\\\"0.0.0.0/0\\\\"}]}]"
	ServicesSpec *string `json:"ServicesSpec,omitempty" xml:"ServicesSpec,omitempty"`
	// The addon version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateClusterRequestAddons) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestAddons) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestAddons) SetName(v string) *CreateClusterRequestAddons {
	s.Name = &v
	return s
}

func (s *CreateClusterRequestAddons) SetResourcesSpec(v string) *CreateClusterRequestAddons {
	s.ResourcesSpec = &v
	return s
}

func (s *CreateClusterRequestAddons) SetServicesSpec(v string) *CreateClusterRequestAddons {
	s.ServicesSpec = &v
	return s
}

func (s *CreateClusterRequestAddons) SetVersion(v string) *CreateClusterRequestAddons {
	s.Version = &v
	return s
}

type CreateClusterRequestClusterCredentials struct {
	// The name of the key pair. The name must be 2 to 128 characters in length. The name must start with a letter but cannot start with `http://` or `https://`. The name can contain digits, letters, colons (:), underscores (_), and hyphens (-).
	//
	// >  For more information, see [Create a key pair](https://help.aliyun.com/document_detail/51793.html).
	//
	// example:
	//
	// ali0824
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The password for the root user to log on to the node. The password must be 8 to 20 characters in length, and must contain at least 3 of the following character types: uppercase letters, lowercase letters, digits, and special characters. The following special characters are supported: `() ~ ! @ # $ % ^ & 	- - = + { } [ ] : ; \\" < > , . ? /`
	//
	// >  We recommend that you use HTTPS to call the API operation to prevent password leakage.
	//
	// example:
	//
	// **********
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s CreateClusterRequestClusterCredentials) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestClusterCredentials) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestClusterCredentials) SetKeyPairName(v string) *CreateClusterRequestClusterCredentials {
	s.KeyPairName = &v
	return s
}

func (s *CreateClusterRequestClusterCredentials) SetPassword(v string) *CreateClusterRequestClusterCredentials {
	s.Password = &v
	return s
}

type CreateClusterRequestClusterCustomConfiguration struct {
	// The runtime parameters of the script after the cluster is created.
	//
	// example:
	//
	// E-HPC cn-hangzhou
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The URL that is used to download the post-processing script.
	//
	// example:
	//
	// http://*****
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s CreateClusterRequestClusterCustomConfiguration) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestClusterCustomConfiguration) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestClusterCustomConfiguration) SetArgs(v string) *CreateClusterRequestClusterCustomConfiguration {
	s.Args = &v
	return s
}

func (s *CreateClusterRequestClusterCustomConfiguration) SetScript(v string) *CreateClusterRequestClusterCustomConfiguration {
	s.Script = &v
	return s
}

type CreateClusterRequestManager struct {
	// The configurations of the domain name resolution service.
	DNS *CreateClusterRequestManagerDNS `json:"DNS,omitempty" xml:"DNS,omitempty" type:"Struct"`
	// The configurations of the domain account service.
	DirectoryService *CreateClusterRequestManagerDirectoryService `json:"DirectoryService,omitempty" xml:"DirectoryService,omitempty" type:"Struct"`
	// The hardware configurations of the management node.
	ManagerNode *NodeTemplate `json:"ManagerNode,omitempty" xml:"ManagerNode,omitempty"`
	// The configurations of the scheduler service.
	Scheduler *CreateClusterRequestManagerScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Struct"`
}

func (s CreateClusterRequestManager) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestManager) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestManager) SetDNS(v *CreateClusterRequestManagerDNS) *CreateClusterRequestManager {
	s.DNS = v
	return s
}

func (s *CreateClusterRequestManager) SetDirectoryService(v *CreateClusterRequestManagerDirectoryService) *CreateClusterRequestManager {
	s.DirectoryService = v
	return s
}

func (s *CreateClusterRequestManager) SetManagerNode(v *NodeTemplate) *CreateClusterRequestManager {
	s.ManagerNode = v
	return s
}

func (s *CreateClusterRequestManager) SetScheduler(v *CreateClusterRequestManagerScheduler) *CreateClusterRequestManager {
	s.Scheduler = v
	return s
}

type CreateClusterRequestManagerDNS struct {
	// The domain name resolution type.
	//
	// Valid value:
	//
	// 	- NIS
	//
	// example:
	//
	// NIS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the domain name resolution service.
	//
	// example:
	//
	// 2.31
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateClusterRequestManagerDNS) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestManagerDNS) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestManagerDNS) SetType(v string) *CreateClusterRequestManagerDNS {
	s.Type = &v
	return s
}

func (s *CreateClusterRequestManagerDNS) SetVersion(v string) *CreateClusterRequestManagerDNS {
	s.Version = &v
	return s
}

type CreateClusterRequestManagerDirectoryService struct {
	// The type of the domain account.
	//
	// Valid value:
	//
	// 	- NIS
	//
	// example:
	//
	// NIS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the domain account service.
	//
	// example:
	//
	// 2.31
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateClusterRequestManagerDirectoryService) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestManagerDirectoryService) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestManagerDirectoryService) SetType(v string) *CreateClusterRequestManagerDirectoryService {
	s.Type = &v
	return s
}

func (s *CreateClusterRequestManagerDirectoryService) SetVersion(v string) *CreateClusterRequestManagerDirectoryService {
	s.Version = &v
	return s
}

type CreateClusterRequestManagerScheduler struct {
	// The scheduler type. Valid values:
	//
	// 	- SLURM
	//
	// 	- PBS
	//
	// 	- OPENGRIDSCHEDULER
	//
	// 	- LSF_PLUGIN
	//
	// 	- PBS_PLUGIN
	//
	// example:
	//
	// SLURM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The scheduler version.
	//
	// example:
	//
	// 22.05.8
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateClusterRequestManagerScheduler) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestManagerScheduler) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestManagerScheduler) SetType(v string) *CreateClusterRequestManagerScheduler {
	s.Type = &v
	return s
}

func (s *CreateClusterRequestManagerScheduler) SetVersion(v string) *CreateClusterRequestManagerScheduler {
	s.Version = &v
	return s
}

type CreateClusterRequestTags struct {
	// The tag key. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// ClusterId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// ehpc-hz-******
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateClusterRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestTags) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestTags) SetKey(v string) *CreateClusterRequestTags {
	s.Key = &v
	return s
}

func (s *CreateClusterRequestTags) SetValue(v string) *CreateClusterRequestTags {
	s.Value = &v
	return s
}

type CreateClusterShrinkRequest struct {
	// The list of software that you want to install in the cluster. Valid values of N: 0 to 10.
	AdditionalPackagesShrink *string `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty"`
	// The configurations of the custom addons in the cluster. Only one addon is supported.
	AddonsShrink *string `json:"Addons,omitempty" xml:"Addons,omitempty"`
	// The client version. By default, the latest version is used.
	//
	// example:
	//
	// 2.1.0
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The cluster type. Valid values:
	//
	// 	- Standard
	//
	// 	- Serverless
	//
	// example:
	//
	// Standard
	ClusterCategory *string `json:"ClusterCategory,omitempty" xml:"ClusterCategory,omitempty"`
	// The access credentials of the cluster.
	ClusterCredentialsShrink *string `json:"ClusterCredentials,omitempty" xml:"ClusterCredentials,omitempty"`
	// The post-processing script of the cluster.
	ClusterCustomConfigurationShrink *string `json:"ClusterCustomConfiguration,omitempty" xml:"ClusterCustomConfiguration,omitempty"`
	// The cluster description. The description must be 1 to 128 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// slurm22.05.8-cluster-20240718
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// The deployment mode of the cluster. Valid values:
	//
	// 	- Integrated
	//
	// 	- Hybrid
	//
	// 	- Custom
	//
	// example:
	//
	// Integrated
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// The cluster name. The name must be 1 to 128 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// slurm22.05.8-cluster-20240718
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The ID of the vSwitch that you want the cluster to use. The vSwitch must reside in the VPC that is specified by the `ClusterVpcId` parameter.
	//
	// You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/448581.html) operation to query information about the created VPCs and vSwitches.
	//
	// example:
	//
	// vsw-f8za5p0mwzgdu3wgx****
	ClusterVSwitchId *string `json:"ClusterVSwitchId,omitempty" xml:"ClusterVSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the cluster resides.
	//
	// example:
	//
	// vpc-m5efjevmclc0xdmys****
	ClusterVpcId *string `json:"ClusterVpcId,omitempty" xml:"ClusterVpcId,omitempty"`
	// Specifies whether to enable deletion protection for the cluster. Deletion protection decides whether the cluster can be deleted in the console or by calling the [DeleteCluster](https://help.aliyun.com/document_detail/424406.html) operation. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// Specifies whether to use an advanced security group. Valid values:
	//
	// 	- true: automatically creates and uses an advanced security group.
	//
	// 	- false: automatically creates and uses a basic security group.
	//
	// For more information, see [Basic security groups and advanced security groups](https://help.aliyun.com/document_detail/605897.html).
	//
	// example:
	//
	// false
	IsEnterpriseSecurityGroup *bool `json:"IsEnterpriseSecurityGroup,omitempty" xml:"IsEnterpriseSecurityGroup,omitempty"`
	// The configurations of the cluster management node.
	ManagerShrink *string `json:"Manager,omitempty" xml:"Manager,omitempty"`
	// The maximum number of vCPUs that can be used by compute nodes in the cluster. Valid values: 0 to 100,000.
	//
	// example:
	//
	// 10000
	MaxCoreCount *int32 `json:"MaxCoreCount,omitempty" xml:"MaxCoreCount,omitempty"`
	// The maximum number of compute nodes that the cluster can manage. Valid values: 0 to 5,000.
	//
	// example:
	//
	// 500
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The queues in the cluster. The number of queues can be 0 to 8.
	QueuesShrink *string `json:"Queues,omitempty" xml:"Queues,omitempty"`
	// The ID of the resource group to which the cluster belongs.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to obtain the IDs of the resource groups.
	//
	// example:
	//
	// rg-acfmxazb4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the security group to which the cluster belongs.
	//
	// You can call the [DescribeSecurityGroups](https://help.aliyun.com/document_detail/25556.html) operation to query available security groups in the current region.
	//
	// example:
	//
	// sg-bp13n61xsydodfyg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The shared storage resources of the cluster.
	SharedStoragesShrink *string `json:"SharedStorages,omitempty" xml:"SharedStorages,omitempty"`
	// The tags of the cluster.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterShrinkRequest) SetAdditionalPackagesShrink(v string) *CreateClusterShrinkRequest {
	s.AdditionalPackagesShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetAddonsShrink(v string) *CreateClusterShrinkRequest {
	s.AddonsShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClientVersion(v string) *CreateClusterShrinkRequest {
	s.ClientVersion = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterCategory(v string) *CreateClusterShrinkRequest {
	s.ClusterCategory = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterCredentialsShrink(v string) *CreateClusterShrinkRequest {
	s.ClusterCredentialsShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterCustomConfigurationShrink(v string) *CreateClusterShrinkRequest {
	s.ClusterCustomConfigurationShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterDescription(v string) *CreateClusterShrinkRequest {
	s.ClusterDescription = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterMode(v string) *CreateClusterShrinkRequest {
	s.ClusterMode = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterName(v string) *CreateClusterShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterVSwitchId(v string) *CreateClusterShrinkRequest {
	s.ClusterVSwitchId = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterVpcId(v string) *CreateClusterShrinkRequest {
	s.ClusterVpcId = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetDeletionProtection(v bool) *CreateClusterShrinkRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetIsEnterpriseSecurityGroup(v bool) *CreateClusterShrinkRequest {
	s.IsEnterpriseSecurityGroup = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetManagerShrink(v string) *CreateClusterShrinkRequest {
	s.ManagerShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetMaxCoreCount(v int32) *CreateClusterShrinkRequest {
	s.MaxCoreCount = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetMaxCount(v int32) *CreateClusterShrinkRequest {
	s.MaxCount = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetQueuesShrink(v string) *CreateClusterShrinkRequest {
	s.QueuesShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetResourceGroupId(v string) *CreateClusterShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetSecurityGroupId(v string) *CreateClusterShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetSharedStoragesShrink(v string) *CreateClusterShrinkRequest {
	s.SharedStoragesShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetTagsShrink(v string) *CreateClusterShrinkRequest {
	s.TagsShrink = &v
	return s
}

type CreateClusterResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F6757FA4-8FED-4602-B7F5-3550C0******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID.
	//
	// example:
	//
	// F6757FA4-8FED-4602-B7F5-3550C0******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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

func (s *CreateClusterResponseBody) SetSuccess(v bool) *CreateClusterResponseBody {
	s.Success = &v
	return s
}

func (s *CreateClusterResponseBody) SetTaskId(v string) *CreateClusterResponseBody {
	s.TaskId = &v
	return s
}

type CreateClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateClusterResponse) SetStatusCode(v int32) *CreateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateJobRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The job name.
	//
	// example:
	//
	// TestJob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The job configurations.
	JobSpec *CreateJobRequestJobSpec `json:"JobSpec,omitempty" xml:"JobSpec,omitempty" type:"Struct"`
}

func (s CreateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) SetClusterId(v string) *CreateJobRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateJobRequest) SetJobName(v string) *CreateJobRequest {
	s.JobName = &v
	return s
}

func (s *CreateJobRequest) SetJobSpec(v *CreateJobRequestJobSpec) *CreateJobRequest {
	s.JobSpec = v
	return s
}

type CreateJobRequestJobSpec struct {
	// The jobs in the queue.
	//
	// Format: X-Y:Z. X is the minimum index value. Y is the maximum index value. Z is the step size. For example, 2-7:2 indicates that three jobs need to be run and their index values are 2, 4, and 6.
	//
	// example:
	//
	// 1-5:2
	ArrayRequest *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	// The command or script that is used to run the job. If you want to use a command, you must specify the full path of the command, for example, /bin/ping.
	//
	// If you want to use a script, you must make sure that you have the execution permissions on it. By default, the user root directory ~/ is used as the default script path on the cluster side. If your script is not in that directory, you must specify the full path in this parameter, such as /home/xxx/job.sh Note that in this mode, if requirements on resources such as CPU and memory are specified in the script, the job will be run based on the resource requirements specified in the script. In this case, do not specify resource requirements in the Resource parameter. Otherwise, the job may fail to run.
	//
	// If you want to run the job directly by using the CLI, you must specify the absolute path of the command and add two hyphens and a space (-- ) before the path, such as -- /bin/ping -c 10 localhost.
	//
	// This parameter is required.
	//
	// example:
	//
	// /home/xxx/test.job
	CommandLine *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	// The queue to which the job belongs.
	//
	// example:
	//
	// comp
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	// The post-processing command of the job.
	//
	// example:
	//
	// /bin/sleep 10
	PostCmdLine *string `json:"PostCmdLine,omitempty" xml:"PostCmdLine,omitempty"`
	// The job priority.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The resource configurations of the job.
	Resources *CreateJobRequestJobSpecResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// The cluster-side user as which you want to submit the job.
	//
	// example:
	//
	// testuser
	RunasUser *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	// The password of the user specified by the RunasUser parameter.
	//
	// example:
	//
	// xxx
	RunasUserPassword *string `json:"RunasUserPassword,omitempty" xml:"RunasUserPassword,omitempty"`
	// The path of the standard error output file of the job. You need to specify the full path.
	//
	// example:
	//
	// /home/xxx/job.err
	StderrPath *string `json:"StderrPath,omitempty" xml:"StderrPath,omitempty"`
	// The path of the standard output file of the job. You need to specify the full path.
	//
	// example:
	//
	// /home/xxx/job.out
	StdoutPath *string `json:"StdoutPath,omitempty" xml:"StdoutPath,omitempty"`
	// The environment variables of the job. The value is a string in the JSON array format. Each array member is a JSON object that contains two members: Name and Value. Name indicates the name of the environment variable, and Value indicates the value of the environment variable.
	//
	// example:
	//
	// [{"Name":"x", "Value":"y"}]
	Variables *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
	// The maximum duration for which the job can be run. Format: `hour: minute: second`. For example, `01:00:00` indicates 1 hour.
	//
	// example:
	//
	// 360:48:50
	WallTime *string `json:"WallTime,omitempty" xml:"WallTime,omitempty"`
}

func (s CreateJobRequestJobSpec) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestJobSpec) GoString() string {
	return s.String()
}

func (s *CreateJobRequestJobSpec) SetArrayRequest(v string) *CreateJobRequestJobSpec {
	s.ArrayRequest = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetCommandLine(v string) *CreateJobRequestJobSpec {
	s.CommandLine = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetJobQueue(v string) *CreateJobRequestJobSpec {
	s.JobQueue = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetPostCmdLine(v string) *CreateJobRequestJobSpec {
	s.PostCmdLine = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetPriority(v string) *CreateJobRequestJobSpec {
	s.Priority = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetResources(v *CreateJobRequestJobSpecResources) *CreateJobRequestJobSpec {
	s.Resources = v
	return s
}

func (s *CreateJobRequestJobSpec) SetRunasUser(v string) *CreateJobRequestJobSpec {
	s.RunasUser = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetRunasUserPassword(v string) *CreateJobRequestJobSpec {
	s.RunasUserPassword = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetStderrPath(v string) *CreateJobRequestJobSpec {
	s.StderrPath = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetStdoutPath(v string) *CreateJobRequestJobSpec {
	s.StdoutPath = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetVariables(v string) *CreateJobRequestJobSpec {
	s.Variables = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetWallTime(v string) *CreateJobRequestJobSpec {
	s.WallTime = &v
	return s
}

type CreateJobRequestJobSpecResources struct {
	// The number of vCPUs to be allocated to each compute node.
	//
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The number of GPUs to be allocated to each compute node.
	//
	// example:
	//
	// 1
	Gpus *int32 `json:"Gpus,omitempty" xml:"Gpus,omitempty"`
	// The memory size to be allocated to each compute node. The memory size is in string format. Unit: MB or GB.
	//
	// example:
	//
	// 4gb
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The number of compute nodes to be allocated to the job.
	//
	// example:
	//
	// 2
	Nodes *int32 `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s CreateJobRequestJobSpecResources) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestJobSpecResources) GoString() string {
	return s.String()
}

func (s *CreateJobRequestJobSpecResources) SetCores(v int32) *CreateJobRequestJobSpecResources {
	s.Cores = &v
	return s
}

func (s *CreateJobRequestJobSpecResources) SetGpus(v int32) *CreateJobRequestJobSpecResources {
	s.Gpus = &v
	return s
}

func (s *CreateJobRequestJobSpecResources) SetMemory(v string) *CreateJobRequestJobSpecResources {
	s.Memory = &v
	return s
}

func (s *CreateJobRequestJobSpecResources) SetNodes(v int32) *CreateJobRequestJobSpecResources {
	s.Nodes = &v
	return s
}

type CreateJobShrinkRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The job name.
	//
	// example:
	//
	// TestJob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The job configurations.
	JobSpecShrink *string `json:"JobSpec,omitempty" xml:"JobSpec,omitempty"`
}

func (s CreateJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateJobShrinkRequest) SetClusterId(v string) *CreateJobShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateJobShrinkRequest) SetJobName(v string) *CreateJobShrinkRequest {
	s.JobName = &v
	return s
}

func (s *CreateJobShrinkRequest) SetJobSpecShrink(v string) *CreateJobShrinkRequest {
	s.JobSpecShrink = &v
	return s
}

type CreateJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// Submitted batch job 10\\n
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A0A38A38-1565-555E-B597-E48A2E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBody) SetJobId(v string) *CreateJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateJobResponseBody) SetRequestId(v string) *CreateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobResponseBody) SetSuccess(v string) *CreateJobResponseBody {
	s.Success = &v
	return s
}

type CreateJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponse) GoString() string {
	return s.String()
}

func (s *CreateJobResponse) SetHeaders(v map[string]*string) *CreateJobResponse {
	s.Headers = v
	return s
}

func (s *CreateJobResponse) SetStatusCode(v int32) *CreateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobResponse) SetBody(v *CreateJobResponseBody) *CreateJobResponse {
	s.Body = v
	return s
}

type CreateNodesRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The hardware configurations of the compute nodes.
	ComputeNode *NodeTemplate `json:"ComputeNode,omitempty" xml:"ComputeNode,omitempty"`
	// The number of compute nodes that you want to add. Valid values: 1 to 99. The MinCount value must be smaller than the Count value.
	//
	// 	- If the number of available Elastic Compute Service (ECS) instances is smaller than the MinCount value, the nodes fail to be added.
	//
	// 	- If the number of available ECS instances is larger than the MinCount value but smaller than the Count value, nodes are added based on the MinCount value.
	//
	// 	- If the number of available ECS instances is larger than the Count value, nodes are added based on the Count value.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Deployment set ID. You can obtain the deployment set ID through [DescribeDeploymentSets](https://help.aliyun.com/document_detail/91313.html). Currently, only deployment sets with a low network latency strategy are supported.
	//
	// example:
	//
	// ds-bp1frxuzdg87zh4pzq****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The type of the network between compute nodes. Valid values:
	//
	// 	- vpc
	//
	// 	- eRDMA
	//
	// example:
	//
	// vpc
	HPCInterConnect *string `json:"HPCInterConnect,omitempty" xml:"HPCInterConnect,omitempty"`
	// The hostname prefix of the added compute nodes.
	//
	// example:
	//
	// compute
	HostnamePrefix *string `json:"HostnamePrefix,omitempty" xml:"HostnamePrefix,omitempty"`
	// The hostname suffix of the added compute nodes.
	//
	// example:
	//
	// demo
	HostnameSuffix *string `json:"HostnameSuffix,omitempty" xml:"HostnameSuffix,omitempty"`
	// Specifies whether to enable deletion protection for the added compute nodes.
	//
	// example:
	//
	// false
	KeepAlive *string `json:"KeepAlive,omitempty" xml:"KeepAlive,omitempty"`
	// The name of the queue for which you want to create compute nodes.
	//
	// example:
	//
	// test1
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The Resource Access Management (RAM) role to be assumed by the added nodes.
	//
	// example:
	//
	// AliyunServiceRoleForOOSBandwidthScheduler
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// Preset node pool ID.
	//
	// example:
	//
	// rnp-756vlp7a
	ReservedNodePoolId *string `json:"ReservedNodePoolId,omitempty" xml:"ReservedNodePoolId,omitempty"`
	// The ID of the vSwitch to be used by the added nodes.
	//
	// example:
	//
	// vsw-bp1lfcjbfb099rrjn****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNodesRequest) GoString() string {
	return s.String()
}

func (s *CreateNodesRequest) SetClusterId(v string) *CreateNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNodesRequest) SetComputeNode(v *NodeTemplate) *CreateNodesRequest {
	s.ComputeNode = v
	return s
}

func (s *CreateNodesRequest) SetCount(v int32) *CreateNodesRequest {
	s.Count = &v
	return s
}

func (s *CreateNodesRequest) SetDeploymentSetId(v string) *CreateNodesRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateNodesRequest) SetHPCInterConnect(v string) *CreateNodesRequest {
	s.HPCInterConnect = &v
	return s
}

func (s *CreateNodesRequest) SetHostnamePrefix(v string) *CreateNodesRequest {
	s.HostnamePrefix = &v
	return s
}

func (s *CreateNodesRequest) SetHostnameSuffix(v string) *CreateNodesRequest {
	s.HostnameSuffix = &v
	return s
}

func (s *CreateNodesRequest) SetKeepAlive(v string) *CreateNodesRequest {
	s.KeepAlive = &v
	return s
}

func (s *CreateNodesRequest) SetQueueName(v string) *CreateNodesRequest {
	s.QueueName = &v
	return s
}

func (s *CreateNodesRequest) SetRamRole(v string) *CreateNodesRequest {
	s.RamRole = &v
	return s
}

func (s *CreateNodesRequest) SetReservedNodePoolId(v string) *CreateNodesRequest {
	s.ReservedNodePoolId = &v
	return s
}

func (s *CreateNodesRequest) SetVSwitchId(v string) *CreateNodesRequest {
	s.VSwitchId = &v
	return s
}

type CreateNodesShrinkRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The hardware configurations of the compute nodes.
	ComputeNodeShrink *string `json:"ComputeNode,omitempty" xml:"ComputeNode,omitempty"`
	// The number of compute nodes that you want to add. Valid values: 1 to 99. The MinCount value must be smaller than the Count value.
	//
	// 	- If the number of available Elastic Compute Service (ECS) instances is smaller than the MinCount value, the nodes fail to be added.
	//
	// 	- If the number of available ECS instances is larger than the MinCount value but smaller than the Count value, nodes are added based on the MinCount value.
	//
	// 	- If the number of available ECS instances is larger than the Count value, nodes are added based on the Count value.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Deployment set ID. You can obtain the deployment set ID through [DescribeDeploymentSets](https://help.aliyun.com/document_detail/91313.html). Currently, only deployment sets with a low network latency strategy are supported.
	//
	// example:
	//
	// ds-bp1frxuzdg87zh4pzq****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The type of the network between compute nodes. Valid values:
	//
	// 	- vpc
	//
	// 	- eRDMA
	//
	// example:
	//
	// vpc
	HPCInterConnect *string `json:"HPCInterConnect,omitempty" xml:"HPCInterConnect,omitempty"`
	// The hostname prefix of the added compute nodes.
	//
	// example:
	//
	// compute
	HostnamePrefix *string `json:"HostnamePrefix,omitempty" xml:"HostnamePrefix,omitempty"`
	// The hostname suffix of the added compute nodes.
	//
	// example:
	//
	// demo
	HostnameSuffix *string `json:"HostnameSuffix,omitempty" xml:"HostnameSuffix,omitempty"`
	// Specifies whether to enable deletion protection for the added compute nodes.
	//
	// example:
	//
	// false
	KeepAlive *string `json:"KeepAlive,omitempty" xml:"KeepAlive,omitempty"`
	// The name of the queue for which you want to create compute nodes.
	//
	// example:
	//
	// test1
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The Resource Access Management (RAM) role to be assumed by the added nodes.
	//
	// example:
	//
	// AliyunServiceRoleForOOSBandwidthScheduler
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// Preset node pool ID.
	//
	// example:
	//
	// rnp-756vlp7a
	ReservedNodePoolId *string `json:"ReservedNodePoolId,omitempty" xml:"ReservedNodePoolId,omitempty"`
	// The ID of the vSwitch to be used by the added nodes.
	//
	// example:
	//
	// vsw-bp1lfcjbfb099rrjn****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateNodesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNodesShrinkRequest) SetClusterId(v string) *CreateNodesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetComputeNodeShrink(v string) *CreateNodesShrinkRequest {
	s.ComputeNodeShrink = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetCount(v int32) *CreateNodesShrinkRequest {
	s.Count = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetDeploymentSetId(v string) *CreateNodesShrinkRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetHPCInterConnect(v string) *CreateNodesShrinkRequest {
	s.HPCInterConnect = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetHostnamePrefix(v string) *CreateNodesShrinkRequest {
	s.HostnamePrefix = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetHostnameSuffix(v string) *CreateNodesShrinkRequest {
	s.HostnameSuffix = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetKeepAlive(v string) *CreateNodesShrinkRequest {
	s.KeepAlive = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetQueueName(v string) *CreateNodesShrinkRequest {
	s.QueueName = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetRamRole(v string) *CreateNodesShrinkRequest {
	s.RamRole = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetReservedNodePoolId(v string) *CreateNodesShrinkRequest {
	s.ReservedNodePoolId = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetVSwitchId(v string) *CreateNodesShrinkRequest {
	s.VSwitchId = &v
	return s
}

type CreateNodesResponseBody struct {
	// The IDs of the compute nodes to be created.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNodesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodesResponseBody) SetInstanceIds(v []*string) *CreateNodesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateNodesResponseBody) SetRequestId(v string) *CreateNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNodesResponseBody) SetSuccess(v bool) *CreateNodesResponseBody {
	s.Success = &v
	return s
}

type CreateNodesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNodesResponse) GoString() string {
	return s.String()
}

func (s *CreateNodesResponse) SetHeaders(v map[string]*string) *CreateNodesResponse {
	s.Headers = v
	return s
}

func (s *CreateNodesResponse) SetStatusCode(v int32) *CreateNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNodesResponse) SetBody(v *CreateNodesResponseBody) *CreateNodesResponse {
	s.Body = v
	return s
}

type CreateQueueRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The configurations of the queue to be created.
	Queue *QueueTemplate `json:"Queue,omitempty" xml:"Queue,omitempty"`
}

func (s CreateQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueRequest) GoString() string {
	return s.String()
}

func (s *CreateQueueRequest) SetClusterId(v string) *CreateQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateQueueRequest) SetQueue(v *QueueTemplate) *CreateQueueRequest {
	s.Queue = v
	return s
}

type CreateQueueShrinkRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The configurations of the queue to be created.
	QueueShrink *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
}

func (s CreateQueueShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateQueueShrinkRequest) SetClusterId(v string) *CreateQueueShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetQueueShrink(v string) *CreateQueueShrinkRequest {
	s.QueueShrink = &v
	return s
}

type CreateQueueResponseBody struct {
	// The name of the created queue.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQueueResponseBody) SetName(v string) *CreateQueueResponseBody {
	s.Name = &v
	return s
}

func (s *CreateQueueResponseBody) SetRequestId(v string) *CreateQueueResponseBody {
	s.RequestId = &v
	return s
}

type CreateQueueResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueResponse) GoString() string {
	return s.String()
}

func (s *CreateQueueResponse) SetHeaders(v map[string]*string) *CreateQueueResponse {
	s.Headers = v
	return s
}

func (s *CreateQueueResponse) SetStatusCode(v int32) *CreateQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQueueResponse) SetBody(v *CreateQueueResponseBody) *CreateQueueResponse {
	s.Body = v
	return s
}

type CreateUsersRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The users that you want to add.
	User []*CreateUsersRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s CreateUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersRequest) GoString() string {
	return s.String()
}

func (s *CreateUsersRequest) SetClusterId(v string) *CreateUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateUsersRequest) SetUser(v []*CreateUsersRequestUser) *CreateUsersRequest {
	s.User = v
	return s
}

type CreateUsersRequestUser struct {
	// The public key of the user.
	//
	// You can add up to 20 users in a call.
	//
	// Specify one of the Password and AuthKey parameters. The AuthKey parameter takes effect only when the cluster authentication method is set to Key. Key authentication is not recommended.
	//
	// example:
	//
	// Abc****
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The permission group to which the user belongs. Valid values:
	//
	// users: ordinary permissions, which are suitable for ordinary users that need only to submit and debug jobs. wheel: sudo permissions, which are suitable for administrators who need to manage clusters. In addition to submitting and debugging jobs, you can also run sudo commands to install software and restart nodes. You can add up to 20 users in a call.
	//
	// example:
	//
	// users
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The password of the user. The password must be 6 to 30 characters in length and must contain three of the following character types:
	//
	// 	- Uppercase letters
	//
	// 	- Lowercase letters
	//
	// 	- Digits
	//
	// 	- Special characters ()~!@#$%^&\\*-_+=|{}[]:;\\"/<>,.?/
	//
	// You can add up to 20 users in a call.
	//
	// Specify one of the Password and AuthKey parameters. The Password parameter takes effect only when the cluster authentication method is set to Password. Password authentication is recommended.
	//
	// example:
	//
	// 1@a2****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The username. The username must be 1 to 30 characters in length. It must start with a letter and can contain digits, letters, and periods (.).
	//
	// You can add up to 20 users in a call.
	//
	// example:
	//
	// testuser
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUsersRequestUser) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersRequestUser) GoString() string {
	return s.String()
}

func (s *CreateUsersRequestUser) SetAuthKey(v string) *CreateUsersRequestUser {
	s.AuthKey = &v
	return s
}

func (s *CreateUsersRequestUser) SetGroup(v string) *CreateUsersRequestUser {
	s.Group = &v
	return s
}

func (s *CreateUsersRequestUser) SetPassword(v string) *CreateUsersRequestUser {
	s.Password = &v
	return s
}

func (s *CreateUsersRequestUser) SetUserName(v string) *CreateUsersRequestUser {
	s.UserName = &v
	return s
}

type CreateUsersShrinkRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The users that you want to add.
	UserShrink *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s CreateUsersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUsersShrinkRequest) SetClusterId(v string) *CreateUsersShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateUsersShrinkRequest) SetUserShrink(v string) *CreateUsersShrinkRequest {
	s.UserShrink = &v
	return s
}

type CreateUsersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUsersResponseBody) SetRequestId(v string) *CreateUsersResponseBody {
	s.RequestId = &v
	return s
}

type CreateUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersResponse) GoString() string {
	return s.String()
}

func (s *CreateUsersResponse) SetHeaders(v map[string]*string) *CreateUsersResponse {
	s.Headers = v
	return s
}

func (s *CreateUsersResponse) SetStatusCode(v int32) *CreateUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUsersResponse) SetBody(v *CreateUsersResponseBody) *CreateUsersResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	// The ID of the E-HPC cluster to be released.
	//
	// example:
	//
	// ehpc-hz-QKKVqO****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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

type DeleteClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F1AB6D8D-E185-4D94-859C-7CE7B8B7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID.
	//
	// example:
	//
	// F1AB6D8D-E185-4D94-859C-7CE7B8B7****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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

func (s *DeleteClusterResponseBody) SetSuccess(v bool) *DeleteClusterResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteClusterResponseBody) SetTaskId(v string) *DeleteClusterResponseBody {
	s.TaskId = &v
	return s
}

type DeleteClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteClusterResponse) SetStatusCode(v int32) *DeleteClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type DeleteNodesRequest struct {
	// The cluster ID. You can call the [listclusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The instance IDs of the compute nodes that you want to delete.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
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

func (s *DeleteNodesRequest) SetInstanceIds(v []*string) *DeleteNodesRequest {
	s.InstanceIds = v
	return s
}

type DeleteNodesShrinkRequest struct {
	// The cluster ID. You can call the [listclusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The instance IDs of the compute nodes that you want to delete.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s DeleteNodesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodesShrinkRequest) SetClusterId(v string) *DeleteNodesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteNodesShrinkRequest) SetInstanceIdsShrink(v string) *DeleteNodesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

type DeleteNodesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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

func (s *DeleteNodesResponseBody) SetSuccess(v bool) *DeleteNodesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteNodesResponseBody) SetTaskId(v string) *DeleteNodesResponseBody {
	s.TaskId = &v
	return s
}

type DeleteNodesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteNodesResponse) SetStatusCode(v int32) *DeleteNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNodesResponse) SetBody(v *DeleteNodesResponseBody) *DeleteNodesResponse {
	s.Body = v
	return s
}

type DeleteQueuesRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The queues that you want to delete.
	QueueNames []*string `json:"QueueNames,omitempty" xml:"QueueNames,omitempty" type:"Repeated"`
}

func (s DeleteQueuesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueuesRequest) GoString() string {
	return s.String()
}

func (s *DeleteQueuesRequest) SetClusterId(v string) *DeleteQueuesRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteQueuesRequest) SetQueueNames(v []*string) *DeleteQueuesRequest {
	s.QueueNames = v
	return s
}

type DeleteQueuesShrinkRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The queues that you want to delete.
	QueueNamesShrink *string `json:"QueueNames,omitempty" xml:"QueueNames,omitempty"`
}

func (s DeleteQueuesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueuesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteQueuesShrinkRequest) SetClusterId(v string) *DeleteQueuesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteQueuesShrinkRequest) SetQueueNamesShrink(v string) *DeleteQueuesShrinkRequest {
	s.QueueNamesShrink = &v
	return s
}

type DeleteQueuesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQueuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueuesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQueuesResponseBody) SetRequestId(v string) *DeleteQueuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQueuesResponseBody) SetSuccess(v bool) *DeleteQueuesResponseBody {
	s.Success = &v
	return s
}

type DeleteQueuesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQueuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQueuesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueuesResponse) GoString() string {
	return s.String()
}

func (s *DeleteQueuesResponse) SetHeaders(v map[string]*string) *DeleteQueuesResponse {
	s.Headers = v
	return s
}

func (s *DeleteQueuesResponse) SetStatusCode(v int32) *DeleteQueuesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQueuesResponse) SetBody(v *DeleteQueuesResponseBody) *DeleteQueuesResponse {
	s.Body = v
	return s
}

type DeleteUsersRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The users that you want to delete.
	//
	// This parameter is required.
	User []*DeleteUsersRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
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
	// The name of user N that you want to delete.
	//
	// Valid values of N: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// testuser
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteUsersRequestUser) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersRequestUser) GoString() string {
	return s.String()
}

func (s *DeleteUsersRequestUser) SetUserName(v string) *DeleteUsersRequestUser {
	s.UserName = &v
	return s
}

type DeleteUsersShrinkRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The users that you want to delete.
	//
	// This parameter is required.
	UserShrink *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DeleteUsersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteUsersShrinkRequest) SetClusterId(v string) *DeleteUsersShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteUsersShrinkRequest) SetUserShrink(v string) *DeleteUsersShrinkRequest {
	s.UserShrink = &v
	return s
}

type DeleteUsersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE****
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteUsersResponse) SetStatusCode(v int32) *DeleteUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUsersResponse) SetBody(v *DeleteUsersResponseBody) *DeleteUsersResponse {
	s.Body = v
	return s
}

type DescribeAddonTemplateRequest struct {
	// The addon name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login
	AddonName *string `json:"AddonName,omitempty" xml:"AddonName,omitempty"`
	// The addon version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	AddonVersion *string `json:"AddonVersion,omitempty" xml:"AddonVersion,omitempty"`
	// The page number of the page returned. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAddonTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateRequest) SetAddonName(v string) *DescribeAddonTemplateRequest {
	s.AddonName = &v
	return s
}

func (s *DescribeAddonTemplateRequest) SetAddonVersion(v string) *DescribeAddonTemplateRequest {
	s.AddonVersion = &v
	return s
}

func (s *DescribeAddonTemplateRequest) SetPageNumber(v int64) *DescribeAddonTemplateRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAddonTemplateRequest) SetPageSize(v int64) *DescribeAddonTemplateRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAddonTemplateRequest) SetRegionId(v string) *DescribeAddonTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAddonTemplateRequest) SetZoneId(v string) *DescribeAddonTemplateRequest {
	s.ZoneId = &v
	return s
}

type DescribeAddonTemplateResponseBody struct {
	// The details of the addon template.
	Addon *DescribeAddonTemplateResponseBodyAddon `json:"Addon,omitempty" xml:"Addon,omitempty" type:"Struct"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAddonTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateResponseBody) SetAddon(v *DescribeAddonTemplateResponseBodyAddon) *DescribeAddonTemplateResponseBody {
	s.Addon = v
	return s
}

func (s *DescribeAddonTemplateResponseBody) SetPageNumber(v int64) *DescribeAddonTemplateResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAddonTemplateResponseBody) SetPageSize(v int64) *DescribeAddonTemplateResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAddonTemplateResponseBody) SetRequestId(v string) *DescribeAddonTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAddonTemplateResponseBody) SetTotalCount(v int32) *DescribeAddonTemplateResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAddonTemplateResponseBodyAddon struct {
	// The addon description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The addon icon.
	//
	// example:
	//
	// /assets/icons/your_icon.svg
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The addon label.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The date when the addon template was last updated.
	//
	// example:
	//
	// 2024-08-22 18:11:17
	LastUpdate *string `json:"LastUpdate,omitempty" xml:"LastUpdate,omitempty"`
	// The addon name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource configurations of the addon.
	ResourcesSpec *DescribeAddonTemplateResponseBodyAddonResourcesSpec `json:"ResourcesSpec,omitempty" xml:"ResourcesSpec,omitempty" type:"Struct"`
	// The addon configurations.
	ServicesSpec []*DescribeAddonTemplateResponseBodyAddonServicesSpec `json:"ServicesSpec,omitempty" xml:"ServicesSpec,omitempty" type:"Repeated"`
	// The addon version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAddonTemplateResponseBodyAddon) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonTemplateResponseBodyAddon) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateResponseBodyAddon) SetDescription(v string) *DescribeAddonTemplateResponseBodyAddon {
	s.Description = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddon) SetIcon(v string) *DescribeAddonTemplateResponseBodyAddon {
	s.Icon = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddon) SetLabel(v string) *DescribeAddonTemplateResponseBodyAddon {
	s.Label = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddon) SetLastUpdate(v string) *DescribeAddonTemplateResponseBodyAddon {
	s.LastUpdate = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddon) SetName(v string) *DescribeAddonTemplateResponseBodyAddon {
	s.Name = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddon) SetResourcesSpec(v *DescribeAddonTemplateResponseBodyAddonResourcesSpec) *DescribeAddonTemplateResponseBodyAddon {
	s.ResourcesSpec = v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddon) SetServicesSpec(v []*DescribeAddonTemplateResponseBodyAddonServicesSpec) *DescribeAddonTemplateResponseBodyAddon {
	s.ServicesSpec = v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddon) SetVersion(v string) *DescribeAddonTemplateResponseBodyAddon {
	s.Version = &v
	return s
}

type DescribeAddonTemplateResponseBodyAddonResourcesSpec struct {
	// The Elastic Compute Service (ECS) resource configurations of the addon.
	EcsResources []*AddonNodeTemplate `json:"EcsResources,omitempty" xml:"EcsResources,omitempty" type:"Repeated"`
	// The Elastic IP Address (EIP) configurations of the service.
	EipResource *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource `json:"EipResource,omitempty" xml:"EipResource,omitempty" type:"Struct"`
}

func (s DescribeAddonTemplateResponseBodyAddonResourcesSpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonTemplateResponseBodyAddonResourcesSpec) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpec) SetEcsResources(v []*AddonNodeTemplate) *DescribeAddonTemplateResponseBodyAddonResourcesSpec {
	s.EcsResources = v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpec) SetEipResource(v *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) *DescribeAddonTemplateResponseBodyAddonResourcesSpec {
	s.EipResource = v
	return s
}

type DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource struct {
	// Indicates whether the EIP is automatically created.
	//
	// example:
	//
	// True
	AutoCreate *bool `json:"AutoCreate,omitempty" xml:"AutoCreate,omitempty"`
	// The maximum bandwidth of the EIP. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The EIP ID.
	//
	// example:
	//
	// eip-bp1jwtsuoiif2qf90****
	EipInstanceId *string `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	// The billing method of the EIP. Valid values:
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// 	- PrePaid: subscription.
	//
	// Default value: PostPaid
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the EIP. Valid values:
	//
	// 	- PayByBandwidth: pay by bandwidth.
	//
	// 	- PayByTraffic: pay by data transfer.
	//
	// Valid values of N: 1 to 10.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
}

func (s DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) SetAutoCreate(v bool) *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource {
	s.AutoCreate = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) SetBandwidth(v string) *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource {
	s.Bandwidth = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) SetEipInstanceId(v string) *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource {
	s.EipInstanceId = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) SetInstanceChargeType(v string) *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) SetInternetChargeType(v string) *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource {
	s.InternetChargeType = &v
	return s
}

type DescribeAddonTemplateResponseBodyAddonServicesSpec struct {
	// The parameter configurations of the service.
	InputParams []*DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams `json:"InputParams,omitempty" xml:"InputParams,omitempty" type:"Repeated"`
	// The security group configurations of the service.
	NetworkACL []*DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL `json:"NetworkACL,omitempty" xml:"NetworkACL,omitempty" type:"Repeated"`
	// The service access type.
	//
	// example:
	//
	// URL
	ServiceAccessType *string `json:"ServiceAccessType,omitempty" xml:"ServiceAccessType,omitempty"`
	// The service access URL.
	//
	// example:
	//
	// https://47.96.xx.xx:12008
	ServiceAccessUrl *string `json:"ServiceAccessUrl,omitempty" xml:"ServiceAccessUrl,omitempty"`
	// The service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Web Portal
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeAddonTemplateResponseBodyAddonServicesSpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonTemplateResponseBodyAddonServicesSpec) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpec) SetInputParams(v []*DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) *DescribeAddonTemplateResponseBodyAddonServicesSpec {
	s.InputParams = v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpec) SetNetworkACL(v []*DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL) *DescribeAddonTemplateResponseBodyAddonServicesSpec {
	s.NetworkACL = v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpec) SetServiceAccessType(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpec {
	s.ServiceAccessType = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpec) SetServiceAccessUrl(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpec {
	s.ServiceAccessUrl = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpec) SetServiceName(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpec {
	s.ServiceName = &v
	return s
}

type DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams struct {
	// The help information of the parameter.
	HelpText *string `json:"HelpText,omitempty" xml:"HelpText,omitempty"`
	// The parameter label.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The parameter name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MYSQL_HOME
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter type.
	//
	// This parameter is required.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The parameter value.
	//
	// This parameter is required.
	//
	// example:
	//
	// usr/local/mysql
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) SetHelpText(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams {
	s.HelpText = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) SetLabel(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams {
	s.Label = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) SetName(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams {
	s.Name = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) SetType(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams {
	s.Type = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) SetValue(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams {
	s.Value = &v
	return s
}

type DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL struct {
	// The protocol type. Valid values:
	//
	// 	- **TCP**: forwards TCP packets.
	//
	// 	- **UDP**: forwards UDP packets.
	//
	// 	- **Any**: forwards all packets.
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The port number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *float32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The source CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.0.0/12
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL) SetIpProtocol(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL {
	s.IpProtocol = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL) SetPort(v float32) *DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL {
	s.Port = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL) SetSourceCidrIp(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL {
	s.SourceCidrIp = &v
	return s
}

type DescribeAddonTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAddonTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAddonTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddonTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateResponse) SetHeaders(v map[string]*string) *DescribeAddonTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeAddonTemplateResponse) SetStatusCode(v int32) *DescribeAddonTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAddonTemplateResponse) SetBody(v *DescribeAddonTemplateResponseBody) *DescribeAddonTemplateResponse {
	s.Body = v
	return s
}

type DetachSharedStoragesRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about mounted shared storage resources.
	//
	// This parameter is required.
	SharedStorages []*DetachSharedStoragesRequestSharedStorages `json:"SharedStorages,omitempty" xml:"SharedStorages,omitempty" type:"Repeated"`
}

func (s DetachSharedStoragesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachSharedStoragesRequest) GoString() string {
	return s.String()
}

func (s *DetachSharedStoragesRequest) SetClusterId(v string) *DetachSharedStoragesRequest {
	s.ClusterId = &v
	return s
}

func (s *DetachSharedStoragesRequest) SetSharedStorages(v []*DetachSharedStoragesRequestSharedStorages) *DetachSharedStoragesRequest {
	s.SharedStorages = v
	return s
}

type DetachSharedStoragesRequestSharedStorages struct {
	// The local mount directory of the mounted file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// /test
	MountDirectory *string `json:"MountDirectory,omitempty" xml:"MountDirectory,omitempty"`
}

func (s DetachSharedStoragesRequestSharedStorages) String() string {
	return tea.Prettify(s)
}

func (s DetachSharedStoragesRequestSharedStorages) GoString() string {
	return s.String()
}

func (s *DetachSharedStoragesRequestSharedStorages) SetMountDirectory(v string) *DetachSharedStoragesRequestSharedStorages {
	s.MountDirectory = &v
	return s
}

type DetachSharedStoragesShrinkRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about mounted shared storage resources.
	//
	// This parameter is required.
	SharedStoragesShrink *string `json:"SharedStorages,omitempty" xml:"SharedStorages,omitempty"`
}

func (s DetachSharedStoragesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachSharedStoragesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetachSharedStoragesShrinkRequest) SetClusterId(v string) *DetachSharedStoragesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DetachSharedStoragesShrinkRequest) SetSharedStoragesShrink(v string) *DetachSharedStoragesShrinkRequest {
	s.SharedStoragesShrink = &v
	return s
}

type DetachSharedStoragesResponseBody struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachSharedStoragesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachSharedStoragesResponseBody) GoString() string {
	return s.String()
}

func (s *DetachSharedStoragesResponseBody) SetClusterId(v string) *DetachSharedStoragesResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DetachSharedStoragesResponseBody) SetRequestId(v string) *DetachSharedStoragesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachSharedStoragesResponseBody) SetSuccess(v string) *DetachSharedStoragesResponseBody {
	s.Success = &v
	return s
}

type DetachSharedStoragesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachSharedStoragesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachSharedStoragesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachSharedStoragesResponse) GoString() string {
	return s.String()
}

func (s *DetachSharedStoragesResponse) SetHeaders(v map[string]*string) *DetachSharedStoragesResponse {
	s.Headers = v
	return s
}

func (s *DetachSharedStoragesResponse) SetStatusCode(v int32) *DetachSharedStoragesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachSharedStoragesResponse) SetBody(v *DetachSharedStoragesResponseBody) *DetachSharedStoragesResponse {
	s.Body = v
	return s
}

type GetAddonRequest struct {
	// The addon ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login-1.0-W4g****
	AddonId *string `json:"AddonId,omitempty" xml:"AddonId,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetAddonRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAddonRequest) GoString() string {
	return s.String()
}

func (s *GetAddonRequest) SetAddonId(v string) *GetAddonRequest {
	s.AddonId = &v
	return s
}

func (s *GetAddonRequest) SetClusterId(v string) *GetAddonRequest {
	s.ClusterId = &v
	return s
}

type GetAddonResponseBody struct {
	// The information about the addon.
	Addon *GetAddonResponseBodyAddon `json:"Addon,omitempty" xml:"Addon,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BBC2F93D-003A-49C4-850C-B826EECF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAddonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAddonResponseBody) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBody) SetAddon(v *GetAddonResponseBodyAddon) *GetAddonResponseBody {
	s.Addon = v
	return s
}

func (s *GetAddonResponseBody) SetRequestId(v string) *GetAddonResponseBody {
	s.RequestId = &v
	return s
}

type GetAddonResponseBodyAddon struct {
	// The addon ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login-1.0-W2g****
	AddonId *string `json:"AddonId,omitempty" xml:"AddonId,omitempty"`
	// The addon description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The addon icon.
	//
	// example:
	//
	// /assets/icons/your_icon.svg
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The time when the addon was installed.
	//
	// example:
	//
	// 2024-08-22 18:11:17
	InstallTime *string `json:"InstallTime,omitempty" xml:"InstallTime,omitempty"`
	// The addon label.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The addon name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource configurations of the addon.
	ResourcesSpec *GetAddonResponseBodyAddonResourcesSpec `json:"ResourcesSpec,omitempty" xml:"ResourcesSpec,omitempty" type:"Struct"`
	// The service configurations of the addon.
	ServicesSpec []*GetAddonResponseBodyAddonServicesSpec `json:"ServicesSpec,omitempty" xml:"ServicesSpec,omitempty" type:"Repeated"`
	// The addon status.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The addon version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAddonResponseBodyAddon) String() string {
	return tea.Prettify(s)
}

func (s GetAddonResponseBodyAddon) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyAddon) SetAddonId(v string) *GetAddonResponseBodyAddon {
	s.AddonId = &v
	return s
}

func (s *GetAddonResponseBodyAddon) SetDescription(v string) *GetAddonResponseBodyAddon {
	s.Description = &v
	return s
}

func (s *GetAddonResponseBodyAddon) SetIcon(v string) *GetAddonResponseBodyAddon {
	s.Icon = &v
	return s
}

func (s *GetAddonResponseBodyAddon) SetInstallTime(v string) *GetAddonResponseBodyAddon {
	s.InstallTime = &v
	return s
}

func (s *GetAddonResponseBodyAddon) SetLabel(v string) *GetAddonResponseBodyAddon {
	s.Label = &v
	return s
}

func (s *GetAddonResponseBodyAddon) SetName(v string) *GetAddonResponseBodyAddon {
	s.Name = &v
	return s
}

func (s *GetAddonResponseBodyAddon) SetResourcesSpec(v *GetAddonResponseBodyAddonResourcesSpec) *GetAddonResponseBodyAddon {
	s.ResourcesSpec = v
	return s
}

func (s *GetAddonResponseBodyAddon) SetServicesSpec(v []*GetAddonResponseBodyAddonServicesSpec) *GetAddonResponseBodyAddon {
	s.ServicesSpec = v
	return s
}

func (s *GetAddonResponseBodyAddon) SetStatus(v string) *GetAddonResponseBodyAddon {
	s.Status = &v
	return s
}

func (s *GetAddonResponseBodyAddon) SetVersion(v string) *GetAddonResponseBodyAddon {
	s.Version = &v
	return s
}

type GetAddonResponseBodyAddonResourcesSpec struct {
	// The Elastic Compute Service (ECS) resource configurations of the addon.
	EcsResources []*AddonNodeTemplate `json:"EcsResources,omitempty" xml:"EcsResources,omitempty" type:"Repeated"`
	// The Elastic IP Address (EIP) configurations.
	EipResource *GetAddonResponseBodyAddonResourcesSpecEipResource `json:"EipResource,omitempty" xml:"EipResource,omitempty" type:"Struct"`
}

func (s GetAddonResponseBodyAddonResourcesSpec) String() string {
	return tea.Prettify(s)
}

func (s GetAddonResponseBodyAddonResourcesSpec) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyAddonResourcesSpec) SetEcsResources(v []*AddonNodeTemplate) *GetAddonResponseBodyAddonResourcesSpec {
	s.EcsResources = v
	return s
}

func (s *GetAddonResponseBodyAddonResourcesSpec) SetEipResource(v *GetAddonResponseBodyAddonResourcesSpecEipResource) *GetAddonResponseBodyAddonResourcesSpec {
	s.EipResource = v
	return s
}

type GetAddonResponseBodyAddonResourcesSpecEipResource struct {
	// Indicates whether the EIP is automatically created.
	//
	// example:
	//
	// True
	AutoCreate *bool `json:"AutoCreate,omitempty" xml:"AutoCreate,omitempty"`
	// The maximum bandwidth of the EIP. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The EIP.
	//
	// example:
	//
	// 39.108.xx.xx
	EipAddress *string `json:"EipAddress,omitempty" xml:"EipAddress,omitempty"`
	// The EIP ID.
	//
	// example:
	//
	// eip-bp1vi9124tbx1g3kr****
	EipInstanceId *string `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	// The billing method of the EIP.
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// 	- PrePaid: subscription.
	//
	// Default value: PostPaid.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the EIP. Valid values:
	//
	// 	- PayByBandwidth: pay by bandwidth.
	//
	// 	- PayByTraffic: pay by data transfer.
	//
	// Valid values of N: 1 to 10.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
}

func (s GetAddonResponseBodyAddonResourcesSpecEipResource) String() string {
	return tea.Prettify(s)
}

func (s GetAddonResponseBodyAddonResourcesSpecEipResource) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) SetAutoCreate(v bool) *GetAddonResponseBodyAddonResourcesSpecEipResource {
	s.AutoCreate = &v
	return s
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) SetBandwidth(v string) *GetAddonResponseBodyAddonResourcesSpecEipResource {
	s.Bandwidth = &v
	return s
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) SetEipAddress(v string) *GetAddonResponseBodyAddonResourcesSpecEipResource {
	s.EipAddress = &v
	return s
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) SetEipInstanceId(v string) *GetAddonResponseBodyAddonResourcesSpecEipResource {
	s.EipInstanceId = &v
	return s
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) SetInstanceChargeType(v string) *GetAddonResponseBodyAddonResourcesSpecEipResource {
	s.InstanceChargeType = &v
	return s
}

func (s *GetAddonResponseBodyAddonResourcesSpecEipResource) SetInternetChargeType(v string) *GetAddonResponseBodyAddonResourcesSpecEipResource {
	s.InternetChargeType = &v
	return s
}

type GetAddonResponseBodyAddonServicesSpec struct {
	// The parameter configurations of the service.
	InputParams []*GetAddonResponseBodyAddonServicesSpecInputParams `json:"InputParams,omitempty" xml:"InputParams,omitempty" type:"Repeated"`
	// The security group configurations of the service.
	NetworkACL []*GetAddonResponseBodyAddonServicesSpecNetworkACL `json:"NetworkACL,omitempty" xml:"NetworkACL,omitempty" type:"Repeated"`
	// The service access type.
	//
	// example:
	//
	// URL
	ServiceAccessType *string `json:"ServiceAccessType,omitempty" xml:"ServiceAccessType,omitempty"`
	// The service access URL.
	//
	// example:
	//
	// https://47.96.xx.xx:12008
	ServiceAccessUrl *string `json:"ServiceAccessUrl,omitempty" xml:"ServiceAccessUrl,omitempty"`
	// The service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Web Portal
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetAddonResponseBodyAddonServicesSpec) String() string {
	return tea.Prettify(s)
}

func (s GetAddonResponseBodyAddonServicesSpec) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyAddonServicesSpec) SetInputParams(v []*GetAddonResponseBodyAddonServicesSpecInputParams) *GetAddonResponseBodyAddonServicesSpec {
	s.InputParams = v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpec) SetNetworkACL(v []*GetAddonResponseBodyAddonServicesSpecNetworkACL) *GetAddonResponseBodyAddonServicesSpec {
	s.NetworkACL = v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpec) SetServiceAccessType(v string) *GetAddonResponseBodyAddonServicesSpec {
	s.ServiceAccessType = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpec) SetServiceAccessUrl(v string) *GetAddonResponseBodyAddonServicesSpec {
	s.ServiceAccessUrl = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpec) SetServiceName(v string) *GetAddonResponseBodyAddonServicesSpec {
	s.ServiceName = &v
	return s
}

type GetAddonResponseBodyAddonServicesSpecInputParams struct {
	// The help information of the parameter.
	HelpText *string `json:"HelpText,omitempty" xml:"HelpText,omitempty"`
	// The parameter label.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The parameter name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MYSQL_HOME
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter type.
	//
	// This parameter is required.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The parameter value.
	//
	// This parameter is required.
	//
	// example:
	//
	// usr/local/mysql
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAddonResponseBodyAddonServicesSpecInputParams) String() string {
	return tea.Prettify(s)
}

func (s GetAddonResponseBodyAddonServicesSpecInputParams) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyAddonServicesSpecInputParams) SetHelpText(v string) *GetAddonResponseBodyAddonServicesSpecInputParams {
	s.HelpText = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpecInputParams) SetLabel(v string) *GetAddonResponseBodyAddonServicesSpecInputParams {
	s.Label = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpecInputParams) SetName(v string) *GetAddonResponseBodyAddonServicesSpecInputParams {
	s.Name = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpecInputParams) SetType(v string) *GetAddonResponseBodyAddonServicesSpecInputParams {
	s.Type = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpecInputParams) SetValue(v string) *GetAddonResponseBodyAddonServicesSpecInputParams {
	s.Value = &v
	return s
}

type GetAddonResponseBodyAddonServicesSpecNetworkACL struct {
	// The protocol type. Valid values:
	//
	// TCP: forwards TCP packets.
	//
	// UDP: forwards UDP packets.
	//
	// Any: forwards all packets.
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The port number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *float32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The source CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.0.0/12
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s GetAddonResponseBodyAddonServicesSpecNetworkACL) String() string {
	return tea.Prettify(s)
}

func (s GetAddonResponseBodyAddonServicesSpecNetworkACL) GoString() string {
	return s.String()
}

func (s *GetAddonResponseBodyAddonServicesSpecNetworkACL) SetIpProtocol(v string) *GetAddonResponseBodyAddonServicesSpecNetworkACL {
	s.IpProtocol = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpecNetworkACL) SetPort(v float32) *GetAddonResponseBodyAddonServicesSpecNetworkACL {
	s.Port = &v
	return s
}

func (s *GetAddonResponseBodyAddonServicesSpecNetworkACL) SetSourceCidrIp(v string) *GetAddonResponseBodyAddonServicesSpecNetworkACL {
	s.SourceCidrIp = &v
	return s
}

type GetAddonResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAddonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAddonResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAddonResponse) GoString() string {
	return s.String()
}

func (s *GetAddonResponse) SetHeaders(v map[string]*string) *GetAddonResponse {
	s.Headers = v
	return s
}

func (s *GetAddonResponse) SetStatusCode(v int32) *GetAddonResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAddonResponse) SetBody(v *GetAddonResponseBody) *GetAddonResponse {
	s.Body = v
	return s
}

type GetClusterRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRequest) GoString() string {
	return s.String()
}

func (s *GetClusterRequest) SetClusterId(v string) *GetClusterRequest {
	s.ClusterId = &v
	return s
}

type GetClusterResponseBody struct {
	// The E-HPC Util version.
	//
	// example:
	//
	// 2.0.31
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The cluster type. Valid values:
	//
	// 	- Standard
	//
	// 	- Serverless
	//
	// example:
	//
	// Standard
	ClusterCategory *string `json:"ClusterCategory,omitempty" xml:"ClusterCategory,omitempty"`
	// The time when the cluster was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2024-08-06T12:43:01.000Z
	ClusterCreateTime *string `json:"ClusterCreateTime,omitempty" xml:"ClusterCreateTime,omitempty"`
	// The post-processing script of the cluster.
	ClusterCustomConfiguration *GetClusterResponseBodyClusterCustomConfiguration `json:"ClusterCustomConfiguration,omitempty" xml:"ClusterCustomConfiguration,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The deployment type of the cluster. Valid values:
	//
	// 	- Integrated: The cluster is deployed on a public cloud.
	//
	// 	- Hybrid: The cluster is deployed on a hybrid cloud.
	//
	// 	- Custom: The cluster is a custom cluster.
	//
	// example:
	//
	// Integrated
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// The time when the cluster was last modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2024-08-06T12:43:01.000Z
	ClusterModifyTime *string `json:"ClusterModifyTime,omitempty" xml:"ClusterModifyTime,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// slurm22.05.8-cluster-20240614
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The cluster state. Valid values:
	//
	// 	- uninit: The cluster is being installed.
	//
	// 	- creating: The cluster is being created.
	//
	// 	- initing: The cluster is being initialized.
	//
	// 	- running: The cluster is running.
	//
	// 	- exception: The cluster has run into an exception.
	//
	// 	- raleasing: The cluster is being released.
	//
	// 	- stopping: The cluster is being stopped.
	//
	// 	- stopped: The cluster is stopped.
	//
	// 	- pending: The cluster is waiting to be configured.
	//
	// example:
	//
	// running
	ClusterStatus *string `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	// The ID of the vSwitch used by the cluster.
	//
	// example:
	//
	// vsw-bp1p2uugqsjppno******
	ClusterVSwitchId *string `json:"ClusterVSwitchId,omitempty" xml:"ClusterVSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) used by the cluster.
	//
	// example:
	//
	// vpc-uf6u3lk1pjy28eg*****
	ClusterVpcId *string `json:"ClusterVpcId,omitempty" xml:"ClusterVpcId,omitempty"`
	// Indicates whether deletion protection is enabled for the cluster. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	DeleteProtection *string `json:"DeleteProtection,omitempty" xml:"DeleteProtection,omitempty"`
	// The E-HPC version.
	//
	// example:
	//
	// 2.0.0
	EhpcVersion *string `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	// Indicates whether automatic scale-in is enabled for the cluster. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableScaleIn *bool `json:"EnableScaleIn,omitempty" xml:"EnableScaleIn,omitempty"`
	// Indicates whether automatic scale-out is enabled for the cluster. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableScaleOut *bool `json:"EnableScaleOut,omitempty" xml:"EnableScaleOut,omitempty"`
	// The interval at which the cluster is automatically scaled out.
	//
	// example:
	//
	// 2
	GrowInterval *int32 `json:"GrowInterval,omitempty" xml:"GrowInterval,omitempty"`
	// The idle duration of the compute nodes allowed by the cluster.
	//
	// example:
	//
	// 4
	IdleInterval *int32 `json:"IdleInterval,omitempty" xml:"IdleInterval,omitempty"`
	// The management node configurations.
	Manager *GetClusterResponseBodyManager `json:"Manager,omitempty" xml:"Manager,omitempty" type:"Struct"`
	// The maximum total number of vCPUs that can be used by all compute nodes managed by the cluster.
	//
	// example:
	//
	// 10000
	MaxCoreCount *string `json:"MaxCoreCount,omitempty" xml:"MaxCoreCount,omitempty"`
	// The maximum number of compute nodes that the cluster can manage.
	//
	// example:
	//
	// 100
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The monitoring details of the cluster.
	MonitorSpec *GetClusterResponseBodyMonitorSpec `json:"MonitorSpec,omitempty" xml:"MonitorSpec,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The scheduler specifications of the cluster.
	SchedulerSpec *GetClusterResponseBodySchedulerSpec `json:"SchedulerSpec,omitempty" xml:"SchedulerSpec,omitempty" type:"Struct"`
	// The security group ID.
	//
	// example:
	//
	// sg-f8z9vb2zaezpkr69a21k
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s GetClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBody) SetClientVersion(v string) *GetClusterResponseBody {
	s.ClientVersion = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterCategory(v string) *GetClusterResponseBody {
	s.ClusterCategory = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterCreateTime(v string) *GetClusterResponseBody {
	s.ClusterCreateTime = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterCustomConfiguration(v *GetClusterResponseBodyClusterCustomConfiguration) *GetClusterResponseBody {
	s.ClusterCustomConfiguration = v
	return s
}

func (s *GetClusterResponseBody) SetClusterId(v string) *GetClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterMode(v string) *GetClusterResponseBody {
	s.ClusterMode = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterModifyTime(v string) *GetClusterResponseBody {
	s.ClusterModifyTime = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterName(v string) *GetClusterResponseBody {
	s.ClusterName = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterStatus(v string) *GetClusterResponseBody {
	s.ClusterStatus = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterVSwitchId(v string) *GetClusterResponseBody {
	s.ClusterVSwitchId = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterVpcId(v string) *GetClusterResponseBody {
	s.ClusterVpcId = &v
	return s
}

func (s *GetClusterResponseBody) SetDeleteProtection(v string) *GetClusterResponseBody {
	s.DeleteProtection = &v
	return s
}

func (s *GetClusterResponseBody) SetEhpcVersion(v string) *GetClusterResponseBody {
	s.EhpcVersion = &v
	return s
}

func (s *GetClusterResponseBody) SetEnableScaleIn(v bool) *GetClusterResponseBody {
	s.EnableScaleIn = &v
	return s
}

func (s *GetClusterResponseBody) SetEnableScaleOut(v bool) *GetClusterResponseBody {
	s.EnableScaleOut = &v
	return s
}

func (s *GetClusterResponseBody) SetGrowInterval(v int32) *GetClusterResponseBody {
	s.GrowInterval = &v
	return s
}

func (s *GetClusterResponseBody) SetIdleInterval(v int32) *GetClusterResponseBody {
	s.IdleInterval = &v
	return s
}

func (s *GetClusterResponseBody) SetManager(v *GetClusterResponseBodyManager) *GetClusterResponseBody {
	s.Manager = v
	return s
}

func (s *GetClusterResponseBody) SetMaxCoreCount(v string) *GetClusterResponseBody {
	s.MaxCoreCount = &v
	return s
}

func (s *GetClusterResponseBody) SetMaxCount(v string) *GetClusterResponseBody {
	s.MaxCount = &v
	return s
}

func (s *GetClusterResponseBody) SetMonitorSpec(v *GetClusterResponseBodyMonitorSpec) *GetClusterResponseBody {
	s.MonitorSpec = v
	return s
}

func (s *GetClusterResponseBody) SetRequestId(v string) *GetClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterResponseBody) SetResourceGroupId(v string) *GetClusterResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetClusterResponseBody) SetSchedulerSpec(v *GetClusterResponseBodySchedulerSpec) *GetClusterResponseBody {
	s.SchedulerSpec = v
	return s
}

func (s *GetClusterResponseBody) SetSecurityGroupId(v string) *GetClusterResponseBody {
	s.SecurityGroupId = &v
	return s
}

type GetClusterResponseBodyClusterCustomConfiguration struct {
	// The arguments that are used to run the script after the scrip is installed.
	//
	// example:
	//
	// E-HPC cn-hangzhou
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The URL that is used to download the post-processing script.
	//
	// example:
	//
	// http://*****
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s GetClusterResponseBodyClusterCustomConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyClusterCustomConfiguration) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyClusterCustomConfiguration) SetArgs(v string) *GetClusterResponseBodyClusterCustomConfiguration {
	s.Args = &v
	return s
}

func (s *GetClusterResponseBodyClusterCustomConfiguration) SetScript(v string) *GetClusterResponseBodyClusterCustomConfiguration {
	s.Script = &v
	return s
}

type GetClusterResponseBodyManager struct {
	// The configurations of the domain name resolution service.
	DNS *GetClusterResponseBodyManagerDNS `json:"DNS,omitempty" xml:"DNS,omitempty" type:"Struct"`
	// The information about the domain account service.
	DirectoryService *GetClusterResponseBodyManagerDirectoryService `json:"DirectoryService,omitempty" xml:"DirectoryService,omitempty" type:"Struct"`
	// The configurations of the management node.
	ManagerNode *GetClusterResponseBodyManagerManagerNode `json:"ManagerNode,omitempty" xml:"ManagerNode,omitempty" type:"Struct"`
	// The information about the scheduler.
	Scheduler *GetClusterResponseBodyManagerScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Struct"`
}

func (s GetClusterResponseBodyManager) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyManager) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyManager) SetDNS(v *GetClusterResponseBodyManagerDNS) *GetClusterResponseBodyManager {
	s.DNS = v
	return s
}

func (s *GetClusterResponseBodyManager) SetDirectoryService(v *GetClusterResponseBodyManagerDirectoryService) *GetClusterResponseBodyManager {
	s.DirectoryService = v
	return s
}

func (s *GetClusterResponseBodyManager) SetManagerNode(v *GetClusterResponseBodyManagerManagerNode) *GetClusterResponseBodyManager {
	s.ManagerNode = v
	return s
}

func (s *GetClusterResponseBodyManager) SetScheduler(v *GetClusterResponseBodyManagerScheduler) *GetClusterResponseBodyManager {
	s.Scheduler = v
	return s
}

type GetClusterResponseBodyManagerDNS struct {
	// The state of the domain name resolution service. Valid values:
	//
	// 	- uninit: The service is being installed.
	//
	// 	- initing: The service is being initialized.
	//
	// 	- running: The service is running.
	//
	// 	- exception: The service has run into an exception.
	//
	// 	- releasing: The service is being released.
	//
	// 	- stopped: The service is stopped.
	//
	// 	- pending: The service is waiting to be configured.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The resolution type.
	//
	// example:
	//
	// nis
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the resolution service.
	//
	// example:
	//
	// 2.31
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetClusterResponseBodyManagerDNS) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyManagerDNS) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyManagerDNS) SetStatus(v string) *GetClusterResponseBodyManagerDNS {
	s.Status = &v
	return s
}

func (s *GetClusterResponseBodyManagerDNS) SetType(v string) *GetClusterResponseBodyManagerDNS {
	s.Type = &v
	return s
}

func (s *GetClusterResponseBodyManagerDNS) SetVersion(v string) *GetClusterResponseBodyManagerDNS {
	s.Version = &v
	return s
}

type GetClusterResponseBodyManagerDirectoryService struct {
	// The state of the domain account service. Valid values:
	//
	// 	- uninit: The service is being installed.
	//
	// 	- initing: The service is being initialized.
	//
	// 	- running: The service is running.
	//
	// 	- exception: The service has run into an exception.
	//
	// 	- releasing: The service is being released.
	//
	// 	- stopped: The service is stopped.
	//
	// 	- pending: The service is waiting to be configured.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the domain account.
	//
	// example:
	//
	// nis
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the domain account service.
	//
	// example:
	//
	// 2.31
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetClusterResponseBodyManagerDirectoryService) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyManagerDirectoryService) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyManagerDirectoryService) SetStatus(v string) *GetClusterResponseBodyManagerDirectoryService {
	s.Status = &v
	return s
}

func (s *GetClusterResponseBodyManagerDirectoryService) SetType(v string) *GetClusterResponseBodyManagerDirectoryService {
	s.Type = &v
	return s
}

func (s *GetClusterResponseBodyManagerDirectoryService) SetVersion(v string) *GetClusterResponseBodyManagerDirectoryService {
	s.Version = &v
	return s
}

type GetClusterResponseBodyManagerManagerNode struct {
	// The expiration time of the management node.
	//
	// example:
	//
	// 2099-12-31T15:59Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The instance billing method of the management node. Valid values:
	//
	// 	- PostPaid: pay-as-you-go
	//
	// 	- PrePaid: subscription
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance ID of the management node.
	//
	// example:
	//
	// i-bp1a170jgea1vl******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance type of the management node.
	//
	// example:
	//
	// ecs.g6.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s GetClusterResponseBodyManagerManagerNode) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyManagerManagerNode) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyManagerManagerNode) SetExpiredTime(v string) *GetClusterResponseBodyManagerManagerNode {
	s.ExpiredTime = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNode) SetInstanceChargeType(v string) *GetClusterResponseBodyManagerManagerNode {
	s.InstanceChargeType = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNode) SetInstanceId(v string) *GetClusterResponseBodyManagerManagerNode {
	s.InstanceId = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNode) SetInstanceType(v string) *GetClusterResponseBodyManagerManagerNode {
	s.InstanceType = &v
	return s
}

type GetClusterResponseBodyManagerScheduler struct {
	// The scheduler state. Valid values:
	//
	// 	- uninit: The scheduler is being installed.
	//
	// 	- initing: The scheduler is being initialized.
	//
	// 	- running: The scheduler is running.
	//
	// 	- exception: The scheduler has run into an exception.
	//
	// 	- releasing: The scheduler is being released.
	//
	// 	- stopped: The scheduler is stopped.
	//
	// 	- pending: The scheduler is waiting to be configured.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The scheduler type. Valid values:
	//
	// 	- SLURM
	//
	// 	- PBS
	//
	// 	- OPENGRIDSCHEDULER
	//
	// 	- LSF_PLUGIN
	//
	// 	- PBS_PLUGIN
	//
	// example:
	//
	// SLURM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The scheduler version.
	//
	// example:
	//
	// 22.05.8
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetClusterResponseBodyManagerScheduler) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyManagerScheduler) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyManagerScheduler) SetStatus(v string) *GetClusterResponseBodyManagerScheduler {
	s.Status = &v
	return s
}

func (s *GetClusterResponseBodyManagerScheduler) SetType(v string) *GetClusterResponseBodyManagerScheduler {
	s.Type = &v
	return s
}

func (s *GetClusterResponseBodyManagerScheduler) SetVersion(v string) *GetClusterResponseBodyManagerScheduler {
	s.Version = &v
	return s
}

type GetClusterResponseBodyMonitorSpec struct {
	// Indicates whether the monitoring component of compute nodes is enabled for the cluster. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableComputeLoadMonitor *bool `json:"EnableComputeLoadMonitor,omitempty" xml:"EnableComputeLoadMonitor,omitempty"`
}

func (s GetClusterResponseBodyMonitorSpec) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyMonitorSpec) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyMonitorSpec) SetEnableComputeLoadMonitor(v bool) *GetClusterResponseBodyMonitorSpec {
	s.EnableComputeLoadMonitor = &v
	return s
}

type GetClusterResponseBodySchedulerSpec struct {
	// Indicates whether the topology awareness feature is enabled for the cluster. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableTopologyAwareness *bool `json:"EnableTopologyAwareness,omitempty" xml:"EnableTopologyAwareness,omitempty"`
}

func (s GetClusterResponseBodySchedulerSpec) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodySchedulerSpec) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodySchedulerSpec) SetEnableTopologyAwareness(v bool) *GetClusterResponseBodySchedulerSpec {
	s.EnableTopologyAwareness = &v
	return s
}

type GetClusterResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponse) GoString() string {
	return s.String()
}

func (s *GetClusterResponse) SetHeaders(v map[string]*string) *GetClusterResponse {
	s.Headers = v
	return s
}

func (s *GetClusterResponse) SetStatusCode(v int32) *GetClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterResponse) SetBody(v *GetClusterResponseBody) *GetClusterResponse {
	s.Body = v
	return s
}

type GetCommonLogDetailRequest struct {
	// The start time of the time range within which the logs that you want to query were generated. The time is a timestamp. This value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1703821542
	From *int64 `json:"From,omitempty" xml:"From,omitempty"`
	// Specifies whether to hide the process of each step. Valid values:
	//
	// 	- true: hides the process and returns only the result log of each step.
	//
	// 	- false: does not hide the process and displays the start and result logs of each step.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	HiddenProcess *bool `json:"HiddenProcess,omitempty" xml:"HiddenProcess,omitempty"`
	// The request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	// The end time of the time range within which the logs that you want to query were generated. The time is a timestamp. This value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1703821666
	To *int64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetCommonLogDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCommonLogDetailRequest) GoString() string {
	return s.String()
}

func (s *GetCommonLogDetailRequest) SetFrom(v int64) *GetCommonLogDetailRequest {
	s.From = &v
	return s
}

func (s *GetCommonLogDetailRequest) SetHiddenProcess(v bool) *GetCommonLogDetailRequest {
	s.HiddenProcess = &v
	return s
}

func (s *GetCommonLogDetailRequest) SetLogRequestId(v string) *GetCommonLogDetailRequest {
	s.LogRequestId = &v
	return s
}

func (s *GetCommonLogDetailRequest) SetTo(v int64) *GetCommonLogDetailRequest {
	s.To = &v
	return s
}

type GetCommonLogDetailResponseBody struct {
	// The action name.
	//
	// example:
	//
	// CreateCluster
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-abc***
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about the logs.
	LogDetail []*GetCommonLogDetailResponseBodyLogDetail `json:"LogDetail,omitempty" xml:"LogDetail,omitempty" type:"Repeated"`
	// The log type.
	//
	// example:
	//
	// operation
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The account ID of the operator.
	//
	// example:
	//
	// 239***
	OperatorUid *string `json:"OperatorUid,omitempty" xml:"OperatorUid,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 464E9919-D04F-4D1D-B375-15989492****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 137***
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetCommonLogDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCommonLogDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetCommonLogDetailResponseBody) SetAction(v string) *GetCommonLogDetailResponseBody {
	s.Action = &v
	return s
}

func (s *GetCommonLogDetailResponseBody) SetClusterId(v string) *GetCommonLogDetailResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetCommonLogDetailResponseBody) SetLogDetail(v []*GetCommonLogDetailResponseBodyLogDetail) *GetCommonLogDetailResponseBody {
	s.LogDetail = v
	return s
}

func (s *GetCommonLogDetailResponseBody) SetLogType(v string) *GetCommonLogDetailResponseBody {
	s.LogType = &v
	return s
}

func (s *GetCommonLogDetailResponseBody) SetOperatorUid(v string) *GetCommonLogDetailResponseBody {
	s.OperatorUid = &v
	return s
}

func (s *GetCommonLogDetailResponseBody) SetRequestId(v string) *GetCommonLogDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCommonLogDetailResponseBody) SetUid(v string) *GetCommonLogDetailResponseBody {
	s.Uid = &v
	return s
}

type GetCommonLogDetailResponseBodyLogDetail struct {
	// The stage name of the log.
	//
	// example:
	//
	// ConfigNetwork
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The information about the log stages.
	Stages []*GetCommonLogDetailResponseBodyLogDetailStages `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
}

func (s GetCommonLogDetailResponseBodyLogDetail) String() string {
	return tea.Prettify(s)
}

func (s GetCommonLogDetailResponseBodyLogDetail) GoString() string {
	return s.String()
}

func (s *GetCommonLogDetailResponseBodyLogDetail) SetStageName(v string) *GetCommonLogDetailResponseBodyLogDetail {
	s.StageName = &v
	return s
}

func (s *GetCommonLogDetailResponseBodyLogDetail) SetStages(v []*GetCommonLogDetailResponseBodyLogDetailStages) *GetCommonLogDetailResponseBodyLogDetail {
	s.Stages = v
	return s
}

type GetCommonLogDetailResponseBodyLogDetailStages struct {
	// The log level.
	//
	// Valid values:
	//
	// 	- ERROR
	//
	// 	- INFO
	//
	// 	- WARN
	//
	// example:
	//
	// INFO
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	// The output message of the log.
	//
	// example:
	//
	// Successfully created security group sg-bcd***
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The method involved in the log.
	//
	// example:
	//
	// CreateSecurityGroup
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The request ID associated with the log.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The action state involved in the log. Valid values:
	//
	// 	- InProgress: The action is being performed.
	//
	// 	- Finished: The action is completed.
	//
	// 	- Failed: The action failed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The resource involved in the log.
	//
	// example:
	//
	// sg-bcd***
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The time when the log was generated.
	//
	// example:
	//
	// 2024-08-22 14:21:54
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetCommonLogDetailResponseBodyLogDetailStages) String() string {
	return tea.Prettify(s)
}

func (s GetCommonLogDetailResponseBodyLogDetailStages) GoString() string {
	return s.String()
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) SetLogLevel(v string) *GetCommonLogDetailResponseBodyLogDetailStages {
	s.LogLevel = &v
	return s
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) SetMessage(v string) *GetCommonLogDetailResponseBodyLogDetailStages {
	s.Message = &v
	return s
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) SetMethod(v string) *GetCommonLogDetailResponseBodyLogDetailStages {
	s.Method = &v
	return s
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) SetRequestId(v string) *GetCommonLogDetailResponseBodyLogDetailStages {
	s.RequestId = &v
	return s
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) SetStatus(v string) *GetCommonLogDetailResponseBodyLogDetailStages {
	s.Status = &v
	return s
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) SetTarget(v string) *GetCommonLogDetailResponseBodyLogDetailStages {
	s.Target = &v
	return s
}

func (s *GetCommonLogDetailResponseBodyLogDetailStages) SetTime(v string) *GetCommonLogDetailResponseBodyLogDetailStages {
	s.Time = &v
	return s
}

type GetCommonLogDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCommonLogDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCommonLogDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCommonLogDetailResponse) GoString() string {
	return s.String()
}

func (s *GetCommonLogDetailResponse) SetHeaders(v map[string]*string) *GetCommonLogDetailResponse {
	s.Headers = v
	return s
}

func (s *GetCommonLogDetailResponse) SetStatusCode(v int32) *GetCommonLogDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCommonLogDetailResponse) SetBody(v *GetCommonLogDetailResponseBody) *GetCommonLogDetailResponse {
	s.Body = v
	return s
}

type GetJobRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The job ID. You can call the ListJobs operation to query the job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.manager
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) SetClusterId(v string) *GetJobRequest {
	s.ClusterId = &v
	return s
}

func (s *GetJobRequest) SetJobId(v string) *GetJobRequest {
	s.JobId = &v
	return s
}

type GetJobResponseBody struct {
	// The job details.
	JobInfo *GetJobResponseBodyJobInfo `json:"JobInfo,omitempty" xml:"JobInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0****-1335-****-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) SetJobInfo(v *GetJobResponseBodyJobInfo) *GetJobResponseBody {
	s.JobInfo = v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobResponseBody) SetSuccess(v string) *GetJobResponseBody {
	s.Success = &v
	return s
}

type GetJobResponseBodyJobInfo struct {
	// The parent job ID. If the return value is a non-empty string, the job is an array job.
	//
	// example:
	//
	// 1
	ArrayJobId *string `json:"ArrayJobId,omitempty" xml:"ArrayJobId,omitempty"`
	// The sub-job ID. This parameter is valid when the ArrayJobId parameter is a non-empty string.
	//
	// example:
	//
	// 3
	ArrayJobSubId *string `json:"ArrayJobSubId,omitempty" xml:"ArrayJobSubId,omitempty"`
	// The job queue. If the job is not in a queue, the output is empty.
	//
	// The format is X-Y:Z. X indicates the first index, Y indicates the final index, and Z indicates the step size. For example, 2-7:2 indicates three sub-jobs numbered 2, 4, and 6.
	//
	// example:
	//
	// 1-5:2
	ArrayRequest *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	// The command that is used to run the job.
	//
	// example:
	//
	// /home/huangsf/ehpc/job_meta.pbs
	CommandLine *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	// The time when the job was submitted.
	//
	// example:
	//
	// 2024-08-16T10:52:48
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error log file of the job.
	//
	// example:
	//
	// /home/xxx/STDIN.e1
	ErrorLog *string `json:"ErrorLog,omitempty" xml:"ErrorLog,omitempty"`
	// Additional information.
	//
	// example:
	//
	// {}
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 1.manager
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job name.
	//
	// example:
	//
	// testJob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The queue to which the job belongs.
	//
	// example:
	//
	// workq
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	// The time when the job was last modified.
	//
	// example:
	//
	// 2024-08-16T10:52:48
	LastModifyTime *string `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The compute nodes that are used to run the job.
	//
	// example:
	//
	// compute000
	NodeList *string `json:"NodeList,omitempty" xml:"NodeList,omitempty"`
	// The standard output log file of the job.
	//
	// example:
	//
	// /home/xxx/STDIN.o1
	OutputLog *string `json:"OutputLog,omitempty" xml:"OutputLog,omitempty"`
	// The priority of the job.
	//
	// example:
	//
	// 0
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The resources that were requested when the job was submitted.
	Resources *GetJobResponseBodyJobInfoResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// The resources that are actually used by the job.
	ResourcesUsed *GetJobResponseBodyJobInfoResourcesUsed `json:"ResourcesUsed,omitempty" xml:"ResourcesUsed,omitempty" type:"Struct"`
	// The user to which the job belongs or that is used to submit the job. This user is a cluster-side user.
	//
	// example:
	//
	// testuser
	RunasUser *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	// The time when the job was started.
	//
	// example:
	//
	// 2024-08-16T10:52:48
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The job state.
	//
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The variables of the job.
	Variables []*GetJobResponseBodyJobInfoVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s GetJobResponseBodyJobInfo) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfo) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfo) SetArrayJobId(v string) *GetJobResponseBodyJobInfo {
	s.ArrayJobId = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetArrayJobSubId(v string) *GetJobResponseBodyJobInfo {
	s.ArrayJobSubId = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetArrayRequest(v string) *GetJobResponseBodyJobInfo {
	s.ArrayRequest = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetCommandLine(v string) *GetJobResponseBodyJobInfo {
	s.CommandLine = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetCreateTime(v string) *GetJobResponseBodyJobInfo {
	s.CreateTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetErrorLog(v string) *GetJobResponseBodyJobInfo {
	s.ErrorLog = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetExtraInfo(v string) *GetJobResponseBodyJobInfo {
	s.ExtraInfo = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetJobId(v string) *GetJobResponseBodyJobInfo {
	s.JobId = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetJobName(v string) *GetJobResponseBodyJobInfo {
	s.JobName = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetJobQueue(v string) *GetJobResponseBodyJobInfo {
	s.JobQueue = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetLastModifyTime(v string) *GetJobResponseBodyJobInfo {
	s.LastModifyTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetNodeList(v string) *GetJobResponseBodyJobInfo {
	s.NodeList = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetOutputLog(v string) *GetJobResponseBodyJobInfo {
	s.OutputLog = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetPriority(v string) *GetJobResponseBodyJobInfo {
	s.Priority = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetResources(v *GetJobResponseBodyJobInfoResources) *GetJobResponseBodyJobInfo {
	s.Resources = v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetResourcesUsed(v *GetJobResponseBodyJobInfoResourcesUsed) *GetJobResponseBodyJobInfo {
	s.ResourcesUsed = v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetRunasUser(v string) *GetJobResponseBodyJobInfo {
	s.RunasUser = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetStartTime(v string) *GetJobResponseBodyJobInfo {
	s.StartTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetState(v string) *GetJobResponseBodyJobInfo {
	s.State = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetVariables(v []*GetJobResponseBodyJobInfoVariables) *GetJobResponseBodyJobInfo {
	s.Variables = v
	return s
}

type GetJobResponseBodyJobInfoResources struct {
	// The number of vCPUs used by the job on each node.
	//
	// example:
	//
	// 2
	Cores *string `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The number of GPUs used by the job on each node.
	//
	// example:
	//
	// 1
	Gpus *string `json:"Gpus,omitempty" xml:"Gpus,omitempty"`
	// The memory size used by the job on each node.
	//
	// example:
	//
	// 1gb
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The number of nodes that are used to run the job.
	//
	// example:
	//
	// 1
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s GetJobResponseBodyJobInfoResources) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfoResources) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoResources) SetCores(v string) *GetJobResponseBodyJobInfoResources {
	s.Cores = &v
	return s
}

func (s *GetJobResponseBodyJobInfoResources) SetGpus(v string) *GetJobResponseBodyJobInfoResources {
	s.Gpus = &v
	return s
}

func (s *GetJobResponseBodyJobInfoResources) SetMemory(v string) *GetJobResponseBodyJobInfoResources {
	s.Memory = &v
	return s
}

func (s *GetJobResponseBodyJobInfoResources) SetNodes(v string) *GetJobResponseBodyJobInfoResources {
	s.Nodes = &v
	return s
}

type GetJobResponseBodyJobInfoResourcesUsed struct {
	// The number of vCPUs used by the job on each node.
	//
	// example:
	//
	// 2
	Cores *string `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The memory size used by the job on each node.
	//
	// example:
	//
	// 512mb
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The number of nodes that are used to run the job.
	//
	// example:
	//
	// 2
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s GetJobResponseBodyJobInfoResourcesUsed) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfoResourcesUsed) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoResourcesUsed) SetCores(v string) *GetJobResponseBodyJobInfoResourcesUsed {
	s.Cores = &v
	return s
}

func (s *GetJobResponseBodyJobInfoResourcesUsed) SetMemory(v string) *GetJobResponseBodyJobInfoResourcesUsed {
	s.Memory = &v
	return s
}

func (s *GetJobResponseBodyJobInfoResourcesUsed) SetNodes(v string) *GetJobResponseBodyJobInfoResourcesUsed {
	s.Nodes = &v
	return s
}

type GetJobResponseBodyJobInfoVariables struct {
	// The name of the environment variable.
	//
	// example:
	//
	// ProxyIP
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the environment variable.
	//
	// example:
	//
	// 10.x.x.x
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetJobResponseBodyJobInfoVariables) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfoVariables) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoVariables) SetName(v string) *GetJobResponseBodyJobInfoVariables {
	s.Name = &v
	return s
}

func (s *GetJobResponseBodyJobInfoVariables) SetValue(v string) *GetJobResponseBodyJobInfoVariables {
	s.Value = &v
	return s
}

type GetJobResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponse) GoString() string {
	return s.String()
}

func (s *GetJobResponse) SetHeaders(v map[string]*string) *GetJobResponse {
	s.Headers = v
	return s
}

func (s *GetJobResponse) SetStatusCode(v int32) *GetJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobResponse) SetBody(v *GetJobResponseBody) *GetJobResponse {
	s.Body = v
	return s
}

type GetJobLogRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-jeJki6****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.manager
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The log type. Valid values:
	//
	// 	- stdout: standard output logs.
	//
	// 	- stderr: error output logs.
	//
	// 	- all: all logs.
	//
	// Default value: all.
	//
	// example:
	//
	// stdout
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The position where logs start to be read.
	//
	// Unit: bytes.
	//
	// Default value: 0.
	//
	// example:
	//
	// 0
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The maximum size of logs that you can read in a single request.
	//
	// Unit: bytes.
	//
	// Default value: 10240.
	//
	// example:
	//
	// 20480
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s GetJobLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobLogRequest) GoString() string {
	return s.String()
}

func (s *GetJobLogRequest) SetClusterId(v string) *GetJobLogRequest {
	s.ClusterId = &v
	return s
}

func (s *GetJobLogRequest) SetJobId(v string) *GetJobLogRequest {
	s.JobId = &v
	return s
}

func (s *GetJobLogRequest) SetLogType(v string) *GetJobLogRequest {
	s.LogType = &v
	return s
}

func (s *GetJobLogRequest) SetOffset(v string) *GetJobLogRequest {
	s.Offset = &v
	return s
}

func (s *GetJobLogRequest) SetSize(v string) *GetJobLogRequest {
	s.Size = &v
	return s
}

type GetJobLogResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// 1.manager
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B745C159-3155-4B94-95D0-4B73D4D2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The error output log that is encoded in Base64.
	//
	// example:
	//
	// aG9zdG5hbWU=
	StderrLog *string `json:"StderrLog,omitempty" xml:"StderrLog,omitempty"`
	// The size of the error output file.
	//
	// example:
	//
	// 0
	StderrLogSize *string `json:"StderrLogSize,omitempty" xml:"StderrLogSize,omitempty"`
	// The standard output log that is encoded in Base64.
	//
	// example:
	//
	// aG9zdG5hbWU=
	StdoutLog *string `json:"StdoutLog,omitempty" xml:"StdoutLog,omitempty"`
	// The size of the standard output file.
	//
	// example:
	//
	// 4096
	StdoutLogSize *string `json:"StdoutLogSize,omitempty" xml:"StdoutLogSize,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJobLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobLogResponseBody) SetJobId(v string) *GetJobLogResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobLogResponseBody) SetRequestId(v string) *GetJobLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobLogResponseBody) SetStderrLog(v string) *GetJobLogResponseBody {
	s.StderrLog = &v
	return s
}

func (s *GetJobLogResponseBody) SetStderrLogSize(v string) *GetJobLogResponseBody {
	s.StderrLogSize = &v
	return s
}

func (s *GetJobLogResponseBody) SetStdoutLog(v string) *GetJobLogResponseBody {
	s.StdoutLog = &v
	return s
}

func (s *GetJobLogResponseBody) SetStdoutLogSize(v string) *GetJobLogResponseBody {
	s.StdoutLogSize = &v
	return s
}

func (s *GetJobLogResponseBody) SetSuccess(v string) *GetJobLogResponseBody {
	s.Success = &v
	return s
}

type GetJobLogResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobLogResponse) GoString() string {
	return s.String()
}

func (s *GetJobLogResponse) SetHeaders(v map[string]*string) *GetJobLogResponse {
	s.Headers = v
	return s
}

func (s *GetJobLogResponse) SetStatusCode(v int32) *GetJobLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobLogResponse) SetBody(v *GetJobLogResponseBody) *GetJobLogResponse {
	s.Body = v
	return s
}

type GetQueueRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The queue name.
	//
	// example:
	//
	// comp
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s GetQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQueueRequest) GoString() string {
	return s.String()
}

func (s *GetQueueRequest) SetClusterId(v string) *GetQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *GetQueueRequest) SetQueueName(v string) *GetQueueRequest {
	s.QueueName = &v
	return s
}

type GetQueueResponseBody struct {
	// The queue configurations.
	Queue *GetQueueResponseBodyQueue `json:"Queue,omitempty" xml:"Queue,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQueueResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueueResponseBody) SetQueue(v *GetQueueResponseBodyQueue) *GetQueueResponseBody {
	s.Queue = v
	return s
}

func (s *GetQueueResponseBody) SetRequestId(v string) *GetQueueResponseBody {
	s.RequestId = &v
	return s
}

type GetQueueResponseBodyQueue struct {
	// The auto scale-out policy of the queue.
	//
	// example:
	//
	// PriorityInstanceType
	AllocationStrategy *string `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	// The hardware configurations of the compute nodes in the queue.
	ComputeNodes []*NodeTemplate `json:"ComputeNodes,omitempty" xml:"ComputeNodes,omitempty" type:"Repeated"`
	// Indicates whether auto scale-in is enabled for the queue. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableScaleIn *bool `json:"EnableScaleIn,omitempty" xml:"EnableScaleIn,omitempty"`
	// Indicates whether auto scale-out is enabled for the queue. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableScaleOut *bool `json:"EnableScaleOut,omitempty" xml:"EnableScaleOut,omitempty"`
	// The hostname prefix of the compute nodes in the queue.
	//
	// example:
	//
	// compute
	HostnamePrefix *string `json:"HostnamePrefix,omitempty" xml:"HostnamePrefix,omitempty"`
	// The hostname suffix of the compute nodes in the queue.
	//
	// example:
	//
	// demo
	HostnameSuffix *string `json:"HostnameSuffix,omitempty" xml:"HostnameSuffix,omitempty"`
	// The initial number of nodes in the queue.
	//
	// example:
	//
	// 0
	InitialCount *int32 `json:"InitialCount,omitempty" xml:"InitialCount,omitempty"`
	// The type of the network between compute nodes in the queue. Valid values:
	//
	// 	- vpc
	//
	// 	- eRDMA
	//
	// example:
	//
	// erdma
	InterConnect *string `json:"InterConnect,omitempty" xml:"InterConnect,omitempty"`
	// The nodes for which deletion protection is enabled in the queue.
	KeepAliveNodes []*string `json:"KeepAliveNodes,omitempty" xml:"KeepAliveNodes,omitempty" type:"Repeated"`
	// The maximum number of compute nodes that the queue can contain.
	//
	// example:
	//
	// 1000
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The minimum number of nodes that are delivered to the queue in each scale-out cycle.
	//
	// example:
	//
	// 99
	MaxCountPerCycle *int64 `json:"MaxCountPerCycle,omitempty" xml:"MaxCountPerCycle,omitempty"`
	// The minimum number of compute nodes that the queue must contain.
	//
	// example:
	//
	// 0
	MinCount *int32 `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// The queue name.
	//
	// This parameter is required.
	//
	// example:
	//
	// comp
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The Resource Access Management (RAM) role that is assumed by compute nodes in the queue.
	//
	// example:
	//
	// AliyunECSInstanceForEHPCRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// Preset node pool ID.
	//
	// example:
	//
	// rnp-756vlp7a
	ReservedNodePoolId *string `json:"ReservedNodePoolId,omitempty" xml:"ReservedNodePoolId,omitempty"`
	// The available vSwitches for compute nodes in the queue. Valid values of N: 1 to 5.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s GetQueueResponseBodyQueue) String() string {
	return tea.Prettify(s)
}

func (s GetQueueResponseBodyQueue) GoString() string {
	return s.String()
}

func (s *GetQueueResponseBodyQueue) SetAllocationStrategy(v string) *GetQueueResponseBodyQueue {
	s.AllocationStrategy = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetComputeNodes(v []*NodeTemplate) *GetQueueResponseBodyQueue {
	s.ComputeNodes = v
	return s
}

func (s *GetQueueResponseBodyQueue) SetEnableScaleIn(v bool) *GetQueueResponseBodyQueue {
	s.EnableScaleIn = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetEnableScaleOut(v bool) *GetQueueResponseBodyQueue {
	s.EnableScaleOut = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetHostnamePrefix(v string) *GetQueueResponseBodyQueue {
	s.HostnamePrefix = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetHostnameSuffix(v string) *GetQueueResponseBodyQueue {
	s.HostnameSuffix = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetInitialCount(v int32) *GetQueueResponseBodyQueue {
	s.InitialCount = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetInterConnect(v string) *GetQueueResponseBodyQueue {
	s.InterConnect = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetKeepAliveNodes(v []*string) *GetQueueResponseBodyQueue {
	s.KeepAliveNodes = v
	return s
}

func (s *GetQueueResponseBodyQueue) SetMaxCount(v int32) *GetQueueResponseBodyQueue {
	s.MaxCount = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetMaxCountPerCycle(v int64) *GetQueueResponseBodyQueue {
	s.MaxCountPerCycle = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetMinCount(v int32) *GetQueueResponseBodyQueue {
	s.MinCount = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetQueueName(v string) *GetQueueResponseBodyQueue {
	s.QueueName = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetRamRole(v string) *GetQueueResponseBodyQueue {
	s.RamRole = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetReservedNodePoolId(v string) *GetQueueResponseBodyQueue {
	s.ReservedNodePoolId = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetVSwitchIds(v []*string) *GetQueueResponseBodyQueue {
	s.VSwitchIds = v
	return s
}

type GetQueueResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQueueResponse) GoString() string {
	return s.String()
}

func (s *GetQueueResponse) SetHeaders(v map[string]*string) *GetQueueResponse {
	s.Headers = v
	return s
}

func (s *GetQueueResponse) SetStatusCode(v int32) *GetQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueueResponse) SetBody(v *GetQueueResponseBody) *GetQueueResponse {
	s.Body = v
	return s
}

type InstallAddonRequest struct {
	// The addon name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login
	AddonName *string `json:"AddonName,omitempty" xml:"AddonName,omitempty"`
	// The addon version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	AddonVersion *string `json:"AddonVersion,omitempty" xml:"AddonVersion,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The resource configurations of the addon.
	//
	// This parameter is required.
	//
	// example:
	//
	// `{"EipResource": {"AutoCreate": true}, "EcsResources": [{"InstanceType": "ecs.c7.xlarge", "ImageId": "centos_7_6_xxx.vhd", "SystemDisk": {"Category": "cloud_essd", "Size": 40, "Level": "PL0"}]}`
	ResourcesSpec *string `json:"ResourcesSpec,omitempty" xml:"ResourcesSpec,omitempty"`
	// The service configurations of the addon.
	//
	// This parameter is required.
	//
	// example:
	//
	// `[{"ServiceName": "SSH", "ServiceAccessType": null, "ServiceAccessUrl": null, "NetworkACL": [{"IpProtocol": "TCP", "Port": 22, "SourceCidrIp": "0.0.0.0/0"}]}, {"ServiceName": "VNC", "ServiceAccessType": null, "ServiceAccessUrl": null, "NetworkACL": [{"IpProtocol": "TCP", "Port": 12016, "SourceCidrIp": "0.0.0.0/0"}]}]`
	ServicesSpec *string `json:"ServicesSpec,omitempty" xml:"ServicesSpec,omitempty"`
}

func (s InstallAddonRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallAddonRequest) GoString() string {
	return s.String()
}

func (s *InstallAddonRequest) SetAddonName(v string) *InstallAddonRequest {
	s.AddonName = &v
	return s
}

func (s *InstallAddonRequest) SetAddonVersion(v string) *InstallAddonRequest {
	s.AddonVersion = &v
	return s
}

func (s *InstallAddonRequest) SetClusterId(v string) *InstallAddonRequest {
	s.ClusterId = &v
	return s
}

func (s *InstallAddonRequest) SetResourcesSpec(v string) *InstallAddonRequest {
	s.ResourcesSpec = &v
	return s
}

func (s *InstallAddonRequest) SetServicesSpec(v string) *InstallAddonRequest {
	s.ServicesSpec = &v
	return s
}

type InstallAddonResponseBody struct {
	// The addon ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login-1.0-W4g****
	AddonId *string `json:"AddonId,omitempty" xml:"AddonId,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B745C159-3155-4B94-95D0-4B73D4D2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallAddonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallAddonResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAddonResponseBody) SetAddonId(v string) *InstallAddonResponseBody {
	s.AddonId = &v
	return s
}

func (s *InstallAddonResponseBody) SetClusterId(v string) *InstallAddonResponseBody {
	s.ClusterId = &v
	return s
}

func (s *InstallAddonResponseBody) SetRequestId(v string) *InstallAddonResponseBody {
	s.RequestId = &v
	return s
}

type InstallAddonResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallAddonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallAddonResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallAddonResponse) GoString() string {
	return s.String()
}

func (s *InstallAddonResponse) SetHeaders(v map[string]*string) *InstallAddonResponse {
	s.Headers = v
	return s
}

func (s *InstallAddonResponse) SetStatusCode(v int32) *InstallAddonResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallAddonResponse) SetBody(v *InstallAddonResponseBody) *InstallAddonResponse {
	s.Body = v
	return s
}

type InstallSoftwaresRequest struct {
	// The information about the software systems that you want to install.
	AdditionalPackages []*InstallSoftwaresRequestAdditionalPackages `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s InstallSoftwaresRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallSoftwaresRequest) GoString() string {
	return s.String()
}

func (s *InstallSoftwaresRequest) SetAdditionalPackages(v []*InstallSoftwaresRequestAdditionalPackages) *InstallSoftwaresRequest {
	s.AdditionalPackages = v
	return s
}

func (s *InstallSoftwaresRequest) SetClusterId(v string) *InstallSoftwaresRequest {
	s.ClusterId = &v
	return s
}

type InstallSoftwaresRequestAdditionalPackages struct {
	// The software name.
	//
	// example:
	//
	// gromacs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The software version.
	//
	// example:
	//
	// 2024.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s InstallSoftwaresRequestAdditionalPackages) String() string {
	return tea.Prettify(s)
}

func (s InstallSoftwaresRequestAdditionalPackages) GoString() string {
	return s.String()
}

func (s *InstallSoftwaresRequestAdditionalPackages) SetName(v string) *InstallSoftwaresRequestAdditionalPackages {
	s.Name = &v
	return s
}

func (s *InstallSoftwaresRequestAdditionalPackages) SetVersion(v string) *InstallSoftwaresRequestAdditionalPackages {
	s.Version = &v
	return s
}

type InstallSoftwaresShrinkRequest struct {
	// The information about the software systems that you want to install.
	AdditionalPackagesShrink *string `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s InstallSoftwaresShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallSoftwaresShrinkRequest) GoString() string {
	return s.String()
}

func (s *InstallSoftwaresShrinkRequest) SetAdditionalPackagesShrink(v string) *InstallSoftwaresShrinkRequest {
	s.AdditionalPackagesShrink = &v
	return s
}

func (s *InstallSoftwaresShrinkRequest) SetClusterId(v string) *InstallSoftwaresShrinkRequest {
	s.ClusterId = &v
	return s
}

type InstallSoftwaresResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallSoftwaresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallSoftwaresResponseBody) GoString() string {
	return s.String()
}

func (s *InstallSoftwaresResponseBody) SetRequestId(v string) *InstallSoftwaresResponseBody {
	s.RequestId = &v
	return s
}

type InstallSoftwaresResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallSoftwaresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallSoftwaresResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallSoftwaresResponse) GoString() string {
	return s.String()
}

func (s *InstallSoftwaresResponse) SetHeaders(v map[string]*string) *InstallSoftwaresResponse {
	s.Headers = v
	return s
}

func (s *InstallSoftwaresResponse) SetStatusCode(v int32) *InstallSoftwaresResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallSoftwaresResponse) SetBody(v *InstallSoftwaresResponseBody) *InstallSoftwaresResponse {
	s.Body = v
	return s
}

type ListAddonTemplatesRequest struct {
	// The addon names.
	AddonNames []*string `json:"AddonNames,omitempty" xml:"AddonNames,omitempty" type:"Repeated"`
	// The cluster type. Valid values:
	//
	// 	- Standard
	//
	// 	- Serverless
	//
	// example:
	//
	// Standard
	ClusterCategory *string `json:"ClusterCategory,omitempty" xml:"ClusterCategory,omitempty"`
	// The page number of the page to return. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAddonTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAddonTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListAddonTemplatesRequest) SetAddonNames(v []*string) *ListAddonTemplatesRequest {
	s.AddonNames = v
	return s
}

func (s *ListAddonTemplatesRequest) SetClusterCategory(v string) *ListAddonTemplatesRequest {
	s.ClusterCategory = &v
	return s
}

func (s *ListAddonTemplatesRequest) SetPageNumber(v int64) *ListAddonTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAddonTemplatesRequest) SetPageSize(v int64) *ListAddonTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAddonTemplatesRequest) SetRegionId(v string) *ListAddonTemplatesRequest {
	s.RegionId = &v
	return s
}

type ListAddonTemplatesResponseBody struct {
	// The information about the addon templates.
	Addons []*ListAddonTemplatesResponseBodyAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAddonTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAddonTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddonTemplatesResponseBody) SetAddons(v []*ListAddonTemplatesResponseBodyAddons) *ListAddonTemplatesResponseBody {
	s.Addons = v
	return s
}

func (s *ListAddonTemplatesResponseBody) SetPageNumber(v int64) *ListAddonTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAddonTemplatesResponseBody) SetPageSize(v int64) *ListAddonTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAddonTemplatesResponseBody) SetRequestId(v string) *ListAddonTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAddonTemplatesResponseBody) SetTotalCount(v int32) *ListAddonTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAddonTemplatesResponseBodyAddons struct {
	// The addon description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The addon label
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The addon name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The addon version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAddonTemplatesResponseBodyAddons) String() string {
	return tea.Prettify(s)
}

func (s ListAddonTemplatesResponseBodyAddons) GoString() string {
	return s.String()
}

func (s *ListAddonTemplatesResponseBodyAddons) SetDescription(v string) *ListAddonTemplatesResponseBodyAddons {
	s.Description = &v
	return s
}

func (s *ListAddonTemplatesResponseBodyAddons) SetLabel(v string) *ListAddonTemplatesResponseBodyAddons {
	s.Label = &v
	return s
}

func (s *ListAddonTemplatesResponseBodyAddons) SetName(v string) *ListAddonTemplatesResponseBodyAddons {
	s.Name = &v
	return s
}

func (s *ListAddonTemplatesResponseBodyAddons) SetVersion(v string) *ListAddonTemplatesResponseBodyAddons {
	s.Version = &v
	return s
}

type ListAddonTemplatesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAddonTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAddonTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAddonTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListAddonTemplatesResponse) SetHeaders(v map[string]*string) *ListAddonTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListAddonTemplatesResponse) SetStatusCode(v int32) *ListAddonTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAddonTemplatesResponse) SetBody(v *ListAddonTemplatesResponseBody) *ListAddonTemplatesResponse {
	s.Body = v
	return s
}

type ListAddonsRequest struct {
	// The addon IDs.
	AddonIds []*string `json:"AddonIds,omitempty" xml:"AddonIds,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number of the page to return. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAddonsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAddonsRequest) GoString() string {
	return s.String()
}

func (s *ListAddonsRequest) SetAddonIds(v []*string) *ListAddonsRequest {
	s.AddonIds = v
	return s
}

func (s *ListAddonsRequest) SetClusterId(v string) *ListAddonsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAddonsRequest) SetPageNumber(v int32) *ListAddonsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAddonsRequest) SetPageSize(v int32) *ListAddonsRequest {
	s.PageSize = &v
	return s
}

type ListAddonsShrinkRequest struct {
	// The addon IDs.
	AddonIdsShrink *string `json:"AddonIds,omitempty" xml:"AddonIds,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number of the page to return. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAddonsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAddonsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAddonsShrinkRequest) SetAddonIdsShrink(v string) *ListAddonsShrinkRequest {
	s.AddonIdsShrink = &v
	return s
}

func (s *ListAddonsShrinkRequest) SetClusterId(v string) *ListAddonsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAddonsShrinkRequest) SetPageNumber(v int32) *ListAddonsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAddonsShrinkRequest) SetPageSize(v int32) *ListAddonsShrinkRequest {
	s.PageSize = &v
	return s
}

type ListAddonsResponseBody struct {
	// The information about the addons.
	Addons []*ListAddonsResponseBodyAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number of the returned page. Default value: 1
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAddonsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBody) SetAddons(v []*ListAddonsResponseBodyAddons) *ListAddonsResponseBody {
	s.Addons = v
	return s
}

func (s *ListAddonsResponseBody) SetPageNumber(v int32) *ListAddonsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAddonsResponseBody) SetPageSize(v int32) *ListAddonsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAddonsResponseBody) SetRequestId(v string) *ListAddonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAddonsResponseBody) SetTotalCount(v int32) *ListAddonsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAddonsResponseBodyAddons struct {
	// The addon ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login-1.0-W4g****
	AddonId *string `json:"AddonId,omitempty" xml:"AddonId,omitempty"`
	// The addon description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the addon was installed.
	//
	// example:
	//
	// 2024-08-22 18:11:17
	InstallTime *string `json:"InstallTime,omitempty" xml:"InstallTime,omitempty"`
	// The addon label.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The addon name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The addon state.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The addon version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAddonsResponseBodyAddons) String() string {
	return tea.Prettify(s)
}

func (s ListAddonsResponseBodyAddons) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyAddons) SetAddonId(v string) *ListAddonsResponseBodyAddons {
	s.AddonId = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetDescription(v string) *ListAddonsResponseBodyAddons {
	s.Description = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetInstallTime(v string) *ListAddonsResponseBodyAddons {
	s.InstallTime = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetLabel(v string) *ListAddonsResponseBodyAddons {
	s.Label = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetName(v string) *ListAddonsResponseBodyAddons {
	s.Name = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetStatus(v string) *ListAddonsResponseBodyAddons {
	s.Status = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetVersion(v string) *ListAddonsResponseBodyAddons {
	s.Version = &v
	return s
}

type ListAddonsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAddonsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAddonsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAddonsResponse) GoString() string {
	return s.String()
}

func (s *ListAddonsResponse) SetHeaders(v map[string]*string) *ListAddonsResponse {
	s.Headers = v
	return s
}

func (s *ListAddonsResponse) SetStatusCode(v int32) *ListAddonsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAddonsResponse) SetBody(v *ListAddonsResponseBody) *ListAddonsResponse {
	s.Body = v
	return s
}

type ListAvailableFileSystemsRequest struct {
	// The page number of the page to return. Page starts from page 1. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50. Default value: 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAvailableFileSystemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableFileSystemsRequest) SetPageNumber(v int32) *ListAvailableFileSystemsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAvailableFileSystemsRequest) SetPageSize(v int32) *ListAvailableFileSystemsRequest {
	s.PageSize = &v
	return s
}

type ListAvailableFileSystemsResponseBody struct {
	// The file systems.
	FileSystemList []*ListAvailableFileSystemsResponseBodyFileSystemList `json:"FileSystemList,omitempty" xml:"FileSystemList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BF4E8AB1-02A3-5ECF-87CC-3AB7BE98**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 65
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAvailableFileSystemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableFileSystemsResponseBody) SetFileSystemList(v []*ListAvailableFileSystemsResponseBodyFileSystemList) *ListAvailableFileSystemsResponseBody {
	s.FileSystemList = v
	return s
}

func (s *ListAvailableFileSystemsResponseBody) SetPageNumber(v int32) *ListAvailableFileSystemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBody) SetPageSize(v int32) *ListAvailableFileSystemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBody) SetRequestId(v string) *ListAvailableFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBody) SetTotalCount(v int32) *ListAvailableFileSystemsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAvailableFileSystemsResponseBodyFileSystemList struct {
	// The time when the file system was created.
	//
	// example:
	//
	// 2024-7-29 15:43:53
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the file system.
	//
	// example:
	//
	// 2fa0248***
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The type of the file system. Valid values:
	//
	// 	- standard: general-purpose network-attached storage (NAS) file system
	//
	// 	- extreme: extreme NAS file system
	//
	// example:
	//
	// cpfs
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The mount targets of the file system.
	MountTargetList []*ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList `json:"MountTargetList,omitempty" xml:"MountTargetList,omitempty" type:"Repeated"`
	// The protocol type of the file system. Valid values:
	//
	// 	- nfs
	//
	// 	- smb
	//
	// 	- cpfs
	//
	// example:
	//
	// cpfs
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The state of the file system. Valid values:
	//
	// 	- Pending: The file system is processing a task.
	//
	// 	- Running: The file system is available. You can perform subsequent operations, such as creating a mount target, only when the file system is in the Running state.
	//
	// 	- Stopped: The file system is unavailable.
	//
	// 	- Extending: The file system is being scaled out.
	//
	// 	- Stopping: The file system is being stopped.
	//
	// 	- Deleting: The file system is being deleted.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage type of the file system.
	//
	// 	- Valid values if FileSystemType is set to standard: Capacity and Performance.
	//
	// 	- Valid values if FileSystemType is set to extreme: standard and advance.
	//
	// example:
	//
	// Performance
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp132kgui8n0targb****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListAvailableFileSystemsResponseBodyFileSystemList) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableFileSystemsResponseBodyFileSystemList) GoString() string {
	return s.String()
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) SetCreateTime(v string) *ListAvailableFileSystemsResponseBodyFileSystemList {
	s.CreateTime = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) SetFileSystemId(v string) *ListAvailableFileSystemsResponseBodyFileSystemList {
	s.FileSystemId = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) SetFileSystemType(v string) *ListAvailableFileSystemsResponseBodyFileSystemList {
	s.FileSystemType = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) SetMountTargetList(v []*ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) *ListAvailableFileSystemsResponseBodyFileSystemList {
	s.MountTargetList = v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) SetProtocolType(v string) *ListAvailableFileSystemsResponseBodyFileSystemList {
	s.ProtocolType = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) SetStatus(v string) *ListAvailableFileSystemsResponseBodyFileSystemList {
	s.Status = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) SetStorageType(v string) *ListAvailableFileSystemsResponseBodyFileSystemList {
	s.StorageType = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) SetVpcId(v string) *ListAvailableFileSystemsResponseBodyFileSystemList {
	s.VpcId = &v
	return s
}

type ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList struct {
	// The address of the mount target.
	//
	// example:
	//
	// c0967****.cn-hangzhou.cpfs.nas.aliyuncs.com
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The network type. Valid values: Valid values:
	//
	// 	- vpc
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The state of the mount target. Valid values:
	//
	// 	- Active: The mount target is available.
	//
	// 	- Inactive: The mount target is unavailable.
	//
	// 	- Pending: The mount target is being mounted.
	//
	// 	- Deleting: The mount target is being deleted.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2ze0c41hwu7lc29ris***
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-8vbvb34rtyh6xb3zrehs1****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) GoString() string {
	return s.String()
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) SetMountTargetDomain(v string) *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList {
	s.MountTargetDomain = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) SetNetworkType(v string) *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList {
	s.NetworkType = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) SetStatus(v string) *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList {
	s.Status = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) SetVSwitchId(v string) *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList {
	s.VSwitchId = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) SetVpcId(v string) *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList {
	s.VpcId = &v
	return s
}

type ListAvailableFileSystemsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvailableFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvailableFileSystemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableFileSystemsResponse) SetHeaders(v map[string]*string) *ListAvailableFileSystemsResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableFileSystemsResponse) SetStatusCode(v int32) *ListAvailableFileSystemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableFileSystemsResponse) SetBody(v *ListAvailableFileSystemsResponseBody) *ListAvailableFileSystemsResponse {
	s.Body = v
	return s
}

type ListAvailableImagesRequest struct {
	// The information about the domain account service.
	DirectoryService *ListAvailableImagesRequestDirectoryService `json:"DirectoryService,omitempty" xml:"DirectoryService,omitempty" type:"Struct"`
	// Specifies whether to return images in which hyper-threading is enabled.
	//
	// example:
	//
	// true
	EnableHT *bool `json:"EnableHT,omitempty" xml:"EnableHT,omitempty"`
	// The network type of the images that you want to query.
	//
	// example:
	//
	// vpc
	HPCInterConnect *string `json:"HPCInterConnect,omitempty" xml:"HPCInterConnect,omitempty"`
	// The image source. Valid values:
	//
	// 	- system: system image.
	//
	// 	- self: custom image.
	//
	// 	- others: shared image.
	//
	// example:
	//
	// self
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The instance type.
	//
	// >  By default, if you do not specify an instance type, the list of images that are supported by all instance types are queried. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// ecs.c7.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether to return published community images. Valid values:
	//
	// 	- true: returns published community images. If you set the value of this parameter to `true`, the `ImageOwnerAlias` parameter must be set to `others`.
	//
	// 	- false: returns non-community images. The value of the `ImageOwnerAlias` parameter prevails.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsPublic *bool `json:"IsPublic,omitempty" xml:"IsPublic,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50. Default value: 10.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The scheduler information about the images that you want to query.
	Scheduler *ListAvailableImagesRequestScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Struct"`
}

func (s ListAvailableImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableImagesRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableImagesRequest) SetDirectoryService(v *ListAvailableImagesRequestDirectoryService) *ListAvailableImagesRequest {
	s.DirectoryService = v
	return s
}

func (s *ListAvailableImagesRequest) SetEnableHT(v bool) *ListAvailableImagesRequest {
	s.EnableHT = &v
	return s
}

func (s *ListAvailableImagesRequest) SetHPCInterConnect(v string) *ListAvailableImagesRequest {
	s.HPCInterConnect = &v
	return s
}

func (s *ListAvailableImagesRequest) SetImageOwnerAlias(v string) *ListAvailableImagesRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListAvailableImagesRequest) SetInstanceType(v string) *ListAvailableImagesRequest {
	s.InstanceType = &v
	return s
}

func (s *ListAvailableImagesRequest) SetIsPublic(v bool) *ListAvailableImagesRequest {
	s.IsPublic = &v
	return s
}

func (s *ListAvailableImagesRequest) SetPageNumber(v int32) *ListAvailableImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAvailableImagesRequest) SetPageSize(v int32) *ListAvailableImagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAvailableImagesRequest) SetScheduler(v *ListAvailableImagesRequestScheduler) *ListAvailableImagesRequest {
	s.Scheduler = v
	return s
}

type ListAvailableImagesRequestDirectoryService struct {
	// The type of the domain account.
	//
	// example:
	//
	// NIS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the domain account service.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAvailableImagesRequestDirectoryService) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableImagesRequestDirectoryService) GoString() string {
	return s.String()
}

func (s *ListAvailableImagesRequestDirectoryService) SetType(v string) *ListAvailableImagesRequestDirectoryService {
	s.Type = &v
	return s
}

func (s *ListAvailableImagesRequestDirectoryService) SetVersion(v string) *ListAvailableImagesRequestDirectoryService {
	s.Version = &v
	return s
}

type ListAvailableImagesRequestScheduler struct {
	// The scheduler type.
	//
	// example:
	//
	// SLURM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The scheduler version.
	//
	// example:
	//
	// 22.05.8
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAvailableImagesRequestScheduler) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableImagesRequestScheduler) GoString() string {
	return s.String()
}

func (s *ListAvailableImagesRequestScheduler) SetType(v string) *ListAvailableImagesRequestScheduler {
	s.Type = &v
	return s
}

func (s *ListAvailableImagesRequestScheduler) SetVersion(v string) *ListAvailableImagesRequestScheduler {
	s.Version = &v
	return s
}

type ListAvailableImagesShrinkRequest struct {
	// The information about the domain account service.
	DirectoryServiceShrink *string `json:"DirectoryService,omitempty" xml:"DirectoryService,omitempty"`
	// Specifies whether to return images in which hyper-threading is enabled.
	//
	// example:
	//
	// true
	EnableHT *bool `json:"EnableHT,omitempty" xml:"EnableHT,omitempty"`
	// The network type of the images that you want to query.
	//
	// example:
	//
	// vpc
	HPCInterConnect *string `json:"HPCInterConnect,omitempty" xml:"HPCInterConnect,omitempty"`
	// The image source. Valid values:
	//
	// 	- system: system image.
	//
	// 	- self: custom image.
	//
	// 	- others: shared image.
	//
	// example:
	//
	// self
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The instance type.
	//
	// >  By default, if you do not specify an instance type, the list of images that are supported by all instance types are queried. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// ecs.c7.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether to return published community images. Valid values:
	//
	// 	- true: returns published community images. If you set the value of this parameter to `true`, the `ImageOwnerAlias` parameter must be set to `others`.
	//
	// 	- false: returns non-community images. The value of the `ImageOwnerAlias` parameter prevails.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsPublic *bool `json:"IsPublic,omitempty" xml:"IsPublic,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50. Default value: 10.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The scheduler information about the images that you want to query.
	SchedulerShrink *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
}

func (s ListAvailableImagesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableImagesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableImagesShrinkRequest) SetDirectoryServiceShrink(v string) *ListAvailableImagesShrinkRequest {
	s.DirectoryServiceShrink = &v
	return s
}

func (s *ListAvailableImagesShrinkRequest) SetEnableHT(v bool) *ListAvailableImagesShrinkRequest {
	s.EnableHT = &v
	return s
}

func (s *ListAvailableImagesShrinkRequest) SetHPCInterConnect(v string) *ListAvailableImagesShrinkRequest {
	s.HPCInterConnect = &v
	return s
}

func (s *ListAvailableImagesShrinkRequest) SetImageOwnerAlias(v string) *ListAvailableImagesShrinkRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListAvailableImagesShrinkRequest) SetInstanceType(v string) *ListAvailableImagesShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *ListAvailableImagesShrinkRequest) SetIsPublic(v bool) *ListAvailableImagesShrinkRequest {
	s.IsPublic = &v
	return s
}

func (s *ListAvailableImagesShrinkRequest) SetPageNumber(v int32) *ListAvailableImagesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAvailableImagesShrinkRequest) SetPageSize(v int32) *ListAvailableImagesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAvailableImagesShrinkRequest) SetSchedulerShrink(v string) *ListAvailableImagesShrinkRequest {
	s.SchedulerShrink = &v
	return s
}

type ListAvailableImagesResponseBody struct {
	// The information about the images.
	Images []*ListAvailableImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// 	- Pages start from page 1.
	//
	// 	- Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAvailableImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableImagesResponseBody) SetImages(v []*ListAvailableImagesResponseBodyImages) *ListAvailableImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListAvailableImagesResponseBody) SetPageNumber(v string) *ListAvailableImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAvailableImagesResponseBody) SetPageSize(v string) *ListAvailableImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAvailableImagesResponseBody) SetRequestId(v string) *ListAvailableImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableImagesResponseBody) SetSuccess(v bool) *ListAvailableImagesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAvailableImagesResponseBody) SetTotalCount(v string) *ListAvailableImagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAvailableImagesResponseBodyImages struct {
	// The OS architecture of the image. Valid values:
	//
	// 	- x86_64
	//
	// 	- arm64
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The boot mode of the image. Valid values:
	//
	// 	- BIOS: Basic Input/Output System (BIOS).
	//
	// 	- UEFI: Unified Extensible Firmware Interface (UEFI).
	//
	// >  When you change the OS boot mode of an instance, you must make sure that the boot mode matches the boot mode of the associated image. Otherwise, the instance fails to be booted.
	//
	// example:
	//
	// BIOS
	BootMode *string `json:"BootMode,omitempty" xml:"BootMode,omitempty"`
	// The image description.
	//
	// example:
	//
	// ExampleDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image ID.
	//
	// example:
	//
	// centos_7_06_64_20G_alibase_2019071****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// CHESS5V5.0.27
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image source. Valid values:
	//
	// 	- system: system image.
	//
	// 	- self: custom image.
	//
	// 	- others: shared image.
	//
	// example:
	//
	// self
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The OS name in Chinese.
	OSName *string `json:"OSName,omitempty" xml:"OSName,omitempty"`
	// The OS name in English.
	//
	// example:
	//
	// CentOS  7.9 64 bit
	OSNameEn *string `json:"OSNameEn,omitempty" xml:"OSNameEn,omitempty"`
	// The OS. Valid values:
	//
	// 	- CentOS
	//
	// 	- windows
	//
	// example:
	//
	// windows
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The image size. Unit: GiB
	//
	// example:
	//
	// 40
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListAvailableImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListAvailableImagesResponseBodyImages) SetArchitecture(v string) *ListAvailableImagesResponseBodyImages {
	s.Architecture = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetBootMode(v string) *ListAvailableImagesResponseBodyImages {
	s.BootMode = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetDescription(v string) *ListAvailableImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetImageId(v string) *ListAvailableImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetImageName(v string) *ListAvailableImagesResponseBodyImages {
	s.ImageName = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetImageOwnerAlias(v string) *ListAvailableImagesResponseBodyImages {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetOSName(v string) *ListAvailableImagesResponseBodyImages {
	s.OSName = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetOSNameEn(v string) *ListAvailableImagesResponseBodyImages {
	s.OSNameEn = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetPlatform(v string) *ListAvailableImagesResponseBodyImages {
	s.Platform = &v
	return s
}

func (s *ListAvailableImagesResponseBodyImages) SetSize(v string) *ListAvailableImagesResponseBodyImages {
	s.Size = &v
	return s
}

type ListAvailableImagesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvailableImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvailableImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableImagesResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableImagesResponse) SetHeaders(v map[string]*string) *ListAvailableImagesResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableImagesResponse) SetStatusCode(v int32) *ListAvailableImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableImagesResponse) SetBody(v *ListAvailableImagesResponseBody) *ListAvailableImagesResponse {
	s.Body = v
	return s
}

type ListClustersRequest struct {
	// The cluster IDs. You can specify up to 20 IDs.
	ClusterIds []*string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty" type:"Repeated"`
	// The cluster names. You can specify up to 20 names.
	ClusterNames []*string `json:"ClusterNames,omitempty" xml:"ClusterNames,omitempty" type:"Repeated"`
	// The page number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 10 to 100. Default value: 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) SetClusterIds(v []*string) *ListClustersRequest {
	s.ClusterIds = v
	return s
}

func (s *ListClustersRequest) SetClusterNames(v []*string) *ListClustersRequest {
	s.ClusterNames = v
	return s
}

func (s *ListClustersRequest) SetPageNumber(v int32) *ListClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClustersRequest) SetPageSize(v int32) *ListClustersRequest {
	s.PageSize = &v
	return s
}

type ListClustersShrinkRequest struct {
	// The cluster IDs. You can specify up to 20 IDs.
	ClusterIdsShrink *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	// The cluster names. You can specify up to 20 names.
	ClusterNamesShrink *string `json:"ClusterNames,omitempty" xml:"ClusterNames,omitempty"`
	// The page number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 10 to 100. Default value: 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClustersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListClustersShrinkRequest) SetClusterIdsShrink(v string) *ListClustersShrinkRequest {
	s.ClusterIdsShrink = &v
	return s
}

func (s *ListClustersShrinkRequest) SetClusterNamesShrink(v string) *ListClustersShrinkRequest {
	s.ClusterNamesShrink = &v
	return s
}

func (s *ListClustersShrinkRequest) SetPageNumber(v int32) *ListClustersShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClustersShrinkRequest) SetPageSize(v int32) *ListClustersShrinkRequest {
	s.PageSize = &v
	return s
}

type ListClustersResponseBody struct {
	// The list of clusters.
	Clusters []*ListClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) SetClusters(v []*ListClustersResponseBodyClusters) *ListClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *ListClustersResponseBody) SetPageNumber(v string) *ListClustersResponseBody {
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
	// The information about installed software in the cluster.
	AdditionalPackages []*ListClustersResponseBodyClustersAdditionalPackages `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty" type:"Repeated"`
	// The information about the addons in the cluster.
	Addons []*ListClustersResponseBodyClustersAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
	// The cluster type. Valid values:
	//
	// 	- Standard
	//
	// 	- Serverless
	//
	// example:
	//
	// Standard
	ClusterCategory *string `json:"ClusterCategory,omitempty" xml:"ClusterCategory,omitempty"`
	// The time when the cluster was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2024-08-06T12:43:01.000Z
	ClusterCreateTime *string `json:"ClusterCreateTime,omitempty" xml:"ClusterCreateTime,omitempty"`
	// The logon credential type of the cluster. Valid values:
	//
	// 	- password: requires passwords for logons.
	//
	// 	- keypair: requires key pairs for logons.
	ClusterCredentials []*string `json:"ClusterCredentials,omitempty" xml:"ClusterCredentials,omitempty" type:"Repeated"`
	// The post-processing script used by the cluster.
	ClusterCustomConfiguration *ListClustersResponseBodyClustersClusterCustomConfiguration `json:"ClusterCustomConfiguration,omitempty" xml:"ClusterCustomConfiguration,omitempty" type:"Struct"`
	// The cluster description.
	//
	// example:
	//
	// Demo
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-VMKe******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The deployment type of the cluster. Valid values:
	//
	// 	- Integrated: public cloud
	//
	// 	- Hybrid: hybrid cloud
	//
	// 	- Custom: a custom cluster
	//
	// example:
	//
	// Integrated
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// The time when the cluster was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2024-08-06T12:43:01.000Z
	ClusterModifyTime *string `json:"ClusterModifyTime,omitempty" xml:"ClusterModifyTime,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// slurm22.05.8-cluster-20240227
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The cluster state. Valid values:
	//
	// 	- uninit: The cluster is being installed.
	//
	// 	- creating: The cluster is being created.
	//
	// 	- initing: The cluster is being initialized.
	//
	// 	- running: The cluster is running.
	//
	// 	- Releasing: The cluster is being released.
	//
	// 	- stopping: The cluster is being stopped.
	//
	// 	- stopped: The cluster is stopped.
	//
	// 	- exception: The cluster has run into an exception.
	//
	// 	- pending: The cluster is waiting to be configured.
	//
	// example:
	//
	// running
	ClusterStatus *string `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	// The vCPU-hour usage of the cluster.
	//
	// example:
	//
	// 1000
	ClusterUsedCoreTime *float32 `json:"ClusterUsedCoreTime,omitempty" xml:"ClusterUsedCoreTime,omitempty"`
	// The ID of the vSwitch used by the cluster.
	//
	// example:
	//
	// vsw-f8za5p0mwzgdu3wgx****
	ClusterVSwitchId *string `json:"ClusterVSwitchId,omitempty" xml:"ClusterVSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) used by the cluster.
	//
	// example:
	//
	// vpc-m5efjevmclc0xdmys****
	ClusterVpcId *string `json:"ClusterVpcId,omitempty" xml:"ClusterVpcId,omitempty"`
	// Indicates whether deletion protection is enabled for the cluster. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The Elastic High Performance Computing (E-HPC) version.
	//
	// example:
	//
	// 2.0.0
	EhpcVersion *string `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	// The configurations of the cluster management node.
	Manager *ListClustersResponseBodyClustersManager `json:"Manager,omitempty" xml:"Manager,omitempty" type:"Struct"`
	// The maximum total number of vCPUs used by the compute nodes that can be managed by the cluster.
	//
	// example:
	//
	// 10000
	MaxCoreCount *int64 `json:"MaxCoreCount,omitempty" xml:"MaxCoreCount,omitempty"`
	// The maximum number of compute nodes that can be managed by the cluster.
	//
	// example:
	//
	// 500
	MaxCount *int64 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The node statistics of the cluster.
	Nodes *ListClustersResponseBodyClustersNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the security group used by the cluster.
	//
	// example:
	//
	// sg-bp13n61xsydodfyg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The user attribute information of the cluster.
	Users *ListClustersResponseBodyClustersUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClusters) SetAdditionalPackages(v []*ListClustersResponseBodyClustersAdditionalPackages) *ListClustersResponseBodyClusters {
	s.AdditionalPackages = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetAddons(v []*ListClustersResponseBodyClustersAddons) *ListClustersResponseBodyClusters {
	s.Addons = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterCategory(v string) *ListClustersResponseBodyClusters {
	s.ClusterCategory = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterCreateTime(v string) *ListClustersResponseBodyClusters {
	s.ClusterCreateTime = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterCredentials(v []*string) *ListClustersResponseBodyClusters {
	s.ClusterCredentials = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterCustomConfiguration(v *ListClustersResponseBodyClustersClusterCustomConfiguration) *ListClustersResponseBodyClusters {
	s.ClusterCustomConfiguration = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterDescription(v string) *ListClustersResponseBodyClusters {
	s.ClusterDescription = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterId(v string) *ListClustersResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterMode(v string) *ListClustersResponseBodyClusters {
	s.ClusterMode = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterModifyTime(v string) *ListClustersResponseBodyClusters {
	s.ClusterModifyTime = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterName(v string) *ListClustersResponseBodyClusters {
	s.ClusterName = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterStatus(v string) *ListClustersResponseBodyClusters {
	s.ClusterStatus = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterUsedCoreTime(v float32) *ListClustersResponseBodyClusters {
	s.ClusterUsedCoreTime = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterVSwitchId(v string) *ListClustersResponseBodyClusters {
	s.ClusterVSwitchId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterVpcId(v string) *ListClustersResponseBodyClusters {
	s.ClusterVpcId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetDeletionProtection(v bool) *ListClustersResponseBodyClusters {
	s.DeletionProtection = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetEhpcVersion(v string) *ListClustersResponseBodyClusters {
	s.EhpcVersion = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetManager(v *ListClustersResponseBodyClustersManager) *ListClustersResponseBodyClusters {
	s.Manager = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetMaxCoreCount(v int64) *ListClustersResponseBodyClusters {
	s.MaxCoreCount = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetMaxCount(v int64) *ListClustersResponseBodyClusters {
	s.MaxCount = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetNodes(v *ListClustersResponseBodyClustersNodes) *ListClustersResponseBodyClusters {
	s.Nodes = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetResourceGroupId(v string) *ListClustersResponseBodyClusters {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetSecurityGroupId(v string) *ListClustersResponseBodyClusters {
	s.SecurityGroupId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetUsers(v *ListClustersResponseBodyClustersUsers) *ListClustersResponseBodyClusters {
	s.Users = v
	return s
}

type ListClustersResponseBodyClustersAdditionalPackages struct {
	// The software name.
	//
	// example:
	//
	// gromacs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The software version.
	//
	// example:
	//
	// 2024.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListClustersResponseBodyClustersAdditionalPackages) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersAdditionalPackages) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersAdditionalPackages) SetName(v string) *ListClustersResponseBodyClustersAdditionalPackages {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyClustersAdditionalPackages) SetVersion(v string) *ListClustersResponseBodyClustersAdditionalPackages {
	s.Version = &v
	return s
}

type ListClustersResponseBodyClustersAddons struct {
	// The addon ID.
	//
	// example:
	//
	// Login-1.0-W2g****
	AddonId *string `json:"AddonId,omitempty" xml:"AddonId,omitempty"`
	// The addon description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The addon label.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The addon name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource configurations of the addon.
	ResourcesSpec *ListClustersResponseBodyClustersAddonsResourcesSpec `json:"ResourcesSpec,omitempty" xml:"ResourcesSpec,omitempty" type:"Struct"`
	// The information about the addon services.
	ServicesSpec []*ListClustersResponseBodyClustersAddonsServicesSpec `json:"ServicesSpec,omitempty" xml:"ServicesSpec,omitempty" type:"Repeated"`
	// The addon state.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The addon version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListClustersResponseBodyClustersAddons) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersAddons) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersAddons) SetAddonId(v string) *ListClustersResponseBodyClustersAddons {
	s.AddonId = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddons) SetDescription(v string) *ListClustersResponseBodyClustersAddons {
	s.Description = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddons) SetLabel(v string) *ListClustersResponseBodyClustersAddons {
	s.Label = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddons) SetName(v string) *ListClustersResponseBodyClustersAddons {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddons) SetResourcesSpec(v *ListClustersResponseBodyClustersAddonsResourcesSpec) *ListClustersResponseBodyClustersAddons {
	s.ResourcesSpec = v
	return s
}

func (s *ListClustersResponseBodyClustersAddons) SetServicesSpec(v []*ListClustersResponseBodyClustersAddonsServicesSpec) *ListClustersResponseBodyClustersAddons {
	s.ServicesSpec = v
	return s
}

func (s *ListClustersResponseBodyClustersAddons) SetStatus(v string) *ListClustersResponseBodyClustersAddons {
	s.Status = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddons) SetVersion(v string) *ListClustersResponseBodyClustersAddons {
	s.Version = &v
	return s
}

type ListClustersResponseBodyClustersAddonsResourcesSpec struct {
	// The instance ID.
	//
	// example:
	//
	// i-bp1bg85d2q6laic8****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The Elastic IP Address (EIP) ID.
	//
	// example:
	//
	// eip-bp1vi9124tbx1g3kr****
	EipInstanceId *string `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
}

func (s ListClustersResponseBodyClustersAddonsResourcesSpec) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersAddonsResourcesSpec) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersAddonsResourcesSpec) SetEcsInstanceId(v string) *ListClustersResponseBodyClustersAddonsResourcesSpec {
	s.EcsInstanceId = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddonsResourcesSpec) SetEipInstanceId(v string) *ListClustersResponseBodyClustersAddonsResourcesSpec {
	s.EipInstanceId = &v
	return s
}

type ListClustersResponseBodyClustersAddonsServicesSpec struct {
	// The service access type.
	//
	// example:
	//
	// URL
	ServiceAccessType *string `json:"ServiceAccessType,omitempty" xml:"ServiceAccessType,omitempty"`
	// The service access URL.
	//
	// example:
	//
	// https://47.96.xx.xx:12008
	ServiceAccessUrl *string `json:"ServiceAccessUrl,omitempty" xml:"ServiceAccessUrl,omitempty"`
	// The service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Web Portal
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListClustersResponseBodyClustersAddonsServicesSpec) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersAddonsServicesSpec) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersAddonsServicesSpec) SetServiceAccessType(v string) *ListClustersResponseBodyClustersAddonsServicesSpec {
	s.ServiceAccessType = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddonsServicesSpec) SetServiceAccessUrl(v string) *ListClustersResponseBodyClustersAddonsServicesSpec {
	s.ServiceAccessUrl = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddonsServicesSpec) SetServiceName(v string) *ListClustersResponseBodyClustersAddonsServicesSpec {
	s.ServiceName = &v
	return s
}

type ListClustersResponseBodyClustersClusterCustomConfiguration struct {
	// The parameters of the post-processing script.
	//
	// example:
	//
	// demo
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The link to the post-processing script.
	//
	// example:
	//
	// https://xxxxx
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterCustomConfiguration) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterCustomConfiguration) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterCustomConfiguration) SetArgs(v string) *ListClustersResponseBodyClustersClusterCustomConfiguration {
	s.Args = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterCustomConfiguration) SetScript(v string) *ListClustersResponseBodyClustersClusterCustomConfiguration {
	s.Script = &v
	return s
}

type ListClustersResponseBodyClustersManager struct {
	// The configurations of the domain name resolution service.
	DNS *ListClustersResponseBodyClustersManagerDNS `json:"DNS,omitempty" xml:"DNS,omitempty" type:"Struct"`
	// The configurations of the directory service.
	DirectoryService *ListClustersResponseBodyClustersManagerDirectoryService `json:"DirectoryService,omitempty" xml:"DirectoryService,omitempty" type:"Struct"`
	// The configurations of the scheduler service.
	Scheduler *ListClustersResponseBodyClustersManagerScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Struct"`
}

func (s ListClustersResponseBodyClustersManager) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersManager) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersManager) SetDNS(v *ListClustersResponseBodyClustersManagerDNS) *ListClustersResponseBodyClustersManager {
	s.DNS = v
	return s
}

func (s *ListClustersResponseBodyClustersManager) SetDirectoryService(v *ListClustersResponseBodyClustersManagerDirectoryService) *ListClustersResponseBodyClustersManager {
	s.DirectoryService = v
	return s
}

func (s *ListClustersResponseBodyClustersManager) SetScheduler(v *ListClustersResponseBodyClustersManagerScheduler) *ListClustersResponseBodyClustersManager {
	s.Scheduler = v
	return s
}

type ListClustersResponseBodyClustersManagerDNS struct {
	// The resolution type.
	//
	// example:
	//
	// NIS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the domain name resolution service.
	//
	// example:
	//
	// 2.31
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListClustersResponseBodyClustersManagerDNS) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersManagerDNS) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersManagerDNS) SetType(v string) *ListClustersResponseBodyClustersManagerDNS {
	s.Type = &v
	return s
}

func (s *ListClustersResponseBodyClustersManagerDNS) SetVersion(v string) *ListClustersResponseBodyClustersManagerDNS {
	s.Version = &v
	return s
}

type ListClustersResponseBodyClustersManagerDirectoryService struct {
	// The type of the domain account.
	//
	// example:
	//
	// NIS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the domain account service.
	//
	// example:
	//
	// 2.31
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListClustersResponseBodyClustersManagerDirectoryService) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersManagerDirectoryService) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersManagerDirectoryService) SetType(v string) *ListClustersResponseBodyClustersManagerDirectoryService {
	s.Type = &v
	return s
}

func (s *ListClustersResponseBodyClustersManagerDirectoryService) SetVersion(v string) *ListClustersResponseBodyClustersManagerDirectoryService {
	s.Version = &v
	return s
}

type ListClustersResponseBodyClustersManagerScheduler struct {
	// The scheduler type.
	//
	// example:
	//
	// SLURM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The scheduler version.
	//
	// example:
	//
	// 22.05.8
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListClustersResponseBodyClustersManagerScheduler) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersManagerScheduler) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersManagerScheduler) SetType(v string) *ListClustersResponseBodyClustersManagerScheduler {
	s.Type = &v
	return s
}

func (s *ListClustersResponseBodyClustersManagerScheduler) SetVersion(v string) *ListClustersResponseBodyClustersManagerScheduler {
	s.Version = &v
	return s
}

type ListClustersResponseBodyClustersNodes struct {
	// The number of malfunctioning compute nodes.
	//
	// example:
	//
	// 0
	AbnormalCounts *int32 `json:"AbnormalCounts,omitempty" xml:"AbnormalCounts,omitempty"`
	// The number of compute nodes that are being created.
	//
	// example:
	//
	// 0
	CreatingCounts *int32 `json:"CreatingCounts,omitempty" xml:"CreatingCounts,omitempty"`
	// The number of running compute nodes.
	//
	// example:
	//
	// 1
	RunningCounts *int32 `json:"RunningCounts,omitempty" xml:"RunningCounts,omitempty"`
}

func (s ListClustersResponseBodyClustersNodes) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersNodes) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersNodes) SetAbnormalCounts(v int32) *ListClustersResponseBodyClustersNodes {
	s.AbnormalCounts = &v
	return s
}

func (s *ListClustersResponseBodyClustersNodes) SetCreatingCounts(v int32) *ListClustersResponseBodyClustersNodes {
	s.CreatingCounts = &v
	return s
}

func (s *ListClustersResponseBodyClustersNodes) SetRunningCounts(v int32) *ListClustersResponseBodyClustersNodes {
	s.RunningCounts = &v
	return s
}

type ListClustersResponseBodyClustersUsers struct {
	// The number of ordinary users.
	//
	// example:
	//
	// 2
	NormalCounts *int32 `json:"NormalCounts,omitempty" xml:"NormalCounts,omitempty"`
	// The number of administrators.
	//
	// example:
	//
	// 2
	SudoCounts *int32 `json:"SudoCounts,omitempty" xml:"SudoCounts,omitempty"`
}

func (s ListClustersResponseBodyClustersUsers) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersUsers) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersUsers) SetNormalCounts(v int32) *ListClustersResponseBodyClustersUsers {
	s.NormalCounts = &v
	return s
}

func (s *ListClustersResponseBodyClustersUsers) SetSudoCounts(v int32) *ListClustersResponseBodyClustersUsers {
	s.SudoCounts = &v
	return s
}

type ListClustersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListClustersResponse) SetStatusCode(v int32) *ListClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClustersResponse) SetBody(v *ListClustersResponseBody) *ListClustersResponse {
	s.Body = v
	return s
}

type ListCommonLogsRequest struct {
	// The action types.
	ActionName []*string `json:"ActionName,omitempty" xml:"ActionName,omitempty" type:"Repeated"`
	// The action status. Logs associated with the specific action status are returned.
	//
	// Valid values:
	//
	// 	- Finished: The action is completed.
	//
	// 	- Failed: The action failed.
	//
	// 	- InProgress: The action is being performed.
	//
	// example:
	//
	// Finished
	ActionStatus *string `json:"ActionStatus,omitempty" xml:"ActionStatus,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The start time of the time range. The time is a timestamp. This value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1703821542
	From *int64 `json:"From,omitempty" xml:"From,omitempty"`
	// Specifies whether to display results in reverse order.
	//
	// Default value: true
	//
	// example:
	//
	// true
	IsReverse *bool `json:"IsReverse,omitempty" xml:"IsReverse,omitempty"`
	// The request ID of the action. Logs associated with the specific request ID are returned.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	// The log type. Logs of the specific type are returned.
	//
	// example:
	//
	// Operation
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The account ID of the operator.
	//
	// example:
	//
	// 137***
	OperatorUid *string `json:"OperatorUid,omitempty" xml:"OperatorUid,omitempty"`
	// The page number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the resource involved in the action. Logs associated with the specific resource are returned. This parameter is not recommended.
	//
	// example:
	//
	// i-abc***
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The end time of the time range. The time is a timestamp. This value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1703821666
	To *int64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s ListCommonLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCommonLogsRequest) GoString() string {
	return s.String()
}

func (s *ListCommonLogsRequest) SetActionName(v []*string) *ListCommonLogsRequest {
	s.ActionName = v
	return s
}

func (s *ListCommonLogsRequest) SetActionStatus(v string) *ListCommonLogsRequest {
	s.ActionStatus = &v
	return s
}

func (s *ListCommonLogsRequest) SetClusterId(v string) *ListCommonLogsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCommonLogsRequest) SetFrom(v int64) *ListCommonLogsRequest {
	s.From = &v
	return s
}

func (s *ListCommonLogsRequest) SetIsReverse(v bool) *ListCommonLogsRequest {
	s.IsReverse = &v
	return s
}

func (s *ListCommonLogsRequest) SetLogRequestId(v string) *ListCommonLogsRequest {
	s.LogRequestId = &v
	return s
}

func (s *ListCommonLogsRequest) SetLogType(v string) *ListCommonLogsRequest {
	s.LogType = &v
	return s
}

func (s *ListCommonLogsRequest) SetOperatorUid(v string) *ListCommonLogsRequest {
	s.OperatorUid = &v
	return s
}

func (s *ListCommonLogsRequest) SetPageNumber(v int32) *ListCommonLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCommonLogsRequest) SetPageSize(v int32) *ListCommonLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCommonLogsRequest) SetResource(v string) *ListCommonLogsRequest {
	s.Resource = &v
	return s
}

func (s *ListCommonLogsRequest) SetTo(v int64) *ListCommonLogsRequest {
	s.To = &v
	return s
}

type ListCommonLogsShrinkRequest struct {
	// The action types.
	ActionNameShrink *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// The action status. Logs associated with the specific action status are returned.
	//
	// Valid values:
	//
	// 	- Finished: The action is completed.
	//
	// 	- Failed: The action failed.
	//
	// 	- InProgress: The action is being performed.
	//
	// example:
	//
	// Finished
	ActionStatus *string `json:"ActionStatus,omitempty" xml:"ActionStatus,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The start time of the time range. The time is a timestamp. This value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1703821542
	From *int64 `json:"From,omitempty" xml:"From,omitempty"`
	// Specifies whether to display results in reverse order.
	//
	// Default value: true
	//
	// example:
	//
	// true
	IsReverse *bool `json:"IsReverse,omitempty" xml:"IsReverse,omitempty"`
	// The request ID of the action. Logs associated with the specific request ID are returned.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	// The log type. Logs of the specific type are returned.
	//
	// example:
	//
	// Operation
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The account ID of the operator.
	//
	// example:
	//
	// 137***
	OperatorUid *string `json:"OperatorUid,omitempty" xml:"OperatorUid,omitempty"`
	// The page number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the resource involved in the action. Logs associated with the specific resource are returned. This parameter is not recommended.
	//
	// example:
	//
	// i-abc***
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The end time of the time range. The time is a timestamp. This value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1703821666
	To *int64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s ListCommonLogsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCommonLogsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCommonLogsShrinkRequest) SetActionNameShrink(v string) *ListCommonLogsShrinkRequest {
	s.ActionNameShrink = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetActionStatus(v string) *ListCommonLogsShrinkRequest {
	s.ActionStatus = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetClusterId(v string) *ListCommonLogsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetFrom(v int64) *ListCommonLogsShrinkRequest {
	s.From = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetIsReverse(v bool) *ListCommonLogsShrinkRequest {
	s.IsReverse = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetLogRequestId(v string) *ListCommonLogsShrinkRequest {
	s.LogRequestId = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetLogType(v string) *ListCommonLogsShrinkRequest {
	s.LogType = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetOperatorUid(v string) *ListCommonLogsShrinkRequest {
	s.OperatorUid = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetPageNumber(v int32) *ListCommonLogsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetPageSize(v int32) *ListCommonLogsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetResource(v string) *ListCommonLogsShrinkRequest {
	s.Resource = &v
	return s
}

func (s *ListCommonLogsShrinkRequest) SetTo(v int64) *ListCommonLogsShrinkRequest {
	s.To = &v
	return s
}

type ListCommonLogsResponseBody struct {
	// The brief information of operation logs.
	Logs []*ListCommonLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 464E9919-D04F-4D1D-B375-15989492****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 137***
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListCommonLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCommonLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommonLogsResponseBody) SetLogs(v []*ListCommonLogsResponseBodyLogs) *ListCommonLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListCommonLogsResponseBody) SetPageNumber(v int32) *ListCommonLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCommonLogsResponseBody) SetPageSize(v int32) *ListCommonLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCommonLogsResponseBody) SetRequestId(v string) *ListCommonLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommonLogsResponseBody) SetTotalCount(v int32) *ListCommonLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCommonLogsResponseBody) SetUid(v string) *ListCommonLogsResponseBody {
	s.Uid = &v
	return s
}

type ListCommonLogsResponseBodyLogs struct {
	// The name of the action corresponding to the log.
	//
	// example:
	//
	// CreaterCluster
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-9T3xPNezoS
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The log type.
	//
	// example:
	//
	// Operation
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The message of the log.
	//
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The account ID of the operator.
	//
	// example:
	//
	// 137***
	OperatorUid *string `json:"OperatorUid,omitempty" xml:"OperatorUid,omitempty"`
	// The request ID associated with the action that generated the log.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The action state in the log. Valid values:
	//
	// 	- InProgress: The action is being performed.
	//
	// 	- Finished: The action is completed.
	//
	// 	- Failed: The action failed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The involved resource.
	//
	// example:
	//
	// i-abc***
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The time when the log was generated.
	//
	// example:
	//
	// 2024-08-22 14:21:54
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s ListCommonLogsResponseBodyLogs) String() string {
	return tea.Prettify(s)
}

func (s ListCommonLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListCommonLogsResponseBodyLogs) SetAction(v string) *ListCommonLogsResponseBodyLogs {
	s.Action = &v
	return s
}

func (s *ListCommonLogsResponseBodyLogs) SetClusterId(v string) *ListCommonLogsResponseBodyLogs {
	s.ClusterId = &v
	return s
}

func (s *ListCommonLogsResponseBodyLogs) SetLogType(v string) *ListCommonLogsResponseBodyLogs {
	s.LogType = &v
	return s
}

func (s *ListCommonLogsResponseBodyLogs) SetMessage(v string) *ListCommonLogsResponseBodyLogs {
	s.Message = &v
	return s
}

func (s *ListCommonLogsResponseBodyLogs) SetOperatorUid(v string) *ListCommonLogsResponseBodyLogs {
	s.OperatorUid = &v
	return s
}

func (s *ListCommonLogsResponseBodyLogs) SetRequestId(v string) *ListCommonLogsResponseBodyLogs {
	s.RequestId = &v
	return s
}

func (s *ListCommonLogsResponseBodyLogs) SetStatus(v string) *ListCommonLogsResponseBodyLogs {
	s.Status = &v
	return s
}

func (s *ListCommonLogsResponseBodyLogs) SetTarget(v string) *ListCommonLogsResponseBodyLogs {
	s.Target = &v
	return s
}

func (s *ListCommonLogsResponseBodyLogs) SetTime(v string) *ListCommonLogsResponseBodyLogs {
	s.Time = &v
	return s
}

type ListCommonLogsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCommonLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCommonLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCommonLogsResponse) GoString() string {
	return s.String()
}

func (s *ListCommonLogsResponse) SetHeaders(v map[string]*string) *ListCommonLogsResponse {
	s.Headers = v
	return s
}

func (s *ListCommonLogsResponse) SetStatusCode(v int32) *ListCommonLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCommonLogsResponse) SetBody(v *ListCommonLogsResponseBody) *ListCommonLogsResponse {
	s.Body = v
	return s
}

type ListInstalledSoftwaresRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInstalledSoftwaresRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwaresRequest) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwaresRequest) SetClusterId(v string) *ListInstalledSoftwaresRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInstalledSoftwaresRequest) SetPageNumber(v string) *ListInstalledSoftwaresRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstalledSoftwaresRequest) SetPageSize(v string) *ListInstalledSoftwaresRequest {
	s.PageSize = &v
	return s
}

type ListInstalledSoftwaresResponseBody struct {
	// The list of installed software.
	AdditionalPackages *ListInstalledSoftwaresResponseBodyAdditionalPackages `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstalledSoftwaresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwaresResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwaresResponseBody) SetAdditionalPackages(v *ListInstalledSoftwaresResponseBodyAdditionalPackages) *ListInstalledSoftwaresResponseBody {
	s.AdditionalPackages = v
	return s
}

func (s *ListInstalledSoftwaresResponseBody) SetPageNumber(v string) *ListInstalledSoftwaresResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBody) SetPageSize(v string) *ListInstalledSoftwaresResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBody) SetRequestId(v string) *ListInstalledSoftwaresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBody) SetTotalCount(v string) *ListInstalledSoftwaresResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstalledSoftwaresResponseBodyAdditionalPackages struct {
	AdditionalPackageInfos []*ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos `json:"AdditionalPackageInfos,omitempty" xml:"AdditionalPackageInfos,omitempty" type:"Repeated"`
}

func (s ListInstalledSoftwaresResponseBodyAdditionalPackages) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwaresResponseBodyAdditionalPackages) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackages) SetAdditionalPackageInfos(v []*ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) *ListInstalledSoftwaresResponseBodyAdditionalPackages {
	s.AdditionalPackageInfos = v
	return s
}

type ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos struct {
	// The category into which the software falls.
	//
	// example:
	//
	// NWP
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the software was installed.
	//
	// example:
	//
	// 2024-03-05 18:24:08
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The software description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the software icon.
	//
	// example:
	//
	// https://gw.alicdn.com/imgextra/i2/O1CN01FIkxZ81LmE0fvrAyR_!!6000000001341-55-tps-6349-1603.svg
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The software name.
	//
	// example:
	//
	// gromacs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The installation status of the software.
	//
	// Valid values:
	//
	// 	- Installed
	//
	// 	- Uninstalled
	//
	// 	- Installing
	//
	// 	- Exception
	//
	// example:
	//
	// Installed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The software version.
	//
	// example:
	//
	// 2024.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetCategory(v string) *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Category = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetCreateTime(v string) *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.CreateTime = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetDescription(v string) *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Description = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetIcon(v string) *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Icon = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetName(v string) *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Name = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetStatus(v string) *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Status = &v
	return s
}

func (s *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetVersion(v string) *ListInstalledSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Version = &v
	return s
}

type ListInstalledSoftwaresResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstalledSoftwaresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstalledSoftwaresResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwaresResponse) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwaresResponse) SetHeaders(v map[string]*string) *ListInstalledSoftwaresResponse {
	s.Headers = v
	return s
}

func (s *ListInstalledSoftwaresResponse) SetStatusCode(v int32) *ListInstalledSoftwaresResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstalledSoftwaresResponse) SetBody(v *ListInstalledSoftwaresResponseBody) *ListInstalledSoftwaresResponse {
	s.Body = v
	return s
}

type ListJobsRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-csbua72***
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The job filter information.
	JobFilter *ListJobsRequestJobFilter `json:"JobFilter,omitempty" xml:"JobFilter,omitempty" type:"Struct"`
	// The page number of the page to return.
	//
	// 	- Pages start from page 1.
	//
	// 	- Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// 	- Maximum value: 50.
	//
	// 	- Default value: 10
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *ListJobsRequest) SetJobFilter(v *ListJobsRequestJobFilter) *ListJobsRequest {
	s.JobFilter = v
	return s
}

func (s *ListJobsRequest) SetPageNumber(v string) *ListJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v string) *ListJobsRequest {
	s.PageSize = &v
	return s
}

type ListJobsRequestJobFilter struct {
	// The time when the job was last updated. The value is a UNIX timestamp representing the number of seconds that have elapsed since 1970-01-01T00:00:00Z.
	//
	// example:
	//
	// 1724123085
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// The time when the job started. The value is a UNIX timestamp representing the number of seconds that have elapsed since 1970-01-01T00:00:00Z.
	//
	// example:
	//
	// 1724122486
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// Job diagnosis and analysis list.
	Diagnosis []*ListJobsRequestJobFilterDiagnosis `json:"Diagnosis,omitempty" xml:"Diagnosis,omitempty" type:"Repeated"`
	// The job name. Fuzzy match is supported.
	//
	// example:
	//
	// testjob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The job status. Valid values:
	//
	// 	- all: returns all jobs.
	//
	// 	- finished: returns completed jobs.
	//
	// 	- notfinish: returns uncompleted jobs.
	//
	// Default value: all.
	//
	// example:
	//
	// all
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The compute nodes that run the jobs.
	Nodes []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The queues to which the jobs belong.
	Queues []*string `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
	// The result sorting configurations.
	SortBy *ListJobsRequestJobFilterSortBy `json:"SortBy,omitempty" xml:"SortBy,omitempty" type:"Struct"`
	// The users that run the jobs.
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListJobsRequestJobFilter) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequestJobFilter) GoString() string {
	return s.String()
}

func (s *ListJobsRequestJobFilter) SetCreateTimeEnd(v string) *ListJobsRequestJobFilter {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListJobsRequestJobFilter) SetCreateTimeStart(v string) *ListJobsRequestJobFilter {
	s.CreateTimeStart = &v
	return s
}

func (s *ListJobsRequestJobFilter) SetDiagnosis(v []*ListJobsRequestJobFilterDiagnosis) *ListJobsRequestJobFilter {
	s.Diagnosis = v
	return s
}

func (s *ListJobsRequestJobFilter) SetJobName(v string) *ListJobsRequestJobFilter {
	s.JobName = &v
	return s
}

func (s *ListJobsRequestJobFilter) SetJobStatus(v string) *ListJobsRequestJobFilter {
	s.JobStatus = &v
	return s
}

func (s *ListJobsRequestJobFilter) SetNodes(v []*string) *ListJobsRequestJobFilter {
	s.Nodes = v
	return s
}

func (s *ListJobsRequestJobFilter) SetQueues(v []*string) *ListJobsRequestJobFilter {
	s.Queues = v
	return s
}

func (s *ListJobsRequestJobFilter) SetSortBy(v *ListJobsRequestJobFilterSortBy) *ListJobsRequestJobFilter {
	s.SortBy = v
	return s
}

func (s *ListJobsRequestJobFilter) SetUsers(v []*string) *ListJobsRequestJobFilter {
	s.Users = v
	return s
}

type ListJobsRequestJobFilterDiagnosis struct {
	// Job diagnosis threshold comparator.
	//
	// example:
	//
	// greater
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Job diagnosis and analysis metrics
	//
	// example:
	//
	// run_duration
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// Job diagnosis threshold.
	//
	// example:
	//
	// 24
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListJobsRequestJobFilterDiagnosis) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequestJobFilterDiagnosis) GoString() string {
	return s.String()
}

func (s *ListJobsRequestJobFilterDiagnosis) SetOperator(v string) *ListJobsRequestJobFilterDiagnosis {
	s.Operator = &v
	return s
}

func (s *ListJobsRequestJobFilterDiagnosis) SetOption(v string) *ListJobsRequestJobFilterDiagnosis {
	s.Option = &v
	return s
}

func (s *ListJobsRequestJobFilterDiagnosis) SetThreshold(v string) *ListJobsRequestJobFilterDiagnosis {
	s.Threshold = &v
	return s
}

type ListJobsRequestJobFilterSortBy struct {
	// The order in which jobs are sorted based on their execution time. Valid values:
	//
	// 	- asc: in ascending order.
	//
	// 	- desc: in descending order.
	//
	// Default value: desc.
	//
	// example:
	//
	// asc
	ExecuteOrder *string `json:"ExecuteOrder,omitempty" xml:"ExecuteOrder,omitempty"`
	// The order in which jobs are sorted based on their queuing time. Valid values:
	//
	// 	- asc: in ascending order.
	//
	// 	- desc: in descending order.
	//
	// Default value: desc.
	//
	// example:
	//
	// desc
	PendOrder *string `json:"PendOrder,omitempty" xml:"PendOrder,omitempty"`
	// The order in which jobs are sorted based on their submitting time. Valid values:
	//
	// 	- asc: in ascending order.
	//
	// 	- desc: in descending order.
	//
	// Default value: desc.
	//
	// example:
	//
	// asc
	SubmitOrder *string `json:"SubmitOrder,omitempty" xml:"SubmitOrder,omitempty"`
}

func (s ListJobsRequestJobFilterSortBy) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequestJobFilterSortBy) GoString() string {
	return s.String()
}

func (s *ListJobsRequestJobFilterSortBy) SetExecuteOrder(v string) *ListJobsRequestJobFilterSortBy {
	s.ExecuteOrder = &v
	return s
}

func (s *ListJobsRequestJobFilterSortBy) SetPendOrder(v string) *ListJobsRequestJobFilterSortBy {
	s.PendOrder = &v
	return s
}

func (s *ListJobsRequestJobFilterSortBy) SetSubmitOrder(v string) *ListJobsRequestJobFilterSortBy {
	s.SubmitOrder = &v
	return s
}

type ListJobsShrinkRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-csbua72***
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The job filter information.
	JobFilterShrink *string `json:"JobFilter,omitempty" xml:"JobFilter,omitempty"`
	// The page number of the page to return.
	//
	// 	- Pages start from page 1.
	//
	// 	- Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// 	- Maximum value: 50.
	//
	// 	- Default value: 10
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListJobsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListJobsShrinkRequest) SetClusterId(v string) *ListJobsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetJobFilterShrink(v string) *ListJobsShrinkRequest {
	s.JobFilterShrink = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPageNumber(v string) *ListJobsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPageSize(v string) *ListJobsShrinkRequest {
	s.PageSize = &v
	return s
}

type ListJobsResponseBody struct {
	// The jobs.
	Jobs []*ListJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The page number. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EABFBD93-58BE-53F3-BBFE-8654BB2E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) SetJobs(v []*ListJobsResponseBodyJobs) *ListJobsResponseBody {
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

func (s *ListJobsResponseBody) SetSuccess(v bool) *ListJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalCount(v int32) *ListJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListJobsResponseBodyJobs struct {
	// The job name.
	//
	// example:
	//
	// testjob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The job configurations.
	JobSpec *ListJobsResponseBodyJobsJobSpec `json:"JobSpec,omitempty" xml:"JobSpec,omitempty" type:"Struct"`
}

func (s ListJobsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobs) SetJobName(v string) *ListJobsResponseBodyJobs {
	s.JobName = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetJobSpec(v *ListJobsResponseBodyJobsJobSpec) *ListJobsResponseBodyJobs {
	s.JobSpec = v
	return s
}

type ListJobsResponseBodyJobsJobSpec struct {
	// The array job ID.
	//
	// example:
	//
	// 4
	ArrayJobId *string `json:"ArrayJobId,omitempty" xml:"ArrayJobId,omitempty"`
	// The ID of the job in the array.
	//
	// example:
	//
	// 1
	ArrayJobSubId *string `json:"ArrayJobSubId,omitempty" xml:"ArrayJobSubId,omitempty"`
	// The queue format of the job.
	//
	// 	- If the job is not in a queue, the output is empty.
	//
	// 	- The format is X-Y:Z. X indicates the first index, Y indicates the final index, and Z indicates the step size. For example, 2-7:2 indicates three sub-jobs numbered 2, 4, and 6.
	//
	// example:
	//
	// 1-5:2
	ArrayRequest *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	// The job description.
	//
	// example:
	//
	// jobDescription
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 12
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The queue name.
	//
	// example:
	//
	// comp
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	// The time when the job was last updated.
	//
	// example:
	//
	// 1724123085
	LastModifyTime *string `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The compute nodes that were used to run the job.
	//
	// example:
	//
	// compute[002,005,003]
	NodeList *string `json:"NodeList,omitempty" xml:"NodeList,omitempty"`
	// The job priority. Valid values: 0 to 9. A larger value indicates a higher priority.
	//
	// example:
	//
	// 0
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The information about the resources required to run the job.
	Resources *ListJobsResponseBodyJobsJobSpecResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// Actual resource usage of the job program
	ResourcesActualOccupied *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied `json:"ResourcesActualOccupied,omitempty" xml:"ResourcesActualOccupied,omitempty" type:"Struct"`
	// The user that ran the job.
	//
	// example:
	//
	// testuser1
	RunasUser *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	// Job start time.
	//
	// example:
	//
	// 1724122486
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The job state. Valid values: (PBS cluster and Slurm cluster)
	//
	// 	- FINISHED/Completed
	//
	// 	- RUNNING/Running
	//
	// 	- QUEUED/Pending
	//
	// 	- FAILED/Failed
	//
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The error output path.
	//
	// example:
	//
	// ./Temp
	StderrPath *string `json:"StderrPath,omitempty" xml:"StderrPath,omitempty"`
	// The standard output path.
	//
	// example:
	//
	// ./Temp
	StdoutPath *string `json:"StdoutPath,omitempty" xml:"StdoutPath,omitempty"`
	// The time when the job was submitted.
	//
	// example:
	//
	// 1724122486
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// The variables of the job.
	//
	// example:
	//
	// {"PBS_O_SHELL":"/bin/bash", 	"PBS_O_HOST":"manager", 	"PBS_O_SYSTEM":"Linux", 	"PBS_O_LANG":"en_US.UTF-8", 	"PBS_O_QUEUE":"workq"}
	Variables *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s ListJobsResponseBodyJobsJobSpec) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobsJobSpec) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsJobSpec) SetArrayJobId(v string) *ListJobsResponseBodyJobsJobSpec {
	s.ArrayJobId = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetArrayJobSubId(v string) *ListJobsResponseBodyJobsJobSpec {
	s.ArrayJobSubId = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetArrayRequest(v string) *ListJobsResponseBodyJobsJobSpec {
	s.ArrayRequest = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetComment(v string) *ListJobsResponseBodyJobsJobSpec {
	s.Comment = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetId(v string) *ListJobsResponseBodyJobsJobSpec {
	s.Id = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetJobQueue(v string) *ListJobsResponseBodyJobsJobSpec {
	s.JobQueue = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetLastModifyTime(v string) *ListJobsResponseBodyJobsJobSpec {
	s.LastModifyTime = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetNodeList(v string) *ListJobsResponseBodyJobsJobSpec {
	s.NodeList = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetPriority(v string) *ListJobsResponseBodyJobsJobSpec {
	s.Priority = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetResources(v *ListJobsResponseBodyJobsJobSpecResources) *ListJobsResponseBodyJobsJobSpec {
	s.Resources = v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetResourcesActualOccupied(v *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) *ListJobsResponseBodyJobsJobSpec {
	s.ResourcesActualOccupied = v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetRunasUser(v string) *ListJobsResponseBodyJobsJobSpec {
	s.RunasUser = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetStartTime(v string) *ListJobsResponseBodyJobsJobSpec {
	s.StartTime = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetState(v string) *ListJobsResponseBodyJobsJobSpec {
	s.State = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetStderrPath(v string) *ListJobsResponseBodyJobsJobSpec {
	s.StderrPath = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetStdoutPath(v string) *ListJobsResponseBodyJobsJobSpec {
	s.StdoutPath = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetSubmitTime(v string) *ListJobsResponseBodyJobsJobSpec {
	s.SubmitTime = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetVariables(v string) *ListJobsResponseBodyJobsJobSpec {
	s.Variables = &v
	return s
}

type ListJobsResponseBodyJobsJobSpecResources struct {
	// The number of vCPUs that were used to run the job.
	//
	// example:
	//
	// 6
	Cores *string `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The number of GPUs that were used to run the job.
	//
	// example:
	//
	// 0
	Gpus *string `json:"Gpus,omitempty" xml:"Gpus,omitempty"`
	// The size of memory that was used to run the job.
	//
	// example:
	//
	// 1536MB
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The number of compute nodes that were used to run the job.
	//
	// example:
	//
	// 3
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s ListJobsResponseBodyJobsJobSpecResources) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobsJobSpecResources) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsJobSpecResources) SetCores(v string) *ListJobsResponseBodyJobsJobSpecResources {
	s.Cores = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpecResources) SetGpus(v string) *ListJobsResponseBodyJobsJobSpecResources {
	s.Gpus = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpecResources) SetMemory(v string) *ListJobsResponseBodyJobsJobSpecResources {
	s.Memory = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpecResources) SetNodes(v string) *ListJobsResponseBodyJobsJobSpecResources {
	s.Nodes = &v
	return s
}

type ListJobsResponseBodyJobsJobSpecResourcesActualOccupied struct {
	// Number of CPU cores.
	//
	// example:
	//
	// 4
	Cores *string `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// Number of CPUs
	//
	// example:
	//
	// 0
	Gpus *string `json:"Gpus,omitempty" xml:"Gpus,omitempty"`
	// Number of memory.
	//
	// example:
	//
	// 982MB
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// Number of compute nodes.
	//
	// example:
	//
	// 2
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) SetCores(v string) *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied {
	s.Cores = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) SetGpus(v string) *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied {
	s.Gpus = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) SetMemory(v string) *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied {
	s.Memory = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) SetNodes(v string) *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied {
	s.Nodes = &v
	return s
}

type ListJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListJobsResponse) SetStatusCode(v int32) *ListJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobsResponse) SetBody(v *ListJobsResponseBody) *ListJobsResponse {
	s.Body = v
	return s
}

type ListNodesRequest struct {
	// The cluster ID. You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The hostnames of the compute nodes that you want to query.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IP addresses of the compute nodes that you want to query.
	PrivateIpAddress []*string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" type:"Repeated"`
	// The queues to which the nodes belong.
	QueueNames []*string `json:"QueueNames,omitempty" xml:"QueueNames,omitempty" type:"Repeated"`
	// Specifies whether the results are sorted in ascending or descending order. Valid values:
	//
	// 	- Forward: ascending
	//
	// 	- Backward: descending
	//
	// Default value: Forward.
	//
	// example:
	//
	// Forward
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The sorting method of the node list. Valid values:
	//
	// 	- AddedTime: sorts the nodes by the time that they were added.
	//
	// 	- HostName: sorts the nodes by their hostnames.
	//
	// Default value: addedtime.
	//
	// example:
	//
	// AddedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The states of the compute nodes to be queried.
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
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

func (s *ListNodesRequest) SetHostnames(v []*string) *ListNodesRequest {
	s.Hostnames = v
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

func (s *ListNodesRequest) SetPrivateIpAddress(v []*string) *ListNodesRequest {
	s.PrivateIpAddress = v
	return s
}

func (s *ListNodesRequest) SetQueueNames(v []*string) *ListNodesRequest {
	s.QueueNames = v
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

func (s *ListNodesRequest) SetStatus(v []*string) *ListNodesRequest {
	s.Status = v
	return s
}

type ListNodesShrinkRequest struct {
	// The cluster ID. You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The hostnames of the compute nodes that you want to query.
	HostnamesShrink *string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IP addresses of the compute nodes that you want to query.
	PrivateIpAddressShrink *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The queues to which the nodes belong.
	QueueNamesShrink *string `json:"QueueNames,omitempty" xml:"QueueNames,omitempty"`
	// Specifies whether the results are sorted in ascending or descending order. Valid values:
	//
	// 	- Forward: ascending
	//
	// 	- Backward: descending
	//
	// Default value: Forward.
	//
	// example:
	//
	// Forward
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The sorting method of the node list. Valid values:
	//
	// 	- AddedTime: sorts the nodes by the time that they were added.
	//
	// 	- HostName: sorts the nodes by their hostnames.
	//
	// Default value: addedtime.
	//
	// example:
	//
	// AddedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The states of the compute nodes to be queried.
	StatusShrink *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListNodesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListNodesShrinkRequest) SetClusterId(v string) *ListNodesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodesShrinkRequest) SetHostnamesShrink(v string) *ListNodesShrinkRequest {
	s.HostnamesShrink = &v
	return s
}

func (s *ListNodesShrinkRequest) SetPageNumber(v int32) *ListNodesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesShrinkRequest) SetPageSize(v int32) *ListNodesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesShrinkRequest) SetPrivateIpAddressShrink(v string) *ListNodesShrinkRequest {
	s.PrivateIpAddressShrink = &v
	return s
}

func (s *ListNodesShrinkRequest) SetQueueNamesShrink(v string) *ListNodesShrinkRequest {
	s.QueueNamesShrink = &v
	return s
}

func (s *ListNodesShrinkRequest) SetSequence(v string) *ListNodesShrinkRequest {
	s.Sequence = &v
	return s
}

func (s *ListNodesShrinkRequest) SetSortBy(v string) *ListNodesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListNodesShrinkRequest) SetStatusShrink(v string) *ListNodesShrinkRequest {
	s.StatusShrink = &v
	return s
}

type ListNodesResponseBody struct {
	// The information about the nodes.
	Nodes []*ListNodesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 89A1AC0F-4A6C-4F3D-98F9-BEF9A823****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) SetNodes(v []*ListNodesResponseBodyNodes) *ListNodesResponseBody {
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
	// The time when the node was created.
	//
	// example:
	//
	// 2020-06-09T06:22:02.000Z
	AddTime *string `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	// The deployment set ID.
	//
	// example:
	//
	// ds-8vbe4av4gededlqg****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The time when the node expires.
	//
	// example:
	//
	// 2020-06-09T06:22:02.000Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The hostname of the node.
	//
	// example:
	//
	// edas.cn-shanghai.aliyuncs.com
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Indicates whether hyper-threading is enabled.
	//
	// example:
	//
	// true
	HtEnabled *bool `json:"HtEnabled,omitempty" xml:"HtEnabled,omitempty"`
	// The instance ID of the node.
	//
	// example:
	//
	// i-bp15707mys2rsy0j****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The image ID of the node.
	//
	// example:
	//
	// centos_7_06_64_20G_alibase_20190711.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance type of the node.
	//
	// example:
	//
	// ecs.c5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The VPC IP address of the node.
	//
	// example:
	//
	// ``172.16.**.**``
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// Indicates whether deletion protection is enabled for the node. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	KeepAlive *bool `json:"KeepAlive,omitempty" xml:"KeepAlive,omitempty"`
	// The public IP address of the node.
	//
	// example:
	//
	// ``172.16.**.**``
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	// The name of the queue to which the node belongs.
	//
	// example:
	//
	// autoque3
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The bidding policy of the node. Valid values:
	//
	// 	- NoSpot: The instances of the compute node are pay-as-you-go instances.
	//
	// 	- SpotWithPriceLimit: The instances are created as preemptible instances with a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The node is a preemptible instance for which the market price at the time of purchase is automatically used as the bidding price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The node state in the scheduler.
	//
	// example:
	//
	// active
	StateInSched *string `json:"StateInSched,omitempty" xml:"StateInSched,omitempty"`
	// The node state. Valid values:
	//
	// 	- uninit: The node is being installed.
	//
	// 	- initing: The node is being initialized.
	//
	// 	- running: The node is running.
	//
	// 	- releasing: The node is being released.
	//
	// 	- stopped: The node is stopped.
	//
	// 	- exception: The node has run into an exception.
	//
	// 	- untracking: The node is not added to the cluster.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The hardware configurations of the node.
	TotalResources *ListNodesResponseBodyNodesTotalResources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty" type:"Struct"`
	// The vSwitch ID of the node.
	//
	// example:
	//
	// vsw-bp1e47optm9g58zcu****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1gnu8br4ay7beb2w****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the node.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListNodesResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodes) SetAddTime(v string) *ListNodesResponseBodyNodes {
	s.AddTime = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetDeploymentSetId(v string) *ListNodesResponseBodyNodes {
	s.DeploymentSetId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetExpiredTime(v string) *ListNodesResponseBodyNodes {
	s.ExpiredTime = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetHostname(v string) *ListNodesResponseBodyNodes {
	s.Hostname = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetHtEnabled(v bool) *ListNodesResponseBodyNodes {
	s.HtEnabled = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetId(v string) *ListNodesResponseBodyNodes {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetImageId(v string) *ListNodesResponseBodyNodes {
	s.ImageId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetInstanceType(v string) *ListNodesResponseBodyNodes {
	s.InstanceType = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetIpAddress(v string) *ListNodesResponseBodyNodes {
	s.IpAddress = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetKeepAlive(v bool) *ListNodesResponseBodyNodes {
	s.KeepAlive = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetPublicIpAddress(v string) *ListNodesResponseBodyNodes {
	s.PublicIpAddress = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetQueueName(v string) *ListNodesResponseBodyNodes {
	s.QueueName = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetSpotStrategy(v string) *ListNodesResponseBodyNodes {
	s.SpotStrategy = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetStateInSched(v string) *ListNodesResponseBodyNodes {
	s.StateInSched = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetStatus(v string) *ListNodesResponseBodyNodes {
	s.Status = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetTotalResources(v *ListNodesResponseBodyNodesTotalResources) *ListNodesResponseBodyNodes {
	s.TotalResources = v
	return s
}

func (s *ListNodesResponseBodyNodes) SetVSwitchId(v string) *ListNodesResponseBodyNodes {
	s.VSwitchId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetVpcId(v string) *ListNodesResponseBodyNodes {
	s.VpcId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetZoneId(v string) *ListNodesResponseBodyNodes {
	s.ZoneId = &v
	return s
}

type ListNodesResponseBodyNodesTotalResources struct {
	// The number of vCPUs.
	//
	// example:
	//
	// 1
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 0
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The amount of memory. Unit: GiB.
	//
	// example:
	//
	// 1024
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesResponseBodyNodesTotalResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodesTotalResources) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodesTotalResources) SetCpu(v int32) *ListNodesResponseBodyNodesTotalResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesTotalResources) SetGpu(v int32) *ListNodesResponseBodyNodesTotalResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesTotalResources) SetMemory(v int32) *ListNodesResponseBodyNodesTotalResources {
	s.Memory = &v
	return s
}

type ListNodesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListNodesResponse) SetStatusCode(v int32) *ListNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesResponse) SetBody(v *ListNodesResponseBody) *ListNodesResponse {
	s.Body = v
	return s
}

type ListQueuesRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The names of the queues that you want to query. You can specify up to eight names.
	QueueNames []*string `json:"QueueNames,omitempty" xml:"QueueNames,omitempty" type:"Repeated"`
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

func (s *ListQueuesRequest) SetQueueNames(v []*string) *ListQueuesRequest {
	s.QueueNames = v
	return s
}

type ListQueuesShrinkRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The names of the queues that you want to query. You can specify up to eight names.
	QueueNamesShrink *string `json:"QueueNames,omitempty" xml:"QueueNames,omitempty"`
}

func (s ListQueuesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListQueuesShrinkRequest) SetClusterId(v string) *ListQueuesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ListQueuesShrinkRequest) SetQueueNamesShrink(v string) *ListQueuesShrinkRequest {
	s.QueueNamesShrink = &v
	return s
}

type ListQueuesResponseBody struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about the queues.
	Queues []*ListQueuesResponseBodyQueues `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C6E5005C-00B0-4F27-98BB-95AB88016C22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQueuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBody) SetClusterId(v string) *ListQueuesResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ListQueuesResponseBody) SetPageNumber(v int32) *ListQueuesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListQueuesResponseBody) SetPageSize(v int32) *ListQueuesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListQueuesResponseBody) SetQueues(v []*ListQueuesResponseBodyQueues) *ListQueuesResponseBody {
	s.Queues = v
	return s
}

func (s *ListQueuesResponseBody) SetRequestId(v string) *ListQueuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueuesResponseBody) SetTotalCount(v int32) *ListQueuesResponseBody {
	s.TotalCount = &v
	return s
}

type ListQueuesResponseBodyQueues struct {
	// The hardware configurations of the compute nodes that are added in auto scale-outs. Up to five nodes are displayed.
	ComputeNodes []*NodeTemplate `json:"ComputeNodes,omitempty" xml:"ComputeNodes,omitempty" type:"Repeated"`
	// The time when the queue was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2023-11-10T02:04:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether auto scale-in is enabled for the queue. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableScaleIn *bool `json:"EnableScaleIn,omitempty" xml:"EnableScaleIn,omitempty"`
	// Indicates whether auto scale-out is enabled for the queue. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableScaleOut *bool `json:"EnableScaleOut,omitempty" xml:"EnableScaleOut,omitempty"`
	// The maximum number of compute nodes that the queue can contain.
	//
	// example:
	//
	// 100
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The minimum number of compute nodes that are added to the queue in each auto scale-out.
	//
	// example:
	//
	// 1
	MaxCountPerCycle *int32 `json:"MaxCountPerCycle,omitempty" xml:"MaxCountPerCycle,omitempty"`
	// The minimum number of compute nodes that the queue must contain.
	//
	// example:
	//
	// 0
	MinCount *int32 `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// The statistics about the compute nodes in the queue.
	Nodes *ListQueuesResponseBodyQueuesNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
	// The queue name.
	//
	// example:
	//
	// comp
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The total number of vCPUs that are used by all compute nodes in the queue.
	//
	// example:
	//
	// 24
	TotalCores *int32 `json:"TotalCores,omitempty" xml:"TotalCores,omitempty"`
	// The time when the queue was updated. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2024-04-25T02:02:32
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The vSwitches that can be used for added nodes during auto scale-outs. Up to three vSwitches are displayed.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s ListQueuesResponseBodyQueues) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyQueues) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyQueues) SetComputeNodes(v []*NodeTemplate) *ListQueuesResponseBodyQueues {
	s.ComputeNodes = v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetCreateTime(v string) *ListQueuesResponseBodyQueues {
	s.CreateTime = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetEnableScaleIn(v bool) *ListQueuesResponseBodyQueues {
	s.EnableScaleIn = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetEnableScaleOut(v bool) *ListQueuesResponseBodyQueues {
	s.EnableScaleOut = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetMaxCount(v int32) *ListQueuesResponseBodyQueues {
	s.MaxCount = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetMaxCountPerCycle(v int32) *ListQueuesResponseBodyQueues {
	s.MaxCountPerCycle = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetMinCount(v int32) *ListQueuesResponseBodyQueues {
	s.MinCount = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetNodes(v *ListQueuesResponseBodyQueuesNodes) *ListQueuesResponseBodyQueues {
	s.Nodes = v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetQueueName(v string) *ListQueuesResponseBodyQueues {
	s.QueueName = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetTotalCores(v int32) *ListQueuesResponseBodyQueues {
	s.TotalCores = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetUpdateTime(v string) *ListQueuesResponseBodyQueues {
	s.UpdateTime = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetVSwitchIds(v []*string) *ListQueuesResponseBodyQueues {
	s.VSwitchIds = v
	return s
}

type ListQueuesResponseBodyQueuesNodes struct {
	// The number of compute nodes that are not ready.
	//
	// example:
	//
	// 2
	CreatingCounts *int32 `json:"CreatingCounts,omitempty" xml:"CreatingCounts,omitempty"`
	// The number of malfunctioning compute nodes.
	//
	// example:
	//
	// 0
	ExceptionCounts *int32 `json:"ExceptionCounts,omitempty" xml:"ExceptionCounts,omitempty"`
	// The number of running compute nodes.
	//
	// example:
	//
	// 1
	RunningCounts *int32 `json:"RunningCounts,omitempty" xml:"RunningCounts,omitempty"`
}

func (s ListQueuesResponseBodyQueuesNodes) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyQueuesNodes) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyQueuesNodes) SetCreatingCounts(v int32) *ListQueuesResponseBodyQueuesNodes {
	s.CreatingCounts = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesNodes) SetExceptionCounts(v int32) *ListQueuesResponseBodyQueuesNodes {
	s.ExceptionCounts = &v
	return s
}

func (s *ListQueuesResponseBodyQueuesNodes) SetRunningCounts(v int32) *ListQueuesResponseBodyQueuesNodes {
	s.RunningCounts = &v
	return s
}

type ListQueuesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListQueuesResponse) SetStatusCode(v int32) *ListQueuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueuesResponse) SetBody(v *ListQueuesResponseBody) *ListQueuesResponse {
	s.Body = v
	return s
}

type ListSharedStoragesRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the attached file system.
	//
	// example:
	//
	// 0bd504b0**
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The type of the attached file system. Valid values:
	//
	// 	- nas
	//
	// 	- cpfs
	//
	// example:
	//
	// nas
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s ListSharedStoragesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSharedStoragesRequest) GoString() string {
	return s.String()
}

func (s *ListSharedStoragesRequest) SetClusterId(v string) *ListSharedStoragesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListSharedStoragesRequest) SetFileSystemId(v string) *ListSharedStoragesRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListSharedStoragesRequest) SetFileSystemType(v string) *ListSharedStoragesRequest {
	s.FileSystemType = &v
	return s
}

type ListSharedStoragesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F6757FA4-8FED-4602-B7F5-3550C084****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the attached shared storage.
	SharedStorages []*ListSharedStoragesResponseBodySharedStorages `json:"SharedStorages,omitempty" xml:"SharedStorages,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSharedStoragesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSharedStoragesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSharedStoragesResponseBody) SetRequestId(v string) *ListSharedStoragesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSharedStoragesResponseBody) SetSharedStorages(v []*ListSharedStoragesResponseBodySharedStorages) *ListSharedStoragesResponseBody {
	s.SharedStorages = v
	return s
}

func (s *ListSharedStoragesResponseBody) SetSuccess(v string) *ListSharedStoragesResponseBody {
	s.Success = &v
	return s
}

type ListSharedStoragesResponseBodySharedStorages struct {
	// The ID of the attached file system.
	//
	// example:
	//
	// 08c7f4b***
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The protocol used by the attached file system. Valid values:
	//
	// 	- nfs3
	//
	// 	- nfs4
	//
	// 	- cpfs
	//
	// example:
	//
	// nfs4
	FileSystemProtocol *string `json:"FileSystemProtocol,omitempty" xml:"FileSystemProtocol,omitempty"`
	// The type of the attached file system. Valid values:
	//
	// 	- nas
	//
	// 	- cpfs
	//
	// example:
	//
	// nas
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The mount information.
	MountInfo []*ListSharedStoragesResponseBodySharedStoragesMountInfo `json:"MountInfo,omitempty" xml:"MountInfo,omitempty" type:"Repeated"`
}

func (s ListSharedStoragesResponseBodySharedStorages) String() string {
	return tea.Prettify(s)
}

func (s ListSharedStoragesResponseBodySharedStorages) GoString() string {
	return s.String()
}

func (s *ListSharedStoragesResponseBodySharedStorages) SetFileSystemId(v string) *ListSharedStoragesResponseBodySharedStorages {
	s.FileSystemId = &v
	return s
}

func (s *ListSharedStoragesResponseBodySharedStorages) SetFileSystemProtocol(v string) *ListSharedStoragesResponseBodySharedStorages {
	s.FileSystemProtocol = &v
	return s
}

func (s *ListSharedStoragesResponseBodySharedStorages) SetFileSystemType(v string) *ListSharedStoragesResponseBodySharedStorages {
	s.FileSystemType = &v
	return s
}

func (s *ListSharedStoragesResponseBodySharedStorages) SetMountInfo(v []*ListSharedStoragesResponseBodySharedStoragesMountInfo) *ListSharedStoragesResponseBodySharedStorages {
	s.MountInfo = v
	return s
}

type ListSharedStoragesResponseBodySharedStoragesMountInfo struct {
	// The local mount directory of the attached file system.
	//
	// example:
	//
	// /test
	MountDirectory *string `json:"MountDirectory,omitempty" xml:"MountDirectory,omitempty"`
	// The mount options for the attached file system. Valid values:
	//
	// 	- \\-t nfs -o vers=3,nolock,proto=tcp,noresvport
	//
	// 	- \\-t nfs -o vers=4.0,noresvport
	//
	// example:
	//
	// -t nfs -o vers=4.0,noresvport
	MountOptions *string `json:"MountOptions,omitempty" xml:"MountOptions,omitempty"`
	// The mount target of the attached file system.
	//
	// example:
	//
	// 0bd504b***-ngq26.cn-hangzhou.nas.aliyuncs.com
	MountTarget *string `json:"MountTarget,omitempty" xml:"MountTarget,omitempty"`
	// The protocol used by the mount target of the attached file system. Valid values:
	//
	// 	- nfs3
	//
	// 	- nfs4
	//
	// 	- cpfs
	//
	// example:
	//
	// nfs3
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The storage directory of the attached file system.
	//
	// example:
	//
	// /testehpc
	StorageDirectory *string `json:"StorageDirectory,omitempty" xml:"StorageDirectory,omitempty"`
}

func (s ListSharedStoragesResponseBodySharedStoragesMountInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSharedStoragesResponseBodySharedStoragesMountInfo) GoString() string {
	return s.String()
}

func (s *ListSharedStoragesResponseBodySharedStoragesMountInfo) SetMountDirectory(v string) *ListSharedStoragesResponseBodySharedStoragesMountInfo {
	s.MountDirectory = &v
	return s
}

func (s *ListSharedStoragesResponseBodySharedStoragesMountInfo) SetMountOptions(v string) *ListSharedStoragesResponseBodySharedStoragesMountInfo {
	s.MountOptions = &v
	return s
}

func (s *ListSharedStoragesResponseBodySharedStoragesMountInfo) SetMountTarget(v string) *ListSharedStoragesResponseBodySharedStoragesMountInfo {
	s.MountTarget = &v
	return s
}

func (s *ListSharedStoragesResponseBodySharedStoragesMountInfo) SetProtocolType(v string) *ListSharedStoragesResponseBodySharedStoragesMountInfo {
	s.ProtocolType = &v
	return s
}

func (s *ListSharedStoragesResponseBodySharedStoragesMountInfo) SetStorageDirectory(v string) *ListSharedStoragesResponseBodySharedStoragesMountInfo {
	s.StorageDirectory = &v
	return s
}

type ListSharedStoragesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSharedStoragesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSharedStoragesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSharedStoragesResponse) GoString() string {
	return s.String()
}

func (s *ListSharedStoragesResponse) SetHeaders(v map[string]*string) *ListSharedStoragesResponse {
	s.Headers = v
	return s
}

func (s *ListSharedStoragesResponse) SetStatusCode(v int32) *ListSharedStoragesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSharedStoragesResponse) SetBody(v *ListSharedStoragesResponseBody) *ListSharedStoragesResponse {
	s.Body = v
	return s
}

type ListSoftwaresRequest struct {
	// The application category.
	//
	// example:
	//
	// NWP
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The software name.
	//
	// example:
	//
	// gromacs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operating system (OS) information.
	OsInfos []*ListSoftwaresRequestOsInfos `json:"OsInfos,omitempty" xml:"OsInfos,omitempty" type:"Repeated"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSoftwaresRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresRequest) GoString() string {
	return s.String()
}

func (s *ListSoftwaresRequest) SetCategory(v string) *ListSoftwaresRequest {
	s.Category = &v
	return s
}

func (s *ListSoftwaresRequest) SetClusterId(v string) *ListSoftwaresRequest {
	s.ClusterId = &v
	return s
}

func (s *ListSoftwaresRequest) SetName(v string) *ListSoftwaresRequest {
	s.Name = &v
	return s
}

func (s *ListSoftwaresRequest) SetOsInfos(v []*ListSoftwaresRequestOsInfos) *ListSoftwaresRequest {
	s.OsInfos = v
	return s
}

func (s *ListSoftwaresRequest) SetPageNumber(v string) *ListSoftwaresRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSoftwaresRequest) SetPageSize(v string) *ListSoftwaresRequest {
	s.PageSize = &v
	return s
}

type ListSoftwaresRequestOsInfos struct {
	// The OS architecture. Valid values:
	//
	// 	- x86_64
	//
	// 	- arm64
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The image tag.
	//
	// example:
	//
	// CentOS_7.9_64
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
}

func (s ListSoftwaresRequestOsInfos) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresRequestOsInfos) GoString() string {
	return s.String()
}

func (s *ListSoftwaresRequestOsInfos) SetArchitecture(v string) *ListSoftwaresRequestOsInfos {
	s.Architecture = &v
	return s
}

func (s *ListSoftwaresRequestOsInfos) SetOsTag(v string) *ListSoftwaresRequestOsInfos {
	s.OsTag = &v
	return s
}

type ListSoftwaresResponseBody struct {
	// The information about the software that can be installed in the cluster.
	AdditionalPackages *ListSoftwaresResponseBodyAdditionalPackages `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty" type:"Struct"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSoftwaresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBody) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBody) SetAdditionalPackages(v *ListSoftwaresResponseBodyAdditionalPackages) *ListSoftwaresResponseBody {
	s.AdditionalPackages = v
	return s
}

func (s *ListSoftwaresResponseBody) SetPageNumber(v string) *ListSoftwaresResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSoftwaresResponseBody) SetPageSize(v string) *ListSoftwaresResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSoftwaresResponseBody) SetRequestId(v string) *ListSoftwaresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSoftwaresResponseBody) SetTotalCount(v string) *ListSoftwaresResponseBody {
	s.TotalCount = &v
	return s
}

type ListSoftwaresResponseBodyAdditionalPackages struct {
	AdditionalPackageInfos []*ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos `json:"AdditionalPackageInfos,omitempty" xml:"AdditionalPackageInfos,omitempty" type:"Repeated"`
}

func (s ListSoftwaresResponseBodyAdditionalPackages) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodyAdditionalPackages) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodyAdditionalPackages) SetAdditionalPackageInfos(v []*ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) *ListSoftwaresResponseBodyAdditionalPackages {
	s.AdditionalPackageInfos = v
	return s
}

type ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos struct {
	// The application category.
	//
	// example:
	//
	// NWP
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The software description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the software icon.
	//
	// example:
	//
	// https://gw.alicdn.com/imgextra/i2/O1CN01FIkxZ81LmE0fvrAyR_!!6000000001341-55-tps-6349-1603.svg
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The software name.
	//
	// example:
	//
	// gromacs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The information about the software versions that can be installed in the cluster.
	Versions *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Struct"`
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetCategory(v string) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Category = &v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetDescription(v string) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Description = &v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetIcon(v string) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Icon = &v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetName(v string) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Name = &v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos) SetVersions(v *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersions) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfos {
	s.Versions = v
	return s
}

type ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersions struct {
	VersionInfos []*ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos `json:"VersionInfos,omitempty" xml:"VersionInfos,omitempty" type:"Repeated"`
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersions) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersions) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersions) SetVersionInfos(v []*ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersions {
	s.VersionInfos = v
	return s
}

type ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos struct {
	// Indicates whether the version is the latest.
	//
	// example:
	//
	// false
	Latest *string `json:"Latest,omitempty" xml:"Latest,omitempty"`
	// The information about the supported OSs.
	SupportOs *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOs `json:"SupportOs,omitempty" xml:"SupportOs,omitempty" type:"Struct"`
	// The software version.
	//
	// example:
	//
	// 2024.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos) SetLatest(v string) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos {
	s.Latest = &v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos) SetSupportOs(v *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOs) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos {
	s.SupportOs = v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos) SetVersion(v string) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfos {
	s.Version = &v
	return s
}

type ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOs struct {
	SupportOsInfos []*ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos `json:"SupportOsInfos,omitempty" xml:"SupportOsInfos,omitempty" type:"Repeated"`
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOs) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOs) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOs) SetSupportOsInfos(v []*ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOs {
	s.SupportOsInfos = v
	return s
}

type ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos struct {
	// The OS architecture. Valid values:
	//
	// 	- x86_64
	//
	// 	- arm64
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The image tag.
	//
	// example:
	//
	// Alibaba Cloud Linux  3.2104 LTS 64 bit ARM Edition
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos) SetArchitecture(v string) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos {
	s.Architecture = &v
	return s
}

func (s *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos) SetOsTag(v string) *ListSoftwaresResponseBodyAdditionalPackagesAdditionalPackageInfosVersionsVersionInfosSupportOsSupportOsInfos {
	s.OsTag = &v
	return s
}

type ListSoftwaresResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSoftwaresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListSoftwaresResponse) SetStatusCode(v int32) *ListSoftwaresResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSoftwaresResponse) SetBody(v *ListSoftwaresResponseBody) *ListSoftwaresResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the users.
	Users *ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
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
	// The time when the user was first added.
	//
	// example:
	//
	// 2014-08-22T17:46:47
	AddTime *string `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	// The name of the permission group. Valid values:
	//
	// users: ordinary permissions, which are suitable for regular users that need only to submit and debug jobs.
	//
	// wheel: sudo permissions, which are suitable for administrators who need to manage clusters. In addition to submitting and debugging jobs, you can also run sudo commands to install software and restart nodes.
	//
	// example:
	//
	// users
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The permission group ID.
	//
	// example:
	//
	// 100
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1001
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// testuser
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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

func (s *ListUsersResponseBodyUsersUserInfo) SetGroupId(v string) *ListUsersResponseBodyUsersUserInfo {
	s.GroupId = &v
	return s
}

func (s *ListUsersResponseBodyUsersUserInfo) SetUserId(v string) *ListUsersResponseBodyUsersUserInfo {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyUsersUserInfo) SetUserName(v string) *ListUsersResponseBodyUsersUserInfo {
	s.UserName = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type StopJobsRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The IDs of the jobs that you want to stop.
	JobIds []*string `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
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

func (s *StopJobsRequest) SetJobIds(v []*string) *StopJobsRequest {
	s.JobIds = v
	return s
}

type StopJobsShrinkRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The IDs of the jobs that you want to stop.
	JobIdsShrink *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
}

func (s StopJobsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s StopJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopJobsShrinkRequest) SetClusterId(v string) *StopJobsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *StopJobsShrinkRequest) SetJobIdsShrink(v string) *StopJobsShrinkRequest {
	s.JobIdsShrink = &v
	return s
}

type StopJobsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F8868A00-6757-5542-BDD6-E1040D94****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *StopJobsResponseBody) SetSuccess(v string) *StopJobsResponseBody {
	s.Success = &v
	return s
}

type StopJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *StopJobsResponse) SetStatusCode(v int32) *StopJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *StopJobsResponse) SetBody(v *StopJobsResponseBody) *StopJobsResponse {
	s.Body = v
	return s
}

type UnInstallAddonRequest struct {
	// The addon ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login-1.0-W2g****
	AddonId *string `json:"AddonId,omitempty" xml:"AddonId,omitempty"`
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UnInstallAddonRequest) String() string {
	return tea.Prettify(s)
}

func (s UnInstallAddonRequest) GoString() string {
	return s.String()
}

func (s *UnInstallAddonRequest) SetAddonId(v string) *UnInstallAddonRequest {
	s.AddonId = &v
	return s
}

func (s *UnInstallAddonRequest) SetClusterId(v string) *UnInstallAddonRequest {
	s.ClusterId = &v
	return s
}

type UnInstallAddonResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnInstallAddonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnInstallAddonResponseBody) GoString() string {
	return s.String()
}

func (s *UnInstallAddonResponseBody) SetRequestId(v string) *UnInstallAddonResponseBody {
	s.RequestId = &v
	return s
}

type UnInstallAddonResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnInstallAddonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnInstallAddonResponse) String() string {
	return tea.Prettify(s)
}

func (s UnInstallAddonResponse) GoString() string {
	return s.String()
}

func (s *UnInstallAddonResponse) SetHeaders(v map[string]*string) *UnInstallAddonResponse {
	s.Headers = v
	return s
}

func (s *UnInstallAddonResponse) SetStatusCode(v int32) *UnInstallAddonResponse {
	s.StatusCode = &v
	return s
}

func (s *UnInstallAddonResponse) SetBody(v *UnInstallAddonResponseBody) *UnInstallAddonResponse {
	s.Body = v
	return s
}

type UninstallSoftwaresRequest struct {
	// The information about the software systems that you want to uninstall.
	AdditionalPackages []*UninstallSoftwaresRequestAdditionalPackages `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UninstallSoftwaresRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallSoftwaresRequest) GoString() string {
	return s.String()
}

func (s *UninstallSoftwaresRequest) SetAdditionalPackages(v []*UninstallSoftwaresRequestAdditionalPackages) *UninstallSoftwaresRequest {
	s.AdditionalPackages = v
	return s
}

func (s *UninstallSoftwaresRequest) SetClusterId(v string) *UninstallSoftwaresRequest {
	s.ClusterId = &v
	return s
}

type UninstallSoftwaresRequestAdditionalPackages struct {
	// The software name.
	//
	// example:
	//
	// gromacs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The software version.
	//
	// example:
	//
	// 2024.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UninstallSoftwaresRequestAdditionalPackages) String() string {
	return tea.Prettify(s)
}

func (s UninstallSoftwaresRequestAdditionalPackages) GoString() string {
	return s.String()
}

func (s *UninstallSoftwaresRequestAdditionalPackages) SetName(v string) *UninstallSoftwaresRequestAdditionalPackages {
	s.Name = &v
	return s
}

func (s *UninstallSoftwaresRequestAdditionalPackages) SetVersion(v string) *UninstallSoftwaresRequestAdditionalPackages {
	s.Version = &v
	return s
}

type UninstallSoftwaresShrinkRequest struct {
	// The information about the software systems that you want to uninstall.
	AdditionalPackagesShrink *string `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty"`
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UninstallSoftwaresShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallSoftwaresShrinkRequest) GoString() string {
	return s.String()
}

func (s *UninstallSoftwaresShrinkRequest) SetAdditionalPackagesShrink(v string) *UninstallSoftwaresShrinkRequest {
	s.AdditionalPackagesShrink = &v
	return s
}

func (s *UninstallSoftwaresShrinkRequest) SetClusterId(v string) *UninstallSoftwaresShrinkRequest {
	s.ClusterId = &v
	return s
}

type UninstallSoftwaresResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UninstallSoftwaresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallSoftwaresResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallSoftwaresResponseBody) SetRequestId(v string) *UninstallSoftwaresResponseBody {
	s.RequestId = &v
	return s
}

type UninstallSoftwaresResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallSoftwaresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallSoftwaresResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallSoftwaresResponse) GoString() string {
	return s.String()
}

func (s *UninstallSoftwaresResponse) SetHeaders(v map[string]*string) *UninstallSoftwaresResponse {
	s.Headers = v
	return s
}

func (s *UninstallSoftwaresResponse) SetStatusCode(v int32) *UninstallSoftwaresResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallSoftwaresResponse) SetBody(v *UninstallSoftwaresResponseBody) *UninstallSoftwaresResponse {
	s.Body = v
	return s
}

type UpdateClusterRequest struct {
	// Specifies whether to enable auto scale-out for the cluster. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// 2.1.0
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// Specifies whether to enable auto scale-in for the cluster. Valid values:
	//
	// 	- true
	//
	// 	- false
	ClusterCustomConfiguration *UpdateClusterRequestClusterCustomConfiguration `json:"ClusterCustomConfiguration,omitempty" xml:"ClusterCustomConfiguration,omitempty" type:"Struct"`
	// The URL that is used to download the post-processing script.
	//
	// example:
	//
	// slurm22.05.8-serverless-cluster-20240805
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// The client version. By default, the latest version is used.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The post-processing script of the cluster.
	//
	// example:
	//
	// slurm22.05.8-serverless-cluster-20240805
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The idle duration of the compute nodes allowed by the cluster.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The request result. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	EnableScaleIn *bool `json:"EnableScaleIn,omitempty" xml:"EnableScaleIn,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// true
	EnableScaleOut *bool `json:"EnableScaleOut,omitempty" xml:"EnableScaleOut,omitempty"`
	// The scheduler specifications of the cluster.
	//
	// example:
	//
	// 2
	GrowInterval *int32 `json:"GrowInterval,omitempty" xml:"GrowInterval,omitempty"`
	// Specifies whether to enable the topology awareness feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// 4
	IdleInterval *int32 `json:"IdleInterval,omitempty" xml:"IdleInterval,omitempty"`
	// The interval at which the cluster is automatically scaled out.
	//
	// example:
	//
	// 10000
	MaxCoreCount *int32 `json:"MaxCoreCount,omitempty" xml:"MaxCoreCount,omitempty"`
	// The arguments that are used to run the post-processing script.
	//
	// example:
	//
	// 500
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The monitoring details of the cluster.
	MonitorSpec *UpdateClusterRequestMonitorSpec `json:"MonitorSpec,omitempty" xml:"MonitorSpec,omitempty" type:"Struct"`
	// The scheduler specifications of the cluster.
	SchedulerSpec *UpdateClusterRequestSchedulerSpec `json:"SchedulerSpec,omitempty" xml:"SchedulerSpec,omitempty" type:"Struct"`
}

func (s UpdateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterRequest) SetClientVersion(v string) *UpdateClusterRequest {
	s.ClientVersion = &v
	return s
}

func (s *UpdateClusterRequest) SetClusterCustomConfiguration(v *UpdateClusterRequestClusterCustomConfiguration) *UpdateClusterRequest {
	s.ClusterCustomConfiguration = v
	return s
}

func (s *UpdateClusterRequest) SetClusterDescription(v string) *UpdateClusterRequest {
	s.ClusterDescription = &v
	return s
}

func (s *UpdateClusterRequest) SetClusterId(v string) *UpdateClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterRequest) SetClusterName(v string) *UpdateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *UpdateClusterRequest) SetDeletionProtection(v bool) *UpdateClusterRequest {
	s.DeletionProtection = &v
	return s
}

func (s *UpdateClusterRequest) SetEnableScaleIn(v bool) *UpdateClusterRequest {
	s.EnableScaleIn = &v
	return s
}

func (s *UpdateClusterRequest) SetEnableScaleOut(v bool) *UpdateClusterRequest {
	s.EnableScaleOut = &v
	return s
}

func (s *UpdateClusterRequest) SetGrowInterval(v int32) *UpdateClusterRequest {
	s.GrowInterval = &v
	return s
}

func (s *UpdateClusterRequest) SetIdleInterval(v int32) *UpdateClusterRequest {
	s.IdleInterval = &v
	return s
}

func (s *UpdateClusterRequest) SetMaxCoreCount(v int32) *UpdateClusterRequest {
	s.MaxCoreCount = &v
	return s
}

func (s *UpdateClusterRequest) SetMaxCount(v int32) *UpdateClusterRequest {
	s.MaxCount = &v
	return s
}

func (s *UpdateClusterRequest) SetMonitorSpec(v *UpdateClusterRequestMonitorSpec) *UpdateClusterRequest {
	s.MonitorSpec = v
	return s
}

func (s *UpdateClusterRequest) SetSchedulerSpec(v *UpdateClusterRequestSchedulerSpec) *UpdateClusterRequest {
	s.SchedulerSpec = v
	return s
}

type UpdateClusterRequestClusterCustomConfiguration struct {
	// Specifies whether to enable the monitoring component of compute nodes. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// E-HPC cn-hangzhou
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The monitoring details of the cluster.
	//
	// example:
	//
	// http://*****
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s UpdateClusterRequestClusterCustomConfiguration) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterRequestClusterCustomConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateClusterRequestClusterCustomConfiguration) SetArgs(v string) *UpdateClusterRequestClusterCustomConfiguration {
	s.Args = &v
	return s
}

func (s *UpdateClusterRequestClusterCustomConfiguration) SetScript(v string) *UpdateClusterRequestClusterCustomConfiguration {
	s.Script = &v
	return s
}

type UpdateClusterRequestMonitorSpec struct {
	// Specifies whether to enable the monitoring component of compute nodes. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableComputeLoadMonitor *bool `json:"EnableComputeLoadMonitor,omitempty" xml:"EnableComputeLoadMonitor,omitempty"`
}

func (s UpdateClusterRequestMonitorSpec) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterRequestMonitorSpec) GoString() string {
	return s.String()
}

func (s *UpdateClusterRequestMonitorSpec) SetEnableComputeLoadMonitor(v bool) *UpdateClusterRequestMonitorSpec {
	s.EnableComputeLoadMonitor = &v
	return s
}

type UpdateClusterRequestSchedulerSpec struct {
	// Specifies whether to enable the topology awareness feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableTopologyAwareness *bool `json:"EnableTopologyAwareness,omitempty" xml:"EnableTopologyAwareness,omitempty"`
}

func (s UpdateClusterRequestSchedulerSpec) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterRequestSchedulerSpec) GoString() string {
	return s.String()
}

func (s *UpdateClusterRequestSchedulerSpec) SetEnableTopologyAwareness(v bool) *UpdateClusterRequestSchedulerSpec {
	s.EnableTopologyAwareness = &v
	return s
}

type UpdateClusterShrinkRequest struct {
	// Specifies whether to enable auto scale-out for the cluster. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// 2.1.0
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// Specifies whether to enable auto scale-in for the cluster. Valid values:
	//
	// 	- true
	//
	// 	- false
	ClusterCustomConfigurationShrink *string `json:"ClusterCustomConfiguration,omitempty" xml:"ClusterCustomConfiguration,omitempty"`
	// The URL that is used to download the post-processing script.
	//
	// example:
	//
	// slurm22.05.8-serverless-cluster-20240805
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// The client version. By default, the latest version is used.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The post-processing script of the cluster.
	//
	// example:
	//
	// slurm22.05.8-serverless-cluster-20240805
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The idle duration of the compute nodes allowed by the cluster.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The request result. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	EnableScaleIn *bool `json:"EnableScaleIn,omitempty" xml:"EnableScaleIn,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// true
	EnableScaleOut *bool `json:"EnableScaleOut,omitempty" xml:"EnableScaleOut,omitempty"`
	// The scheduler specifications of the cluster.
	//
	// example:
	//
	// 2
	GrowInterval *int32 `json:"GrowInterval,omitempty" xml:"GrowInterval,omitempty"`
	// Specifies whether to enable the topology awareness feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// 4
	IdleInterval *int32 `json:"IdleInterval,omitempty" xml:"IdleInterval,omitempty"`
	// The interval at which the cluster is automatically scaled out.
	//
	// example:
	//
	// 10000
	MaxCoreCount *int32 `json:"MaxCoreCount,omitempty" xml:"MaxCoreCount,omitempty"`
	// The arguments that are used to run the post-processing script.
	//
	// example:
	//
	// 500
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The monitoring details of the cluster.
	MonitorSpecShrink *string `json:"MonitorSpec,omitempty" xml:"MonitorSpec,omitempty"`
	// The scheduler specifications of the cluster.
	SchedulerSpecShrink *string `json:"SchedulerSpec,omitempty" xml:"SchedulerSpec,omitempty"`
}

func (s UpdateClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterShrinkRequest) SetClientVersion(v string) *UpdateClusterShrinkRequest {
	s.ClientVersion = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetClusterCustomConfigurationShrink(v string) *UpdateClusterShrinkRequest {
	s.ClusterCustomConfigurationShrink = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetClusterDescription(v string) *UpdateClusterShrinkRequest {
	s.ClusterDescription = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetClusterId(v string) *UpdateClusterShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetClusterName(v string) *UpdateClusterShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetDeletionProtection(v bool) *UpdateClusterShrinkRequest {
	s.DeletionProtection = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetEnableScaleIn(v bool) *UpdateClusterShrinkRequest {
	s.EnableScaleIn = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetEnableScaleOut(v bool) *UpdateClusterShrinkRequest {
	s.EnableScaleOut = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetGrowInterval(v int32) *UpdateClusterShrinkRequest {
	s.GrowInterval = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetIdleInterval(v int32) *UpdateClusterShrinkRequest {
	s.IdleInterval = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetMaxCoreCount(v int32) *UpdateClusterShrinkRequest {
	s.MaxCoreCount = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetMaxCount(v int32) *UpdateClusterShrinkRequest {
	s.MaxCount = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetMonitorSpecShrink(v string) *UpdateClusterShrinkRequest {
	s.MonitorSpecShrink = &v
	return s
}

func (s *UpdateClusterShrinkRequest) SetSchedulerSpecShrink(v string) *UpdateClusterShrinkRequest {
	s.SchedulerSpecShrink = &v
	return s
}

type UpdateClusterResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Request result, possible values:
	//
	// - true: request succeeded
	//
	// - false: request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterResponseBody) SetRequestId(v string) *UpdateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClusterResponseBody) SetSuccess(v bool) *UpdateClusterResponseBody {
	s.Success = &v
	return s
}

type UpdateClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterResponse) SetHeaders(v map[string]*string) *UpdateClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterResponse) SetStatusCode(v int32) *UpdateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClusterResponse) SetBody(v *UpdateClusterResponseBody) *UpdateClusterResponse {
	s.Body = v
	return s
}

type UpdateNodesRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about the compute nodes that you want to update.
	Instances []*UpdateNodesRequestInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
}

func (s UpdateNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodesRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodesRequest) SetClusterId(v string) *UpdateNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateNodesRequest) SetInstances(v []*UpdateNodesRequestInstances) *UpdateNodesRequest {
	s.Instances = v
	return s
}

type UpdateNodesRequestInstances struct {
	// The instance ID of the compute node.
	//
	// example:
	//
	// i-bp1bzqq1ddeemuddn****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to enable deletion protection for the compute node.
	//
	// example:
	//
	// true
	KeepAlive *bool `json:"KeepAlive,omitempty" xml:"KeepAlive,omitempty"`
}

func (s UpdateNodesRequestInstances) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodesRequestInstances) GoString() string {
	return s.String()
}

func (s *UpdateNodesRequestInstances) SetInstanceId(v string) *UpdateNodesRequestInstances {
	s.InstanceId = &v
	return s
}

func (s *UpdateNodesRequestInstances) SetKeepAlive(v bool) *UpdateNodesRequestInstances {
	s.KeepAlive = &v
	return s
}

type UpdateNodesShrinkRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about the compute nodes that you want to update.
	InstancesShrink *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
}

func (s UpdateNodesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodesShrinkRequest) SetClusterId(v string) *UpdateNodesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateNodesShrinkRequest) SetInstancesShrink(v string) *UpdateNodesShrinkRequest {
	s.InstancesShrink = &v
	return s
}

type UpdateNodesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodesResponseBody) SetRequestId(v string) *UpdateNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodesResponseBody) SetSuccess(v bool) *UpdateNodesResponseBody {
	s.Success = &v
	return s
}

type UpdateNodesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodesResponse) GoString() string {
	return s.String()
}

func (s *UpdateNodesResponse) SetHeaders(v map[string]*string) *UpdateNodesResponse {
	s.Headers = v
	return s
}

func (s *UpdateNodesResponse) SetStatusCode(v int32) *UpdateNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNodesResponse) SetBody(v *UpdateNodesResponseBody) *UpdateNodesResponse {
	s.Body = v
	return s
}

type UpdateQueueRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about the queue to be updated.
	Queue *UpdateQueueRequestQueue `json:"Queue,omitempty" xml:"Queue,omitempty" type:"Struct"`
}

func (s UpdateQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQueueRequest) GoString() string {
	return s.String()
}

func (s *UpdateQueueRequest) SetClusterId(v string) *UpdateQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateQueueRequest) SetQueue(v *UpdateQueueRequestQueue) *UpdateQueueRequest {
	s.Queue = v
	return s
}

type UpdateQueueRequestQueue struct {
	// The policy based on which instance types are selected for compute nodes during auto scale-outs. Valid values:
	//
	// 	- PriorityInstanceType
	//
	// example:
	//
	// PriorityInstanceType
	AllocationStrategy *string `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	// The hardware configurations of the compute nodes in the queue. Valid values of N: 1 to 10.
	ComputeNodes []*NodeTemplate `json:"ComputeNodes,omitempty" xml:"ComputeNodes,omitempty" type:"Repeated"`
	// Specifies whether to enable auto scale-in for the queue. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableScaleIn *bool `json:"EnableScaleIn,omitempty" xml:"EnableScaleIn,omitempty"`
	// Specifies whether to enable auto scale-out for the queue. Valid values:
	//
	// 	- true: deletion protection is enabled.
	//
	// 	- false: disables ASM.
	//
	// example:
	//
	// true
	EnableScaleOut *bool `json:"EnableScaleOut,omitempty" xml:"EnableScaleOut,omitempty"`
	// The hostname prefix of the added compute nodes.
	//
	// example:
	//
	// compute
	HostnamePrefix *string `json:"HostnamePrefix,omitempty" xml:"HostnamePrefix,omitempty"`
	// The hostname suffix of the compute nodes in the queue.
	//
	// example:
	//
	// hpc
	HostnameSuffix *string `json:"HostnameSuffix,omitempty" xml:"HostnameSuffix,omitempty"`
	// The initial number of compute nodes in the queue.
	//
	// example:
	//
	// 0
	InitialCount *int32 `json:"InitialCount,omitempty" xml:"InitialCount,omitempty"`
	// The type of the network for interconnecting compute nodes in the queue.
	//
	// example:
	//
	// erdma
	InterConnect *string `json:"InterConnect,omitempty" xml:"InterConnect,omitempty"`
	// The nodes for which deletion protection is enabled in the queue.
	KeepAliveNodes []*string `json:"KeepAliveNodes,omitempty" xml:"KeepAliveNodes,omitempty" type:"Repeated"`
	// The maximum number of compute nodes that the queue can contain.
	//
	// example:
	//
	// 1000
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The minimum number of compute nodes that are added to the queue during an automatic scale-out.
	//
	// example:
	//
	// 99
	MaxCountPerCycle *int64 `json:"MaxCountPerCycle,omitempty" xml:"MaxCountPerCycle,omitempty"`
	// The minimum number of compute nodes that the queue must contain.
	//
	// example:
	//
	// 0
	MinCount *int32 `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// The queue name.
	//
	// This parameter is required.
	//
	// example:
	//
	// comp
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The Resource Access Management (RAM) role that is assumed by compute nodes in the queue.
	//
	// example:
	//
	// AliyunECSInstanceForEHPCRole
	RamRole            *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	ReservedNodePoolId *string `json:"ReservedNodePoolId,omitempty" xml:"ReservedNodePoolId,omitempty"`
	// The vSwitches available for use by compute nodes in the queue.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s UpdateQueueRequestQueue) String() string {
	return tea.Prettify(s)
}

func (s UpdateQueueRequestQueue) GoString() string {
	return s.String()
}

func (s *UpdateQueueRequestQueue) SetAllocationStrategy(v string) *UpdateQueueRequestQueue {
	s.AllocationStrategy = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetComputeNodes(v []*NodeTemplate) *UpdateQueueRequestQueue {
	s.ComputeNodes = v
	return s
}

func (s *UpdateQueueRequestQueue) SetEnableScaleIn(v bool) *UpdateQueueRequestQueue {
	s.EnableScaleIn = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetEnableScaleOut(v bool) *UpdateQueueRequestQueue {
	s.EnableScaleOut = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetHostnamePrefix(v string) *UpdateQueueRequestQueue {
	s.HostnamePrefix = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetHostnameSuffix(v string) *UpdateQueueRequestQueue {
	s.HostnameSuffix = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetInitialCount(v int32) *UpdateQueueRequestQueue {
	s.InitialCount = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetInterConnect(v string) *UpdateQueueRequestQueue {
	s.InterConnect = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetKeepAliveNodes(v []*string) *UpdateQueueRequestQueue {
	s.KeepAliveNodes = v
	return s
}

func (s *UpdateQueueRequestQueue) SetMaxCount(v int32) *UpdateQueueRequestQueue {
	s.MaxCount = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetMaxCountPerCycle(v int64) *UpdateQueueRequestQueue {
	s.MaxCountPerCycle = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetMinCount(v int32) *UpdateQueueRequestQueue {
	s.MinCount = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetQueueName(v string) *UpdateQueueRequestQueue {
	s.QueueName = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetRamRole(v string) *UpdateQueueRequestQueue {
	s.RamRole = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetReservedNodePoolId(v string) *UpdateQueueRequestQueue {
	s.ReservedNodePoolId = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetVSwitchIds(v []*string) *UpdateQueueRequestQueue {
	s.VSwitchIds = v
	return s
}

type UpdateQueueShrinkRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about the queue to be updated.
	QueueShrink *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
}

func (s UpdateQueueShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQueueShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateQueueShrinkRequest) SetClusterId(v string) *UpdateQueueShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateQueueShrinkRequest) SetQueueShrink(v string) *UpdateQueueShrinkRequest {
	s.QueueShrink = &v
	return s
}

type UpdateQueueResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateQueueResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQueueResponseBody) SetRequestId(v string) *UpdateQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQueueResponseBody) SetSuccess(v bool) *UpdateQueueResponseBody {
	s.Success = &v
	return s
}

type UpdateQueueResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateQueueResponse) GoString() string {
	return s.String()
}

func (s *UpdateQueueResponse) SetHeaders(v map[string]*string) *UpdateQueueResponse {
	s.Headers = v
	return s
}

func (s *UpdateQueueResponse) SetStatusCode(v int32) *UpdateQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQueueResponse) SetBody(v *UpdateQueueResponseBody) *UpdateQueueResponse {
	s.Body = v
	return s
}

type UpdateUserRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The user group attribute of the user that you want to update. Valid values:
	//
	// 	- users: ordinary permissions, which are suitable for ordinary users that need only to submit and debug jobs.
	//
	// 	- wheel: sudo permissions, which are suitable for administrators who need to manage clusters. In addition to submitting and debugging jobs, you can also run sudo commands to install software and restart nodes.
	//
	// example:
	//
	// users
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The password attribute of the user that you want to update. The password must be 6 to 30 characters in length and must contain three of the following four character types:
	//
	// 	- Uppercase letters
	//
	// 	- Lowercase letters
	//
	// 	- Digits
	//
	// 	- Special characters ()~!@#$%^&\\*-_+=|{}[]:;\\"/<>,.?/
	//
	// example:
	//
	// 123****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// testuser
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) SetClusterId(v string) *UpdateUserRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateUserRequest) SetGroup(v string) *UpdateUserRequest {
	s.Group = &v
	return s
}

func (s *UpdateUserRequest) SetPassword(v string) *UpdateUserRequest {
	s.Password = &v
	return s
}

func (s *UpdateUserRequest) SetUserName(v string) *UpdateUserRequest {
	s.UserName = &v
	return s
}

type UpdateUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponseBody) SetSuccess(v string) *UpdateUserResponseBody {
	s.Success = &v
	return s
}

type UpdateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) SetHeaders(v map[string]*string) *UpdateUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserResponse) SetStatusCode(v int32) *UpdateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserResponse) SetBody(v *UpdateUserResponseBody) *UpdateUserResponse {
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

// Summary:
//
// Adds Elastic Compute Service (ECS) instances as compute nodes to Elastic High Performance Computing (E-HPC) clusters.
//
// Description:
//
// The ECS instances must meet the following requirements:
//
//   - The ECS instances do not belong to any E-HPC cluster.
//
//   - The ECS instances reside in the same virtual private cloud (VPC) as the cluster.
//
//   - The ECS instances are in the Stopped state.
//
// Take of the following limits:
//
//   - You can specify multiple instance IDs to add them at a time. However, the instances must be of the same type.
//
//   - When an instance is added to the cluster, [the system disk is reset](https://help.aliyun.com/zh/ecs/user-guide/re-initialize-a-system-disk) by using the image specified by the input parameters.
//
//   - If the instance has data disks, they are not automatically created and mounted after the instance is added.
//
//   - The hostname of the instance remains the same. Therefore, you must ensure that the hostname of the instance to be added is different from the hostname of an existing node in the cluster.
//
// @param tmpReq - AttachNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachNodesResponse
func (client *Client) AttachNodesWithOptions(tmpReq *AttachNodesRequest, runtime *util.RuntimeOptions) (_result *AttachNodesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AttachNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ComputeNode)) {
		request.ComputeNodeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ComputeNode, tea.String("ComputeNode"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ComputeNodeShrink)) {
		query["ComputeNode"] = request.ComputeNodeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		query["QueueName"] = request.QueueName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachNodes"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds Elastic Compute Service (ECS) instances as compute nodes to Elastic High Performance Computing (E-HPC) clusters.
//
// Description:
//
// The ECS instances must meet the following requirements:
//
//   - The ECS instances do not belong to any E-HPC cluster.
//
//   - The ECS instances reside in the same virtual private cloud (VPC) as the cluster.
//
//   - The ECS instances are in the Stopped state.
//
// Take of the following limits:
//
//   - You can specify multiple instance IDs to add them at a time. However, the instances must be of the same type.
//
//   - When an instance is added to the cluster, [the system disk is reset](https://help.aliyun.com/zh/ecs/user-guide/re-initialize-a-system-disk) by using the image specified by the input parameters.
//
//   - If the instance has data disks, they are not automatically created and mounted after the instance is added.
//
//   - The hostname of the instance remains the same. Therefore, you must ensure that the hostname of the instance to be added is different from the hostname of an existing node in the cluster.
//
// @param request - AttachNodesRequest
//
// @return AttachNodesResponse
func (client *Client) AttachNodes(request *AttachNodesRequest) (_result *AttachNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachNodesResponse{}
	_body, _err := client.AttachNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches shared storage to an Elastic High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// When you call this operation, take note of the following items:
//
//   - The file system that you want to attach must be created in advance in the same virtual private cloud (VPC) as the destination cluster. For more information, see [Create a file system](https://help.aliyun.com/document_detail/27530.html) and [Manage mount targets](https://help.aliyun.com/document_detail/27531.html).
//
//   - E-HPC clusters support Apsara File Storage NAS file systems.
//
// @param tmpReq - AttachSharedStoragesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachSharedStoragesResponse
func (client *Client) AttachSharedStoragesWithOptions(tmpReq *AttachSharedStoragesRequest, runtime *util.RuntimeOptions) (_result *AttachSharedStoragesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AttachSharedStoragesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SharedStorages)) {
		request.SharedStoragesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SharedStorages, tea.String("SharedStorages"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.SharedStoragesShrink)) {
		query["SharedStorages"] = request.SharedStoragesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachSharedStorages"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachSharedStoragesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches shared storage to an Elastic High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// When you call this operation, take note of the following items:
//
//   - The file system that you want to attach must be created in advance in the same virtual private cloud (VPC) as the destination cluster. For more information, see [Create a file system](https://help.aliyun.com/document_detail/27530.html) and [Manage mount targets](https://help.aliyun.com/document_detail/27531.html).
//
//   - E-HPC clusters support Apsara File Storage NAS file systems.
//
// @param request - AttachSharedStoragesRequest
//
// @return AttachSharedStoragesResponse
func (client *Client) AttachSharedStorages(request *AttachSharedStoragesRequest) (_result *AttachSharedStoragesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachSharedStoragesResponse{}
	_body, _err := client.AttachSharedStoragesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a pay-as-you-go or subscription Elastic High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// Before you call this operation, make sure that you are familiar with the billing and pricing of E-HPC. For more information, see [Overview](https://help.aliyun.com/document_detail/2842985.html).
//
// @param tmpReq - CreateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithOptions(tmpReq *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AdditionalPackages)) {
		request.AdditionalPackagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdditionalPackages, tea.String("AdditionalPackages"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Addons)) {
		request.AddonsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Addons, tea.String("Addons"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ClusterCredentials)) {
		request.ClusterCredentialsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClusterCredentials, tea.String("ClusterCredentials"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ClusterCustomConfiguration)) {
		request.ClusterCustomConfigurationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClusterCustomConfiguration, tea.String("ClusterCustomConfiguration"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Manager)) {
		request.ManagerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Manager, tea.String("Manager"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Queues)) {
		request.QueuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Queues, tea.String("Queues"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SharedStorages)) {
		request.SharedStoragesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SharedStorages, tea.String("SharedStorages"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdditionalPackagesShrink)) {
		query["AdditionalPackages"] = request.AdditionalPackagesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AddonsShrink)) {
		query["Addons"] = request.AddonsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterCategory)) {
		query["ClusterCategory"] = request.ClusterCategory
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterCredentialsShrink)) {
		query["ClusterCredentials"] = request.ClusterCredentialsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterCustomConfigurationShrink)) {
		query["ClusterCustomConfiguration"] = request.ClusterCustomConfigurationShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterDescription)) {
		query["ClusterDescription"] = request.ClusterDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterMode)) {
		query["ClusterMode"] = request.ClusterMode
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterVSwitchId)) {
		query["ClusterVSwitchId"] = request.ClusterVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterVpcId)) {
		query["ClusterVpcId"] = request.ClusterVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.DeletionProtection)) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !tea.BoolValue(util.IsUnset(request.IsEnterpriseSecurityGroup)) {
		query["IsEnterpriseSecurityGroup"] = request.IsEnterpriseSecurityGroup
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerShrink)) {
		query["Manager"] = request.ManagerShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MaxCoreCount)) {
		query["MaxCoreCount"] = request.MaxCoreCount
	}

	if !tea.BoolValue(util.IsUnset(request.MaxCount)) {
		query["MaxCount"] = request.MaxCount
	}

	if !tea.BoolValue(util.IsUnset(request.QueuesShrink)) {
		query["Queues"] = request.QueuesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SharedStoragesShrink)) {
		query["SharedStorages"] = request.SharedStoragesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCluster"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
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

// Summary:
//
// Creates a pay-as-you-go or subscription Elastic High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// Before you call this operation, make sure that you are familiar with the billing and pricing of E-HPC. For more information, see [Overview](https://help.aliyun.com/document_detail/2842985.html).
//
// @param request - CreateClusterRequest
//
// @return CreateClusterResponse
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

// Summary:
//
// Creates a job for a cluster.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing and [pricing](https://www.aliyun.com/price/product#/ecs/detail) of E-HPC.
//
// @param tmpReq - CreateJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobResponse
func (client *Client) CreateJobWithOptions(tmpReq *CreateJobRequest, runtime *util.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.JobSpec)) {
		request.JobSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobSpec, tea.String("JobSpec"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.JobName)) {
		query["JobName"] = request.JobName
	}

	if !tea.BoolValue(util.IsUnset(request.JobSpecShrink)) {
		query["JobSpec"] = request.JobSpecShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateJob"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a job for a cluster.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing and [pricing](https://www.aliyun.com/price/product#/ecs/detail) of E-HPC.
//
// @param request - CreateJobRequest
//
// @return CreateJobResponse
func (client *Client) CreateJob(request *CreateJobRequest) (_result *CreateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateJobResponse{}
	_body, _err := client.CreateJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates compute nodes for an Elastic High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)
//
// @param tmpReq - CreateNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodesResponse
func (client *Client) CreateNodesWithOptions(tmpReq *CreateNodesRequest, runtime *util.RuntimeOptions) (_result *CreateNodesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ComputeNode)) {
		request.ComputeNodeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ComputeNode, tea.String("ComputeNode"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ComputeNodeShrink)) {
		query["ComputeNode"] = request.ComputeNodeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Count)) {
		query["Count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.DeploymentSetId)) {
		query["DeploymentSetId"] = request.DeploymentSetId
	}

	if !tea.BoolValue(util.IsUnset(request.HPCInterConnect)) {
		query["HPCInterConnect"] = request.HPCInterConnect
	}

	if !tea.BoolValue(util.IsUnset(request.HostnamePrefix)) {
		query["HostnamePrefix"] = request.HostnamePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.HostnameSuffix)) {
		query["HostnameSuffix"] = request.HostnameSuffix
	}

	if !tea.BoolValue(util.IsUnset(request.KeepAlive)) {
		query["KeepAlive"] = request.KeepAlive
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		query["QueueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.RamRole)) {
		query["RamRole"] = request.RamRole
	}

	if !tea.BoolValue(util.IsUnset(request.ReservedNodePoolId)) {
		query["ReservedNodePoolId"] = request.ReservedNodePoolId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNodes"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates compute nodes for an Elastic High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)
//
// @param request - CreateNodesRequest
//
// @return CreateNodesResponse
func (client *Client) CreateNodes(request *CreateNodesRequest) (_result *CreateNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNodesResponse{}
	_body, _err := client.CreateNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a queue for an Enterprise High Performance Computing (E-HPC) cluster.
//
// @param tmpReq - CreateQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQueueResponse
func (client *Client) CreateQueueWithOptions(tmpReq *CreateQueueRequest, runtime *util.RuntimeOptions) (_result *CreateQueueResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateQueueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Queue)) {
		request.QueueShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Queue, tea.String("Queue"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.QueueShrink)) {
		query["Queue"] = request.QueueShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateQueue"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a queue for an Enterprise High Performance Computing (E-HPC) cluster.
//
// @param request - CreateQueueRequest
//
// @return CreateQueueResponse
func (client *Client) CreateQueue(request *CreateQueueRequest) (_result *CreateQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateQueueResponse{}
	_body, _err := client.CreateQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds users to an Elastic High Performance Computing (E-HPC) cluster.
//
// @param tmpReq - CreateUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUsersResponse
func (client *Client) CreateUsersWithOptions(tmpReq *CreateUsersRequest, runtime *util.RuntimeOptions) (_result *CreateUsersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.User)) {
		request.UserShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.User, tea.String("User"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.UserShrink)) {
		query["User"] = request.UserShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUsers"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds users to an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - CreateUsersRequest
//
// @return CreateUsersResponse
func (client *Client) CreateUsers(request *CreateUsersRequest) (_result *CreateUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUsersResponse{}
	_body, _err := client.CreateUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases an Enterprise High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// Make sure that data of the cluster to be deleted is backed up before you call this operation.
//
// > After a cluster is released, you cannot restore the data stored in the cluster. Exercise caution when you release a cluster.
//
// @param request - DeleteClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterResponse
func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCluster"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
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

// Summary:
//
// Releases an Enterprise High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// Make sure that data of the cluster to be deleted is backed up before you call this operation.
//
// > After a cluster is released, you cannot restore the data stored in the cluster. Exercise caution when you release a cluster.
//
// @param request - DeleteClusterRequest
//
// @return DeleteClusterResponse
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

// Summary:
//
// Deletes compute nodes from an Enterprise High Performance Computing (E-HPC) cluster at a time.
//
// Description:
//
// ## [](#)Usage notes
//
// Before you delete a compute node, we recommend that you export all job data from the node to prevent data loss.
//
// @param tmpReq - DeleteNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNodesResponse
func (client *Client) DeleteNodesWithOptions(tmpReq *DeleteNodesRequest, runtime *util.RuntimeOptions) (_result *DeleteNodesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceIds)) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, tea.String("InstanceIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIdsShrink)) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNodes"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
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

// Summary:
//
// Deletes compute nodes from an Enterprise High Performance Computing (E-HPC) cluster at a time.
//
// Description:
//
// ## [](#)Usage notes
//
// Before you delete a compute node, we recommend that you export all job data from the node to prevent data loss.
//
// @param request - DeleteNodesRequest
//
// @return DeleteNodesResponse
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

// Summary:
//
// Deletes queues from an Enterprise High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// Before you delete a queue, you must delete all compute nodes in the queue.
//
// @param tmpReq - DeleteQueuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQueuesResponse
func (client *Client) DeleteQueuesWithOptions(tmpReq *DeleteQueuesRequest, runtime *util.RuntimeOptions) (_result *DeleteQueuesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteQueuesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.QueueNames)) {
		request.QueueNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueueNames, tea.String("QueueNames"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.QueueNamesShrink)) {
		query["QueueNames"] = request.QueueNamesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteQueues"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteQueuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes queues from an Enterprise High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// Before you delete a queue, you must delete all compute nodes in the queue.
//
// @param request - DeleteQueuesRequest
//
// @return DeleteQueuesResponse
func (client *Client) DeleteQueues(request *DeleteQueuesRequest) (_result *DeleteQueuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteQueuesResponse{}
	_body, _err := client.DeleteQueuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes users from a cluster.
//
// @param tmpReq - DeleteUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUsersResponse
func (client *Client) DeleteUsersWithOptions(tmpReq *DeleteUsersRequest, runtime *util.RuntimeOptions) (_result *DeleteUsersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.User)) {
		request.UserShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.User, tea.String("User"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUsers"),
		Version:     tea.String("2024-07-30"),
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

// Summary:
//
// Deletes users from a cluster.
//
// @param request - DeleteUsersRequest
//
// @return DeleteUsersResponse
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

// Summary:
//
// Queries the details of an addon template.
//
// @param request - DescribeAddonTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAddonTemplateResponse
func (client *Client) DescribeAddonTemplateWithOptions(request *DescribeAddonTemplateRequest, runtime *util.RuntimeOptions) (_result *DescribeAddonTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddonName)) {
		query["AddonName"] = request.AddonName
	}

	if !tea.BoolValue(util.IsUnset(request.AddonVersion)) {
		query["AddonVersion"] = request.AddonVersion
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

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAddonTemplate"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAddonTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an addon template.
//
// @param request - DescribeAddonTemplateRequest
//
// @return DescribeAddonTemplateResponse
func (client *Client) DescribeAddonTemplate(request *DescribeAddonTemplateRequest) (_result *DescribeAddonTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAddonTemplateResponse{}
	_body, _err := client.DescribeAddonTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unmounts shared storage from the mount directory of a cluster.
//
// @param tmpReq - DetachSharedStoragesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachSharedStoragesResponse
func (client *Client) DetachSharedStoragesWithOptions(tmpReq *DetachSharedStoragesRequest, runtime *util.RuntimeOptions) (_result *DetachSharedStoragesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetachSharedStoragesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SharedStorages)) {
		request.SharedStoragesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SharedStorages, tea.String("SharedStorages"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.SharedStoragesShrink)) {
		query["SharedStorages"] = request.SharedStoragesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachSharedStorages"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachSharedStoragesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unmounts shared storage from the mount directory of a cluster.
//
// @param request - DetachSharedStoragesRequest
//
// @return DetachSharedStoragesResponse
func (client *Client) DetachSharedStorages(request *DetachSharedStoragesRequest) (_result *DetachSharedStoragesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachSharedStoragesResponse{}
	_body, _err := client.DetachSharedStoragesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an installed addon.
//
// @param request - GetAddonRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAddonResponse
func (client *Client) GetAddonWithOptions(request *GetAddonRequest, runtime *util.RuntimeOptions) (_result *GetAddonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddonId)) {
		query["AddonId"] = request.AddonId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAddon"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAddonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an installed addon.
//
// @param request - GetAddonRequest
//
// @return GetAddonResponse
func (client *Client) GetAddon(request *GetAddonRequest) (_result *GetAddonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAddonResponse{}
	_body, _err := client.GetAddonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - GetClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterResponse
func (client *Client) GetClusterWithOptions(request *GetClusterRequest, runtime *util.RuntimeOptions) (_result *GetClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCluster"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - GetClusterRequest
//
// @return GetClusterResponse
func (client *Client) GetCluster(request *GetClusterRequest) (_result *GetClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClusterResponse{}
	_body, _err := client.GetClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query logs based on a request ID. Logs for specific actions can be queried thanks to an Action-Stage-Method three-layer log splitting structure.
//
// @param request - GetCommonLogDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCommonLogDetailResponse
func (client *Client) GetCommonLogDetailWithOptions(request *GetCommonLogDetailRequest, runtime *util.RuntimeOptions) (_result *GetCommonLogDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.HiddenProcess)) {
		query["HiddenProcess"] = request.HiddenProcess
	}

	if !tea.BoolValue(util.IsUnset(request.LogRequestId)) {
		query["LogRequestId"] = request.LogRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["To"] = request.To
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCommonLogDetail"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCommonLogDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query logs based on a request ID. Logs for specific actions can be queried thanks to an Action-Stage-Method three-layer log splitting structure.
//
// @param request - GetCommonLogDetailRequest
//
// @return GetCommonLogDetailResponse
func (client *Client) GetCommonLogDetail(request *GetCommonLogDetailRequest) (_result *GetCommonLogDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCommonLogDetailResponse{}
	_body, _err := client.GetCommonLogDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of a job.
//
// @param request - GetJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
func (client *Client) GetJobWithOptions(request *GetJobRequest, runtime *util.RuntimeOptions) (_result *GetJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJob"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a job.
//
// @param request - GetJobRequest
//
// @return GetJobResponse
func (client *Client) GetJob(request *GetJobRequest) (_result *GetJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the output logs of a job, including standard output logs and error output logs.
//
// Description:
//
// ## [](#)Usage notes
//
// Currently, only Slurm and PBS Pro schedulers for Standard Edition clusters are supported.
//
// @param request - GetJobLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobLogResponse
func (client *Client) GetJobLogWithOptions(request *GetJobLogRequest, runtime *util.RuntimeOptions) (_result *GetJobLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		query["LogType"] = request.LogType
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobLog"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the output logs of a job, including standard output logs and error output logs.
//
// Description:
//
// ## [](#)Usage notes
//
// Currently, only Slurm and PBS Pro schedulers for Standard Edition clusters are supported.
//
// @param request - GetJobLogRequest
//
// @return GetJobLogResponse
func (client *Client) GetJobLog(request *GetJobLogRequest) (_result *GetJobLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobLogResponse{}
	_body, _err := client.GetJobLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a queue in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - GetQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueueResponse
func (client *Client) GetQueueWithOptions(request *GetQueueRequest, runtime *util.RuntimeOptions) (_result *GetQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		query["QueueName"] = request.QueueName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQueue"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a queue in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - GetQueueRequest
//
// @return GetQueueResponse
func (client *Client) GetQueue(request *GetQueueRequest) (_result *GetQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQueueResponse{}
	_body, _err := client.GetQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs an addon.
//
// Description:
//
// ## [](#)Usage notes
//
// Take note of the following items when you call this operation:
//
//   - The cluster must be in the `Running` state.
//
//   - Clusters fall into two types:
//
//   - Regular clusters on Alibaba Cloud Public Cloud
//
//   - Managed clusters on Alibaba Cloud Public Cloud
//
// @param request - InstallAddonRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAddonResponse
func (client *Client) InstallAddonWithOptions(request *InstallAddonRequest, runtime *util.RuntimeOptions) (_result *InstallAddonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddonName)) {
		query["AddonName"] = request.AddonName
	}

	if !tea.BoolValue(util.IsUnset(request.AddonVersion)) {
		query["AddonVersion"] = request.AddonVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcesSpec)) {
		query["ResourcesSpec"] = request.ResourcesSpec
	}

	if !tea.BoolValue(util.IsUnset(request.ServicesSpec)) {
		query["ServicesSpec"] = request.ServicesSpec
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallAddon"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallAddonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs an addon.
//
// Description:
//
// ## [](#)Usage notes
//
// Take note of the following items when you call this operation:
//
//   - The cluster must be in the `Running` state.
//
//   - Clusters fall into two types:
//
//   - Regular clusters on Alibaba Cloud Public Cloud
//
//   - Managed clusters on Alibaba Cloud Public Cloud
//
// @param request - InstallAddonRequest
//
// @return InstallAddonResponse
func (client *Client) InstallAddon(request *InstallAddonRequest) (_result *InstallAddonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallAddonResponse{}
	_body, _err := client.InstallAddonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Install software for the specified cluster.
//
// Description:
//
// ## Interface Description
//
// When calling this interface, please note the following:
//
// - The cluster status must be `Running`.
//
// - If the cluster series is `Serverless`, ensure that there is at least one login node or compute node in the cluster; otherwise, software cannot be added to the target cluster.
//
// @param tmpReq - InstallSoftwaresRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallSoftwaresResponse
func (client *Client) InstallSoftwaresWithOptions(tmpReq *InstallSoftwaresRequest, runtime *util.RuntimeOptions) (_result *InstallSoftwaresResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &InstallSoftwaresShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AdditionalPackages)) {
		request.AdditionalPackagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdditionalPackages, tea.String("AdditionalPackages"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallSoftwares"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallSoftwaresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Install software for the specified cluster.
//
// Description:
//
// ## Interface Description
//
// When calling this interface, please note the following:
//
// - The cluster status must be `Running`.
//
// - If the cluster series is `Serverless`, ensure that there is at least one login node or compute node in the cluster; otherwise, software cannot be added to the target cluster.
//
// @param request - InstallSoftwaresRequest
//
// @return InstallSoftwaresResponse
func (client *Client) InstallSoftwares(request *InstallSoftwaresRequest) (_result *InstallSoftwaresResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallSoftwaresResponse{}
	_body, _err := client.InstallSoftwaresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries supported addon templates.
//
// @param request - ListAddonTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddonTemplatesResponse
func (client *Client) ListAddonTemplatesWithOptions(request *ListAddonTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListAddonTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddonNames)) {
		query["AddonNames"] = request.AddonNames
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterCategory)) {
		query["ClusterCategory"] = request.ClusterCategory
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
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAddonTemplates"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAddonTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries supported addon templates.
//
// @param request - ListAddonTemplatesRequest
//
// @return ListAddonTemplatesResponse
func (client *Client) ListAddonTemplates(request *ListAddonTemplatesRequest) (_result *ListAddonTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAddonTemplatesResponse{}
	_body, _err := client.ListAddonTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries installed addons of an Elastic High Performance Computing (E-HPC) cluster.
//
// @param tmpReq - ListAddonsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddonsResponse
func (client *Client) ListAddonsWithOptions(tmpReq *ListAddonsRequest, runtime *util.RuntimeOptions) (_result *ListAddonsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListAddonsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AddonIds)) {
		request.AddonIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddonIds, tea.String("AddonIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddonIdsShrink)) {
		query["AddonIds"] = request.AddonIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAddons"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAddonsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries installed addons of an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - ListAddonsRequest
//
// @return ListAddonsResponse
func (client *Client) ListAddons(request *ListAddonsRequest) (_result *ListAddonsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAddonsResponse{}
	_body, _err := client.ListAddonsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the file systems that can be attached in a region.
//
// @param request - ListAvailableFileSystemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvailableFileSystemsResponse
func (client *Client) ListAvailableFileSystemsWithOptions(request *ListAvailableFileSystemsRequest, runtime *util.RuntimeOptions) (_result *ListAvailableFileSystemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAvailableFileSystems"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAvailableFileSystemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the file systems that can be attached in a region.
//
// @param request - ListAvailableFileSystemsRequest
//
// @return ListAvailableFileSystemsResponse
func (client *Client) ListAvailableFileSystems(request *ListAvailableFileSystemsRequest) (_result *ListAvailableFileSystemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAvailableFileSystemsResponse{}
	_body, _err := client.ListAvailableFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries images that are available for Elastic High Performance Computing (E-HPC) clusters.
//
// @param tmpReq - ListAvailableImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvailableImagesResponse
func (client *Client) ListAvailableImagesWithOptions(tmpReq *ListAvailableImagesRequest, runtime *util.RuntimeOptions) (_result *ListAvailableImagesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListAvailableImagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DirectoryService)) {
		request.DirectoryServiceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DirectoryService, tea.String("DirectoryService"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Scheduler)) {
		request.SchedulerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Scheduler, tea.String("Scheduler"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAvailableImages"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAvailableImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries images that are available for Elastic High Performance Computing (E-HPC) clusters.
//
// @param request - ListAvailableImagesRequest
//
// @return ListAvailableImagesResponse
func (client *Client) ListAvailableImages(request *ListAvailableImagesRequest) (_result *ListAvailableImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAvailableImagesResponse{}
	_body, _err := client.ListAvailableImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all clusters of a user in each region.
//
// @param tmpReq - ListClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClustersResponse
func (client *Client) ListClustersWithOptions(tmpReq *ListClustersRequest, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListClustersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ClusterIds)) {
		request.ClusterIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClusterIds, tea.String("ClusterIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ClusterNames)) {
		request.ClusterNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClusterNames, tea.String("ClusterNames"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterIdsShrink)) {
		query["ClusterIds"] = request.ClusterIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterNamesShrink)) {
		query["ClusterNames"] = request.ClusterNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusters"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
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

// Summary:
//
// Queries all clusters of a user in each region.
//
// @param request - ListClustersRequest
//
// @return ListClustersResponse
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

// Summary:
//
// Queries the logs of a cluster that are generated within a time range.
//
// @param tmpReq - ListCommonLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCommonLogsResponse
func (client *Client) ListCommonLogsWithOptions(tmpReq *ListCommonLogsRequest, runtime *util.RuntimeOptions) (_result *ListCommonLogsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListCommonLogsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ActionName)) {
		request.ActionNameShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ActionName, tea.String("ActionName"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionNameShrink)) {
		query["ActionName"] = request.ActionNameShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ActionStatus)) {
		query["ActionStatus"] = request.ActionStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.IsReverse)) {
		query["IsReverse"] = request.IsReverse
	}

	if !tea.BoolValue(util.IsUnset(request.LogRequestId)) {
		query["LogRequestId"] = request.LogRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		query["LogType"] = request.LogType
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		query["OperatorUid"] = request.OperatorUid
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["To"] = request.To
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCommonLogs"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCommonLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logs of a cluster that are generated within a time range.
//
// @param request - ListCommonLogsRequest
//
// @return ListCommonLogsResponse
func (client *Client) ListCommonLogs(request *ListCommonLogsRequest) (_result *ListCommonLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCommonLogsResponse{}
	_body, _err := client.ListCommonLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the installed software of a cluster.
//
// @param request - ListInstalledSoftwaresRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstalledSoftwaresResponse
func (client *Client) ListInstalledSoftwaresWithOptions(request *ListInstalledSoftwaresRequest, runtime *util.RuntimeOptions) (_result *ListInstalledSoftwaresResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstalledSoftwares"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstalledSoftwaresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the installed software of a cluster.
//
// @param request - ListInstalledSoftwaresRequest
//
// @return ListInstalledSoftwaresResponse
func (client *Client) ListInstalledSoftwares(request *ListInstalledSoftwaresRequest) (_result *ListInstalledSoftwaresResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstalledSoftwaresResponse{}
	_body, _err := client.ListInstalledSoftwaresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the jobs in a cluster.
//
// @param tmpReq - ListJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
func (client *Client) ListJobsWithOptions(tmpReq *ListJobsRequest, runtime *util.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.JobFilter)) {
		request.JobFilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobFilter, tea.String("JobFilter"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.JobFilterShrink)) {
		query["JobFilter"] = request.JobFilterShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobs"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
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

// Summary:
//
// Queries the jobs in a cluster.
//
// @param request - ListJobsRequest
//
// @return ListJobsResponse
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

// Summary:
//
// Queries the nodes of an Elastic High Performance Computing (E-HPC) cluster.
//
// @param tmpReq - ListNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithOptions(tmpReq *ListNodesRequest, runtime *util.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Hostnames)) {
		request.HostnamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Hostnames, tea.String("Hostnames"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PrivateIpAddress)) {
		request.PrivateIpAddressShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PrivateIpAddress, tea.String("PrivateIpAddress"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.QueueNames)) {
		request.QueueNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueueNames, tea.String("QueueNames"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Status)) {
		request.StatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Status, tea.String("Status"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.HostnamesShrink)) {
		query["Hostnames"] = request.HostnamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIpAddressShrink)) {
		query["PrivateIpAddress"] = request.PrivateIpAddressShrink
	}

	if !tea.BoolValue(util.IsUnset(request.QueueNamesShrink)) {
		query["QueueNames"] = request.QueueNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Sequence)) {
		query["Sequence"] = request.Sequence
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.StatusShrink)) {
		query["Status"] = request.StatusShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodes"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
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

// Summary:
//
// Queries the nodes of an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - ListNodesRequest
//
// @return ListNodesResponse
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

// Summary:
//
// Queries queues in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param tmpReq - ListQueuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueuesResponse
func (client *Client) ListQueuesWithOptions(tmpReq *ListQueuesRequest, runtime *util.RuntimeOptions) (_result *ListQueuesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListQueuesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.QueueNames)) {
		request.QueueNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueueNames, tea.String("QueueNames"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.QueueNamesShrink)) {
		query["QueueNames"] = request.QueueNamesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueues"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
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

// Summary:
//
// Queries queues in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - ListQueuesRequest
//
// @return ListQueuesResponse
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

// Summary:
//
// Queries the shared storage that is attached to a cluster.
//
// @param request - ListSharedStoragesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSharedStoragesResponse
func (client *Client) ListSharedStoragesWithOptions(request *ListSharedStoragesRequest, runtime *util.RuntimeOptions) (_result *ListSharedStoragesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemType)) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSharedStorages"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSharedStoragesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the shared storage that is attached to a cluster.
//
// @param request - ListSharedStoragesRequest
//
// @return ListSharedStoragesResponse
func (client *Client) ListSharedStorages(request *ListSharedStoragesRequest) (_result *ListSharedStoragesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSharedStoragesResponse{}
	_body, _err := client.ListSharedStoragesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the software that can be installed in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - ListSoftwaresRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSoftwaresResponse
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
		Version:     tea.String("2024-07-30"),
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

// Summary:
//
// Queries the software that can be installed in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - ListSoftwaresRequest
//
// @return ListSoftwaresResponse
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

// Summary:
//
// Queries the users of a cluster.
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
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
		Version:     tea.String("2024-07-30"),
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

// Summary:
//
// Queries the users of a cluster.
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
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

// Summary:
//
// Stops uncompleted jobs in a batch in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param tmpReq - StopJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopJobsResponse
func (client *Client) StopJobsWithOptions(tmpReq *StopJobsRequest, runtime *util.RuntimeOptions) (_result *StopJobsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &StopJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.JobIds)) {
		request.JobIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobIds, tea.String("JobIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.JobIdsShrink)) {
		query["JobIds"] = request.JobIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopJobs"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
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

// Summary:
//
// Stops uncompleted jobs in a batch in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - StopJobsRequest
//
// @return StopJobsResponse
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

// Summary:
//
// Uninstalls an addon.
//
// Description:
//
// ## [](#)Usage notes
//
// Take note of the following items when you call this operation:
//
//   - The cluster must be in the `Running` state.
//
//   - Clusters fall into the following types:
//
//   - Regular clusters on Alibaba Cloud Public Cloud
//
//   - Managed clusters on Alibaba Cloud Public Cloud
//
// @param request - UnInstallAddonRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnInstallAddonResponse
func (client *Client) UnInstallAddonWithOptions(request *UnInstallAddonRequest, runtime *util.RuntimeOptions) (_result *UnInstallAddonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddonId)) {
		query["AddonId"] = request.AddonId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnInstallAddon"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnInstallAddonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls an addon.
//
// Description:
//
// ## [](#)Usage notes
//
// Take note of the following items when you call this operation:
//
//   - The cluster must be in the `Running` state.
//
//   - Clusters fall into the following types:
//
//   - Regular clusters on Alibaba Cloud Public Cloud
//
//   - Managed clusters on Alibaba Cloud Public Cloud
//
// @param request - UnInstallAddonRequest
//
// @return UnInstallAddonResponse
func (client *Client) UnInstallAddon(request *UnInstallAddonRequest) (_result *UnInstallAddonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnInstallAddonResponse{}
	_body, _err := client.UnInstallAddonWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uninstalls software systems from an Enterprise High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## Interface Description
//
// When calling this interface, please note:
//
// The cluster status must be `Running`.
//
// @param tmpReq - UninstallSoftwaresRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallSoftwaresResponse
func (client *Client) UninstallSoftwaresWithOptions(tmpReq *UninstallSoftwaresRequest, runtime *util.RuntimeOptions) (_result *UninstallSoftwaresResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UninstallSoftwaresShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AdditionalPackages)) {
		request.AdditionalPackagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdditionalPackages, tea.String("AdditionalPackages"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UninstallSoftwares"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UninstallSoftwaresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls software systems from an Enterprise High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## Interface Description
//
// When calling this interface, please note:
//
// The cluster status must be `Running`.
//
// @param request - UninstallSoftwaresRequest
//
// @return UninstallSoftwaresResponse
func (client *Client) UninstallSoftwares(request *UninstallSoftwaresRequest) (_result *UninstallSoftwaresResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UninstallSoftwaresResponse{}
	_body, _err := client.UninstallSoftwaresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modify the basic information of a specified cluster.
//
// @param tmpReq - UpdateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClusterResponse
func (client *Client) UpdateClusterWithOptions(tmpReq *UpdateClusterRequest, runtime *util.RuntimeOptions) (_result *UpdateClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ClusterCustomConfiguration)) {
		request.ClusterCustomConfigurationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClusterCustomConfiguration, tea.String("ClusterCustomConfiguration"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.MonitorSpec)) {
		request.MonitorSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MonitorSpec, tea.String("MonitorSpec"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SchedulerSpec)) {
		request.SchedulerSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SchedulerSpec, tea.String("SchedulerSpec"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterCustomConfigurationShrink)) {
		query["ClusterCustomConfiguration"] = request.ClusterCustomConfigurationShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterDescription)) {
		query["ClusterDescription"] = request.ClusterDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.DeletionProtection)) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !tea.BoolValue(util.IsUnset(request.EnableScaleIn)) {
		query["EnableScaleIn"] = request.EnableScaleIn
	}

	if !tea.BoolValue(util.IsUnset(request.EnableScaleOut)) {
		query["EnableScaleOut"] = request.EnableScaleOut
	}

	if !tea.BoolValue(util.IsUnset(request.GrowInterval)) {
		query["GrowInterval"] = request.GrowInterval
	}

	if !tea.BoolValue(util.IsUnset(request.IdleInterval)) {
		query["IdleInterval"] = request.IdleInterval
	}

	if !tea.BoolValue(util.IsUnset(request.MaxCoreCount)) {
		query["MaxCoreCount"] = request.MaxCoreCount
	}

	if !tea.BoolValue(util.IsUnset(request.MaxCount)) {
		query["MaxCount"] = request.MaxCount
	}

	if !tea.BoolValue(util.IsUnset(request.MonitorSpecShrink)) {
		query["MonitorSpec"] = request.MonitorSpecShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SchedulerSpecShrink)) {
		query["SchedulerSpec"] = request.SchedulerSpecShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCluster"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the basic information of a specified cluster.
//
// @param request - UpdateClusterRequest
//
// @return UpdateClusterResponse
func (client *Client) UpdateCluster(request *UpdateClusterRequest) (_result *UpdateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateClusterResponse{}
	_body, _err := client.UpdateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configurations of compute nodes in an Enterprise High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// Before you delete a compute node, we recommend that you export all job data from the node to prevent data loss.
//
// @param tmpReq - UpdateNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNodesResponse
func (client *Client) UpdateNodesWithOptions(tmpReq *UpdateNodesRequest, runtime *util.RuntimeOptions) (_result *UpdateNodesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Instances)) {
		request.InstancesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Instances, tea.String("Instances"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.InstancesShrink)) {
		query["Instances"] = request.InstancesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNodes"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of compute nodes in an Enterprise High Performance Computing (E-HPC) cluster.
//
// Description:
//
// ## [](#)Usage notes
//
// Before you delete a compute node, we recommend that you export all job data from the node to prevent data loss.
//
// @param request - UpdateNodesRequest
//
// @return UpdateNodesResponse
func (client *Client) UpdateNodes(request *UpdateNodesRequest) (_result *UpdateNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNodesResponse{}
	_body, _err := client.UpdateNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a queue in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param tmpReq - UpdateQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQueueResponse
func (client *Client) UpdateQueueWithOptions(tmpReq *UpdateQueueRequest, runtime *util.RuntimeOptions) (_result *UpdateQueueResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateQueueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Queue)) {
		request.QueueShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Queue, tea.String("Queue"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.QueueShrink)) {
		query["Queue"] = request.QueueShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateQueue"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a queue in an Elastic High Performance Computing (E-HPC) cluster.
//
// @param request - UpdateQueueRequest
//
// @return UpdateQueueResponse
func (client *Client) UpdateQueue(request *UpdateQueueRequest) (_result *UpdateQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateQueueResponse{}
	_body, _err := client.UpdateQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information of a user in an Elastic High Performance Computing (E-HPC) cluster, including the user group and password.
//
// @param request - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Group)) {
		query["Group"] = request.Group
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUser"),
		Version:     tea.String("2024-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information of a user in an Elastic High Performance Computing (E-HPC) cluster, including the user group and password.
//
// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
