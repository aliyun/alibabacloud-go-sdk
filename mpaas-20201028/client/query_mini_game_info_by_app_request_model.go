// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMiniGameInfoByAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryMiniGameInfoByAppRequest
	GetAppId() *string
	SetMiniProgramCode(v string) *QueryMiniGameInfoByAppRequest
	GetMiniProgramCode() *string
	SetTenantId(v string) *QueryMiniGameInfoByAppRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *QueryMiniGameInfoByAppRequest
	GetWorkspaceId() *string
}

type QueryMiniGameInfoByAppRequest struct {
	// This parameter is required.
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	MiniProgramCode *string `json:"MiniProgramCode,omitempty" xml:"MiniProgramCode,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMiniGameInfoByAppRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMiniGameInfoByAppRequest) GoString() string {
	return s.String()
}

func (s *QueryMiniGameInfoByAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMiniGameInfoByAppRequest) GetMiniProgramCode() *string {
	return s.MiniProgramCode
}

func (s *QueryMiniGameInfoByAppRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMiniGameInfoByAppRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMiniGameInfoByAppRequest) SetAppId(v string) *QueryMiniGameInfoByAppRequest {
	s.AppId = &v
	return s
}

func (s *QueryMiniGameInfoByAppRequest) SetMiniProgramCode(v string) *QueryMiniGameInfoByAppRequest {
	s.MiniProgramCode = &v
	return s
}

func (s *QueryMiniGameInfoByAppRequest) SetTenantId(v string) *QueryMiniGameInfoByAppRequest {
	s.TenantId = &v
	return s
}

func (s *QueryMiniGameInfoByAppRequest) SetWorkspaceId(v string) *QueryMiniGameInfoByAppRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMiniGameInfoByAppRequest) Validate() error {
	return dara.Validate(s)
}
