// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeWhitelistForIdeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMcubeWhitelistForIdeRequest
	GetAppId() *string
	SetTenantId(v string) *CreateMcubeWhitelistForIdeRequest
	GetTenantId() *string
	SetUserId(v string) *CreateMcubeWhitelistForIdeRequest
	GetUserId() *string
	SetWhitelistValue(v string) *CreateMcubeWhitelistForIdeRequest
	GetWhitelistValue() *string
	SetWorkspaceId(v string) *CreateMcubeWhitelistForIdeRequest
	GetWorkspaceId() *string
}

type CreateMcubeWhitelistForIdeRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// This parameter is required.
	WhitelistValue *string `json:"WhitelistValue,omitempty" xml:"WhitelistValue,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMcubeWhitelistForIdeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeWhitelistForIdeRequest) GoString() string {
	return s.String()
}

func (s *CreateMcubeWhitelistForIdeRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMcubeWhitelistForIdeRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMcubeWhitelistForIdeRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateMcubeWhitelistForIdeRequest) GetWhitelistValue() *string {
	return s.WhitelistValue
}

func (s *CreateMcubeWhitelistForIdeRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMcubeWhitelistForIdeRequest) SetAppId(v string) *CreateMcubeWhitelistForIdeRequest {
	s.AppId = &v
	return s
}

func (s *CreateMcubeWhitelistForIdeRequest) SetTenantId(v string) *CreateMcubeWhitelistForIdeRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMcubeWhitelistForIdeRequest) SetUserId(v string) *CreateMcubeWhitelistForIdeRequest {
	s.UserId = &v
	return s
}

func (s *CreateMcubeWhitelistForIdeRequest) SetWhitelistValue(v string) *CreateMcubeWhitelistForIdeRequest {
	s.WhitelistValue = &v
	return s
}

func (s *CreateMcubeWhitelistForIdeRequest) SetWorkspaceId(v string) *CreateMcubeWhitelistForIdeRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMcubeWhitelistForIdeRequest) Validate() error {
	return dara.Validate(s)
}
