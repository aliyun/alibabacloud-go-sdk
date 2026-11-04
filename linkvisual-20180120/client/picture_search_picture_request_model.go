// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPictureSearchPictureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceId(v string) *PictureSearchPictureRequest
	GetAppInstanceId() *string
	SetContainPicUrl(v bool) *PictureSearchPictureRequest
	GetContainPicUrl() *bool
	SetCurrentPage(v int32) *PictureSearchPictureRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *PictureSearchPictureRequest
	GetEndTime() *int64
	SetPageSize(v int32) *PictureSearchPictureRequest
	GetPageSize() *int32
	SetPictureSearchType(v int32) *PictureSearchPictureRequest
	GetPictureSearchType() *int32
	SetSearchPicUrl(v string) *PictureSearchPictureRequest
	GetSearchPicUrl() *string
	SetStartTime(v int64) *PictureSearchPictureRequest
	GetStartTime() *int64
	SetThreshold(v float32) *PictureSearchPictureRequest
	GetThreshold() *float32
}

type PictureSearchPictureRequest struct {
	// Instance ID of the image search application instance. You can call the [QueryPictureSearchApps](https://help.aliyun.com/document_detail/206608.html) operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a502d3fbab8410e8fd4be9037c7****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// Specifies whether the response includes the image URL.
	//
	// > Currently, the response is returned regardless of whether this parameter is set to true or false.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ContainPicUrl *bool `json:"ContainPicUrl,omitempty" xml:"ContainPicUrl,omitempty"`
	// The page number of the page to return. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time of the search range. The time is a 13-digit timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1614850651299
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries to return on each page. Maximum value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search type. Valid values:
	//
	// - **1*	- (default): face
	//
	// - **0**: human body
	//
	// example:
	//
	// 1
	PictureSearchType *int32 `json:"PictureSearchType,omitempty" xml:"PictureSearchType,omitempty"`
	// The URL of the image to search for.
	//
	// This parameter is required.
	SearchPicUrl *string `json:"SearchPicUrl,omitempty" xml:"SearchPicUrl,omitempty"`
	// The start time of the search range. The time is a 13-digit timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1614297000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The search threshold. Valid values: 0 to 1.
	//
	// > Set the search threshold to a value greater than or equal to 0.3 to ensure a match with the detected object.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.5
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s PictureSearchPictureRequest) String() string {
	return dara.Prettify(s)
}

func (s PictureSearchPictureRequest) GoString() string {
	return s.String()
}

func (s *PictureSearchPictureRequest) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *PictureSearchPictureRequest) GetContainPicUrl() *bool {
	return s.ContainPicUrl
}

func (s *PictureSearchPictureRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *PictureSearchPictureRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *PictureSearchPictureRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *PictureSearchPictureRequest) GetPictureSearchType() *int32 {
	return s.PictureSearchType
}

func (s *PictureSearchPictureRequest) GetSearchPicUrl() *string {
	return s.SearchPicUrl
}

func (s *PictureSearchPictureRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *PictureSearchPictureRequest) GetThreshold() *float32 {
	return s.Threshold
}

func (s *PictureSearchPictureRequest) SetAppInstanceId(v string) *PictureSearchPictureRequest {
	s.AppInstanceId = &v
	return s
}

func (s *PictureSearchPictureRequest) SetContainPicUrl(v bool) *PictureSearchPictureRequest {
	s.ContainPicUrl = &v
	return s
}

func (s *PictureSearchPictureRequest) SetCurrentPage(v int32) *PictureSearchPictureRequest {
	s.CurrentPage = &v
	return s
}

func (s *PictureSearchPictureRequest) SetEndTime(v int64) *PictureSearchPictureRequest {
	s.EndTime = &v
	return s
}

func (s *PictureSearchPictureRequest) SetPageSize(v int32) *PictureSearchPictureRequest {
	s.PageSize = &v
	return s
}

func (s *PictureSearchPictureRequest) SetPictureSearchType(v int32) *PictureSearchPictureRequest {
	s.PictureSearchType = &v
	return s
}

func (s *PictureSearchPictureRequest) SetSearchPicUrl(v string) *PictureSearchPictureRequest {
	s.SearchPicUrl = &v
	return s
}

func (s *PictureSearchPictureRequest) SetStartTime(v int64) *PictureSearchPictureRequest {
	s.StartTime = &v
	return s
}

func (s *PictureSearchPictureRequest) SetThreshold(v float32) *PictureSearchPictureRequest {
	s.Threshold = &v
	return s
}

func (s *PictureSearchPictureRequest) Validate() error {
	return dara.Validate(s)
}
