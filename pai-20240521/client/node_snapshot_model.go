// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeSnapshot interface {
	dara.Model
	String() string
	GoString() string
	SetAncestorQuotaWorkloadNum(v int32) *NodeSnapshot
	GetAncestorQuotaWorkloadNum() *int32
	SetDescendantQuotaWorkloadNum(v int32) *NodeSnapshot
	GetDescendantQuotaWorkloadNum() *int32
	SetNodeName(v string) *NodeSnapshot
	GetNodeName() *string
	SetRequestCPU(v string) *NodeSnapshot
	GetRequestCPU() *string
	SetRequestGPU(v string) *NodeSnapshot
	GetRequestGPU() *string
	SetRequestMemory(v string) *NodeSnapshot
	GetRequestMemory() *string
	SetSelfQuotaWorkloadNum(v int32) *NodeSnapshot
	GetSelfQuotaWorkloadNum() *int32
	SetWorkloadNum(v int32) *NodeSnapshot
	GetWorkloadNum() *int32
	SetWorkloads(v []*NodeSnapshotWorkloads) *NodeSnapshot
	GetWorkloads() []*NodeSnapshotWorkloads
}

type NodeSnapshot struct {
	AncestorQuotaWorkloadNum   *int32                   `json:"AncestorQuotaWorkloadNum,omitempty" xml:"AncestorQuotaWorkloadNum,omitempty"`
	DescendantQuotaWorkloadNum *int32                   `json:"DescendantQuotaWorkloadNum,omitempty" xml:"DescendantQuotaWorkloadNum,omitempty"`
	NodeName                   *string                  `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	RequestCPU                 *string                  `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	RequestGPU                 *string                  `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	RequestMemory              *string                  `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	SelfQuotaWorkloadNum       *int32                   `json:"SelfQuotaWorkloadNum,omitempty" xml:"SelfQuotaWorkloadNum,omitempty"`
	WorkloadNum                *int32                   `json:"WorkloadNum,omitempty" xml:"WorkloadNum,omitempty"`
	Workloads                  []*NodeSnapshotWorkloads `json:"Workloads,omitempty" xml:"Workloads,omitempty" type:"Repeated"`
}

func (s NodeSnapshot) String() string {
	return dara.Prettify(s)
}

func (s NodeSnapshot) GoString() string {
	return s.String()
}

func (s *NodeSnapshot) GetAncestorQuotaWorkloadNum() *int32 {
	return s.AncestorQuotaWorkloadNum
}

func (s *NodeSnapshot) GetDescendantQuotaWorkloadNum() *int32 {
	return s.DescendantQuotaWorkloadNum
}

func (s *NodeSnapshot) GetNodeName() *string {
	return s.NodeName
}

func (s *NodeSnapshot) GetRequestCPU() *string {
	return s.RequestCPU
}

func (s *NodeSnapshot) GetRequestGPU() *string {
	return s.RequestGPU
}

func (s *NodeSnapshot) GetRequestMemory() *string {
	return s.RequestMemory
}

func (s *NodeSnapshot) GetSelfQuotaWorkloadNum() *int32 {
	return s.SelfQuotaWorkloadNum
}

func (s *NodeSnapshot) GetWorkloadNum() *int32 {
	return s.WorkloadNum
}

func (s *NodeSnapshot) GetWorkloads() []*NodeSnapshotWorkloads {
	return s.Workloads
}

func (s *NodeSnapshot) SetAncestorQuotaWorkloadNum(v int32) *NodeSnapshot {
	s.AncestorQuotaWorkloadNum = &v
	return s
}

func (s *NodeSnapshot) SetDescendantQuotaWorkloadNum(v int32) *NodeSnapshot {
	s.DescendantQuotaWorkloadNum = &v
	return s
}

func (s *NodeSnapshot) SetNodeName(v string) *NodeSnapshot {
	s.NodeName = &v
	return s
}

func (s *NodeSnapshot) SetRequestCPU(v string) *NodeSnapshot {
	s.RequestCPU = &v
	return s
}

func (s *NodeSnapshot) SetRequestGPU(v string) *NodeSnapshot {
	s.RequestGPU = &v
	return s
}

func (s *NodeSnapshot) SetRequestMemory(v string) *NodeSnapshot {
	s.RequestMemory = &v
	return s
}

func (s *NodeSnapshot) SetSelfQuotaWorkloadNum(v int32) *NodeSnapshot {
	s.SelfQuotaWorkloadNum = &v
	return s
}

func (s *NodeSnapshot) SetWorkloadNum(v int32) *NodeSnapshot {
	s.WorkloadNum = &v
	return s
}

func (s *NodeSnapshot) SetWorkloads(v []*NodeSnapshotWorkloads) *NodeSnapshot {
	s.Workloads = v
	return s
}

func (s *NodeSnapshot) Validate() error {
	if s.Workloads != nil {
		for _, item := range s.Workloads {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type NodeSnapshotWorkloads struct {
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	WorkloadId   *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
}

func (s NodeSnapshotWorkloads) String() string {
	return dara.Prettify(s)
}

func (s NodeSnapshotWorkloads) GoString() string {
	return s.String()
}

func (s *NodeSnapshotWorkloads) GetName() *string {
	return s.Name
}

func (s *NodeSnapshotWorkloads) GetWorkloadId() *string {
	return s.WorkloadId
}

func (s *NodeSnapshotWorkloads) GetWorkloadType() *string {
	return s.WorkloadType
}

func (s *NodeSnapshotWorkloads) SetName(v string) *NodeSnapshotWorkloads {
	s.Name = &v
	return s
}

func (s *NodeSnapshotWorkloads) SetWorkloadId(v string) *NodeSnapshotWorkloads {
	s.WorkloadId = &v
	return s
}

func (s *NodeSnapshotWorkloads) SetWorkloadType(v string) *NodeSnapshotWorkloads {
	s.WorkloadType = &v
	return s
}

func (s *NodeSnapshotWorkloads) Validate() error {
	return dara.Validate(s)
}
