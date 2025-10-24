// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaCensorJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBarrages(v string) *SubmitMediaCensorJobRequest
	GetBarrages() *string
	SetCoverImages(v string) *SubmitMediaCensorJobRequest
	GetCoverImages() *string
	SetDescription(v string) *SubmitMediaCensorJobRequest
	GetDescription() *string
	SetExternalUrl(v string) *SubmitMediaCensorJobRequest
	GetExternalUrl() *string
	SetInput(v string) *SubmitMediaCensorJobRequest
	GetInput() *string
	SetOwnerAccount(v string) *SubmitMediaCensorJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SubmitMediaCensorJobRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *SubmitMediaCensorJobRequest
	GetPipelineId() *string
	SetResourceOwnerAccount(v string) *SubmitMediaCensorJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitMediaCensorJobRequest
	GetResourceOwnerId() *int64
	SetTitle(v string) *SubmitMediaCensorJobRequest
	GetTitle() *string
	SetUserData(v string) *SubmitMediaCensorJobRequest
	GetUserData() *string
	SetVideoCensorConfig(v string) *SubmitMediaCensorJobRequest
	GetVideoCensorConfig() *string
}

type SubmitMediaCensorJobRequest struct {
	// The live comments.
	//
	// example:
	//
	// hello world
	Barrages *string `json:"Barrages,omitempty" xml:"Barrages,omitempty"`
	// The OSS URL of the image file that is used as the thumbnail. To view the OSS URL of the image file, you can log on to the **MPS console*	- and choose **Media Management*	- > **Media List*	- in the left-side navigation pane. You can specify up to five thumbnails in a JSON array.
	//
	// 	- Bucket: the name of the OSS bucket that stores the input file.
	//
	// 	- Location: the OSS region. The OSS region must be the same as the region in which your MPS service is activated.
	//
	// 	- Object: the OSS object to be moderated.
	//
	//     **
	//
	//     **Note**The name of the object cannot start with a forward slash (/). Otherwise, the operation fails to be called.
	//
	// example:
	//
	// [{"Bucket":"example-bucket-****","Location":"oss-cn-shanghai","Object":"example-****.jpeg"}]
	CoverImages *string `json:"CoverImages,omitempty" xml:"CoverImages,omitempty"`
	// The description of the video. The value can be up to 128 bytes in size.
	//
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the video.
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	// The Object Storage Service (OSS) URL of the media file to be moderated. To view the OSS URL of the media file, you can log on to the **MPS console*	- and choose **Media Management*	- > **Media List*	- in the left-side navigation pane. To moderate an image file, use the `CoverImage` parameter to specify the OSS URL of the image file. The value is a JSON object. For more information, see the "Input" section of the [Parameter details](https://help.aliyun.com/document_detail/29253.html) topic.
	//
	// 	- Bucket: the name of the OSS bucket that stores the input file.
	//
	// 	- Location: the OSS region. The OSS region must be the same as the region in which your MPS service is activated.
	//
	// 	- Object: the OSS object to be moderated.
	//
	//     **
	//
	//     **Note**The name of the object cannot start with a forward slash (/). Otherwise, the operation fails to be called.
	//
	// example:
	//
	// {"Bucket":"example-bucket-****","Location":"oss-cn-shanghai","Object":"example-****.flv"}
	Input        *string `json:"Input,omitempty" xml:"Input,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the MPS queue. This ID can be used to associate the job with a notification method. To view the ID of the MPS queue, you can log on to the **MPS console*	- and choose **Global Settings*	- > **Pipelines*	- in the left-side navigation pane. An empty string ("") indicates that the default MPS queue is used to run the job. By default, an MPS queue can process a maximum of 10 concurrent content moderation jobs. To increase the limit, [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.12246746.top-nav.ditem-sub.35da7bbcitpQnr#/ticket/createIndex).
	//
	// > MPS queues are automatically created by the system. For more information about how to query and update MPS queues, see the [UpdatePipeline](https://help.aliyun.com/document_detail/188374.html) topic.
	//
	// example:
	//
	// b22c173cced04565b1f38f1ecc39****
	PipelineId           *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The title of the video. The value can be up to 64 bytes in size.
	//
	// example:
	//
	// Hello World
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The custom data. The value can be up to 128 bytes in size.
	//
	// example:
	//
	// UserDatatestid-001-****
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The video moderation configurations and the OSS URLs of the output snapshots. To view the OSS URL of the media file, you can log on to the **MPS console*	- and choose **Media Management*	- > **Media List*	- in the left-side navigation pane.
	//
	// The value is a JSON object.
	//
	// 	- OutputFile:
	//
	//     	- Bucket: the name of the OSS bucket that stores the output file.
	//
	//     	- Location: the OSS region. The OSS region must be the same as the region in which your MPS service is activated.
	//
	//     	- Object: the OSS object to be generated. In the value, {Count} indicates the sequence number of the frame snapshot.
	//
	// 	- StoreVideoTimeline: specifies whether to generate the `{jobId}.video_timeline` file. The file is stored in OSS. A value of true indicates that the file is generated. A value of false indicates that the file is not generated. If you do not specify this parameter, the file is not generated by default. For more information about the format of the file, see the "VideoTimelines" parameter in the [QueryMediaCensorJobDetail](https://help.aliyun.com/document_detail/91779.html) topic.
	//
	// 	- SaveType: the output mode. A value of abnormal indicates that snapshots are generated only for illegal frames. A value of all indicates that snapshots are generated for all frames.
	//
	// 	- Biztype: the moderation template. If you do not specify this parameter or set the value to common, the default template is used. You can submit a ticket to create a custom moderation template. Then, set this parameter to your user ID to use the custom moderation template.
	//
	// 	- Scenes: the moderation scenarios. You can specify the moderation scenarios that you want to use. If you do not specify this parameter, the terrorism and porn moderation scenarios are used by default. Valid values:
	//
	//     	- porn: pornographic content detection
	//
	//     	- terrorism: terrorist content detection
	//
	//     	- ad: ad violation detection
	//
	//     	- live: undesirable scene detection
	//
	//     	- logo: special logo detection
	//
	//     	- audio: audio anti-spam
	//
	// > If the input file contains audio tracks and the audio moderation scenario is specified, the audio tracks are moderated. If the input file does not contain audio tracks, you do not need to specify the audio moderation scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Scenes" : ["porn"], "OutputFile":{"Bucket": "example-001","Location": "oss-cn-hangzhou","Object": "test/example-{Count}.jpg"},"SaveType" : "abnormal","BizType":"common"}
	VideoCensorConfig *string `json:"VideoCensorConfig,omitempty" xml:"VideoCensorConfig,omitempty"`
}

func (s SubmitMediaCensorJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaCensorJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitMediaCensorJobRequest) GetBarrages() *string {
	return s.Barrages
}

func (s *SubmitMediaCensorJobRequest) GetCoverImages() *string {
	return s.CoverImages
}

func (s *SubmitMediaCensorJobRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitMediaCensorJobRequest) GetExternalUrl() *string {
	return s.ExternalUrl
}

func (s *SubmitMediaCensorJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitMediaCensorJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitMediaCensorJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitMediaCensorJobRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitMediaCensorJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitMediaCensorJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitMediaCensorJobRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitMediaCensorJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitMediaCensorJobRequest) GetVideoCensorConfig() *string {
	return s.VideoCensorConfig
}

func (s *SubmitMediaCensorJobRequest) SetBarrages(v string) *SubmitMediaCensorJobRequest {
	s.Barrages = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetCoverImages(v string) *SubmitMediaCensorJobRequest {
	s.CoverImages = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetDescription(v string) *SubmitMediaCensorJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetExternalUrl(v string) *SubmitMediaCensorJobRequest {
	s.ExternalUrl = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetInput(v string) *SubmitMediaCensorJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetOwnerAccount(v string) *SubmitMediaCensorJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetOwnerId(v int64) *SubmitMediaCensorJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetPipelineId(v string) *SubmitMediaCensorJobRequest {
	s.PipelineId = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetResourceOwnerAccount(v string) *SubmitMediaCensorJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetResourceOwnerId(v int64) *SubmitMediaCensorJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetTitle(v string) *SubmitMediaCensorJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetUserData(v string) *SubmitMediaCensorJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetVideoCensorConfig(v string) *SubmitMediaCensorJobRequest {
	s.VideoCensorConfig = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) Validate() error {
	return dara.Validate(s)
}
