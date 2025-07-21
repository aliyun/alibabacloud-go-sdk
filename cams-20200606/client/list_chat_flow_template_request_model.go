// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatFlowTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *ListChatFlowTemplateRequest
	GetBizCode() *string
	SetKeyword(v string) *ListChatFlowTemplateRequest
	GetKeyword() *string
	SetOwnerId(v int64) *ListChatFlowTemplateRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *ListChatFlowTemplateRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListChatFlowTemplateRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *ListChatFlowTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListChatFlowTemplateRequest
	GetResourceOwnerId() *int64
	SetTriggerType(v string) *ListChatFlowTemplateRequest
	GetTriggerType() *string
}

type ListChatFlowTemplateRequest struct {
	// Business tenant code, default is “ALICOM_OPAAS”.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Search keyword.
	//
	// example:
	//
	// LLM
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Number of records per page.
	//
	// example:
	//
	// 10
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Trigger type, with the following enum values:
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
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s ListChatFlowTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatFlowTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListChatFlowTemplateRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *ListChatFlowTemplateRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListChatFlowTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListChatFlowTemplateRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListChatFlowTemplateRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListChatFlowTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListChatFlowTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListChatFlowTemplateRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListChatFlowTemplateRequest) SetBizCode(v string) *ListChatFlowTemplateRequest {
	s.BizCode = &v
	return s
}

func (s *ListChatFlowTemplateRequest) SetKeyword(v string) *ListChatFlowTemplateRequest {
	s.Keyword = &v
	return s
}

func (s *ListChatFlowTemplateRequest) SetOwnerId(v int64) *ListChatFlowTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *ListChatFlowTemplateRequest) SetPageNo(v int64) *ListChatFlowTemplateRequest {
	s.PageNo = &v
	return s
}

func (s *ListChatFlowTemplateRequest) SetPageSize(v int64) *ListChatFlowTemplateRequest {
	s.PageSize = &v
	return s
}

func (s *ListChatFlowTemplateRequest) SetResourceOwnerAccount(v string) *ListChatFlowTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListChatFlowTemplateRequest) SetResourceOwnerId(v int64) *ListChatFlowTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListChatFlowTemplateRequest) SetTriggerType(v string) *ListChatFlowTemplateRequest {
	s.TriggerType = &v
	return s
}

func (s *ListChatFlowTemplateRequest) Validate() error {
	return dara.Validate(s)
}
