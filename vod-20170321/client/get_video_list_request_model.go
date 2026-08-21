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
	// The category ID. You can obtain the category ID by using the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Management Configuration*	- > **Category Management*	- to view the category ID.
	//
	// - Obtain the value of CateId from the response when you call the [CreateCategory](https://help.aliyun.com/document_detail/56401.html) operation.
	//
	// - Obtain the value of CateId from the response when you call the [GetCategories](https://help.aliyun.com/document_detail/56406.html) operation.
	//
	// example:
	//
	// 781111
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The end of the time range to query based on CreationTime. The end time must be later than the start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2017-01-11T12:59:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: **10**. Maximum value: **100**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of custom IDs. Specify one or more custom IDs separated by commas (,). A maximum of 20 IDs are supported.
	//
	// example:
	//
	// 123-123,1234-1234
	ReferenceIds *string `json:"ReferenceIds,omitempty" xml:"ReferenceIds,omitempty"`
	// The sorting rule of the results. Valid values:
	//
	// - **CreationTime:Desc*	- (default): sorted by creation time in descending order.
	//
	// - **CreationTime:Asc**: sorted by creation time in ascending order.
	//
	// example:
	//
	// CreationTime:Asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The beginning of the time range to query based on CreationTime (creation time). Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The video status. You can specify multiple statuses. Separate multiple statuses with commas (,). Valid values:
	//
	// - **Uploading**: The video is being uploaded.
	//
	// - **UploadFail**: The video failed to be uploaded.
	//
	// - **UploadSucc**: The video has been uploaded.
	//
	// - **Transcoding**: The video is being transcoded.
	//
	// - **TranscodeFail**: The video failed to be transcoded.
	//
	// - **Checking**: The video is being reviewed.
	//
	// - **Blocked**: The video is blocked.
	//
	// - **Normal**: The video is in a normal state.
	//
	// - **ProduceFail**: The video failed to be produced.
	//
	// For more information about video statuses and related limits, see [Status: video status](~~52839#section-p7c-jgy-070~~).
	//
	// example:
	//
	// Uploading,Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage address of the audio or video file.
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
