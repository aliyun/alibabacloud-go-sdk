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
	// The live comments of the video.
	//
	// >  If this parameter is specified, the system checks the live comments specified by this parameter instead of the live comments of the input file specified by Media.
	//
	// example:
	//
	// hello world
	Barrages *string `json:"Barrages,omitempty" xml:"Barrages,omitempty"`
	// The Object Storage Service (OSS) objects that are used as the thumbnails. Specify the thumbnails in a JSON array. A maximum of five thumbnails are supported.
	//
	// >  If this parameter is specified, the system checks the thumbnails specified by this parameter instead of the thumbnails of the input file specified by **Media**.
	//
	// example:
	//
	// [{"Bucket":"example-bucket-****","Location":"oss-cn-shanghai","Object":"example-****.jpeg","RoleArn":"acs:ram::1997018457688683:role/AliyunICEDefaultRole"}]
	CoverImages *string `json:"CoverImages,omitempty" xml:"CoverImages,omitempty"`
	// The video description, which can be up to 128 bytes in length.
	//
	// >  If this parameter is specified, the system checks the description specified by this parameter instead of the description of the input file specified by Media.
	//
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about the file to be moderated.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The callback URL. Simple Message Queue (SMQ, formerly MNS) and HTTP callbacks are supported.
	//
	// example:
	//
	// mns://125340688170****.oss-cn-shanghai.aliyuncs.com/queues/example-pipeline
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// The output snapshots. The moderation job generates output snapshots and the result JSON file in the path corresponding to the input file.
	//
	// 	- File name format of output snapshots: oss://bucket/snapshot-{Count}.jpg. In the path, bucket indicates an OSS bucket that resides in the same region as the current project, and {Count} is the sequence number of the snapshot.
	//
	// 	- The detailed moderation results are stored in the {jobId}.output file in the same OSS folder as the output snapshots. For more information about the parameters in the output file, see [Output parameters of media moderation jobs](https://help.aliyun.com/document_detail/609211.html).
	//
	// example:
	//
	// oss://sashimi-cn-shanghai/censor/snapshot-{Count}.jpg
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The scheduling configurations.
	ScheduleConfigShrink *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
	// The template ID. If this parameter is not specified, the default template is used for moderation.
	//
	// example:
	//
	// S00000001-100060
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The video title, which can be up to 64 bytes in length.
	//
	// >  If this parameter is specified, the system checks the title specified by this parameter instead of the title of the input file specified by Media.
	//
	// example:
	//
	// Hello World
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The user-defined data, which can be up to 128 bytes in length.
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
