// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlidingAssistantShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistantId(v string) *DeleteAlidingAssistantShrinkRequest
	GetAssistantId() *string
	SetTenantContextShrink(v string) *DeleteAlidingAssistantShrinkRequest
	GetTenantContextShrink() *string
}

type DeleteAlidingAssistantShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
	AssistantId         *string `json:"AssistantId,omitempty" xml:"AssistantId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s DeleteAlidingAssistantShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlidingAssistantShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlidingAssistantShrinkRequest) GetAssistantId() *string {
	return s.AssistantId
}

func (s *DeleteAlidingAssistantShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DeleteAlidingAssistantShrinkRequest) SetAssistantId(v string) *DeleteAlidingAssistantShrinkRequest {
	s.AssistantId = &v
	return s
}

func (s *DeleteAlidingAssistantShrinkRequest) SetTenantContextShrink(v string) *DeleteAlidingAssistantShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeleteAlidingAssistantShrinkRequest) Validate() error {
	return dara.Validate(s)
}
