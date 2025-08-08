// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeMiniTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMcubeMiniTaskRequest
	GetAppId() *string
	SetGreyConfigInfo(v string) *CreateMcubeMiniTaskRequest
	GetGreyConfigInfo() *string
	SetGreyEndtimeData(v string) *CreateMcubeMiniTaskRequest
	GetGreyEndtimeData() *string
	SetGreyNum(v int64) *CreateMcubeMiniTaskRequest
	GetGreyNum() *int64
	SetMemo(v string) *CreateMcubeMiniTaskRequest
	GetMemo() *string
	SetPackageId(v int64) *CreateMcubeMiniTaskRequest
	GetPackageId() *int64
	SetPublishMode(v int64) *CreateMcubeMiniTaskRequest
	GetPublishMode() *int64
	SetPublishType(v int64) *CreateMcubeMiniTaskRequest
	GetPublishType() *int64
	SetTenantId(v string) *CreateMcubeMiniTaskRequest
	GetTenantId() *string
	SetWhitelistIds(v string) *CreateMcubeMiniTaskRequest
	GetWhitelistIds() *string
	SetWorkspaceId(v string) *CreateMcubeMiniTaskRequest
	GetWorkspaceId() *string
}

type CreateMcubeMiniTaskRequest struct {
	// This parameter is required.
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	GreyConfigInfo  *string `json:"GreyConfigInfo,omitempty" xml:"GreyConfigInfo,omitempty"`
	GreyEndtimeData *string `json:"GreyEndtimeData,omitempty" xml:"GreyEndtimeData,omitempty"`
	GreyNum         *int64  `json:"GreyNum,omitempty" xml:"GreyNum,omitempty"`
	// This parameter is required.
	Memo *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	// This parameter is required.
	PackageId *int64 `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// This parameter is required.
	PublishMode *int64 `json:"PublishMode,omitempty" xml:"PublishMode,omitempty"`
	// This parameter is required.
	PublishType *int64 `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	// This parameter is required.
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WhitelistIds *string `json:"WhitelistIds,omitempty" xml:"WhitelistIds,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMcubeMiniTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeMiniTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateMcubeMiniTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMcubeMiniTaskRequest) GetGreyConfigInfo() *string {
	return s.GreyConfigInfo
}

func (s *CreateMcubeMiniTaskRequest) GetGreyEndtimeData() *string {
	return s.GreyEndtimeData
}

func (s *CreateMcubeMiniTaskRequest) GetGreyNum() *int64 {
	return s.GreyNum
}

func (s *CreateMcubeMiniTaskRequest) GetMemo() *string {
	return s.Memo
}

func (s *CreateMcubeMiniTaskRequest) GetPackageId() *int64 {
	return s.PackageId
}

func (s *CreateMcubeMiniTaskRequest) GetPublishMode() *int64 {
	return s.PublishMode
}

func (s *CreateMcubeMiniTaskRequest) GetPublishType() *int64 {
	return s.PublishType
}

func (s *CreateMcubeMiniTaskRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMcubeMiniTaskRequest) GetWhitelistIds() *string {
	return s.WhitelistIds
}

func (s *CreateMcubeMiniTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMcubeMiniTaskRequest) SetAppId(v string) *CreateMcubeMiniTaskRequest {
	s.AppId = &v
	return s
}

func (s *CreateMcubeMiniTaskRequest) SetGreyConfigInfo(v string) *CreateMcubeMiniTaskRequest {
	s.GreyConfigInfo = &v
	return s
}

func (s *CreateMcubeMiniTaskRequest) SetGreyEndtimeData(v string) *CreateMcubeMiniTaskRequest {
	s.GreyEndtimeData = &v
	return s
}

func (s *CreateMcubeMiniTaskRequest) SetGreyNum(v int64) *CreateMcubeMiniTaskRequest {
	s.GreyNum = &v
	return s
}

func (s *CreateMcubeMiniTaskRequest) SetMemo(v string) *CreateMcubeMiniTaskRequest {
	s.Memo = &v
	return s
}

func (s *CreateMcubeMiniTaskRequest) SetPackageId(v int64) *CreateMcubeMiniTaskRequest {
	s.PackageId = &v
	return s
}

func (s *CreateMcubeMiniTaskRequest) SetPublishMode(v int64) *CreateMcubeMiniTaskRequest {
	s.PublishMode = &v
	return s
}

func (s *CreateMcubeMiniTaskRequest) SetPublishType(v int64) *CreateMcubeMiniTaskRequest {
	s.PublishType = &v
	return s
}

func (s *CreateMcubeMiniTaskRequest) SetTenantId(v string) *CreateMcubeMiniTaskRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMcubeMiniTaskRequest) SetWhitelistIds(v string) *CreateMcubeMiniTaskRequest {
	s.WhitelistIds = &v
	return s
}

func (s *CreateMcubeMiniTaskRequest) SetWorkspaceId(v string) *CreateMcubeMiniTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMcubeMiniTaskRequest) Validate() error {
	return dara.Validate(s)
}
