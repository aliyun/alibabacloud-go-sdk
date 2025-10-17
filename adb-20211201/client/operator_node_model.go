// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperatorNode interface {
	dara.Model
	String() string
	GoString() string
	SetChildren(v []*OperatorNode) *OperatorNode
	GetChildren() []*OperatorNode
	SetId(v int32) *OperatorNode
	GetId() *int32
	SetLevelWidth(v int32) *OperatorNode
	GetLevelWidth() *int32
	SetNodeDepth(v int32) *OperatorNode
	GetNodeDepth() *int32
	SetNodeName(v string) *OperatorNode
	GetNodeName() *string
	SetNodeWidth(v int32) *OperatorNode
	GetNodeWidth() *int32
	SetParentId(v int32) *OperatorNode
	GetParentId() *int32
	SetStats(v *OperatorNodeStats) *OperatorNode
	GetStats() *OperatorNodeStats
}

type OperatorNode struct {
	Children   []*OperatorNode    `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	Id         *int32             `json:"id,omitempty" xml:"id,omitempty"`
	LevelWidth *int32             `json:"levelWidth,omitempty" xml:"levelWidth,omitempty"`
	NodeDepth  *int32             `json:"nodeDepth,omitempty" xml:"nodeDepth,omitempty"`
	NodeName   *string            `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	NodeWidth  *int32             `json:"nodeWidth,omitempty" xml:"nodeWidth,omitempty"`
	ParentId   *int32             `json:"parentId,omitempty" xml:"parentId,omitempty"`
	Stats      *OperatorNodeStats `json:"stats,omitempty" xml:"stats,omitempty" type:"Struct"`
}

func (s OperatorNode) String() string {
	return dara.Prettify(s)
}

func (s OperatorNode) GoString() string {
	return s.String()
}

func (s *OperatorNode) GetChildren() []*OperatorNode {
	return s.Children
}

func (s *OperatorNode) GetId() *int32 {
	return s.Id
}

func (s *OperatorNode) GetLevelWidth() *int32 {
	return s.LevelWidth
}

func (s *OperatorNode) GetNodeDepth() *int32 {
	return s.NodeDepth
}

func (s *OperatorNode) GetNodeName() *string {
	return s.NodeName
}

func (s *OperatorNode) GetNodeWidth() *int32 {
	return s.NodeWidth
}

func (s *OperatorNode) GetParentId() *int32 {
	return s.ParentId
}

func (s *OperatorNode) GetStats() *OperatorNodeStats {
	return s.Stats
}

func (s *OperatorNode) SetChildren(v []*OperatorNode) *OperatorNode {
	s.Children = v
	return s
}

func (s *OperatorNode) SetId(v int32) *OperatorNode {
	s.Id = &v
	return s
}

func (s *OperatorNode) SetLevelWidth(v int32) *OperatorNode {
	s.LevelWidth = &v
	return s
}

func (s *OperatorNode) SetNodeDepth(v int32) *OperatorNode {
	s.NodeDepth = &v
	return s
}

func (s *OperatorNode) SetNodeName(v string) *OperatorNode {
	s.NodeName = &v
	return s
}

func (s *OperatorNode) SetNodeWidth(v int32) *OperatorNode {
	s.NodeWidth = &v
	return s
}

func (s *OperatorNode) SetParentId(v int32) *OperatorNode {
	s.ParentId = &v
	return s
}

func (s *OperatorNode) SetStats(v *OperatorNodeStats) *OperatorNode {
	s.Stats = v
	return s
}

func (s *OperatorNode) Validate() error {
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Stats != nil {
		if err := s.Stats.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OperatorNodeStats struct {
	Bytes      *int64  `json:"bytes,omitempty" xml:"bytes,omitempty"`
	OutputRows *int64  `json:"outputRows,omitempty" xml:"outputRows,omitempty"`
	Parameters *string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	PeakMemory *int64  `json:"peakMemory,omitempty" xml:"peakMemory,omitempty"`
	TimeCost   *int64  `json:"timeCost,omitempty" xml:"timeCost,omitempty"`
}

func (s OperatorNodeStats) String() string {
	return dara.Prettify(s)
}

func (s OperatorNodeStats) GoString() string {
	return s.String()
}

func (s *OperatorNodeStats) GetBytes() *int64 {
	return s.Bytes
}

func (s *OperatorNodeStats) GetOutputRows() *int64 {
	return s.OutputRows
}

func (s *OperatorNodeStats) GetParameters() *string {
	return s.Parameters
}

func (s *OperatorNodeStats) GetPeakMemory() *int64 {
	return s.PeakMemory
}

func (s *OperatorNodeStats) GetTimeCost() *int64 {
	return s.TimeCost
}

func (s *OperatorNodeStats) SetBytes(v int64) *OperatorNodeStats {
	s.Bytes = &v
	return s
}

func (s *OperatorNodeStats) SetOutputRows(v int64) *OperatorNodeStats {
	s.OutputRows = &v
	return s
}

func (s *OperatorNodeStats) SetParameters(v string) *OperatorNodeStats {
	s.Parameters = &v
	return s
}

func (s *OperatorNodeStats) SetPeakMemory(v int64) *OperatorNodeStats {
	s.PeakMemory = &v
	return s
}

func (s *OperatorNodeStats) SetTimeCost(v int64) *OperatorNodeStats {
	s.TimeCost = &v
	return s
}

func (s *OperatorNodeStats) Validate() error {
	return dara.Validate(s)
}
