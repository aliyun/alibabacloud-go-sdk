// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStructSyncOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateStructSyncOrderShrinkRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateStructSyncOrderShrinkRequest
	GetComment() *string
	SetParamShrink(v string) *CreateStructSyncOrderShrinkRequest
	GetParamShrink() *string
	SetRelatedUserListShrink(v string) *CreateStructSyncOrderShrinkRequest
	GetRelatedUserListShrink() *string
	SetTid(v int64) *CreateStructSyncOrderShrinkRequest
	GetTid() *int64
}

type CreateStructSyncOrderShrinkRequest struct {
	// The key of an attachment that is returned after the attachment is uploaded. You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to query the key of the attachment.
	//
	// example:
	//
	// upload_3c7edea3-e4c3-4403-857d-737043036f69_test.sql
	AttachmentKey *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	// The remarks of the ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The parameters of the ticket.
	//
	// This parameter is required.
	ParamShrink *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The IDs of the stakeholders.
	RelatedUserListShrink *string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the DMS console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateStructSyncOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStructSyncOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStructSyncOrderShrinkRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateStructSyncOrderShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateStructSyncOrderShrinkRequest) GetParamShrink() *string {
	return s.ParamShrink
}

func (s *CreateStructSyncOrderShrinkRequest) GetRelatedUserListShrink() *string {
	return s.RelatedUserListShrink
}

func (s *CreateStructSyncOrderShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateStructSyncOrderShrinkRequest) SetAttachmentKey(v string) *CreateStructSyncOrderShrinkRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateStructSyncOrderShrinkRequest) SetComment(v string) *CreateStructSyncOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateStructSyncOrderShrinkRequest) SetParamShrink(v string) *CreateStructSyncOrderShrinkRequest {
	s.ParamShrink = &v
	return s
}

func (s *CreateStructSyncOrderShrinkRequest) SetRelatedUserListShrink(v string) *CreateStructSyncOrderShrinkRequest {
	s.RelatedUserListShrink = &v
	return s
}

func (s *CreateStructSyncOrderShrinkRequest) SetTid(v int64) *CreateStructSyncOrderShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateStructSyncOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
