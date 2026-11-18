// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableFirstFrameCover(v bool) *RegisterMediaRequest
	GetEnableFirstFrameCover() *bool
	SetGenerateThumbnail(v bool) *RegisterMediaRequest
	GetGenerateThumbnail() *bool
	SetRegisterMetadatas(v string) *RegisterMediaRequest
	GetRegisterMetadatas() *string
	SetTemplateGroupId(v string) *RegisterMediaRequest
	GetTemplateGroupId() *string
	SetUserData(v string) *RegisterMediaRequest
	GetUserData() *string
	SetWorkflowId(v string) *RegisterMediaRequest
	GetWorkflowId() *string
}

type RegisterMediaRequest struct {
	// Specifies whether to use the first frame of the video as the thumbnail.
	EnableFirstFrameCover *bool `json:"EnableFirstFrameCover,omitempty" xml:"EnableFirstFrameCover,omitempty"`
	// Specifies whether to generate a thumbnail for the video.
	GenerateThumbnail *bool `json:"GenerateThumbnail,omitempty" xml:"GenerateThumbnail,omitempty"`
	// The metadata of the media assets to be registered. The value is a JSON string. You can specify the metadata for up to 10 media assets at a time. For more information about the parameter structure, see the **RegisterMetadata*	- table below.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"FileURL":"https://****.oss-cn-shanghai.aliyuncs.com/video/test/video123.m3u8","Title":"VideoName"}]
	RegisterMetadatas *string `json:"RegisterMetadatas,omitempty" xml:"RegisterMetadatas,omitempty"`
	// The ID of the transcoding template group. You can obtain the ID by using one of the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Processing Configuration*	- > **Transcoding Template Groups*	- to view the ID of the transcoding template group.
	//
	// - When you create a transcoding template group by calling the [Create a transcoding template group](https://help.aliyun.com/document_detail/102665.html) operation, the ID of the transcoding template group is the value of the TranscodeTemplateGroupId response parameter.
	//
	// - When you query transcoding template groups by calling the [Query transcoding configurations](https://help.aliyun.com/document_detail/102669.html) operation, the ID of the transcoding template group is the value of the TranscodeTemplateGroupId response parameter.
	//
	// > - If transcoding is not required, set this parameter to VOD_NO_TRANSCODE, which specifies the no-transcoding template group. Otherwise, the video status is **Upload Completed**, and you cannot use the playback service to play the audio or video. If transcoding is required, specify the ID of the corresponding transcoding template group.
	//
	// > - If both WorkflowId and TemplateGroupId are specified, WorkflowId takes precedence. For more information, see [Media workflows](https://help.aliyun.com/document_detail/115347.html).
	//
	// > - The task for this parameter is an [asynchronous task](https://help.aliyun.com/document_detail/3027551.html). After the task is submitted, it is not immediately completed but queued for asynchronous execution in the background.
	//
	// example:
	//
	// ca3a8f6e49c87b65806709586****
	TemplateGroupId *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	// The custom settings. The value is a JSON string. You can configure settings such as message callbacks. For more information about the parameter, see [UserData](~~86952#section_6fg_qll_v3w~~).
	//
	// > This operation does not support callbacks. Even if a message callback is configured in this parameter, no callback message is generated after the media asset registration is completed. When you subsequently initiate media processing such as transcoding and snapshot generation for the registered media assets, if you separately configure a message callback in UserData when initiating the transcoding or snapshot generation, that callback URL takes precedence. If no callback is separately configured, the callback URL configured in UserData during the media asset registration is used.
	//
	// example:
	//
	// {"Extend":{"localId":"****","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the media workflow. You can log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Processing Configuration*	- > **Workflow Management*	- to view the workflow ID.
	//
	// > - If both WorkflowId and TemplateGroupId are specified, WorkflowId takes precedence. For more information, see [Media workflows](https://help.aliyun.com/document_detail/115347.html).
	//
	// > - The task for this parameter is an [asynchronous task](https://help.aliyun.com/document_detail/3027551.html). After the task is submitted, it is not immediately completed but queued for asynchronous execution in the background.
	//
	// example:
	//
	// 637adc2b7ba51a83d841606f8****
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s RegisterMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterMediaRequest) GoString() string {
	return s.String()
}

func (s *RegisterMediaRequest) GetEnableFirstFrameCover() *bool {
	return s.EnableFirstFrameCover
}

func (s *RegisterMediaRequest) GetGenerateThumbnail() *bool {
	return s.GenerateThumbnail
}

func (s *RegisterMediaRequest) GetRegisterMetadatas() *string {
	return s.RegisterMetadatas
}

func (s *RegisterMediaRequest) GetTemplateGroupId() *string {
	return s.TemplateGroupId
}

func (s *RegisterMediaRequest) GetUserData() *string {
	return s.UserData
}

func (s *RegisterMediaRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *RegisterMediaRequest) SetEnableFirstFrameCover(v bool) *RegisterMediaRequest {
	s.EnableFirstFrameCover = &v
	return s
}

func (s *RegisterMediaRequest) SetGenerateThumbnail(v bool) *RegisterMediaRequest {
	s.GenerateThumbnail = &v
	return s
}

func (s *RegisterMediaRequest) SetRegisterMetadatas(v string) *RegisterMediaRequest {
	s.RegisterMetadatas = &v
	return s
}

func (s *RegisterMediaRequest) SetTemplateGroupId(v string) *RegisterMediaRequest {
	s.TemplateGroupId = &v
	return s
}

func (s *RegisterMediaRequest) SetUserData(v string) *RegisterMediaRequest {
	s.UserData = &v
	return s
}

func (s *RegisterMediaRequest) SetWorkflowId(v string) *RegisterMediaRequest {
	s.WorkflowId = &v
	return s
}

func (s *RegisterMediaRequest) Validate() error {
	return dara.Validate(s)
}
