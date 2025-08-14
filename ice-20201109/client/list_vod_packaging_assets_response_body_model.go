// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodPackagingAssetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssets(v []*ListVodPackagingAssetsResponseBodyAssets) *ListVodPackagingAssetsResponseBody
	GetAssets() []*ListVodPackagingAssetsResponseBodyAssets
	SetPageNo(v int32) *ListVodPackagingAssetsResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListVodPackagingAssetsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListVodPackagingAssetsResponseBody
	GetRequestId() *string
	SetSortBy(v string) *ListVodPackagingAssetsResponseBody
	GetSortBy() *string
	SetTotalCount(v int32) *ListVodPackagingAssetsResponseBody
	GetTotalCount() *int32
}

type ListVodPackagingAssetsResponseBody struct {
	// The VOD packaging assets.
	Assets []*ListVodPackagingAssetsResponseBodyAssets `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sorting order of the assets based on the time when they were ingested. Valid values:
	//
	// 	- desc: descending order.
	//
	// 	- asc: ascending order.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVodPackagingAssetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVodPackagingAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVodPackagingAssetsResponseBody) GetAssets() []*ListVodPackagingAssetsResponseBodyAssets {
	return s.Assets
}

func (s *ListVodPackagingAssetsResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListVodPackagingAssetsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVodPackagingAssetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVodPackagingAssetsResponseBody) GetSortBy() *string {
	return s.SortBy
}

func (s *ListVodPackagingAssetsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVodPackagingAssetsResponseBody) SetAssets(v []*ListVodPackagingAssetsResponseBodyAssets) *ListVodPackagingAssetsResponseBody {
	s.Assets = v
	return s
}

func (s *ListVodPackagingAssetsResponseBody) SetPageNo(v int32) *ListVodPackagingAssetsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListVodPackagingAssetsResponseBody) SetPageSize(v int32) *ListVodPackagingAssetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVodPackagingAssetsResponseBody) SetRequestId(v string) *ListVodPackagingAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVodPackagingAssetsResponseBody) SetSortBy(v string) *ListVodPackagingAssetsResponseBody {
	s.SortBy = &v
	return s
}

func (s *ListVodPackagingAssetsResponseBody) SetTotalCount(v int32) *ListVodPackagingAssetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVodPackagingAssetsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVodPackagingAssetsResponseBodyAssets struct {
	// The name of the VOD packaging asset.
	//
	// example:
	//
	// 30min_movie
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The time when the asset was ingested. It follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-11-21T06:45:32Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The asset description.
	//
	// example:
	//
	// movie 30min
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the packaging group.
	//
	// example:
	//
	// vod_hls
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The asset input configurations.
	Input *ListVodPackagingAssetsResponseBodyAssetsInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
}

func (s ListVodPackagingAssetsResponseBodyAssets) String() string {
	return dara.Prettify(s)
}

func (s ListVodPackagingAssetsResponseBodyAssets) GoString() string {
	return s.String()
}

func (s *ListVodPackagingAssetsResponseBodyAssets) GetAssetName() *string {
	return s.AssetName
}

func (s *ListVodPackagingAssetsResponseBodyAssets) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListVodPackagingAssetsResponseBodyAssets) GetDescription() *string {
	return s.Description
}

func (s *ListVodPackagingAssetsResponseBodyAssets) GetGroupName() *string {
	return s.GroupName
}

func (s *ListVodPackagingAssetsResponseBodyAssets) GetInput() *ListVodPackagingAssetsResponseBodyAssetsInput {
	return s.Input
}

func (s *ListVodPackagingAssetsResponseBodyAssets) SetAssetName(v string) *ListVodPackagingAssetsResponseBodyAssets {
	s.AssetName = &v
	return s
}

func (s *ListVodPackagingAssetsResponseBodyAssets) SetCreateTime(v string) *ListVodPackagingAssetsResponseBodyAssets {
	s.CreateTime = &v
	return s
}

func (s *ListVodPackagingAssetsResponseBodyAssets) SetDescription(v string) *ListVodPackagingAssetsResponseBodyAssets {
	s.Description = &v
	return s
}

func (s *ListVodPackagingAssetsResponseBodyAssets) SetGroupName(v string) *ListVodPackagingAssetsResponseBodyAssets {
	s.GroupName = &v
	return s
}

func (s *ListVodPackagingAssetsResponseBodyAssets) SetInput(v *ListVodPackagingAssetsResponseBodyAssetsInput) *ListVodPackagingAssetsResponseBodyAssets {
	s.Input = v
	return s
}

func (s *ListVodPackagingAssetsResponseBodyAssets) Validate() error {
	return dara.Validate(s)
}

type ListVodPackagingAssetsResponseBodyAssetsInput struct {
	// The URL of the media file. Only M3U8 files stored in OSS are supported.
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The input type. Only Object Storage Service (OSS) is supported.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListVodPackagingAssetsResponseBodyAssetsInput) String() string {
	return dara.Prettify(s)
}

func (s ListVodPackagingAssetsResponseBodyAssetsInput) GoString() string {
	return s.String()
}

func (s *ListVodPackagingAssetsResponseBodyAssetsInput) GetMedia() *string {
	return s.Media
}

func (s *ListVodPackagingAssetsResponseBodyAssetsInput) GetType() *string {
	return s.Type
}

func (s *ListVodPackagingAssetsResponseBodyAssetsInput) SetMedia(v string) *ListVodPackagingAssetsResponseBodyAssetsInput {
	s.Media = &v
	return s
}

func (s *ListVodPackagingAssetsResponseBodyAssetsInput) SetType(v string) *ListVodPackagingAssetsResponseBodyAssetsInput {
	s.Type = &v
	return s
}

func (s *ListVodPackagingAssetsResponseBodyAssetsInput) Validate() error {
	return dara.Validate(s)
}
