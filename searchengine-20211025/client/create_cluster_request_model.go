// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoLoad(v bool) *CreateClusterRequest
	GetAutoLoad() *bool
	SetDataNode(v *CreateClusterRequestDataNode) *CreateClusterRequest
	GetDataNode() *CreateClusterRequestDataNode
	SetDescription(v string) *CreateClusterRequest
	GetDescription() *string
	SetName(v string) *CreateClusterRequest
	GetName() *string
	SetQueryNode(v *CreateClusterRequestQueryNode) *CreateClusterRequest
	GetQueryNode() *CreateClusterRequestQueryNode
}

type CreateClusterRequest struct {
	// Specifies whether to enable automatic connection.
	//
	// example:
	//
	// true
	AutoLoad *bool `json:"autoLoad,omitempty" xml:"autoLoad,omitempty"`
	// The details of the Searcher workers.
	DataNode *CreateClusterRequestDataNode `json:"dataNode,omitempty" xml:"dataNode,omitempty" type:"Struct"`
	// The description of the cluster.
	//
	// example:
	//
	// "ha-tets"
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// ha-cn-zvp2qr1sk01_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The details of the Query Result Searcher (QRS) workers.
	QueryNode *CreateClusterRequestQueryNode `json:"queryNode,omitempty" xml:"queryNode,omitempty" type:"Struct"`
}

func (s CreateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) GetAutoLoad() *bool {
	return s.AutoLoad
}

func (s *CreateClusterRequest) GetDataNode() *CreateClusterRequestDataNode {
	return s.DataNode
}

func (s *CreateClusterRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateClusterRequest) GetName() *string {
	return s.Name
}

func (s *CreateClusterRequest) GetQueryNode() *CreateClusterRequestQueryNode {
	return s.QueryNode
}

func (s *CreateClusterRequest) SetAutoLoad(v bool) *CreateClusterRequest {
	s.AutoLoad = &v
	return s
}

func (s *CreateClusterRequest) SetDataNode(v *CreateClusterRequestDataNode) *CreateClusterRequest {
	s.DataNode = v
	return s
}

func (s *CreateClusterRequest) SetDescription(v string) *CreateClusterRequest {
	s.Description = &v
	return s
}

func (s *CreateClusterRequest) SetName(v string) *CreateClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateClusterRequest) SetQueryNode(v *CreateClusterRequestQueryNode) *CreateClusterRequest {
	s.QueryNode = v
	return s
}

func (s *CreateClusterRequest) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestDataNode struct {
	// The number of Searcher workers.
	//
	// example:
	//
	// 2
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
	// The number of shards.
	//
	// example:
	//
	// 2
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s CreateClusterRequestDataNode) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestDataNode) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestDataNode) GetNumber() *int32 {
	return s.Number
}

func (s *CreateClusterRequestDataNode) GetPartition() *string {
	return s.Partition
}

func (s *CreateClusterRequestDataNode) SetNumber(v int32) *CreateClusterRequestDataNode {
	s.Number = &v
	return s
}

func (s *CreateClusterRequestDataNode) SetPartition(v string) *CreateClusterRequestDataNode {
	s.Partition = &v
	return s
}

func (s *CreateClusterRequestDataNode) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestQueryNode struct {
	// The number of QRS workers.
	//
	// example:
	//
	// 2
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
}

func (s CreateClusterRequestQueryNode) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestQueryNode) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestQueryNode) GetNumber() *int32 {
	return s.Number
}

func (s *CreateClusterRequestQueryNode) SetNumber(v int32) *CreateClusterRequestQueryNode {
	s.Number = &v
	return s
}

func (s *CreateClusterRequestQueryNode) Validate() error {
	return dara.Validate(s)
}
