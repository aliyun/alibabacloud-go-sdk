// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseExportOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateDatabaseExportOrderShrinkRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateDatabaseExportOrderShrinkRequest
	GetComment() *string
	SetParentId(v int64) *CreateDatabaseExportOrderShrinkRequest
	GetParentId() *int64
	SetPluginParamShrink(v string) *CreateDatabaseExportOrderShrinkRequest
	GetPluginParamShrink() *string
	SetRelatedUserListShrink(v string) *CreateDatabaseExportOrderShrinkRequest
	GetRelatedUserListShrink() *string
	SetTid(v int64) *CreateDatabaseExportOrderShrinkRequest
	GetTid() *int64
}

type CreateDatabaseExportOrderShrinkRequest struct {
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
	// document_test
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

func (s CreateDatabaseExportOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseExportOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDatabaseExportOrderShrinkRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateDatabaseExportOrderShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDatabaseExportOrderShrinkRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateDatabaseExportOrderShrinkRequest) GetPluginParamShrink() *string {
	return s.PluginParamShrink
}

func (s *CreateDatabaseExportOrderShrinkRequest) GetRelatedUserListShrink() *string {
	return s.RelatedUserListShrink
}

func (s *CreateDatabaseExportOrderShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDatabaseExportOrderShrinkRequest) SetAttachmentKey(v string) *CreateDatabaseExportOrderShrinkRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateDatabaseExportOrderShrinkRequest) SetComment(v string) *CreateDatabaseExportOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateDatabaseExportOrderShrinkRequest) SetParentId(v int64) *CreateDatabaseExportOrderShrinkRequest {
	s.ParentId = &v
	return s
}

func (s *CreateDatabaseExportOrderShrinkRequest) SetPluginParamShrink(v string) *CreateDatabaseExportOrderShrinkRequest {
	s.PluginParamShrink = &v
	return s
}

func (s *CreateDatabaseExportOrderShrinkRequest) SetRelatedUserListShrink(v string) *CreateDatabaseExportOrderShrinkRequest {
	s.RelatedUserListShrink = &v
	return s
}

func (s *CreateDatabaseExportOrderShrinkRequest) SetTid(v int64) *CreateDatabaseExportOrderShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateDatabaseExportOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
