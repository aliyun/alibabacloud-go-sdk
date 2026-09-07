// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGrafanaWorkspaceAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNotes(v string) *CreateGrafanaWorkspaceAccountRequest
	GetAccountNotes() *string
	SetAccountPassword(v string) *CreateGrafanaWorkspaceAccountRequest
	GetAccountPassword() *string
	SetAliyunLang(v string) *CreateGrafanaWorkspaceAccountRequest
	GetAliyunLang() *string
	SetAliyunUid(v string) *CreateGrafanaWorkspaceAccountRequest
	GetAliyunUid() *string
	SetGrafanaWorkspaceId(v string) *CreateGrafanaWorkspaceAccountRequest
	GetGrafanaWorkspaceId() *string
	SetOrgId(v int32) *CreateGrafanaWorkspaceAccountRequest
	GetOrgId() *int32
	SetRegionId(v string) *CreateGrafanaWorkspaceAccountRequest
	GetRegionId() *string
	SetRole(v string) *CreateGrafanaWorkspaceAccountRequest
	GetRole() *string
}

type CreateGrafanaWorkspaceAccountRequest struct {
	// example:
	//
	// notes
	AccountNotes *string `json:"AccountNotes,omitempty" xml:"AccountNotes,omitempty"`
	// example:
	//
	// 123456
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1449570186405787
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// grafana-cn-06f4xyxjo01
	GrafanaWorkspaceId *string `json:"GrafanaWorkspaceId,omitempty" xml:"GrafanaWorkspaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	OrgId *int32 `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// admin
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s CreateGrafanaWorkspaceAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGrafanaWorkspaceAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateGrafanaWorkspaceAccountRequest) GetAccountNotes() *string {
	return s.AccountNotes
}

func (s *CreateGrafanaWorkspaceAccountRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateGrafanaWorkspaceAccountRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *CreateGrafanaWorkspaceAccountRequest) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *CreateGrafanaWorkspaceAccountRequest) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *CreateGrafanaWorkspaceAccountRequest) GetOrgId() *int32 {
	return s.OrgId
}

func (s *CreateGrafanaWorkspaceAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGrafanaWorkspaceAccountRequest) GetRole() *string {
	return s.Role
}

func (s *CreateGrafanaWorkspaceAccountRequest) SetAccountNotes(v string) *CreateGrafanaWorkspaceAccountRequest {
	s.AccountNotes = &v
	return s
}

func (s *CreateGrafanaWorkspaceAccountRequest) SetAccountPassword(v string) *CreateGrafanaWorkspaceAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateGrafanaWorkspaceAccountRequest) SetAliyunLang(v string) *CreateGrafanaWorkspaceAccountRequest {
	s.AliyunLang = &v
	return s
}

func (s *CreateGrafanaWorkspaceAccountRequest) SetAliyunUid(v string) *CreateGrafanaWorkspaceAccountRequest {
	s.AliyunUid = &v
	return s
}

func (s *CreateGrafanaWorkspaceAccountRequest) SetGrafanaWorkspaceId(v string) *CreateGrafanaWorkspaceAccountRequest {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *CreateGrafanaWorkspaceAccountRequest) SetOrgId(v int32) *CreateGrafanaWorkspaceAccountRequest {
	s.OrgId = &v
	return s
}

func (s *CreateGrafanaWorkspaceAccountRequest) SetRegionId(v string) *CreateGrafanaWorkspaceAccountRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGrafanaWorkspaceAccountRequest) SetRole(v string) *CreateGrafanaWorkspaceAccountRequest {
	s.Role = &v
	return s
}

func (s *CreateGrafanaWorkspaceAccountRequest) Validate() error {
	return dara.Validate(s)
}
