// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncMiniGameOnlineStatusByAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SyncMiniGameOnlineStatusByAppRequest
	GetAppId() *string
	SetMiniProgramCode(v string) *SyncMiniGameOnlineStatusByAppRequest
	GetMiniProgramCode() *string
	SetOnlineStatus(v int32) *SyncMiniGameOnlineStatusByAppRequest
	GetOnlineStatus() *int32
	SetTenantId(v string) *SyncMiniGameOnlineStatusByAppRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *SyncMiniGameOnlineStatusByAppRequest
	GetWorkspaceId() *string
}

type SyncMiniGameOnlineStatusByAppRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	MiniProgramCode *string `json:"MiniProgramCode,omitempty" xml:"MiniProgramCode,omitempty"`
	// This parameter is required.
	OnlineStatus *int32 `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SyncMiniGameOnlineStatusByAppRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncMiniGameOnlineStatusByAppRequest) GoString() string {
	return s.String()
}

func (s *SyncMiniGameOnlineStatusByAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *SyncMiniGameOnlineStatusByAppRequest) GetMiniProgramCode() *string {
	return s.MiniProgramCode
}

func (s *SyncMiniGameOnlineStatusByAppRequest) GetOnlineStatus() *int32 {
	return s.OnlineStatus
}

func (s *SyncMiniGameOnlineStatusByAppRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SyncMiniGameOnlineStatusByAppRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SyncMiniGameOnlineStatusByAppRequest) SetAppId(v string) *SyncMiniGameOnlineStatusByAppRequest {
	s.AppId = &v
	return s
}

func (s *SyncMiniGameOnlineStatusByAppRequest) SetMiniProgramCode(v string) *SyncMiniGameOnlineStatusByAppRequest {
	s.MiniProgramCode = &v
	return s
}

func (s *SyncMiniGameOnlineStatusByAppRequest) SetOnlineStatus(v int32) *SyncMiniGameOnlineStatusByAppRequest {
	s.OnlineStatus = &v
	return s
}

func (s *SyncMiniGameOnlineStatusByAppRequest) SetTenantId(v string) *SyncMiniGameOnlineStatusByAppRequest {
	s.TenantId = &v
	return s
}

func (s *SyncMiniGameOnlineStatusByAppRequest) SetWorkspaceId(v string) *SyncMiniGameOnlineStatusByAppRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SyncMiniGameOnlineStatusByAppRequest) Validate() error {
	return dara.Validate(s)
}
