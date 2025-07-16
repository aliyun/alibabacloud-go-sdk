// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLabelTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLabelTables(v []*ListLabelTablesResponseBodyLabelTables) *ListLabelTablesResponseBody
	GetLabelTables() []*ListLabelTablesResponseBodyLabelTables
	SetRequestId(v string) *ListLabelTablesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListLabelTablesResponseBody
	GetTotalCount() *int64
}

type ListLabelTablesResponseBody struct {
	LabelTables []*ListLabelTablesResponseBodyLabelTables `json:"LabelTables,omitempty" xml:"LabelTables,omitempty" type:"Repeated"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 21
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLabelTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLabelTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLabelTablesResponseBody) GetLabelTables() []*ListLabelTablesResponseBodyLabelTables {
	return s.LabelTables
}

func (s *ListLabelTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLabelTablesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListLabelTablesResponseBody) SetLabelTables(v []*ListLabelTablesResponseBodyLabelTables) *ListLabelTablesResponseBody {
	s.LabelTables = v
	return s
}

func (s *ListLabelTablesResponseBody) SetRequestId(v string) *ListLabelTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLabelTablesResponseBody) SetTotalCount(v int64) *ListLabelTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLabelTablesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLabelTablesResponseBodyLabelTables struct {
	// example:
	//
	// 3
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// datasource1
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 3
	LabelTableId *string `json:"LabelTableId,omitempty" xml:"LabelTableId,omitempty"`
	// example:
	//
	// label_table1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 123214213214
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListLabelTablesResponseBodyLabelTables) String() string {
	return dara.Prettify(s)
}

func (s ListLabelTablesResponseBodyLabelTables) GoString() string {
	return s.String()
}

func (s *ListLabelTablesResponseBodyLabelTables) GetDatasourceId() *string {
	return s.DatasourceId
}

func (s *ListLabelTablesResponseBodyLabelTables) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *ListLabelTablesResponseBodyLabelTables) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListLabelTablesResponseBodyLabelTables) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListLabelTablesResponseBodyLabelTables) GetLabelTableId() *string {
	return s.LabelTableId
}

func (s *ListLabelTablesResponseBodyLabelTables) GetName() *string {
	return s.Name
}

func (s *ListLabelTablesResponseBodyLabelTables) GetOwner() *string {
	return s.Owner
}

func (s *ListLabelTablesResponseBodyLabelTables) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListLabelTablesResponseBodyLabelTables) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListLabelTablesResponseBodyLabelTables) SetDatasourceId(v string) *ListLabelTablesResponseBodyLabelTables {
	s.DatasourceId = &v
	return s
}

func (s *ListLabelTablesResponseBodyLabelTables) SetDatasourceName(v string) *ListLabelTablesResponseBodyLabelTables {
	s.DatasourceName = &v
	return s
}

func (s *ListLabelTablesResponseBodyLabelTables) SetGmtCreateTime(v string) *ListLabelTablesResponseBodyLabelTables {
	s.GmtCreateTime = &v
	return s
}

func (s *ListLabelTablesResponseBodyLabelTables) SetGmtModifiedTime(v string) *ListLabelTablesResponseBodyLabelTables {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListLabelTablesResponseBodyLabelTables) SetLabelTableId(v string) *ListLabelTablesResponseBodyLabelTables {
	s.LabelTableId = &v
	return s
}

func (s *ListLabelTablesResponseBodyLabelTables) SetName(v string) *ListLabelTablesResponseBodyLabelTables {
	s.Name = &v
	return s
}

func (s *ListLabelTablesResponseBodyLabelTables) SetOwner(v string) *ListLabelTablesResponseBodyLabelTables {
	s.Owner = &v
	return s
}

func (s *ListLabelTablesResponseBodyLabelTables) SetProjectId(v string) *ListLabelTablesResponseBodyLabelTables {
	s.ProjectId = &v
	return s
}

func (s *ListLabelTablesResponseBodyLabelTables) SetProjectName(v string) *ListLabelTablesResponseBodyLabelTables {
	s.ProjectName = &v
	return s
}

func (s *ListLabelTablesResponseBodyLabelTables) Validate() error {
	return dara.Validate(s)
}
