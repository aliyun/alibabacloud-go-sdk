// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityWatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpsertQualityWatchRequest
	GetOpTenantId() *int64
	SetUpsertCommand(v *UpsertQualityWatchRequestUpsertCommand) *UpsertQualityWatchRequest
	GetUpsertCommand() *UpsertQualityWatchRequestUpsertCommand
}

type UpsertQualityWatchRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpsertCommand *UpsertQualityWatchRequestUpsertCommand `json:"UpsertCommand,omitempty" xml:"UpsertCommand,omitempty" type:"Struct"`
}

func (s UpsertQualityWatchRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityWatchRequest) GoString() string {
	return s.String()
}

func (s *UpsertQualityWatchRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpsertQualityWatchRequest) GetUpsertCommand() *UpsertQualityWatchRequestUpsertCommand {
	return s.UpsertCommand
}

func (s *UpsertQualityWatchRequest) SetOpTenantId(v int64) *UpsertQualityWatchRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpsertQualityWatchRequest) SetUpsertCommand(v *UpsertQualityWatchRequestUpsertCommand) *UpsertQualityWatchRequest {
	s.UpsertCommand = v
	return s
}

func (s *UpsertQualityWatchRequest) Validate() error {
	if s.UpsertCommand != nil {
		if err := s.UpsertCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpsertQualityWatchRequestUpsertCommand struct {
	DataSourceInfo *UpsertQualityWatchRequestUpsertCommandDataSourceInfo `json:"DataSourceInfo,omitempty" xml:"DataSourceInfo,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Id        *int64                                           `json:"Id,omitempty" xml:"Id,omitempty"`
	IndexInfo *UpsertQualityWatchRequestUpsertCommandIndexInfo `json:"IndexInfo,omitempty" xml:"IndexInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	QualityOwner *string                                          `json:"QualityOwner,omitempty" xml:"QualityOwner,omitempty"`
	TableInfo    *UpsertQualityWatchRequestUpsertCommandTableInfo `json:"TableInfo,omitempty" xml:"TableInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// TABLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpsertQualityWatchRequestUpsertCommand) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityWatchRequestUpsertCommand) GoString() string {
	return s.String()
}

func (s *UpsertQualityWatchRequestUpsertCommand) GetDataSourceInfo() *UpsertQualityWatchRequestUpsertCommandDataSourceInfo {
	return s.DataSourceInfo
}

func (s *UpsertQualityWatchRequestUpsertCommand) GetId() *int64 {
	return s.Id
}

func (s *UpsertQualityWatchRequestUpsertCommand) GetIndexInfo() *UpsertQualityWatchRequestUpsertCommandIndexInfo {
	return s.IndexInfo
}

func (s *UpsertQualityWatchRequestUpsertCommand) GetQualityOwner() *string {
	return s.QualityOwner
}

func (s *UpsertQualityWatchRequestUpsertCommand) GetTableInfo() *UpsertQualityWatchRequestUpsertCommandTableInfo {
	return s.TableInfo
}

func (s *UpsertQualityWatchRequestUpsertCommand) GetType() *string {
	return s.Type
}

func (s *UpsertQualityWatchRequestUpsertCommand) SetDataSourceInfo(v *UpsertQualityWatchRequestUpsertCommandDataSourceInfo) *UpsertQualityWatchRequestUpsertCommand {
	s.DataSourceInfo = v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommand) SetId(v int64) *UpsertQualityWatchRequestUpsertCommand {
	s.Id = &v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommand) SetIndexInfo(v *UpsertQualityWatchRequestUpsertCommandIndexInfo) *UpsertQualityWatchRequestUpsertCommand {
	s.IndexInfo = v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommand) SetQualityOwner(v string) *UpsertQualityWatchRequestUpsertCommand {
	s.QualityOwner = &v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommand) SetTableInfo(v *UpsertQualityWatchRequestUpsertCommandTableInfo) *UpsertQualityWatchRequestUpsertCommand {
	s.TableInfo = v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommand) SetType(v string) *UpsertQualityWatchRequestUpsertCommand {
	s.Type = &v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommand) Validate() error {
	if s.DataSourceInfo != nil {
		if err := s.DataSourceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.IndexInfo != nil {
		if err := s.IndexInfo.Validate(); err != nil {
			return err
		}
	}
	if s.TableInfo != nil {
		if err := s.TableInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpsertQualityWatchRequestUpsertCommandDataSourceInfo struct {
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpsertQualityWatchRequestUpsertCommandDataSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityWatchRequestUpsertCommandDataSourceInfo) GoString() string {
	return s.String()
}

func (s *UpsertQualityWatchRequestUpsertCommandDataSourceInfo) GetId() *string {
	return s.Id
}

func (s *UpsertQualityWatchRequestUpsertCommandDataSourceInfo) SetId(v string) *UpsertQualityWatchRequestUpsertCommandDataSourceInfo {
	s.Id = &v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommandDataSourceInfo) Validate() error {
	return dara.Validate(s)
}

type UpsertQualityWatchRequestUpsertCommandIndexInfo struct {
	// example:
	//
	// 1121
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// example:
	//
	// dws_all
	CellSumLogicTableName *string `json:"CellSumLogicTableName,omitempty" xml:"CellSumLogicTableName,omitempty"`
	// example:
	//
	// AUTO
	ComputeType *string `json:"ComputeType,omitempty" xml:"ComputeType,omitempty"`
	// example:
	//
	// bigint
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// logic
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 全站汇总表
	GranularityDisplayName *string `json:"GranularityDisplayName,omitempty" xml:"GranularityDisplayName,omitempty"`
	// example:
	//
	// 18755764
	GranularityId *int64 `json:"GranularityId,omitempty" xml:"GranularityId,omitempty"`
	// example:
	//
	// 11
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test_idx_
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1121
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// INDEX
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpsertQualityWatchRequestUpsertCommandIndexInfo) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityWatchRequestUpsertCommandIndexInfo) GoString() string {
	return s.String()
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) GetCellSumLogicTableName() *string {
	return s.CellSumLogicTableName
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) GetComputeType() *string {
	return s.ComputeType
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) GetDateType() *string {
	return s.DateType
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) GetDescription() *string {
	return s.Description
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) GetGranularityDisplayName() *string {
	return s.GranularityDisplayName
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) GetGranularityId() *int64 {
	return s.GranularityId
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) GetId() *string {
	return s.Id
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) GetName() *string {
	return s.Name
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) GetType() *string {
	return s.Type
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) SetBizUnitId(v int64) *UpsertQualityWatchRequestUpsertCommandIndexInfo {
	s.BizUnitId = &v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) SetCellSumLogicTableName(v string) *UpsertQualityWatchRequestUpsertCommandIndexInfo {
	s.CellSumLogicTableName = &v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) SetComputeType(v string) *UpsertQualityWatchRequestUpsertCommandIndexInfo {
	s.ComputeType = &v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) SetDateType(v string) *UpsertQualityWatchRequestUpsertCommandIndexInfo {
	s.DateType = &v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) SetDescription(v string) *UpsertQualityWatchRequestUpsertCommandIndexInfo {
	s.Description = &v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) SetDisplayName(v string) *UpsertQualityWatchRequestUpsertCommandIndexInfo {
	s.DisplayName = &v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) SetGranularityDisplayName(v string) *UpsertQualityWatchRequestUpsertCommandIndexInfo {
	s.GranularityDisplayName = &v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) SetGranularityId(v int64) *UpsertQualityWatchRequestUpsertCommandIndexInfo {
	s.GranularityId = &v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) SetId(v string) *UpsertQualityWatchRequestUpsertCommandIndexInfo {
	s.Id = &v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) SetName(v string) *UpsertQualityWatchRequestUpsertCommandIndexInfo {
	s.Name = &v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) SetProjectId(v int64) *UpsertQualityWatchRequestUpsertCommandIndexInfo {
	s.ProjectId = &v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) SetType(v string) *UpsertQualityWatchRequestUpsertCommandIndexInfo {
	s.Type = &v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommandIndexInfo) Validate() error {
	return dara.Validate(s)
}

type UpsertQualityWatchRequestUpsertCommandTableInfo struct {
	// example:
	//
	// test
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpsertQualityWatchRequestUpsertCommandTableInfo) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityWatchRequestUpsertCommandTableInfo) GoString() string {
	return s.String()
}

func (s *UpsertQualityWatchRequestUpsertCommandTableInfo) GetId() *string {
	return s.Id
}

func (s *UpsertQualityWatchRequestUpsertCommandTableInfo) SetId(v string) *UpsertQualityWatchRequestUpsertCommandTableInfo {
	s.Id = &v
	return s
}

func (s *UpsertQualityWatchRequestUpsertCommandTableInfo) Validate() error {
	return dara.Validate(s)
}
