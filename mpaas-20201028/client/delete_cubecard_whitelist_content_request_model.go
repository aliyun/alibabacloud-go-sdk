// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCubecardWhitelistContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteCubecardWhitelistContentRequest
	GetAppId() *string
	SetTenantId(v string) *DeleteCubecardWhitelistContentRequest
	GetTenantId() *string
	SetWhitelistId(v string) *DeleteCubecardWhitelistContentRequest
	GetWhitelistId() *string
	SetWhitelistValue(v string) *DeleteCubecardWhitelistContentRequest
	GetWhitelistValue() *string
	SetWorkspaceId(v string) *DeleteCubecardWhitelistContentRequest
	GetWorkspaceId() *string
}

type DeleteCubecardWhitelistContentRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TenantId       *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WhitelistId    *string `json:"WhitelistId,omitempty" xml:"WhitelistId,omitempty"`
	WhitelistValue *string `json:"WhitelistValue,omitempty" xml:"WhitelistValue,omitempty"`
	WorkspaceId    *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteCubecardWhitelistContentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCubecardWhitelistContentRequest) GoString() string {
	return s.String()
}

func (s *DeleteCubecardWhitelistContentRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteCubecardWhitelistContentRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteCubecardWhitelistContentRequest) GetWhitelistId() *string {
	return s.WhitelistId
}

func (s *DeleteCubecardWhitelistContentRequest) GetWhitelistValue() *string {
	return s.WhitelistValue
}

func (s *DeleteCubecardWhitelistContentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteCubecardWhitelistContentRequest) SetAppId(v string) *DeleteCubecardWhitelistContentRequest {
	s.AppId = &v
	return s
}

func (s *DeleteCubecardWhitelistContentRequest) SetTenantId(v string) *DeleteCubecardWhitelistContentRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteCubecardWhitelistContentRequest) SetWhitelistId(v string) *DeleteCubecardWhitelistContentRequest {
	s.WhitelistId = &v
	return s
}

func (s *DeleteCubecardWhitelistContentRequest) SetWhitelistValue(v string) *DeleteCubecardWhitelistContentRequest {
	s.WhitelistValue = &v
	return s
}

func (s *DeleteCubecardWhitelistContentRequest) SetWorkspaceId(v string) *DeleteCubecardWhitelistContentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteCubecardWhitelistContentRequest) Validate() error {
	return dara.Validate(s)
}
