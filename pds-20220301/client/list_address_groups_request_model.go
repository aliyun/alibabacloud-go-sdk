// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddressGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *ListAddressGroupsRequest
	GetDriveId() *string
	SetImageThumbnailProcess(v string) *ListAddressGroupsRequest
	GetImageThumbnailProcess() *string
	SetLimit(v int32) *ListAddressGroupsRequest
	GetLimit() *int32
	SetMarker(v string) *ListAddressGroupsRequest
	GetMarker() *string
	SetVideoThumbnailProcess(v string) *ListAddressGroupsRequest
	GetVideoThumbnailProcess() *string
}

type ListAddressGroupsRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The method that is used to generate a thumbnail of an image.
	//
	// example:
	//
	// image/resize,w_200
	ImageThumbnailProcess *string `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The method that is used to generate a thumbnail of a video.
	//
	// example:
	//
	// video_thumbnail_process
	VideoThumbnailProcess *string `json:"video_thumbnail_process,omitempty" xml:"video_thumbnail_process,omitempty"`
}

func (s ListAddressGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAddressGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListAddressGroupsRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *ListAddressGroupsRequest) GetImageThumbnailProcess() *string {
	return s.ImageThumbnailProcess
}

func (s *ListAddressGroupsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListAddressGroupsRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListAddressGroupsRequest) GetVideoThumbnailProcess() *string {
	return s.VideoThumbnailProcess
}

func (s *ListAddressGroupsRequest) SetDriveId(v string) *ListAddressGroupsRequest {
	s.DriveId = &v
	return s
}

func (s *ListAddressGroupsRequest) SetImageThumbnailProcess(v string) *ListAddressGroupsRequest {
	s.ImageThumbnailProcess = &v
	return s
}

func (s *ListAddressGroupsRequest) SetLimit(v int32) *ListAddressGroupsRequest {
	s.Limit = &v
	return s
}

func (s *ListAddressGroupsRequest) SetMarker(v string) *ListAddressGroupsRequest {
	s.Marker = &v
	return s
}

func (s *ListAddressGroupsRequest) SetVideoThumbnailProcess(v string) *ListAddressGroupsRequest {
	s.VideoThumbnailProcess = &v
	return s
}

func (s *ListAddressGroupsRequest) Validate() error {
	return dara.Validate(s)
}
