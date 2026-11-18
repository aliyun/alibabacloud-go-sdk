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
	SetEnableFirstFrameCover(v bool) *CreateUploadVideoRequest
	GetEnableFirstFrameCover() *bool
	SetFileName(v string) *CreateUploadVideoRequest
	GetFileName() *string
	SetFileSize(v int64) *CreateUploadVideoRequest
	GetFileSize() *int64
	SetGenerateThumbnail(v bool) *CreateUploadVideoRequest
	GetGenerateThumbnail() *bool
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
	// The application ID. Default value: **app-1000000**. For more information, see [Multi-application](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The category ID. You can obtain the category ID by using one of the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Configuration Management*	- > **Media Asset Management Configuration*	- > **Category Management*	- to view the category ID.
	//
	// - When you create a category by calling the [AddCategory](~~AddCategory~~) operation, the category ID is the value of the CateId parameter in the response.
	//
	// - When you query categories by calling the [GetCategories](~~GetCategories~~) operation, the category ID is the value of the CateId parameter in the response.
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
	// The description of the audio or video file displayed in ApsaraVideo VOD after the upload is complete.
	//
	// - The description can be up to 1,024 characters in length.
	//
	// - The description must be encoded in UTF-8.
	//
	// example:
	//
	// UploadTest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to use the first frame of the video as the thumbnail.
	EnableFirstFrameCover *bool `json:"EnableFirstFrameCover,omitempty" xml:"EnableFirstFrameCover,omitempty"`
	// The name of the audio or video source file to be uploaded.
	//
	// - The file name must contain an extension. The extension is not case-sensitive.
	//
	// - For supported extensions, see [Upload overview](https://help.aliyun.com/document_detail/55396.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// D:\\video_01.mp4
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The size of the audio or video source file to be uploaded. Unit: bytes.
	//
	// example:
	//
	// 123
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// Specifies whether to generate a thumbnail for the video.
	GenerateThumbnail *bool `json:"GenerateThumbnail,omitempty" xml:"GenerateThumbnail,omitempty"`
	// The custom ID. The ID can be 6 to 64 characters in length and can contain only lowercase letters, uppercase letters, digits, hyphens (-), and underscores (_). The ID must be unique within the user dimension.
	//
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The storage address. You can obtain the storage address by using the following method:
	//
	// Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Asset Management Configuration*	- > **Storage Management*	- to view the storage address.
	//
	// > If you do not specify this parameter, the audio or video file is uploaded to the default storage address. If no default storage address is configured, the file is uploaded to the first storage address in the storage list. If you specify this parameter, the audio or video file is uploaded to the specified storage address.
	//
	// example:
	//
	// out-****.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags of the audio or video file.
	//
	// - You can specify up to 16 tags.
	//
	// - Separate multiple tags with commas (,).
	//
	// - Each tag can be up to 32 characters in length.
	//
	// - The tags must be encoded in UTF-8.
	//
	// example:
	//
	// tag1,tag2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the transcoding template group. You can obtain the ID by using one of the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Configuration Management*	- > **Media Processing Configuration*	- > **Transcoding Template Groups*	- to view the ID of the transcoding template group.
	//
	// - When you create a transcoding template group by calling the [Create a transcoding template group](https://help.aliyun.com/document_detail/102665.html) operation, the ID is the value of the TranscodeTemplateGroupId parameter in the response.
	//
	// - When you query transcoding template groups by calling the [Query transcoding configuration list](https://help.aliyun.com/document_detail/102669.html) operation, the ID is the value of the TranscodeTemplateGroupId parameter in the response.
	//
	// >- If both WorkflowId and TemplateGroupId are specified, WorkflowId takes precedence.
	//
	// >- If you do not specify this parameter, the default transcoding template group is used. If you specify this parameter, the specified transcoding template group is used.
	//
	// >- If you set this parameter to the built-in **No Transcoding*	- template group, only the [video upload completion](https://help.aliyun.com/document_detail/55630.html) event notification is generated after the audio or video file is uploaded. The [single definition transcoding completion](https://help.aliyun.com/document_detail/55636.html) event notification is not generated.
	//
	// > - The task triggered by this parameter is an [asynchronous task](https://help.aliyun.com/document_detail/3027551.html). After the task is submitted, it is queued and executed asynchronously in the background.
	//
	// >- To ensure normal playback, if you set this parameter to the built-in **No Transcoding*	- template group, only the following formats support direct playback without transcoding after the audio or video file is uploaded: MP4, FLV, MP3, M3U8, and WEBM. Other formats are only supported for storage (pay attention to the extension of FileName). If you use ApsaraVideo Player for playback, the version must be 3.1.0 or later.
	//
	// example:
	//
	// 405477f9e214d19ea2c7c854****
	TemplateGroupId *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	// The title of the audio or video file displayed in ApsaraVideo VOD after the upload is complete.
	//
	// - The title can be up to 128 characters in length.
	//
	// - The title must be encoded in UTF-8.
	//
	// This parameter is required.
	//
	// example:
	//
	// UploadTest
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The custom settings. The value is a JSON string that supports settings such as message callback and upload acceleration. For more information, see [UserData](https://help.aliyun.com/document_detail/86952.html).
	//
	// > - To use the message callback feature, you must configure an HTTP callback URL and select the corresponding callback event types in the console. Otherwise, the callback settings do not take effect. If no callback URL is specified for subsequent tasks, notifications are pushed to this URL by default. For more information about how to configure HTTP callbacks in the console, see [Callback settings](https://help.aliyun.com/document_detail/86071.html).
	//
	// > - To use the upload acceleration feature, you must [submit an application](https://yida.alibaba-inc.com/o/ticketapply) to enable it. For more information, see [Upload instructions](https://help.aliyun.com/document_detail/55396.html).
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://example.aliyundoc.com"},"Extend":{"localId":"*****","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The workflow ID. You can log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Processing Configuration*	- > **Workflows*	- to view the workflow ID.
	//
	// > - If both WorkflowId and TemplateGroupId are specified, WorkflowId takes precedence. For more information, see [Workflows](https://help.aliyun.com/document_detail/115347.html).
	//
	// > - The task triggered by this parameter is an [asynchronous task](https://help.aliyun.com/document_detail/3027551.html). After the task is submitted, it is queued and executed asynchronously in the background.
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

func (s *CreateUploadVideoRequest) GetEnableFirstFrameCover() *bool {
	return s.EnableFirstFrameCover
}

func (s *CreateUploadVideoRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateUploadVideoRequest) GetFileSize() *int64 {
	return s.FileSize
}

func (s *CreateUploadVideoRequest) GetGenerateThumbnail() *bool {
	return s.GenerateThumbnail
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

func (s *CreateUploadVideoRequest) SetEnableFirstFrameCover(v bool) *CreateUploadVideoRequest {
	s.EnableFirstFrameCover = &v
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

func (s *CreateUploadVideoRequest) SetGenerateThumbnail(v bool) *CreateUploadVideoRequest {
	s.GenerateThumbnail = &v
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
