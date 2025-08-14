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
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ****12e8864746a0a398****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The material parameters of the template, in the JSON format. If TemplateId is specified, ClipsParam must also be specified. For more information, see [Create and use a regular template](https://help.aliyun.com/document_detail/445399.html) and [Create and use advanced templates](https://help.aliyun.com/document_detail/445389.html).
	ClipsParam *string `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	// The parameters for editing and production. For more information, see [EditingProduceConfig](https://help.aliyun.com/document_detail/357745.html).
	//
	// >  If no thumbnail is specified in EditingProduceConfig, the first frame of the video is used as the thumbnail.
	//
	// 	- AutoRegisterInputVodMedia: specifies whether to automatically register the ApsaraVideo VOD media assets in your timeline with IMS. Default value: true.
	//
	// 	- OutputWebmTransparentChannel: specifies whether the output video contains alpha channels. Default value: false.
	//
	// 	- CoverConfig: the custom thumbnail parameters.
	//
	// *
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
	// The metadata of the produced video, in the JSON format. For more information about the parameters, see [MediaMetadata](https://help.aliyun.com/document_detail/357745.html).
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
	// The configurations of the output file, in the JSON format. You can specify an OSS URL or a storage location in a storage bucket of ApsaraVideo VOD.
	//
	// To store the output file in OSS, you must specify MediaURL. To store the output file in ApsaraVideo VOD, you must specify StorageLocation and FileName.
	//
	// For more information, see [OutputMediaConfig](https://help.aliyun.com/document_detail/357745.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"MediaURL":"https://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4"}
	OutputMediaConfig *string `json:"OutputMediaConfig,omitempty" xml:"OutputMediaConfig,omitempty"`
	// The type of the output file. Valid values:
	//
	// 	- oss-object: OSS object in an OSS bucket.
	//
	// 	- vod-media: media asset in ApsaraVideo VOD.
	//
	// 	- S3: output file based on the Amazon Simple Storage Service (S3) protocol.
	//
	// example:
	//
	// oss-object
	OutputMediaTarget *string `json:"OutputMediaTarget,omitempty" xml:"OutputMediaTarget,omitempty"`
	// The ID of the editing project.
	//
	// > : You must specify one of ProgectId, Timeline, and TempalteId and leave the other two parameters empty.
	//
	// example:
	//
	// xxxxxfb2101cb318xxxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The metadata of the editing project, in the JSON format. For more information about the parameters, see [ProjectMetadata](https://help.aliyun.com/document_detail/357745.html).
	ProjectMetadata *string `json:"ProjectMetadata,omitempty" xml:"ProjectMetadata,omitempty"`
	// The source of the editing and production request. Valid values:
	//
	// 	- OpenAPI
	//
	// 	- AliyunConsole
	//
	// 	- WebSDK
	//
	// example:
	//
	// OPENAPI
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The template ID. The template is used to build a timeline with ease.
	//
	// > : You must specify one of ProgectId, Timeline, and TempalteId and leave the other two parameters empty. If TemplateId is specified, ClipsParam must also be specified.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Timeline   *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// The user-defined data in the JSON format, which can be up to 512 bytes in length. You can specify a custom callback URL. For more information, see [Configure a callback upon editing completion](https://help.aliyun.com/document_detail/451631.html).
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
