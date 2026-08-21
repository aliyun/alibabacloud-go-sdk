// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteImageType(v string) *DeleteImageRequest
	GetDeleteImageType() *string
	SetImageIds(v string) *DeleteImageRequest
	GetImageIds() *string
	SetImageType(v string) *DeleteImageRequest
	GetImageType() *string
	SetImageURLs(v string) *DeleteImageRequest
	GetImageURLs() *string
	SetVideoId(v string) *DeleteImageRequest
	GetVideoId() *string
}

type DeleteImageRequest struct {
	// The type of image deletion operation. Valid values:
	//
	// - **ImageURL**: deletes images based on image URLs.
	//
	// - **ImageId**: deletes images based on image IDs.
	//
	// - **VideoId**: deletes images associated with a video based on the video ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// VideoId
	DeleteImageType *string `json:"DeleteImageType,omitempty" xml:"DeleteImageType,omitempty"`
	// The image IDs. Separate multiple IDs with commas (,). A maximum of 20 IDs are supported. You can obtain image IDs by using the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Image*	- to view the IDs.
	//
	// - Obtain the IDs from the response of the [CreateUploadImage](~~CreateUploadImage~~) operation that is called to obtain the upload URL and credential.
	//
	// - Obtain the IDs from the response of the [SearchMedia](~~SearchMedia~~) operation that is called to query images.
	//
	// > This parameter is available and required only when **DeleteImageType*	- is set to **ImageId**.
	//
	// example:
	//
	// bbc65bba53fed90de118a7849****,594228cdd14b4d069fc17a8c4a****
	ImageIds *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	// The type of images associated with the video that you want to delete. Valid values:
	//
	// - **CoverSnapshot**: thumbnail snapshot.
	//
	// - **NormalSnapshot**: regular snapshot.
	//
	// - **SpriteSnapshot**: sprite snapshot.
	//
	// - **SpriteOriginSnapshot**: sprite source image.
	//
	// - **All**: all of the preceding image types. If the value is not `All`, you can specify multiple image types. Separate multiple values with commas (,).
	//
	// > This parameter is available and required only when **DeleteImageType*	- is set to **VideoId**.
	//
	// example:
	//
	// All
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The image URLs. The value is the `ImageURL` parameter returned by the [CreateUploadImage](~~CreateUploadImage~~) operation. Separate multiple URLs with commas (,). A maximum of 20 URLs are supported.
	//
	// > This parameter is available and required only when **DeleteImageType*	- is set to **ImageURL**.
	//
	// example:
	//
	// https://example.aliyundoc.com/image/default/41AE7ADABBE*****.png
	ImageURLs *string `json:"ImageURLs,omitempty" xml:"ImageURLs,omitempty"`
	// The video ID. Only a single video ID is supported. You can obtain the video ID by using the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the video ID.
	//
	// - Obtain the ID from the response of the [CreateUploadVideo](~~CreateUploadVideo~~) operation that is called to obtain the upload URL and credential.
	//
	// - Obtain the ID from the response of the [SearchMedia](~~SearchMedia~~) operation that is called to query videos.
	//
	// > This parameter is available and required only when **DeleteImageType*	- is set to **VideoId**.
	//
	// example:
	//
	// eb1861d2c9a8842340e989dd56****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s DeleteImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) GetDeleteImageType() *string {
	return s.DeleteImageType
}

func (s *DeleteImageRequest) GetImageIds() *string {
	return s.ImageIds
}

func (s *DeleteImageRequest) GetImageType() *string {
	return s.ImageType
}

func (s *DeleteImageRequest) GetImageURLs() *string {
	return s.ImageURLs
}

func (s *DeleteImageRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *DeleteImageRequest) SetDeleteImageType(v string) *DeleteImageRequest {
	s.DeleteImageType = &v
	return s
}

func (s *DeleteImageRequest) SetImageIds(v string) *DeleteImageRequest {
	s.ImageIds = &v
	return s
}

func (s *DeleteImageRequest) SetImageType(v string) *DeleteImageRequest {
	s.ImageType = &v
	return s
}

func (s *DeleteImageRequest) SetImageURLs(v string) *DeleteImageRequest {
	s.ImageURLs = &v
	return s
}

func (s *DeleteImageRequest) SetVideoId(v string) *DeleteImageRequest {
	s.VideoId = &v
	return s
}

func (s *DeleteImageRequest) Validate() error {
	return dara.Validate(s)
}
