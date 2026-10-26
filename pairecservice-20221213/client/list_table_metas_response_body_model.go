// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableMetasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTableMetasResponseBody
	GetRequestId() *string
	SetTableMetas(v []*ListTableMetasResponseBodyTableMetas) *ListTableMetasResponseBody
	GetTableMetas() []*ListTableMetasResponseBodyTableMetas
	SetTotalCount(v int64) *ListTableMetasResponseBody
	GetTotalCount() *int64
}

type ListTableMetasResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of data tables.
	TableMetas []*ListTableMetasResponseBodyTableMetas `json:"TableMetas,omitempty" xml:"TableMetas,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTableMetasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTableMetasResponseBody) GoString() string {
	return s.String()
}

func (s *ListTableMetasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTableMetasResponseBody) GetTableMetas() []*ListTableMetasResponseBodyTableMetas {
	return s.TableMetas
}

func (s *ListTableMetasResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTableMetasResponseBody) SetRequestId(v string) *ListTableMetasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTableMetasResponseBody) SetTableMetas(v []*ListTableMetasResponseBodyTableMetas) *ListTableMetasResponseBody {
	s.TableMetas = v
	return s
}

func (s *ListTableMetasResponseBody) SetTotalCount(v int64) *ListTableMetasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTableMetasResponseBody) Validate() error {
	if s.TableMetas != nil {
		for _, item := range s.TableMetas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTableMetasResponseBodyTableMetas struct {
	// Indicates whether the data table can be deleted. The value is false if the table is referenced elsewhere.
	//
	// example:
	//
	// true
	CanDelete *bool `json:"CanDelete,omitempty" xml:"CanDelete,omitempty"`
	// The additional configurations of the data table.
	//
	// example:
	//
	// {"table_type":"item"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The description of the data table.
	//
	// example:
	//
	// this is a test table
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of fields in the table.
	Fields []*ListTableMetasResponseBodyTableMetasFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The creation time. The format is YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2021-12-12 12:24:33
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the fields were imported.
	//
	// example:
	//
	// imprecation
	GmtImportedTime *string `json:"GmtImportedTime,omitempty" xml:"GmtImportedTime,omitempty"`
	// The update time. The format is YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2021-12-12 12:24:33
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The module of the data table.
	//
	// - ABTest: Experiment data tables
	//
	// - ExperimentTool: Experiment tool tables
	//
	// - DataDiagnosis: Data diagnosis
	//
	// example:
	//
	// ABTest
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The name of the data table.
	//
	// example:
	//
	// test_table
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// reso-2s416t***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// Maxcompute
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The resource URI.
	//
	// example:
	//
	// test
	ResourceUri *string `json:"ResourceUri,omitempty" xml:"ResourceUri,omitempty"`
	// The ID of the data table.
	//
	// example:
	//
	// 3
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// The name of the table in the database.
	//
	// example:
	//
	// table_mysql
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The type of the data table.
	//
	// - MaxCompute
	//
	// - Hologres
	//
	// example:
	//
	// MaxCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the table in DataWorks Data Map. This parameter is returned only if the table is a MaxCompute table.
	//
	// example:
	//
	// https://dmc-xxx.com/dm/table/xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListTableMetasResponseBodyTableMetas) String() string {
	return dara.Prettify(s)
}

func (s ListTableMetasResponseBodyTableMetas) GoString() string {
	return s.String()
}

func (s *ListTableMetasResponseBodyTableMetas) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *ListTableMetasResponseBodyTableMetas) GetConfig() *string {
	return s.Config
}

func (s *ListTableMetasResponseBodyTableMetas) GetDescription() *string {
	return s.Description
}

func (s *ListTableMetasResponseBodyTableMetas) GetFields() []*ListTableMetasResponseBodyTableMetasFields {
	return s.Fields
}

func (s *ListTableMetasResponseBodyTableMetas) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListTableMetasResponseBodyTableMetas) GetGmtImportedTime() *string {
	return s.GmtImportedTime
}

func (s *ListTableMetasResponseBodyTableMetas) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListTableMetasResponseBodyTableMetas) GetModule() *string {
	return s.Module
}

func (s *ListTableMetasResponseBodyTableMetas) GetName() *string {
	return s.Name
}

func (s *ListTableMetasResponseBodyTableMetas) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTableMetasResponseBodyTableMetas) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTableMetasResponseBodyTableMetas) GetResourceUri() *string {
	return s.ResourceUri
}

func (s *ListTableMetasResponseBodyTableMetas) GetTableMetaId() *string {
	return s.TableMetaId
}

func (s *ListTableMetasResponseBodyTableMetas) GetTableName() *string {
	return s.TableName
}

func (s *ListTableMetasResponseBodyTableMetas) GetType() *string {
	return s.Type
}

func (s *ListTableMetasResponseBodyTableMetas) GetUrl() *string {
	return s.Url
}

func (s *ListTableMetasResponseBodyTableMetas) SetCanDelete(v bool) *ListTableMetasResponseBodyTableMetas {
	s.CanDelete = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetConfig(v string) *ListTableMetasResponseBodyTableMetas {
	s.Config = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetDescription(v string) *ListTableMetasResponseBodyTableMetas {
	s.Description = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetFields(v []*ListTableMetasResponseBodyTableMetasFields) *ListTableMetasResponseBodyTableMetas {
	s.Fields = v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetGmtCreateTime(v string) *ListTableMetasResponseBodyTableMetas {
	s.GmtCreateTime = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetGmtImportedTime(v string) *ListTableMetasResponseBodyTableMetas {
	s.GmtImportedTime = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetGmtModifiedTime(v string) *ListTableMetasResponseBodyTableMetas {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetModule(v string) *ListTableMetasResponseBodyTableMetas {
	s.Module = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetName(v string) *ListTableMetasResponseBodyTableMetas {
	s.Name = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetResourceId(v string) *ListTableMetasResponseBodyTableMetas {
	s.ResourceId = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetResourceType(v string) *ListTableMetasResponseBodyTableMetas {
	s.ResourceType = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetResourceUri(v string) *ListTableMetasResponseBodyTableMetas {
	s.ResourceUri = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetTableMetaId(v string) *ListTableMetasResponseBodyTableMetas {
	s.TableMetaId = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetTableName(v string) *ListTableMetasResponseBodyTableMetas {
	s.TableName = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetType(v string) *ListTableMetasResponseBodyTableMetas {
	s.Type = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) SetUrl(v string) *ListTableMetasResponseBodyTableMetas {
	s.Url = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetas) Validate() error {
	if s.Fields != nil {
		for _, item := range s.Fields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTableMetasResponseBodyTableMetasFields struct {
	// Indicates whether the field is a dimension field.
	//
	// example:
	//
	// false
	IsDimensionField *bool `json:"IsDimensionField,omitempty" xml:"IsDimensionField,omitempty"`
	// The meaning of the field.
	//
	// example:
	//
	// the gender of people
	Meaning *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// gender
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The field type.
	//
	// example:
	//
	// BIGINT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTableMetasResponseBodyTableMetasFields) String() string {
	return dara.Prettify(s)
}

func (s ListTableMetasResponseBodyTableMetasFields) GoString() string {
	return s.String()
}

func (s *ListTableMetasResponseBodyTableMetasFields) GetIsDimensionField() *bool {
	return s.IsDimensionField
}

func (s *ListTableMetasResponseBodyTableMetasFields) GetMeaning() *string {
	return s.Meaning
}

func (s *ListTableMetasResponseBodyTableMetasFields) GetName() *string {
	return s.Name
}

func (s *ListTableMetasResponseBodyTableMetasFields) GetType() *string {
	return s.Type
}

func (s *ListTableMetasResponseBodyTableMetasFields) SetIsDimensionField(v bool) *ListTableMetasResponseBodyTableMetasFields {
	s.IsDimensionField = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetasFields) SetMeaning(v string) *ListTableMetasResponseBodyTableMetasFields {
	s.Meaning = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetasFields) SetName(v string) *ListTableMetasResponseBodyTableMetasFields {
	s.Name = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetasFields) SetType(v string) *ListTableMetasResponseBodyTableMetasFields {
	s.Type = &v
	return s
}

func (s *ListTableMetasResponseBodyTableMetasFields) Validate() error {
	return dara.Validate(s)
}
