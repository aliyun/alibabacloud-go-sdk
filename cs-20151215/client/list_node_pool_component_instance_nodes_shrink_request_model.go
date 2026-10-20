// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodePoolComponentInstanceNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRevision(v string) *ListNodePoolComponentInstanceNodesShrinkRequest
	GetConfigRevision() *string
	SetMaxResults(v int32) *ListNodePoolComponentInstanceNodesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodePoolComponentInstanceNodesShrinkRequest
	GetNextToken() *string
	SetNodeNamesShrink(v string) *ListNodePoolComponentInstanceNodesShrinkRequest
	GetNodeNamesShrink() *string
	SetVersion(v string) *ListNodePoolComponentInstanceNodesShrinkRequest
	GetVersion() *string
}

type ListNodePoolComponentInstanceNodesShrinkRequest struct {
	// Filters nodes by the specified configuration version number of the node component.
	//
	// example:
	//
	// 1
	ConfigRevision *string `json:"config_revision,omitempty" xml:"config_revision,omitempty"`
	// The maximum number of records to return in a single response for pagination. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"max_results,omitempty" xml:"max_results,omitempty"`
	// The pagination cursor. Do not specify this parameter for the first request. For subsequent requests, specify the token returned in the previous response.
	//
	// example:
	//
	// 5c0a1c0f91c14c6****
	NextToken *string `json:"next_token,omitempty" xml:"next_token,omitempty"`
	// Specifies the list of node names to query.
	//
	// example:
	//
	// ["cn-hangzhou.10.91.xx.xx"]
	NodeNamesShrink *string `json:"node_names,omitempty" xml:"node_names,omitempty"`
	// Filters nodes by the specified node component version.
	//
	// example:
	//
	// 1.28.9-aliyun.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListNodePoolComponentInstanceNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodePoolComponentInstanceNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListNodePoolComponentInstanceNodesShrinkRequest) GetConfigRevision() *string {
	return s.ConfigRevision
}

func (s *ListNodePoolComponentInstanceNodesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodePoolComponentInstanceNodesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodePoolComponentInstanceNodesShrinkRequest) GetNodeNamesShrink() *string {
	return s.NodeNamesShrink
}

func (s *ListNodePoolComponentInstanceNodesShrinkRequest) GetVersion() *string {
	return s.Version
}

func (s *ListNodePoolComponentInstanceNodesShrinkRequest) SetConfigRevision(v string) *ListNodePoolComponentInstanceNodesShrinkRequest {
	s.ConfigRevision = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesShrinkRequest) SetMaxResults(v int32) *ListNodePoolComponentInstanceNodesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesShrinkRequest) SetNextToken(v string) *ListNodePoolComponentInstanceNodesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesShrinkRequest) SetNodeNamesShrink(v string) *ListNodePoolComponentInstanceNodesShrinkRequest {
	s.NodeNamesShrink = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesShrinkRequest) SetVersion(v string) *ListNodePoolComponentInstanceNodesShrinkRequest {
	s.Version = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
