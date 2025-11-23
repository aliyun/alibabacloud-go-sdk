// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseExportOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateDatabaseExportOrderRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateDatabaseExportOrderRequest
	GetComment() *string
	SetParentId(v int64) *CreateDatabaseExportOrderRequest
	GetParentId() *int64
	SetPluginParam(v *CreateDatabaseExportOrderRequestPluginParam) *CreateDatabaseExportOrderRequest
	GetPluginParam() *CreateDatabaseExportOrderRequestPluginParam
	SetRelatedUserList(v []*int64) *CreateDatabaseExportOrderRequest
	GetRelatedUserList() []*int64
	SetTid(v int64) *CreateDatabaseExportOrderRequest
	GetTid() *int64
}

type CreateDatabaseExportOrderRequest struct {
	// The key of the attachment that provides more instructions for the ticket. You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to obtain the attachment key.
	//
	// example:
	//
	// order_attachment.txt
	AttachmentKey *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	// The purpose or objective of the ticket. This parameter helps reduce unnecessary communication.
	//
	// This parameter is required.
	//
	// example:
	//
	// document_test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the parent ticket.
	//
	// example:
	//
	// 877****
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The parameters of the ticket.
	//
	// This parameter is required.
	PluginParam *CreateDatabaseExportOrderRequestPluginParam `json:"PluginParam,omitempty" xml:"PluginParam,omitempty" type:"Struct"`
	// The stakeholders involved in this operation.
	RelatedUserList []*int64 `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// > To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the DMS console. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDatabaseExportOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseExportOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateDatabaseExportOrderRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateDatabaseExportOrderRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDatabaseExportOrderRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateDatabaseExportOrderRequest) GetPluginParam() *CreateDatabaseExportOrderRequestPluginParam {
	return s.PluginParam
}

func (s *CreateDatabaseExportOrderRequest) GetRelatedUserList() []*int64 {
	return s.RelatedUserList
}

func (s *CreateDatabaseExportOrderRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDatabaseExportOrderRequest) SetAttachmentKey(v string) *CreateDatabaseExportOrderRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateDatabaseExportOrderRequest) SetComment(v string) *CreateDatabaseExportOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateDatabaseExportOrderRequest) SetParentId(v int64) *CreateDatabaseExportOrderRequest {
	s.ParentId = &v
	return s
}

func (s *CreateDatabaseExportOrderRequest) SetPluginParam(v *CreateDatabaseExportOrderRequestPluginParam) *CreateDatabaseExportOrderRequest {
	s.PluginParam = v
	return s
}

func (s *CreateDatabaseExportOrderRequest) SetRelatedUserList(v []*int64) *CreateDatabaseExportOrderRequest {
	s.RelatedUserList = v
	return s
}

func (s *CreateDatabaseExportOrderRequest) SetTid(v int64) *CreateDatabaseExportOrderRequest {
	s.Tid = &v
	return s
}

func (s *CreateDatabaseExportOrderRequest) Validate() error {
	if s.PluginParam != nil {
		if err := s.PluginParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDatabaseExportOrderRequestPluginParam struct {
	// The reason for the database export.
	//
	// This parameter is required.
	//
	// example:
	//
	// document_test
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The configurations for database export.
	//
	// This parameter is required.
	Config *CreateDatabaseExportOrderRequestPluginParamConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The database ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 17****
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 137****
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the database is a logical database. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The name that is used to search for the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@xxx.xxx.xxx.xxx:3306
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s CreateDatabaseExportOrderRequestPluginParam) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseExportOrderRequestPluginParam) GoString() string {
	return s.String()
}

func (s *CreateDatabaseExportOrderRequestPluginParam) GetClassify() *string {
	return s.Classify
}

func (s *CreateDatabaseExportOrderRequestPluginParam) GetConfig() *CreateDatabaseExportOrderRequestPluginParamConfig {
	return s.Config
}

func (s *CreateDatabaseExportOrderRequestPluginParam) GetDbId() *int64 {
	return s.DbId
}

func (s *CreateDatabaseExportOrderRequestPluginParam) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *CreateDatabaseExportOrderRequestPluginParam) GetLogic() *bool {
	return s.Logic
}

func (s *CreateDatabaseExportOrderRequestPluginParam) GetSearchName() *string {
	return s.SearchName
}

func (s *CreateDatabaseExportOrderRequestPluginParam) SetClassify(v string) *CreateDatabaseExportOrderRequestPluginParam {
	s.Classify = &v
	return s
}

func (s *CreateDatabaseExportOrderRequestPluginParam) SetConfig(v *CreateDatabaseExportOrderRequestPluginParamConfig) *CreateDatabaseExportOrderRequestPluginParam {
	s.Config = v
	return s
}

func (s *CreateDatabaseExportOrderRequestPluginParam) SetDbId(v int64) *CreateDatabaseExportOrderRequestPluginParam {
	s.DbId = &v
	return s
}

func (s *CreateDatabaseExportOrderRequestPluginParam) SetInstanceId(v int64) *CreateDatabaseExportOrderRequestPluginParam {
	s.InstanceId = &v
	return s
}

func (s *CreateDatabaseExportOrderRequestPluginParam) SetLogic(v bool) *CreateDatabaseExportOrderRequestPluginParam {
	s.Logic = &v
	return s
}

func (s *CreateDatabaseExportOrderRequestPluginParam) SetSearchName(v string) *CreateDatabaseExportOrderRequestPluginParam {
	s.SearchName = &v
	return s
}

func (s *CreateDatabaseExportOrderRequestPluginParam) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDatabaseExportOrderRequestPluginParamConfig struct {
	// The export options for big data. The options are used to filter the big data to be exported. You can leave this parameter empty.
	//
	// This parameter is required.
	DataOption []*string `json:"DataOption,omitempty" xml:"DataOption,omitempty" type:"Repeated"`
	// The type of data that you want to export. Valid values:
	//
	// 	- **DATA**: The data of the database is exported.
	//
	// 	- **STRUCT**: The schema of the database is exported.
	//
	// 	- **DATA_STRUCT**: The data and schema of the database are exported.
	//
	// This parameter is required.
	//
	// example:
	//
	// DATA
	ExportContent *string `json:"ExportContent,omitempty" xml:"ExportContent,omitempty"`
	// The types of schemas that you want to export.
	ExportTypes []*string `json:"ExportTypes,omitempty" xml:"ExportTypes,omitempty" type:"Repeated"`
	// The extension options of the SQL script. You can leave this parameter empty.
	//
	// This parameter is required.
	SQLExtOption []*string `json:"SQLExtOption,omitempty" xml:"SQLExtOption,omitempty" type:"Repeated"`
	// The tables that you want to export.
	SelectedTables []*string `json:"SelectedTables,omitempty" xml:"SelectedTables,omitempty" type:"Repeated"`
	// The conditions used to filter the tables to be exported.
	Tables map[string]*string `json:"Tables,omitempty" xml:"Tables,omitempty"`
	// The format in which the database is exported. Valid values:
	//
	// 	- **SQL**
	//
	// 	- **CSV**
	//
	// 	- **XLSX**
	//
	// This parameter is required.
	//
	// example:
	//
	// SQL
	TargetOption *string `json:"TargetOption,omitempty" xml:"TargetOption,omitempty"`
}

func (s CreateDatabaseExportOrderRequestPluginParamConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseExportOrderRequestPluginParamConfig) GoString() string {
	return s.String()
}

func (s *CreateDatabaseExportOrderRequestPluginParamConfig) GetDataOption() []*string {
	return s.DataOption
}

func (s *CreateDatabaseExportOrderRequestPluginParamConfig) GetExportContent() *string {
	return s.ExportContent
}

func (s *CreateDatabaseExportOrderRequestPluginParamConfig) GetExportTypes() []*string {
	return s.ExportTypes
}

func (s *CreateDatabaseExportOrderRequestPluginParamConfig) GetSQLExtOption() []*string {
	return s.SQLExtOption
}

func (s *CreateDatabaseExportOrderRequestPluginParamConfig) GetSelectedTables() []*string {
	return s.SelectedTables
}

func (s *CreateDatabaseExportOrderRequestPluginParamConfig) GetTables() map[string]*string {
	return s.Tables
}

func (s *CreateDatabaseExportOrderRequestPluginParamConfig) GetTargetOption() *string {
	return s.TargetOption
}

func (s *CreateDatabaseExportOrderRequestPluginParamConfig) SetDataOption(v []*string) *CreateDatabaseExportOrderRequestPluginParamConfig {
	s.DataOption = v
	return s
}

func (s *CreateDatabaseExportOrderRequestPluginParamConfig) SetExportContent(v string) *CreateDatabaseExportOrderRequestPluginParamConfig {
	s.ExportContent = &v
	return s
}

func (s *CreateDatabaseExportOrderRequestPluginParamConfig) SetExportTypes(v []*string) *CreateDatabaseExportOrderRequestPluginParamConfig {
	s.ExportTypes = v
	return s
}

func (s *CreateDatabaseExportOrderRequestPluginParamConfig) SetSQLExtOption(v []*string) *CreateDatabaseExportOrderRequestPluginParamConfig {
	s.SQLExtOption = v
	return s
}

func (s *CreateDatabaseExportOrderRequestPluginParamConfig) SetSelectedTables(v []*string) *CreateDatabaseExportOrderRequestPluginParamConfig {
	s.SelectedTables = v
	return s
}

func (s *CreateDatabaseExportOrderRequestPluginParamConfig) SetTables(v map[string]*string) *CreateDatabaseExportOrderRequestPluginParamConfig {
	s.Tables = v
	return s
}

func (s *CreateDatabaseExportOrderRequestPluginParamConfig) SetTargetOption(v string) *CreateDatabaseExportOrderRequestPluginParamConfig {
	s.TargetOption = &v
	return s
}

func (s *CreateDatabaseExportOrderRequestPluginParamConfig) Validate() error {
	return dara.Validate(s)
}
