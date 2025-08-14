// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppendTags(v bool) *UpdateMediaInfoRequest
	GetAppendTags() *bool
	SetBusinessType(v string) *UpdateMediaInfoRequest
	GetBusinessType() *string
	SetCateId(v int64) *UpdateMediaInfoRequest
	GetCateId() *int64
	SetCategory(v string) *UpdateMediaInfoRequest
	GetCategory() *string
	SetCoverURL(v string) *UpdateMediaInfoRequest
	GetCoverURL() *string
	SetDescription(v string) *UpdateMediaInfoRequest
	GetDescription() *string
	SetInputURL(v string) *UpdateMediaInfoRequest
	GetInputURL() *string
	SetMediaId(v string) *UpdateMediaInfoRequest
	GetMediaId() *string
	SetMediaTags(v string) *UpdateMediaInfoRequest
	GetMediaTags() *string
	SetReferenceId(v string) *UpdateMediaInfoRequest
	GetReferenceId() *string
	SetTitle(v string) *UpdateMediaInfoRequest
	GetTitle() *string
	SetUserData(v string) *UpdateMediaInfoRequest
	GetUserData() *string
}

type UpdateMediaInfoRequest struct {
	// Specifies whether to append tags. Default value: false. Valid values:
	//
	// 	- true: updates the MediaTags parameter by appending new tags.
	//
	// 	- false: updates the MediaTags parameter by overwriting existing tags with new tags.
	//
	// example:
	//
	// true
	AppendTags *bool `json:"AppendTags,omitempty" xml:"AppendTags,omitempty"`
	// The business type. Valid values:
	//
	// 	- subtitles
	//
	// 	- watermark
	//
	// 	- opening
	//
	// 	- ending
	//
	// 	- general
	//
	// example:
	//
	// video
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The category ID.
	//
	// example:
	//
	// 3048
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The category.
	//
	// 	- The value can be up to 64 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// defaultCategory
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The URL of the thumbnail.
	//
	// 	- The value can be up to 128 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.png
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The content description.
	//
	// 	- The value can be up to 1,024 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// defaultDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The input URL of the media asset in another service. The URL must be bound to the ID of the media asset in IMS. The URL cannot be modified once registered.
	//
	// For a media asset from Object Storage Service (OSS), the URL may have one of the following formats:
	//
	// 1\\. http(s)://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	//
	// 2\\. oss://example-bucket/example.mp4. This format indicates that the region in which the OSS bucket of the media asset resides is the same as the region in which OSS is activated.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// The ID of the media asset. If this parameter is left empty, you must specify the input URL of the media asset, which has been registered in the IMS content library.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The tags.
	//
	// 	- Up to 16 tags are supported.
	//
	// 	- Separate multiple tags with commas (,).
	//
	// 	- Each tag can be up to 32 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// updateTags1,updateTags2
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// The custom ID. The ID can be 6 to 64 characters in length and can contain only letters, digits, hyphens (-), and underscores (_). Make sure that the ID is unique among users.
	//
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The title.
	//
	// 	- The value can be up to 128 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// defaultTitle
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The user data. It can be up to 1,024 bytes in size.
	//
	// example:
	//
	// userData
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s UpdateMediaInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaInfoRequest) GetAppendTags() *bool {
	return s.AppendTags
}

func (s *UpdateMediaInfoRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *UpdateMediaInfoRequest) GetCateId() *int64 {
	return s.CateId
}

func (s *UpdateMediaInfoRequest) GetCategory() *string {
	return s.Category
}

func (s *UpdateMediaInfoRequest) GetCoverURL() *string {
	return s.CoverURL
}

func (s *UpdateMediaInfoRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateMediaInfoRequest) GetInputURL() *string {
	return s.InputURL
}

func (s *UpdateMediaInfoRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *UpdateMediaInfoRequest) GetMediaTags() *string {
	return s.MediaTags
}

func (s *UpdateMediaInfoRequest) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *UpdateMediaInfoRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateMediaInfoRequest) GetUserData() *string {
	return s.UserData
}

func (s *UpdateMediaInfoRequest) SetAppendTags(v bool) *UpdateMediaInfoRequest {
	s.AppendTags = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetBusinessType(v string) *UpdateMediaInfoRequest {
	s.BusinessType = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetCateId(v int64) *UpdateMediaInfoRequest {
	s.CateId = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetCategory(v string) *UpdateMediaInfoRequest {
	s.Category = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetCoverURL(v string) *UpdateMediaInfoRequest {
	s.CoverURL = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetDescription(v string) *UpdateMediaInfoRequest {
	s.Description = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetInputURL(v string) *UpdateMediaInfoRequest {
	s.InputURL = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetMediaId(v string) *UpdateMediaInfoRequest {
	s.MediaId = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetMediaTags(v string) *UpdateMediaInfoRequest {
	s.MediaTags = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetReferenceId(v string) *UpdateMediaInfoRequest {
	s.ReferenceId = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetTitle(v string) *UpdateMediaInfoRequest {
	s.Title = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetUserData(v string) *UpdateMediaInfoRequest {
	s.UserData = &v
	return s
}

func (s *UpdateMediaInfoRequest) Validate() error {
	return dara.Validate(s)
}
