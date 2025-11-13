// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiRtcLicenseInfoListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLicenseItemId(v string) *GetAiRtcLicenseInfoListRequest
	GetLicenseItemId() *string
	SetNeedTotalCount(v bool) *GetAiRtcLicenseInfoListRequest
	GetNeedTotalCount() *bool
	SetPageNo(v int64) *GetAiRtcLicenseInfoListRequest
	GetPageNo() *int64
	SetPageSize(v int64) *GetAiRtcLicenseInfoListRequest
	GetPageSize() *int64
	SetStatus(v int32) *GetAiRtcLicenseInfoListRequest
	GetStatus() *int32
	SetType(v int32) *GetAiRtcLicenseInfoListRequest
	GetType() *int32
}

type GetAiRtcLicenseInfoListRequest struct {
	// example:
	//
	// 17712***
	LicenseItemId *string `json:"LicenseItemId,omitempty" xml:"LicenseItemId,omitempty"`
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
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAiRtcLicenseInfoListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiRtcLicenseInfoListRequest) GoString() string {
	return s.String()
}

func (s *GetAiRtcLicenseInfoListRequest) GetLicenseItemId() *string {
	return s.LicenseItemId
}

func (s *GetAiRtcLicenseInfoListRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *GetAiRtcLicenseInfoListRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetAiRtcLicenseInfoListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetAiRtcLicenseInfoListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *GetAiRtcLicenseInfoListRequest) GetType() *int32 {
	return s.Type
}

func (s *GetAiRtcLicenseInfoListRequest) SetLicenseItemId(v string) *GetAiRtcLicenseInfoListRequest {
	s.LicenseItemId = &v
	return s
}

func (s *GetAiRtcLicenseInfoListRequest) SetNeedTotalCount(v bool) *GetAiRtcLicenseInfoListRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *GetAiRtcLicenseInfoListRequest) SetPageNo(v int64) *GetAiRtcLicenseInfoListRequest {
	s.PageNo = &v
	return s
}

func (s *GetAiRtcLicenseInfoListRequest) SetPageSize(v int64) *GetAiRtcLicenseInfoListRequest {
	s.PageSize = &v
	return s
}

func (s *GetAiRtcLicenseInfoListRequest) SetStatus(v int32) *GetAiRtcLicenseInfoListRequest {
	s.Status = &v
	return s
}

func (s *GetAiRtcLicenseInfoListRequest) SetType(v int32) *GetAiRtcLicenseInfoListRequest {
	s.Type = &v
	return s
}

func (s *GetAiRtcLicenseInfoListRequest) Validate() error {
	return dara.Validate(s)
}
