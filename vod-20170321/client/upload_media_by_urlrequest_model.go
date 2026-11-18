// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMediaByURLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UploadMediaByURLRequest
	GetAppId() *string
	SetEnableFirstFrameCover(v bool) *UploadMediaByURLRequest
	GetEnableFirstFrameCover() *bool
	SetGenerateThumbnail(v bool) *UploadMediaByURLRequest
	GetGenerateThumbnail() *bool
	SetSessionId(v string) *UploadMediaByURLRequest
	GetSessionId() *string
	SetStorageLocation(v string) *UploadMediaByURLRequest
	GetStorageLocation() *string
	SetTemplateGroupId(v string) *UploadMediaByURLRequest
	GetTemplateGroupId() *string
	SetUploadMetadatas(v string) *UploadMediaByURLRequest
	GetUploadMetadatas() *string
	SetUploadURLs(v string) *UploadMediaByURLRequest
	GetUploadURLs() *string
	SetUserData(v string) *UploadMediaByURLRequest
	GetUserData() *string
	SetWorkflowId(v string) *UploadMediaByURLRequest
	GetWorkflowId() *string
}

type UploadMediaByURLRequest struct {
	// The application ID. Default value: **app-1000000**. For more information, see [Multi-application](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to use the first frame of the video as the cover.
	EnableFirstFrameCover *bool `json:"EnableFirstFrameCover,omitempty" xml:"EnableFirstFrameCover,omitempty"`
	// Specifies whether to generate a thumbnail for the video.
	GenerateThumbnail *bool `json:"GenerateThumbnail,omitempty" xml:"GenerateThumbnail,omitempty"`
	// The custom deduplication identifier. If this parameter is passed in a request and a request with the same identifier was detected within the past 10 minutes, an error is returned for the current request.
	//
	// >
	//
	// > - The deduplication identifier is custom. It can be up to 50 characters in length and supports uppercase and lowercase letters, digits, hyphens (-), and underscores (_). If you do not pass this parameter or pass an empty string, deduplication is not performed.
	//
	// example:
	//
	// 5c62d40299034bbaa4c195da330****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The storage address of the media file.
	//
	// Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com/?spm=a2c4g.11186623.2.15.6948257eaZ4m54#/vod/settings/censored) and choose **Configuration Management*	- > **Media Asset Management Configuration*	- > **Storage Management*	- to view the storage address. If you do not specify this parameter, the default storage address is used.
	//
	// example:
	//
	// outin-bfefbb90a47c******163e1c7426.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The ID of the transcoding template group. You can obtain the ID by using one of the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Processing Configuration*	- > **Transcoding Template Groups*	- to view the ID of the transcoding template group.
	//
	// - When you create a transcoding template group by calling the [Create a transcoding template group](https://help.aliyun.com/document_detail/102665.html) operation, the ID of the transcoding template group is the value of the TranscodeTemplateGroupId response parameter.
	//
	// - When you query transcoding template groups by calling the [Query transcoding template groups](https://help.aliyun.com/document_detail/102669.html) operation, the ID is the value of the TranscodeTemplateGroupId response parameter.
	//
	// >- If you do not specify the ID of a transcoding template group, the default transcoding template group is used for transcoding. If you specify the ID of a transcoding template group, the specified template group is used for transcoding.
	//
	// >- You can also set this parameter in `UploadMetadatas`. If TemplateGroupId is set in both `UploadMetadatas` and here, the TemplateGroupId set in `UploadMetadatas` prevails.
	//
	// example:
	//
	// ca3a8f6e4957b65806709586****
	TemplateGroupId *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	// The metadata of the media files to be uploaded. The value is a JSON string.
	//
	// - The metadata takes effect only when it matches the URL in UploadURLs.
	//
	// - JSON format: `[UploadMetadata, UploadMetadata,…]`. The value must be converted to a JSON string.
	//
	// - For more information, see the **UploadMetadata*	- table below.
	//
	// example:
	//
	// [{"SourceURL":"https://example.aliyundoc.com/video01.mp4","Title":"urlUploadTest"}]
	UploadMetadatas *string `json:"UploadMetadatas,omitempty" xml:"UploadMetadatas,omitempty"`
	// The URLs of the source media files.
	//
	// - The URL must include a file name extension. For example, in `https://****.mp4`, mp4 is the file name extension.
	//
	//     - If the URL does not contain a file name extension, you can specify it by passing `FileExtension` in `UploadMetadatas`.
	//
	//     - If the URL contains a file name extension and `FileExtension` is also passed, the passed `FileExtension` prevails.
	//
	//     - For supported file name extensions, see [Upload overview](https://help.aliyun.com/document_detail/55396.html).
	//
	// > - The URLs must be URL-encoded. Separate multiple URLs with commas (,). You can specify up to 20 URLs. To prevent upload failures caused by special characters, **URL-encode*	- the URLs before concatenating them with commas.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://****.mp4
	UploadURLs *string `json:"UploadURLs,omitempty" xml:"UploadURLs,omitempty"`
	// The custom settings. The value is a JSON string. You can configure settings such as message callbacks and upload acceleration. For more information, see [UserData](~~86952#UserData~~).
	//
	// > - To use the message callback feature in this parameter, you must configure an HTTP callback URL and select the corresponding callback event types in the console. Otherwise, the callback settings do not take effect. For more information about how to configure HTTP callbacks in the console, see [Callback settings](https://help.aliyun.com/document_detail/86071.html).
	//
	// > - To use the upload acceleration feature, you must submit a ticket to apply for activation. For more information, see [Upload instructions](https://help.aliyun.com/document_detail/55396.html). For more information about how to submit a ticket, see [Contact us](https://help.aliyun.com/document_detail/464625.html).
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://example.aliyundoc.com"},"Extend":{"localId":"xxx","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The workflow ID. You can log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Processing Configuration*	- > **Workflow Management*	- to view the workflow ID.
	//
	// > If both WorkflowId and TemplateGroupId are passed, WorkflowId prevails. For usage instructions, see [Workflows](https://help.aliyun.com/document_detail/115347.html).
	//
	// example:
	//
	// e1e243b42548248197d6f74f9****
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s UploadMediaByURLRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadMediaByURLRequest) GoString() string {
	return s.String()
}

