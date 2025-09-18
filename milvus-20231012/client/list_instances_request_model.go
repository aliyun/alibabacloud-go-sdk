// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListInstancesRequest
	GetClusterId() *string
	SetClusterName(v string) *ListInstancesRequest
	GetClusterName() *string
	SetPageNumber(v int32) *ListInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListInstancesRequest
	GetResourceGroupId() *string
	SetTag(v []*ListInstancesRequestTag) *ListInstancesRequest
	GetTag() []*ListInstancesRequestTag
}

type ListInstancesRequest struct {
	// The instance ID.
	//
	// example:
	//
	// c-123xxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// milvus-test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-123xxx
	ResourceGroupId *string                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*ListInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListInstancesRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesRequest) GetTag() []*ListInstancesRequestTag {
	return s.Tag
}

func (s *ListInstancesRequest) SetClusterId(v string) *ListInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInstancesRequest) SetClusterName(v string) *ListInstancesRequest {
	s.ClusterName = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetRegionId(v string) *ListInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstancesRequest) SetResourceGroupId(v string) *ListInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesRequest) SetTag(v []*ListInstancesRequestTag) *ListInstancesRequest {
	s.Tag = v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	return dara.Validate(s)
}

type ListInstancesRequestTag struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListInstancesRequestTag) SetKey(v string) *ListInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListInstancesRequestTag) SetValue(v string) *ListInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *ListInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
