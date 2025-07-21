// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatFlowByImportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *CreateChatFlowByImportShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *CreateChatFlowByImportShrinkRequest
	GetBizExtendShrink() *string
	SetFlowViewModel(v string) *CreateChatFlowByImportShrinkRequest
	GetFlowViewModel() *string
	SetOwnerId(v int64) *CreateChatFlowByImportShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *CreateChatFlowByImportShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *CreateChatFlowByImportShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateChatFlowByImportShrinkRequest
	GetResourceOwnerId() *int64
	SetTitle(v string) *CreateChatFlowByImportShrinkRequest
	GetTitle() *string
}

type CreateChatFlowByImportShrinkRequest struct {
	// Business tenant code, default is “ALICOM_OPAAS”.
	//
	// example:
	//
	// 示例值示例值
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Business extension information, default is “{}”.
	//
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// Imported flow DSL data
	//
	// example:
	//
	// 示例值示例值
	FlowViewModel *string `json:"FlowViewModel,omitempty" xml:"FlowViewModel,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Flow remarks
	//
	// example:
	//
	// 示例值示例值
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Flow title
	//
	// example:
	//
	// 示例值示例值示例值
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateChatFlowByImportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatFlowByImportShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateChatFlowByImportShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *CreateChatFlowByImportShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *CreateChatFlowByImportShrinkRequest) GetFlowViewModel() *string {
	return s.FlowViewModel
}

func (s *CreateChatFlowByImportShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateChatFlowByImportShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateChatFlowByImportShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateChatFlowByImportShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateChatFlowByImportShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateChatFlowByImportShrinkRequest) SetBizCode(v string) *CreateChatFlowByImportShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *CreateChatFlowByImportShrinkRequest) SetBizExtendShrink(v string) *CreateChatFlowByImportShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *CreateChatFlowByImportShrinkRequest) SetFlowViewModel(v string) *CreateChatFlowByImportShrinkRequest {
	s.FlowViewModel = &v
	return s
}

func (s *CreateChatFlowByImportShrinkRequest) SetOwnerId(v int64) *CreateChatFlowByImportShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateChatFlowByImportShrinkRequest) SetRemark(v string) *CreateChatFlowByImportShrinkRequest {
	s.Remark = &v
	return s
}

func (s *CreateChatFlowByImportShrinkRequest) SetResourceOwnerAccount(v string) *CreateChatFlowByImportShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateChatFlowByImportShrinkRequest) SetResourceOwnerId(v int64) *CreateChatFlowByImportShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateChatFlowByImportShrinkRequest) SetTitle(v string) *CreateChatFlowByImportShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateChatFlowByImportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
