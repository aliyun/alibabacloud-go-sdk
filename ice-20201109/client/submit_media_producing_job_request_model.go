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
	// A client-generated token that ensures request idempotence. This token must be a unique value of up to 64 ASCII characters.
	//
	// example:
	//
	// ****12e8864746a0a398****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The clip parameters that correspond to the template, in JSON format. If `TemplateId` is specified, this parameter is required. For details about the format, see [Create and use basic templates](https://help.aliyun.com/document_detail/445399.html) and [Create and use advanced templates](https://help.aliyun.com/document_detail/445389.html).
	//
	// example:
	//
	// See the template user guide.
	ClipsParam *string `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	// The parameters for the media producing job. For configuration details, see [EditingProduceConfig parameter details](~~357745#section-8a4-pb2-hkv~~).
	//
	// > If a cover is not configured in `EditingProduceConfig`, the first frame of the video is used as the default cover.
	//
	// - `AutoRegisterInputVodMedia`: Specifies whether to automatically register the VOD media assets in your timeline to IMS. Default value: true.
	//
	// - `OutputWebmTransparentChannel`: Specifies whether to output a video with a transparent channel. Default value: false.
	//
	// - `CoverConfig`: The parameters for a custom cover.
	//
	// - ...
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
	// The metadata of the output video, in JSON format. For details about the structure, see [MediaMetadata](~~357745#97ff26d0e3c28~~).
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
	// The configuration for the output media target, in JSON format. You can set the URL for the output media in OSS or the storage location in a VOD bucket.
	//
	// - When outputting to OSS, the `MediaURL` parameter is required.
	//
	// - When outputting to VOD, both the `StorageLocation` and `FileName` parameters are required.
	//
	// For more information, see [OutputMediaConfig parameter examples](~~357745#title-4j6-ve7-g31~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"MediaURL":"https://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4"}
	OutputMediaConfig *string `json:"OutputMediaConfig,omitempty" xml:"OutputMediaConfig,omitempty"`
	// The target type for the output media. Valid values:
	//
	// - `oss-object`: An object in your Alibaba Cloud OSS bucket.
	//
	// - `vod-media`: A media asset in Alibaba Cloud VOD.
	//
	// - `S3`: A destination that supports the S3 protocol.
	//
	// example:
	//
	// oss-object
	OutputMediaTarget *string `json:"OutputMediaTarget,omitempty" xml:"OutputMediaTarget,omitempty"`
	// The ID of the editing project. Call the [CreateEditingProject](https://help.aliyun.com/document_detail/441137.html) operation to create an editing project and obtain the `ProjectId` to submit a media producing job.
	//
	// 	Notice: You must specify one of the `ProjectId`, `Timeline`, or `TemplateId` parameters. The other two parameters must be left empty.
	//
	// example:
	//
	// xxxxxfb2101cb318xxxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The metadata of the editing project, in JSON format. For details about the structure, see [ProjectMetadata](~~357745#title-yvp-81k-wff~~).
	//
	// example:
	//
	// {"Description":"Video editing description","Title":"Editing title test"}
	ProjectMetadata *string `json:"ProjectMetadata,omitempty" xml:"ProjectMetadata,omitempty"`
	// The source of the media producing job request. Valid values:
	//
	// - `OpenAPI`: A request initiated through an API call.
	//
	// - `AliyunConsole`: A request that originates from the Alibaba Cloud console.
	//
	// - `WebSDK`: A request sent from a front-end page that integrates the WebSDK.
	//
	// example:
	//
	// OPENAPI
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of a template for quickly building a timeline. You can use basic and advanced templates for video editing.
	//
	// - When you submit a media producing job using a template ID, you must provide the `ClipsParam` parameter to adjust or replace clips in the template.
	//
	// - Call the [GetTemplate](https://help.aliyun.com/document_detail/441164.html) operation to obtain template information.
	//
	// 	Notice:
	//
	// You must specify one of the `ProjectId`, `Timeline`, or `TemplateId` parameters. The other two parameters must be left empty.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The timeline for the cloud editing job. To arrange clips and design effects, manually construct the `Timeline` parameter.
	//
	// - A timeline primarily consists of three types of objects: tracks, clips, and effects. For more information, see [Timeline configuration](https://help.aliyun.com/document_detail/198823.html).
	//
	// - For more examples of timeline configurations, see [Best practices](https://help.aliyun.com/document_detail/2766669.html).
	//
	// 	Notice:
	//
	// You must specify one of the `ProjectId`, `Timeline`, or `TemplateId` parameters. The other two parameters must be left empty.
	//
	// example:
	//
	// {"VideoTracks":[{"VideoTrackClips":[{"MediaId":"****4d7cf14dc7b83b0e801c****"},{"MediaId":"****4d7cf14dc7b83b0e801c****"}]}]}
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// Custom user data in JSON format. The value can be up to 512 bytes in length. This parameter supports [job completion callback configuration](https://help.aliyun.com/document_detail/451631.html). The fields include:
	//
	// - `NotifyAddress`: The callback for job completion.
	//
	// - `RegisterMediaNotifyAddress`: The callback invoked when the analysis of the output media asset is complete.
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
