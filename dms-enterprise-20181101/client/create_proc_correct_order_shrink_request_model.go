// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProcCorrectOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateProcCorrectOrderShrinkRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateProcCorrectOrderShrinkRequest
	GetComment() *string
	SetParamShrink(v string) *CreateProcCorrectOrderShrinkRequest
	GetParamShrink() *string
	SetRelatedUserListShrink(v string) *CreateProcCorrectOrderShrinkRequest
	GetRelatedUserListShrink() *string
	SetTid(v int64) *CreateProcCorrectOrderShrinkRequest
	GetTid() *int64
}

type CreateProcCorrectOrderShrinkRequest struct {
	// example:
	//
	// order_attachment.txt
	AttachmentKey *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	ParamShrink           *string `json:"Param,omitempty" xml:"Param,omitempty"`
	RelatedUserListShrink *string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	// example:
	//
	// 4***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateProcCorrectOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProcCorrectOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProcCorrectOrderShrinkRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateProcCorrectOrderShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateProcCorrectOrderShrinkRequest) GetParamShrink() *string {
	return s.ParamShrink
}

func (s *CreateProcCorrectOrderShrinkRequest) GetRelatedUserListShrink() *string {
	return s.RelatedUserListShrink
}

func (s *CreateProcCorrectOrderShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateProcCorrectOrderShrinkRequest) SetAttachmentKey(v string) *CreateProcCorrectOrderShrinkRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateProcCorrectOrderShrinkRequest) SetComment(v string) *CreateProcCorrectOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateProcCorrectOrderShrinkRequest) SetParamShrink(v string) *CreateProcCorrectOrderShrinkRequest {
	s.ParamShrink = &v
	return s
}

func (s *CreateProcCorrectOrderShrinkRequest) SetRelatedUserListShrink(v string) *CreateProcCorrectOrderShrinkRequest {
	s.RelatedUserListShrink = &v
	return s
}

func (s *CreateProcCorrectOrderShrinkRequest) SetTid(v int64) *CreateProcCorrectOrderShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateProcCorrectOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
