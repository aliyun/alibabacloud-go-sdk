// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveMgsApirestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SaveMgsApirestRequest
	GetAppId() *string
	SetMpaasMappcenterMgsApirestSaveJsonStr(v string) *SaveMgsApirestRequest
	GetMpaasMappcenterMgsApirestSaveJsonStr() *string
	SetTenantId(v string) *SaveMgsApirestRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *SaveMgsApirestRequest
	GetWorkspaceId() *string
}

type SaveMgsApirestRequest struct {
	AppId                                *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	MpaasMappcenterMgsApirestSaveJsonStr *string `json:"MpaasMappcenterMgsApirestSaveJsonStr,omitempty" xml:"MpaasMappcenterMgsApirestSaveJsonStr,omitempty"`
	TenantId                             *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId                          *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SaveMgsApirestRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveMgsApirestRequest) GoString() string {
	return s.String()
}

func (s *SaveMgsApirestRequest) GetAppId() *string {
	return s.AppId
}

func (s *SaveMgsApirestRequest) GetMpaasMappcenterMgsApirestSaveJsonStr() *string {
	return s.MpaasMappcenterMgsApirestSaveJsonStr
}

func (s *SaveMgsApirestRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveMgsApirestRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SaveMgsApirestRequest) SetAppId(v string) *SaveMgsApirestRequest {
	s.AppId = &v
	return s
}

func (s *SaveMgsApirestRequest) SetMpaasMappcenterMgsApirestSaveJsonStr(v string) *SaveMgsApirestRequest {
	s.MpaasMappcenterMgsApirestSaveJsonStr = &v
	return s
}

func (s *SaveMgsApirestRequest) SetTenantId(v string) *SaveMgsApirestRequest {
	s.TenantId = &v
	return s
}

func (s *SaveMgsApirestRequest) SetWorkspaceId(v string) *SaveMgsApirestRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SaveMgsApirestRequest) Validate() error {
	return dara.Validate(s)
}
