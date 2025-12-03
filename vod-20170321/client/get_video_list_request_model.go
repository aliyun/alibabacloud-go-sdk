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
	SetReferenceIds(v string) *GetVideoListRequest
	GetReferenceIds() *string
	SetSortBy(v string) *GetVideoListRequest
	GetSortBy() *string
	SetStartTime(v string) *GetVideoListRequest
	GetStartTime() *string
	SetStatus(v string) *GetVideoListRequest
	GetStatus() *string
	SetStorageLocation(v string) *GetVideoListRequest
	GetStorageLocation() *string
}

type GetVideoListRequest struct {
	// The ID of the category. You can use one of the following methods to obtain the category ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Configuration Management*	- > **Media Management*	- > **Categories*	- to view the category ID.
	//
	// 	- Obtain the value of CateId from the response to the [AddCategory](https://help.aliyun.com/document_detail/56401.html) operation.
	//
	// 	- Obtain the value of CateId from the response to the [GetCategories](https://help.aliyun.com/document_detail/56406.html) operation.
	//
	// example:
	//
	// 781111
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2017-01-11T12:59:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page. Default value: **10**. Maximum value: **100**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 123-123,1234-1234
	ReferenceIds *string `json:"ReferenceIds,omitempty" xml:"ReferenceIds,omitempty"`
	// The sorting method of the results. Valid values:
	//
	// 	- **CreationTime:Desc*	- (default): The results are sorted in reverse chronological order based on the creation time.
	//
	// 	- **CreationTime:Asc**: The results are sorted in chronological order based on the creation time.
	//
	// example:
	//
	// CreationTime:Asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the video. You can specify multiple video statuses and separate them with commas (,). Valid values:
	//
	// 	- **Uploading**: The video is being uploaded.
	//
	// 	- **UploadFail**: The video failed to be uploaded.
	//
	// 	- **UploadSucc**: The video has been uploaded.
	//
	// 	- **Transcoding**: The video is being transcoded.
	//
	// 	- **TranscodeFail**: The video failed to be transcoded.
	//
	// 	- **checking**: The video is being reviewed.
	//
	// 	- **Blocked**: The video is blocked.
	//
	// 	- **Normal**: The video is normal.
	//
	// 	- **ProduceFail**: The video failed to be produced.
	//
	// For more information about each video status, see the "Status: the status of a video" section of the [Basic data types](~~52839#section-p7c-jgy-070~~) topic.
	//
	// example:
	//
	// Uploading,Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage address of the media file.
	//
	// example:
	//
	// out-****.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
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

func (s *GetVideoListRequest) GetReferenceIds() *string {
	return s.ReferenceIds
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

func (s *GetVideoListRequest) GetStorageLocation() *string {
	return s.StorageLocation
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

func (s *GetVideoListRequest) SetReferenceIds(v string) *GetVideoListRequest {
	s.ReferenceIds = &v
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

func (s *GetVideoListRequest) SetStorageLocation(v string) *GetVideoListRequest {
	s.StorageLocation = &v
	return s
}

func (s *GetVideoListRequest) Validate() error {
	return dara.Validate(s)
}
