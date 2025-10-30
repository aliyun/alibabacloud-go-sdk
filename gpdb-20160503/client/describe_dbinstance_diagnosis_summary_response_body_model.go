// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceDiagnosisSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeDBInstanceDiagnosisSummaryResponseBodyItems) *DescribeDBInstanceDiagnosisSummaryResponseBody
	GetItems() []*DescribeDBInstanceDiagnosisSummaryResponseBodyItems
	SetMasterStatusInfo(v *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) *DescribeDBInstanceDiagnosisSummaryResponseBody
	GetMasterStatusInfo() *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo
	SetPageNumber(v string) *DescribeDBInstanceDiagnosisSummaryResponseBody
	GetPageNumber() *string
	SetRequestId(v string) *DescribeDBInstanceDiagnosisSummaryResponseBody
	GetRequestId() *string
	SetSegmentStatusInfo(v *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) *DescribeDBInstanceDiagnosisSummaryResponseBody
	GetSegmentStatusInfo() *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo
	SetTotalCount(v string) *DescribeDBInstanceDiagnosisSummaryResponseBody
	GetTotalCount() *string
}

type DescribeDBInstanceDiagnosisSummaryResponseBody struct {
	// The group ID.
	Items []*DescribeDBInstanceDiagnosisSummaryResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The state information about the coordinator node.
	MasterStatusInfo *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo `json:"MasterStatusInfo,omitempty" xml:"MasterStatusInfo,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 070534EC-78D5-5519-83CC-E7B1A9213483
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state information about compute nodes.
	SegmentStatusInfo *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo `json:"SegmentStatusInfo,omitempty" xml:"SegmentStatusInfo,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 6
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) GetItems() []*DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) GetMasterStatusInfo() *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	return s.MasterStatusInfo
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) GetSegmentStatusInfo() *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	return s.SegmentStatusInfo
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetItems(v []*DescribeDBInstanceDiagnosisSummaryResponseBodyItems) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetMasterStatusInfo(v *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.MasterStatusInfo = v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetPageNumber(v string) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetRequestId(v string) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetSegmentStatusInfo(v *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.SegmentStatusInfo = v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) SetTotalCount(v string) *DescribeDBInstanceDiagnosisSummaryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MasterStatusInfo != nil {
		if err := s.MasterStatusInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SegmentStatusInfo != nil {
		if err := s.SegmentStatusInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceDiagnosisSummaryResponseBodyItems struct {
	// The name of the node.
	//
	// example:
	//
	// gp-t4np568qe9710****-master-100984919
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The IP address of the node.
	//
	// example:
	//
	// 192.168.XX.XX
	NodeAddress *string `json:"NodeAddress,omitempty" xml:"NodeAddress,omitempty"`
	// The node group ID.
	//
	// example:
	//
	// -1
	NodeCID *string `json:"NodeCID,omitempty" xml:"NodeCID,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 1
	NodeID *string `json:"NodeID,omitempty" xml:"NodeID,omitempty"`
	// The name of the host where the node resides.
	//
	// example:
	//
	// ap-southeast-1.i-t4n4c4ryr0yr441d****
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The port number of the node.
	//
	// example:
	//
	// 3000
	NodePort *string `json:"NodePort,omitempty" xml:"NodePort,omitempty"`
	// The initial role of the node. Valid values:
	//
	// 	- **primary**: primary node.
	//
	// 	- **mirror**: secondary node.
	//
	// If the value of this parameter is the same as that of **NodeRole**, no primary/secondary switchover occurs. If the value of this parameter is different from that of **NodeRole**, a primary/secondary switchover occurs.
	//
	// example:
	//
	// primary
	NodePreferredRole *string `json:"NodePreferredRole,omitempty" xml:"NodePreferredRole,omitempty"`
	// The data synchronization state of the node. Valid values:
	//
	// 	- **Synced**: The node data is synchronized.
	//
	// 	- **Not Syncing**: The node data is not synchronized.
	//
	// 	- **No sync required**: Data synchronization is not required. This value may be returned only for the coordinator node.
	//
	// example:
	//
	// Synced
	NodeReplicationMode *string `json:"NodeReplicationMode,omitempty" xml:"NodeReplicationMode,omitempty"`
	// The current role of the node. Valid values:
	//
	// 	- **primary**: primary node.
	//
	// 	- **mirror**: secondary node.
	//
	// example:
	//
	// primary
	NodeRole *string `json:"NodeRole,omitempty" xml:"NodeRole,omitempty"`
	// The running state of the node. Valid values:
	//
	// 	- **UP**: The node is running.
	//
	// 	- **DOWN**: The node is faulty.
	//
	// example:
	//
	// UP
	NodeStatus *string `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	// The type of the node. Valid values:
	//
	// 	- **master**: primary coordinator node.
	//
	// 	- **slave**: standby coordinator node.
	//
	// 	- **segment**: compute node.
	//
	// example:
	//
	// master
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) GetHostname() *string {
	return s.Hostname
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) GetNodeAddress() *string {
	return s.NodeAddress
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) GetNodeCID() *string {
	return s.NodeCID
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) GetNodeID() *string {
	return s.NodeID
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) GetNodePort() *string {
	return s.NodePort
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) GetNodePreferredRole() *string {
	return s.NodePreferredRole
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) GetNodeReplicationMode() *string {
	return s.NodeReplicationMode
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) GetNodeRole() *string {
	return s.NodeRole
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) GetNodeStatus() *string {
	return s.NodeStatus
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetHostname(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.Hostname = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeAddress(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeAddress = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeCID(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeCID = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeID(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeID = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeName(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeName = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodePort(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodePort = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodePreferredRole(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodePreferredRole = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeReplicationMode(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeReplicationMode = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeRole(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeRole = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeStatus(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeStatus = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) SetNodeType(v string) *DescribeDBInstanceDiagnosisSummaryResponseBodyItems {
	s.NodeType = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo struct {
	// The number of abnormal nodes.
	//
	// example:
	//
	// 0
	ExceptionNodeNum *int32 `json:"ExceptionNodeNum,omitempty" xml:"ExceptionNodeNum,omitempty"`
	// The number of normal nodes.
	//
	// example:
	//
	// 2
	NormalNodeNum *int32 `json:"NormalNodeNum,omitempty" xml:"NormalNodeNum,omitempty"`
	// The number of nodes whose roles are reversed.
	//
	// example:
	//
	// 0
	NotPreferredNodeNum *int32 `json:"NotPreferredNodeNum,omitempty" xml:"NotPreferredNodeNum,omitempty"`
	// The number of unsynchronized nodes.
	//
	// example:
	//
	// 0
	NotSyncingNodeNum *int32 `json:"NotSyncingNodeNum,omitempty" xml:"NotSyncingNodeNum,omitempty"`
	// The number of nodes whose roles are normal.
	//
	// example:
	//
	// 2
	PreferredNodeNum *int32 `json:"PreferredNodeNum,omitempty" xml:"PreferredNodeNum,omitempty"`
	// The number of synchronized nodes.
	//
	// example:
	//
	// 1
	SyncedNodeNum *int32 `json:"SyncedNodeNum,omitempty" xml:"SyncedNodeNum,omitempty"`
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) GetExceptionNodeNum() *int32 {
	return s.ExceptionNodeNum
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) GetNormalNodeNum() *int32 {
	return s.NormalNodeNum
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) GetNotPreferredNodeNum() *int32 {
	return s.NotPreferredNodeNum
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) GetNotSyncingNodeNum() *int32 {
	return s.NotSyncingNodeNum
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) GetPreferredNodeNum() *int32 {
	return s.PreferredNodeNum
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) GetSyncedNodeNum() *int32 {
	return s.SyncedNodeNum
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetExceptionNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.ExceptionNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetNormalNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.NormalNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetNotPreferredNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.NotPreferredNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetNotSyncingNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.NotSyncingNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetPreferredNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.PreferredNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) SetSyncedNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo {
	s.SyncedNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodyMasterStatusInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo struct {
	// The number of abnormal nodes.
	//
	// example:
	//
	// 0
	ExceptionNodeNum *int32 `json:"ExceptionNodeNum,omitempty" xml:"ExceptionNodeNum,omitempty"`
	// The number of normal nodes.
	//
	// example:
	//
	// 4
	NormalNodeNum *int32 `json:"NormalNodeNum,omitempty" xml:"NormalNodeNum,omitempty"`
	// The number of nodes whose roles are reversed.
	//
	// example:
	//
	// 0
	NotPreferredNodeNum *int32 `json:"NotPreferredNodeNum,omitempty" xml:"NotPreferredNodeNum,omitempty"`
	// The number of unsynchronized nodes.
	//
	// example:
	//
	// 4
	NotSyncingNodeNum *int32 `json:"NotSyncingNodeNum,omitempty" xml:"NotSyncingNodeNum,omitempty"`
	// The number of nodes whose roles are normal.
	//
	// example:
	//
	// 4
	PreferredNodeNum *int32 `json:"PreferredNodeNum,omitempty" xml:"PreferredNodeNum,omitempty"`
	// The number of synchronized nodes.
	//
	// example:
	//
	// 0
	SyncedNodeNum *int32 `json:"SyncedNodeNum,omitempty" xml:"SyncedNodeNum,omitempty"`
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) GetExceptionNodeNum() *int32 {
	return s.ExceptionNodeNum
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) GetNormalNodeNum() *int32 {
	return s.NormalNodeNum
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) GetNotPreferredNodeNum() *int32 {
	return s.NotPreferredNodeNum
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) GetNotSyncingNodeNum() *int32 {
	return s.NotSyncingNodeNum
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) GetPreferredNodeNum() *int32 {
	return s.PreferredNodeNum
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) GetSyncedNodeNum() *int32 {
	return s.SyncedNodeNum
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetExceptionNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.ExceptionNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetNormalNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.NormalNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetNotPreferredNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.NotPreferredNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetNotSyncingNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.NotSyncingNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetPreferredNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.PreferredNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) SetSyncedNodeNum(v int32) *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo {
	s.SyncedNodeNum = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryResponseBodySegmentStatusInfo) Validate() error {
	return dara.Validate(s)
}
