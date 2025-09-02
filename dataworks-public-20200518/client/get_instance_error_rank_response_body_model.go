// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceErrorRankResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceErrorRank(v *GetInstanceErrorRankResponseBodyInstanceErrorRank) *GetInstanceErrorRankResponseBody
	GetInstanceErrorRank() *GetInstanceErrorRankResponseBodyInstanceErrorRank
	SetRequestId(v string) *GetInstanceErrorRankResponseBody
	GetRequestId() *string
}

type GetInstanceErrorRankResponseBody struct {
	// The ranking data of nodes on which errors occurred.
	InstanceErrorRank *GetInstanceErrorRankResponseBodyInstanceErrorRank `json:"InstanceErrorRank,omitempty" xml:"InstanceErrorRank,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 952795279527****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceErrorRankResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceErrorRankResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceErrorRankResponseBody) GetInstanceErrorRank() *GetInstanceErrorRankResponseBodyInstanceErrorRank {
	return s.InstanceErrorRank
}

func (s *GetInstanceErrorRankResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceErrorRankResponseBody) SetInstanceErrorRank(v *GetInstanceErrorRankResponseBodyInstanceErrorRank) *GetInstanceErrorRankResponseBody {
	s.InstanceErrorRank = v
	return s
}

func (s *GetInstanceErrorRankResponseBody) SetRequestId(v string) *GetInstanceErrorRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceErrorRankResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceErrorRankResponseBodyInstanceErrorRank struct {
	// The ranking data of nodes on which errors occurred within the last month.
	ErrorRank []*GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank `json:"ErrorRank,omitempty" xml:"ErrorRank,omitempty" type:"Repeated"`
	// The timestamp at which the rankings were updated.
	//
	// example:
	//
	// 1600963200000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetInstanceErrorRankResponseBodyInstanceErrorRank) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceErrorRankResponseBodyInstanceErrorRank) GoString() string {
	return s.String()
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRank) GetErrorRank() []*GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank {
	return s.ErrorRank
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRank) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRank) SetErrorRank(v []*GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank) *GetInstanceErrorRankResponseBodyInstanceErrorRank {
	s.ErrorRank = v
	return s
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRank) SetUpdateTime(v int64) *GetInstanceErrorRankResponseBodyInstanceErrorRank {
	s.UpdateTime = &v
	return s
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRank) Validate() error {
	return dara.Validate(s)
}

type GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank struct {
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
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The ID of the Alibaba Cloud account used by the node owner.
	//
	// example:
	//
	// 952795279527
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The type of the node.
	//
	// example:
	//
	// 10
	PrgType *int32 `json:"PrgType,omitempty" xml:"PrgType,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 9527
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank) GoString() string {
	return s.String()
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank) GetCount() *int32 {
	return s.Count
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank) GetNodeName() *string {
	return s.NodeName
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank) GetOwner() *string {
	return s.Owner
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank) GetPrgType() *int32 {
	return s.PrgType
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank) SetCount(v int32) *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank {
	s.Count = &v
	return s
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank) SetNodeId(v int64) *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank {
	s.NodeId = &v
	return s
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank) SetNodeName(v string) *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank {
	s.NodeName = &v
	return s
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank) SetOwner(v string) *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank {
	s.Owner = &v
	return s
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank) SetPrgType(v int32) *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank {
	s.PrgType = &v
	return s
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank) SetProjectId(v int64) *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank {
	s.ProjectId = &v
	return s
}

func (s *GetInstanceErrorRankResponseBodyInstanceErrorRankErrorRank) Validate() error {
	return dara.Validate(s)
}
