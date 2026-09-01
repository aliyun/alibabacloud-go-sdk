// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterInterceptionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterCNNFStatus(v int32) *ListClusterInterceptionConfigRequest
	GetClusterCNNFStatus() *int32
	SetClusterId(v string) *ListClusterInterceptionConfigRequest
	GetClusterId() *string
	SetClusterName(v string) *ListClusterInterceptionConfigRequest
	GetClusterName() *string
	SetCurrentPage(v int32) *ListClusterInterceptionConfigRequest
	GetCurrentPage() *int32
	SetExcludeClusterTypes(v []*string) *ListClusterInterceptionConfigRequest
	GetExcludeClusterTypes() []*string
	SetPageSize(v int32) *ListClusterInterceptionConfigRequest
	GetPageSize() *int32
}

type ListClusterInterceptionConfigRequest struct {
	// The container firewall status of the cluster. Valid values:
	//
	// - **-1**: unknown
	//
	// - **0**: abnormal
	//
	// - **1**: normal
	//
	// - **2**: normal pending confirmation
	//
	// example:
	//
	// 1
	ClusterCNNFStatus *int32 `json:"ClusterCNNFStatus,omitempty" xml:"ClusterCNNFStatus,omitempty"`
	// The ID of the container cluster.
	//
	// >You can call the [DescribeContainerInstances](~~DescribeContainerInstances~~) operation to obtain this parameter.
	//
	// example:
	//
	// c22143730ab6e40b09ec7c1c51d4d****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// sas
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The page number of the page to return. Default value: 1, which indicates the first page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of excluded cluster types.
	ExcludeClusterTypes []*string `json:"ExcludeClusterTypes,omitempty" xml:"ExcludeClusterTypes,omitempty" type:"Repeated"`
	// The maximum number of entries per page in a paging query. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClusterInterceptionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterInterceptionConfigRequest) GoString() string {
	return s.String()
}

func (s *ListClusterInterceptionConfigRequest) GetClusterCNNFStatus() *int32 {
	return s.ClusterCNNFStatus
}

func (s *ListClusterInterceptionConfigRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterInterceptionConfigRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListClusterInterceptionConfigRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListClusterInterceptionConfigRequest) GetExcludeClusterTypes() []*string {
	return s.ExcludeClusterTypes
}

func (s *ListClusterInterceptionConfigRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClusterInterceptionConfigRequest) SetClusterCNNFStatus(v int32) *ListClusterInterceptionConfigRequest {
	s.ClusterCNNFStatus = &v
	return s
}

func (s *ListClusterInterceptionConfigRequest) SetClusterId(v string) *ListClusterInterceptionConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterInterceptionConfigRequest) SetClusterName(v string) *ListClusterInterceptionConfigRequest {
	s.ClusterName = &v
	return s
}

func (s *ListClusterInterceptionConfigRequest) SetCurrentPage(v int32) *ListClusterInterceptionConfigRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListClusterInterceptionConfigRequest) SetExcludeClusterTypes(v []*string) *ListClusterInterceptionConfigRequest {
	s.ExcludeClusterTypes = v
	return s
}

func (s *ListClusterInterceptionConfigRequest) SetPageSize(v int32) *ListClusterInterceptionConfigRequest {
	s.PageSize = &v
	return s
}

func (s *ListClusterInterceptionConfigRequest) Validate() error {
	return dara.Validate(s)
}
