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
	// The method that is used to delete images. Valid values:
	//
	// 	- **ImageURL**: deletes images based on URLs.
	//
	// 	- **ImageId**: deletes images based on image IDs.
	//
	// 	- **VideoId**: deletes images associated with a video based on the video ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// VideoId
	DeleteImageType *string `json:"DeleteImageType,omitempty" xml:"DeleteImageType,omitempty"`
	// The ID of the image. You can specify up to 20 image IDs and separate them with commas (,). You can use one of the following methods to obtain the image ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Media Files*	- > **Image*	- to view the image ID.
	//
	// 	- Obtain the image ID from the response to the [CreateUploadImage](~~CreateUploadImage~~) operation that you call to obtain the upload credential and URL.
	//
	// 	- Obtain the image ID from the response to the [SearchMedia](~~SearchMedia~~) operation that you call to query images.
	//
	// >  This parameter takes effect and is required only if you set **DeleteImageType*	- to **ImageId**.
	//
	// example:
	//
	// bbc65bba53fed90de118a7849****,594228cdd14b4d069fc17a8c4a****
	ImageIds *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	// The type of images that you want to delete. The images are associated with the video. Valid values:
	//
	// 	- **CoverSnapshot**: thumbnail snapshot.
	//
	// 	- **NormalSnapshot**: regular snapshot.
	//
	// 	- **SpriteSnapshot**: sprite snapshot.
	//
	// 	- **SpriteOriginSnapshot**: sprite source snapshot.
	//
	// 	- **All**: images of all the preceding types. You can specify multiple types other than `All` for this parameter. Separate multiple types with commas (,).
	//
	// >  This parameter takes effect and is required only if you set **DeleteImageType*	- to **VideoId**.
	//
	// example:
	//
	// All
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The URL of the image. You can obtain the value of `ImageURL` from the response to the [CreateUploadImage](~~CreateUploadImage~~) operation. You can specify up to 20 URLs and separate them with commas (,).
	//
	// >  This parameter takes effect and is required only if you set **DeleteImageType*	- to **ImageURL**.
	//
	// example:
	//
	// https://example.aliyundoc.com/image/default/41AE7ADABBE*****.png
	ImageURLs *string `json:"ImageURLs,omitempty" xml:"ImageURLs,omitempty"`
	// The ID of the video. You can specify only one ID. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Media Files*	- > **Audio/Video**. On the Video and Audio page, view the ID of the media file.
	//
	// 	- Obtain the video ID from the response to the [CreateUploadVideo](~~CreateUploadVideo~~) operation that you call to obtain the upload credential and URL.
	//
	// 	- Obtain the video ID from the response to the [SearchMedia](~~SearchMedia~~) operation that you call to query videos.
	//
	// >  This parameter takes effect and is required only if you set **DeleteImageType*	- to **VideoId**.
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
