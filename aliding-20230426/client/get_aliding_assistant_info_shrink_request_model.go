// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlidingAssistantInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistantId(v string) *GetAlidingAssistantInfoShrinkRequest
	GetAssistantId() *string
	SetTenantContextShrink(v string) *GetAlidingAssistantInfoShrinkRequest
	GetTenantContextShrink() *string
}

type GetAlidingAssistantInfoShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AssistantId         *string `json:"AssistantId,omitempty" xml:"AssistantId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetAlidingAssistantInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlidingAssistantInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAlidingAssistantInfoShrinkRequest) GetAssistantId() *string {
	return s.AssistantId
}

func (s *GetAlidingAssistantInfoShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetAlidingAssistantInfoShrinkRequest) SetAssistantId(v string) *GetAlidingAssistantInfoShrinkRequest {
	s.AssistantId = &v
	return s
}

func (s *GetAlidingAssistantInfoShrinkRequest) SetTenantContextShrink(v string) *GetAlidingAssistantInfoShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetAlidingAssistantInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
