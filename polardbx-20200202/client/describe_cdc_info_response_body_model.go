// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdcInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeCdcInfoResponseBodyData) *DescribeCdcInfoResponseBody
	GetData() *DescribeCdcInfoResponseBodyData
	SetHttpStatusCode(v int32) *DescribeCdcInfoResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeCdcInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCdcInfoResponseBody
	GetSuccess() *bool
}

type DescribeCdcInfoResponseBody struct {
	// The instance data.
	Data *DescribeCdcInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCdcInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdcInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdcInfoResponseBody) GetData() *DescribeCdcInfoResponseBodyData {
	return s.Data
}

func (s *DescribeCdcInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeCdcInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdcInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCdcInfoResponseBody) SetData(v *DescribeCdcInfoResponseBodyData) *DescribeCdcInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCdcInfoResponseBody) SetHttpStatusCode(v int32) *DescribeCdcInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeCdcInfoResponseBody) SetRequestId(v string) *DescribeCdcInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdcInfoResponseBody) SetSuccess(v bool) *DescribeCdcInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCdcInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdcInfoResponseBodyData struct {
	// The number of days that binlogs are retained.
	//
	// example:
	//
	// 15
	BinlogPersistTime *int32 `json:"BinlogPersistTime,omitempty" xml:"BinlogPersistTime,omitempty"`
	// The size of the binlog.
	//
	// This parameter is required.
	//
	// example:
	//
	// 524288000
	BinlogSize *int32 `json:"BinlogSize,omitempty" xml:"BinlogSize,omitempty"`
	// The latest CDC version.
	//
	// This parameter is required.
	//
	// example:
	//
	// polarx-cdc-kernel-***
	CdcNewVersion *string `json:"CdcNewVersion,omitempty" xml:"CdcNewVersion,omitempty"`
	// The checksum switch status.
	//
	// example:
	//
	// ON
	CheckSumSwitch *string `json:"CheckSumSwitch,omitempty" xml:"CheckSumSwitch,omitempty"`
	// Indicates whether bidirectional replication is supported.
	//
	// example:
	//
	// true
	EnableCyclicReplication *bool `json:"EnableCyclicReplication,omitempty" xml:"EnableCyclicReplication,omitempty"`
	// The list of instances.
	InstanceTopologyList []*DescribeCdcInfoResponseBodyDataInstanceTopologyList `json:"InstanceTopologyList,omitempty" xml:"InstanceTopologyList,omitempty" type:"Repeated"`
	// server id
	//
	// This parameter is required.
	//
	// example:
	//
	// 3014767486
	ServerId *int32 `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// Indicates whether multi-stream CDC is supported.
	//
	// example:
	//
	// true
	VersionSupportMultiCdc *bool `json:"VersionSupportMultiCdc,omitempty" xml:"VersionSupportMultiCdc,omitempty"`
}

func (s DescribeCdcInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdcInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCdcInfoResponseBodyData) GetBinlogPersistTime() *int32 {
	return s.BinlogPersistTime
}

func (s *DescribeCdcInfoResponseBodyData) GetBinlogSize() *int32 {
	return s.BinlogSize
}

func (s *DescribeCdcInfoResponseBodyData) GetCdcNewVersion() *string {
	return s.CdcNewVersion
}

func (s *DescribeCdcInfoResponseBodyData) GetCheckSumSwitch() *string {
	return s.CheckSumSwitch
}

func (s *DescribeCdcInfoResponseBodyData) GetEnableCyclicReplication() *bool {
	return s.EnableCyclicReplication
}

func (s *DescribeCdcInfoResponseBodyData) GetInstanceTopologyList() []*DescribeCdcInfoResponseBodyDataInstanceTopologyList {
	return s.InstanceTopologyList
}

func (s *DescribeCdcInfoResponseBodyData) GetServerId() *int32 {
	return s.ServerId
}

func (s *DescribeCdcInfoResponseBodyData) GetVersionSupportMultiCdc() *bool {
	return s.VersionSupportMultiCdc
}

func (s *DescribeCdcInfoResponseBodyData) SetBinlogPersistTime(v int32) *DescribeCdcInfoResponseBodyData {
	s.BinlogPersistTime = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyData) SetBinlogSize(v int32) *DescribeCdcInfoResponseBodyData {
	s.BinlogSize = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyData) SetCdcNewVersion(v string) *DescribeCdcInfoResponseBodyData {
	s.CdcNewVersion = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyData) SetCheckSumSwitch(v string) *DescribeCdcInfoResponseBodyData {
	s.CheckSumSwitch = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyData) SetEnableCyclicReplication(v bool) *DescribeCdcInfoResponseBodyData {
	s.EnableCyclicReplication = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyData) SetInstanceTopologyList(v []*DescribeCdcInfoResponseBodyDataInstanceTopologyList) *DescribeCdcInfoResponseBodyData {
	s.InstanceTopologyList = v
	return s
}

func (s *DescribeCdcInfoResponseBodyData) SetServerId(v int32) *DescribeCdcInfoResponseBodyData {
	s.ServerId = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyData) SetVersionSupportMultiCdc(v bool) *DescribeCdcInfoResponseBodyData {
	s.VersionSupportMultiCdc = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyData) Validate() error {
	if s.InstanceTopologyList != nil {
		for _, item := range s.InstanceTopologyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdcInfoResponseBodyDataInstanceTopologyList struct {
	// The cluster type.
	//
	// example:
	//
	// BINLOG_X
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The remarks on the instance type.
	//
	// example:
	//
	// ***
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the multi-stream group.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The split level.
	//
	// example:
	//
	// RECORD
	HashLevel *string `json:"HashLevel,omitempty" xml:"HashLevel,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// pxc-***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The array of nodes.
	PhysicalNodes []*DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes `json:"PhysicalNodes,omitempty" xml:"PhysicalNodes,omitempty" type:"Repeated"`
	// The number of splits.
	//
	// example:
	//
	// 2
	StreamNum *int32 `json:"StreamNum,omitempty" xml:"StreamNum,omitempty"`
}

