// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceHAConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceHAConfigResponseBody
	GetDBInstanceId() *string
	SetHAMode(v string) *DescribeDBInstanceHAConfigResponseBody
	GetHAMode() *string
	SetHostInstanceInfos(v *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfos) *DescribeDBInstanceHAConfigResponseBody
	GetHostInstanceInfos() *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfos
	SetRequestId(v string) *DescribeDBInstanceHAConfigResponseBody
	GetRequestId() *string
	SetSyncMode(v string) *DescribeDBInstanceHAConfigResponseBody
	GetSyncMode() *string
}

type DescribeDBInstanceHAConfigResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The high availability mode of the instance. Valid values:
	//
	// 	- **RPO**: Data consistency is preferred. The instance ensures data reliability to minimize data losses. If you have high requirements on data consistency, select this mode.
	//
	// 	- **RTO**: Service availability is preferred. The instance restores the database service at the earliest opportunity to ensure service availability. If you have high requirements on instance availability, select this mode.
	//
	// > This parameter is returned only for instances that run MySQL.
	//
	// example:
	//
	// RPO
	HAMode            *string                                                  `json:"HAMode,omitempty" xml:"HAMode,omitempty"`
	HostInstanceInfos *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfos `json:"HostInstanceInfos,omitempty" xml:"HostInstanceInfos,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data replication mode of the instance. Valid values:
	//
	// 	- **Sync**: the synchronous mode
	//
	// 	- **Semi-sync**: the semi-synchronous replication mode
	//
	// 	- **Async**: the asynchronous mode
	//
	// > This parameter is returned only for instances that run MySQL.
	//
	// example:
	//
	// Sync
	SyncMode *string `json:"SyncMode,omitempty" xml:"SyncMode,omitempty"`
}

func (s DescribeDBInstanceHAConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceHAConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceHAConfigResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceHAConfigResponseBody) GetHAMode() *string {
	return s.HAMode
}

func (s *DescribeDBInstanceHAConfigResponseBody) GetHostInstanceInfos() *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfos {
	return s.HostInstanceInfos
}

func (s *DescribeDBInstanceHAConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceHAConfigResponseBody) GetSyncMode() *string {
	return s.SyncMode
}

func (s *DescribeDBInstanceHAConfigResponseBody) SetDBInstanceId(v string) *DescribeDBInstanceHAConfigResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceHAConfigResponseBody) SetHAMode(v string) *DescribeDBInstanceHAConfigResponseBody {
	s.HAMode = &v
	return s
}

func (s *DescribeDBInstanceHAConfigResponseBody) SetHostInstanceInfos(v *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfos) *DescribeDBInstanceHAConfigResponseBody {
	s.HostInstanceInfos = v
	return s
}

func (s *DescribeDBInstanceHAConfigResponseBody) SetRequestId(v string) *DescribeDBInstanceHAConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceHAConfigResponseBody) SetSyncMode(v string) *DescribeDBInstanceHAConfigResponseBody {
	s.SyncMode = &v
	return s
}

func (s *DescribeDBInstanceHAConfigResponseBody) Validate() error {
	if s.HostInstanceInfos != nil {
		if err := s.HostInstanceInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceHAConfigResponseBodyHostInstanceInfos struct {
	NodeInfo []*DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceHAConfigResponseBodyHostInstanceInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceHAConfigResponseBodyHostInstanceInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfos) GetNodeInfo() []*DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo {
	return s.NodeInfo
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfos) SetNodeInfo(v []*DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfos {
	s.NodeInfo = v
	return s
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfos) Validate() error {
	if s.NodeInfo != nil {
		for _, item := range s.NodeInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo struct {
	DataSyncTime *string `json:"DataSyncTime,omitempty" xml:"DataSyncTime,omitempty"`
	LogSyncTime  *string `json:"LogSyncTime,omitempty" xml:"LogSyncTime,omitempty"`
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SyncStatus   *string `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) GetDataSyncTime() *string {
	return s.DataSyncTime
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) GetLogSyncTime() *string {
	return s.LogSyncTime
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) GetSyncStatus() *string {
	return s.SyncStatus
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) SetDataSyncTime(v string) *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo {
	s.DataSyncTime = &v
	return s
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) SetLogSyncTime(v string) *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo {
	s.LogSyncTime = &v
	return s
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) SetNodeId(v string) *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo {
	s.NodeId = &v
	return s
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) SetNodeType(v string) *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo {
	s.NodeType = &v
	return s
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) SetRegionId(v string) *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) SetSyncStatus(v string) *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo {
	s.SyncStatus = &v
	return s
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) SetZoneId(v string) *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceHAConfigResponseBodyHostInstanceInfosNodeInfo) Validate() error {
	return dara.Validate(s)
}
