// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataExportOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateDataExportOrderShrinkRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateDataExportOrderShrinkRequest
	GetComment() *string
	SetParentId(v int64) *CreateDataExportOrderShrinkRequest
	GetParentId() *int64
	SetPluginParamShrink(v string) *CreateDataExportOrderShrinkRequest
	GetPluginParamShrink() *string
	SetRealLoginUserUid(v string) *CreateDataExportOrderShrinkRequest
	GetRealLoginUserUid() *string
	SetRelatedUserListShrink(v string) *CreateDataExportOrderShrinkRequest
	GetRelatedUserListShrink() *string
	SetTid(v int64) *CreateDataExportOrderShrinkRequest
	GetTid() *int64
}

type CreateDataExportOrderShrinkRequest struct {
	// The key of the attachment that provides more instructions for the ticket. You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to obtain the attachment key.
	//
	// example:
	//
	// order_attachment.txt
	AttachmentKey *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	// The purpose or objective of the ticket. This parameter helps reduce unnecessary communication.
	//
	// This parameter is required.
	//
	// example:
	//
	// business_test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the parent ticket.
	//
	// example:
	//
	// 877****
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The parameters of the ticket.
	//
	// This parameter is required.
	PluginParamShrink *string `json:"PluginParam,omitempty" xml:"PluginParam,omitempty"`
	// The UID of the Alibaba Cloud account that actually calls the API.
	//
	// example:
	//
	// 21400447956867****
	RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
	// The stakeholders involved in this operation.
	RelatedUserListShrink *string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	// The tenant ID.
	//
	// > To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the DMS console. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDataExportOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataExportOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataExportOrderShrinkRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateDataExportOrderShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDataExportOrderShrinkRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateDataExportOrderShrinkRequest) GetPluginParamShrink() *string {
	return s.PluginParamShrink
}

func (s *CreateDataExportOrderShrinkRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *CreateDataExportOrderShrinkRequest) GetRelatedUserListShrink() *string {
	return s.RelatedUserListShrink
}

func (s *CreateDataExportOrderShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataExportOrderShrinkRequest) SetAttachmentKey(v string) *CreateDataExportOrderShrinkRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateDataExportOrderShrinkRequest) SetComment(v string) *CreateDataExportOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataExportOrderShrinkRequest) SetParentId(v int64) *CreateDataExportOrderShrinkRequest {
	s.ParentId = &v
	return s
}

func (s *CreateDataExportOrderShrinkRequest) SetPluginParamShrink(v string) *CreateDataExportOrderShrinkRequest {
	s.PluginParamShrink = &v
	return s
}

func (s *CreateDataExportOrderShrinkRequest) SetRealLoginUserUid(v string) *CreateDataExportOrderShrinkRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *CreateDataExportOrderShrinkRequest) SetRelatedUserListShrink(v string) *CreateDataExportOrderShrinkRequest {
	s.RelatedUserListShrink = &v
	return s
}

func (s *CreateDataExportOrderShrinkRequest) SetTid(v int64) *CreateDataExportOrderShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataExportOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
