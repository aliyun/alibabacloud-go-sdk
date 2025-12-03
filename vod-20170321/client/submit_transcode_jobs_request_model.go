// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTranscodeJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptConfig(v string) *SubmitTranscodeJobsRequest
	GetEncryptConfig() *string
	SetOverrideParams(v string) *SubmitTranscodeJobsRequest
	GetOverrideParams() *string
	SetPipelineId(v string) *SubmitTranscodeJobsRequest
	GetPipelineId() *string
	SetPriority(v string) *SubmitTranscodeJobsRequest
	GetPriority() *string
	SetReferenceId(v string) *SubmitTranscodeJobsRequest
	GetReferenceId() *string
	SetSessionId(v string) *SubmitTranscodeJobsRequest
	GetSessionId() *string
	SetTemplateGroupId(v string) *SubmitTranscodeJobsRequest
	GetTemplateGroupId() *string
	SetUserData(v string) *SubmitTranscodeJobsRequest
	GetUserData() *string
	SetVideoId(v string) *SubmitTranscodeJobsRequest
	GetVideoId() *string
}

type SubmitTranscodeJobsRequest struct {
	// The encryption configurations. The value must be a JSON string. This parameter is required only when you use HLS encryption.
	//
	// >
	//
	// 	- You must set **CipherText*	- in [EncrptConfig](https://help.aliyun.com/document_detail/86952.html) to the AES_128 cipher text that is obtained from the response to the [GenerateKMSDataKey](https://help.aliyun.com/document_detail/455051.html) operation. Otherwise, the HLS encryption fails. For more information about how to use HLS encryption, see [HLS encryption](https://help.aliyun.com/document_detail/68612.html).
	//
	// 	- You must select HLS encryption for the template specified by **TemplateGroupId*	- no matter you use HLS encryption or Alibaba Cloud proprietary cryptography. Otherwise, the transcoded file is not encrypted.
	//
	// example:
	//
	// {"CipherText":"ZjJmZGViNzUtZWY1Mi00Y2RlLTk3****", "DecryptKeyUri":"http://demo.aliyundoc.com?CipherText=ZjJmZGViNzUtZWY1Mi00Y2RlLTk3****","KeyServiceType":"KMS"}
	EncryptConfig *string `json:"EncryptConfig,omitempty" xml:"EncryptConfig,omitempty"`
	// The override parameter. The value must be a JSON string. You can use this parameter to override the image watermark, text watermark, or subtitle file specified in the transcoding template, or override the encoding format of the subtitle file. For more information about the data structure, see [OverrideParams](https://help.aliyun.com/document_detail/98618.html).
	//
	// example:
	//
	// {"Watermarks":[{"WatermarkId":"af2afe4761992c47dae973374****","FileUrl":"http://developer.aliyundoc.com/image/image.png"},{"WatermarkId":"e8e5b8038d7ada85b376c2707****","Content":"watermark test"}]}
	OverrideParams *string `json:"OverrideParams,omitempty" xml:"OverrideParams,omitempty"`
	// The ID of the queue that you want to use to run the job.
	//
	// example:
	//
	// d3e680e618708erf45fbf2cae7c****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the transcoding job in all queued jobs.
	//
	// 	- Valid values: **1*	- to **10**.
	//
	// 	- A value of **10*	- indicates the highest priority.
	//
	// 	- Default value: **6**.
	//
	// >  This parameter takes effect only on the queued transcoding jobs. The priorities of the in-progress transcoding jobs are not affected.
	//
	// example:
	//
	// 6
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The custom identifier for deduplication. If you send a request, an error is returned if a request with the same identifier was sent in the last 7 days. A custom identifier can be up to 50 characters in length and can contain letters, digits, hyphens (-), and underscores (_). If you do not specify this parameter or leave this parameter empty, duplicate requests are not filtered.
	//
	// example:
	//
	// 5c62d40299034bbaa4c195da330****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The ID of the transcoding template group that you want to use. To view the template group ID, perform the following operations: Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Configuration Management*	- > **Media Processing*	- > **Transcoding Template Groups**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0e408c803baf658ee637790c5d9f****
	TemplateGroupId *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	// The custom settings. The value must be a JSON string. You can configure settings such as message callbacks. For more information, see [UserData](https://help.aliyun.com/document_detail/86952.html).
	//
	// >  To use the callback configurations specified by this parameter, you must configure an HTTP callback URL and specify the types of the callback events in the ApsaraVideo VOD console. Otherwise, the callback configurations do not take effect.
	//
	// example:
	//
	// {"Extend":{"localId":"****","test":"***"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the video file. You can use one of the following methods to obtain the video ID:
	//
	// 	- Log on to the [ApsaraVideo VOD](https://vod.console.aliyun.com) console. In the left-side navigation pane, choose **Media Files*	- > **Audio/Video**. On the Video and Audio page, view the ID of the video file. This method is applicable to files that are uploaded by using the ApsaraVideo VOD console.
	//
	// 	- Obtain the value of VideoId from the response to the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation that you call to upload the video.
	//
	// 	- Obtain the value of VideoId from the response to the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation after you upload the video.
	//
	// example:
	//
	// 142710f878bd42508932f660d7b1****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s SubmitTranscodeJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobsRequest) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobsRequest) GetEncryptConfig() *string {
	return s.EncryptConfig
}

func (s *SubmitTranscodeJobsRequest) GetOverrideParams() *string {
	return s.OverrideParams
}

func (s *SubmitTranscodeJobsRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitTranscodeJobsRequest) GetPriority() *string {
	return s.Priority
}

func (s *SubmitTranscodeJobsRequest) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *SubmitTranscodeJobsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *SubmitTranscodeJobsRequest) GetTemplateGroupId() *string {
	return s.TemplateGroupId
}

func (s *SubmitTranscodeJobsRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitTranscodeJobsRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *SubmitTranscodeJobsRequest) SetEncryptConfig(v string) *SubmitTranscodeJobsRequest {
	s.EncryptConfig = &v
	return s
}

func (s *SubmitTranscodeJobsRequest) SetOverrideParams(v string) *SubmitTranscodeJobsRequest {
	s.OverrideParams = &v
	return s
}

func (s *SubmitTranscodeJobsRequest) SetPipelineId(v string) *SubmitTranscodeJobsRequest {
	s.PipelineId = &v
	return s
}

func (s *SubmitTranscodeJobsRequest) SetPriority(v string) *SubmitTranscodeJobsRequest {
	s.Priority = &v
	return s
}

func (s *SubmitTranscodeJobsRequest) SetReferenceId(v string) *SubmitTranscodeJobsRequest {
	s.ReferenceId = &v
	return s
}

func (s *SubmitTranscodeJobsRequest) SetSessionId(v string) *SubmitTranscodeJobsRequest {
	s.SessionId = &v
	return s
}

func (s *SubmitTranscodeJobsRequest) SetTemplateGroupId(v string) *SubmitTranscodeJobsRequest {
	s.TemplateGroupId = &v
	return s
}

func (s *SubmitTranscodeJobsRequest) SetUserData(v string) *SubmitTranscodeJobsRequest {
	s.UserData = &v
	return s
}

func (s *SubmitTranscodeJobsRequest) SetVideoId(v string) *SubmitTranscodeJobsRequest {
	s.VideoId = &v
	return s
}

func (s *SubmitTranscodeJobsRequest) Validate() error {
	return dara.Validate(s)
}
