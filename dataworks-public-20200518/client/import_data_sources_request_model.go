// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDataSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSources(v string) *ImportDataSourcesRequest
	GetDataSources() *string
	SetProjectId(v int64) *ImportDataSourcesRequest
	GetProjectId() *int64
}

type ImportDataSourcesRequest struct {
	// The configurations of the data sources that you want to import. The Name, DataSourceType, SubType, Description, Content, and EnvType parameters are required. For more information about the parameters, see [CreateDataSource](https://help.aliyun.com/document_detail/211429.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"SubType":"","DataSourceType":"mysql","EnvType":1,"Name":"mysql_dms2","Description":"aaa","Content":"{\\"database\\":\\"mysql_chengdu_old\\",\\"password\\":\\"***\\",\\"instanceName\\":\\"rm-2vcrckb37163g7l3w\\",\\"regionId\\":\\"cn-chengdu\\",\\"tag\\":\\"rds\\",\\"rdsOwnerId\\":\\"333\\",\\"username\\":\\"mysql_chengdu2\\"}"},{"SubType":"","DataSourceType":"mysql","EnvType":1,"Name":"mysql_dms2","Description":"aaa","Content":"{\\"database\\":\\"mysql_chengdu_old\\",\\"password\\":\\"***\\",\\"instanceName\\":\\"rm-2vcrckb37163g7l3w\\",\\"regionId\\":\\"cn-chengdu\\",\\"tag\\":\\"rds\\",\\"rdsOwnerId\\":\\"143\\",\\"username\\":\\"mysql_chengdu2\\"}"}]
	DataSources *string `json:"DataSources,omitempty" xml:"DataSources,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ImportDataSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ImportDataSourcesRequest) GetDataSources() *string {
	return s.DataSources
}

func (s *ImportDataSourcesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ImportDataSourcesRequest) SetDataSources(v string) *ImportDataSourcesRequest {
	s.DataSources = &v
	return s
}

func (s *ImportDataSourcesRequest) SetProjectId(v int64) *ImportDataSourcesRequest {
	s.ProjectId = &v
	return s
}

func (s *ImportDataSourcesRequest) Validate() error {
	return dara.Validate(s)
}
