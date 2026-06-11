// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePromptShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizTagsShrink(v string) *UpdatePromptShrinkRequest
	GetBizTagsShrink() *string
	SetDescription(v string) *UpdatePromptShrinkRequest
	GetDescription() *string
	SetLabelsShrink(v string) *UpdatePromptShrinkRequest
	GetLabelsShrink() *string
	SetNamespaceId(v string) *UpdatePromptShrinkRequest
	GetNamespaceId() *string
	SetPromptKey(v string) *UpdatePromptShrinkRequest
	GetPromptKey() *string
}

type UpdatePromptShrinkRequest struct {
	// example:
	//
	// ["cs","qa","support"]
	BizTagsShrink *string `json:"BizTags,omitempty" xml:"BizTags,omitempty"`
	// example:
	//
	// 客服问答 Prompt
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LabelsShrink *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// customer-service-qa
	PromptKey *string `json:"PromptKey,omitempty" xml:"PromptKey,omitempty"`
}

func (s UpdatePromptShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePromptShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePromptShrinkRequest) GetBizTagsShrink() *string {
	return s.BizTagsShrink
}

func (s *UpdatePromptShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePromptShrinkRequest) GetLabelsShrink() *string {
	return s.LabelsShrink
}

func (s *UpdatePromptShrinkRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdatePromptShrinkRequest) GetPromptKey() *string {
	return s.PromptKey
}

func (s *UpdatePromptShrinkRequest) SetBizTagsShrink(v string) *UpdatePromptShrinkRequest {
	s.BizTagsShrink = &v
	return s
}

func (s *UpdatePromptShrinkRequest) SetDescription(v string) *UpdatePromptShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdatePromptShrinkRequest) SetLabelsShrink(v string) *UpdatePromptShrinkRequest {
	s.LabelsShrink = &v
	return s
}

func (s *UpdatePromptShrinkRequest) SetNamespaceId(v string) *UpdatePromptShrinkRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdatePromptShrinkRequest) SetPromptKey(v string) *UpdatePromptShrinkRequest {
	s.PromptKey = &v
	return s
}

func (s *UpdatePromptShrinkRequest) Validate() error {
	return dara.Validate(s)
}
