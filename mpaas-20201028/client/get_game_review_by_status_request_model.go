// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGameReviewByStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetGameReviewByStatusRequest
	GetAppId() *string
	SetKeyword(v string) *GetGameReviewByStatusRequest
	GetKeyword() *string
	SetPageNum(v int32) *GetGameReviewByStatusRequest
	GetPageNum() *int32
	SetPageSize(v int32) *GetGameReviewByStatusRequest
	GetPageSize() *int32
	SetReviewStatus(v string) *GetGameReviewByStatusRequest
	GetReviewStatus() *string
	SetSortMode(v string) *GetGameReviewByStatusRequest
	GetSortMode() *string
	SetTenantId(v string) *GetGameReviewByStatusRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *GetGameReviewByStatusRequest
	GetWorkspaceId() *string
}

type GetGameReviewByStatusRequest struct {
	// This parameter is required.
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Keyword      *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNum      *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ReviewStatus *string `json:"ReviewStatus,omitempty" xml:"ReviewStatus,omitempty"`
	SortMode     *string `json:"SortMode,omitempty" xml:"SortMode,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetGameReviewByStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGameReviewByStatusRequest) GoString() string {
	return s.String()
}

func (s *GetGameReviewByStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetGameReviewByStatusRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetGameReviewByStatusRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetGameReviewByStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetGameReviewByStatusRequest) GetReviewStatus() *string {
	return s.ReviewStatus
}

func (s *GetGameReviewByStatusRequest) GetSortMode() *string {
	return s.SortMode
}

func (s *GetGameReviewByStatusRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetGameReviewByStatusRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetGameReviewByStatusRequest) SetAppId(v string) *GetGameReviewByStatusRequest {
	s.AppId = &v
	return s
}

func (s *GetGameReviewByStatusRequest) SetKeyword(v string) *GetGameReviewByStatusRequest {
	s.Keyword = &v
	return s
}

func (s *GetGameReviewByStatusRequest) SetPageNum(v int32) *GetGameReviewByStatusRequest {
	s.PageNum = &v
	return s
}

func (s *GetGameReviewByStatusRequest) SetPageSize(v int32) *GetGameReviewByStatusRequest {
	s.PageSize = &v
	return s
}

func (s *GetGameReviewByStatusRequest) SetReviewStatus(v string) *GetGameReviewByStatusRequest {
	s.ReviewStatus = &v
	return s
}

func (s *GetGameReviewByStatusRequest) SetSortMode(v string) *GetGameReviewByStatusRequest {
	s.SortMode = &v
	return s
}

func (s *GetGameReviewByStatusRequest) SetTenantId(v string) *GetGameReviewByStatusRequest {
	s.TenantId = &v
	return s
}

func (s *GetGameReviewByStatusRequest) SetWorkspaceId(v string) *GetGameReviewByStatusRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetGameReviewByStatusRequest) Validate() error {
	return dara.Validate(s)
}
