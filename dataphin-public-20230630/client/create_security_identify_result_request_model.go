// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityIdentifyResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateSecurityIdentifyResultRequestCreateCommand) *CreateSecurityIdentifyResultRequest
	GetCreateCommand() *CreateSecurityIdentifyResultRequestCreateCommand
	SetOpTenantId(v int64) *CreateSecurityIdentifyResultRequest
	GetOpTenantId() *int64
}

type CreateSecurityIdentifyResultRequest struct {
	// The create request.
	//
	// This parameter is required.
	CreateCommand *CreateSecurityIdentifyResultRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateSecurityIdentifyResultRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityIdentifyResultRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityIdentifyResultRequest) GetCreateCommand() *CreateSecurityIdentifyResultRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateSecurityIdentifyResultRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateSecurityIdentifyResultRequest) SetCreateCommand(v *CreateSecurityIdentifyResultRequestCreateCommand) *CreateSecurityIdentifyResultRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateSecurityIdentifyResultRequest) SetOpTenantId(v int64) *CreateSecurityIdentifyResultRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateSecurityIdentifyResultRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSecurityIdentifyResultRequestCreateCommand struct {
	// The classification ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ClassifyId *int64 `json:"ClassifyId,omitempty" xml:"ClassifyId,omitempty"`
	// The deduplication strategy. Valid values:
	//
	// - COVER_ALL: overwrites all online tags.
	//
	// - COVER_UNLOCKED: overwrites only unlocked online tags.
	//
	// This parameter is required.
	//
	// example:
	//
	// COVER_ALL
	ConflictStrategy *string `json:"ConflictStrategy,omitempty" xml:"ConflictStrategy,omitempty"`
	// The datasource environment identifier. This parameter is required only for datasource tables.
	//
	// example:
	//
	// DEV
	DatasourceEnv *string `json:"DatasourceEnv,omitempty" xml:"DatasourceEnv,omitempty"`
	// The datasource name. This parameter is required only for datasource tables.
	//
	// example:
	//
	// test
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// Specifies whether to enable the result. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The field name.
	//
	// This parameter is required.
	//
	// example:
	//
	// col1
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// Specifies whether the table is a datasource table. Default value: false (treated as a Dataphin table).
	//
	// example:
	//
	// true
	IsDatasourceTable *bool `json:"IsDatasourceTable,omitempty" xml:"IsDatasourceTable,omitempty"`
	// The table catalog. For a datasource table, specify the database or schema name. For a Dataphin physical table, specify the project name in English. For a Dataphin logical table, specify the business unit name in English.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TableCatalog *string `json:"TableCatalog,omitempty" xml:"TableCatalog,omitempty"`
	// The table name.
	//
	// This parameter is required.
	//
	// example:
	//
	// t_test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s CreateSecurityIdentifyResultRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityIdentifyResultRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) GetClassifyId() *int64 {
	return s.ClassifyId
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) GetConflictStrategy() *string {
	return s.ConflictStrategy
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) GetDatasourceEnv() *string {
	return s.DatasourceEnv
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) GetEnable() *bool {
	return s.Enable
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) GetFieldName() *string {
	return s.FieldName
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) GetIsDatasourceTable() *bool {
	return s.IsDatasourceTable
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) GetTableCatalog() *string {
	return s.TableCatalog
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) GetTableName() *string {
	return s.TableName
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) SetClassifyId(v int64) *CreateSecurityIdentifyResultRequestCreateCommand {
	s.ClassifyId = &v
	return s
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) SetConflictStrategy(v string) *CreateSecurityIdentifyResultRequestCreateCommand {
	s.ConflictStrategy = &v
	return s
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) SetDatasourceEnv(v string) *CreateSecurityIdentifyResultRequestCreateCommand {
	s.DatasourceEnv = &v
	return s
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) SetDatasourceName(v string) *CreateSecurityIdentifyResultRequestCreateCommand {
	s.DatasourceName = &v
	return s
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) SetEnable(v bool) *CreateSecurityIdentifyResultRequestCreateCommand {
	s.Enable = &v
	return s
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) SetFieldName(v string) *CreateSecurityIdentifyResultRequestCreateCommand {
	s.FieldName = &v
	return s
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) SetIsDatasourceTable(v bool) *CreateSecurityIdentifyResultRequestCreateCommand {
	s.IsDatasourceTable = &v
	return s
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) SetTableCatalog(v string) *CreateSecurityIdentifyResultRequestCreateCommand {
	s.TableCatalog = &v
	return s
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) SetTableName(v string) *CreateSecurityIdentifyResultRequestCreateCommand {
	s.TableName = &v
	return s
}

func (s *CreateSecurityIdentifyResultRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}
