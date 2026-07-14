// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArchiveChatappTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchiveType(v string) *ArchiveChatappTemplateShrinkRequest
	GetArchiveType() *string
	SetChannelType(v string) *ArchiveChatappTemplateShrinkRequest
	GetChannelType() *string
	SetCustSpaceId(v string) *ArchiveChatappTemplateShrinkRequest
	GetCustSpaceId() *string
	SetTemplateListShrink(v string) *ArchiveChatappTemplateShrinkRequest
	GetTemplateListShrink() *string
}

type ArchiveChatappTemplateShrinkRequest struct {
	// The archive type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ARCHIVED
	ArchiveType *string `json:"ArchiveType,omitempty" xml:"ArchiveType,omitempty"`
	// The channel type. Valid values:
	//
	// - **WHATSAPP**.
	//
	// > Only the WhatsApp channel type is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The space ID of the ISV sub-customer or the instance ID of the direct customer. You can view the Space ID on the <props="china">[Channel Management](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[Channel Management](https://chatapp.console.alibabacloud.com/CustomerList) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The template list.
	//
	// This parameter is required.
	TemplateListShrink *string `json:"TemplateList,omitempty" xml:"TemplateList,omitempty"`
}

func (s ArchiveChatappTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ArchiveChatappTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ArchiveChatappTemplateShrinkRequest) GetArchiveType() *string {
	return s.ArchiveType
}

func (s *ArchiveChatappTemplateShrinkRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *ArchiveChatappTemplateShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ArchiveChatappTemplateShrinkRequest) GetTemplateListShrink() *string {
	return s.TemplateListShrink
}

func (s *ArchiveChatappTemplateShrinkRequest) SetArchiveType(v string) *ArchiveChatappTemplateShrinkRequest {
	s.ArchiveType = &v
	return s
}

func (s *ArchiveChatappTemplateShrinkRequest) SetChannelType(v string) *ArchiveChatappTemplateShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *ArchiveChatappTemplateShrinkRequest) SetCustSpaceId(v string) *ArchiveChatappTemplateShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ArchiveChatappTemplateShrinkRequest) SetTemplateListShrink(v string) *ArchiveChatappTemplateShrinkRequest {
	s.TemplateListShrink = &v
	return s
}

func (s *ArchiveChatappTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
