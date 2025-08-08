// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeNebulaTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *CreateMcubeNebulaTaskRequest
	GetAppCode() *string
	SetAppId(v string) *CreateMcubeNebulaTaskRequest
	GetAppId() *string
	SetBizType(v string) *CreateMcubeNebulaTaskRequest
	GetBizType() *string
	SetCreator(v string) *CreateMcubeNebulaTaskRequest
	GetCreator() *string
	SetGmtCreate(v string) *CreateMcubeNebulaTaskRequest
	GetGmtCreate() *string
	SetGmtModified(v string) *CreateMcubeNebulaTaskRequest
	GetGmtModified() *string
	SetGmtModifiedStr(v string) *CreateMcubeNebulaTaskRequest
	GetGmtModifiedStr() *string
	SetGreyConfigInfo(v string) *CreateMcubeNebulaTaskRequest
	GetGreyConfigInfo() *string
	SetGreyEndtime(v string) *CreateMcubeNebulaTaskRequest
	GetGreyEndtime() *string
	SetGreyEndtimeData(v string) *CreateMcubeNebulaTaskRequest
	GetGreyEndtimeData() *string
	SetGreyEndtimeStr(v string) *CreateMcubeNebulaTaskRequest
	GetGreyEndtimeStr() *string
	SetGreyNum(v int32) *CreateMcubeNebulaTaskRequest
	GetGreyNum() *int32
	SetGreyUrl(v string) *CreateMcubeNebulaTaskRequest
	GetGreyUrl() *string
	SetId(v int64) *CreateMcubeNebulaTaskRequest
	GetId() *int64
	SetMemo(v string) *CreateMcubeNebulaTaskRequest
	GetMemo() *string
	SetModifier(v string) *CreateMcubeNebulaTaskRequest
	GetModifier() *string
	SetPackageId(v int64) *CreateMcubeNebulaTaskRequest
	GetPackageId() *int64
	SetPercent(v int32) *CreateMcubeNebulaTaskRequest
	GetPercent() *int32
	SetPlatform(v string) *CreateMcubeNebulaTaskRequest
	GetPlatform() *string
	SetProductId(v string) *CreateMcubeNebulaTaskRequest
	GetProductId() *string
	SetProductVersion(v string) *CreateMcubeNebulaTaskRequest
	GetProductVersion() *string
	SetPublishMode(v int32) *CreateMcubeNebulaTaskRequest
	GetPublishMode() *int32
	SetPublishType(v int32) *CreateMcubeNebulaTaskRequest
	GetPublishType() *int32
	SetReleaseVersion(v string) *CreateMcubeNebulaTaskRequest
	GetReleaseVersion() *string
	SetResIds(v string) *CreateMcubeNebulaTaskRequest
	GetResIds() *string
	SetSerialVersionUID(v int64) *CreateMcubeNebulaTaskRequest
	GetSerialVersionUID() *int64
	SetStatus(v int32) *CreateMcubeNebulaTaskRequest
	GetStatus() *int32
	SetSyncMode(v string) *CreateMcubeNebulaTaskRequest
	GetSyncMode() *string
	SetSyncResult(v string) *CreateMcubeNebulaTaskRequest
	GetSyncResult() *string
	SetTaskName(v string) *CreateMcubeNebulaTaskRequest
	GetTaskName() *string
	SetTaskStatus(v int32) *CreateMcubeNebulaTaskRequest
	GetTaskStatus() *int32
	SetTaskType(v int32) *CreateMcubeNebulaTaskRequest
	GetTaskType() *int32
	SetTaskVersion(v int64) *CreateMcubeNebulaTaskRequest
	GetTaskVersion() *int64
	SetTenantId(v string) *CreateMcubeNebulaTaskRequest
	GetTenantId() *string
	SetUpgradeNoticeNum(v int64) *CreateMcubeNebulaTaskRequest
	GetUpgradeNoticeNum() *int64
	SetUpgradeProgress(v string) *CreateMcubeNebulaTaskRequest
	GetUpgradeProgress() *string
	SetWhitelistIds(v string) *CreateMcubeNebulaTaskRequest
	GetWhitelistIds() *string
	SetWorkspaceId(v string) *CreateMcubeNebulaTaskRequest
	GetWorkspaceId() *string
}

