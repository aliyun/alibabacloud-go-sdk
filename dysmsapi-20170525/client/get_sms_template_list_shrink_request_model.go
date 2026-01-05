// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsTemplateListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditStatus(v int64) *GetSmsTemplateListShrinkRequest
	GetAuditStatus() *int64
	SetOwnerId(v int64) *GetSmsTemplateListShrinkRequest
	GetOwnerId() *int64
	SetPageIndex(v int64) *GetSmsTemplateListShrinkRequest
	GetPageIndex() *int64
	SetPageSize(v int64) *GetSmsTemplateListShrinkRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *GetSmsTemplateListShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetSmsTemplateListShrinkRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *GetSmsTemplateListShrinkRequest
	GetSignName() *string
	SetTemplateCode(v string) *GetSmsTemplateListShrinkRequest
	GetTemplateCode() *string
	SetTemplateName(v string) *GetSmsTemplateListShrinkRequest
	GetTemplateName() *string
	SetTemplateTag(v string) *GetSmsTemplateListShrinkRequest
	GetTemplateTag() *string
	SetTemplateType(v int64) *GetSmsTemplateListShrinkRequest
	GetTemplateType() *int64
	SetUsableStateListShrink(v string) *GetSmsTemplateListShrinkRequest
	GetUsableStateListShrink() *string
}

type GetSmsTemplateListShrinkRequest struct {
	// 模板审核状态
	//
	// example:
	//
	// 1
	AuditStatus *int64 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 页码, 默认1
	//
	// example:
	//
	// 1
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// 每页数量，默认10
	//
	// example:
	//
	// 10
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 签名
	//
	// example:
	//
	// 示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// 模板code
	//
	// example:
	//
	// SMS_1688168
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板名称
	//
	// example:
	//
	// 示例值
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 模板标签
	//
	// example:
	//
	// 示例值示例值
	TemplateTag *string `json:"TemplateTag,omitempty" xml:"TemplateTag,omitempty"`
	// 模板类型
	//
	// example:
	//
	// 0
	TemplateType *int64 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// 模板状态
	UsableStateListShrink *string `json:"UsableStateList,omitempty" xml:"UsableStateList,omitempty"`
}

func (s GetSmsTemplateListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSmsTemplateListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateListShrinkRequest) GetAuditStatus() *int64 {
	return s.AuditStatus
}

func (s *GetSmsTemplateListShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetSmsTemplateListShrinkRequest) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *GetSmsTemplateListShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetSmsTemplateListShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetSmsTemplateListShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetSmsTemplateListShrinkRequest) GetSignName() *string {
	return s.SignName
}

func (s *GetSmsTemplateListShrinkRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetSmsTemplateListShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetSmsTemplateListShrinkRequest) GetTemplateTag() *string {
	return s.TemplateTag
}

func (s *GetSmsTemplateListShrinkRequest) GetTemplateType() *int64 {
	return s.TemplateType
}

func (s *GetSmsTemplateListShrinkRequest) GetUsableStateListShrink() *string {
	return s.UsableStateListShrink
}

func (s *GetSmsTemplateListShrinkRequest) SetAuditStatus(v int64) *GetSmsTemplateListShrinkRequest {
	s.AuditStatus = &v
	return s
}

func (s *GetSmsTemplateListShrinkRequest) SetOwnerId(v int64) *GetSmsTemplateListShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSmsTemplateListShrinkRequest) SetPageIndex(v int64) *GetSmsTemplateListShrinkRequest {
	s.PageIndex = &v
	return s
}

func (s *GetSmsTemplateListShrinkRequest) SetPageSize(v int64) *GetSmsTemplateListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSmsTemplateListShrinkRequest) SetResourceOwnerAccount(v string) *GetSmsTemplateListShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSmsTemplateListShrinkRequest) SetResourceOwnerId(v int64) *GetSmsTemplateListShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetSmsTemplateListShrinkRequest) SetSignName(v string) *GetSmsTemplateListShrinkRequest {
	s.SignName = &v
	return s
}

func (s *GetSmsTemplateListShrinkRequest) SetTemplateCode(v string) *GetSmsTemplateListShrinkRequest {
	s.TemplateCode = &v
	return s
}

func (s *GetSmsTemplateListShrinkRequest) SetTemplateName(v string) *GetSmsTemplateListShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *GetSmsTemplateListShrinkRequest) SetTemplateTag(v string) *GetSmsTemplateListShrinkRequest {
	s.TemplateTag = &v
	return s
}

func (s *GetSmsTemplateListShrinkRequest) SetTemplateType(v int64) *GetSmsTemplateListShrinkRequest {
	s.TemplateType = &v
	return s
}

func (s *GetSmsTemplateListShrinkRequest) SetUsableStateListShrink(v string) *GetSmsTemplateListShrinkRequest {
	s.UsableStateListShrink = &v
	return s
}

func (s *GetSmsTemplateListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
