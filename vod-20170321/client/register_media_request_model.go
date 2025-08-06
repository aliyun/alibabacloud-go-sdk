// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterMediaRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// The metadata of the media files. The value must be a JSON string. You can specify the metadata for up to 10 media files at a time. For more information about the metadata of media files, see the **RegisterMetadata*	- section of this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"FileURL":"https://****.oss-cn-shanghai.aliyuncs.com/video/test/video123.m3u8","Title":"VideoName"}]
	RegisterMetadatas *string `json:"RegisterMetadatas,omitempty" xml:"RegisterMetadatas,omitempty"`
	// The ID of the transcoding template group. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Configuration Management*	- > **Media Processing*	- > **Transcoding Template Groups**. On the Transcoding Template Groups page, you can view the ID of the transcoding template group.
	//
	// 	- Obtain the value of the TranscodeTemplateGroupId parameter from the response to the [AddTranscodeTemplateGroup](https://help.aliyun.com/document_detail/102665.html) operation that you called to create a transcoding template group.
	//
	// 	- Obtain the value of the TranscodeTemplateGroupId parameter from the response to the [ListTranscodeTemplateGroup](https://help.aliyun.com/document_detail/102669.html) operation that you called to query transcoding template groups.
	//
	// >
	//
	// 	- If you do not need to transcode media files, set the TemplateGroupId parameter to VOD_NO_TRANSCODE. If you do not specify this configuration, errors occur on your files. If you need to transcode media files, specify the ID of the transcoding template group.
	//
	// 	- If you specify both WorkflowId and TemplateGroupId, the value of the WorkflowId parameter takes effect. For more information, see [Workflows](https://help.aliyun.com/document_detail/115347.html).
	//
	// example:
	//
	// ca3a8f6e49c87b65806709586****
	TemplateGroupId *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	// The custom settings. The value must be a JSON string. You can configure settings such as message callbacks. For more information, see [UserData](~~86952#section_6fg_qll_v3w~~).
	//
	// >  You cannot configure callbacks for this operation. No callback message is returned after the media files are registered even if you configure callback settings for this parameter. If you configure callback settings for the UserData parameter when you create media processing jobs such as transcoding and snapshot capture jobs for the media file, the callback URL that you specified is used. If you do not configure callback settings when you create media processing jobs, the callback URL that you specified for the UserData parameter when you register the media file is used.
	//
	// example:
	//
	// {"Extend":{"localId":"****","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the workflow. To view the workflow ID, perform the following steps: Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Configuration Management*	- > **Media Processing*	- > **Workflows**.
	//
	// >  If you specify both WorkflowId and TemplateGroupId, the value of WorkflowId parameter takes effect. For more information, see [Workflows](https://help.aliyun.com/document_detail/115347.html).
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
