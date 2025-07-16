// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendBannerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentShrink(v string) *SendBannerShrinkRequest
	GetContentShrink() *string
	SetEndTime(v int64) *SendBannerShrinkRequest
	GetEndTime() *int64
	SetStartTime(v int64) *SendBannerShrinkRequest
	GetStartTime() *int64
	SetTenantContextShrink(v string) *SendBannerShrinkRequest
	GetTenantContextShrink() *string
}

type SendBannerShrinkRequest struct {
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1693881641000L
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1693881641000L
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s SendBannerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendBannerShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendBannerShrinkRequest) GetContentShrink() *string {
	return s.ContentShrink
}

func (s *SendBannerShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SendBannerShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SendBannerShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *SendBannerShrinkRequest) SetContentShrink(v string) *SendBannerShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *SendBannerShrinkRequest) SetEndTime(v int64) *SendBannerShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *SendBannerShrinkRequest) SetStartTime(v int64) *SendBannerShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *SendBannerShrinkRequest) SetTenantContextShrink(v string) *SendBannerShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *SendBannerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
