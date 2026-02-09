// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeHotpatchTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMcubeHotpatchTaskRequest
	GetAppId() *string
	SetGreyConfigInfo(v string) *CreateMcubeHotpatchTaskRequest
	GetGreyConfigInfo() *string
	SetGreyEndtimeData(v string) *CreateMcubeHotpatchTaskRequest
	GetGreyEndtimeData() *string
	SetGreyNum(v int64) *CreateMcubeHotpatchTaskRequest
	GetGreyNum() *int64
	SetMemo(v string) *CreateMcubeHotpatchTaskRequest
	GetMemo() *string
	SetPackageId(v int64) *CreateMcubeHotpatchTaskRequest
	GetPackageId() *int64
	SetPlatform(v string) *CreateMcubeHotpatchTaskRequest
	GetPlatform() *string
	SetPublishMode(v int64) *CreateMcubeHotpatchTaskRequest
	GetPublishMode() *int64
	SetPublishType(v int64) *CreateMcubeHotpatchTaskRequest
	GetPublishType() *int64
	SetSyncMode(v string) *CreateMcubeHotpatchTaskRequest
	GetSyncMode() *string
	SetTenantId(v string) *CreateMcubeHotpatchTaskRequest
	GetTenantId() *string
	SetWhitelistIds(v string) *CreateMcubeHotpatchTaskRequest
	GetWhitelistIds() *string
	SetWorkspaceId(v string) *CreateMcubeHotpatchTaskRequest
	GetWorkspaceId() *string
}

type CreateMcubeHotpatchTaskRequest struct {
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	GreyConfigInfo  *string `json:"GreyConfigInfo,omitempty" xml:"GreyConfigInfo,omitempty"`
	GreyEndtimeData *string `json:"GreyEndtimeData,omitempty" xml:"GreyEndtimeData,omitempty"`
	// example:
	//
	// 100
	GreyNum *int64  `json:"GreyNum,omitempty" xml:"GreyNum,omitempty"`
	Memo    *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1662218
	PackageId *int64 `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// example:
	//
	// iOS
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// 2
	PublishMode *int64 `json:"PublishMode,omitempty" xml:"PublishMode,omitempty"`
	// example:
	//
	// 3
	PublishType *int64 `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	// example:
	//
	// 0
	SyncMode *string `json:"SyncMode,omitempty" xml:"SyncMode,omitempty"`
	// example:
	//
	// ZXCXMAHQ
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 825827
	WhitelistIds *string `json:"WhitelistIds,omitempty" xml:"WhitelistIds,omitempty"`
	// example:
	//
	// dev
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMcubeHotpatchTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeHotpatchTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateMcubeHotpatchTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMcubeHotpatchTaskRequest) GetGreyConfigInfo() *string {
	return s.GreyConfigInfo
}

func (s *CreateMcubeHotpatchTaskRequest) GetGreyEndtimeData() *string {
	return s.GreyEndtimeData
}

func (s *CreateMcubeHotpatchTaskRequest) GetGreyNum() *int64 {
	return s.GreyNum
}

func (s *CreateMcubeHotpatchTaskRequest) GetMemo() *string {
	return s.Memo
}

func (s *CreateMcubeHotpatchTaskRequest) GetPackageId() *int64 {
	return s.PackageId
}

func (s *CreateMcubeHotpatchTaskRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateMcubeHotpatchTaskRequest) GetPublishMode() *int64 {
	return s.PublishMode
}

func (s *CreateMcubeHotpatchTaskRequest) GetPublishType() *int64 {
	return s.PublishType
}

func (s *CreateMcubeHotpatchTaskRequest) GetSyncMode() *string {
	return s.SyncMode
}

func (s *CreateMcubeHotpatchTaskRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMcubeHotpatchTaskRequest) GetWhitelistIds() *string {
	return s.WhitelistIds
}

func (s *CreateMcubeHotpatchTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMcubeHotpatchTaskRequest) SetAppId(v string) *CreateMcubeHotpatchTaskRequest {
	s.AppId = &v
	return s
}

func (s *CreateMcubeHotpatchTaskRequest) SetGreyConfigInfo(v string) *CreateMcubeHotpatchTaskRequest {
	s.GreyConfigInfo = &v
	return s
}

func (s *CreateMcubeHotpatchTaskRequest) SetGreyEndtimeData(v string) *CreateMcubeHotpatchTaskRequest {
	s.GreyEndtimeData = &v
	return s
}

func (s *CreateMcubeHotpatchTaskRequest) SetGreyNum(v int64) *CreateMcubeHotpatchTaskRequest {
	s.GreyNum = &v
	return s
}

func (s *CreateMcubeHotpatchTaskRequest) SetMemo(v string) *CreateMcubeHotpatchTaskRequest {
	s.Memo = &v
	return s
}

func (s *CreateMcubeHotpatchTaskRequest) SetPackageId(v int64) *CreateMcubeHotpatchTaskRequest {
	s.PackageId = &v
	return s
}

func (s *CreateMcubeHotpatchTaskRequest) SetPlatform(v string) *CreateMcubeHotpatchTaskRequest {
	s.Platform = &v
	return s
}

func (s *CreateMcubeHotpatchTaskRequest) SetPublishMode(v int64) *CreateMcubeHotpatchTaskRequest {
	s.PublishMode = &v
	return s
}

func (s *CreateMcubeHotpatchTaskRequest) SetPublishType(v int64) *CreateMcubeHotpatchTaskRequest {
	s.PublishType = &v
	return s
}

func (s *CreateMcubeHotpatchTaskRequest) SetSyncMode(v string) *CreateMcubeHotpatchTaskRequest {
	s.SyncMode = &v
	return s
}

func (s *CreateMcubeHotpatchTaskRequest) SetTenantId(v string) *CreateMcubeHotpatchTaskRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMcubeHotpatchTaskRequest) SetWhitelistIds(v string) *CreateMcubeHotpatchTaskRequest {
	s.WhitelistIds = &v
	return s
}

func (s *CreateMcubeHotpatchTaskRequest) SetWorkspaceId(v string) *CreateMcubeHotpatchTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMcubeHotpatchTaskRequest) Validate() error {
	return dara.Validate(s)
}
