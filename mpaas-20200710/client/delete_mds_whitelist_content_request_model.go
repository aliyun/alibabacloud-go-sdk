// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMdsWhitelistContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteMdsWhitelistContentRequest
	GetAppId() *string
	SetTenantId(v string) *DeleteMdsWhitelistContentRequest
	GetTenantId() *string
	SetWhitelistId(v string) *DeleteMdsWhitelistContentRequest
	GetWhitelistId() *string
	SetWhitelistValue(v string) *DeleteMdsWhitelistContentRequest
	GetWhitelistValue() *string
	SetWorkspaceId(v string) *DeleteMdsWhitelistContentRequest
	GetWorkspaceId() *string
}

type DeleteMdsWhitelistContentRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TenantId       *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WhitelistId    *string `json:"WhitelistId,omitempty" xml:"WhitelistId,omitempty"`
	WhitelistValue *string `json:"WhitelistValue,omitempty" xml:"WhitelistValue,omitempty"`
	WorkspaceId    *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteMdsWhitelistContentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMdsWhitelistContentRequest) GoString() string {
	return s.String()
}

func (s *DeleteMdsWhitelistContentRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMdsWhitelistContentRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteMdsWhitelistContentRequest) GetWhitelistId() *string {
	return s.WhitelistId
}

func (s *DeleteMdsWhitelistContentRequest) GetWhitelistValue() *string {
	return s.WhitelistValue
}

func (s *DeleteMdsWhitelistContentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteMdsWhitelistContentRequest) SetAppId(v string) *DeleteMdsWhitelistContentRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMdsWhitelistContentRequest) SetTenantId(v string) *DeleteMdsWhitelistContentRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteMdsWhitelistContentRequest) SetWhitelistId(v string) *DeleteMdsWhitelistContentRequest {
	s.WhitelistId = &v
	return s
}

func (s *DeleteMdsWhitelistContentRequest) SetWhitelistValue(v string) *DeleteMdsWhitelistContentRequest {
	s.WhitelistValue = &v
	return s
}

func (s *DeleteMdsWhitelistContentRequest) SetWorkspaceId(v string) *DeleteMdsWhitelistContentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteMdsWhitelistContentRequest) Validate() error {
	return dara.Validate(s)
}
