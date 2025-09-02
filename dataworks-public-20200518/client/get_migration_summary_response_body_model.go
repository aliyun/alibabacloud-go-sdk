// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrationSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMigrationSummaryResponseBodyData) *GetMigrationSummaryResponseBody
	GetData() *GetMigrationSummaryResponseBodyData
	SetRequestId(v string) *GetMigrationSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMigrationSummaryResponseBody
	GetSuccess() *bool
}

type GetMigrationSummaryResponseBody struct {
	// The details of the migration task.
	Data *GetMigrationSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. You can use the request ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 19999A96-71BA-2845-B455-ED620EF4E37F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMigrationSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetMigrationSummaryResponseBody) GetData() *GetMigrationSummaryResponseBodyData {
	return s.Data
}

func (s *GetMigrationSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMigrationSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMigrationSummaryResponseBody) SetData(v *GetMigrationSummaryResponseBodyData) *GetMigrationSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetMigrationSummaryResponseBody) SetRequestId(v string) *GetMigrationSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMigrationSummaryResponseBody) SetSuccess(v bool) *GetMigrationSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetMigrationSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMigrationSummaryResponseBodyData struct {
	// The ID of the user who created the task.
	//
	// example:
	//
	// 982293332403****
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The URL that is used to download the package of the export task. You can use the URL to download the package of the export task.
	//
	// example:
	//
	// https://shanghai-xxx-oss.oss-cn-shanghai.aliyuncs.com/pre/store/f10_bf47_b4fa7df0860f.zip?Expires=1639540903&OSSAccessKeyId=XXXXXXeF4Lv5j&Signature=qxxxxx
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 1589904000000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the task was modified.
	//
	// example:
	//
	// 1589904000000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1234
	MigrationId *int64 `json:"MigrationId,omitempty" xml:"MigrationId,omitempty"`
	// The task name.
	//
	// example:
	//
	// test_export_01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the user who managed the task.
	//
	// example:
	//
	// 982293332403****
	OpUser *string `json:"OpUser,omitempty" xml:"OpUser,omitempty"`
	// The ID of the DataWorks workspace to which the migration task belongs.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The status of the migration task. Valid values:
	//
	// 	- INIT
	//
	// 	- EDITING
	//
	// 	- IMPORTING
	//
	// 	- IMPORT_ERROR
	//
	// 	- IMPORT_SUCCESS
	//
	// 	- EXPORTING
	//
	// 	- EXPORT_ERROR
	//
	// 	- EXPORT_SUCCESS
	//
	// 	- REVOKED
	//
	// 	- PARTIAL_SUCCESS
	//
	// example:
	//
	// EXPORT_SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetMigrationSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMigrationSummaryResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetMigrationSummaryResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetMigrationSummaryResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetMigrationSummaryResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetMigrationSummaryResponseBodyData) GetMigrationId() *int64 {
	return s.MigrationId
}

func (s *GetMigrationSummaryResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetMigrationSummaryResponseBodyData) GetOpUser() *string {
	return s.OpUser
}

func (s *GetMigrationSummaryResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetMigrationSummaryResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetMigrationSummaryResponseBodyData) SetCreateUser(v string) *GetMigrationSummaryResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *GetMigrationSummaryResponseBodyData) SetDownloadUrl(v string) *GetMigrationSummaryResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GetMigrationSummaryResponseBodyData) SetGmtCreate(v int64) *GetMigrationSummaryResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetMigrationSummaryResponseBodyData) SetGmtModified(v int64) *GetMigrationSummaryResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetMigrationSummaryResponseBodyData) SetMigrationId(v int64) *GetMigrationSummaryResponseBodyData {
	s.MigrationId = &v
	return s
}

func (s *GetMigrationSummaryResponseBodyData) SetName(v string) *GetMigrationSummaryResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMigrationSummaryResponseBodyData) SetOpUser(v string) *GetMigrationSummaryResponseBodyData {
	s.OpUser = &v
	return s
}

func (s *GetMigrationSummaryResponseBodyData) SetProjectId(v int64) *GetMigrationSummaryResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetMigrationSummaryResponseBodyData) SetStatus(v string) *GetMigrationSummaryResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMigrationSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
