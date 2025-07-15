// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmartAuditShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubCodesShrink(v string) *SubmitSmartAuditShrinkRequest
	GetSubCodesShrink() *string
	SetText(v string) *SubmitSmartAuditShrinkRequest
	GetText() *string
	SetWorkspaceId(v string) *SubmitSmartAuditShrinkRequest
	GetWorkspaceId() *string
}

type SubmitSmartAuditShrinkRequest struct {
	SubCodesShrink *string `json:"SubCodes,omitempty" xml:"SubCodes,omitempty"`
	Text           *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitSmartAuditShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartAuditShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmartAuditShrinkRequest) GetSubCodesShrink() *string {
	return s.SubCodesShrink
}

func (s *SubmitSmartAuditShrinkRequest) GetText() *string {
	return s.Text
}

func (s *SubmitSmartAuditShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitSmartAuditShrinkRequest) SetSubCodesShrink(v string) *SubmitSmartAuditShrinkRequest {
	s.SubCodesShrink = &v
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

func (s *SubmitSmartAuditShrinkRequest) Validate() error {
	return dara.Validate(s)
}
