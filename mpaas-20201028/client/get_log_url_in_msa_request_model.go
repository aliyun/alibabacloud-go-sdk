// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogUrlInMsaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetLogUrlInMsaRequest
	GetAppId() *string
	SetId(v int64) *GetLogUrlInMsaRequest
	GetId() *int64
	SetTenantId(v string) *GetLogUrlInMsaRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *GetLogUrlInMsaRequest
	GetWorkspaceId() *string
}

type GetLogUrlInMsaRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetLogUrlInMsaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLogUrlInMsaRequest) GoString() string {
	return s.String()
}

func (s *GetLogUrlInMsaRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetLogUrlInMsaRequest) GetId() *int64 {
	return s.Id
}

func (s *GetLogUrlInMsaRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetLogUrlInMsaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetLogUrlInMsaRequest) SetAppId(v string) *GetLogUrlInMsaRequest {
	s.AppId = &v
	return s
}

func (s *GetLogUrlInMsaRequest) SetId(v int64) *GetLogUrlInMsaRequest {
	s.Id = &v
	return s
}

func (s *GetLogUrlInMsaRequest) SetTenantId(v string) *GetLogUrlInMsaRequest {
	s.TenantId = &v
	return s
}

func (s *GetLogUrlInMsaRequest) SetWorkspaceId(v string) *GetLogUrlInMsaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetLogUrlInMsaRequest) Validate() error {
	return dara.Validate(s)
}
