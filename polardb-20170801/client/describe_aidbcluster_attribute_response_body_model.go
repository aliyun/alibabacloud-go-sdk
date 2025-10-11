// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAiNodeType(v string) *DescribeAIDBClusterAttributeResponseBody
	GetAiNodeType() *string
	SetCreationTime(v string) *DescribeAIDBClusterAttributeResponseBody
	GetCreationTime() *string
	SetDBClusterDescription(v string) *DescribeAIDBClusterAttributeResponseBody
	GetDBClusterDescription() *string
	SetDBClusterId(v string) *DescribeAIDBClusterAttributeResponseBody
	GetDBClusterId() *string
	SetDBClusterStatus(v string) *DescribeAIDBClusterAttributeResponseBody
	GetDBClusterStatus() *string
	SetDBNodes(v []*DescribeAIDBClusterAttributeResponseBodyDBNodes) *DescribeAIDBClusterAttributeResponseBody
	GetDBNodes() []*DescribeAIDBClusterAttributeResponseBodyDBNodes
	SetDBVersion(v string) *DescribeAIDBClusterAttributeResponseBody
	GetDBVersion() *string
	SetEndpointList(v []*DescribeAIDBClusterAttributeResponseBodyEndpointList) *DescribeAIDBClusterAttributeResponseBody
	GetEndpointList() []*DescribeAIDBClusterAttributeResponseBodyEndpointList
	SetExpireTime(v string) *DescribeAIDBClusterAttributeResponseBody
	GetExpireTime() *string
	SetExpired(v bool) *DescribeAIDBClusterAttributeResponseBody
	GetExpired() *bool
	SetInternalIp(v string) *DescribeAIDBClusterAttributeResponseBody
	GetInternalIp() *string
	SetKubeClusterId(v string) *DescribeAIDBClusterAttributeResponseBody
	GetKubeClusterId() *string
	SetLockMode(v string) *DescribeAIDBClusterAttributeResponseBody
	GetLockMode() *string
	SetMaxQPM(v string) *DescribeAIDBClusterAttributeResponseBody
	GetMaxQPM() *string
	SetModelName(v string) *DescribeAIDBClusterAttributeResponseBody
	GetModelName() *string
	SetPayType(v string) *DescribeAIDBClusterAttributeResponseBody
	GetPayType() *string
	SetPublicIp(v string) *DescribeAIDBClusterAttributeResponseBody
	GetPublicIp() *string
	SetRegionId(v string) *DescribeAIDBClusterAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeAIDBClusterAttributeResponseBody
	GetRequestId() *string
	SetRunType(v string) *DescribeAIDBClusterAttributeResponseBody
	GetRunType() *string
	SetStorageType(v string) *DescribeAIDBClusterAttributeResponseBody
	GetStorageType() *string
	SetVPCId(v string) *DescribeAIDBClusterAttributeResponseBody
	GetVPCId() *string
	SetVSwitchId(v string) *DescribeAIDBClusterAttributeResponseBody
	GetVSwitchId() *string
	SetVolumes(v []*DescribeAIDBClusterAttributeResponseBodyVolumes) *DescribeAIDBClusterAttributeResponseBody
	GetVolumes() []*DescribeAIDBClusterAttributeResponseBodyVolumes
	SetZoneId(v string) *DescribeAIDBClusterAttributeResponseBody
	GetZoneId() *string
	SetZoneIds(v string) *DescribeAIDBClusterAttributeResponseBody
	GetZoneIds() *string
}

