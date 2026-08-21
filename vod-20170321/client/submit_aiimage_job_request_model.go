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
	// The AI task pipeline ID.
	//
	// > A default ID is available, so this parameter is optional. If you need to perform batch imports, use a separate task pipeline. Submit a ticket to request configuration or contact Alibaba Cloud after-sales support for configuration. For more information about how to submit a ticket, see [Contact us](https://help.aliyun.com/document_detail/464625.html).
	//
	// example:
	//
	// 6492025b8f*****6ba5bb755a33438
	AIPipelineId *string `json:"AIPipelineId,omitempty" xml:"AIPipelineId,omitempty"`
	// The AI image template ID. You can obtain the template ID by using one of the following methods:
	//
	// - When you create an image template by calling the [AddAITemplate](https://help.aliyun.com/document_detail/102930.html) operation, the template ID is the value of the TemplateId parameter in the response.
	//
	// - After the template is created, you can call the [ListAITemplate](https://help.aliyun.com/document_detail/102936.html) operation to query the AI image template ID, which is the value of the TemplateId parameter in the response.
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
	// The custom settings.
	//
	// - The value must be a JSON string.
	//
	// - The value must contain the MessageCallback or Extend parameter.
	//
	// - The maximum length is 512 bytes.
	//
	// For more information about the parameter structure, see [UserData](https://help.aliyun.com/document_detail/86952.html).
	//
	// example:
	//
	// {"Extend":{"localId":"****","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The video ID. You can obtain the video ID by using one of the following methods:
	//
	// - For videos uploaded through the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the video ID.
	//
	// - When you upload a video by calling the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation, the video ID is the value of the VideoId parameter in the response.
	//
	// - After the video is uploaded, you can call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation to query the video ID, which is the value of the VideoId parameter in the response.
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
