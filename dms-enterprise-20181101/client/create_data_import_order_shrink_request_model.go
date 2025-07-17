// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataImportOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateDataImportOrderShrinkRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateDataImportOrderShrinkRequest
	GetComment() *string
	SetParamShrink(v string) *CreateDataImportOrderShrinkRequest
	GetParamShrink() *string
	SetRealLoginUserUid(v string) *CreateDataImportOrderShrinkRequest
	GetRealLoginUserUid() *string
	SetRelatedUserListShrink(v string) *CreateDataImportOrderShrinkRequest
	GetRelatedUserListShrink() *string
	SetTid(v int64) *CreateDataImportOrderShrinkRequest
	GetTid() *int64
}

type CreateDataImportOrderShrinkRequest struct {
	// The key of the attachment that provides more instructions for the ticket. You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to obtain the attachment key from the value of the AttachmentKey parameter.
	//
	// example:
	//
	// order_attachment.txt
	AttachmentKey *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	// The purpose or objective of the data import. This parameter is used to help reduce unnecessary communication.
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
	// The stakeholders of the data import. All stakeholders can view the ticket details and assist in the approval process. Irrelevant users other than DMS administrators and database administrators (DBAs) are not allowed to view the ticket details.
	RelatedUserListShrink *string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDataImportOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataImportOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataImportOrderShrinkRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateDataImportOrderShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDataImportOrderShrinkRequest) GetParamShrink() *string {
	return s.ParamShrink
}

func (s *CreateDataImportOrderShrinkRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *CreateDataImportOrderShrinkRequest) GetRelatedUserListShrink() *string {
	return s.RelatedUserListShrink
}

func (s *CreateDataImportOrderShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataImportOrderShrinkRequest) SetAttachmentKey(v string) *CreateDataImportOrderShrinkRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateDataImportOrderShrinkRequest) SetComment(v string) *CreateDataImportOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataImportOrderShrinkRequest) SetParamShrink(v string) *CreateDataImportOrderShrinkRequest {
	s.ParamShrink = &v
	return s
}

func (s *CreateDataImportOrderShrinkRequest) SetRealLoginUserUid(v string) *CreateDataImportOrderShrinkRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *CreateDataImportOrderShrinkRequest) SetRelatedUserListShrink(v string) *CreateDataImportOrderShrinkRequest {
	s.RelatedUserListShrink = &v
	return s
}

func (s *CreateDataImportOrderShrinkRequest) SetTid(v int64) *CreateDataImportOrderShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataImportOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
