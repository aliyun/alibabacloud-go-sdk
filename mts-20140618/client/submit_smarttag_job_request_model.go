// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmarttagJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *SubmitSmarttagJobRequest
	GetContent() *string
	SetContentAddr(v string) *SubmitSmarttagJobRequest
	GetContentAddr() *string
	SetContentType(v string) *SubmitSmarttagJobRequest
	GetContentType() *string
	SetInput(v string) *SubmitSmarttagJobRequest
	GetInput() *string
	SetNotifyUrl(v string) *SubmitSmarttagJobRequest
	GetNotifyUrl() *string
	SetOwnerAccount(v string) *SubmitSmarttagJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SubmitSmarttagJobRequest
	GetOwnerId() *int64
	SetParams(v string) *SubmitSmarttagJobRequest
	GetParams() *string
	SetPipelineId(v string) *SubmitSmarttagJobRequest
	GetPipelineId() *string
	SetPriority(v string) *SubmitSmarttagJobRequest
	GetPriority() *string
	SetResourceOwnerAccount(v string) *SubmitSmarttagJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitSmarttagJobRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v string) *SubmitSmarttagJobRequest
	GetTemplateId() *string
	SetTitle(v string) *SubmitSmarttagJobRequest
	GetTitle() *string
	SetUserData(v string) *SubmitSmarttagJobRequest
	GetUserData() *string
}

type SubmitSmarttagJobRequest struct {
	// example:
	//
	// example content ****
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// http://exampleBucket.oss-cn-shanghai.aliyuncs.com/mps-test/ai-tag.mp4
	ContentAddr *string `json:"ContentAddr,omitempty" xml:"ContentAddr,omitempty"`
	// example:
	//
	// application/zip
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	Input       *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// https://example.com/endpoint/aliyun/ai?id=76401125000***
	NotifyUrl    *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// false
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the job in the ApsaraVideo Media Processing (MPS) queue to which the job is added. Valid values: 0 to 9. Default value: 5.
	//
	// example:
	//
	// 5
	Priority             *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The template ID, which is used to specify the analysis algorithm of the smart tagging job. For more information about how to manage templates, see [AddSmarttagTemplate](https://help.aliyun.com/document_detail/602910.html), [QuerySmarttagTemplateList](https://help.aliyun.com/document_detail/187770.html), [UpdateSmarttagTemplate](https://help.aliyun.com/document_detail/187776.html), and [DeleteSmarttagTemplate](https://help.aliyun.com/document_detail/187775.html).
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example-title-****
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// {"key":"value"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitSmarttagJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmarttagJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmarttagJobRequest) GetContent() *string {
	return s.Content
}

func (s *SubmitSmarttagJobRequest) GetContentAddr() *string {
	return s.ContentAddr
}

func (s *SubmitSmarttagJobRequest) GetContentType() *string {
	return s.ContentType
}

func (s *SubmitSmarttagJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitSmarttagJobRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *SubmitSmarttagJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitSmarttagJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitSmarttagJobRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitSmarttagJobRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitSmarttagJobRequest) GetPriority() *string {
	return s.Priority
}

func (s *SubmitSmarttagJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitSmarttagJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitSmarttagJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitSmarttagJobRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitSmarttagJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSmarttagJobRequest) SetContent(v string) *SubmitSmarttagJobRequest {
	s.Content = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetContentAddr(v string) *SubmitSmarttagJobRequest {
	s.ContentAddr = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetContentType(v string) *SubmitSmarttagJobRequest {
	s.ContentType = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetInput(v string) *SubmitSmarttagJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetNotifyUrl(v string) *SubmitSmarttagJobRequest {
	s.NotifyUrl = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetOwnerAccount(v string) *SubmitSmarttagJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetOwnerId(v int64) *SubmitSmarttagJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetParams(v string) *SubmitSmarttagJobRequest {
	s.Params = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetPipelineId(v string) *SubmitSmarttagJobRequest {
	s.PipelineId = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetPriority(v string) *SubmitSmarttagJobRequest {
	s.Priority = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetResourceOwnerAccount(v string) *SubmitSmarttagJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetResourceOwnerId(v int64) *SubmitSmarttagJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetTemplateId(v string) *SubmitSmarttagJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetTitle(v string) *SubmitSmarttagJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetUserData(v string) *SubmitSmarttagJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSmarttagJobRequest) Validate() error {
	return dara.Validate(s)
}
