// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSQLReviewOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateSQLReviewOrderShrinkRequest
	GetComment() *string
	SetParamShrink(v string) *CreateSQLReviewOrderShrinkRequest
	GetParamShrink() *string
	SetRelatedUserListShrink(v string) *CreateSQLReviewOrderShrinkRequest
	GetRelatedUserListShrink() *string
	SetTid(v int64) *CreateSQLReviewOrderShrinkRequest
	GetTid() *int64
}

type CreateSQLReviewOrderShrinkRequest struct {
	// The purpose or objective of the SQL review. This reduces unnecessary communication.
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
	// The stakeholders involved in this operation. All the specified stakeholders can view the ticket details and take part in the approval process. Irrelevant users other than DMS administrators and database administrators (DBAs) are not allowed to view the ticket details.
	RelatedUserListShrink *string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateSQLReviewOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSQLReviewOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSQLReviewOrderShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateSQLReviewOrderShrinkRequest) GetParamShrink() *string {
	return s.ParamShrink
}

func (s *CreateSQLReviewOrderShrinkRequest) GetRelatedUserListShrink() *string {
	return s.RelatedUserListShrink
}

func (s *CreateSQLReviewOrderShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateSQLReviewOrderShrinkRequest) SetComment(v string) *CreateSQLReviewOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateSQLReviewOrderShrinkRequest) SetParamShrink(v string) *CreateSQLReviewOrderShrinkRequest {
	s.ParamShrink = &v
	return s
}

func (s *CreateSQLReviewOrderShrinkRequest) SetRelatedUserListShrink(v string) *CreateSQLReviewOrderShrinkRequest {
	s.RelatedUserListShrink = &v
	return s
}

func (s *CreateSQLReviewOrderShrinkRequest) SetTid(v int64) *CreateSQLReviewOrderShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateSQLReviewOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
