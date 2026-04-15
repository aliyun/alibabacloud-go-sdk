// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveOrderRelationInfoToMsenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *SaveOrderRelationInfoToMsenceRequest
	GetAmount() *int32
	SetAppId(v string) *SaveOrderRelationInfoToMsenceRequest
	GetAppId() *string
	SetBizOrderId(v string) *SaveOrderRelationInfoToMsenceRequest
	GetBizOrderId() *string
	SetBizOrderStatus(v int32) *SaveOrderRelationInfoToMsenceRequest
	GetBizOrderStatus() *int32
	SetCustomId(v string) *SaveOrderRelationInfoToMsenceRequest
	GetCustomId() *string
	SetMiniProgramId(v string) *SaveOrderRelationInfoToMsenceRequest
	GetMiniProgramId() *string
	SetOpenUid(v string) *SaveOrderRelationInfoToMsenceRequest
	GetOpenUid() *string
	SetPlatformId(v string) *SaveOrderRelationInfoToMsenceRequest
	GetPlatformId() *string
	SetTenantId(v string) *SaveOrderRelationInfoToMsenceRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *SaveOrderRelationInfoToMsenceRequest
	GetWorkspaceId() *string
}

type SaveOrderRelationInfoToMsenceRequest struct {
	// example:
	//
	// 100
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 3929520
	BizOrderId *string `json:"BizOrderId,omitempty" xml:"BizOrderId,omitempty"`
	// example:
	//
	// 2
	BizOrderStatus *int32 `json:"BizOrderStatus,omitempty" xml:"BizOrderStatus,omitempty"`
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
	// 123456
	OpenUid *string `json:"OpenUid,omitempty" xml:"OpenUid,omitempty"`
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

func (s SaveOrderRelationInfoToMsenceRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveOrderRelationInfoToMsenceRequest) GoString() string {
	return s.String()
}

func (s *SaveOrderRelationInfoToMsenceRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *SaveOrderRelationInfoToMsenceRequest) GetAppId() *string {
	return s.AppId
}

func (s *SaveOrderRelationInfoToMsenceRequest) GetBizOrderId() *string {
	return s.BizOrderId
}

func (s *SaveOrderRelationInfoToMsenceRequest) GetBizOrderStatus() *int32 {
	return s.BizOrderStatus
}

func (s *SaveOrderRelationInfoToMsenceRequest) GetCustomId() *string {
	return s.CustomId
}

func (s *SaveOrderRelationInfoToMsenceRequest) GetMiniProgramId() *string {
	return s.MiniProgramId
}

func (s *SaveOrderRelationInfoToMsenceRequest) GetOpenUid() *string {
	return s.OpenUid
}

func (s *SaveOrderRelationInfoToMsenceRequest) GetPlatformId() *string {
	return s.PlatformId
}

func (s *SaveOrderRelationInfoToMsenceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveOrderRelationInfoToMsenceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SaveOrderRelationInfoToMsenceRequest) SetAmount(v int32) *SaveOrderRelationInfoToMsenceRequest {
	s.Amount = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceRequest) SetAppId(v string) *SaveOrderRelationInfoToMsenceRequest {
	s.AppId = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceRequest) SetBizOrderId(v string) *SaveOrderRelationInfoToMsenceRequest {
	s.BizOrderId = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceRequest) SetBizOrderStatus(v int32) *SaveOrderRelationInfoToMsenceRequest {
	s.BizOrderStatus = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceRequest) SetCustomId(v string) *SaveOrderRelationInfoToMsenceRequest {
	s.CustomId = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceRequest) SetMiniProgramId(v string) *SaveOrderRelationInfoToMsenceRequest {
	s.MiniProgramId = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceRequest) SetOpenUid(v string) *SaveOrderRelationInfoToMsenceRequest {
	s.OpenUid = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceRequest) SetPlatformId(v string) *SaveOrderRelationInfoToMsenceRequest {
	s.PlatformId = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceRequest) SetTenantId(v string) *SaveOrderRelationInfoToMsenceRequest {
	s.TenantId = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceRequest) SetWorkspaceId(v string) *SaveOrderRelationInfoToMsenceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceRequest) Validate() error {
	return dara.Validate(s)
}
