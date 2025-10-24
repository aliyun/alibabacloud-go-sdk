// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCateId(v int64) *AddMediaRequest
	GetCateId() *int64
	SetCoverURL(v string) *AddMediaRequest
	GetCoverURL() *string
	SetDescription(v string) *AddMediaRequest
	GetDescription() *string
	SetFileURL(v string) *AddMediaRequest
	GetFileURL() *string
	SetInputUnbind(v bool) *AddMediaRequest
	GetInputUnbind() *bool
	SetMediaWorkflowId(v string) *AddMediaRequest
	GetMediaWorkflowId() *string
	SetMediaWorkflowUserData(v string) *AddMediaRequest
	GetMediaWorkflowUserData() *string
	SetOverrideParams(v string) *AddMediaRequest
	GetOverrideParams() *string
	SetOwnerAccount(v string) *AddMediaRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddMediaRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddMediaRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddMediaRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *AddMediaRequest
	GetTags() *string
	SetTitle(v string) *AddMediaRequest
	GetTitle() *string
}

type AddMediaRequest struct {
	// The ID of the category to which the media file belongs. The value cannot be negative.
	//
	// example:
	//
	// 123
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The URL of the thumbnail. To obtain the URL, you can log on to the **MPS console*	- and choose **Workflows*	- > **Media Buckets**. Alternatively, you can log on to the **OSS console*	- and click **My OSS Paths**.
	//
	// 	- The value can be up to 3,200 bytes in length.
	//
	// 	- The URL complies with RFC 2396 and is encoded in UTF-8, with reserved characters being percent-encoded. For more information, see [URL encoding](https://help.aliyun.com/document_detail/423796.html).
	//
	// example:
	//
	// http://bucket.oss-cn-hangzhou.aliyuncs.com/example/1.png
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The description of the media file.
	//
	// 	- The description can be up to 1,024 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// A test video
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the input file. You can obtain the URL in the MPS or OSS console. For more information, see the **Triggering and matching rule for a workflow*	- section of this topic.
	//
	// 	- Only OSS HTTP URLs are supported. Alibaba Cloud CDN URLs and HTTPS URLs are not supported.
	//
	// 	- The value can be up to 3,200 bytes in size.
	//
	// 	- The URL complies with RFC 2396 and is encoded in UTF-8. For more information, see [URL encoding](https://help.aliyun.com/document_detail/423796.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// http://bucket.oss-cn-hangzhou.aliyuncs.com/A/B/C/test.mp4
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// Specifies whether to check if the media workflow supports the specified input path. We recommend that you set this parameter to true to prevent errors that may result from invalid paths. Valid values:
	//
	// 	- **true**: checks whether the workflow supports the specified input path.
	//
	// 	- **false**: does not check whether the workflow supports the specified input path.
	//
	// example:
	//
	// false
	InputUnbind *bool `json:"InputUnbind,omitempty" xml:"InputUnbind,omitempty"`
	// The ID of the media workflow that you want to run for the media file. To query the ID of a media workflow, you can log on to the MPS console or call the [AddMediaWorkflow](https://help.aliyun.com/document_detail/44437.html) operation.
	//
	// example:
	//
	// 07da6c65da7f458997336e0de192****
	MediaWorkflowId *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	// The custom data of the media workflow.
	//
	// 	- The value can be up to 1,024 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// test
	MediaWorkflowUserData *string `json:"MediaWorkflowUserData,omitempty" xml:"MediaWorkflowUserData,omitempty"`
	// The subtitle settings that are used to overwrite the original settings.
	//
	// 	- Example 1: Use `{"WebVTTSubtitleOverrides",[{"RefActivityName":"subtitleNode","WebVTTSubtitleURL":"http://test.oss-cn-hangzhou.aliyuncs.com/example1.vtt"}]}` to overwrite the original subtitle settings during HTTP Live Streaming (HLS) packaging.
	//
	// 	- Example 2: Use `{"subtitleTransNodeName":{"InputConfig":{"Format":"stl","InputFile":{"URL":"http://subtitleBucket.oss-cn-hangzhou.aliyuncs.com/package/example/CENG.stl"}}}}` to overwrite the original subtitle settings during Dynamic Adaptive Streaming over HTTP (DASH) packaging.
	//
	// example:
	//
	// {“subtitleTransNodeName”:{“InputConfig”:{“Format”:”stl”,”InputFile”:{“URL”:”http://exampleBucket.oss-cn-hangzhou.aliyuncs.com/package/example/CENG.stl"}}}}
	OverrideParams       *string `json:"OverrideParams,omitempty" xml:"OverrideParams,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags that you want to add to the media file.
	//
	// > In MPS, each tag that is specified for a media file is independent. You can search for all the media files that have the same tags in the Media Library.
	//
	// 	- You can specify up to 16 tags for a media file. Separate multiple tags with commas (,).
	//
	// 	- Each tag can be up to 32 bytes in size
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// tag1,tag2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the media file.
	//
	// 	- The title can be up to 128 bytes in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// mytest
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s AddMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMediaRequest) GoString() string {
	return s.String()
}

func (s *AddMediaRequest) GetCateId() *int64 {
	return s.CateId
}

func (s *AddMediaRequest) GetCoverURL() *string {
	return s.CoverURL
}

func (s *AddMediaRequest) GetDescription() *string {
	return s.Description
}

func (s *AddMediaRequest) GetFileURL() *string {
	return s.FileURL
}

func (s *AddMediaRequest) GetInputUnbind() *bool {
	return s.InputUnbind
}

func (s *AddMediaRequest) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *AddMediaRequest) GetMediaWorkflowUserData() *string {
	return s.MediaWorkflowUserData
}

func (s *AddMediaRequest) GetOverrideParams() *string {
	return s.OverrideParams
}

func (s *AddMediaRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddMediaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddMediaRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddMediaRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddMediaRequest) GetTags() *string {
	return s.Tags
}

func (s *AddMediaRequest) GetTitle() *string {
	return s.Title
}

func (s *AddMediaRequest) SetCateId(v int64) *AddMediaRequest {
	s.CateId = &v
	return s
}

func (s *AddMediaRequest) SetCoverURL(v string) *AddMediaRequest {
	s.CoverURL = &v
	return s
}

func (s *AddMediaRequest) SetDescription(v string) *AddMediaRequest {
	s.Description = &v
	return s
}

func (s *AddMediaRequest) SetFileURL(v string) *AddMediaRequest {
	s.FileURL = &v
	return s
}

func (s *AddMediaRequest) SetInputUnbind(v bool) *AddMediaRequest {
	s.InputUnbind = &v
	return s
}

func (s *AddMediaRequest) SetMediaWorkflowId(v string) *AddMediaRequest {
	s.MediaWorkflowId = &v
	return s
}

func (s *AddMediaRequest) SetMediaWorkflowUserData(v string) *AddMediaRequest {
	s.MediaWorkflowUserData = &v
	return s
}

func (s *AddMediaRequest) SetOverrideParams(v string) *AddMediaRequest {
	s.OverrideParams = &v
	return s
}

func (s *AddMediaRequest) SetOwnerAccount(v string) *AddMediaRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddMediaRequest) SetOwnerId(v int64) *AddMediaRequest {
	s.OwnerId = &v
	return s
}

func (s *AddMediaRequest) SetResourceOwnerAccount(v string) *AddMediaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddMediaRequest) SetResourceOwnerId(v int64) *AddMediaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddMediaRequest) SetTags(v string) *AddMediaRequest {
	s.Tags = &v
	return s
}

func (s *AddMediaRequest) SetTitle(v string) *AddMediaRequest {
	s.Title = &v
	return s
}

func (s *AddMediaRequest) Validate() error {
	return dara.Validate(s)
}
