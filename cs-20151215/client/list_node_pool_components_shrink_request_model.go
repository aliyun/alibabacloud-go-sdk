// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodePoolComponentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *ListNodePoolComponentsShrinkRequest
	GetImageId() *string
	SetImageType(v string) *ListNodePoolComponentsShrinkRequest
	GetImageType() *string
	SetInstanceTypesShrink(v string) *ListNodePoolComponentsShrinkRequest
	GetInstanceTypesShrink() *string
	SetMaxResults(v int32) *ListNodePoolComponentsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodePoolComponentsShrinkRequest
	GetNextToken() *string
	SetNodepoolId(v string) *ListNodePoolComponentsShrinkRequest
	GetNodepoolId() *string
	SetNodepoolType(v string) *ListNodePoolComponentsShrinkRequest
	GetNodepoolType() *string
}

type ListNodePoolComponentsShrinkRequest struct {
	// example:
	//
	// aliyun_3_x64_20G_container_optimized_alibase_20250629.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// example:
	//
	// AliyunLinux3
	ImageType *string `json:"image_type,omitempty" xml:"image_type,omitempty"`
	// example:
	//
	// ["ecs.c6.xlarge"]
	InstanceTypesShrink *string `json:"instance_types,omitempty" xml:"instance_types,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"max_results,omitempty" xml:"max_results,omitempty"`
	// example:
	//
	// 5c0a1c0f91c14c6****
	NextToken *string `json:"next_token,omitempty" xml:"next_token,omitempty"`
	// example:
	//
	// np1855b102ac434f5990d87b77a****
	NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	// example:
	//
	// ess
	NodepoolType *string `json:"nodepool_type,omitempty" xml:"nodepool_type,omitempty"`
}

func (s ListNodePoolComponentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodePoolComponentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListNodePoolComponentsShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ListNodePoolComponentsShrinkRequest) GetImageType() *string {
	return s.ImageType
}

func (s *ListNodePoolComponentsShrinkRequest) GetInstanceTypesShrink() *string {
	return s.InstanceTypesShrink
}

func (s *ListNodePoolComponentsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodePoolComponentsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodePoolComponentsShrinkRequest) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *ListNodePoolComponentsShrinkRequest) GetNodepoolType() *string {
	return s.NodepoolType
}

func (s *ListNodePoolComponentsShrinkRequest) SetImageId(v string) *ListNodePoolComponentsShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *ListNodePoolComponentsShrinkRequest) SetImageType(v string) *ListNodePoolComponentsShrinkRequest {
	s.ImageType = &v
	return s
}

func (s *ListNodePoolComponentsShrinkRequest) SetInstanceTypesShrink(v string) *ListNodePoolComponentsShrinkRequest {
	s.InstanceTypesShrink = &v
	return s
}

func (s *ListNodePoolComponentsShrinkRequest) SetMaxResults(v int32) *ListNodePoolComponentsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodePoolComponentsShrinkRequest) SetNextToken(v string) *ListNodePoolComponentsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodePoolComponentsShrinkRequest) SetNodepoolId(v string) *ListNodePoolComponentsShrinkRequest {
	s.NodepoolId = &v
	return s
}

func (s *ListNodePoolComponentsShrinkRequest) SetNodepoolType(v string) *ListNodePoolComponentsShrinkRequest {
	s.NodepoolType = &v
	return s
}

func (s *ListNodePoolComponentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
