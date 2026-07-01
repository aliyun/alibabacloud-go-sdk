// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaCensorJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBarrages(v string) *SubmitMediaCensorJobShrinkRequest
	GetBarrages() *string
	SetCoverImages(v string) *SubmitMediaCensorJobShrinkRequest
	GetCoverImages() *string
	SetDescription(v string) *SubmitMediaCensorJobShrinkRequest
	GetDescription() *string
	SetInputShrink(v string) *SubmitMediaCensorJobShrinkRequest
	GetInputShrink() *string
	SetNotifyUrl(v string) *SubmitMediaCensorJobShrinkRequest
	GetNotifyUrl() *string
	SetOutput(v string) *SubmitMediaCensorJobShrinkRequest
	GetOutput() *string
	SetScheduleConfigShrink(v string) *SubmitMediaCensorJobShrinkRequest
	GetScheduleConfigShrink() *string
	SetTemplateId(v string) *SubmitMediaCensorJobShrinkRequest
	GetTemplateId() *string
	SetTitle(v string) *SubmitMediaCensorJobShrinkRequest
	GetTitle() *string
	SetUserData(v string) *SubmitMediaCensorJobShrinkRequest
	GetUserData() *string
}

type SubmitMediaCensorJobShrinkRequest struct {
	// The video barrages (on-screen comments).
	//
	// > If specified, it overrides the barrages specified in the Media object.
	//
	// example:
	//
	// hello world
	Barrages *string `json:"Barrages,omitempty" xml:"Barrages,omitempty"`
	// The Object Storage Service (OSS) files for the cover images, specified as a JSON array. You can specify up to five cover images.
	//
	// > If specified, this parameter overrides the cover image information in the **Media*	- object.
	//
	// example:
	//
	// [{"Bucket":"example-bucket-****","Location":"oss-cn-shanghai","Object":"example-****.jpeg","RoleArn":"acs:ram::1997018457688683:role/AliyunICEDefaultRole"}]
	CoverImages *string `json:"CoverImages,omitempty" xml:"CoverImages,omitempty"`
	// The video description. The maximum length is 128 bytes.
	//
	// > If specified, this parameter overrides the description specified in the Media object.
	//
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The input file to censor.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The callback path. Both Message Service (MNS) and HTTP callbacks are supported.
	//
	// example:
	//
	// mns://125340688170****.oss-cn-shanghai.aliyuncs.com/queues/example-pipeline
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// The output location for screenshots. The censor job generates screenshots and a result JSON file in the OSS location specified by this parameter.
	//
	// - Example format: `oss://bucket/snapshot-{Count}.jpg`, where `bucket` is the name of an OSS bucket in the same region as the project, and `{Count}` is a placeholder for the screenshot sequence number.
	//
	// - The detailed censor results are saved to a file named `{jobId}.output` in the same OSS folder as the value of `Output`. For information about the fields in the output file, see [Media censor result file fields](https://help.aliyun.com/document_detail/609211.html).
	//
	// example:
	//
	// oss://sashimi-cn-shanghai/censor/snapshot-{Count}.jpg
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The scheduling configuration.
	ScheduleConfigShrink *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
	// The template ID. If this parameter is left empty, the service uses the default template for the censor job.
	//
	// example:
	//
	// S00000001-100060
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The video title. The maximum length is 64 bytes.
	//
	// > If specified, this parameter overrides the title specified in the Media object.
	//
	// example:
	//
	// Hello World
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The user-defined data. The maximum length is 128 bytes.
	//
	// example:
	//
	// UserDatatestid-001-****
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitMediaCensorJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaCensorJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitMediaCensorJobShrinkRequest) GetBarrages() *string {
	return s.Barrages
}

func (s *SubmitMediaCensorJobShrinkRequest) GetCoverImages() *string {
	return s.CoverImages
}

func (s *SubmitMediaCensorJobShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitMediaCensorJobShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *SubmitMediaCensorJobShrinkRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *SubmitMediaCensorJobShrinkRequest) GetOutput() *string {
	return s.Output
}

func (s *SubmitMediaCensorJobShrinkRequest) GetScheduleConfigShrink() *string {
	return s.ScheduleConfigShrink
}

func (s *SubmitMediaCensorJobShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitMediaCensorJobShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitMediaCensorJobShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitMediaCensorJobShrinkRequest) SetBarrages(v string) *SubmitMediaCensorJobShrinkRequest {
	s.Barrages = &v
	return s
}

func (s *SubmitMediaCensorJobShrinkRequest) SetCoverImages(v string) *SubmitMediaCensorJobShrinkRequest {
	s.CoverImages = &v
	return s
}

func (s *SubmitMediaCensorJobShrinkRequest) SetDescription(v string) *SubmitMediaCensorJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *SubmitMediaCensorJobShrinkRequest) SetInputShrink(v string) *SubmitMediaCensorJobShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *SubmitMediaCensorJobShrinkRequest) SetNotifyUrl(v string) *SubmitMediaCensorJobShrinkRequest {
	s.NotifyUrl = &v
	return s
}

func (s *SubmitMediaCensorJobShrinkRequest) SetOutput(v string) *SubmitMediaCensorJobShrinkRequest {
	s.Output = &v
	return s
}

func (s *SubmitMediaCensorJobShrinkRequest) SetScheduleConfigShrink(v string) *SubmitMediaCensorJobShrinkRequest {
	s.ScheduleConfigShrink = &v
	return s
}

func (s *SubmitMediaCensorJobShrinkRequest) SetTemplateId(v string) *SubmitMediaCensorJobShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitMediaCensorJobShrinkRequest) SetTitle(v string) *SubmitMediaCensorJobShrinkRequest {
	s.Title = &v
	return s
}

func (s *SubmitMediaCensorJobShrinkRequest) SetUserData(v string) *SubmitMediaCensorJobShrinkRequest {
	s.UserData = &v
	return s
}

func (s *SubmitMediaCensorJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
