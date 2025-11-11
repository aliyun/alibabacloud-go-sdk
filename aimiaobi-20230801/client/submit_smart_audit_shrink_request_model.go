// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmartAuditShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrlListShrink(v string) *SubmitSmartAuditShrinkRequest
	GetImageUrlListShrink() *string
	SetNoteId(v string) *SubmitSmartAuditShrinkRequest
	GetNoteId() *string
	SetSubCodesShrink(v string) *SubmitSmartAuditShrinkRequest
	GetSubCodesShrink() *string
	SetTermsName(v string) *SubmitSmartAuditShrinkRequest
	GetTermsName() *string
	SetText(v string) *SubmitSmartAuditShrinkRequest
	GetText() *string
	SetWorkspaceId(v string) *SubmitSmartAuditShrinkRequest
	GetWorkspaceId() *string
	SetImageUrlsShrink(v string) *SubmitSmartAuditShrinkRequest
	GetImageUrlsShrink() *string
}

type SubmitSmartAuditShrinkRequest struct {
	ImageUrlListShrink *string `json:"ImageUrlList,omitempty" xml:"ImageUrlList,omitempty"`
	NoteId             *string `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
	SubCodesShrink     *string `json:"SubCodes,omitempty" xml:"SubCodes,omitempty"`
	TermsName          *string `json:"TermsName,omitempty" xml:"TermsName,omitempty"`
	Text               *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// xxxx
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	ImageUrlsShrink *string `json:"imageUrls,omitempty" xml:"imageUrls,omitempty"`
}

func (s SubmitSmartAuditShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartAuditShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmartAuditShrinkRequest) GetImageUrlListShrink() *string {
	return s.ImageUrlListShrink
}

func (s *SubmitSmartAuditShrinkRequest) GetNoteId() *string {
	return s.NoteId
}

func (s *SubmitSmartAuditShrinkRequest) GetSubCodesShrink() *string {
	return s.SubCodesShrink
}

func (s *SubmitSmartAuditShrinkRequest) GetTermsName() *string {
	return s.TermsName
}

func (s *SubmitSmartAuditShrinkRequest) GetText() *string {
	return s.Text
}

func (s *SubmitSmartAuditShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitSmartAuditShrinkRequest) GetImageUrlsShrink() *string {
	return s.ImageUrlsShrink
}

func (s *SubmitSmartAuditShrinkRequest) SetImageUrlListShrink(v string) *SubmitSmartAuditShrinkRequest {
	s.ImageUrlListShrink = &v
	return s
}

func (s *SubmitSmartAuditShrinkRequest) SetNoteId(v string) *SubmitSmartAuditShrinkRequest {
	s.NoteId = &v
	return s
}

func (s *SubmitSmartAuditShrinkRequest) SetSubCodesShrink(v string) *SubmitSmartAuditShrinkRequest {
	s.SubCodesShrink = &v
	return s
}

func (s *SubmitSmartAuditShrinkRequest) SetTermsName(v string) *SubmitSmartAuditShrinkRequest {
	s.TermsName = &v
	return s
}

func (s *SubmitSmartAuditShrinkRequest) SetText(v string) *SubmitSmartAuditShrinkRequest {
	s.Text = &v
	return s
}

func (s *SubmitSmartAuditShrinkRequest) SetWorkspaceId(v string) *SubmitSmartAuditShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitSmartAuditShrinkRequest) SetImageUrlsShrink(v string) *SubmitSmartAuditShrinkRequest {
	s.ImageUrlsShrink = &v
	return s
}

func (s *SubmitSmartAuditShrinkRequest) Validate() error {
	return dara.Validate(s)
}
