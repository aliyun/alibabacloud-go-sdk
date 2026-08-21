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
	EnableFirstFrameCover *bool `json:"EnableFirstFrameCover,omitempty" xml:"EnableFirstFrameCover,omitempty"`
	GenerateThumbnail     *bool `json:"GenerateThumbnail,omitempty" xml:"GenerateThumbnail,omitempty"`
	// The metadata of the media assets to register. The value is a JSON string. You can specify metadata for up to 10 media assets at a time. For more information about the parameter structure, see the **RegisterMetadata*	- table below.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"FileURL":"https://****.oss-cn-shanghai.aliyuncs.com/video/test/video123.m3u8","Title":"VideoName"}]
	RegisterMetadatas *string `json:"RegisterMetadatas,omitempty" xml:"RegisterMetadatas,omitempty"`
	// The transcoding template group ID. You can obtain the ID by using one of the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Processing*	- > **Transcoding Template Groups*	- to view the transcoding template group ID.
	//
	// - Obtain the value of TranscodeTemplateGroupId from the response when you call the [CreateTranscodeTemplateGroup](https://help.aliyun.com/document_detail/102665.html) operation.
	//
	// - Obtain the value of TranscodeTemplateGroupId from the response when you call the [ListTranscodeTemplateGroup](https://help.aliyun.com/document_detail/102669.html) operation.
	//
	// > - If transcoding is not required, set this parameter to VOD_NO_TRANSCODE (the no-transcoding template group). Otherwise, the video status is **UploadSucc*	- and the video cannot be played by using the playback service. If transcoding is required, specify the corresponding transcoding template group ID.
	//
	// > - If both WorkflowId and TemplateGroupId are specified, WorkflowId takes precedence. For more information, see [Workflows](https://help.aliyun.com/document_detail/115347.html).
	//
	// > - This parameter triggers an [asynchronous task](https://help.aliyun.com/document_detail/3027551.html). After submission, the task enters a background queue for asynchronous execution.
	//
	// example:
	//
	// ca3a8f6e49c87b65806709586****
	TemplateGroupId *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	// The custom settings. The value is a JSON string that supports settings such as message callbacks. For more information, see [UserData](~~86952#section_6fg_qll_v3w~~).
	//
	// >This operation does not support callbacks. Even if you configure a message callback in this parameter, no callback message is generated after media asset registration is complete. When you subsequently initiate media processing such as transcoding or snapshotting on the registered media asset, if you specify a message callback in UserData at that time, that callback URL takes precedence. Otherwise, the callback URL specified in UserData during media asset registration is used.
	//
	// example:
	//
	// {"Extend":{"localId":"****","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The workflow ID. Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Processing*	- > **Workflow Management*	- to view the workflow ID.
	//
	// > - If both WorkflowId and TemplateGroupId are specified, WorkflowId takes precedence. For more information, see [Workflows](https://help.aliyun.com/document_detail/115347.html).
	//
	// > - This parameter triggers an [asynchronous task](https://help.aliyun.com/document_detail/3027551.html). After submission, the task enters a background queue for asynchronous execution.
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
