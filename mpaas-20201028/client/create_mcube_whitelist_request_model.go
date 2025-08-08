// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMcubeWhitelistRequest
	GetAppId() *string
	SetTenantId(v string) *CreateMcubeWhitelistRequest
	GetTenantId() *string
	SetWhiteListName(v string) *CreateMcubeWhitelistRequest
	GetWhiteListName() *string
	SetWhitelistType(v string) *CreateMcubeWhitelistRequest
	GetWhitelistType() *string
	SetWorkspaceId(v string) *CreateMcubeWhitelistRequest
	GetWorkspaceId() *string
}

type CreateMcubeWhitelistRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WhiteListName *string `json:"WhiteListName,omitempty" xml:"WhiteListName,omitempty"`
	// This parameter is required.
	WhitelistType *string `json:"WhitelistType,omitempty" xml:"WhitelistType,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMcubeWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeWhitelistRequest) GoString() string {
	return s.String()
}

func (s *CreateMcubeWhitelistRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMcubeWhitelistRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMcubeWhitelistRequest) GetWhiteListName() *string {
	return s.WhiteListName
}

func (s *CreateMcubeWhitelistRequest) GetWhitelistType() *string {
	return s.WhitelistType
}

func (s *CreateMcubeWhitelistRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMcubeWhitelistRequest) SetAppId(v string) *CreateMcubeWhitelistRequest {
	s.AppId = &v
	return s
}

func (s *CreateMcubeWhitelistRequest) SetTenantId(v string) *CreateMcubeWhitelistRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMcubeWhitelistRequest) SetWhiteListName(v string) *CreateMcubeWhitelistRequest {
	s.WhiteListName = &v
	return s
}

func (s *CreateMcubeWhitelistRequest) SetWhitelistType(v string) *CreateMcubeWhitelistRequest {
	s.WhitelistType = &v
	return s
}

func (s *CreateMcubeWhitelistRequest) SetWorkspaceId(v string) *CreateMcubeWhitelistRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMcubeWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
