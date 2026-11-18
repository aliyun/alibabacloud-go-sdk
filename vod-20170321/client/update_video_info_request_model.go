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
	SetReferenceId(v string) *UpdateVideoInfoRequest
	GetReferenceId() *string
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
	// The category ID. You can obtain the ID by using one of the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Asset Management Settings*	- > **Category Management*	- to view the category ID.
	//
	// - When you call the [AddCategory](https://help.aliyun.com/document_detail/56401.html) operation to create a category, the category ID is the value of the CateId response parameter.
	//
	// - When you call the [GetCategories](https://help.aliyun.com/document_detail/56406.html) operation to query categories, the category ID is the value of the CateId response parameter.
	//
	// example:
	//
	// 384761111
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The URL of the thumbnail for the audio or video file.
	//
	// example:
	//
	// https://example.aliyundoc.com/****.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The description of the audio or video file.
	//
	// - The description can be up to 1,024 bytes in length.
	//
	// - The description must be encoded in UTF-8.
	//
	// example:
	//
	// Alibaba Cloud VOD video description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The custom ID. The ID must be 6 to 64 characters in length and can contain only lowercase letters, uppercase letters, digits, hyphens, and underscores. The ID must be unique per user.
	//
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The tags.
	//
	// - A single tag can be up to 32 bytes in length. You can add up to 16 tags.
	//
	// - Separate multiple tags with commas (,).
	//
	// - Tags must be encoded in UTF-8.
	//
	// example:
	//
	// tag1,tag2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the audio or video file.
	//
	// - The title can be up to 128 bytes in length.
	//
	// - The title must be encoded in UTF-8.
	//
	// example:
	//
	// Alibaba Cloud VOD Video Title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The custom settings. The value is a JSON string that supports settings such as message callbacks and upload acceleration. For more information, see [UserData](https://help.aliyun.com/document_detail/86952.html).
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://example.aliyundoc.com"},"Extend":{"localId":"*****","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The audio or video file ID. You can obtain the ID by using one of the following methods:
	//
	// - For videos uploaded through the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Assets*	- > **Audio and Video*	- to view the audio or video file ID.
	//
	// - When you call the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation to obtain the upload URL and credential, the audio or video file ID is the value of the VideoId response parameter.
	//
	// - After a video is uploaded, you can call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation to query the audio or video file ID. The ID is the value of the VideoId response parameter.
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

func (s *UpdateVideoInfoRequest) GetReferenceId() *string {
	return s.ReferenceId
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

func (s *UpdateVideoInfoRequest) SetReferenceId(v string) *UpdateVideoInfoRequest {
	s.ReferenceId = &v
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
