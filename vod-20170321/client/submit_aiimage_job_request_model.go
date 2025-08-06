// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIImageJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIPipelineId(v string) *SubmitAIImageJobRequest
	GetAIPipelineId() *string
	SetAITemplateId(v string) *SubmitAIImageJobRequest
	GetAITemplateId() *string
	SetOwnerAccount(v string) *SubmitAIImageJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SubmitAIImageJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SubmitAIImageJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SubmitAIImageJobRequest
	GetResourceOwnerId() *string
	SetUserData(v string) *SubmitAIImageJobRequest
	GetUserData() *string
	SetVideoId(v string) *SubmitAIImageJobRequest
	GetVideoId() *string
}

type SubmitAIImageJobRequest struct {
	// The ID of the pipeline that is used for the AI processing job.
	//
	// >  This parameter is optional if you specify a default pipeline ID. If you want to use a separate pipeline to submit multiple AI processing jobs., submit a ticket or contact Alibaba Cloud after-sales engineers. For more information about how to submit a ticket, see [Contact us](https://help.aliyun.com/document_detail/464625.html).
	//
	// example:
	//
	// 6492025b8f*****6ba5bb755a33438
	AIPipelineId *string `json:"AIPipelineId,omitempty" xml:"AIPipelineId,omitempty"`
	// The ID of the AI template. You can use one of the following methods to obtain the ID:
	//
	// 	- Obtain the value of TemplateId from the response to the [AddAITemplate](https://help.aliyun.com/document_detail/102930.html) that you call to create the template.
	//
	// 	- Obtain the value of TemplateId from the response to the [ListAITemplate](https://help.aliyun.com/document_detail/102936.html) operation after you create the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ef1a8842cb9f*****cea80cad902e416
	AITemplateId         *string `json:"AITemplateId,omitempty" xml:"AITemplateId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The user data.
	//
	// 	- The value must be a JSON string.
	//
	// 	- You must specify the MessageCallback or Extend parameter.
	//
	// 	- The value can contain a maximum of 512 bytes.
	//
	// For more information, see the "UserData: specifies the custom configurations for media upload" section of the [Request parameters](https://help.aliyun.com/document_detail/86952.html) topic.
	//
	// example:
	//
	// {"Extend":{"localId":"****","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the video. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD](https://vod.console.aliyun.com) console. In the left-side navigation pane, choose **Media Files*	- > **Audio/Video**. On the Video and Audio page, view the ID of the video file. This method is applicable to files that are uploaded by using the ApsaraVideo VOD console.
	//
	// 	- Obtain the value of VideoId from the response to the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation that you call to upload the video.
	//
	// 	- Obtain the value of VideoId from the response to the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation after you upload the video.
	//
	// This parameter is required.
	//
	// example:
	//
	// 357a8748c5774*****89d2726e6436aa
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s SubmitAIImageJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIImageJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIImageJobRequest) GetAIPipelineId() *string {
	return s.AIPipelineId
}

func (s *SubmitAIImageJobRequest) GetAITemplateId() *string {
	return s.AITemplateId
}

func (s *SubmitAIImageJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitAIImageJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SubmitAIImageJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitAIImageJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SubmitAIImageJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAIImageJobRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *SubmitAIImageJobRequest) SetAIPipelineId(v string) *SubmitAIImageJobRequest {
	s.AIPipelineId = &v
	return s
}

func (s *SubmitAIImageJobRequest) SetAITemplateId(v string) *SubmitAIImageJobRequest {
	s.AITemplateId = &v
	return s
}

func (s *SubmitAIImageJobRequest) SetOwnerAccount(v string) *SubmitAIImageJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitAIImageJobRequest) SetOwnerId(v string) *SubmitAIImageJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitAIImageJobRequest) SetResourceOwnerAccount(v string) *SubmitAIImageJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitAIImageJobRequest) SetResourceOwnerId(v string) *SubmitAIImageJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitAIImageJobRequest) SetUserData(v string) *SubmitAIImageJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAIImageJobRequest) SetVideoId(v string) *SubmitAIImageJobRequest {
	s.VideoId = &v
	return s
}

func (s *SubmitAIImageJobRequest) Validate() error {
	return dara.Validate(s)
}
