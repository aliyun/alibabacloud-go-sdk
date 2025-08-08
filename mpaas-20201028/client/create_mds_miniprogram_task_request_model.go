// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMdsMiniprogramTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMdsMiniprogramTaskRequest
	GetAppId() *string
	SetGreyConfigInfo(v string) *CreateMdsMiniprogramTaskRequest
	GetGreyConfigInfo() *string
	SetGreyEndtimeData(v string) *CreateMdsMiniprogramTaskRequest
	GetGreyEndtimeData() *string
	SetGreyNum(v string) *CreateMdsMiniprogramTaskRequest
	GetGreyNum() *string
	SetId(v int64) *CreateMdsMiniprogramTaskRequest
	GetId() *int64
	SetMemo(v string) *CreateMdsMiniprogramTaskRequest
	GetMemo() *string
	SetPackageId(v int64) *CreateMdsMiniprogramTaskRequest
	GetPackageId() *int64
	SetPublishMode(v string) *CreateMdsMiniprogramTaskRequest
	GetPublishMode() *string
	SetPublishType(v int64) *CreateMdsMiniprogramTaskRequest
	GetPublishType() *int64
	SetSyncMode(v string) *CreateMdsMiniprogramTaskRequest
	GetSyncMode() *string
	SetTenantId(v string) *CreateMdsMiniprogramTaskRequest
	GetTenantId() *string
	SetWhitelistIds(v string) *CreateMdsMiniprogramTaskRequest
	GetWhitelistIds() *string
	SetWorkspaceId(v string) *CreateMdsMiniprogramTaskRequest
	GetWorkspaceId() *string
}

type CreateMdsMiniprogramTaskRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	GreyConfigInfo  *string `json:"GreyConfigInfo,omitempty" xml:"GreyConfigInfo,omitempty"`
	GreyEndtimeData *string `json:"GreyEndtimeData,omitempty" xml:"GreyEndtimeData,omitempty"`
	GreyNum         *string `json:"GreyNum,omitempty" xml:"GreyNum,omitempty"`
	// This parameter is required.
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Memo *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	// This parameter is required.
	PackageId   *int64  `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	PublishMode *string `json:"PublishMode,omitempty" xml:"PublishMode,omitempty"`
	// This parameter is required.
	PublishType  *int64  `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	SyncMode     *string `json:"SyncMode,omitempty" xml:"SyncMode,omitempty"`
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WhitelistIds *string `json:"WhitelistIds,omitempty" xml:"WhitelistIds,omitempty"`
	WorkspaceId  *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMdsMiniprogramTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsMiniprogramTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateMdsMiniprogramTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMdsMiniprogramTaskRequest) GetGreyConfigInfo() *string {
	return s.GreyConfigInfo
}

func (s *CreateMdsMiniprogramTaskRequest) GetGreyEndtimeData() *string {
	return s.GreyEndtimeData
}

func (s *CreateMdsMiniprogramTaskRequest) GetGreyNum() *string {
	return s.GreyNum
}

func (s *CreateMdsMiniprogramTaskRequest) GetId() *int64 {
	return s.Id
}

func (s *CreateMdsMiniprogramTaskRequest) GetMemo() *string {
	return s.Memo
}

func (s *CreateMdsMiniprogramTaskRequest) GetPackageId() *int64 {
	return s.PackageId
}

func (s *CreateMdsMiniprogramTaskRequest) GetPublishMode() *string {
	return s.PublishMode
}

func (s *CreateMdsMiniprogramTaskRequest) GetPublishType() *int64 {
	return s.PublishType
}

func (s *CreateMdsMiniprogramTaskRequest) GetSyncMode() *string {
	return s.SyncMode
}

func (s *CreateMdsMiniprogramTaskRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMdsMiniprogramTaskRequest) GetWhitelistIds() *string {
	return s.WhitelistIds
}

func (s *CreateMdsMiniprogramTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMdsMiniprogramTaskRequest) SetAppId(v string) *CreateMdsMiniprogramTaskRequest {
	s.AppId = &v
	return s
}

func (s *CreateMdsMiniprogramTaskRequest) SetGreyConfigInfo(v string) *CreateMdsMiniprogramTaskRequest {
	s.GreyConfigInfo = &v
	return s
}

func (s *CreateMdsMiniprogramTaskRequest) SetGreyEndtimeData(v string) *CreateMdsMiniprogramTaskRequest {
	s.GreyEndtimeData = &v
	return s
}

func (s *CreateMdsMiniprogramTaskRequest) SetGreyNum(v string) *CreateMdsMiniprogramTaskRequest {
	s.GreyNum = &v
	return s
}

func (s *CreateMdsMiniprogramTaskRequest) SetId(v int64) *CreateMdsMiniprogramTaskRequest {
	s.Id = &v
	return s
}

func (s *CreateMdsMiniprogramTaskRequest) SetMemo(v string) *CreateMdsMiniprogramTaskRequest {
	s.Memo = &v
	return s
}

func (s *CreateMdsMiniprogramTaskRequest) SetPackageId(v int64) *CreateMdsMiniprogramTaskRequest {
	s.PackageId = &v
	return s
}

func (s *CreateMdsMiniprogramTaskRequest) SetPublishMode(v string) *CreateMdsMiniprogramTaskRequest {
	s.PublishMode = &v
	return s
}

func (s *CreateMdsMiniprogramTaskRequest) SetPublishType(v int64) *CreateMdsMiniprogramTaskRequest {
	s.PublishType = &v
	return s
}

func (s *CreateMdsMiniprogramTaskRequest) SetSyncMode(v string) *CreateMdsMiniprogramTaskRequest {
	s.SyncMode = &v
	return s
}

func (s *CreateMdsMiniprogramTaskRequest) SetTenantId(v string) *CreateMdsMiniprogramTaskRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMdsMiniprogramTaskRequest) SetWhitelistIds(v string) *CreateMdsMiniprogramTaskRequest {
	s.WhitelistIds = &v
	return s
}

func (s *CreateMdsMiniprogramTaskRequest) SetWorkspaceId(v string) *CreateMdsMiniprogramTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMdsMiniprogramTaskRequest) Validate() error {
	return dara.Validate(s)
}
