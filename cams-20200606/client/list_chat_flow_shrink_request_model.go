// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *ListChatFlowShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *ListChatFlowShrinkRequest
	GetBizExtendShrink() *string
	SetFlowCode(v string) *ListChatFlowShrinkRequest
	GetFlowCode() *string
	SetFlowStatus(v string) *ListChatFlowShrinkRequest
	GetFlowStatus() *string
	SetFlowTriggerType(v string) *ListChatFlowShrinkRequest
	GetFlowTriggerType() *string
	SetKeyword(v string) *ListChatFlowShrinkRequest
	GetKeyword() *string
	SetOwnerId(v int64) *ListChatFlowShrinkRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *ListChatFlowShrinkRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListChatFlowShrinkRequest
	GetPageSize() *int64
	SetPhoneNumber(v string) *ListChatFlowShrinkRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *ListChatFlowShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListChatFlowShrinkRequest
	GetResourceOwnerId() *int64
	SetReturnWithOnlineVersion(v bool) *ListChatFlowShrinkRequest
	GetReturnWithOnlineVersion() *bool
	SetStatus(v string) *ListChatFlowShrinkRequest
	GetStatus() *string
	SetTitle(v string) *ListChatFlowShrinkRequest
	GetTitle() *string
}

type ListChatFlowShrinkRequest struct {
	// The business tenant code. Default value: ALICOM_OPAAS.
	//
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The business extension information. Default value: an empty collection.
	//
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// flowCode
	//
	// example:
	//
	// SampleValueSampleValue
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	// The flow status.
	//
	// example:
	//
	// SampleValueSampleValue
	FlowStatus *string `json:"FlowStatus,omitempty" xml:"FlowStatus,omitempty"`
	// The flow trigger type. Valid values:
	//
	// - TriggeredManually
	//
	// - TriggeredByWhatsApp
	//
	// - TriggeredByInstagram
	//
	// - TriggeredByViber
	//
	// - TriggeredByMessenger
	//
	// example:
	//
	// TriggeredByWhatsApp
	FlowTriggerType *string `json:"FlowTriggerType,omitempty" xml:"FlowTriggerType,omitempty"`
	// The search keyword. This parameter is used for fuzzy match of flow names.
	//
	// example:
	//
	// LLM
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The phone number.
	//
	// example:
	//
	// SampleValue
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// 1
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// example:
	//
	// 1
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to return the online status. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// true
	ReturnWithOnlineVersion *bool `json:"ReturnWithOnlineVersion,omitempty" xml:"ReturnWithOnlineVersion,omitempty"`
	// The flow status. Default value: NORMAL.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The title.
	//
	// example:
	//
	// SampleValueSampleValue
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListChatFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListChatFlowShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *ListChatFlowShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *ListChatFlowShrinkRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *ListChatFlowShrinkRequest) GetFlowStatus() *string {
	return s.FlowStatus
}

func (s *ListChatFlowShrinkRequest) GetFlowTriggerType() *string {
	return s.FlowTriggerType
}

func (s *ListChatFlowShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListChatFlowShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListChatFlowShrinkRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListChatFlowShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListChatFlowShrinkRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ListChatFlowShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListChatFlowShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListChatFlowShrinkRequest) GetReturnWithOnlineVersion() *bool {
	return s.ReturnWithOnlineVersion
}

func (s *ListChatFlowShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListChatFlowShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *ListChatFlowShrinkRequest) SetBizCode(v string) *ListChatFlowShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *ListChatFlowShrinkRequest) SetBizExtendShrink(v string) *ListChatFlowShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *ListChatFlowShrinkRequest) SetFlowCode(v string) *ListChatFlowShrinkRequest {
	s.FlowCode = &v
	return s
}

func (s *ListChatFlowShrinkRequest) SetFlowStatus(v string) *ListChatFlowShrinkRequest {
	s.FlowStatus = &v
	return s
}

func (s *ListChatFlowShrinkRequest) SetFlowTriggerType(v string) *ListChatFlowShrinkRequest {
	s.FlowTriggerType = &v
	return s
}

func (s *ListChatFlowShrinkRequest) SetKeyword(v string) *ListChatFlowShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListChatFlowShrinkRequest) SetOwnerId(v int64) *ListChatFlowShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ListChatFlowShrinkRequest) SetPageNo(v int64) *ListChatFlowShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *ListChatFlowShrinkRequest) SetPageSize(v int64) *ListChatFlowShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListChatFlowShrinkRequest) SetPhoneNumber(v string) *ListChatFlowShrinkRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ListChatFlowShrinkRequest) SetResourceOwnerAccount(v string) *ListChatFlowShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListChatFlowShrinkRequest) SetResourceOwnerId(v int64) *ListChatFlowShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListChatFlowShrinkRequest) SetReturnWithOnlineVersion(v bool) *ListChatFlowShrinkRequest {
	s.ReturnWithOnlineVersion = &v
	return s
}

func (s *ListChatFlowShrinkRequest) SetStatus(v string) *ListChatFlowShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListChatFlowShrinkRequest) SetTitle(v string) *ListChatFlowShrinkRequest {
	s.Title = &v
	return s
}

func (s *ListChatFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
