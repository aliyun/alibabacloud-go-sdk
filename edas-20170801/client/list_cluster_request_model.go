// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogicalRegionId(v string) *ListClusterRequest
	GetLogicalRegionId() *string
	SetResourceGroupId(v string) *ListClusterRequest
	GetResourceGroupId() *string
}

type ListClusterRequest struct {
	// The ID of the namespace. You can call the ListUserDefineRegion operation to query the namespace ID. For more information, see [ListUserDefineRegion](https://help.aliyun.com/document_detail/149377.html).
	//
	// 	- If this parameter is left empty, the clusters in the default namespace are queried.
	//
	// 	- If this parameter is specified, the clusters in the specified namespace are queried.
	//
	// example:
	//
	// cn-beijing:test
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
	// The ID of the resource group. You can call the ListResourceGroup operation to query the resource group ID. For more information, see [ListResourceGroup](https://help.aliyun.com/document_detail/62055.html).
	//
	// 	- If this parameter is left empty, the clusters in the default resource group are queried.
	//
	// 	- If this parameter is specified, the clusters in the specified resource group are queried.
	//
	// example:
	//
	// 461
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterRequest) GoString() string {
	return s.String()
}

func (s *ListClusterRequest) GetLogicalRegionId() *string {
	return s.LogicalRegionId
}

func (s *ListClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListClusterRequest) SetLogicalRegionId(v string) *ListClusterRequest {
	s.LogicalRegionId = &v
	return s
}

func (s *ListClusterRequest) SetResourceGroupId(v string) *ListClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClusterRequest) Validate() error {
	return dara.Validate(s)
}
