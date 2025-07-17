// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataCronClearOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateDataCronClearOrderShrinkRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateDataCronClearOrderShrinkRequest
	GetComment() *string
	SetParamShrink(v string) *CreateDataCronClearOrderShrinkRequest
	GetParamShrink() *string
	SetRelatedUserListShrink(v string) *CreateDataCronClearOrderShrinkRequest
	GetRelatedUserListShrink() *string
	SetTid(v int64) *CreateDataCronClearOrderShrinkRequest
	GetTid() *int64
}

type CreateDataCronClearOrderShrinkRequest struct {
	// The key of the attachment for the ticket. The attachment provides more instructions for this operation.
	//
	// You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to query the key of the attachment.
	//
	// example:
	//
	// order_attachement.txt
	AttachmentKey *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	// The purpose or objective of the data change. This reduces unnecessary communication.
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
	// The stakeholders of this operation. All stakeholders can view the ticket details and assist in the approval process. Irrelevant users other than Data Management (DMS) administrators and database administrators (DBAs) are not allowed to view the ticket details.
	RelatedUserListShrink *string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	// The ID of the tenant.
	//
	// >  The ID of the tenant is displayed when you move the pointer over the profile picture in the upper-right corner of the DMS console. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the Manage DMS tenants topic.
	//
	// example:
	//
	// 123454324
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDataCronClearOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataCronClearOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataCronClearOrderShrinkRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateDataCronClearOrderShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDataCronClearOrderShrinkRequest) GetParamShrink() *string {
	return s.ParamShrink
}

func (s *CreateDataCronClearOrderShrinkRequest) GetRelatedUserListShrink() *string {
	return s.RelatedUserListShrink
}

func (s *CreateDataCronClearOrderShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataCronClearOrderShrinkRequest) SetAttachmentKey(v string) *CreateDataCronClearOrderShrinkRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateDataCronClearOrderShrinkRequest) SetComment(v string) *CreateDataCronClearOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataCronClearOrderShrinkRequest) SetParamShrink(v string) *CreateDataCronClearOrderShrinkRequest {
	s.ParamShrink = &v
	return s
}

func (s *CreateDataCronClearOrderShrinkRequest) SetRelatedUserListShrink(v string) *CreateDataCronClearOrderShrinkRequest {
	s.RelatedUserListShrink = &v
	return s
}

func (s *CreateDataCronClearOrderShrinkRequest) SetTid(v int64) *CreateDataCronClearOrderShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataCronClearOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
