// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocatedHyperNodeDetail interface {
	dara.Model
	String() string
	GoString() string
	SetAllocatedNodeNum(v int64) *AllocatedHyperNodeDetail
	GetAllocatedNodeNum() *int64
	SetEmptyNodeNum(v int64) *AllocatedHyperNodeDetail
	GetEmptyNodeNum() *int64
	SetHyperNodeName(v string) *AllocatedHyperNodeDetail
	GetHyperNodeName() *string
	SetTotalNodeNum(v int64) *AllocatedHyperNodeDetail
	GetTotalNodeNum() *int64
}

type AllocatedHyperNodeDetail struct {
	// The number of allocated nodes in the hyper node.
	AllocatedNodeNum *int64 `json:"AllocatedNodeNum,omitempty" xml:"AllocatedNodeNum,omitempty"`
	// The number of idle nodes in the hyper node.
	EmptyNodeNum *int64 `json:"EmptyNodeNum,omitempty" xml:"EmptyNodeNum,omitempty"`
	// The name of the hyper node.
	HyperNodeName *string `json:"HyperNodeName,omitempty" xml:"HyperNodeName,omitempty"`
	// The total number of nodes in the hyper node.
	TotalNodeNum *int64 `json:"TotalNodeNum,omitempty" xml:"TotalNodeNum,omitempty"`
}

func (s AllocatedHyperNodeDetail) String() string {
	return dara.Prettify(s)
}

func (s AllocatedHyperNodeDetail) GoString() string {
	return s.String()
}

func (s *AllocatedHyperNodeDetail) GetAllocatedNodeNum() *int64 {
	return s.AllocatedNodeNum
}

func (s *AllocatedHyperNodeDetail) GetEmptyNodeNum() *int64 {
	return s.EmptyNodeNum
}

func (s *AllocatedHyperNodeDetail) GetHyperNodeName() *string {
	return s.HyperNodeName
}

func (s *AllocatedHyperNodeDetail) GetTotalNodeNum() *int64 {
	return s.TotalNodeNum
}

func (s *AllocatedHyperNodeDetail) SetAllocatedNodeNum(v int64) *AllocatedHyperNodeDetail {
	s.AllocatedNodeNum = &v
	return s
}

func (s *AllocatedHyperNodeDetail) SetEmptyNodeNum(v int64) *AllocatedHyperNodeDetail {
	s.EmptyNodeNum = &v
	return s
}

func (s *AllocatedHyperNodeDetail) SetHyperNodeName(v string) *AllocatedHyperNodeDetail {
	s.HyperNodeName = &v
	return s
}

func (s *AllocatedHyperNodeDetail) SetTotalNodeNum(v int64) *AllocatedHyperNodeDetail {
	s.TotalNodeNum = &v
	return s
}

func (s *AllocatedHyperNodeDetail) Validate() error {
	return dara.Validate(s)
}
