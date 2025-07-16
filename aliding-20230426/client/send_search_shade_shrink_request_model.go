// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSearchShadeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentShrink(v string) *SendSearchShadeShrinkRequest
	GetContentShrink() *string
	SetEndTime(v int64) *SendSearchShadeShrinkRequest
	GetEndTime() *int64
	SetStartTime(v int64) *SendSearchShadeShrinkRequest
	GetStartTime() *int64
	SetTenantContextShrink(v string) *SendSearchShadeShrinkRequest
	GetTenantContextShrink() *string
}

type SendSearchShadeShrinkRequest struct {
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

func (s SendSearchShadeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendSearchShadeShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendSearchShadeShrinkRequest) GetContentShrink() *string {
	return s.ContentShrink
}

func (s *SendSearchShadeShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SendSearchShadeShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SendSearchShadeShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *SendSearchShadeShrinkRequest) SetContentShrink(v string) *SendSearchShadeShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *SendSearchShadeShrinkRequest) SetEndTime(v int64) *SendSearchShadeShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *SendSearchShadeShrinkRequest) SetStartTime(v int64) *SendSearchShadeShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *SendSearchShadeShrinkRequest) SetTenantContextShrink(v string) *SendSearchShadeShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *SendSearchShadeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
