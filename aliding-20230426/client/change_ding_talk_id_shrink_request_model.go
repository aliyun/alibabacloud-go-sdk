// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDingTalkIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDingTalkId(v string) *ChangeDingTalkIdShrinkRequest
	GetDingTalkId() *string
	SetTenantContextShrink(v string) *ChangeDingTalkIdShrinkRequest
	GetTenantContextShrink() *string
}

type ChangeDingTalkIdShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4uf_iw54grufg9
	DingTalkId          *string `json:"DingTalkId,omitempty" xml:"DingTalkId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s ChangeDingTalkIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeDingTalkIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChangeDingTalkIdShrinkRequest) GetDingTalkId() *string {
	return s.DingTalkId
}

func (s *ChangeDingTalkIdShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *ChangeDingTalkIdShrinkRequest) SetDingTalkId(v string) *ChangeDingTalkIdShrinkRequest {
	s.DingTalkId = &v
	return s
}

func (s *ChangeDingTalkIdShrinkRequest) SetTenantContextShrink(v string) *ChangeDingTalkIdShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *ChangeDingTalkIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
