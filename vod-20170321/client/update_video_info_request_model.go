// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCateId(v int64) *UpdateVideoInfoRequest
	GetCateId() *int64
	SetCoverURL(v string) *UpdateVideoInfoRequest
	GetCoverURL() *string
	SetDescription(v string) *UpdateVideoInfoRequest
	GetDescription() *string
	SetTags(v string) *UpdateVideoInfoRequest
	GetTags() *string
	SetTitle(v string) *UpdateVideoInfoRequest
	GetTitle() *string
	SetUserData(v string) *UpdateVideoInfoRequest
	GetUserData() *string
	SetVideoId(v string) *UpdateVideoInfoRequest
	GetVideoId() *string
}

type UpdateVideoInfoRequest struct {
	// The category ID. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Configuration Management*	- > **Media Management*	- > **Categories*	- to view the category ID of the media file.
	//
	// 	- View the value of the CateId parameter returned by the [AddCategory](https://help.aliyun.com/document_detail/56401.html) operation that you called to create a category.
	//
	// 	- View the value of the CateId parameter returned by the [GetCategories](https://help.aliyun.com/document_detail/56406.html) operation that you called to query a category.
	//
	// example:
	//
	// 384761111
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The URL of the audio/video thumbnail.
	//
	// example:
	//
	// https://example.aliyundoc.com/****.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The description of the audio or video file.
	//
	// 	- The description can be up to 1,024 bytes in length.
	//
	// 	- The value is encoded in UTF-8.
	//
	// example:
	//
	// video description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tags of the media file.
	//
	// 	- Each tag can be up to 32 bytes in length. You can specify up to 16 tags.
	//
	// 	- Separate multiple tags with commas (,).
	//
	// 	- The value is encoded in UTF-8.
	//
	// example:
	//
	// tag1,tag2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the audio or video file.
	//
	// 	- The name cannot exceed 128 bytes.
	//
	// 	- The value is encoded in UTF-8.
	//
	// example:
	//
	// video title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Custom settings. This is a JSON string that supports message callbacks, upload acceleration, and other settings. For more information, please refer to [UserData](https://help.aliyun.com/document_detail/86952.html).
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://example.aliyundoc.com"},"Extend":{"localId":"*****","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the audio or video file. Perform the following operations to obtain the storage address:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Media Files*	- > **Audio/Video**. On the Video and Audio page, view the ID of the audio or video file. This method is applicable to files that are uploaded by using the ApsaraVideo VOD console.
	//
	// 	- Obtain the value of VideoId from the response to the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation that you called to obtain the upload URL and credential.
	//
	// 	- View the value of the VideoId parameter returned by the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation that you called to query media information after the audio or video file is uploaded.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2deda93265312baf9b0ed810d****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s UpdateVideoInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateVideoInfoRequest) GetCateId() *int64 {
	return s.CateId
}

func (s *UpdateVideoInfoRequest) GetCoverURL() *string {
	return s.CoverURL
}

func (s *UpdateVideoInfoRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateVideoInfoRequest) GetTags() *string {
	return s.Tags
}

func (s *UpdateVideoInfoRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateVideoInfoRequest) GetUserData() *string {
	return s.UserData
}

func (s *UpdateVideoInfoRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *UpdateVideoInfoRequest) SetCateId(v int64) *UpdateVideoInfoRequest {
	s.CateId = &v
	return s
}

func (s *UpdateVideoInfoRequest) SetCoverURL(v string) *UpdateVideoInfoRequest {
	s.CoverURL = &v
	return s
}

func (s *UpdateVideoInfoRequest) SetDescription(v string) *UpdateVideoInfoRequest {
	s.Description = &v
	return s
}

func (s *UpdateVideoInfoRequest) SetTags(v string) *UpdateVideoInfoRequest {
	s.Tags = &v
	return s
}

func (s *UpdateVideoInfoRequest) SetTitle(v string) *UpdateVideoInfoRequest {
	s.Title = &v
	return s
}

func (s *UpdateVideoInfoRequest) SetUserData(v string) *UpdateVideoInfoRequest {
	s.UserData = &v
	return s
}

func (s *UpdateVideoInfoRequest) SetVideoId(v string) *UpdateVideoInfoRequest {
	s.VideoId = &v
	return s
}

func (s *UpdateVideoInfoRequest) Validate() error {
	return dara.Validate(s)
}
