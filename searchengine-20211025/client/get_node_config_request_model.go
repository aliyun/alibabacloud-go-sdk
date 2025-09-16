// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *GetNodeConfigRequest
	GetClusterName() *string
	SetName(v string) *GetNodeConfigRequest
	GetName() *string
	SetType(v string) *GetNodeConfigRequest
	GetType() *string
}

type GetNodeConfigRequest struct {
	// The name of the cluster
	//
	// example:
	//
	// vpc_sh_domain_2
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	// The node name.
	//
	// example:
	//
	// ha-cn-30174dhoz53_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The node type. Valid values:
	//
	// 	- qrs: Query Result Searcher (QRS) worker
	//
	// 	- search: Search worker
	//
	// 	- index: index
	//
	// 	- cluster: cluster
	//
	// example:
	//
	// index
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetNodeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeConfigRequest) GoString() string {
	return s.String()
}

func (s *GetNodeConfigRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetNodeConfigRequest) GetName() *string {
	return s.Name
}

func (s *GetNodeConfigRequest) GetType() *string {
	return s.Type
}

func (s *GetNodeConfigRequest) SetClusterName(v string) *GetNodeConfigRequest {
	s.ClusterName = &v
	return s
}

func (s *GetNodeConfigRequest) SetName(v string) *GetNodeConfigRequest {
	s.Name = &v
	return s
}

func (s *GetNodeConfigRequest) SetType(v string) *GetNodeConfigRequest {
	s.Type = &v
	return s
}

func (s *GetNodeConfigRequest) Validate() error {
	return dara.Validate(s)
}
