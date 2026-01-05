// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsTemplateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditStatus(v int64) *GetSmsTemplateListRequest
	GetAuditStatus() *int64
	SetOwnerId(v int64) *GetSmsTemplateListRequest
	GetOwnerId() *int64
	SetPageIndex(v int64) *GetSmsTemplateListRequest
	GetPageIndex() *int64
	SetPageSize(v int64) *GetSmsTemplateListRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *GetSmsTemplateListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetSmsTemplateListRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *GetSmsTemplateListRequest
	GetSignName() *string
	SetTemplateCode(v string) *GetSmsTemplateListRequest
	GetTemplateCode() *string
	SetTemplateName(v string) *GetSmsTemplateListRequest
	GetTemplateName() *string
	SetTemplateTag(v string) *GetSmsTemplateListRequest
	GetTemplateTag() *string
	SetTemplateType(v int64) *GetSmsTemplateListRequest
	GetTemplateType() *int64
	SetUsableStateList(v []*string) *GetSmsTemplateListRequest
	GetUsableStateList() []*string
}

type GetSmsTemplateListRequest struct {
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
	UsableStateList []*string `json:"UsableStateList,omitempty" xml:"UsableStateList,omitempty" type:"Repeated"`
}

func (s GetSmsTemplateListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSmsTemplateListRequest) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateListRequest) GetAuditStatus() *int64 {
	return s.AuditStatus
}

func (s *GetSmsTemplateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetSmsTemplateListRequest) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *GetSmsTemplateListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetSmsTemplateListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetSmsTemplateListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetSmsTemplateListRequest) GetSignName() *string {
	return s.SignName
}

func (s *GetSmsTemplateListRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetSmsTemplateListRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetSmsTemplateListRequest) GetTemplateTag() *string {
	return s.TemplateTag
}

func (s *GetSmsTemplateListRequest) GetTemplateType() *int64 {
	return s.TemplateType
}

func (s *GetSmsTemplateListRequest) GetUsableStateList() []*string {
	return s.UsableStateList
}

func (s *GetSmsTemplateListRequest) SetAuditStatus(v int64) *GetSmsTemplateListRequest {
	s.AuditStatus = &v
	return s
}

func (s *GetSmsTemplateListRequest) SetOwnerId(v int64) *GetSmsTemplateListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSmsTemplateListRequest) SetPageIndex(v int64) *GetSmsTemplateListRequest {
	s.PageIndex = &v
	return s
}

func (s *GetSmsTemplateListRequest) SetPageSize(v int64) *GetSmsTemplateListRequest {
	s.PageSize = &v
	return s
}

func (s *GetSmsTemplateListRequest) SetResourceOwnerAccount(v string) *GetSmsTemplateListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSmsTemplateListRequest) SetResourceOwnerId(v int64) *GetSmsTemplateListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetSmsTemplateListRequest) SetSignName(v string) *GetSmsTemplateListRequest {
	s.SignName = &v
	return s
}

func (s *GetSmsTemplateListRequest) SetTemplateCode(v string) *GetSmsTemplateListRequest {
	s.TemplateCode = &v
	return s
}

func (s *GetSmsTemplateListRequest) SetTemplateName(v string) *GetSmsTemplateListRequest {
	s.TemplateName = &v
	return s
}

func (s *GetSmsTemplateListRequest) SetTemplateTag(v string) *GetSmsTemplateListRequest {
	s.TemplateTag = &v
	return s
}

func (s *GetSmsTemplateListRequest) SetTemplateType(v int64) *GetSmsTemplateListRequest {
	s.TemplateType = &v
	return s
}

func (s *GetSmsTemplateListRequest) SetUsableStateList(v []*string) *GetSmsTemplateListRequest {
	s.UsableStateList = v
	return s
}

func (s *GetSmsTemplateListRequest) Validate() error {
	return dara.Validate(s)
}
