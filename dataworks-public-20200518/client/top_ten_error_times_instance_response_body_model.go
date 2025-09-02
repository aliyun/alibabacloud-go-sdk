// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTopTenErrorTimesInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceErrorRank(v *TopTenErrorTimesInstanceResponseBodyInstanceErrorRank) *TopTenErrorTimesInstanceResponseBody
	GetInstanceErrorRank() *TopTenErrorTimesInstanceResponseBodyInstanceErrorRank
	SetRequestId(v string) *TopTenErrorTimesInstanceResponseBody
	GetRequestId() *string
}

type TopTenErrorTimesInstanceResponseBody struct {
	// The ranking data of nodes on which errors occurred.
	InstanceErrorRank *TopTenErrorTimesInstanceResponseBodyInstanceErrorRank `json:"InstanceErrorRank,omitempty" xml:"InstanceErrorRank,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 952795279527****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TopTenErrorTimesInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TopTenErrorTimesInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *TopTenErrorTimesInstanceResponseBody) GetInstanceErrorRank() *TopTenErrorTimesInstanceResponseBodyInstanceErrorRank {
	return s.InstanceErrorRank
}

func (s *TopTenErrorTimesInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TopTenErrorTimesInstanceResponseBody) SetInstanceErrorRank(v *TopTenErrorTimesInstanceResponseBodyInstanceErrorRank) *TopTenErrorTimesInstanceResponseBody {
	s.InstanceErrorRank = v
	return s
}

func (s *TopTenErrorTimesInstanceResponseBody) SetRequestId(v string) *TopTenErrorTimesInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *TopTenErrorTimesInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type TopTenErrorTimesInstanceResponseBodyInstanceErrorRank struct {
	// The ranking data of nodes on which errors occurred within the last month.
	ErrorRank []*TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank `json:"ErrorRank,omitempty" xml:"ErrorRank,omitempty" type:"Repeated"`
	// The timestamp at which the rankings were updated.
	//
	// example:
	//
	// 1600963200000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s TopTenErrorTimesInstanceResponseBodyInstanceErrorRank) String() string {
	return dara.Prettify(s)
}

func (s TopTenErrorTimesInstanceResponseBodyInstanceErrorRank) GoString() string {
	return s.String()
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRank) GetErrorRank() []*TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank {
	return s.ErrorRank
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRank) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRank) SetErrorRank(v []*TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank) *TopTenErrorTimesInstanceResponseBodyInstanceErrorRank {
	s.ErrorRank = v
	return s
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRank) SetUpdateTime(v int64) *TopTenErrorTimesInstanceResponseBodyInstanceErrorRank {
	s.UpdateTime = &v
	return s
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRank) Validate() error {
	return dara.Validate(s)
}

type TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank struct {
	// The number of errors that occurred on the node.
	//
	// example:
	//
	// 5
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
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
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 9527
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank) String() string {
	return dara.Prettify(s)
}

func (s TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank) GoString() string {
	return s.String()
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank) GetCount() *int32 {
	return s.Count
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank) GetNodeId() *int64 {
	return s.NodeId
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank) GetNodeName() *string {
	return s.NodeName
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank) GetOwner() *string {
	return s.Owner
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank) GetProgramType() *int32 {
	return s.ProgramType
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank) SetCount(v int32) *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank {
	s.Count = &v
	return s
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank) SetNodeId(v int64) *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank {
	s.NodeId = &v
	return s
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank) SetNodeName(v string) *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank {
	s.NodeName = &v
	return s
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank) SetOwner(v string) *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank {
	s.Owner = &v
	return s
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank) SetProgramType(v int32) *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank {
	s.ProgramType = &v
	return s
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank) SetProjectId(v int64) *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank {
	s.ProjectId = &v
	return s
}

func (s *TopTenErrorTimesInstanceResponseBodyInstanceErrorRankErrorRank) Validate() error {
	return dara.Validate(s)
}
