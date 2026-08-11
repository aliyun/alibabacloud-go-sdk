// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaProducingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SubmitMediaProducingJobRequest
	GetClientToken() *string
	SetClipsParam(v string) *SubmitMediaProducingJobRequest
	GetClipsParam() *string
	SetEditingProduceConfig(v string) *SubmitMediaProducingJobRequest
	GetEditingProduceConfig() *string
	SetMediaMetadata(v string) *SubmitMediaProducingJobRequest
	GetMediaMetadata() *string
	SetOutputMediaConfig(v string) *SubmitMediaProducingJobRequest
	GetOutputMediaConfig() *string
	SetOutputMediaTarget(v string) *SubmitMediaProducingJobRequest
	GetOutputMediaTarget() *string
	SetProjectId(v string) *SubmitMediaProducingJobRequest
	GetProjectId() *string
	SetProjectMetadata(v string) *SubmitMediaProducingJobRequest
	GetProjectMetadata() *string
	SetSource(v string) *SubmitMediaProducingJobRequest
	GetSource() *string
	SetTemplateId(v string) *SubmitMediaProducingJobRequest
	GetTemplateId() *string
	SetTimeline(v string) *SubmitMediaProducingJobRequest
	GetTimeline() *string
	SetUserData(v string) *SubmitMediaProducingJobRequest
	GetUserData() *string
}

type SubmitMediaProducingJobRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ****12e8864746a0a398****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The material parameters corresponding to the template, in JSON format. When TemplateId is not empty, ClipsParam cannot be empty. For the specific format, see [Create and use a standard template](https://help.aliyun.com/document_detail/445399.html) and [Create and use an advanced template](https://help.aliyun.com/document_detail/445389.html).
	//
	// example:
	//
	// See the template user guide.
	ClipsParam *string `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	// The editing and compositing configuration. For more information, see [EditingProduceConfig parameter details](~~357745#section-8a4-pb2-hkv~~).
	//
	// >
	//
	// >If no cover image is configured in EditingProduceConfig, the first frame of the video is used as the cover by default.
	//
	// - AutoRegisterInputVodMedia: specifies whether to automatically register VOD media assets in your timeline to IMS. Default value: true.
	//
	// - OutputWebmTransparentChannel: specifies whether to output video with a transparent channel. Default value: false.
	//
	// - CoverConfig: custom cover image parameters.
	//
	// - ......
	//
	// example:
	//
	// {
	//
	//       "AutoRegisterInputVodMedia": "true",
	//
	//       "OutputWebmTransparentChannel": "true"
	//
	// }
	EditingProduceConfig *string `json:"EditingProduceConfig,omitempty" xml:"EditingProduceConfig,omitempty"`
	// The metadata of the produced video, in JSON format. For the specific structure definition, see [MediaMetadata](~~357745#97ff26d0e3c28~~).
	//
	// example:
	//
	// {
	//
	//       "Title":"test-title",
	//
	//       "Tags":"test-tags1,tags2"
	//
	// }
	MediaMetadata *string `json:"MediaMetadata,omitempty" xml:"MediaMetadata,omitempty"`
	// The target configuration of the output media, in JSON format. You can set the OSS URL or the storage location in a VOD bucket for the output media.
	//
	// - When outputting to OSS, the MediaURL of the output target is required.
	//
	// - When outputting to VOD, the StorageLocation and FileName parameters are required.
	//
	// [OutputMediaConfig parameter examples](~~357745#title-4j6-ve7-g31~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"MediaURL":"https://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4"}
	OutputMediaConfig *string `json:"OutputMediaConfig,omitempty" xml:"OutputMediaConfig,omitempty"`
	// The target type of the output media. Valid values:
	//
	// - oss-object: an OSS object in your Alibaba Cloud OSS bucket.
	//
	// - vod-media: a media asset in ApsaraVideo VOD.
	//
	// - S3: output using the S3 protocol.
	//
	// example:
	//
	// oss-object
	OutputMediaTarget *string `json:"OutputMediaTarget,omitempty" xml:"OutputMediaTarget,omitempty"`
	// The editing project ID. You can call the [CreateEditingProject](https://help.aliyun.com/document_detail/441137.html) operation to create an editing project and obtain the ProjectId to submit an editing task.
	//
	// 	Notice: You must specify one of the following three parameters: ProjectId, Timeline, or TemplateId. Leave the other two parameters empty.
	//
	// example:
	//
	// xxxxxfb2101cb318xxxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The metadata of the editing project, in JSON format. For the specific structure definition, see [ProjectMetadata](~~357745#title-yvp-81k-wff~~).
	//
	// example:
	//
	// {"Description":"Video editing description","Title":"Editing title test"}
	ProjectMetadata *string `json:"ProjectMetadata,omitempty" xml:"ProjectMetadata,omitempty"`
	// The source of the editing and compositing request. Valid values:
	//
	// - OpenAPI: a direct API request.
	//
	// - AliyunConsole: a request from the Alibaba Cloud Management Console.
	//
	// - WebSDK: a request from a frontend page integrated with WebSDK.
	//
	// example:
	//
	// OPENAPI
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The template ID, which is used to quickly build a timeline with minimal effort. Video clip editing based on both standard templates and advanced templates is supported.
	//
	// - When you commit a media producing job by using a template ID, you must provide the ClipsParam parameter to flexibly adjust or replace materials in the template.
	//
	// - You can invoke [GetTemplate](https://help.aliyun.com/document_detail/441164.html) to obtain template information.
	//
	// 	Notice: You must specify one of the following three parameters: ProjectId, Timeline, or TemplateId. Leave the other two parameters empty.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The timeline of the cloud editing task. When you need to arrange materials and design effects based on your video creative ideas, you can manually construct the Timeline parameter.
	//
	// - A timeline mainly contains three types of objects: tracks, materials, and effects. For more information, see [Timeline configuration](https://help.aliyun.com/document_detail/198823.html).
	//
	// - For more timeline configuration examples, see [Best Practices](https://help.aliyun.com/document_detail/2766669.html).
	//
	// 	Notice: You must specify one of the following three parameters: ProjectId, Timeline, or TemplateId. Leave the other two parameters empty.
	//
	// example:
	//
	// {"VideoTracks":[{"VideoTrackClips":[{"MediaId":"****4d7cf14dc7b83b0e801c****"},{"MediaId":"****4d7cf14dc7b83b0e801c****"}]}]}
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// Custom settings, in JSON format, with a maximum length of 512 bytes. Supports [task completion callback configuration](https://help.aliyun.com/document_detail/451631.html). The fields include:
	//
	// - NotifyAddress: the callback URL for task completion.
	//
	// - RegisterMediaNotifyAddress: the callback URL for media asset analysis completion.
	//
	// example:
	//
	// {"NotifyAddress":"https://xx.com/xx","RegisterMediaNotifyAddress":"https://xxx.com/xx"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitMediaProducingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaProducingJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitMediaProducingJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SubmitMediaProducingJobRequest) GetClipsParam() *string {
	return s.ClipsParam
}

func (s *SubmitMediaProducingJobRequest) GetEditingProduceConfig() *string {
	return s.EditingProduceConfig
}

func (s *SubmitMediaProducingJobRequest) GetMediaMetadata() *string {
	return s.MediaMetadata
}

func (s *SubmitMediaProducingJobRequest) GetOutputMediaConfig() *string {
	return s.OutputMediaConfig
}

func (s *SubmitMediaProducingJobRequest) GetOutputMediaTarget() *string {
	return s.OutputMediaTarget
}

func (s *SubmitMediaProducingJobRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *SubmitMediaProducingJobRequest) GetProjectMetadata() *string {
	return s.ProjectMetadata
}

func (s *SubmitMediaProducingJobRequest) GetSource() *string {
	return s.Source
}

func (s *SubmitMediaProducingJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitMediaProducingJobRequest) GetTimeline() *string {
	return s.Timeline
}

func (s *SubmitMediaProducingJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitMediaProducingJobRequest) SetClientToken(v string) *SubmitMediaProducingJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetClipsParam(v string) *SubmitMediaProducingJobRequest {
	s.ClipsParam = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetEditingProduceConfig(v string) *SubmitMediaProducingJobRequest {
	s.EditingProduceConfig = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetMediaMetadata(v string) *SubmitMediaProducingJobRequest {
	s.MediaMetadata = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetOutputMediaConfig(v string) *SubmitMediaProducingJobRequest {
	s.OutputMediaConfig = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetOutputMediaTarget(v string) *SubmitMediaProducingJobRequest {
	s.OutputMediaTarget = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetProjectId(v string) *SubmitMediaProducingJobRequest {
	s.ProjectId = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetProjectMetadata(v string) *SubmitMediaProducingJobRequest {
	s.ProjectMetadata = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetSource(v string) *SubmitMediaProducingJobRequest {
	s.Source = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetTemplateId(v string) *SubmitMediaProducingJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetTimeline(v string) *SubmitMediaProducingJobRequest {
	s.Timeline = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetUserData(v string) *SubmitMediaProducingJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) Validate() error {
	return dara.Validate(s)
}
