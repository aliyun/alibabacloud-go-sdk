// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableThemeLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceType(v string) *GetMetaTableThemeLevelRequest
	GetDataSourceType() *string
	SetTableGuid(v string) *GetMetaTableThemeLevelRequest
	GetTableGuid() *string
}

type GetMetaTableThemeLevelRequest struct {
	// The type of the data source. Set the value to odps.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The GUID of the metatable. Specify the GUID in the format of odps.${projectName}.${tableName}.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps.project1.name1
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s GetMetaTableThemeLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableThemeLevelRequest) GoString() string {
	return s.String()
}

func (s *GetMetaTableThemeLevelRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *GetMetaTableThemeLevelRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTableThemeLevelRequest) SetDataSourceType(v string) *GetMetaTableThemeLevelRequest {
	s.DataSourceType = &v
	return s
}

func (s *GetMetaTableThemeLevelRequest) SetTableGuid(v string) *GetMetaTableThemeLevelRequest {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTableThemeLevelRequest) Validate() error {
	return dara.Validate(s)
}
