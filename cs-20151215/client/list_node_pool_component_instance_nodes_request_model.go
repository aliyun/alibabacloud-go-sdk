// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodePoolComponentInstanceNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRevision(v string) *ListNodePoolComponentInstanceNodesRequest
	GetConfigRevision() *string
	SetMaxResults(v int32) *ListNodePoolComponentInstanceNodesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodePoolComponentInstanceNodesRequest
	GetNextToken() *string
	SetNodeNames(v []*string) *ListNodePoolComponentInstanceNodesRequest
	GetNodeNames() []*string
	SetVersion(v string) *ListNodePoolComponentInstanceNodesRequest
	GetVersion() *string
}

type ListNodePoolComponentInstanceNodesRequest struct {
	// example:
	//
	// 1
	ConfigRevision *string `json:"config_revision,omitempty" xml:"config_revision,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"max_results,omitempty" xml:"max_results,omitempty"`
	// example:
	//
	// 5c0a1c0f91c14c6****
	NextToken *string `json:"next_token,omitempty" xml:"next_token,omitempty"`
	// example:
	//
	// ["cn-hangzhou.10.91.xx.xx"]
	NodeNames []*string `json:"node_names,omitempty" xml:"node_names,omitempty" type:"Repeated"`
	// example:
	//
	// 1.28.9-aliyun.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListNodePoolComponentInstanceNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodePoolComponentInstanceNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodePoolComponentInstanceNodesRequest) GetConfigRevision() *string {
	return s.ConfigRevision
}

func (s *ListNodePoolComponentInstanceNodesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodePoolComponentInstanceNodesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodePoolComponentInstanceNodesRequest) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *ListNodePoolComponentInstanceNodesRequest) GetVersion() *string {
	return s.Version
}

func (s *ListNodePoolComponentInstanceNodesRequest) SetConfigRevision(v string) *ListNodePoolComponentInstanceNodesRequest {
	s.ConfigRevision = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesRequest) SetMaxResults(v int32) *ListNodePoolComponentInstanceNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesRequest) SetNextToken(v string) *ListNodePoolComponentInstanceNodesRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesRequest) SetNodeNames(v []*string) *ListNodePoolComponentInstanceNodesRequest {
	s.NodeNames = v
	return s
}

func (s *ListNodePoolComponentInstanceNodesRequest) SetVersion(v string) *ListNodePoolComponentInstanceNodesRequest {
	s.Version = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesRequest) Validate() error {
	return dara.Validate(s)
}
