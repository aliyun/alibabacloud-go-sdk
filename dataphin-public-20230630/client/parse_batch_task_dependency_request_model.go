// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iParseBatchTaskDependencyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *ParseBatchTaskDependencyRequest
	GetOpTenantId() *int64
	SetParseCommand(v *ParseBatchTaskDependencyRequestParseCommand) *ParseBatchTaskDependencyRequest
	GetParseCommand() *ParseBatchTaskDependencyRequestParseCommand
}

type ParseBatchTaskDependencyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	ParseCommand *ParseBatchTaskDependencyRequestParseCommand `json:"ParseCommand,omitempty" xml:"ParseCommand,omitempty" type:"Struct"`
}

func (s ParseBatchTaskDependencyRequest) String() string {
	return dara.Prettify(s)
}

func (s ParseBatchTaskDependencyRequest) GoString() string {
	return s.String()
}

func (s *ParseBatchTaskDependencyRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ParseBatchTaskDependencyRequest) GetParseCommand() *ParseBatchTaskDependencyRequestParseCommand {
	return s.ParseCommand
}

func (s *ParseBatchTaskDependencyRequest) SetOpTenantId(v int64) *ParseBatchTaskDependencyRequest {
	s.OpTenantId = &v
	return s
}

func (s *ParseBatchTaskDependencyRequest) SetParseCommand(v *ParseBatchTaskDependencyRequestParseCommand) *ParseBatchTaskDependencyRequest {
	s.ParseCommand = v
	return s
}

func (s *ParseBatchTaskDependencyRequest) Validate() error {
	return dara.Validate(s)
}

type ParseBatchTaskDependencyRequestParseCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// select 	- from t_test limit 1;
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// mysql_catalog
	DataSourceCatalog *string `json:"DataSourceCatalog,omitempty" xml:"DataSourceCatalog,omitempty"`
	// example:
	//
	// 12131111
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// erp
	DataSourceSchema      *string `json:"DataSourceSchema,omitempty" xml:"DataSourceSchema,omitempty"`
	IncludeAllInputTables *bool   `json:"IncludeAllInputTables,omitempty" xml:"IncludeAllInputTables,omitempty"`
	NeedQueryLineages     *bool   `json:"NeedQueryLineages,omitempty" xml:"NeedQueryLineages,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MaxCompute_SQL
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 131211211
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ParseBatchTaskDependencyRequestParseCommand) String() string {
	return dara.Prettify(s)
}

func (s ParseBatchTaskDependencyRequestParseCommand) GoString() string {
	return s.String()
}

func (s *ParseBatchTaskDependencyRequestParseCommand) GetCode() *string {
	return s.Code
}

func (s *ParseBatchTaskDependencyRequestParseCommand) GetDataSourceCatalog() *string {
	return s.DataSourceCatalog
}

func (s *ParseBatchTaskDependencyRequestParseCommand) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *ParseBatchTaskDependencyRequestParseCommand) GetDataSourceSchema() *string {
	return s.DataSourceSchema
}

func (s *ParseBatchTaskDependencyRequestParseCommand) GetIncludeAllInputTables() *bool {
	return s.IncludeAllInputTables
}

func (s *ParseBatchTaskDependencyRequestParseCommand) GetNeedQueryLineages() *bool {
	return s.NeedQueryLineages
}

func (s *ParseBatchTaskDependencyRequestParseCommand) GetOperatorType() *string {
	return s.OperatorType
}

func (s *ParseBatchTaskDependencyRequestParseCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ParseBatchTaskDependencyRequestParseCommand) SetCode(v string) *ParseBatchTaskDependencyRequestParseCommand {
	s.Code = &v
	return s
}

func (s *ParseBatchTaskDependencyRequestParseCommand) SetDataSourceCatalog(v string) *ParseBatchTaskDependencyRequestParseCommand {
	s.DataSourceCatalog = &v
	return s
}

func (s *ParseBatchTaskDependencyRequestParseCommand) SetDataSourceId(v int64) *ParseBatchTaskDependencyRequestParseCommand {
	s.DataSourceId = &v
	return s
}

func (s *ParseBatchTaskDependencyRequestParseCommand) SetDataSourceSchema(v string) *ParseBatchTaskDependencyRequestParseCommand {
	s.DataSourceSchema = &v
	return s
}

func (s *ParseBatchTaskDependencyRequestParseCommand) SetIncludeAllInputTables(v bool) *ParseBatchTaskDependencyRequestParseCommand {
	s.IncludeAllInputTables = &v
	return s
}

func (s *ParseBatchTaskDependencyRequestParseCommand) SetNeedQueryLineages(v bool) *ParseBatchTaskDependencyRequestParseCommand {
	s.NeedQueryLineages = &v
	return s
}

func (s *ParseBatchTaskDependencyRequestParseCommand) SetOperatorType(v string) *ParseBatchTaskDependencyRequestParseCommand {
	s.OperatorType = &v
	return s
}

func (s *ParseBatchTaskDependencyRequestParseCommand) SetProjectId(v int64) *ParseBatchTaskDependencyRequestParseCommand {
	s.ProjectId = &v
	return s
}

func (s *ParseBatchTaskDependencyRequestParseCommand) Validate() error {
	return dara.Validate(s)
}
