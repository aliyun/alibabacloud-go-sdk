// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveReplayUrlShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiveId(v string) *GetLiveReplayUrlShrinkRequest
	GetLiveId() *string
	SetTenantContextShrink(v string) *GetLiveReplayUrlShrinkRequest
	GetTenantContextShrink() *string
}

type GetLiveReplayUrlShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4d38xxxxx
	LiveId              *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetLiveReplayUrlShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLiveReplayUrlShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetLiveReplayUrlShrinkRequest) GetLiveId() *string {
	return s.LiveId
}

func (s *GetLiveReplayUrlShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetLiveReplayUrlShrinkRequest) SetLiveId(v string) *GetLiveReplayUrlShrinkRequest {
	s.LiveId = &v
	return s
}

func (s *GetLiveReplayUrlShrinkRequest) SetTenantContextShrink(v string) *GetLiveReplayUrlShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetLiveReplayUrlShrinkRequest) Validate() error {
	return dara.Validate(s)
}
