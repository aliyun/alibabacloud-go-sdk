// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIImageAuditJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCensorProvider(v string) *SubmitAIImageAuditJobRequest
	GetCensorProvider() *string
	SetImageService(v string) *SubmitAIImageAuditJobRequest
	GetImageService() *string
	SetMediaAuditConfiguration(v string) *SubmitAIImageAuditJobRequest
	GetMediaAuditConfiguration() *string
	SetMediaId(v string) *SubmitAIImageAuditJobRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *SubmitAIImageAuditJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SubmitAIImageAuditJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SubmitAIImageAuditJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SubmitAIImageAuditJobRequest
	GetResourceOwnerId() *string
	SetServiceParameters(v string) *SubmitAIImageAuditJobRequest
	GetServiceParameters() *string
	SetTemplateId(v string) *SubmitAIImageAuditJobRequest
	GetTemplateId() *string
}

type SubmitAIImageAuditJobRequest struct {
	CensorProvider *string `json:"CensorProvider,omitempty" xml:"CensorProvider,omitempty"`
	ImageService   *string `json:"ImageService,omitempty" xml:"ImageService,omitempty"`
	// The review node configuration.
	//
	// - Other configuration items of the review node. Currently, only the ResourceType field is supported, which is used to specify the media file type. You can adjust the review standards and rules for the specified type.
	//
	// - Usage notes for ResourceType: Only letters, digits, and underscores (_) are allowed.
	//
	// >- You can customize the ResourceType field based on the usage notes. After customization, [submit a Yida form](https://yida.alibaba-inc.com/o/ticketapply) to commit to Alibaba Cloud for spooling before the configuration takes effect.
	//
	// >- To adjust the review standards and rules for a specific ResourceType, [submit a Yida form](https://yida.alibaba-inc.com/o/ticketapply) to request technical support.
	//
	// example:
	//
	// {"ResourceType":"****_short_video"}
	MediaAuditConfiguration *string `json:"MediaAuditConfiguration,omitempty" xml:"MediaAuditConfiguration,omitempty"`
	// The image ID.
	//
	// The unique identifier of the image returned after the image is uploaded to ApsaraVideo VOD.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1aa3024aee64*****6dc8ca20dbc320
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ServiceParameters    *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
	// The AI template ID. You can obtain the ID by using one of the following methods:
	//
	// - Call the [Add AI template](https://help.aliyun.com/document_detail/102930.html) operation to add an AI template. The AI template ID is the value of TemplateId in the response.
	//
	// - After the AI template is added, call the [Query AI template list](https://help.aliyun.com/document_detail/102936.html) operation to query the AI template ID, which is the value of TemplateId in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// a07a7f7d7d10eb9fd999e56ecc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitAIImageAuditJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIImageAuditJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIImageAuditJobRequest) GetCensorProvider() *string {
	return s.CensorProvider
}

func (s *SubmitAIImageAuditJobRequest) GetImageService() *string {
	return s.ImageService
}

func (s *SubmitAIImageAuditJobRequest) GetMediaAuditConfiguration() *string {
	return s.MediaAuditConfiguration
}

func (s *SubmitAIImageAuditJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAIImageAuditJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitAIImageAuditJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SubmitAIImageAuditJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitAIImageAuditJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SubmitAIImageAuditJobRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *SubmitAIImageAuditJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitAIImageAuditJobRequest) SetCensorProvider(v string) *SubmitAIImageAuditJobRequest {
	s.CensorProvider = &v
	return s
}

func (s *SubmitAIImageAuditJobRequest) SetImageService(v string) *SubmitAIImageAuditJobRequest {
	s.ImageService = &v
	return s
}

func (s *SubmitAIImageAuditJobRequest) SetMediaAuditConfiguration(v string) *SubmitAIImageAuditJobRequest {
	s.MediaAuditConfiguration = &v
	return s
}

func (s *SubmitAIImageAuditJobRequest) SetMediaId(v string) *SubmitAIImageAuditJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitAIImageAuditJobRequest) SetOwnerAccount(v string) *SubmitAIImageAuditJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitAIImageAuditJobRequest) SetOwnerId(v string) *SubmitAIImageAuditJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitAIImageAuditJobRequest) SetResourceOwnerAccount(v string) *SubmitAIImageAuditJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitAIImageAuditJobRequest) SetResourceOwnerId(v string) *SubmitAIImageAuditJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitAIImageAuditJobRequest) SetServiceParameters(v string) *SubmitAIImageAuditJobRequest {
	s.ServiceParameters = &v
	return s
}

func (s *SubmitAIImageAuditJobRequest) SetTemplateId(v string) *SubmitAIImageAuditJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitAIImageAuditJobRequest) Validate() error {
	return dara.Validate(s)
}
