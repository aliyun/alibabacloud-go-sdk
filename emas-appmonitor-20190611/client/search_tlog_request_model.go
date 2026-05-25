// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTlogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *SearchTlogRequest
	GetAppKey() *int64
	SetBeginDate(v int64) *SearchTlogRequest
	GetBeginDate() *int64
	SetDeviceId(v string) *SearchTlogRequest
	GetDeviceId() *string
	SetEndDate(v int64) *SearchTlogRequest
	GetEndDate() *int64
	SetKeyword(v string) *SearchTlogRequest
	GetKeyword() *string
	SetLevelJson(v string) *SearchTlogRequest
	GetLevelJson() *string
	SetOs(v string) *SearchTlogRequest
	GetOs() *string
	SetPageIndex(v int32) *SearchTlogRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *SearchTlogRequest
	GetPageSize() *int32
}

type SearchTlogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 233588686
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1779206400000
	BeginDate *int64 `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ad-123-136
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1779206400000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// keyword
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// ["debug","info","warning","error"]
	LevelJson *string `json:"LevelJson,omitempty" xml:"LevelJson,omitempty"`
	// This parameter is required.
	//
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
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s SearchTlogRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchTlogRequest) GoString() string {
	return s.String()
}

func (s *SearchTlogRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *SearchTlogRequest) GetBeginDate() *int64 {
	return s.BeginDate
}

func (s *SearchTlogRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *SearchTlogRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *SearchTlogRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *SearchTlogRequest) GetLevelJson() *string {
	return s.LevelJson
}

func (s *SearchTlogRequest) GetOs() *string {
	return s.Os
}

func (s *SearchTlogRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *SearchTlogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchTlogRequest) SetAppKey(v int64) *SearchTlogRequest {
	s.AppKey = &v
	return s
}

func (s *SearchTlogRequest) SetBeginDate(v int64) *SearchTlogRequest {
	s.BeginDate = &v
	return s
}

func (s *SearchTlogRequest) SetDeviceId(v string) *SearchTlogRequest {
	s.DeviceId = &v
	return s
}

func (s *SearchTlogRequest) SetEndDate(v int64) *SearchTlogRequest {
	s.EndDate = &v
	return s
}

func (s *SearchTlogRequest) SetKeyword(v string) *SearchTlogRequest {
	s.Keyword = &v
	return s
}

func (s *SearchTlogRequest) SetLevelJson(v string) *SearchTlogRequest {
	s.LevelJson = &v
	return s
}

func (s *SearchTlogRequest) SetOs(v string) *SearchTlogRequest {
	s.Os = &v
	return s
}

func (s *SearchTlogRequest) SetPageIndex(v int32) *SearchTlogRequest {
	s.PageIndex = &v
	return s
}

func (s *SearchTlogRequest) SetPageSize(v int32) *SearchTlogRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTlogRequest) Validate() error {
	return dara.Validate(s)
}
