// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *ListChatFlowRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *ListChatFlowRequest
	GetBizExtend() map[string]interface{}
	SetFlowTriggerType(v string) *ListChatFlowRequest
	GetFlowTriggerType() *string
	SetKeyword(v string) *ListChatFlowRequest
	GetKeyword() *string
	SetOwnerId(v int64) *ListChatFlowRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *ListChatFlowRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListChatFlowRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *ListChatFlowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListChatFlowRequest
	GetResourceOwnerId() *int64
	SetReturnWithOnlineVersion(v bool) *ListChatFlowRequest
	GetReturnWithOnlineVersion() *bool
	SetStatus(v string) *ListChatFlowRequest
	GetStatus() *string
}

type ListChatFlowRequest struct {
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
	BizExtend map[string]interface{} `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
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
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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
}

func (s ListChatFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatFlowRequest) GoString() string {
	return s.String()
}

func (s *ListChatFlowRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *ListChatFlowRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *ListChatFlowRequest) GetFlowTriggerType() *string {
	return s.FlowTriggerType
}

func (s *ListChatFlowRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListChatFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListChatFlowRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListChatFlowRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListChatFlowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListChatFlowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListChatFlowRequest) GetReturnWithOnlineVersion() *bool {
	return s.ReturnWithOnlineVersion
}

func (s *ListChatFlowRequest) GetStatus() *string {
	return s.Status
}

func (s *ListChatFlowRequest) SetBizCode(v string) *ListChatFlowRequest {
	s.BizCode = &v
	return s
}

func (s *ListChatFlowRequest) SetBizExtend(v map[string]interface{}) *ListChatFlowRequest {
	s.BizExtend = v
	return s
}

func (s *ListChatFlowRequest) SetFlowTriggerType(v string) *ListChatFlowRequest {
	s.FlowTriggerType = &v
	return s
}

func (s *ListChatFlowRequest) SetKeyword(v string) *ListChatFlowRequest {
	s.Keyword = &v
	return s
}

func (s *ListChatFlowRequest) SetOwnerId(v int64) *ListChatFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *ListChatFlowRequest) SetPageNo(v int64) *ListChatFlowRequest {
	s.PageNo = &v
	return s
}

func (s *ListChatFlowRequest) SetPageSize(v int64) *ListChatFlowRequest {
	s.PageSize = &v
	return s
}

func (s *ListChatFlowRequest) SetResourceOwnerAccount(v string) *ListChatFlowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListChatFlowRequest) SetResourceOwnerId(v int64) *ListChatFlowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListChatFlowRequest) SetReturnWithOnlineVersion(v bool) *ListChatFlowRequest {
	s.ReturnWithOnlineVersion = &v
	return s
}

func (s *ListChatFlowRequest) SetStatus(v string) *ListChatFlowRequest {
	s.Status = &v
	return s
}

func (s *ListChatFlowRequest) Validate() error {
	return dara.Validate(s)
}
