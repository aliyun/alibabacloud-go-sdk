// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFreeLockCorrectOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateFreeLockCorrectOrderShrinkRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateFreeLockCorrectOrderShrinkRequest
	GetComment() *string
	SetParamShrink(v string) *CreateFreeLockCorrectOrderShrinkRequest
	GetParamShrink() *string
	SetRealLoginUserUid(v string) *CreateFreeLockCorrectOrderShrinkRequest
	GetRealLoginUserUid() *string
	SetRelatedUserListShrink(v string) *CreateFreeLockCorrectOrderShrinkRequest
	GetRelatedUserListShrink() *string
	SetTid(v int64) *CreateFreeLockCorrectOrderShrinkRequest
	GetTid() *int64
}

type CreateFreeLockCorrectOrderShrinkRequest struct {
	// The key of the attachment that provides more instructions for the ticket. You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to obtain the attachment key from the value of the AttachmentKey parameter.
	//
	// example:
	//
	// order_attachment.txt
	AttachmentKey *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	// The purpose or objective of the data change. This parameter is used to help reduce unnecessary communication.
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
	ParamShrink      *string `json:"Param,omitempty" xml:"Param,omitempty"`
	RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
	// The stakeholders of the data change. All stakeholders can view the ticket details and assist in the approval process. Irrelevant users other than DMS administrators and database administrators (DBAs) are not allowed to view the ticket details.
	RelatedUserListShrink *string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateFreeLockCorrectOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFreeLockCorrectOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) GetParamShrink() *string {
	return s.ParamShrink
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) GetRelatedUserListShrink() *string {
	return s.RelatedUserListShrink
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) SetAttachmentKey(v string) *CreateFreeLockCorrectOrderShrinkRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) SetComment(v string) *CreateFreeLockCorrectOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) SetParamShrink(v string) *CreateFreeLockCorrectOrderShrinkRequest {
	s.ParamShrink = &v
	return s
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) SetRealLoginUserUid(v string) *CreateFreeLockCorrectOrderShrinkRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) SetRelatedUserListShrink(v string) *CreateFreeLockCorrectOrderShrinkRequest {
	s.RelatedUserListShrink = &v
	return s
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) SetTid(v int64) *CreateFreeLockCorrectOrderShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateFreeLockCorrectOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
