// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityArchiveTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpsertQualityArchiveTableRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpsertQualityArchiveTableRequest
	GetOpUserId() *string
	SetUpsertCommand(v *UpsertQualityArchiveTableRequestUpsertCommand) *UpsertQualityArchiveTableRequest
	GetUpsertCommand() *UpsertQualityArchiveTableRequestUpsertCommand
}

type UpsertQualityArchiveTableRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The upsert command.
	//
	// This parameter is required.
	UpsertCommand *UpsertQualityArchiveTableRequestUpsertCommand `json:"UpsertCommand,omitempty" xml:"UpsertCommand,omitempty" type:"Struct"`
}

func (s UpsertQualityArchiveTableRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityArchiveTableRequest) GoString() string {
	return s.String()
}

func (s *UpsertQualityArchiveTableRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpsertQualityArchiveTableRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpsertQualityArchiveTableRequest) GetUpsertCommand() *UpsertQualityArchiveTableRequestUpsertCommand {
	return s.UpsertCommand
}

func (s *UpsertQualityArchiveTableRequest) SetOpTenantId(v int64) *UpsertQualityArchiveTableRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpsertQualityArchiveTableRequest) SetOpUserId(v string) *UpsertQualityArchiveTableRequest {
	s.OpUserId = &v
	return s
}

func (s *UpsertQualityArchiveTableRequest) SetUpsertCommand(v *UpsertQualityArchiveTableRequestUpsertCommand) *UpsertQualityArchiveTableRequest {
	s.UpsertCommand = v
	return s
}

