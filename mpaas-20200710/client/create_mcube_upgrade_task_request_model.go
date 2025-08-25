// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeUpgradeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMcubeUpgradeTaskRequest
	GetAppId() *string
	SetGreyConfigInfo(v string) *CreateMcubeUpgradeTaskRequest
	GetGreyConfigInfo() *string
	SetGreyEndtimeData(v string) *CreateMcubeUpgradeTaskRequest
	GetGreyEndtimeData() *string
	SetGreyNum(v int32) *CreateMcubeUpgradeTaskRequest
	GetGreyNum() *int32
	SetHistoryForce(v int32) *CreateMcubeUpgradeTaskRequest
	GetHistoryForce() *int32
	SetMemo(v string) *CreateMcubeUpgradeTaskRequest
	GetMemo() *string
	SetPackageInfoId(v int64) *CreateMcubeUpgradeTaskRequest
	GetPackageInfoId() *int64
	SetPublishMode(v int32) *CreateMcubeUpgradeTaskRequest
	GetPublishMode() *int32
	SetPublishType(v int32) *CreateMcubeUpgradeTaskRequest
	GetPublishType() *int32
	SetTenantId(v string) *CreateMcubeUpgradeTaskRequest
	GetTenantId() *string
	SetUpgradeContent(v string) *CreateMcubeUpgradeTaskRequest
	GetUpgradeContent() *string
	SetUpgradeType(v int32) *CreateMcubeUpgradeTaskRequest
	GetUpgradeType() *int32
	SetWhitelistIds(v string) *CreateMcubeUpgradeTaskRequest
	GetWhitelistIds() *string
	SetWorkspaceId(v string) *CreateMcubeUpgradeTaskRequest
	GetWorkspaceId() *string
}

type CreateMcubeUpgradeTaskRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	GreyConfigInfo  *string `json:"GreyConfigInfo,omitempty" xml:"GreyConfigInfo,omitempty"`
	GreyEndtimeData *string `json:"GreyEndtimeData,omitempty" xml:"GreyEndtimeData,omitempty"`
	GreyNum         *int32  `json:"GreyNum,omitempty" xml:"GreyNum,omitempty"`
	HistoryForce    *int32  `json:"HistoryForce,omitempty" xml:"HistoryForce,omitempty"`
	Memo            *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	PackageInfoId   *int64  `json:"PackageInfoId,omitempty" xml:"PackageInfoId,omitempty"`
	PublishMode     *int32  `json:"PublishMode,omitempty" xml:"PublishMode,omitempty"`
	PublishType     *int32  `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	TenantId        *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UpgradeContent  *string `json:"UpgradeContent,omitempty" xml:"UpgradeContent,omitempty"`
	UpgradeType     *int32  `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
	WhitelistIds    *string `json:"WhitelistIds,omitempty" xml:"WhitelistIds,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMcubeUpgradeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeUpgradeTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateMcubeUpgradeTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMcubeUpgradeTaskRequest) GetGreyConfigInfo() *string {
	return s.GreyConfigInfo
}

func (s *CreateMcubeUpgradeTaskRequest) GetGreyEndtimeData() *string {
	return s.GreyEndtimeData
}

func (s *CreateMcubeUpgradeTaskRequest) GetGreyNum() *int32 {
	return s.GreyNum
}

func (s *CreateMcubeUpgradeTaskRequest) GetHistoryForce() *int32 {
	return s.HistoryForce
}

func (s *CreateMcubeUpgradeTaskRequest) GetMemo() *string {
	return s.Memo
}

func (s *CreateMcubeUpgradeTaskRequest) GetPackageInfoId() *int64 {
	return s.PackageInfoId
}

func (s *CreateMcubeUpgradeTaskRequest) GetPublishMode() *int32 {
	return s.PublishMode
}

func (s *CreateMcubeUpgradeTaskRequest) GetPublishType() *int32 {
	return s.PublishType
}

func (s *CreateMcubeUpgradeTaskRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMcubeUpgradeTaskRequest) GetUpgradeContent() *string {
	return s.UpgradeContent
}

func (s *CreateMcubeUpgradeTaskRequest) GetUpgradeType() *int32 {
	return s.UpgradeType
}

func (s *CreateMcubeUpgradeTaskRequest) GetWhitelistIds() *string {
	return s.WhitelistIds
}

func (s *CreateMcubeUpgradeTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMcubeUpgradeTaskRequest) SetAppId(v string) *CreateMcubeUpgradeTaskRequest {
	s.AppId = &v
	return s
}

func (s *CreateMcubeUpgradeTaskRequest) SetGreyConfigInfo(v string) *CreateMcubeUpgradeTaskRequest {
	s.GreyConfigInfo = &v
	return s
}

func (s *CreateMcubeUpgradeTaskRequest) SetGreyEndtimeData(v string) *CreateMcubeUpgradeTaskRequest {
	s.GreyEndtimeData = &v
	return s
}

func (s *CreateMcubeUpgradeTaskRequest) SetGreyNum(v int32) *CreateMcubeUpgradeTaskRequest {
	s.GreyNum = &v
	return s
}

func (s *CreateMcubeUpgradeTaskRequest) SetHistoryForce(v int32) *CreateMcubeUpgradeTaskRequest {
	s.HistoryForce = &v
	return s
}

func (s *CreateMcubeUpgradeTaskRequest) SetMemo(v string) *CreateMcubeUpgradeTaskRequest {
	s.Memo = &v
	return s
}

func (s *CreateMcubeUpgradeTaskRequest) SetPackageInfoId(v int64) *CreateMcubeUpgradeTaskRequest {
	s.PackageInfoId = &v
	return s
}

func (s *CreateMcubeUpgradeTaskRequest) SetPublishMode(v int32) *CreateMcubeUpgradeTaskRequest {
	s.PublishMode = &v
	return s
}

func (s *CreateMcubeUpgradeTaskRequest) SetPublishType(v int32) *CreateMcubeUpgradeTaskRequest {
	s.PublishType = &v
	return s
}

func (s *CreateMcubeUpgradeTaskRequest) SetTenantId(v string) *CreateMcubeUpgradeTaskRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMcubeUpgradeTaskRequest) SetUpgradeContent(v string) *CreateMcubeUpgradeTaskRequest {
	s.UpgradeContent = &v
	return s
}

func (s *CreateMcubeUpgradeTaskRequest) SetUpgradeType(v int32) *CreateMcubeUpgradeTaskRequest {
	s.UpgradeType = &v
	return s
}

func (s *CreateMcubeUpgradeTaskRequest) SetWhitelistIds(v string) *CreateMcubeUpgradeTaskRequest {
	s.WhitelistIds = &v
	return s
}

func (s *CreateMcubeUpgradeTaskRequest) SetWorkspaceId(v string) *CreateMcubeUpgradeTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMcubeUpgradeTaskRequest) Validate() error {
	return dara.Validate(s)
}
