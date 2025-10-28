// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetK8sClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterType(v int32) *GetK8sClusterRequest
	GetClusterType() *int32
	SetCurrentPage(v int32) *GetK8sClusterRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *GetK8sClusterRequest
	GetPageSize() *int32
	SetRegionTag(v string) *GetK8sClusterRequest
	GetRegionTag() *string
	SetSubClusterType(v string) *GetK8sClusterRequest
	GetSubClusterType() *string
}

type GetK8sClusterRequest struct {
	// The type of the Kubernetes cluster. Valid values:
	//
	// 	- 5: ACK cluster
	//
	// 	- 7: self-managed Kubernetes cluster
	//
	// example:
	//
	// 5
	ClusterType *int32 `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page. Default value: 1000.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionTag *string `json:"RegionTag,omitempty" xml:"RegionTag,omitempty"`
	// The subtype of the cluster. Valid values:
	//
	// 	- Ask: Serverless Kubernetes cluster
	//
	// 	- ManagedKubernetes: ACK cluster
	//
	// example:
	//
	// Ask
	SubClusterType *string `json:"SubClusterType,omitempty" xml:"SubClusterType,omitempty"`
}

func (s GetK8sClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s GetK8sClusterRequest) GoString() string {
	return s.String()
}

func (s *GetK8sClusterRequest) GetClusterType() *int32 {
	return s.ClusterType
}

func (s *GetK8sClusterRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetK8sClusterRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetK8sClusterRequest) GetRegionTag() *string {
	return s.RegionTag
}

func (s *GetK8sClusterRequest) GetSubClusterType() *string {
	return s.SubClusterType
}

func (s *GetK8sClusterRequest) SetClusterType(v int32) *GetK8sClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *GetK8sClusterRequest) SetCurrentPage(v int32) *GetK8sClusterRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetK8sClusterRequest) SetPageSize(v int32) *GetK8sClusterRequest {
	s.PageSize = &v
	return s
}

func (s *GetK8sClusterRequest) SetRegionTag(v string) *GetK8sClusterRequest {
	s.RegionTag = &v
	return s
}

func (s *GetK8sClusterRequest) SetSubClusterType(v string) *GetK8sClusterRequest {
	s.SubClusterType = &v
	return s
}

func (s *GetK8sClusterRequest) Validate() error {
	return dara.Validate(s)
}
