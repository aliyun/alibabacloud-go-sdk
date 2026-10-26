// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanDelete(v bool) *GetTableMetaResponseBody
	GetCanDelete() *bool
	SetConfig(v string) *GetTableMetaResponseBody
	GetConfig() *string
	SetDescription(v string) *GetTableMetaResponseBody
	GetDescription() *string
	SetFields(v []*GetTableMetaResponseBodyFields) *GetTableMetaResponseBody
	GetFields() []*GetTableMetaResponseBodyFields
	SetGmtCreateTime(v string) *GetTableMetaResponseBody
	GetGmtCreateTime() *string
	SetGmtImportedTime(v string) *GetTableMetaResponseBody
	GetGmtImportedTime() *string
	SetGmtModifiedTime(v string) *GetTableMetaResponseBody
	GetGmtModifiedTime() *string
	SetModule(v string) *GetTableMetaResponseBody
	GetModule() *string
	SetName(v string) *GetTableMetaResponseBody
	GetName() *string
	SetRequestId(v string) *GetTableMetaResponseBody
	GetRequestId() *string
	SetResourceId(v string) *GetTableMetaResponseBody
	GetResourceId() *string
	SetResourceType(v string) *GetTableMetaResponseBody
	GetResourceType() *string
	SetResourceUri(v string) *GetTableMetaResponseBody
	GetResourceUri() *string
	SetTableMetaId(v string) *GetTableMetaResponseBody
	GetTableMetaId() *string
	SetTableName(v string) *GetTableMetaResponseBody
	GetTableName() *string
	SetType(v string) *GetTableMetaResponseBody
	GetType() *string
	SetUrl(v string) *GetTableMetaResponseBody
	GetUrl() *string
}

type GetTableMetaResponseBody struct {
	// Indicates whether the data table can be deleted. If the table is referenced elsewhere, false is returned.
	//
	// example:
	//
	// false
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
	Fields []*GetTableMetaResponseBodyFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 2021-12-15:24:33
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the field was imported. The time is in ISO 8601 format and includes the time zone offset.
	//
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtImportedTime *string `json:"GmtImportedTime,omitempty" xml:"GmtImportedTime,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2021-12-15:24:33
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The module of the data table. Valid values:
	//
	// - ABTest: experiment data table.
	//
	// - ExperimentTool: experiment tool table.
	//
	// - DataDiagnosis: data diagnostics.
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
	// The ID of the request.
	//
	// example:
	//
	// 28C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// reso-wkgo***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// Maxcompute
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The URI of the resource.
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
	// The type of the data table. Valid values:
	//
	// - MaxCompute
	//
	// - Hologres
	//
	// example:
	//
	// MaxCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the table in Data Map if the table is a MaxCompute table. Otherwise, this field is empty.
	//
	// example:
	//
	// https://dmc-xxx.com/dm/table/xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetTableMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableMetaResponseBody) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *GetTableMetaResponseBody) GetConfig() *string {
	return s.Config
}

func (s *GetTableMetaResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetTableMetaResponseBody) GetFields() []*GetTableMetaResponseBodyFields {
	return s.Fields
}

func (s *GetTableMetaResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetTableMetaResponseBody) GetGmtImportedTime() *string {
	return s.GmtImportedTime
}

func (s *GetTableMetaResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetTableMetaResponseBody) GetModule() *string {
	return s.Module
}

func (s *GetTableMetaResponseBody) GetName() *string {
	return s.Name
}

func (s *GetTableMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableMetaResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetTableMetaResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetTableMetaResponseBody) GetResourceUri() *string {
	return s.ResourceUri
}

func (s *GetTableMetaResponseBody) GetTableMetaId() *string {
	return s.TableMetaId
}

func (s *GetTableMetaResponseBody) GetTableName() *string {
	return s.TableName
}

func (s *GetTableMetaResponseBody) GetType() *string {
	return s.Type
}

func (s *GetTableMetaResponseBody) GetUrl() *string {
	return s.Url
}

func (s *GetTableMetaResponseBody) SetCanDelete(v bool) *GetTableMetaResponseBody {
	s.CanDelete = &v
	return s
}

func (s *GetTableMetaResponseBody) SetConfig(v string) *GetTableMetaResponseBody {
	s.Config = &v
	return s
}

func (s *GetTableMetaResponseBody) SetDescription(v string) *GetTableMetaResponseBody {
	s.Description = &v
	return s
}

func (s *GetTableMetaResponseBody) SetFields(v []*GetTableMetaResponseBodyFields) *GetTableMetaResponseBody {
	s.Fields = v
	return s
}

func (s *GetTableMetaResponseBody) SetGmtCreateTime(v string) *GetTableMetaResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetTableMetaResponseBody) SetGmtImportedTime(v string) *GetTableMetaResponseBody {
	s.GmtImportedTime = &v
	return s
}

func (s *GetTableMetaResponseBody) SetGmtModifiedTime(v string) *GetTableMetaResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetTableMetaResponseBody) SetModule(v string) *GetTableMetaResponseBody {
	s.Module = &v
	return s
}

func (s *GetTableMetaResponseBody) SetName(v string) *GetTableMetaResponseBody {
	s.Name = &v
	return s
}

func (s *GetTableMetaResponseBody) SetRequestId(v string) *GetTableMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableMetaResponseBody) SetResourceId(v string) *GetTableMetaResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetTableMetaResponseBody) SetResourceType(v string) *GetTableMetaResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetTableMetaResponseBody) SetResourceUri(v string) *GetTableMetaResponseBody {
	s.ResourceUri = &v
	return s
}

func (s *GetTableMetaResponseBody) SetTableMetaId(v string) *GetTableMetaResponseBody {
	s.TableMetaId = &v
	return s
}

func (s *GetTableMetaResponseBody) SetTableName(v string) *GetTableMetaResponseBody {
	s.TableName = &v
	return s
}

func (s *GetTableMetaResponseBody) SetType(v string) *GetTableMetaResponseBody {
	s.Type = &v
	return s
}

func (s *GetTableMetaResponseBody) SetUrl(v string) *GetTableMetaResponseBody {
	s.Url = &v
	return s
}

func (s *GetTableMetaResponseBody) Validate() error {
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

type GetTableMetaResponseBodyFields struct {
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
	// The type of the field.
	//
	// example:
	//
	// BIGINT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTableMetaResponseBodyFields) String() string {
	return dara.Prettify(s)
}

func (s GetTableMetaResponseBodyFields) GoString() string {
	return s.String()
}

func (s *GetTableMetaResponseBodyFields) GetIsDimensionField() *bool {
	return s.IsDimensionField
}

func (s *GetTableMetaResponseBodyFields) GetMeaning() *string {
	return s.Meaning
}

func (s *GetTableMetaResponseBodyFields) GetName() *string {
	return s.Name
}

func (s *GetTableMetaResponseBodyFields) GetType() *string {
	return s.Type
}

func (s *GetTableMetaResponseBodyFields) SetIsDimensionField(v bool) *GetTableMetaResponseBodyFields {
	s.IsDimensionField = &v
	return s
}

func (s *GetTableMetaResponseBodyFields) SetMeaning(v string) *GetTableMetaResponseBodyFields {
	s.Meaning = &v
	return s
}

func (s *GetTableMetaResponseBodyFields) SetName(v string) *GetTableMetaResponseBodyFields {
	s.Name = &v
	return s
}

func (s *GetTableMetaResponseBodyFields) SetType(v string) *GetTableMetaResponseBodyFields {
	s.Type = &v
	return s
}

func (s *GetTableMetaResponseBodyFields) Validate() error {
	return dara.Validate(s)
}
