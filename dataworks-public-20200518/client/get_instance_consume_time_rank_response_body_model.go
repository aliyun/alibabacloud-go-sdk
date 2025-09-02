// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceConsumeTimeRankResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceConsumeTimeRank(v *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRank) *GetInstanceConsumeTimeRankResponseBody
	GetInstanceConsumeTimeRank() *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRank
	SetRequestId(v string) *GetInstanceConsumeTimeRankResponseBody
	GetRequestId() *string
}

type GetInstanceConsumeTimeRankResponseBody struct {
	// The ranking record of the running durations of instances.
	InstanceConsumeTimeRank *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRank `json:"InstanceConsumeTimeRank,omitempty" xml:"InstanceConsumeTimeRank,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6347364dadsfadf****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceConsumeTimeRankResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceConsumeTimeRankResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceConsumeTimeRankResponseBody) GetInstanceConsumeTimeRank() *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRank {
	return s.InstanceConsumeTimeRank
}

func (s *GetInstanceConsumeTimeRankResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceConsumeTimeRankResponseBody) SetInstanceConsumeTimeRank(v *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRank) *GetInstanceConsumeTimeRankResponseBody {
	s.InstanceConsumeTimeRank = v
	return s
}

func (s *GetInstanceConsumeTimeRankResponseBody) SetRequestId(v string) *GetInstanceConsumeTimeRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceConsumeTimeRankResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRank struct {
	// The ranking data of the running durations of instances.
	ConsumeTimeRank []*GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank `json:"ConsumeTimeRank,omitempty" xml:"ConsumeTimeRank,omitempty" type:"Repeated"`
	// The timestamp when the ranking was updated.
	//
	// example:
	//
	// 1600963200000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRank) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRank) GoString() string {
	return s.String()
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRank) GetConsumeTimeRank() []*GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank {
	return s.ConsumeTimeRank
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRank) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRank) SetConsumeTimeRank(v []*GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRank {
	s.ConsumeTimeRank = v
	return s
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRank) SetUpdateTime(v int64) *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRank {
	s.UpdateTime = &v
	return s
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRank) Validate() error {
	return dara.Validate(s)
}

type GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank struct {
	// The data timestamp of the instance.
	//
	// example:
	//
	// 1600963200000
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The running duration of the instance. Unit: seconds.
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
}

func (s GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) GoString() string {
	return s.String()
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) GetConsumed() *int64 {
	return s.Consumed
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) GetNodeName() *string {
	return s.NodeName
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) GetOwner() *string {
	return s.Owner
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) GetPrgType() *int32 {
	return s.PrgType
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) SetBizdate(v int64) *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank {
	s.Bizdate = &v
	return s
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) SetConsumed(v int64) *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank {
	s.Consumed = &v
	return s
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) SetInstanceId(v int64) *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) SetNodeId(v int64) *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank {
	s.NodeId = &v
	return s
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) SetNodeName(v string) *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank {
	s.NodeName = &v
	return s
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) SetOwner(v string) *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank {
	s.Owner = &v
	return s
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) SetPrgType(v int32) *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank {
	s.PrgType = &v
	return s
}

func (s *GetInstanceConsumeTimeRankResponseBodyInstanceConsumeTimeRankConsumeTimeRank) Validate() error {
	return dara.Validate(s)
}