func (s DescribeCdcInfoResponseBodyDataInstanceTopologyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdcInfoResponseBodyDataInstanceTopologyList) GoString() string {
	return s.String()
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyList) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyList) GetComment() *string {
	return s.Comment
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyList) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyList) GetHashLevel() *string {
	return s.HashLevel
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyList) GetPhysicalNodes() []*DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes {
	return s.PhysicalNodes
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyList) GetStreamNum() *int32 {
	return s.StreamNum
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyList) SetClusterType(v string) *DescribeCdcInfoResponseBodyDataInstanceTopologyList {
	s.ClusterType = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyList) SetComment(v string) *DescribeCdcInfoResponseBodyDataInstanceTopologyList {
	s.Comment = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyList) SetGroupName(v string) *DescribeCdcInfoResponseBodyDataInstanceTopologyList {
	s.GroupName = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyList) SetHashLevel(v string) *DescribeCdcInfoResponseBodyDataInstanceTopologyList {
	s.HashLevel = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyList) SetInstanceName(v string) *DescribeCdcInfoResponseBodyDataInstanceTopologyList {
	s.InstanceName = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyList) SetPhysicalNodes(v []*DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) *DescribeCdcInfoResponseBodyDataInstanceTopologyList {
	s.PhysicalNodes = v
	return s
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyList) SetStreamNum(v int32) *DescribeCdcInfoResponseBodyDataInstanceTopologyList {
	s.StreamNum = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyList) Validate() error {
	if s.PhysicalNodes != nil {
		for _, item := range s.PhysicalNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes struct {
	// The zone in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou-h
	AZone *string `json:"AZone,omitempty" xml:"AZone,omitempty"`
	// The disk size. Unit: MB.
	//
	// example:
	//
	// 204800
	Disk *int32 `json:"Disk,omitempty" xml:"Disk,omitempty"`
	// The instance specifications.
	//
	// example:
	//
	// polarx.x4.large.2e.cdc
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// The node ID.
	//
	// example:
	//
	// ***
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// pxc-c-***
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The instance status.
	//
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The instance version.
	//
	// example:
	//
	// polarx-cdc-kernel-***
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) GoString() string {
	return s.String()
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) GetAZone() *string {
	return s.AZone
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) GetDisk() *int32 {
	return s.Disk
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) GetNodeClass() *string {
	return s.NodeClass
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) GetStatus() *string {
	return s.Status
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) GetVersion() *string {
	return s.Version
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) SetAZone(v string) *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes {
	s.AZone = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) SetDisk(v int32) *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes {
	s.Disk = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) SetNodeClass(v string) *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes {
	s.NodeClass = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) SetNodeId(v string) *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes {
	s.NodeId = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) SetNodeName(v string) *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes {
	s.NodeName = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) SetStatus(v string) *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes {
	s.Status = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) SetVersion(v string) *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes {
	s.Version = &v
	return s
}

func (s *DescribeCdcInfoResponseBodyDataInstanceTopologyListPhysicalNodes) Validate() error {
	return dara.Validate(s)
}
