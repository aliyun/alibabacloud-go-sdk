// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrgLiveListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCorpId(v string) *GetOrgLiveListShrinkRequest
	GetCorpId() *string
	SetEndTime(v int64) *GetOrgLiveListShrinkRequest
	GetEndTime() *int64
	SetPageNumber(v int64) *GetOrgLiveListShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetOrgLiveListShrinkRequest
	GetPageSize() *int64
	SetStartTime(v int64) *GetOrgLiveListShrinkRequest
	GetStartTime() *int64
	SetTenantContextShrink(v string) *GetOrgLiveListShrinkRequest
	GetTenantContextShrink() *string
	SetUserId(v string) *GetOrgLiveListShrinkRequest
	GetUserId() *string
}

type GetOrgLiveListShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// db4d318xxxxx
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// 1720211800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1719211800000
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetOrgLiveListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOrgLiveListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListShrinkRequest) GetCorpId() *string {
	return s.CorpId
}

func (s *GetOrgLiveListShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetOrgLiveListShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetOrgLiveListShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetOrgLiveListShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetOrgLiveListShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetOrgLiveListShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetOrgLiveListShrinkRequest) SetCorpId(v string) *GetOrgLiveListShrinkRequest {
	s.CorpId = &v
	return s
}

func (s *GetOrgLiveListShrinkRequest) SetEndTime(v int64) *GetOrgLiveListShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetOrgLiveListShrinkRequest) SetPageNumber(v int64) *GetOrgLiveListShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *GetOrgLiveListShrinkRequest) SetPageSize(v int64) *GetOrgLiveListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetOrgLiveListShrinkRequest) SetStartTime(v int64) *GetOrgLiveListShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *GetOrgLiveListShrinkRequest) SetTenantContextShrink(v string) *GetOrgLiveListShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetOrgLiveListShrinkRequest) SetUserId(v string) *GetOrgLiveListShrinkRequest {
	s.UserId = &v
	return s
}

func (s *GetOrgLiveListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
