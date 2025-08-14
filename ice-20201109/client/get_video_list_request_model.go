// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCateId(v int64) *GetVideoListRequest
	GetCateId() *int64
	SetEndTime(v string) *GetVideoListRequest
	GetEndTime() *string
	SetPageNo(v int32) *GetVideoListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetVideoListRequest
	GetPageSize() *int32
	SetSortBy(v string) *GetVideoListRequest
	GetSortBy() *string
	SetStartTime(v string) *GetVideoListRequest
	GetStartTime() *string
	SetStatus(v string) *GetVideoListRequest
	GetStatus() *string
}

type GetVideoListRequest struct {
	// The ID of the category.
	//
	// example:
	//
	// 781111
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-01-11T12:59:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting method of the results. Valid values:
	//
	// 	- CreationTime:Desc (default): sorts results in reverse chronological order.
	//
	// 	- CreationTime:Asc: sorts results in chronological order.
	//
	// example:
	//
	// CreationTime:Asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the video. You can specify multiple video statuses and separate them with commas (,).
	//
	// Valid values:
	//
	// 	- PrepareFail: The file is abnormal.
	//
	// 	- UploadFail: The video failed to be uploaded.
	//
	// 	- UploadSucc: The video is uploaded.
	//
	// 	- Transcoding: The video is being transcoded.
	//
	// 	- TranscodeFail: The video failed to be transcoded.
	//
	// 	- ProduceFail: The video failed to be produced.
	//
	// 	- Normal: The video is normal.
	//
	// 	- Uploading: The video is being uploaded.
	//
	// 	- Preparing: The file is being generated.
	//
	// 	- Blocked: The video is blocked.
	//
	// 	- checking: The video is being reviewed.
	//
	// example:
	//
	// Uploading,Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetVideoListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoListRequest) GoString() string {
	return s.String()
}

func (s *GetVideoListRequest) GetCateId() *int64 {
	return s.CateId
}

func (s *GetVideoListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetVideoListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetVideoListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetVideoListRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *GetVideoListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetVideoListRequest) GetStatus() *string {
	return s.Status
}

func (s *GetVideoListRequest) SetCateId(v int64) *GetVideoListRequest {
	s.CateId = &v
	return s
}

func (s *GetVideoListRequest) SetEndTime(v string) *GetVideoListRequest {
	s.EndTime = &v
	return s
}

func (s *GetVideoListRequest) SetPageNo(v int32) *GetVideoListRequest {
	s.PageNo = &v
	return s
}

func (s *GetVideoListRequest) SetPageSize(v int32) *GetVideoListRequest {
	s.PageSize = &v
	return s
}

func (s *GetVideoListRequest) SetSortBy(v string) *GetVideoListRequest {
	s.SortBy = &v
	return s
}

func (s *GetVideoListRequest) SetStartTime(v string) *GetVideoListRequest {
	s.StartTime = &v
	return s
}

func (s *GetVideoListRequest) SetStatus(v string) *GetVideoListRequest {
	s.Status = &v
	return s
}

func (s *GetVideoListRequest) Validate() error {
	return dara.Validate(s)
}
