// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterIds(v []*string) *ListClustersRequest
	GetClusterIds() []*string
	SetClusterName(v string) *ListClustersRequest
	GetClusterName() *string
	SetClusterStates(v []*string) *ListClustersRequest
	GetClusterStates() []*string
	SetClusterTypes(v []*string) *ListClustersRequest
	GetClusterTypes() []*string
	SetMaxResults(v int32) *ListClustersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListClustersRequest
	GetNextToken() *string
	SetPaymentTypes(v []*string) *ListClustersRequest
	GetPaymentTypes() []*string
	SetRegionId(v string) *ListClustersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListClustersRequest
	GetResourceGroupId() *string
	SetTags(v []*Tag) *ListClustersRequest
	GetTags() []*Tag
}

type ListClustersRequest struct {
	// The cluster IDs. Number of elements in the array: 1 to 100.
	//
	// example:
	//
	// rg-acfmzabjyop****
	ClusterIds []*string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty" type:"Repeated"`
	// The name of the cluster.
	//
	// example:
	//
	// emrtest
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The states of the clusters. Number of elements in the array: 1 to 100.
	//
	// example:
	//
	// ["HADOOP"]
	ClusterStates []*string `json:"ClusterStates,omitempty" xml:"ClusterStates,omitempty" type:"Repeated"`
	// The list of cluster types. Number of elements in the array: 1 to 100.
	//
	// example:
	//
	// ["c-b933c5aac8fe****"]
	ClusterTypes []*string `json:"ClusterTypes,omitempty" xml:"ClusterTypes,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The starting point of the current query. If you do not configure this parameter, the query starts from the beginning.
	//
	// example:
	//
	// eyJlY21OZXh0VG9rZW4iOiIxIiwidGFpaGFvTmV4dFRva2VuIjoiNTYiLCJ0YWloYW9OZXh0VG9rZW5JbnQiOjU2LCJlY21OZXh0VG9rZW5JbnQiOjF9
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The billing methods. You can specify a maximum of 2 items.
	//
	// example:
	//
	// ["ECS"]
	PaymentTypes []*string `json:"PaymentTypes,omitempty" xml:"PaymentTypes,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzabjyop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag list. Number of elements in the array: 1 to 20.
	//
	// example:
	//
	// ["PayAsYouGo"]
	Tags []*Tag `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) GetClusterIds() []*string {
	return s.ClusterIds
}

func (s *ListClustersRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListClustersRequest) GetClusterStates() []*string {
	return s.ClusterStates
}

func (s *ListClustersRequest) GetClusterTypes() []*string {
	return s.ClusterTypes
}

func (s *ListClustersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListClustersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListClustersRequest) GetPaymentTypes() []*string {
	return s.PaymentTypes
}

func (s *ListClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListClustersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListClustersRequest) GetTags() []*Tag {
	return s.Tags
}

func (s *ListClustersRequest) SetClusterIds(v []*string) *ListClustersRequest {
	s.ClusterIds = v
	return s
}

func (s *ListClustersRequest) SetClusterName(v string) *ListClustersRequest {
	s.ClusterName = &v
	return s
}

func (s *ListClustersRequest) SetClusterStates(v []*string) *ListClustersRequest {
	s.ClusterStates = v
	return s
}

func (s *ListClustersRequest) SetClusterTypes(v []*string) *ListClustersRequest {
	s.ClusterTypes = v
	return s
}

func (s *ListClustersRequest) SetMaxResults(v int32) *ListClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListClustersRequest) SetNextToken(v string) *ListClustersRequest {
	s.NextToken = &v
	return s
}

func (s *ListClustersRequest) SetPaymentTypes(v []*string) *ListClustersRequest {
	s.PaymentTypes = v
	return s
}

func (s *ListClustersRequest) SetRegionId(v string) *ListClustersRequest {
	s.RegionId = &v
	return s
}

func (s *ListClustersRequest) SetResourceGroupId(v string) *ListClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClustersRequest) SetTags(v []*Tag) *ListClustersRequest {
	s.Tags = v
	return s
}

func (s *ListClustersRequest) Validate() error {
	return dara.Validate(s)
}
