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
	// The category ID. You can obtain the category ID by using one of the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Asset Configuration*	- > **Category Management*	- to view the category ID.
	//
	// - Call the [Create a category](https://help.aliyun.com/document_detail/56401.html) operation to create a category. The category ID is the value of the CateId response parameter.
	//
	// - Call the [Get categories and subcategories](https://help.aliyun.com/document_detail/56406.html) operation to query categories. The category ID is the value of the CateId response parameter.
	//
	// example:
	//
	// 781111
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The end time of the creation time, which must be later than the start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time must be in UTC.
	//
	// example:
	//
	// 2017-01-11T12:59:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The current page number. Default value: **1**.
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
	// The list of custom IDs. You can specify one or more custom IDs separated by commas (,). You can specify up to 20 custom IDs.
	//
	// example:
	//
	// 123-123,1234-1234
	ReferenceIds *string `json:"ReferenceIds,omitempty" xml:"ReferenceIds,omitempty"`
	// The sorting method of the query results. Valid values:
	//
	// - **CreationTime:Desc*	- (default): Sort by creation time in descending order.
	//
	// - **CreationTime:Asc**: Sort by creation time in ascending order.
	//
	// example:
	//
	// CreationTime:Asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The start time of the creation time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time must be in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The video status. You can specify multiple statuses separated by commas (,). Valid values:
	//
	// - **Uploading**: The video is being uploaded.
	//
	// - **UploadFail**: The video failed to be uploaded.
	//
	// - **UploadSucc**: The video is uploaded.
	//
	// - **Transcoding**: The video is being transcoded.
	//
	// - **TranscodeFail**: The video failed to be transcoded.
	//
	// - **Checking**: The video is being reviewed.
	//
	// - **Blocked**: The video is blocked.
	//
	// - **Normal**: The video is normal.
	//
	// - **ProduceFail**: The video failed to be produced.
	//
	// For more information about the definitions and limits of video statuses, refer to [Status: video status](~~52839#section-p7c-jgy-070~~).
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
