// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAppEnhanceProcessInMsaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetUserAppEnhanceProcessInMsaRequest
	GetAppId() *string
	SetId(v int64) *GetUserAppEnhanceProcessInMsaRequest
	GetId() *int64
	SetTenantId(v string) *GetUserAppEnhanceProcessInMsaRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *GetUserAppEnhanceProcessInMsaRequest
	GetWorkspaceId() *string
}

type GetUserAppEnhanceProcessInMsaRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetUserAppEnhanceProcessInMsaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppEnhanceProcessInMsaRequest) GoString() string {
	return s.String()
}

func (s *GetUserAppEnhanceProcessInMsaRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetUserAppEnhanceProcessInMsaRequest) GetId() *int64 {
	return s.Id
}

func (s *GetUserAppEnhanceProcessInMsaRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetUserAppEnhanceProcessInMsaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetUserAppEnhanceProcessInMsaRequest) SetAppId(v string) *GetUserAppEnhanceProcessInMsaRequest {
	s.AppId = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaRequest) SetId(v int64) *GetUserAppEnhanceProcessInMsaRequest {
	s.Id = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaRequest) SetTenantId(v string) *GetUserAppEnhanceProcessInMsaRequest {
	s.TenantId = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaRequest) SetWorkspaceId(v string) *GetUserAppEnhanceProcessInMsaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetUserAppEnhanceProcessInMsaRequest) Validate() error {
	return dara.Validate(s)
}
