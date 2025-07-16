// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupLiveListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetGroupLiveListShrinkRequest
	GetEndTime() *int64
	SetOpenConversationId(v string) *GetGroupLiveListShrinkRequest
	GetOpenConversationId() *string
	SetStartTime(v int64) *GetGroupLiveListShrinkRequest
	GetStartTime() *int64
	SetTenantContextShrink(v string) *GetGroupLiveListShrinkRequest
	GetTenantContextShrink() *string
}

type GetGroupLiveListShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1398324600000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidt*****Xa4K10w==
	OpenConversationId *string `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1398324600000
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetGroupLiveListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGroupLiveListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetGroupLiveListShrinkRequest) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *GetGroupLiveListShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetGroupLiveListShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetGroupLiveListShrinkRequest) SetEndTime(v int64) *GetGroupLiveListShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetGroupLiveListShrinkRequest) SetOpenConversationId(v string) *GetGroupLiveListShrinkRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetGroupLiveListShrinkRequest) SetStartTime(v int64) *GetGroupLiveListShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *GetGroupLiveListShrinkRequest) SetTenantContextShrink(v string) *GetGroupLiveListShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetGroupLiveListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
