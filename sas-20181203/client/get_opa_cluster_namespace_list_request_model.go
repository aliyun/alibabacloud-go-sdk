// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaClusterNamespaceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetOpaClusterNamespaceListRequest
	GetClusterId() *string
	SetCurrentPage(v int32) *GetOpaClusterNamespaceListRequest
	GetCurrentPage() *int32
	SetNameSpaceName(v string) *GetOpaClusterNamespaceListRequest
	GetNameSpaceName() *string
	SetPageSize(v int32) *GetOpaClusterNamespaceListRequest
	GetPageSize() *int32
}

type GetOpaClusterNamespaceListRequest struct {
	// The ID of the cluster that you want to query.
	//
	// >  You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of clusters.
	//
	// example:
	//
	// ca0faf43e33904ecbb6c695df7906****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// sit
	NameSpaceName *string `json:"NameSpaceName,omitempty" xml:"NameSpaceName,omitempty"`
	// The number of entries per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetOpaClusterNamespaceListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOpaClusterNamespaceListRequest) GoString() string {
	return s.String()
}

func (s *GetOpaClusterNamespaceListRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetOpaClusterNamespaceListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetOpaClusterNamespaceListRequest) GetNameSpaceName() *string {
	return s.NameSpaceName
}

func (s *GetOpaClusterNamespaceListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOpaClusterNamespaceListRequest) SetClusterId(v string) *GetOpaClusterNamespaceListRequest {
	s.ClusterId = &v
	return s
}

func (s *GetOpaClusterNamespaceListRequest) SetCurrentPage(v int32) *GetOpaClusterNamespaceListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetOpaClusterNamespaceListRequest) SetNameSpaceName(v string) *GetOpaClusterNamespaceListRequest {
	s.NameSpaceName = &v
	return s
}

func (s *GetOpaClusterNamespaceListRequest) SetPageSize(v int32) *GetOpaClusterNamespaceListRequest {
	s.PageSize = &v
	return s
}

func (s *GetOpaClusterNamespaceListRequest) Validate() error {
	return dara.Validate(s)
}
