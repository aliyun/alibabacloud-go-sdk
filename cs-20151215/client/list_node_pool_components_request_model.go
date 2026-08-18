// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodePoolComponentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *ListNodePoolComponentsRequest
	GetImageId() *string
	SetImageType(v string) *ListNodePoolComponentsRequest
	GetImageType() *string
	SetInstanceTypes(v []*string) *ListNodePoolComponentsRequest
	GetInstanceTypes() []*string
	SetMaxResults(v int32) *ListNodePoolComponentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodePoolComponentsRequest
	GetNextToken() *string
	SetNodepoolId(v string) *ListNodePoolComponentsRequest
	GetNodepoolId() *string
	SetNodepoolType(v string) *ListNodePoolComponentsRequest
	GetNodepoolType() *string
}

type ListNodePoolComponentsRequest struct {
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
	InstanceTypes []*string `json:"instance_types,omitempty" xml:"instance_types,omitempty" type:"Repeated"`
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

func (s ListNodePoolComponentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodePoolComponentsRequest) GoString() string {
	return s.String()
}

func (s *ListNodePoolComponentsRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ListNodePoolComponentsRequest) GetImageType() *string {
	return s.ImageType
}

func (s *ListNodePoolComponentsRequest) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *ListNodePoolComponentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodePoolComponentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodePoolComponentsRequest) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *ListNodePoolComponentsRequest) GetNodepoolType() *string {
	return s.NodepoolType
}

func (s *ListNodePoolComponentsRequest) SetImageId(v string) *ListNodePoolComponentsRequest {
	s.ImageId = &v
	return s
}

func (s *ListNodePoolComponentsRequest) SetImageType(v string) *ListNodePoolComponentsRequest {
	s.ImageType = &v
	return s
}

func (s *ListNodePoolComponentsRequest) SetInstanceTypes(v []*string) *ListNodePoolComponentsRequest {
	s.InstanceTypes = v
	return s
}

func (s *ListNodePoolComponentsRequest) SetMaxResults(v int32) *ListNodePoolComponentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodePoolComponentsRequest) SetNextToken(v string) *ListNodePoolComponentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodePoolComponentsRequest) SetNodepoolId(v string) *ListNodePoolComponentsRequest {
	s.NodepoolId = &v
	return s
}

func (s *ListNodePoolComponentsRequest) SetNodepoolType(v string) *ListNodePoolComponentsRequest {
	s.NodepoolType = &v
	return s
}

func (s *ListNodePoolComponentsRequest) Validate() error {
	return dara.Validate(s)
}