type DescribeAIDBClusterAttributeResponseBody struct {
	// example:
	//
	// vnode
	AiNodeType *string `json:"AiNodeType,omitempty" xml:"AiNodeType,omitempty"`
	// example:
	//
	// 2020-08-14T05:58:42Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// Running
	DBClusterStatus *string                                            `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	DBNodes         []*DescribeAIDBClusterAttributeResponseBodyDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
	// example:
	//
	// 1.0
	DBVersion    *string                                                 `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	EndpointList []*DescribeAIDBClusterAttributeResponseBodyEndpointList `json:"EndpointList,omitempty" xml:"EndpointList,omitempty" type:"Repeated"`
	// example:
	//
	// 2020-11-14T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// false
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// example:
	//
	// 10.*.*.72
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// example:
	//
	// xxxxxxxxxxxxxxxxxxxxxxx
	KubeClusterId *string `json:"KubeClusterId,omitempty" xml:"KubeClusterId,omitempty"`
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// example:
	//
	// 20
	MaxQPM *string `json:"MaxQPM,omitempty" xml:"MaxQPM,omitempty"`
	// example:
	//
	// Qwen3-Embedding-8B
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// Postpaid
	PayType  *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// EBEAA83D-1734-42E3-85E3-E25F6E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// container
	RunType *string `json:"RunType,omitempty" xml:"RunType,omitempty"`
	// example:
	//
	// essdpl1
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// example:
	//
	// vpc-*******************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-*********************
	VSwitchId *string                                            `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Volumes   []*DescribeAIDBClusterAttributeResponseBodyVolumes `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou-d
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// example:
	//
	// cn-hangzhou-i,cn-hangzhou-g
	ZoneIds *string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty"`
}

