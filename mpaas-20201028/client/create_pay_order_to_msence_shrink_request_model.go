// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePayOrderToMsenceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *CreatePayOrderToMsenceShrinkRequest
	GetAmount() *int32
	SetAppId(v string) *CreatePayOrderToMsenceShrinkRequest
	GetAppId() *string
	SetAuthToken(v string) *CreatePayOrderToMsenceShrinkRequest
	GetAuthToken() *string
	SetCustomId(v string) *CreatePayOrderToMsenceShrinkRequest
	GetCustomId() *string
	SetExtraInfoShrink(v string) *CreatePayOrderToMsenceShrinkRequest
	GetExtraInfoShrink() *string
	SetMiniProgramId(v string) *CreatePayOrderToMsenceShrinkRequest
	GetMiniProgramId() *string
	SetPlatformId(v string) *CreatePayOrderToMsenceShrinkRequest
	GetPlatformId() *string
	SetTenantId(v string) *CreatePayOrderToMsenceShrinkRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *CreatePayOrderToMsenceShrinkRequest
	GetWorkspaceId() *string
}

type CreatePayOrderToMsenceShrinkRequest struct {
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
	ExtraInfoShrink *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
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

func (s CreatePayOrderToMsenceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePayOrderToMsenceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePayOrderToMsenceShrinkRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CreatePayOrderToMsenceShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreatePayOrderToMsenceShrinkRequest) GetAuthToken() *string {
	return s.AuthToken
}

func (s *CreatePayOrderToMsenceShrinkRequest) GetCustomId() *string {
	return s.CustomId
}

func (s *CreatePayOrderToMsenceShrinkRequest) GetExtraInfoShrink() *string {
	return s.ExtraInfoShrink
}

func (s *CreatePayOrderToMsenceShrinkRequest) GetMiniProgramId() *string {
	return s.MiniProgramId
}

func (s *CreatePayOrderToMsenceShrinkRequest) GetPlatformId() *string {
	return s.PlatformId
}

func (s *CreatePayOrderToMsenceShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePayOrderToMsenceShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreatePayOrderToMsenceShrinkRequest) SetAmount(v int32) *CreatePayOrderToMsenceShrinkRequest {
	s.Amount = &v
	return s
}

func (s *CreatePayOrderToMsenceShrinkRequest) SetAppId(v string) *CreatePayOrderToMsenceShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreatePayOrderToMsenceShrinkRequest) SetAuthToken(v string) *CreatePayOrderToMsenceShrinkRequest {
	s.AuthToken = &v
	return s
}

func (s *CreatePayOrderToMsenceShrinkRequest) SetCustomId(v string) *CreatePayOrderToMsenceShrinkRequest {
	s.CustomId = &v
	return s
}

func (s *CreatePayOrderToMsenceShrinkRequest) SetExtraInfoShrink(v string) *CreatePayOrderToMsenceShrinkRequest {
	s.ExtraInfoShrink = &v
	return s
}

func (s *CreatePayOrderToMsenceShrinkRequest) SetMiniProgramId(v string) *CreatePayOrderToMsenceShrinkRequest {
	s.MiniProgramId = &v
	return s
}

func (s *CreatePayOrderToMsenceShrinkRequest) SetPlatformId(v string) *CreatePayOrderToMsenceShrinkRequest {
	s.PlatformId = &v
	return s
}

func (s *CreatePayOrderToMsenceShrinkRequest) SetTenantId(v string) *CreatePayOrderToMsenceShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePayOrderToMsenceShrinkRequest) SetWorkspaceId(v string) *CreatePayOrderToMsenceShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreatePayOrderToMsenceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
