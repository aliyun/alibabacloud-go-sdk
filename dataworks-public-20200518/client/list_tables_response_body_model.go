// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListTablesResponseBodyData) *ListTablesResponseBody
	GetData() *ListTablesResponseBodyData
	SetRequestId(v string) *ListTablesResponseBody
	GetRequestId() *string
}

type ListTablesResponseBody struct {
	// The returned data.
	Data *ListTablesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E25887B7-579C-54A5-9C4F-83A0DE367DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBody) GetData() *ListTablesResponseBodyData {
	return s.Data
}

func (s *ListTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTablesResponseBody) SetData(v *ListTablesResponseBodyData) *ListTablesResponseBody {
	s.Data = v
	return s
}

func (s *ListTablesResponseBody) SetRequestId(v string) *ListTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTablesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTablesResponseBodyData struct {
	// Pagination information, which specifies the starting point of the next read.
	//
	// example:
	//
	// AAAAAVY3rYiv9VoUJQSiCitgjgSwg+byk0FIjirFkm4zfM4G0xYwM/FQvOhgrTHsCPIZ5yqXYu2NG6qRCRC52HvwbOA=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// An array of entities.
	TableEntityList []*ListTablesResponseBodyDataTableEntityList `json:"TableEntityList,omitempty" xml:"TableEntityList,omitempty" type:"Repeated"`
	// The total number.
	//
	// example:
	//
	// 100
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListTablesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTablesResponseBodyData) GetTableEntityList() []*ListTablesResponseBodyDataTableEntityList {
	return s.TableEntityList
}

func (s *ListTablesResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListTablesResponseBodyData) SetNextToken(v string) *ListTablesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListTablesResponseBodyData) SetTableEntityList(v []*ListTablesResponseBodyDataTableEntityList) *ListTablesResponseBodyData {
	s.TableEntityList = v
	return s
}

func (s *ListTablesResponseBodyData) SetTotal(v int64) *ListTablesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListTablesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListTablesResponseBodyDataTableEntityList struct {
	// The information about the table.
	EntityContent *ListTablesResponseBodyDataTableEntityListEntityContent `json:"EntityContent,omitempty" xml:"EntityContent,omitempty" type:"Struct"`
	// The unique identifier of the table entity.
	//
	// example:
	//
	// maxcompute-table.project.table
	EntityQualifiedName *string `json:"EntityQualifiedName,omitempty" xml:"EntityQualifiedName,omitempty"`
}

func (s ListTablesResponseBodyDataTableEntityList) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyDataTableEntityList) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyDataTableEntityList) GetEntityContent() *ListTablesResponseBodyDataTableEntityListEntityContent {
	return s.EntityContent
}

func (s *ListTablesResponseBodyDataTableEntityList) GetEntityQualifiedName() *string {
	return s.EntityQualifiedName
}

func (s *ListTablesResponseBodyDataTableEntityList) SetEntityContent(v *ListTablesResponseBodyDataTableEntityListEntityContent) *ListTablesResponseBodyDataTableEntityList {
	s.EntityContent = v
	return s
}

func (s *ListTablesResponseBodyDataTableEntityList) SetEntityQualifiedName(v string) *ListTablesResponseBodyDataTableEntityList {
	s.EntityQualifiedName = &v
	return s
}

func (s *ListTablesResponseBodyDataTableEntityList) Validate() error {
	return dara.Validate(s)
}

type ListTablesResponseBodyDataTableEntityListEntityContent struct {
	// The unique identifier of the data source.
	//
	// example:
	//
	// accountId:cn-shanghai:odps:project
	DataSourceQualifiedName *string `json:"DataSourceQualifiedName,omitempty" xml:"DataSourceQualifiedName,omitempty"`
	// The unique ID of the data source identifier.
	//
	// example:
	//
	// e70f92239d491057f6a2563b545bdaf8cc6b537d9dc55ec84c55f7cfefg
	DataSourceUniqueId *string `json:"DataSourceUniqueId,omitempty" xml:"DataSourceUniqueId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// database
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The ID of the data source instance.
	//
	// example:
	//
	// rm-uf6rn0123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the ODPS project.
	//
	// example:
	//
	// project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListTablesResponseBodyDataTableEntityListEntityContent) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyDataTableEntityListEntityContent) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyDataTableEntityListEntityContent) GetDataSourceQualifiedName() *string {
	return s.DataSourceQualifiedName
}

func (s *ListTablesResponseBodyDataTableEntityListEntityContent) GetDataSourceUniqueId() *string {
	return s.DataSourceUniqueId
}

func (s *ListTablesResponseBodyDataTableEntityListEntityContent) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListTablesResponseBodyDataTableEntityListEntityContent) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTablesResponseBodyDataTableEntityListEntityContent) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListTablesResponseBodyDataTableEntityListEntityContent) GetTableName() *string {
	return s.TableName
}

func (s *ListTablesResponseBodyDataTableEntityListEntityContent) SetDataSourceQualifiedName(v string) *ListTablesResponseBodyDataTableEntityListEntityContent {
	s.DataSourceQualifiedName = &v
	return s
}

func (s *ListTablesResponseBodyDataTableEntityListEntityContent) SetDataSourceUniqueId(v string) *ListTablesResponseBodyDataTableEntityListEntityContent {
	s.DataSourceUniqueId = &v
	return s
}

func (s *ListTablesResponseBodyDataTableEntityListEntityContent) SetDatabaseName(v string) *ListTablesResponseBodyDataTableEntityListEntityContent {
	s.DatabaseName = &v
	return s
}

func (s *ListTablesResponseBodyDataTableEntityListEntityContent) SetInstanceId(v string) *ListTablesResponseBodyDataTableEntityListEntityContent {
	s.InstanceId = &v
	return s
}

func (s *ListTablesResponseBodyDataTableEntityListEntityContent) SetProjectName(v string) *ListTablesResponseBodyDataTableEntityListEntityContent {
	s.ProjectName = &v
	return s
}

func (s *ListTablesResponseBodyDataTableEntityListEntityContent) SetTableName(v string) *ListTablesResponseBodyDataTableEntityListEntityContent {
	s.TableName = &v
	return s
}

func (s *ListTablesResponseBodyDataTableEntityListEntityContent) Validate() error {
	return dara.Validate(s)
}
