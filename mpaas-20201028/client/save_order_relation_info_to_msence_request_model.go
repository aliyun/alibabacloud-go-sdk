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
	SetClientType(v string) *SaveOrderRelationInfoToMsenceRequest
	GetClientType() *string
	SetCpExtra(v string) *SaveOrderRelationInfoToMsenceRequest
	GetCpExtra() *string
	SetCustomId(v string) *SaveOrderRelationInfoToMsenceRequest
	GetCustomId() *string
	SetItemId(v string) *SaveOrderRelationInfoToMsenceRequest
	GetItemId() *string
	SetItemTitle(v string) *SaveOrderRelationInfoToMsenceRequest
	GetItemTitle() *string
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
	Amount         *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BizOrderId     *string `json:"BizOrderId,omitempty" xml:"BizOrderId,omitempty"`
	BizOrderStatus *int32  `json:"BizOrderStatus,omitempty" xml:"BizOrderStatus,omitempty"`
	ClientType     *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	CpExtra        *string `json:"CpExtra,omitempty" xml:"CpExtra,omitempty"`
	CustomId       *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	ItemId         *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemTitle      *string `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	MiniProgramId  *string `json:"MiniProgramId,omitempty" xml:"MiniProgramId,omitempty"`
	OpenUid        *string `json:"OpenUid,omitempty" xml:"OpenUid,omitempty"`
	PlatformId     *string `json:"PlatformId,omitempty" xml:"PlatformId,omitempty"`
	TenantId       *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId    *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *SaveOrderRelationInfoToMsenceRequest) GetClientType() *string {
	return s.ClientType
}

func (s *SaveOrderRelationInfoToMsenceRequest) GetCpExtra() *string {
	return s.CpExtra
}

func (s *SaveOrderRelationInfoToMsenceRequest) GetCustomId() *string {
	return s.CustomId
}

func (s *SaveOrderRelationInfoToMsenceRequest) GetItemId() *string {
	return s.ItemId
}

func (s *SaveOrderRelationInfoToMsenceRequest) GetItemTitle() *string {
	return s.ItemTitle
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

func (s *SaveOrderRelationInfoToMsenceRequest) SetClientType(v string) *SaveOrderRelationInfoToMsenceRequest {
	s.ClientType = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceRequest) SetCpExtra(v string) *SaveOrderRelationInfoToMsenceRequest {
	s.CpExtra = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceRequest) SetCustomId(v string) *SaveOrderRelationInfoToMsenceRequest {
	s.CustomId = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceRequest) SetItemId(v string) *SaveOrderRelationInfoToMsenceRequest {
	s.ItemId = &v
	return s
}

func (s *SaveOrderRelationInfoToMsenceRequest) SetItemTitle(v string) *SaveOrderRelationInfoToMsenceRequest {
	s.ItemTitle = &v
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