func (s *UploadMediaByURLRequest) GetAppId() *string {
	return s.AppId
}

func (s *UploadMediaByURLRequest) GetEnableFirstFrameCover() *bool {
	return s.EnableFirstFrameCover
}

func (s *UploadMediaByURLRequest) GetGenerateThumbnail() *bool {
	return s.GenerateThumbnail
}

func (s *UploadMediaByURLRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *UploadMediaByURLRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *UploadMediaByURLRequest) GetTemplateGroupId() *string {
	return s.TemplateGroupId
}

func (s *UploadMediaByURLRequest) GetUploadMetadatas() *string {
	return s.UploadMetadatas
}

func (s *UploadMediaByURLRequest) GetUploadURLs() *string {
	return s.UploadURLs
}

func (s *UploadMediaByURLRequest) GetUserData() *string {
	return s.UserData
}

func (s *UploadMediaByURLRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *UploadMediaByURLRequest) SetAppId(v string) *UploadMediaByURLRequest {
	s.AppId = &v
	return s
}

func (s *UploadMediaByURLRequest) SetEnableFirstFrameCover(v bool) *UploadMediaByURLRequest {
	s.EnableFirstFrameCover = &v
	return s
}

func (s *UploadMediaByURLRequest) SetGenerateThumbnail(v bool) *UploadMediaByURLRequest {
	s.GenerateThumbnail = &v
	return s
}

func (s *UploadMediaByURLRequest) SetSessionId(v string) *UploadMediaByURLRequest {
	s.SessionId = &v
	return s
}

func (s *UploadMediaByURLRequest) SetStorageLocation(v string) *UploadMediaByURLRequest {
	s.StorageLocation = &v
	return s
}

func (s *UploadMediaByURLRequest) SetTemplateGroupId(v string) *UploadMediaByURLRequest {
	s.TemplateGroupId = &v
	return s
}

func (s *UploadMediaByURLRequest) SetUploadMetadatas(v string) *UploadMediaByURLRequest {
	s.UploadMetadatas = &v
	return s
}

func (s *UploadMediaByURLRequest) SetUploadURLs(v string) *UploadMediaByURLRequest {
	s.UploadURLs = &v
	return s
}

func (s *UploadMediaByURLRequest) SetUserData(v string) *UploadMediaByURLRequest {
	s.UserData = &v
	return s
}

func (s *UploadMediaByURLRequest) SetWorkflowId(v string) *UploadMediaByURLRequest {
	s.WorkflowId = &v
	return s
}

func (s *UploadMediaByURLRequest) Validate() error {
	return dara.Validate(s)
}