func (s *UpsertQualityArchiveTableRequest) Validate() error {
	if s.UpsertCommand != nil {
		if err := s.UpsertCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpsertQualityArchiveTableRequestUpsertCommand struct {
	// The mode for adding the archived table. Valid values:
	//
	// - CREATE_NEW_TABLE: creates a new table.
	//
	// - BIND_EXIST_TABLE: binds an existing table.
	//
	// example:
	//
	// CREATE_NEW_TABLE
	AddMode *string `json:"AddMode,omitempty" xml:"AddMode,omitempty"`
	// The ID of the archived table. If this parameter is specified, the operation runs in update mode, and you cannot specify AddMode or NewTableNamePrefix. If this parameter is not specified, the operation runs in create mode.
	//
	// example:
	//
	// 88012
	ArchiveTableId *int64 `json:"ArchiveTableId,omitempty" xml:"ArchiveTableId,omitempty"`
	// The name of the existing table. This parameter is required when AddMode is set to BIND_EXIST_TABLE. For Dataphin tables, use the format "project_name.table_name" (for example, dataphin03.ads_region_order_summary). For datasource tables, use the format "database/schema.table_name" (for example, order_db.order_exception_data). The table must belong to the same project or datasource as the monitored object, and the table schema must contain system fields with the dataphin_quality_ prefix.
	//
	// example:
	//
	// dataphin03.ads_region_order_summary
	ExistTableName *string `json:"ExistTableName,omitempty" xml:"ExistTableName,omitempty"`
	// The lifecycle of the table, in days. The value must be a positive integer. If this parameter is not specified, no lifecycle is set. This parameter is valid only when creating a new table or in edit pattern, and only when the table belongs to MaxCompute, Hadoop series, or Hive. This parameter cannot be specified when AddMode is set to BIND_EXIST_TABLE.
	//
	// example:
	//
	// 30
	Lifecycle *int32 `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// The maximum number of archived rows. A positive integer specifies the limit on the number of archived rows. The console provides options of 10,000, 100,000, and 500,000. A value of -1 indicates full archiving. Default value: 10000. This parameter is supported only for MaxCompute, Hadoop series, or Hive.
	//
	// example:
	//
	// 100000
	MaxArchiveCount *int64 `json:"MaxArchiveCount,omitempty" xml:"MaxArchiveCount,omitempty"`
	// The table name prefix for the new archived table. This parameter is required when AddMode is set to CREATE_NEW_TABLE. The system automatically appends the _exception_data suffix. For example, if you specify vip_user_tips112, the actual table name is vip_user_tips112_exception_data.
	//
	// example:
	//
	// vip_user_tips112
	NewTableNamePrefix *string `json:"NewTableNamePrefix,omitempty" xml:"NewTableNamePrefix,omitempty"`
	// Specifies whether to set the archived table as the active table. Only the value true is supported. After the table is set as active, the previously active table under the same monitored object is automatically deactivated (only one active table is allowed at a time). If you set this parameter to false, an InvalidParameter error is returned. If this parameter is not specified, the default value true is used. If this parameter is left empty, the active status remains unchanged.
	//
	// example:
	//
	// true
	SetActive *bool `json:"SetActive,omitempty" xml:"SetActive,omitempty"`
	// The ID of the monitored object to which the archived table belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
}

func (s UpsertQualityArchiveTableRequestUpsertCommand) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityArchiveTableRequestUpsertCommand) GoString() string {
	return s.String()
}

func (s *UpsertQualityArchiveTableRequestUpsertCommand) GetAddMode() *string {
	return s.AddMode
}

func (s *UpsertQualityArchiveTableRequestUpsertCommand) GetArchiveTableId() *int64 {
	return s.ArchiveTableId
}

func (s *UpsertQualityArchiveTableRequestUpsertCommand) GetExistTableName() *string {
	return s.ExistTableName
}

func (s *UpsertQualityArchiveTableRequestUpsertCommand) GetLifecycle() *int32 {
	return s.Lifecycle
}

func (s *UpsertQualityArchiveTableRequestUpsertCommand) GetMaxArchiveCount() *int64 {
	return s.MaxArchiveCount
}

func (s *UpsertQualityArchiveTableRequestUpsertCommand) GetNewTableNamePrefix() *string {
	return s.NewTableNamePrefix
}

func (s *UpsertQualityArchiveTableRequestUpsertCommand) GetSetActive() *bool {
	return s.SetActive
}

func (s *UpsertQualityArchiveTableRequestUpsertCommand) GetWatchId() *int64 {
	return s.WatchId
}

func (s *UpsertQualityArchiveTableRequestUpsertCommand) SetAddMode(v string) *UpsertQualityArchiveTableRequestUpsertCommand {
	s.AddMode = &v
	return s
}

func (s *UpsertQualityArchiveTableRequestUpsertCommand) SetArchiveTableId(v int64) *UpsertQualityArchiveTableRequestUpsertCommand {
	s.ArchiveTableId = &v
	return s
}

func (s *UpsertQualityArchiveTableRequestUpsertCommand) SetExistTableName(v string) *UpsertQualityArchiveTableRequestUpsertCommand {
	s.ExistTableName = &v
	return s
}

func (s *UpsertQualityArchiveTableRequestUpsertCommand) SetLifecycle(v int32) *UpsertQualityArchiveTableRequestUpsertCommand {
	s.Lifecycle = &v
	return s
}

func (s *UpsertQualityArchiveTableRequestUpsertCommand) SetMaxArchiveCount(v int64) *UpsertQualityArchiveTableRequestUpsertCommand {
	s.MaxArchiveCount = &v
	return s
}

func (s *UpsertQualityArchiveTableRequestUpsertCommand) SetNewTableNamePrefix(v string) *UpsertQualityArchiveTableRequestUpsertCommand {
	s.NewTableNamePrefix = &v
	return s
}

func (s *UpsertQualityArchiveTableRequestUpsertCommand) SetSetActive(v bool) *UpsertQualityArchiveTableRequestUpsertCommand {
	s.SetActive = &v
	return s
}

func (s *UpsertQualityArchiveTableRequestUpsertCommand) SetWatchId(v int64) *UpsertQualityArchiveTableRequestUpsertCommand {
	s.WatchId = &v
	return s
}

func (s *UpsertQualityArchiveTableRequestUpsertCommand) Validate() error {
	return dara.Validate(s)
}
