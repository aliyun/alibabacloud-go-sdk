// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveWatchUserListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiveId(v string) *QueryLiveWatchUserListRequest
	GetLiveId() *string
	SetPageNumber(v int32) *QueryLiveWatchUserListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryLiveWatchUserListRequest
	GetPageSize() *int32
	SetTenantContext(v *QueryLiveWatchUserListRequestTenantContext) *QueryLiveWatchUserListRequest
	GetTenantContext() *QueryLiveWatchUserListRequestTenantContext
}

type QueryLiveWatchUserListRequest struct {
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
	PageSize      *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TenantContext *QueryLiveWatchUserListRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s QueryLiveWatchUserListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchUserListRequest) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchUserListRequest) GetLiveId() *string {
	return s.LiveId
}

func (s *QueryLiveWatchUserListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryLiveWatchUserListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryLiveWatchUserListRequest) GetTenantContext() *QueryLiveWatchUserListRequestTenantContext {
	return s.TenantContext
}

func (s *QueryLiveWatchUserListRequest) SetLiveId(v string) *QueryLiveWatchUserListRequest {
	s.LiveId = &v
	return s
}

func (s *QueryLiveWatchUserListRequest) SetPageNumber(v int32) *QueryLiveWatchUserListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryLiveWatchUserListRequest) SetPageSize(v int32) *QueryLiveWatchUserListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryLiveWatchUserListRequest) SetTenantContext(v *QueryLiveWatchUserListRequestTenantContext) *QueryLiveWatchUserListRequest {
	s.TenantContext = v
	return s
}

func (s *QueryLiveWatchUserListRequest) Validate() error {
	return dara.Validate(s)
}

type QueryLiveWatchUserListRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryLiveWatchUserListRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchUserListRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchUserListRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryLiveWatchUserListRequestTenantContext) SetTenantId(v string) *QueryLiveWatchUserListRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryLiveWatchUserListRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
