// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLicensesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppItemId(v string) *GetLicensesRequest
	GetAppItemId() *string
	SetAppName(v string) *GetLicensesRequest
	GetAppName() *string
	SetBusinessType(v string) *GetLicensesRequest
	GetBusinessType() *string
	SetNeedTotalCount(v bool) *GetLicensesRequest
	GetNeedTotalCount() *bool
	SetOffset(v int64) *GetLicensesRequest
	GetOffset() *int64
	SetPageNo(v int64) *GetLicensesRequest
	GetPageNo() *int64
	SetPageSize(v int64) *GetLicensesRequest
	GetPageSize() *int64
	SetPkgName(v string) *GetLicensesRequest
	GetPkgName() *string
	SetPlatformType(v int64) *GetLicensesRequest
	GetPlatformType() *int64
}

type GetLicensesRequest struct {
	AppItemId      *string `json:"AppItemId,omitempty" xml:"AppItemId,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BusinessType   *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	NeedTotalCount *bool   `json:"NeedTotalCount,omitempty" xml:"NeedTotalCount,omitempty"`
	Offset         *int64  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	PageNo         *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PkgName        *string `json:"PkgName,omitempty" xml:"PkgName,omitempty"`
	PlatformType   *int64  `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
}

func (s GetLicensesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLicensesRequest) GoString() string {
	return s.String()
}

func (s *GetLicensesRequest) GetAppItemId() *string {
	return s.AppItemId
}

func (s *GetLicensesRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetLicensesRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetLicensesRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *GetLicensesRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *GetLicensesRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetLicensesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetLicensesRequest) GetPkgName() *string {
	return s.PkgName
}

func (s *GetLicensesRequest) GetPlatformType() *int64 {
	return s.PlatformType
}

func (s *GetLicensesRequest) SetAppItemId(v string) *GetLicensesRequest {
	s.AppItemId = &v
	return s
}

func (s *GetLicensesRequest) SetAppName(v string) *GetLicensesRequest {
	s.AppName = &v
	return s
}

func (s *GetLicensesRequest) SetBusinessType(v string) *GetLicensesRequest {
	s.BusinessType = &v
	return s
}

func (s *GetLicensesRequest) SetNeedTotalCount(v bool) *GetLicensesRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *GetLicensesRequest) SetOffset(v int64) *GetLicensesRequest {
	s.Offset = &v
	return s
}

func (s *GetLicensesRequest) SetPageNo(v int64) *GetLicensesRequest {
	s.PageNo = &v
	return s
}

func (s *GetLicensesRequest) SetPageSize(v int64) *GetLicensesRequest {
	s.PageSize = &v
	return s
}

func (s *GetLicensesRequest) SetPkgName(v string) *GetLicensesRequest {
	s.PkgName = &v
	return s
}

func (s *GetLicensesRequest) SetPlatformType(v int64) *GetLicensesRequest {
	s.PlatformType = &v
	return s
}

func (s *GetLicensesRequest) Validate() error {
	return dara.Validate(s)
}
