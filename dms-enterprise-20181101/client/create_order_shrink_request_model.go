// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateOrderShrinkRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateOrderShrinkRequest
	GetComment() *string
	SetPluginParamShrink(v string) *CreateOrderShrinkRequest
	GetPluginParamShrink() *string
	SetPluginType(v string) *CreateOrderShrinkRequest
	GetPluginType() *string
	SetRelatedUserList(v string) *CreateOrderShrinkRequest
	GetRelatedUserList() *string
	SetTid(v int64) *CreateOrderShrinkRequest
	GetTid() *int64
}

type CreateOrderShrinkRequest struct {
	// The key of an attachment that is returned after the attachment is uploaded. You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to query the key of the attachment.
	//
	// example:
	//
	// test_AttachmentKey
	AttachmentKey *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	// The description of the ticket to be created.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ticket creation parameter. The value is a JSON string. The value of this parameter differs based on the type of the ticket. For more information, see the **PluginParam parameter*	- section in this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// {PluginParam_test}
	PluginParamShrink *string `json:"PluginParam,omitempty" xml:"PluginParam,omitempty"`
	// The type of the ticket. For more information, see [PluginType parameter](https://help.aliyun.com/document_detail/429109.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// DATA_EXPORT
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	// The IDs of the stakeholders that are involved in the ticket. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// user1,user2
	RelatedUserList *string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderShrinkRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateOrderShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateOrderShrinkRequest) GetPluginParamShrink() *string {
	return s.PluginParamShrink
}

func (s *CreateOrderShrinkRequest) GetPluginType() *string {
	return s.PluginType
}

func (s *CreateOrderShrinkRequest) GetRelatedUserList() *string {
	return s.RelatedUserList
}

func (s *CreateOrderShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateOrderShrinkRequest) SetAttachmentKey(v string) *CreateOrderShrinkRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetComment(v string) *CreateOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetPluginParamShrink(v string) *CreateOrderShrinkRequest {
	s.PluginParamShrink = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetPluginType(v string) *CreateOrderShrinkRequest {
	s.PluginType = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetRelatedUserList(v string) *CreateOrderShrinkRequest {
	s.RelatedUserList = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetTid(v int64) *CreateOrderShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
