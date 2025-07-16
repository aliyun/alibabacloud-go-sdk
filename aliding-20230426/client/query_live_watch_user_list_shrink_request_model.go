// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveWatchUserListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiveId(v string) *QueryLiveWatchUserListShrinkRequest
	GetLiveId() *string
	SetPageNumber(v int32) *QueryLiveWatchUserListShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryLiveWatchUserListShrinkRequest
	GetPageSize() *int32
	SetTenantContextShrink(v string) *QueryLiveWatchUserListShrinkRequest
	GetTenantContextShrink() *string
}

type QueryLiveWatchUserListShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4d38xxxxx
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// example:
	//
	// 0
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s QueryLiveWatchUserListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchUserListShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchUserListShrinkRequest) GetLiveId() *string {
	return s.LiveId
}

func (s *QueryLiveWatchUserListShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryLiveWatchUserListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryLiveWatchUserListShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryLiveWatchUserListShrinkRequest) SetLiveId(v string) *QueryLiveWatchUserListShrinkRequest {
	s.LiveId = &v
	return s
}

func (s *QueryLiveWatchUserListShrinkRequest) SetPageNumber(v int32) *QueryLiveWatchUserListShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryLiveWatchUserListShrinkRequest) SetPageSize(v int32) *QueryLiveWatchUserListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryLiveWatchUserListShrinkRequest) SetTenantContextShrink(v string) *QueryLiveWatchUserListShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryLiveWatchUserListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
