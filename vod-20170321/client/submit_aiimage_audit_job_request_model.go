// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIImageAuditJobRequest interface {
	dara.Model
	String() string
	GoString() string
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
	SetTemplateId(v string) *SubmitAIImageAuditJobRequest
	GetTemplateId() *string
}

type SubmitAIImageAuditJobRequest struct {
	// The configuration information about the review job.
	//
	// 	- Other configuration items of the review job. Only the ResourceType field is supported. This field is used to specify the type of media files. You can adjust review standards and rules based on the type of media files.
	//
	// 	- The value of ResourceType can contain only letters, digits, and underscores (_).
	//
	// > 	- You can specify a value for the ResourceType field based on the preceding limits. After you specify a value for the ResourceType field, you must [submit a ticket](https://yida.alibaba-inc.com/o/ticketapply). The value takes effect after Alibaba Cloud processes your ticket.
	//
	// >	- If you want to change moderation policies and rules based on ResourceType, [submit a ticket](https://yida.alibaba-inc.com/o/ticketapply) to contact technical support.
	//
	// example:
	//
	// {"ResourceType":"****_short_video"}
	MediaAuditConfiguration *string `json:"MediaAuditConfiguration,omitempty" xml:"MediaAuditConfiguration,omitempty"`
	// The ID of the image.
	//
	// The unique ID of the image is returned after the image is uploaded to ApsaraVideo VOD.
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
	// The ID of the AI template. You can use one of the following methods to obtain the ID:
	//
	// 	- Obtain the value of TemplateId from the response to the [AddAITemplate](https://help.aliyun.com/document_detail/102930.html) operation that you call to create an AI template.
	//
	// 	- Obtain the value of TemplateId from the response to the [ListAITemplate](https://help.aliyun.com/document_detail/102936.html) operation that you call to create an AI template.
	//
	// This parameter is required.
	//
	// example:
	//
	// VOD-0003-00****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitAIImageAuditJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIImageAuditJobRequest) GoString() string {
	return s.String()
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

func (s *SubmitAIImageAuditJobRequest) GetTemplateId() *string {
	return s.TemplateId
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

func (s *SubmitAIImageAuditJobRequest) SetTemplateId(v string) *SubmitAIImageAuditJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitAIImageAuditJobRequest) Validate() error {
	return dara.Validate(s)
}
