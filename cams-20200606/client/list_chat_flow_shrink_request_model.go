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
	SetResourceOwnerAccount(v string) *ListChatFlowShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListChatFlowShrinkRequest
	GetResourceOwnerId() *int64
	SetReturnWithOnlineVersion(v bool) *ListChatFlowShrinkRequest
	GetReturnWithOnlineVersion() *bool
	SetStatus(v string) *ListChatFlowShrinkRequest
	GetStatus() *string
}

type ListChatFlowShrinkRequest struct {
	// Business tenant code, default is “ALICOM_OPAAS”.
	//
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Business extension information, default is “{}”.
	//
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// Flow trigger type, enum values:
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
	// 示例值
	FlowTriggerType *string `json:"FlowTriggerType,omitempty" xml:"FlowTriggerType,omitempty"`
	// Search keyword.
	//
	// example:
	//
	// 示例值示例值
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Page size.
	//
	// example:
	//
	// 20
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Whether to return the online status
	//
	// example:
	//
	// true
	ReturnWithOnlineVersion *bool `json:"ReturnWithOnlineVersion,omitempty" xml:"ReturnWithOnlineVersion,omitempty"`
	// Flow status
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *ListChatFlowShrinkRequest) SetBizCode(v string) *ListChatFlowShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *ListChatFlowShrinkRequest) SetBizExtendShrink(v string) *ListChatFlowShrinkRequest {
	s.BizExtendShrink = &v
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

func (s *ListChatFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
