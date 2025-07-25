// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnschedulableNodeDetail interface {
	dara.Model
	String() string
	GoString() string
	SetNodes(v []*string) *UnschedulableNodeDetail
	GetNodes() []*string
	SetReason(v string) *UnschedulableNodeDetail
	GetReason() *string
}

type UnschedulableNodeDetail struct {
	Nodes  []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Reason *string   `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s UnschedulableNodeDetail) String() string {
	return dara.Prettify(s)
}

func (s UnschedulableNodeDetail) GoString() string {
	return s.String()
}

func (s *UnschedulableNodeDetail) GetNodes() []*string {
	return s.Nodes
}

func (s *UnschedulableNodeDetail) GetReason() *string {
	return s.Reason
}

func (s *UnschedulableNodeDetail) SetNodes(v []*string) *UnschedulableNodeDetail {
	s.Nodes = v
	return s
}

func (s *UnschedulableNodeDetail) SetReason(v string) *UnschedulableNodeDetail {
	s.Reason = &v
	return s
}

func (s *UnschedulableNodeDetail) Validate() error {
	return dara.Validate(s)
}
