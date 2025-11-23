// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseExportOrderDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseExportOrderDetail(v *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) *GetDatabaseExportOrderDetailResponseBody
	GetDatabaseExportOrderDetail() *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail
	SetErrorCode(v string) *GetDatabaseExportOrderDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDatabaseExportOrderDetailResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDatabaseExportOrderDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDatabaseExportOrderDetailResponseBody
	GetSuccess() *bool
}

type GetDatabaseExportOrderDetailResponseBody struct {
	// The details of the database export ticket.
	DatabaseExportOrderDetail *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail `json:"DatabaseExportOrderDetail,omitempty" xml:"DatabaseExportOrderDetail,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// CFD8FE00-36D9-4C1B-940D-65A7B73D9066
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDatabaseExportOrderDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseExportOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatabaseExportOrderDetailResponseBody) GetDatabaseExportOrderDetail() *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail {
	return s.DatabaseExportOrderDetail
}

func (s *GetDatabaseExportOrderDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDatabaseExportOrderDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDatabaseExportOrderDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatabaseExportOrderDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDatabaseExportOrderDetailResponseBody) SetDatabaseExportOrderDetail(v *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) *GetDatabaseExportOrderDetailResponseBody {
	s.DatabaseExportOrderDetail = v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBody) SetErrorCode(v string) *GetDatabaseExportOrderDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBody) SetErrorMessage(v string) *GetDatabaseExportOrderDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBody) SetRequestId(v string) *GetDatabaseExportOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBody) SetSuccess(v bool) *GetDatabaseExportOrderDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBody) Validate() error {
	if s.DatabaseExportOrderDetail != nil {
		if err := s.DatabaseExportOrderDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail struct {
	// The business background information of the database export ticket.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The user who submitted the ticket.
	//
	// example:
	//
	// dmsuser
	Committer *string `json:"Committer,omitempty" xml:"Committer,omitempty"`
	// The ID of the user who submitted the ticket. This ID is a user ID and is not the ID of an Alibaba Cloud account.
	//
	// example:
	//
	// 12***
	CommitterId *string `json:"CommitterId,omitempty" xml:"CommitterId,omitempty"`
	// The ticket ID.
	//
	// example:
	//
	// 821****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The key information about the ticket.
	KeyInfo *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo `json:"KeyInfo,omitempty" xml:"KeyInfo,omitempty" type:"Struct"`
	// The execution logs.
	//
	// example:
	//
	// 2023-04-12 14:58:32:015 Database Dump Start.
	//
	// 2023-04-12 14:58:32:096 set server side query timeout, sql : set max_execution_time = 0
	Log *string `json:"Log,omitempty" xml:"Log,omitempty"`
	// The name that is used to search for the database.
	//
	// example:
	//
	// test@xxx.xxx.xxx.xxx:3306
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The status description of the ticket.
	//
	// example:
	//
	// ticket approval
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// The status description of the workflow.
	//
	// example:
	//
	// ticket approval
	WorkflowStatusDesc *string `json:"WorkflowStatusDesc,omitempty" xml:"WorkflowStatusDesc,omitempty"`
}

func (s GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) GoString() string {
	return s.String()
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) GetComment() *string {
	return s.Comment
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) GetCommitter() *string {
	return s.Committer
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) GetCommitterId() *string {
	return s.CommitterId
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) GetId() *int64 {
	return s.Id
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) GetKeyInfo() *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo {
	return s.KeyInfo
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) GetLog() *string {
	return s.Log
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) GetSearchName() *string {
	return s.SearchName
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) GetWorkflowStatusDesc() *string {
	return s.WorkflowStatusDesc
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) SetComment(v string) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail {
	s.Comment = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) SetCommitter(v string) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail {
	s.Committer = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) SetCommitterId(v string) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail {
	s.CommitterId = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) SetId(v int64) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail {
	s.Id = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) SetKeyInfo(v *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail {
	s.KeyInfo = v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) SetLog(v string) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail {
	s.Log = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) SetSearchName(v string) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail {
	s.SearchName = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) SetStatusDesc(v string) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail {
	s.StatusDesc = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) SetWorkflowStatusDesc(v string) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail {
	s.WorkflowStatusDesc = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetail) Validate() error {
	if s.KeyInfo != nil {
		if err := s.KeyInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo struct {
	// The time when the ticket was submitted.
	//
	// example:
	//
	// 2023-04-13 13:44:59
	AuditDate *string `json:"AuditDate,omitempty" xml:"AuditDate,omitempty"`
	// The configuration information about the ticket.
	Config *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The database ID.
	//
	// example:
	//
	// 2583****
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The URL that is used to download the export result.
	//
	// example:
	//
	// https://oss.xxx.com
	DownloadURL *string `json:"DownloadURL,omitempty" xml:"DownloadURL,omitempty"`
}

func (s GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo) GoString() string {
	return s.String()
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo) GetAuditDate() *string {
	return s.AuditDate
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo) GetConfig() *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig {
	return s.Config
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo) GetDbId() *int64 {
	return s.DbId
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo) GetDownloadURL() *string {
	return s.DownloadURL
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo) SetAuditDate(v string) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo {
	s.AuditDate = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo) SetConfig(v *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo {
	s.Config = v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo) SetDbId(v int64) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo {
	s.DbId = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo) SetDownloadURL(v string) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo {
	s.DownloadURL = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfo) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig struct {
	// The database name.
	//
	// example:
	//
	// dmstest
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The type of data that was exported. Valid values:
	//
	// 	- **DATA**: The data of the database was exported.
	//
	// 	- **STRUCT**: The schema of the database was exported.
	//
	// 	- **DATA_STRUCT**: The data and schema of the database were exported.
	//
	// example:
	//
	// DATA
	ExportContent *string `json:"ExportContent,omitempty" xml:"ExportContent,omitempty"`
	// The type of schema that was exported.
	ExportTypes *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigExportTypes `json:"ExportTypes,omitempty" xml:"ExportTypes,omitempty" type:"Struct"`
	// The extension options of the SQL script.
	SQLExtOption *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSQLExtOption `json:"SQLExtOption,omitempty" xml:"SQLExtOption,omitempty" type:"Struct"`
	// The tables that were exported from the database.
	SelectedTables *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSelectedTables `json:"SelectedTables,omitempty" xml:"SelectedTables,omitempty" type:"Struct"`
	// The format in which the database was exported. Valid values:
	//
	// 	- **SQL**
	//
	// 	- **CSV**
	//
	// 	- **XLSX**
	//
	// example:
	//
	// SQL
	TargetOption *string `json:"TargetOption,omitempty" xml:"TargetOption,omitempty"`
}

func (s GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig) GoString() string {
	return s.String()
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig) GetDbName() *string {
	return s.DbName
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig) GetExportContent() *string {
	return s.ExportContent
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig) GetExportTypes() *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigExportTypes {
	return s.ExportTypes
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig) GetSQLExtOption() *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSQLExtOption {
	return s.SQLExtOption
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig) GetSelectedTables() *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSelectedTables {
	return s.SelectedTables
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig) GetTargetOption() *string {
	return s.TargetOption
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig) SetDbName(v string) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig {
	s.DbName = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig) SetExportContent(v string) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig {
	s.ExportContent = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig) SetExportTypes(v *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigExportTypes) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig {
	s.ExportTypes = v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig) SetSQLExtOption(v *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSQLExtOption) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig {
	s.SQLExtOption = v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig) SetSelectedTables(v *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSelectedTables) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig {
	s.SelectedTables = v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig) SetTargetOption(v string) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig {
	s.TargetOption = &v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfig) Validate() error {
	if s.ExportTypes != nil {
		if err := s.ExportTypes.Validate(); err != nil {
			return err
		}
	}
	if s.SQLExtOption != nil {
		if err := s.SQLExtOption.Validate(); err != nil {
			return err
		}
	}
	if s.SelectedTables != nil {
		if err := s.SelectedTables.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigExportTypes struct {
	ExportTypes []*string `json:"ExportTypes,omitempty" xml:"ExportTypes,omitempty" type:"Repeated"`
}

func (s GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigExportTypes) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigExportTypes) GoString() string {
	return s.String()
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigExportTypes) GetExportTypes() []*string {
	return s.ExportTypes
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigExportTypes) SetExportTypes(v []*string) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigExportTypes {
	s.ExportTypes = v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigExportTypes) Validate() error {
	return dara.Validate(s)
}

type GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSQLExtOption struct {
	SQLExtOption []*string `json:"SQLExtOption,omitempty" xml:"SQLExtOption,omitempty" type:"Repeated"`
}

func (s GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSQLExtOption) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSQLExtOption) GoString() string {
	return s.String()
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSQLExtOption) GetSQLExtOption() []*string {
	return s.SQLExtOption
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSQLExtOption) SetSQLExtOption(v []*string) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSQLExtOption {
	s.SQLExtOption = v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSQLExtOption) Validate() error {
	return dara.Validate(s)
}

type GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSelectedTables struct {
	SelectedTables []*string `json:"SelectedTables,omitempty" xml:"SelectedTables,omitempty" type:"Repeated"`
}

func (s GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSelectedTables) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSelectedTables) GoString() string {
	return s.String()
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSelectedTables) GetSelectedTables() []*string {
	return s.SelectedTables
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSelectedTables) SetSelectedTables(v []*string) *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSelectedTables {
	s.SelectedTables = v
	return s
}

func (s *GetDatabaseExportOrderDetailResponseBodyDatabaseExportOrderDetailKeyInfoConfigSelectedTables) Validate() error {
	return dara.Validate(s)
}
