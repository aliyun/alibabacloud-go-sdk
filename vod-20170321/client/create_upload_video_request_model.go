// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadVideoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateUploadVideoRequest
	GetAppId() *string
	SetCateId(v int64) *CreateUploadVideoRequest
	GetCateId() *int64
	SetCoverURL(v string) *CreateUploadVideoRequest
	GetCoverURL() *string
	SetDescription(v string) *CreateUploadVideoRequest
	GetDescription() *string
	SetFileName(v string) *CreateUploadVideoRequest
	GetFileName() *string
	SetFileSize(v int64) *CreateUploadVideoRequest
	GetFileSize() *int64
	SetReferenceId(v string) *CreateUploadVideoRequest
	GetReferenceId() *string
	SetStorageLocation(v string) *CreateUploadVideoRequest
	GetStorageLocation() *string
	SetTags(v string) *CreateUploadVideoRequest
	GetTags() *string
	SetTemplateGroupId(v string) *CreateUploadVideoRequest
	GetTemplateGroupId() *string
	SetTitle(v string) *CreateUploadVideoRequest
	GetTitle() *string
	SetUserData(v string) *CreateUploadVideoRequest
	GetUserData() *string
	SetWorkflowId(v string) *CreateUploadVideoRequest
	GetWorkflowId() *string
}

type CreateUploadVideoRequest struct {
	// The ID of the application. Default value: **app-1000000**. For more information, see [Overview](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the category. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Configuration Management*	- > **Media Management*	- > **Categories*	- to view the category ID of the media file.
	//
	// 	- Obtain the value of CateId from the response to the [AddCategory](~~AddCategory~~) operation.
	//
	// 	- Obtain the value of CateId from the response to the [GetCategories](~~GetCategories~~) operation.
	//
	// example:
	//
	// 100036****
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The URL of the custom video thumbnail.
	//
	// example:
	//
	// https://example.aliyundoc.com/image/D22F553TEST****.jpeg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The description of the audio or video file.
	//
	// 	- The value can be up to 1,024 characters in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// UploadTest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the source file.
	//
	// 	- The name must contain a file name extension, which is not case-sensitive.
	//
	// 	- For more information about file name extensions supported by ApsaraVideo VOD, see [Overview](https://help.aliyun.com/document_detail/55396.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// D:\\video_01.mp4
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The size of the source file. Unit: bytes.
	//
	// example:
	//
	// 123
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The storage address. Perform the following operations to obtain the storage address: Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Configuration Management*	- > **Media Management*	- > **Storage**. On the Storage page, view the storage address.
	//
	// >  If you leave this parameter empty, audio and video files are uploaded to the default storage address. If you specify a storage address, audio and video files are uploaded to the specified address.
	//
	// example:
	//
	// out-****.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags of the audio or video file.
	//
	// 	- You can specify a maximum of 16 tags.
	//
	// 	- If you want to specify multiple tags, separate the tags with commas (,).
	//
	// 	- Each tag can be up to 32 characters in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// tag1,tag2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the transcoding template group. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the ApsaraVideo VOD console. In the left-side navigation pane, choose Configuration Management > Media Processing > Transcoding Template Groups. On the Transcoding Template Groups page, you can view the ID of the transcoding template group.[](https://vod.console.aliyun.com)************
	//
	// 	- Obtain the value of the TranscodeTemplateGroupId parameter from the response to the [AddTranscodeTemplateGroup](https://help.aliyun.com/document_detail/102665.html) operation that you called to create a transcoding template group.
	//
	// 	- Obtain the value of the TranscodeTemplateGroupId parameter from the response to the [ListTranscodeTemplateGroup](https://help.aliyun.com/document_detail/102669.html) operation that you called to query transcoding template groups.
	//
	// > 	- If you specify both WorkflowId and TemplateGroupId, the value of the WorkflowId parameter takes effect.
	//
	// > 	- If this parameter is not specified, transcoding is performed based on the default transcoding template group. If the transcoding template group ID is specified, transcoding is performed based on the specified template group.
	//
	// > 	- If the **No Transcoding*	- template group is used, only the [FileUploadComplete](https://help.aliyun.com/document_detail/55630.html) event notification is returned after a video is uploaded. The [StreamTranscodeComplete](https://help.aliyun.com/document_detail/55636.html) event notification is not returned.
	//
	// > 	- If you use the **No Transcoding*	- template group to upload videos, only videos in the format of MP4, FLV, MP3, M3U8, or WebM can be played. Videos in other formats can only be stored in ApsaraVideo VOD. You can view the file name extension to obtain the video format. If you want to use ApsaraVideo Player, make sure that the version of the player is V3.1.0 or later.
	//
	// example:
	//
	// 405477f9e214d19ea2c7c854****
	TemplateGroupId *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	// The title of the audio or video file.
	//
	// 	- The title can be up to 128 characters in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// This parameter is required.
	//
	// example:
	//
	// UploadTest
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The custom configurations such as callback configurations and upload acceleration configurations. The value must be a JSON string. For more information, see [Request parameters](https://help.aliyun.com/document_detail/86952.html).
	//
	// > 	- The callback configurations take effect only after you specify the HTTP callback URL and select specific callback events in the ApsaraVideo VOD console. For more information about how to configure HTTP callback settings in the ApsaraVideo VOD console, see [Configure callback settings](https://help.aliyun.com/document_detail/86071.html).
	//
	// >	- If you want to enable the upload acceleration feature, [submit a request on Yida](https://yida.alibaba-inc.com/o/ticketapply). For more information, see [Overview](https://help.aliyun.com/document_detail/55396.html).
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://example.aliyundoc.com"},"Extend":{"localId":"*****","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the workflow. To view the ID of the workflow, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Configuration Management*	- > **Media Processing*	- > **Workflows**.
	//
	// > If you specify the WorkflowId and TemplateGroupId parameters, the value of the WorkflowId parameter takes effect. For more information, see [Workflows](https://help.aliyun.com/document_detail/115347.html).
	//
	// example:
	//
	// 613efff3887ec34af685714cc461****
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s CreateUploadVideoRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadVideoRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadVideoRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateUploadVideoRequest) GetCateId() *int64 {
	return s.CateId
}

func (s *CreateUploadVideoRequest) GetCoverURL() *string {
	return s.CoverURL
}

func (s *CreateUploadVideoRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUploadVideoRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateUploadVideoRequest) GetFileSize() *int64 {
	return s.FileSize
}

func (s *CreateUploadVideoRequest) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *CreateUploadVideoRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *CreateUploadVideoRequest) GetTags() *string {
	return s.Tags
}