type CreateMcubeNebulaTaskRequest struct {
	AppCode          *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BizType          *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Creator          *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	GmtCreate        *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GmtModifiedStr   *string `json:"GmtModifiedStr,omitempty" xml:"GmtModifiedStr,omitempty"`
	GreyConfigInfo   *string `json:"GreyConfigInfo,omitempty" xml:"GreyConfigInfo,omitempty"`
	GreyEndtime      *string `json:"GreyEndtime,omitempty" xml:"GreyEndtime,omitempty"`
	GreyEndtimeData  *string `json:"GreyEndtimeData,omitempty" xml:"GreyEndtimeData,omitempty"`
	GreyEndtimeStr   *string `json:"GreyEndtimeStr,omitempty" xml:"GreyEndtimeStr,omitempty"`
	GreyNum          *int32  `json:"GreyNum,omitempty" xml:"GreyNum,omitempty"`
	GreyUrl          *string `json:"GreyUrl,omitempty" xml:"GreyUrl,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Memo             *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	Modifier         *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	PackageId        *int64  `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	Percent          *int32  `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Platform         *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ProductId        *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductVersion   *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	PublishMode      *int32  `json:"PublishMode,omitempty" xml:"PublishMode,omitempty"`
	PublishType      *int32  `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	ReleaseVersion   *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	ResIds           *string `json:"ResIds,omitempty" xml:"ResIds,omitempty"`
	SerialVersionUID *int64  `json:"SerialVersionUID,omitempty" xml:"SerialVersionUID,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SyncMode         *string `json:"SyncMode,omitempty" xml:"SyncMode,omitempty"`
	SyncResult       *string `json:"SyncResult,omitempty" xml:"SyncResult,omitempty"`
	TaskName         *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskStatus       *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskType         *int32  `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskVersion      *int64  `json:"TaskVersion,omitempty" xml:"TaskVersion,omitempty"`
	TenantId         *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UpgradeNoticeNum *int64  `json:"UpgradeNoticeNum,omitempty" xml:"UpgradeNoticeNum,omitempty"`
	UpgradeProgress  *string `json:"UpgradeProgress,omitempty" xml:"UpgradeProgress,omitempty"`
	WhitelistIds     *string `json:"WhitelistIds,omitempty" xml:"WhitelistIds,omitempty"`
	WorkspaceId      *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMcubeNebulaTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeNebulaTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateMcubeNebulaTaskRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *CreateMcubeNebulaTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMcubeNebulaTaskRequest) GetBizType() *string {
	return s.BizType
}

func (s *CreateMcubeNebulaTaskRequest) GetCreator() *string {
	return s.Creator
}

func (s *CreateMcubeNebulaTaskRequest) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateMcubeNebulaTaskRequest) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateMcubeNebulaTaskRequest) GetGmtModifiedStr() *string {
	return s.GmtModifiedStr
}

func (s *CreateMcubeNebulaTaskRequest) GetGreyConfigInfo() *string {
	return s.GreyConfigInfo
}

func (s *CreateMcubeNebulaTaskRequest) GetGreyEndtime() *string {
	return s.GreyEndtime
}

func (s *CreateMcubeNebulaTaskRequest) GetGreyEndtimeData() *string {
	return s.GreyEndtimeData
}

func (s *CreateMcubeNebulaTaskRequest) GetGreyEndtimeStr() *string {
	return s.GreyEndtimeStr
}

func (s *CreateMcubeNebulaTaskRequest) GetGreyNum() *int32 {
	return s.GreyNum
}

func (s *CreateMcubeNebulaTaskRequest) GetGreyUrl() *string {
	return s.GreyUrl
}

func (s *CreateMcubeNebulaTaskRequest) GetId() *int64 {
	return s.Id
}

func (s *CreateMcubeNebulaTaskRequest) GetMemo() *string {
	return s.Memo
}

func (s *CreateMcubeNebulaTaskRequest) GetModifier() *string {
	return s.Modifier
}

func (s *CreateMcubeNebulaTaskRequest) GetPackageId() *int64 {
	return s.PackageId
}

func (s *CreateMcubeNebulaTaskRequest) GetPercent() *int32 {
	return s.Percent
}

func (s *CreateMcubeNebulaTaskRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateMcubeNebulaTaskRequest) GetProductId() *string {
	return s.ProductId
}

func (s *CreateMcubeNebulaTaskRequest) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *CreateMcubeNebulaTaskRequest) GetPublishMode() *int32 {
	return s.PublishMode
}

func (s *CreateMcubeNebulaTaskRequest) GetPublishType() *int32 {
	return s.PublishType
}

func (s *CreateMcubeNebulaTaskRequest) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *CreateMcubeNebulaTaskRequest) GetResIds() *string {
	return s.ResIds
}

func (s *CreateMcubeNebulaTaskRequest) GetSerialVersionUID() *int64 {
	return s.SerialVersionUID
}

func (s *CreateMcubeNebulaTaskRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateMcubeNebulaTaskRequest) GetSyncMode() *string {
	return s.SyncMode
}

func (s *CreateMcubeNebulaTaskRequest) GetSyncResult() *string {
	return s.SyncResult
}

