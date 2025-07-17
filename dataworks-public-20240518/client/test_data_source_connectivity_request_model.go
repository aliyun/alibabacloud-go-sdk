// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestDataSourceConnectivityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v int64) *TestDataSourceConnectivityRequest
	GetDataSourceId() *int64
	SetProjectId(v int64) *TestDataSourceConnectivityRequest
	GetProjectId() *int64
	SetResourceGroupId(v string) *TestDataSourceConnectivityRequest
	GetResourceGroupId() *string
}

type TestDataSourceConnectivityRequest struct {
	// The ID of the data source for which you want to test the network connectivity.
	//
	// This parameter is required.
	//
	// example:
	//
	// 144544
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s TestDataSourceConnectivityRequest) String() string {
	return dara.Prettify(s)
}

func (s TestDataSourceConnectivityRequest) GoString() string {
	return s.String()
}

func (s *TestDataSourceConnectivityRequest) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *TestDataSourceConnectivityRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *TestDataSourceConnectivityRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *TestDataSourceConnectivityRequest) SetDataSourceId(v int64) *TestDataSourceConnectivityRequest {
	s.DataSourceId = &v
	return s
}

func (s *TestDataSourceConnectivityRequest) SetProjectId(v int64) *TestDataSourceConnectivityRequest {
	s.ProjectId = &v
	return s
}

func (s *TestDataSourceConnectivityRequest) SetResourceGroupId(v string) *TestDataSourceConnectivityRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *TestDataSourceConnectivityRequest) Validate() error {
	return dara.Validate(s)
}
