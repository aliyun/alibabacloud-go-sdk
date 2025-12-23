// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResourceGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceGroupsRequest
	GetNextToken() *string
	SetResourceGroupName(v string) *ListResourceGroupsRequest
	GetResourceGroupName() *string
	SetResourceGroupType(v string) *ListResourceGroupsRequest
	GetResourceGroupType() *string
	SetWorkspaceId(v string) *ListResourceGroupsRequest
	GetWorkspaceId() *string
}

type ListResourceGroupsRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// name
	ResourceGroupName *string `json:"resourceGroupName,omitempty" xml:"resourceGroupName,omitempty"`
	// example:
	//
	// CLUSTER_RESOURCE_GROUP
	ResourceGroupType *string `json:"resourceGroupType,omitempty" xml:"resourceGroupType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// w-lxyy60mpgpg****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListResourceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceGroupsRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ListResourceGroupsRequest) GetResourceGroupType() *string {
	return s.ResourceGroupType
}

func (s *ListResourceGroupsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListResourceGroupsRequest) SetMaxResults(v int32) *ListResourceGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceGroupsRequest) SetNextToken(v string) *ListResourceGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceGroupsRequest) SetResourceGroupName(v string) *ListResourceGroupsRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *ListResourceGroupsRequest) SetResourceGroupType(v string) *ListResourceGroupsRequest {
	s.ResourceGroupType = &v
	return s
}

func (s *ListResourceGroupsRequest) SetWorkspaceId(v string) *ListResourceGroupsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListResourceGroupsRequest) Validate() error {
	return dara.Validate(s)
}
