// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiveId(v string) *QueryLiveInfoShrinkRequest
	GetLiveId() *string
	SetTenantContextShrink(v string) *QueryLiveInfoShrinkRequest
	GetTenantContextShrink() *string
}

type QueryLiveInfoShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4d38xxxxx
	LiveId              *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s QueryLiveInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryLiveInfoShrinkRequest) GetLiveId() *string {
	return s.LiveId
}

func (s *QueryLiveInfoShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryLiveInfoShrinkRequest) SetLiveId(v string) *QueryLiveInfoShrinkRequest {
	s.LiveId = &v
	return s
}

func (s *QueryLiveInfoShrinkRequest) SetTenantContextShrink(v string) *QueryLiveInfoShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryLiveInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
