// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPayOrderToMsenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryPayOrderToMsenceRequest
	GetAppId() *string
	SetCustomId(v string) *QueryPayOrderToMsenceRequest
	GetCustomId() *string
	SetMiniProgramId(v string) *QueryPayOrderToMsenceRequest
	GetMiniProgramId() *string
	SetPlatformId(v string) *QueryPayOrderToMsenceRequest
	GetPlatformId() *string
	SetTenantId(v string) *QueryPayOrderToMsenceRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *QueryPayOrderToMsenceRequest
	GetWorkspaceId() *string
}

type QueryPayOrderToMsenceRequest struct {
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// test_custom_id
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// example:
	//
	// 123321
	MiniProgramId *string `json:"MiniProgramId,omitempty" xml:"MiniProgramId,omitempty"`
	// example:
	//
	// mPaaS_Goosefish
	PlatformId *string `json:"PlatformId,omitempty" xml:"PlatformId,omitempty"`
	// example:
	//
	// CBIGWCFH
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryPayOrderToMsenceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPayOrderToMsenceRequest) GoString() string {
	return s.String()
}

func (s *QueryPayOrderToMsenceRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryPayOrderToMsenceRequest) GetCustomId() *string {
	return s.CustomId
}

func (s *QueryPayOrderToMsenceRequest) GetMiniProgramId() *string {
	return s.MiniProgramId
}

func (s *QueryPayOrderToMsenceRequest) GetPlatformId() *string {
	return s.PlatformId
}

func (s *QueryPayOrderToMsenceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryPayOrderToMsenceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryPayOrderToMsenceRequest) SetAppId(v string) *QueryPayOrderToMsenceRequest {
	s.AppId = &v
	return s
}

func (s *QueryPayOrderToMsenceRequest) SetCustomId(v string) *QueryPayOrderToMsenceRequest {
	s.CustomId = &v
	return s
}

func (s *QueryPayOrderToMsenceRequest) SetMiniProgramId(v string) *QueryPayOrderToMsenceRequest {
	s.MiniProgramId = &v
	return s
}

func (s *QueryPayOrderToMsenceRequest) SetPlatformId(v string) *QueryPayOrderToMsenceRequest {
	s.PlatformId = &v
	return s
}

func (s *QueryPayOrderToMsenceRequest) SetTenantId(v string) *QueryPayOrderToMsenceRequest {
	s.TenantId = &v
	return s
}

func (s *QueryPayOrderToMsenceRequest) SetWorkspaceId(v string) *QueryPayOrderToMsenceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryPayOrderToMsenceRequest) Validate() error {
	return dara.Validate(s)
}
