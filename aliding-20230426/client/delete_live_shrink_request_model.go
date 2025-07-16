// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiveId(v string) *DeleteLiveShrinkRequest
	GetLiveId() *string
	SetTenantContextShrink(v string) *DeleteLiveShrinkRequest
	GetTenantContextShrink() *string
}

type DeleteLiveShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4d38xxxxx
	LiveId              *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s DeleteLiveShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveShrinkRequest) GetLiveId() *string {
	return s.LiveId
}

func (s *DeleteLiveShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DeleteLiveShrinkRequest) SetLiveId(v string) *DeleteLiveShrinkRequest {
	s.LiveId = &v
	return s
}

func (s *DeleteLiveShrinkRequest) SetTenantContextShrink(v string) *DeleteLiveShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeleteLiveShrinkRequest) Validate() error {
	return dara.Validate(s)
}
