// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLineageInfo interface {
	dara.Model
	String() string
	GoString() string
	SetEdges(v *Edge) *LineageInfo
	GetEdges() *Edge
	SetJobInfos(v []*JobInfo) *LineageInfo
	GetJobInfos() []*JobInfo
	SetNodes(v []*Node) *LineageInfo
	GetNodes() []*Node
}

type LineageInfo struct {
	Edges    *Edge      `json:"edges,omitempty" xml:"edges,omitempty"`
	JobInfos []*JobInfo `json:"jobInfos,omitempty" xml:"jobInfos,omitempty" type:"Repeated"`
	Nodes    []*Node    `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
}

func (s LineageInfo) String() string {
	return dara.Prettify(s)
}

func (s LineageInfo) GoString() string {
	return s.String()
}

func (s *LineageInfo) GetEdges() *Edge {
	return s.Edges
}

func (s *LineageInfo) GetJobInfos() []*JobInfo {
	return s.JobInfos
}

func (s *LineageInfo) GetNodes() []*Node {
	return s.Nodes
}

func (s *LineageInfo) SetEdges(v *Edge) *LineageInfo {
	s.Edges = v
	return s
}

func (s *LineageInfo) SetJobInfos(v []*JobInfo) *LineageInfo {
	s.JobInfos = v
	return s
}

func (s *LineageInfo) SetNodes(v []*Node) *LineageInfo {
	s.Nodes = v
	return s
}

func (s *LineageInfo) Validate() error {
	if s.Edges != nil {
		if err := s.Edges.Validate(); err != nil {
			return err
		}
	}
	if s.JobInfos != nil {
		for _, item := range s.JobInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
