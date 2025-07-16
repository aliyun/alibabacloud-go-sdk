// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveWatchDetailShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiveId(v string) *QueryLiveWatchDetailShrinkRequest
	GetLiveId() *string
	SetTenantContextShrink(v string) *QueryLiveWatchDetailShrinkRequest
	GetTenantContextShrink() *string
}

type QueryLiveWatchDetailShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4d38xxxxx
	LiveId              *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s QueryLiveWatchDetailShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchDetailShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchDetailShrinkRequest) GetLiveId() *string {
	return s.LiveId
}

func (s *QueryLiveWatchDetailShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryLiveWatchDetailShrinkRequest) SetLiveId(v string) *QueryLiveWatchDetailShrinkRequest {
	s.LiveId = &v
	return s
}

func (s *QueryLiveWatchDetailShrinkRequest) SetTenantContextShrink(v string) *QueryLiveWatchDetailShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryLiveWatchDetailShrinkRequest) Validate() error {
	return dara.Validate(s)
}
