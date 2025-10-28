// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListClusterMembersRequest
	GetClusterId() *string
	SetCurrentPage(v int32) *ListClusterMembersRequest
	GetCurrentPage() *int32
	SetEcsList(v string) *ListClusterMembersRequest
	GetEcsList() *string
	SetPageSize(v int32) *ListClusterMembersRequest
	GetPageSize() *int32
}

type ListClusterMembersRequest struct {
	// The ID of the cluster. You can call the ListCluster operation to query the cluster ID. For more information, see [ListCluster](https://help.aliyun.com/document_detail/154995.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 52984524-6d48-4bbd-85f2-a34b0e5b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of the page to return. If you do not specify this parameter, the first page is returned.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of ECS instances.
	//
	// example:
	//
	// 1
	EcsList *string `json:"EcsList,omitempty" xml:"EcsList,omitempty"`
	// The number of ECS instances to return on each page. If you do not specify this parameter, all ECS instances in the specified cluster are returned on one page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClusterMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterMembersRequest) GoString() string {
	return s.String()
}

func (s *ListClusterMembersRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterMembersRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListClusterMembersRequest) GetEcsList() *string {
	return s.EcsList
}

func (s *ListClusterMembersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClusterMembersRequest) SetClusterId(v string) *ListClusterMembersRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterMembersRequest) SetCurrentPage(v int32) *ListClusterMembersRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListClusterMembersRequest) SetEcsList(v string) *ListClusterMembersRequest {
	s.EcsList = &v
	return s
}

func (s *ListClusterMembersRequest) SetPageSize(v int32) *ListClusterMembersRequest {
	s.PageSize = &v
	return s
}

func (s *ListClusterMembersRequest) Validate() error {
	return dara.Validate(s)
}