func (s *CreateMcubeNebulaTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateMcubeNebulaTaskRequest) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *CreateMcubeNebulaTaskRequest) GetTaskType() *int32 {
	return s.TaskType
}

func (s *CreateMcubeNebulaTaskRequest) GetTaskVersion() *int64 {
	return s.TaskVersion
}

func (s *CreateMcubeNebulaTaskRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMcubeNebulaTaskRequest) GetUpgradeNoticeNum() *int64 {
	return s.UpgradeNoticeNum
}

func (s *CreateMcubeNebulaTaskRequest) GetUpgradeProgress() *string {
	return s.UpgradeProgress
}

func (s *CreateMcubeNebulaTaskRequest) GetWhitelistIds() *string {
	return s.WhitelistIds
}

func (s *CreateMcubeNebulaTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMcubeNebulaTaskRequest) SetAppCode(v string) *CreateMcubeNebulaTaskRequest {
	s.AppCode = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetAppId(v string) *CreateMcubeNebulaTaskRequest {
	s.AppId = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetBizType(v string) *CreateMcubeNebulaTaskRequest {
	s.BizType = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetCreator(v string) *CreateMcubeNebulaTaskRequest {
	s.Creator = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetGmtCreate(v string) *CreateMcubeNebulaTaskRequest {
	s.GmtCreate = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetGmtModified(v string) *CreateMcubeNebulaTaskRequest {
	s.GmtModified = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetGmtModifiedStr(v string) *CreateMcubeNebulaTaskRequest {
	s.GmtModifiedStr = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetGreyConfigInfo(v string) *CreateMcubeNebulaTaskRequest {
	s.GreyConfigInfo = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetGreyEndtime(v string) *CreateMcubeNebulaTaskRequest {
	s.GreyEndtime = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetGreyEndtimeData(v string) *CreateMcubeNebulaTaskRequest {
	s.GreyEndtimeData = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetGreyEndtimeStr(v string) *CreateMcubeNebulaTaskRequest {
	s.GreyEndtimeStr = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetGreyNum(v int32) *CreateMcubeNebulaTaskRequest {
	s.GreyNum = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetGreyUrl(v string) *CreateMcubeNebulaTaskRequest {
	s.GreyUrl = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetId(v int64) *CreateMcubeNebulaTaskRequest {
	s.Id = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetMemo(v string) *CreateMcubeNebulaTaskRequest {
	s.Memo = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetModifier(v string) *CreateMcubeNebulaTaskRequest {
	s.Modifier = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetPackageId(v int64) *CreateMcubeNebulaTaskRequest {
	s.PackageId = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetPercent(v int32) *CreateMcubeNebulaTaskRequest {
	s.Percent = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetPlatform(v string) *CreateMcubeNebulaTaskRequest {
	s.Platform = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetProductId(v string) *CreateMcubeNebulaTaskRequest {
	s.ProductId = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetProductVersion(v string) *CreateMcubeNebulaTaskRequest {
	s.ProductVersion = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetPublishMode(v int32) *CreateMcubeNebulaTaskRequest {
	s.PublishMode = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetPublishType(v int32) *CreateMcubeNebulaTaskRequest {
	s.PublishType = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetReleaseVersion(v string) *CreateMcubeNebulaTaskRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetResIds(v string) *CreateMcubeNebulaTaskRequest {
	s.ResIds = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetSerialVersionUID(v int64) *CreateMcubeNebulaTaskRequest {
	s.SerialVersionUID = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetStatus(v int32) *CreateMcubeNebulaTaskRequest {
	s.Status = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetSyncMode(v string) *CreateMcubeNebulaTaskRequest {
	s.SyncMode = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetSyncResult(v string) *CreateMcubeNebulaTaskRequest {
	s.SyncResult = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetTaskName(v string) *CreateMcubeNebulaTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetTaskStatus(v int32) *CreateMcubeNebulaTaskRequest {
	s.TaskStatus = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetTaskType(v int32) *CreateMcubeNebulaTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetTaskVersion(v int64) *CreateMcubeNebulaTaskRequest {
	s.TaskVersion = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetTenantId(v string) *CreateMcubeNebulaTaskRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetUpgradeNoticeNum(v int64) *CreateMcubeNebulaTaskRequest {
	s.UpgradeNoticeNum = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetUpgradeProgress(v string) *CreateMcubeNebulaTaskRequest {
	s.UpgradeProgress = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetWhitelistIds(v string) *CreateMcubeNebulaTaskRequest {
	s.WhitelistIds = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) SetWorkspaceId(v string) *CreateMcubeNebulaTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMcubeNebulaTaskRequest) Validate() error {
	return dara.Validate(s)
}
