// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiRtcAuthCodeListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLicenseItemId(v string) *GetAiRtcAuthCodeListRequest
	GetLicenseItemId() *string
	SetNeedTotalCount(v bool) *GetAiRtcAuthCodeListRequest
	GetNeedTotalCount() *bool
	SetPageNo(v int64) *GetAiRtcAuthCodeListRequest
	GetPageNo() *int64
	SetPageSize(v int64) *GetAiRtcAuthCodeListRequest
	GetPageSize() *int64
	SetStatus(v int32) *GetAiRtcAuthCodeListRequest
	GetStatus() *int32
	SetType(v int32) *GetAiRtcAuthCodeListRequest
	GetType() *int32
}

type GetAiRtcAuthCodeListRequest struct {
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

func (s GetAiRtcAuthCodeListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiRtcAuthCodeListRequest) GoString() string {
	return s.String()
}

func (s *GetAiRtcAuthCodeListRequest) GetLicenseItemId() *string {
	return s.LicenseItemId
}

func (s *GetAiRtcAuthCodeListRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *GetAiRtcAuthCodeListRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetAiRtcAuthCodeListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetAiRtcAuthCodeListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *GetAiRtcAuthCodeListRequest) GetType() *int32 {
	return s.Type
}

func (s *GetAiRtcAuthCodeListRequest) SetLicenseItemId(v string) *GetAiRtcAuthCodeListRequest {
	s.LicenseItemId = &v
	return s
}

func (s *GetAiRtcAuthCodeListRequest) SetNeedTotalCount(v bool) *GetAiRtcAuthCodeListRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *GetAiRtcAuthCodeListRequest) SetPageNo(v int64) *GetAiRtcAuthCodeListRequest {
	s.PageNo = &v
	return s
}

func (s *GetAiRtcAuthCodeListRequest) SetPageSize(v int64) *GetAiRtcAuthCodeListRequest {
	s.PageSize = &v
	return s
}

func (s *GetAiRtcAuthCodeListRequest) SetStatus(v int32) *GetAiRtcAuthCodeListRequest {
	s.Status = &v
	return s
}

func (s *GetAiRtcAuthCodeListRequest) SetType(v int32) *GetAiRtcAuthCodeListRequest {
	s.Type = &v
	return s
}

func (s *GetAiRtcAuthCodeListRequest) Validate() error {
	return dara.Validate(s)
}
