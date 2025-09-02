// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTopTenElapsedTimeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceConsumeTimeRank(v *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRank) *TopTenElapsedTimeInstanceResponseBody
	GetInstanceConsumeTimeRank() *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRank
	SetRequestId(v string) *TopTenElapsedTimeInstanceResponseBody
	GetRequestId() *string
}

type TopTenElapsedTimeInstanceResponseBody struct {
	// The ranking record of the running durations of the instances.
	InstanceConsumeTimeRank *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRank `json:"InstanceConsumeTimeRank,omitempty" xml:"InstanceConsumeTimeRank,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6347364dadsfadf****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TopTenElapsedTimeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TopTenElapsedTimeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *TopTenElapsedTimeInstanceResponseBody) GetInstanceConsumeTimeRank() *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRank {
	return s.InstanceConsumeTimeRank
}

func (s *TopTenElapsedTimeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TopTenElapsedTimeInstanceResponseBody) SetInstanceConsumeTimeRank(v *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRank) *TopTenElapsedTimeInstanceResponseBody {
	s.InstanceConsumeTimeRank = v
	return s
}

func (s *TopTenElapsedTimeInstanceResponseBody) SetRequestId(v string) *TopTenElapsedTimeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *TopTenElapsedTimeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRank struct {
	// The ranking data of the running durations of the instances.
	ConsumeTimeRank []*TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank `json:"ConsumeTimeRank,omitempty" xml:"ConsumeTimeRank,omitempty" type:"Repeated"`
	// The timestamp at which the ranking of the running durations of the instances was updated.
	//
	// example:
	//
	// 1600963200000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRank) String() string {
	return dara.Prettify(s)
}

func (s TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRank) GoString() string {
	return s.String()
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRank) GetConsumeTimeRank() []*TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank {
	return s.ConsumeTimeRank
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRank) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRank) SetConsumeTimeRank(v []*TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRank {
	s.ConsumeTimeRank = v
	return s
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRank) SetUpdateTime(v int64) *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRank {
	s.UpdateTime = &v
	return s
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRank) Validate() error {
	return dara.Validate(s)
}

type TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank struct {
	// The data timestamp of the instance.
	//
	// example:
	//
	// 1600963200000
	BusinessDate *int64 `json:"BusinessDate,omitempty" xml:"BusinessDate,omitempty"`
	// The run time length of the instance. Unit: seconds.
	//
	// example:
	//
	// 1000
	Consumed *int64 `json:"Consumed,omitempty" xml:"Consumed,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 95279527
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 9527
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// Node name
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The ID of the Alibaba Cloud account used by the node owner.
	//
	// example:
	//
	// 952795279527
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The type of the node. Valid values: 6 (Shell), 10 (ODPS SQL), 11 (ODPS MR), 23 (Data Integration), 24 (ODPS Script), 99 (zero load), 221 (PyODPS 2), 225 (ODPS Spark), 227 (EMR Hive), 228 (EMR Spark), 229 (EMR Spark SQL), 230 (EMR MR), 239 (OSS object inspection), 257 (EMR Shell), 258 (EMR Spark Shell), 259 (EMR Presto), 260 (EMR Impala), 900 (real-time synchronization), 1089 (cross-tenant collaboration), 1091 (Hologres development), 1093 (Hologres SQL), 1100 (assignment), and 1221 (PyODPS 3)
	//
	// example:
	//
	// 10
	ProgramType *int32 `json:"ProgramType,omitempty" xml:"ProgramType,omitempty"`
}

func (s TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) String() string {
	return dara.Prettify(s)
}

func (s TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) GoString() string {
	return s.String()
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) GetBusinessDate() *int64 {
	return s.BusinessDate
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) GetConsumed() *int64 {
	return s.Consumed
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) GetNodeId() *int64 {
	return s.NodeId
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) GetNodeName() *string {
	return s.NodeName
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) GetOwner() *string {
	return s.Owner
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) GetProgramType() *int32 {
	return s.ProgramType
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) SetBusinessDate(v int64) *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank {
	s.BusinessDate = &v
	return s
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) SetConsumed(v int64) *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank {
	s.Consumed = &v
	return s
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) SetInstanceId(v int64) *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank {
	s.InstanceId = &v
	return s
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) SetNodeId(v int64) *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank {
	s.NodeId = &v
	return s
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) SetNodeName(v string) *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank {
	s.NodeName = &v
	return s
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) SetOwner(v string) *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank {
	s.Owner = &v
	return s
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) SetProgramType(v int32) *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank {
	s.ProgramType = &v
	return s
}

func (s *TopTenElapsedTimeInstanceResponseBodyInstanceConsumeTimeRankConsumeTimeRank) Validate() error {
	return dara.Validate(s)
}
