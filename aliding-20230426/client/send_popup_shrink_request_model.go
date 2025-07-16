// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendPopupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentShrink(v string) *SendPopupShrinkRequest
	GetContentShrink() *string
	SetEndTime(v int64) *SendPopupShrinkRequest
	GetEndTime() *int64
	SetStartTime(v int64) *SendPopupShrinkRequest
	GetStartTime() *int64
	SetTenantContextShrink(v string) *SendPopupShrinkRequest
	GetTenantContextShrink() *string
}

type SendPopupShrinkRequest struct {
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1693881641000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1693881641000
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s SendPopupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendPopupShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendPopupShrinkRequest) GetContentShrink() *string {
	return s.ContentShrink
}

func (s *SendPopupShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SendPopupShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SendPopupShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *SendPopupShrinkRequest) SetContentShrink(v string) *SendPopupShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *SendPopupShrinkRequest) SetEndTime(v int64) *SendPopupShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *SendPopupShrinkRequest) SetStartTime(v int64) *SendPopupShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *SendPopupShrinkRequest) SetTenantContextShrink(v string) *SendPopupShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *SendPopupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