func (s *CreateUploadVideoRequest) GetTemplateGroupId() *string {
	return s.TemplateGroupId
}

func (s *CreateUploadVideoRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateUploadVideoRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateUploadVideoRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *CreateUploadVideoRequest) SetAppId(v string) *CreateUploadVideoRequest {
	s.AppId = &v
	return s
}

func (s *CreateUploadVideoRequest) SetCateId(v int64) *CreateUploadVideoRequest {
	s.CateId = &v
	return s
}

func (s *CreateUploadVideoRequest) SetCoverURL(v string) *CreateUploadVideoRequest {
	s.CoverURL = &v
	return s
}

func (s *CreateUploadVideoRequest) SetDescription(v string) *CreateUploadVideoRequest {
	s.Description = &v
	return s
}

func (s *CreateUploadVideoRequest) SetFileName(v string) *CreateUploadVideoRequest {
	s.FileName = &v
	return s
}

func (s *CreateUploadVideoRequest) SetFileSize(v int64) *CreateUploadVideoRequest {
	s.FileSize = &v
	return s
}

func (s *CreateUploadVideoRequest) SetReferenceId(v string) *CreateUploadVideoRequest {
	s.ReferenceId = &v
	return s
}

func (s *CreateUploadVideoRequest) SetStorageLocation(v string) *CreateUploadVideoRequest {
	s.StorageLocation = &v
	return s
}

func (s *CreateUploadVideoRequest) SetTags(v string) *CreateUploadVideoRequest {
	s.Tags = &v
	return s
}

func (s *CreateUploadVideoRequest) SetTemplateGroupId(v string) *CreateUploadVideoRequest {
	s.TemplateGroupId = &v
	return s
}

func (s *CreateUploadVideoRequest) SetTitle(v string) *CreateUploadVideoRequest {
	s.Title = &v
	return s
}

func (s *CreateUploadVideoRequest) SetUserData(v string) *CreateUploadVideoRequest {
	s.UserData = &v
	return s
}

func (s *CreateUploadVideoRequest) SetWorkflowId(v string) *CreateUploadVideoRequest {
	s.WorkflowId = &v
	return s
}

func (s *CreateUploadVideoRequest) Validate() error {
	return dara.Validate(s)
}
