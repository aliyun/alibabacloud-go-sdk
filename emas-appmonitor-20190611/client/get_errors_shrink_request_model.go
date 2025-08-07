// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErrorsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *GetErrorsShrinkRequest
	GetAppKey() *int64
	SetBizModule(v string) *GetErrorsShrinkRequest
	GetBizModule() *string
	SetDigestHash(v string) *GetErrorsShrinkRequest
	GetDigestHash() *string
	SetFilterShrink(v string) *GetErrorsShrinkRequest
	GetFilterShrink() *string
	SetOs(v string) *GetErrorsShrinkRequest
	GetOs() *string
	SetPageIndex(v int32) *GetErrorsShrinkRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *GetErrorsShrinkRequest
	GetPageSize() *int32
	SetTimeRange(v *GetErrorsShrinkRequestTimeRange) *GetErrorsShrinkRequest
	GetTimeRange() *GetErrorsShrinkRequestTimeRange
	SetUtdid(v string) *GetErrorsShrinkRequest
	GetUtdid() *string
}

type GetErrorsShrinkRequest struct {
	// appKey
	//
	// This parameter is required.
	//
	// example:
	//
	// 24780725
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// crash
	BizModule *string `json:"BizModule,omitempty" xml:"BizModule,omitempty"`
	// example:
	//
	// -3481243636390427020
	DigestHash   *string `json:"DigestHash,omitempty" xml:"DigestHash,omitempty"`
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// example:
	//
	// android
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	TimeRange *GetErrorsShrinkRequestTimeRange `json:"TimeRange,omitempty" xml:"TimeRange,omitempty" type:"Struct"`
	// utdid
	//
	// example:
	//
	// Z70g6V/MXJ8DABtD53eHzn4X
	Utdid *string `json:"Utdid,omitempty" xml:"Utdid,omitempty"`
}

func (s GetErrorsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetErrorsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetErrorsShrinkRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *GetErrorsShrinkRequest) GetBizModule() *string {
	return s.BizModule
}

func (s *GetErrorsShrinkRequest) GetDigestHash() *string {
	return s.DigestHash
}

func (s *GetErrorsShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *GetErrorsShrinkRequest) GetOs() *string {
	return s.Os
}

func (s *GetErrorsShrinkRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetErrorsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetErrorsShrinkRequest) GetTimeRange() *GetErrorsShrinkRequestTimeRange {
	return s.TimeRange
}

func (s *GetErrorsShrinkRequest) GetUtdid() *string {
	return s.Utdid
}

func (s *GetErrorsShrinkRequest) SetAppKey(v int64) *GetErrorsShrinkRequest {
	s.AppKey = &v
	return s
}

func (s *GetErrorsShrinkRequest) SetBizModule(v string) *GetErrorsShrinkRequest {
	s.BizModule = &v
	return s
}

func (s *GetErrorsShrinkRequest) SetDigestHash(v string) *GetErrorsShrinkRequest {
	s.DigestHash = &v
	return s
}

func (s *GetErrorsShrinkRequest) SetFilterShrink(v string) *GetErrorsShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *GetErrorsShrinkRequest) SetOs(v string) *GetErrorsShrinkRequest {
	s.Os = &v
	return s
}

func (s *GetErrorsShrinkRequest) SetPageIndex(v int32) *GetErrorsShrinkRequest {
	s.PageIndex = &v
	return s
}

func (s *GetErrorsShrinkRequest) SetPageSize(v int32) *GetErrorsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetErrorsShrinkRequest) SetTimeRange(v *GetErrorsShrinkRequestTimeRange) *GetErrorsShrinkRequest {
	s.TimeRange = v
	return s
}

func (s *GetErrorsShrinkRequest) SetUtdid(v string) *GetErrorsShrinkRequest {
	s.Utdid = &v
	return s
}

func (s *GetErrorsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type GetErrorsShrinkRequestTimeRange struct {
	// This parameter is required.
	//
	// example:
	//
	// 1740499200000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1739894400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetErrorsShrinkRequestTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetErrorsShrinkRequestTimeRange) GoString() string {
	return s.String()
}

func (s *GetErrorsShrinkRequestTimeRange) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetErrorsShrinkRequestTimeRange) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetErrorsShrinkRequestTimeRange) SetEndTime(v int64) *GetErrorsShrinkRequestTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetErrorsShrinkRequestTimeRange) SetStartTime(v int64) *GetErrorsShrinkRequestTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetErrorsShrinkRequestTimeRange) Validate() error {
	return dara.Validate(s)
}
