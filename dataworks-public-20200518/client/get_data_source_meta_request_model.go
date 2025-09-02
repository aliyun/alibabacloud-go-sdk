// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasourceName(v string) *GetDataSourceMetaRequest
	GetDatasourceName() *string
	SetEnvType(v string) *GetDataSourceMetaRequest
	GetEnvType() *string
	SetPageNumber(v int64) *GetDataSourceMetaRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetDataSourceMetaRequest
	GetPageSize() *int64
	SetProjectId(v int64) *GetDataSourceMetaRequest
	GetProjectId() *int64
}

type GetDataSourceMetaRequest struct {
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
	// example:
	//
	// 1
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetDataSourceMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceMetaRequest) GoString() string {
	return s.String()
}

func (s *GetDataSourceMetaRequest) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *GetDataSourceMetaRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *GetDataSourceMetaRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetDataSourceMetaRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetDataSourceMetaRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataSourceMetaRequest) SetDatasourceName(v string) *GetDataSourceMetaRequest {
	s.DatasourceName = &v
	return s
}

func (s *GetDataSourceMetaRequest) SetEnvType(v string) *GetDataSourceMetaRequest {
	s.EnvType = &v
	return s
}

func (s *GetDataSourceMetaRequest) SetPageNumber(v int64) *GetDataSourceMetaRequest {
	s.PageNumber = &v
	return s
}

func (s *GetDataSourceMetaRequest) SetPageSize(v int64) *GetDataSourceMetaRequest {
	s.PageSize = &v
	return s
}

func (s *GetDataSourceMetaRequest) SetProjectId(v int64) *GetDataSourceMetaRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDataSourceMetaRequest) Validate() error {
	return dara.Validate(s)
}
