// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMgsTestreqbodyautogenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryMgsTestreqbodyautogenRequest
	GetAppId() *string
	SetFormat(v string) *QueryMgsTestreqbodyautogenRequest
	GetFormat() *string
	SetMpaasMappcenterMgsTestreqbodyautogenQueryJsonStr(v string) *QueryMgsTestreqbodyautogenRequest
	GetMpaasMappcenterMgsTestreqbodyautogenQueryJsonStr() *string
	SetTenantId(v string) *QueryMgsTestreqbodyautogenRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *QueryMgsTestreqbodyautogenRequest
	GetWorkspaceId() *string
}

type QueryMgsTestreqbodyautogenRequest struct {
	AppId                                            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Format                                           *string `json:"Format,omitempty" xml:"Format,omitempty"`
	MpaasMappcenterMgsTestreqbodyautogenQueryJsonStr *string `json:"MpaasMappcenterMgsTestreqbodyautogenQueryJsonStr,omitempty" xml:"MpaasMappcenterMgsTestreqbodyautogenQueryJsonStr,omitempty"`
	TenantId                                         *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId                                      *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMgsTestreqbodyautogenRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsTestreqbodyautogenRequest) GoString() string {
	return s.String()
}

func (s *QueryMgsTestreqbodyautogenRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMgsTestreqbodyautogenRequest) GetFormat() *string {
	return s.Format
}

func (s *QueryMgsTestreqbodyautogenRequest) GetMpaasMappcenterMgsTestreqbodyautogenQueryJsonStr() *string {
	return s.MpaasMappcenterMgsTestreqbodyautogenQueryJsonStr
}

func (s *QueryMgsTestreqbodyautogenRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMgsTestreqbodyautogenRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMgsTestreqbodyautogenRequest) SetAppId(v string) *QueryMgsTestreqbodyautogenRequest {
	s.AppId = &v
	return s
}

func (s *QueryMgsTestreqbodyautogenRequest) SetFormat(v string) *QueryMgsTestreqbodyautogenRequest {
	s.Format = &v
	return s
}

func (s *QueryMgsTestreqbodyautogenRequest) SetMpaasMappcenterMgsTestreqbodyautogenQueryJsonStr(v string) *QueryMgsTestreqbodyautogenRequest {
	s.MpaasMappcenterMgsTestreqbodyautogenQueryJsonStr = &v
	return s
}

func (s *QueryMgsTestreqbodyautogenRequest) SetTenantId(v string) *QueryMgsTestreqbodyautogenRequest {
	s.TenantId = &v
	return s
}

func (s *QueryMgsTestreqbodyautogenRequest) SetWorkspaceId(v string) *QueryMgsTestreqbodyautogenRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMgsTestreqbodyautogenRequest) Validate() error {
	return dara.Validate(s)
}
