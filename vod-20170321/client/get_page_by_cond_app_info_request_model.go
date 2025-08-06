// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageByCondAppInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppItemId(v string) *GetPageByCondAppInfoRequest
	GetAppItemId() *string
	SetAppName(v string) *GetPageByCondAppInfoRequest
	GetAppName() *string
	SetAppType(v int32) *GetPageByCondAppInfoRequest
	GetAppType() *int32
	SetBusinessType(v string) *GetPageByCondAppInfoRequest
	GetBusinessType() *string
	SetNeedTotalCount(v bool) *GetPageByCondAppInfoRequest
	GetNeedTotalCount() *bool
	SetPageNo(v int64) *GetPageByCondAppInfoRequest
	GetPageNo() *int64
	SetPageSize(v int64) *GetPageByCondAppInfoRequest
	GetPageSize() *int64
	SetPkgName(v string) *GetPageByCondAppInfoRequest
	GetPkgName() *string
	SetPkgSignature(v string) *GetPageByCondAppInfoRequest
	GetPkgSignature() *string
	SetPlatformType(v int64) *GetPageByCondAppInfoRequest
	GetPlatformType() *int64
	SetSortBy(v string) *GetPageByCondAppInfoRequest
	GetSortBy() *string
}

type GetPageByCondAppInfoRequest struct {
	AppItemId *string `json:"AppItemId,omitempty" xml:"AppItemId,omitempty"`
	AppName   *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	AppType *int32 `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// ShortVideo
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// true
	NeedTotalCount *bool `json:"NeedTotalCount,omitempty" xml:"NeedTotalCount,omitempty"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PkgName      *string `json:"PkgName,omitempty" xml:"PkgName,omitempty"`
	PkgSignature *string `json:"PkgSignature,omitempty" xml:"PkgSignature,omitempty"`
	// example:
	//
	// 1
	PlatformType *int64  `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	SortBy       *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s GetPageByCondAppInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPageByCondAppInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPageByCondAppInfoRequest) GetAppItemId() *string {
	return s.AppItemId
}

func (s *GetPageByCondAppInfoRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetPageByCondAppInfoRequest) GetAppType() *int32 {
	return s.AppType
}

func (s *GetPageByCondAppInfoRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetPageByCondAppInfoRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *GetPageByCondAppInfoRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetPageByCondAppInfoRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetPageByCondAppInfoRequest) GetPkgName() *string {
	return s.PkgName
}

func (s *GetPageByCondAppInfoRequest) GetPkgSignature() *string {
	return s.PkgSignature
}

func (s *GetPageByCondAppInfoRequest) GetPlatformType() *int64 {
	return s.PlatformType
}

func (s *GetPageByCondAppInfoRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *GetPageByCondAppInfoRequest) SetAppItemId(v string) *GetPageByCondAppInfoRequest {
	s.AppItemId = &v
	return s
}

func (s *GetPageByCondAppInfoRequest) SetAppName(v string) *GetPageByCondAppInfoRequest {
	s.AppName = &v
	return s
}

func (s *GetPageByCondAppInfoRequest) SetAppType(v int32) *GetPageByCondAppInfoRequest {
	s.AppType = &v
	return s
}

func (s *GetPageByCondAppInfoRequest) SetBusinessType(v string) *GetPageByCondAppInfoRequest {
	s.BusinessType = &v
	return s
}

func (s *GetPageByCondAppInfoRequest) SetNeedTotalCount(v bool) *GetPageByCondAppInfoRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *GetPageByCondAppInfoRequest) SetPageNo(v int64) *GetPageByCondAppInfoRequest {
	s.PageNo = &v
	return s
}

func (s *GetPageByCondAppInfoRequest) SetPageSize(v int64) *GetPageByCondAppInfoRequest {
	s.PageSize = &v
	return s
}

func (s *GetPageByCondAppInfoRequest) SetPkgName(v string) *GetPageByCondAppInfoRequest {
	s.PkgName = &v
	return s
}

func (s *GetPageByCondAppInfoRequest) SetPkgSignature(v string) *GetPageByCondAppInfoRequest {
	s.PkgSignature = &v
	return s
}

func (s *GetPageByCondAppInfoRequest) SetPlatformType(v int64) *GetPageByCondAppInfoRequest {
	s.PlatformType = &v
	return s
}

func (s *GetPageByCondAppInfoRequest) SetSortBy(v string) *GetPageByCondAppInfoRequest {
	s.SortBy = &v
	return s
}

func (s *GetPageByCondAppInfoRequest) Validate() error {
	return dara.Validate(s)
}
