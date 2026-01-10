// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateOrderRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateOrderRequest
	GetComment() *string
	SetPluginParam(v map[string]interface{}) *CreateOrderRequest
	GetPluginParam() map[string]interface{}
	SetPluginType(v string) *CreateOrderRequest
	GetPluginType() *string
	SetRealLoginUserUid(v string) *CreateOrderRequest
	GetRealLoginUserUid() *string
	SetRelatedUserList(v string) *CreateOrderRequest
	GetRelatedUserList() *string
	SetTid(v int64) *CreateOrderRequest
	GetTid() *int64
}

type CreateOrderRequest struct {
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
	PluginParam map[string]interface{} `json:"PluginParam,omitempty" xml:"PluginParam,omitempty"`
	// The type of the ticket. For more information, see [PluginType parameter](https://help.aliyun.com/document_detail/429109.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// DATA_EXPORT
	PluginType       *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
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

func (s CreateOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateOrderRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateOrderRequest) GetPluginParam() map[string]interface{} {
	return s.PluginParam
}

func (s *CreateOrderRequest) GetPluginType() *string {
	return s.PluginType
}

func (s *CreateOrderRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *CreateOrderRequest) GetRelatedUserList() *string {
	return s.RelatedUserList
}

func (s *CreateOrderRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateOrderRequest) SetAttachmentKey(v string) *CreateOrderRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateOrderRequest) SetComment(v string) *CreateOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateOrderRequest) SetPluginParam(v map[string]interface{}) *CreateOrderRequest {
	s.PluginParam = v
	return s
}

func (s *CreateOrderRequest) SetPluginType(v string) *CreateOrderRequest {
	s.PluginType = &v
	return s
}

func (s *CreateOrderRequest) SetRealLoginUserUid(v string) *CreateOrderRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *CreateOrderRequest) SetRelatedUserList(v string) *CreateOrderRequest {
	s.RelatedUserList = &v
	return s
}

func (s *CreateOrderRequest) SetTid(v int64) *CreateOrderRequest {
	s.Tid = &v
	return s
}

func (s *CreateOrderRequest) Validate() error {
	return dara.Validate(s)
}
