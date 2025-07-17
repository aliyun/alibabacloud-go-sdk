// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataCorrectOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateDataCorrectOrderShrinkRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateDataCorrectOrderShrinkRequest
	GetComment() *string
	SetParamShrink(v string) *CreateDataCorrectOrderShrinkRequest
	GetParamShrink() *string
	SetRealLoginUserUid(v string) *CreateDataCorrectOrderShrinkRequest
	GetRealLoginUserUid() *string
	SetRelatedUserListShrink(v string) *CreateDataCorrectOrderShrinkRequest
	GetRelatedUserListShrink() *string
	SetTid(v int64) *CreateDataCorrectOrderShrinkRequest
	GetTid() *int64
}

type CreateDataCorrectOrderShrinkRequest struct {
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
	ParamShrink *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The ID of the Alibaba Cloud account that is used to call the API operation.
	//
	// example:
	//
	// 21400447956867****
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

func (s CreateDataCorrectOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataCorrectOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataCorrectOrderShrinkRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateDataCorrectOrderShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDataCorrectOrderShrinkRequest) GetParamShrink() *string {
	return s.ParamShrink
}

func (s *CreateDataCorrectOrderShrinkRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *CreateDataCorrectOrderShrinkRequest) GetRelatedUserListShrink() *string {
	return s.RelatedUserListShrink
}

func (s *CreateDataCorrectOrderShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataCorrectOrderShrinkRequest) SetAttachmentKey(v string) *CreateDataCorrectOrderShrinkRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateDataCorrectOrderShrinkRequest) SetComment(v string) *CreateDataCorrectOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataCorrectOrderShrinkRequest) SetParamShrink(v string) *CreateDataCorrectOrderShrinkRequest {
	s.ParamShrink = &v
	return s
}

func (s *CreateDataCorrectOrderShrinkRequest) SetRealLoginUserUid(v string) *CreateDataCorrectOrderShrinkRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *CreateDataCorrectOrderShrinkRequest) SetRelatedUserListShrink(v string) *CreateDataCorrectOrderShrinkRequest {
	s.RelatedUserListShrink = &v
	return s
}

func (s *CreateDataCorrectOrderShrinkRequest) SetTid(v int64) *CreateDataCorrectOrderShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataCorrectOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
