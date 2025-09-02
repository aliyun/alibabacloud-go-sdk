// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestNetworkConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasourceName(v string) *TestNetworkConnectionRequest
	GetDatasourceName() *string
	SetEnvType(v string) *TestNetworkConnectionRequest
	GetEnvType() *string
	SetProjectId(v int64) *TestNetworkConnectionRequest
	GetProjectId() *int64
	SetResourceGroup(v string) *TestNetworkConnectionRequest
	GetResourceGroup() *string
}

type TestNetworkConnectionRequest struct {
	// The name of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql_name
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// The environment in which the data source resides. Valid values:
	//
	// 	- 0: development environment
	//
	// 	- 1: production environment
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The ID of the DataWorks workspace to which the data source belongs. You can call the [ListProjects](https://help.aliyun.com/document_detail/178393.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The identifier of the resource group. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/173913.html) operation to query the identifier of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// S_res_group_2XXXX4_1619100XXXXX
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
}

func (s TestNetworkConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s TestNetworkConnectionRequest) GoString() string {
	return s.String()
}

func (s *TestNetworkConnectionRequest) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *TestNetworkConnectionRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *TestNetworkConnectionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *TestNetworkConnectionRequest) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *TestNetworkConnectionRequest) SetDatasourceName(v string) *TestNetworkConnectionRequest {
	s.DatasourceName = &v
	return s
}

func (s *TestNetworkConnectionRequest) SetEnvType(v string) *TestNetworkConnectionRequest {
	s.EnvType = &v
	return s
}

func (s *TestNetworkConnectionRequest) SetProjectId(v int64) *TestNetworkConnectionRequest {
	s.ProjectId = &v
	return s
}

func (s *TestNetworkConnectionRequest) SetResourceGroup(v string) *TestNetworkConnectionRequest {
	s.ResourceGroup = &v
	return s
}

func (s *TestNetworkConnectionRequest) Validate() error {
	return dara.Validate(s)
}
