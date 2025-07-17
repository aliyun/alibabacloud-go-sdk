// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataTrackOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateDataTrackOrderShrinkRequest
	GetComment() *string
	SetParamShrink(v string) *CreateDataTrackOrderShrinkRequest
	GetParamShrink() *string
	SetRelatedUserListShrink(v string) *CreateDataTrackOrderShrinkRequest
	GetRelatedUserListShrink() *string
	SetTid(v int64) *CreateDataTrackOrderShrinkRequest
	GetTid() *int64
}

type CreateDataTrackOrderShrinkRequest struct {
	// The purpose or objective of the data tracking ticket. This parameter is used to help reduce unnecessary communication.
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
	// The IDs of the operators that are related to the ticket.
	RelatedUserListShrink *string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDataTrackOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataTrackOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataTrackOrderShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDataTrackOrderShrinkRequest) GetParamShrink() *string {
	return s.ParamShrink
}

func (s *CreateDataTrackOrderShrinkRequest) GetRelatedUserListShrink() *string {
	return s.RelatedUserListShrink
}

func (s *CreateDataTrackOrderShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataTrackOrderShrinkRequest) SetComment(v string) *CreateDataTrackOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataTrackOrderShrinkRequest) SetParamShrink(v string) *CreateDataTrackOrderShrinkRequest {
	s.ParamShrink = &v
	return s
}

func (s *CreateDataTrackOrderShrinkRequest) SetRelatedUserListShrink(v string) *CreateDataTrackOrderShrinkRequest {
	s.RelatedUserListShrink = &v
	return s
}

func (s *CreateDataTrackOrderShrinkRequest) SetTid(v int64) *CreateDataTrackOrderShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataTrackOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