func (s DescribeAIDBClusterAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetAiNodeType() *string {
	return s.AiNodeType
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetDBNodes() []*DescribeAIDBClusterAttributeResponseBodyDBNodes {
	return s.DBNodes
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetEndpointList() []*DescribeAIDBClusterAttributeResponseBodyEndpointList {
	return s.EndpointList
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetExpired() *bool {
	return s.Expired
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetInternalIp() *string {
	return s.InternalIp
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetKubeClusterId() *string {
	return s.KubeClusterId
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetMaxQPM() *string {
	return s.MaxQPM
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetModelName() *string {
	return s.ModelName
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetPublicIp() *string {
	return s.PublicIp
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetRunType() *string {
	return s.RunType
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetVolumes() []*DescribeAIDBClusterAttributeResponseBodyVolumes {
	return s.Volumes
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAIDBClusterAttributeResponseBody) GetZoneIds() *string {
	return s.ZoneIds
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetAiNodeType(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.AiNodeType = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetCreationTime(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetDBClusterDescription(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetDBClusterId(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetDBClusterStatus(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetDBNodes(v []*DescribeAIDBClusterAttributeResponseBodyDBNodes) *DescribeAIDBClusterAttributeResponseBody {
	s.DBNodes = v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetDBVersion(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetEndpointList(v []*DescribeAIDBClusterAttributeResponseBodyEndpointList) *DescribeAIDBClusterAttributeResponseBody {
	s.EndpointList = v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetExpireTime(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetExpired(v bool) *DescribeAIDBClusterAttributeResponseBody {
	s.Expired = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetInternalIp(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.InternalIp = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetKubeClusterId(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.KubeClusterId = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetLockMode(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetMaxQPM(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.MaxQPM = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetModelName(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.ModelName = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetPayType(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetPublicIp(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.PublicIp = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetRegionId(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetRequestId(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetRunType(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.RunType = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetStorageType(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.StorageType = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetVPCId(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.VPCId = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetVSwitchId(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetVolumes(v []*DescribeAIDBClusterAttributeResponseBodyVolumes) *DescribeAIDBClusterAttributeResponseBody {
	s.Volumes = v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetZoneId(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) SetZoneIds(v string) *DescribeAIDBClusterAttributeResponseBody {
	s.ZoneIds = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAIDBClusterAttributeResponseBodyDBNodes struct {
	// example:
	//
	// 2
	CpuCores *string `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	// example:
	//
	// polar.pg.g8.8xlarge.gu30
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// example:
	//
	// test
	DBNodeDescription *string `json:"DBNodeDescription,omitempty" xml:"DBNodeDescription,omitempty"`
	// example:
	//
	// pi-****************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// example:
	//
	// Running
	DBNodeStatus *string `json:"DBNodeStatus,omitempty" xml:"DBNodeStatus,omitempty"`
	// example:
	//
	// 2
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// example:
	//
	// 10.*.*12
	LinkIP *string `json:"LinkIP,omitempty" xml:"LinkIP,omitempty"`
	// example:
	//
	// 8192
	MemorySize *string `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	// example:
	//
	// vn-***************
	VNodeId *string `json:"VNodeId,omitempty" xml:"VNodeId,omitempty"`
	// example:
	//
	// vpc-*******************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-*********************
	VSwitchId *string                                                   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Volumes   []*DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou-d
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAIDBClusterAttributeResponseBodyDBNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterAttributeResponseBodyDBNodes) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) GetCpuCores() *string {
	return s.CpuCores
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) GetDBNodeDescription() *string {
	return s.DBNodeDescription
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) GetDBNodeStatus() *string {
	return s.DBNodeStatus
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) GetGPU() *string {
	return s.GPU
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) GetLinkIP() *string {
	return s.LinkIP
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) GetMemorySize() *string {
	return s.MemorySize
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) GetVNodeId() *string {
	return s.VNodeId
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) GetVolumes() []*DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes {
	return s.Volumes
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) SetCpuCores(v string) *DescribeAIDBClusterAttributeResponseBodyDBNodes {
	s.CpuCores = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) SetDBNodeClass(v string) *DescribeAIDBClusterAttributeResponseBodyDBNodes {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) SetDBNodeDescription(v string) *DescribeAIDBClusterAttributeResponseBodyDBNodes {
	s.DBNodeDescription = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) SetDBNodeId(v string) *DescribeAIDBClusterAttributeResponseBodyDBNodes {
	s.DBNodeId = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) SetDBNodeStatus(v string) *DescribeAIDBClusterAttributeResponseBodyDBNodes {
	s.DBNodeStatus = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) SetGPU(v string) *DescribeAIDBClusterAttributeResponseBodyDBNodes {
	s.GPU = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) SetLinkIP(v string) *DescribeAIDBClusterAttributeResponseBodyDBNodes {
	s.LinkIP = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) SetMemorySize(v string) *DescribeAIDBClusterAttributeResponseBodyDBNodes {
	s.MemorySize = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) SetVNodeId(v string) *DescribeAIDBClusterAttributeResponseBodyDBNodes {
	s.VNodeId = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) SetVPCId(v string) *DescribeAIDBClusterAttributeResponseBodyDBNodes {
	s.VPCId = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) SetVSwitchId(v string) *DescribeAIDBClusterAttributeResponseBodyDBNodes {
	s.VSwitchId = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) SetVolumes(v []*DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes) *DescribeAIDBClusterAttributeResponseBodyDBNodes {
	s.Volumes = v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) SetZoneId(v string) *DescribeAIDBClusterAttributeResponseBodyDBNodes {
	s.ZoneId = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes struct {
	// example:
	//
	// /tmp/CrowdStrike
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// example:
	//
	// jueming
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 8192
	SizeGB *string `json:"SizeGB,omitempty" xml:"SizeGB,omitempty"`
	// example:
	//
	// PolarFs
	StorageCategory *string `json:"StorageCategory,omitempty" xml:"StorageCategory,omitempty"`
	// example:
	//
	// PL1
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes) GetName() *string {
	return s.Name
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes) GetSizeGB() *string {
	return s.SizeGB
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes) GetStorageCategory() *string {
	return s.StorageCategory
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes) SetMountPath(v string) *DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes {
	s.MountPath = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes) SetName(v string) *DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes {
	s.Name = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes) SetSizeGB(v string) *DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes {
	s.SizeGB = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes) SetStorageCategory(v string) *DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes {
	s.StorageCategory = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes) SetStorageType(v string) *DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes {
	s.StorageType = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyDBNodesVolumes) Validate() error {
	return dara.Validate(s)
}

type DescribeAIDBClusterAttributeResponseBodyEndpointList struct {
	NetInfoItems []*DescribeAIDBClusterAttributeResponseBodyEndpointListNetInfoItems `json:"NetInfoItems,omitempty" xml:"NetInfoItems,omitempty" type:"Repeated"`
}

func (s DescribeAIDBClusterAttributeResponseBodyEndpointList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterAttributeResponseBodyEndpointList) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterAttributeResponseBodyEndpointList) GetNetInfoItems() []*DescribeAIDBClusterAttributeResponseBodyEndpointListNetInfoItems {
	return s.NetInfoItems
}

func (s *DescribeAIDBClusterAttributeResponseBodyEndpointList) SetNetInfoItems(v []*DescribeAIDBClusterAttributeResponseBodyEndpointListNetInfoItems) *DescribeAIDBClusterAttributeResponseBodyEndpointList {
	s.NetInfoItems = v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyEndpointList) Validate() error {
	return dara.Validate(s)
}

type DescribeAIDBClusterAttributeResponseBodyEndpointListNetInfoItems struct {
	// example:
	//
	// pc-**********.rwlb.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// Public
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeAIDBClusterAttributeResponseBodyEndpointListNetInfoItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterAttributeResponseBodyEndpointListNetInfoItems) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterAttributeResponseBodyEndpointListNetInfoItems) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeAIDBClusterAttributeResponseBodyEndpointListNetInfoItems) GetNetType() *string {
	return s.NetType
}

func (s *DescribeAIDBClusterAttributeResponseBodyEndpointListNetInfoItems) GetPort() *string {
	return s.Port
}

func (s *DescribeAIDBClusterAttributeResponseBodyEndpointListNetInfoItems) SetConnectionString(v string) *DescribeAIDBClusterAttributeResponseBodyEndpointListNetInfoItems {
	s.ConnectionString = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyEndpointListNetInfoItems) SetNetType(v string) *DescribeAIDBClusterAttributeResponseBodyEndpointListNetInfoItems {
	s.NetType = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyEndpointListNetInfoItems) SetPort(v string) *DescribeAIDBClusterAttributeResponseBodyEndpointListNetInfoItems {
	s.Port = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyEndpointListNetInfoItems) Validate() error {
	return dara.Validate(s)
}

type DescribeAIDBClusterAttributeResponseBodyVolumes struct {
	// example:
	//
	// /var/run/secrets/kubernetes.io/serviceaccount
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// example:
	//
	// jueming
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 8192
	SizeGB *string `json:"SizeGB,omitempty" xml:"SizeGB,omitempty"`
	// example:
	//
	// PL1
	StorageCategory *string `json:"StorageCategory,omitempty" xml:"StorageCategory,omitempty"`
	// example:
	//
	// PolarFs
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeAIDBClusterAttributeResponseBodyVolumes) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterAttributeResponseBodyVolumes) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterAttributeResponseBodyVolumes) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeAIDBClusterAttributeResponseBodyVolumes) GetName() *string {
	return s.Name
}

func (s *DescribeAIDBClusterAttributeResponseBodyVolumes) GetSizeGB() *string {
	return s.SizeGB
}

func (s *DescribeAIDBClusterAttributeResponseBodyVolumes) GetStorageCategory() *string {
	return s.StorageCategory
}

func (s *DescribeAIDBClusterAttributeResponseBodyVolumes) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeAIDBClusterAttributeResponseBodyVolumes) SetMountPath(v string) *DescribeAIDBClusterAttributeResponseBodyVolumes {
	s.MountPath = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyVolumes) SetName(v string) *DescribeAIDBClusterAttributeResponseBodyVolumes {
	s.Name = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyVolumes) SetSizeGB(v string) *DescribeAIDBClusterAttributeResponseBodyVolumes {
	s.SizeGB = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyVolumes) SetStorageCategory(v string) *DescribeAIDBClusterAttributeResponseBodyVolumes {
	s.StorageCategory = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyVolumes) SetStorageType(v string) *DescribeAIDBClusterAttributeResponseBodyVolumes {
	s.StorageType = &v
	return s
}

func (s *DescribeAIDBClusterAttributeResponseBodyVolumes) Validate() error {
	return dara.Validate(s)
}
