// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePayOrderToMsenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *CreatePayOrderToMsenceRequest
	GetAmount() *int32
	SetAppId(v string) *CreatePayOrderToMsenceRequest
	GetAppId() *string
	SetAuthToken(v string) *CreatePayOrderToMsenceRequest
	GetAuthToken() *string
	SetCustomId(v string) *CreatePayOrderToMsenceRequest
	GetCustomId() *string
	SetExtraInfo(v map[string]interface{}) *CreatePayOrderToMsenceRequest
	GetExtraInfo() map[string]interface{}
	SetMiniProgramId(v string) *CreatePayOrderToMsenceRequest
	GetMiniProgramId() *string
	SetPlatformId(v string) *CreatePayOrderToMsenceRequest
	GetPlatformId() *string
	SetTenantId(v string) *CreatePayOrderToMsenceRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *CreatePayOrderToMsenceRequest
	GetWorkspaceId() *string
}

type CreatePayOrderToMsenceRequest struct {
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// NPHTGKNR
	AuthToken *string `json:"AuthToken,omitempty" xml:"AuthToken,omitempty"`
	// example:
	//
	// test_custom_id
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// example:
	//
	// {}
	ExtraInfo map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
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
	// NPHTGKNR
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreatePayOrderToMsenceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePayOrderToMsenceRequest) GoString() string {
	return s.String()
}

func (s *CreatePayOrderToMsenceRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CreatePayOrderToMsenceRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreatePayOrderToMsenceRequest) GetAuthToken() *string {
	return s.AuthToken
}

func (s *CreatePayOrderToMsenceRequest) GetCustomId() *string {
	return s.CustomId
}

func (s *CreatePayOrderToMsenceRequest) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *CreatePayOrderToMsenceRequest) GetMiniProgramId() *string {
	return s.MiniProgramId
}

func (s *CreatePayOrderToMsenceRequest) GetPlatformId() *string {
	return s.PlatformId
}

func (s *CreatePayOrderToMsenceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePayOrderToMsenceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreatePayOrderToMsenceRequest) SetAmount(v int32) *CreatePayOrderToMsenceRequest {
	s.Amount = &v
	return s
}

func (s *CreatePayOrderToMsenceRequest) SetAppId(v string) *CreatePayOrderToMsenceRequest {
	s.AppId = &v
	return s
}

func (s *CreatePayOrderToMsenceRequest) SetAuthToken(v string) *CreatePayOrderToMsenceRequest {
	s.AuthToken = &v
	return s
}

func (s *CreatePayOrderToMsenceRequest) SetCustomId(v string) *CreatePayOrderToMsenceRequest {
	s.CustomId = &v
	return s
}

func (s *CreatePayOrderToMsenceRequest) SetExtraInfo(v map[string]interface{}) *CreatePayOrderToMsenceRequest {
	s.ExtraInfo = v
	return s
}

func (s *CreatePayOrderToMsenceRequest) SetMiniProgramId(v string) *CreatePayOrderToMsenceRequest {
	s.MiniProgramId = &v
	return s
}

func (s *CreatePayOrderToMsenceRequest) SetPlatformId(v string) *CreatePayOrderToMsenceRequest {
	s.PlatformId = &v
	return s
}

func (s *CreatePayOrderToMsenceRequest) SetTenantId(v string) *CreatePayOrderToMsenceRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePayOrderToMsenceRequest) SetWorkspaceId(v string) *CreatePayOrderToMsenceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreatePayOrderToMsenceRequest) Validate() error {
	return dara.Validate(s)
}
