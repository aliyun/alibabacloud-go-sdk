// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArchiveChatappTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchiveType(v string) *ArchiveChatappTemplateRequest
	GetArchiveType() *string
	SetChannelType(v string) *ArchiveChatappTemplateRequest
	GetChannelType() *string
	SetCustSpaceId(v string) *ArchiveChatappTemplateRequest
	GetCustSpaceId() *string
	SetTemplateList(v []*ArchiveChatappTemplateRequestTemplateList) *ArchiveChatappTemplateRequest
	GetTemplateList() []*ArchiveChatappTemplateRequestTemplateList
}

type ArchiveChatappTemplateRequest struct {
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
	TemplateList []*ArchiveChatappTemplateRequestTemplateList `json:"TemplateList,omitempty" xml:"TemplateList,omitempty" type:"Repeated"`
}

func (s ArchiveChatappTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ArchiveChatappTemplateRequest) GoString() string {
	return s.String()
}

func (s *ArchiveChatappTemplateRequest) GetArchiveType() *string {
	return s.ArchiveType
}

func (s *ArchiveChatappTemplateRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *ArchiveChatappTemplateRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ArchiveChatappTemplateRequest) GetTemplateList() []*ArchiveChatappTemplateRequestTemplateList {
	return s.TemplateList
}

func (s *ArchiveChatappTemplateRequest) SetArchiveType(v string) *ArchiveChatappTemplateRequest {
	s.ArchiveType = &v
	return s
}

func (s *ArchiveChatappTemplateRequest) SetChannelType(v string) *ArchiveChatappTemplateRequest {
	s.ChannelType = &v
	return s
}

func (s *ArchiveChatappTemplateRequest) SetCustSpaceId(v string) *ArchiveChatappTemplateRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ArchiveChatappTemplateRequest) SetTemplateList(v []*ArchiveChatappTemplateRequestTemplateList) *ArchiveChatappTemplateRequest {
	s.TemplateList = v
	return s
}

func (s *ArchiveChatappTemplateRequest) Validate() error {
	if s.TemplateList != nil {
		for _, item := range s.TemplateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ArchiveChatappTemplateRequestTemplateList struct {
	// The template language. For detailed language codes, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// zh_CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The template code. You can view the template code on the <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[**Channel Management**](https://chatapp.console.alibabacloud.com/CustomerList) > **Management*	- > **Template Design*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 939938****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s ArchiveChatappTemplateRequestTemplateList) String() string {
	return dara.Prettify(s)
}

func (s ArchiveChatappTemplateRequestTemplateList) GoString() string {
	return s.String()
}

func (s *ArchiveChatappTemplateRequestTemplateList) GetLanguage() *string {
	return s.Language
}

func (s *ArchiveChatappTemplateRequestTemplateList) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ArchiveChatappTemplateRequestTemplateList) SetLanguage(v string) *ArchiveChatappTemplateRequestTemplateList {
	s.Language = &v
	return s
}

func (s *ArchiveChatappTemplateRequestTemplateList) SetTemplateCode(v string) *ArchiveChatappTemplateRequestTemplateList {
	s.TemplateCode = &v
	return s
}

func (s *ArchiveChatappTemplateRequestTemplateList) Validate() error {
	return dara.Validate(s)
}
