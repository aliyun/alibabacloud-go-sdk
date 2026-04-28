// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAddressGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressLevel(v string) *SearchAddressGroupsRequest
	GetAddressLevel() *string
	SetAddressNames(v []*string) *SearchAddressGroupsRequest
	GetAddressNames() []*string
	SetBrGeoPoint(v string) *SearchAddressGroupsRequest
	GetBrGeoPoint() *string
	SetDriveId(v string) *SearchAddressGroupsRequest
	GetDriveId() *string
	SetImageThumbnailProcess(v string) *SearchAddressGroupsRequest
	GetImageThumbnailProcess() *string
	SetTlGeoPoint(v string) *SearchAddressGroupsRequest
	GetTlGeoPoint() *string
	SetVideoThumbnailProcess(v string) *SearchAddressGroupsRequest
	GetVideoThumbnailProcess() *string
}

type SearchAddressGroupsRequest struct {
	// The level of the location.
	//
	// Valid values:
	//
	// 	- country
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- province
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- city
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- district
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- township
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	AddressLevel *string `json:"address_level,omitempty" xml:"address_level,omitempty"`
	// The locations.
	AddressNames []*string `json:"address_names,omitempty" xml:"address_names,omitempty" type:"Repeated"`
	// The coordinates of the bottom right vertex of the rectangle. Set the value in the format of latitude,longitude.
	//
	// example:
	//
	// 40.121,105.2121
	BrGeoPoint *string `json:"br_geo_point,omitempty" xml:"br_geo_point,omitempty"`
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The method used to generate the thumbnail of an image.
	//
	// example:
	//
	// image/resize,w_200
	ImageThumbnailProcess *string `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
	// The coordinates of the top left vertex of the rectangle. Set the value in the format of latitude,longitude.
	//
	// example:
	//
	// 39.121,101.2121
	TlGeoPoint *string `json:"tl_geo_point,omitempty" xml:"tl_geo_point,omitempty"`
	// The method used to generate the thumbnail of a video.
	//
	// example:
	//
	// video/snapshot,t_7000,f_jpg,w_800,h_600,m_fast
	VideoThumbnailProcess *string `json:"video_thumbnail_process,omitempty" xml:"video_thumbnail_process,omitempty"`
}

func (s SearchAddressGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchAddressGroupsRequest) GoString() string {
	return s.String()
}

func (s *SearchAddressGroupsRequest) GetAddressLevel() *string {
	return s.AddressLevel
}

func (s *SearchAddressGroupsRequest) GetAddressNames() []*string {
	return s.AddressNames
}

func (s *SearchAddressGroupsRequest) GetBrGeoPoint() *string {
	return s.BrGeoPoint
}

func (s *SearchAddressGroupsRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *SearchAddressGroupsRequest) GetImageThumbnailProcess() *string {
	return s.ImageThumbnailProcess
}

func (s *SearchAddressGroupsRequest) GetTlGeoPoint() *string {
	return s.TlGeoPoint
}

func (s *SearchAddressGroupsRequest) GetVideoThumbnailProcess() *string {
	return s.VideoThumbnailProcess
}

func (s *SearchAddressGroupsRequest) SetAddressLevel(v string) *SearchAddressGroupsRequest {
	s.AddressLevel = &v
	return s
}

func (s *SearchAddressGroupsRequest) SetAddressNames(v []*string) *SearchAddressGroupsRequest {
	s.AddressNames = v
	return s
}

func (s *SearchAddressGroupsRequest) SetBrGeoPoint(v string) *SearchAddressGroupsRequest {
	s.BrGeoPoint = &v
	return s
}

func (s *SearchAddressGroupsRequest) SetDriveId(v string) *SearchAddressGroupsRequest {
	s.DriveId = &v
	return s
}

func (s *SearchAddressGroupsRequest) SetImageThumbnailProcess(v string) *SearchAddressGroupsRequest {
	s.ImageThumbnailProcess = &v
	return s
}

func (s *SearchAddressGroupsRequest) SetTlGeoPoint(v string) *SearchAddressGroupsRequest {
	s.TlGeoPoint = &v
	return s
}

func (s *SearchAddressGroupsRequest) SetVideoThumbnailProcess(v string) *SearchAddressGroupsRequest {
	s.VideoThumbnailProcess = &v
	return s
}

func (s *SearchAddressGroupsRequest) Validate() error {
	return dara.Validate(s)
}
