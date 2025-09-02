// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgQuerySensResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCol(v string) *DsgQuerySensResultRequest
	GetCol() *string
	SetDbType(v string) *DsgQuerySensResultRequest
	GetDbType() *string
	SetLevel(v string) *DsgQuerySensResultRequest
	GetLevel() *string
	SetNodeName(v string) *DsgQuerySensResultRequest
	GetNodeName() *string
	SetOrder(v string) *DsgQuerySensResultRequest
	GetOrder() *string
	SetOrderField(v string) *DsgQuerySensResultRequest
	GetOrderField() *string
	SetPageNo(v int32) *DsgQuerySensResultRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DsgQuerySensResultRequest
	GetPageSize() *int32
	SetProjectName(v string) *DsgQuerySensResultRequest
	GetProjectName() *string
	SetSchemaName(v string) *DsgQuerySensResultRequest
	GetSchemaName() *string
	SetSensStatus(v string) *DsgQuerySensResultRequest
	GetSensStatus() *string
	SetSensitiveId(v string) *DsgQuerySensResultRequest
	GetSensitiveId() *string
	SetSensitiveName(v string) *DsgQuerySensResultRequest
	GetSensitiveName() *string
	SetTable(v string) *DsgQuerySensResultRequest
	GetTable() *string
	SetTenantId(v string) *DsgQuerySensResultRequest
	GetTenantId() *string
}

type DsgQuerySensResultRequest struct {
	// The name of the field.
	//
	// example:
	//
	// col
	Col *string `json:"Col,omitempty" xml:"Col,omitempty"`
	// The type of the database. Valid values:
	//
	// 	- **ODPS.ODPS**
	//
	// 	- **HOLO.POSTGRES**
	//
	// 	- **EMR**
	//
	// example:
	//
	// ODPS.ODPS
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The sensitivity level of the field.
	//
	// example:
	//
	// 3
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The name of a data category.
	//
	// example:
	//
	// Personal information
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The sorting method. Valid values:
	//
	// 	- DESC
	//
	// 	- ASC
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The field used for sorting.
	//
	// 	- gmt_create
	//
	// 	- gmt_modified
	//
	// example:
	//
	// gmt_create
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the workspace name.
	//
	// example:
	//
	// project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The name of the schema.
	//
	// example:
	//
	// schema
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The sensitivity status of the field.
	//
	// 	- 1: indicates sensitive.
	//
	// 	- \\-1: indicates non-sensitive.
	//
	// example:
	//
	// 1
	SensStatus *string `json:"SensStatus,omitempty" xml:"SensStatus,omitempty"`
	// The sensitive field ID.
	//
	// example:
	//
	// 10241024
	SensitiveId *string `json:"SensitiveId,omitempty" xml:"SensitiveId,omitempty"`
	// The name of the sensitive field.
	//
	// example:
	//
	// name
	SensitiveName *string `json:"SensitiveName,omitempty" xml:"SensitiveName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// table
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the DataStudio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10241024
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DsgQuerySensResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgQuerySensResultRequest) GoString() string {
	return s.String()
}

func (s *DsgQuerySensResultRequest) GetCol() *string {
	return s.Col
}

func (s *DsgQuerySensResultRequest) GetDbType() *string {
	return s.DbType
}

func (s *DsgQuerySensResultRequest) GetLevel() *string {
	return s.Level
}

func (s *DsgQuerySensResultRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *DsgQuerySensResultRequest) GetOrder() *string {
	return s.Order
}

func (s *DsgQuerySensResultRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *DsgQuerySensResultRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DsgQuerySensResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DsgQuerySensResultRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DsgQuerySensResultRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DsgQuerySensResultRequest) GetSensStatus() *string {
	return s.SensStatus
}

func (s *DsgQuerySensResultRequest) GetSensitiveId() *string {
	return s.SensitiveId
}

func (s *DsgQuerySensResultRequest) GetSensitiveName() *string {
	return s.SensitiveName
}

func (s *DsgQuerySensResultRequest) GetTable() *string {
	return s.Table
}

func (s *DsgQuerySensResultRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DsgQuerySensResultRequest) SetCol(v string) *DsgQuerySensResultRequest {
	s.Col = &v
	return s
}

func (s *DsgQuerySensResultRequest) SetDbType(v string) *DsgQuerySensResultRequest {
	s.DbType = &v
	return s
}

func (s *DsgQuerySensResultRequest) SetLevel(v string) *DsgQuerySensResultRequest {
	s.Level = &v
	return s
}

func (s *DsgQuerySensResultRequest) SetNodeName(v string) *DsgQuerySensResultRequest {
	s.NodeName = &v
	return s
}

func (s *DsgQuerySensResultRequest) SetOrder(v string) *DsgQuerySensResultRequest {
	s.Order = &v
	return s
}

func (s *DsgQuerySensResultRequest) SetOrderField(v string) *DsgQuerySensResultRequest {
	s.OrderField = &v
	return s
}

func (s *DsgQuerySensResultRequest) SetPageNo(v int32) *DsgQuerySensResultRequest {
	s.PageNo = &v
	return s
}

func (s *DsgQuerySensResultRequest) SetPageSize(v int32) *DsgQuerySensResultRequest {
	s.PageSize = &v
	return s
}

func (s *DsgQuerySensResultRequest) SetProjectName(v string) *DsgQuerySensResultRequest {
	s.ProjectName = &v
	return s
}

func (s *DsgQuerySensResultRequest) SetSchemaName(v string) *DsgQuerySensResultRequest {
	s.SchemaName = &v
	return s
}

func (s *DsgQuerySensResultRequest) SetSensStatus(v string) *DsgQuerySensResultRequest {
	s.SensStatus = &v
	return s
}

func (s *DsgQuerySensResultRequest) SetSensitiveId(v string) *DsgQuerySensResultRequest {
	s.SensitiveId = &v
	return s
}

func (s *DsgQuerySensResultRequest) SetSensitiveName(v string) *DsgQuerySensResultRequest {
	s.SensitiveName = &v
	return s
}

func (s *DsgQuerySensResultRequest) SetTable(v string) *DsgQuerySensResultRequest {
	s.Table = &v
	return s
}

func (s *DsgQuerySensResultRequest) SetTenantId(v string) *DsgQuerySensResultRequest {
	s.TenantId = &v
	return s
}

func (s *DsgQuerySensResultRequest) Validate() error {
	return dara.Validate(s)
}
