// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAnalysisJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisConfig(v string) *SubmitAnalysisJobRequest
	GetAnalysisConfig() *string
	SetInput(v string) *SubmitAnalysisJobRequest
	GetInput() *string
	SetOwnerAccount(v string) *SubmitAnalysisJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SubmitAnalysisJobRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *SubmitAnalysisJobRequest
	GetPipelineId() *string
	SetPriority(v string) *SubmitAnalysisJobRequest
	GetPriority() *string
	SetResourceOwnerAccount(v string) *SubmitAnalysisJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitAnalysisJobRequest
	GetResourceOwnerId() *int64
	SetUserData(v string) *SubmitAnalysisJobRequest
	GetUserData() *string
}

type SubmitAnalysisJobRequest struct {
	// The job configurations. Set this parameter as required. For more information, see the "AnalysisConfig" section of the [Parameter details](https://help.aliyun.com/document_detail/29253.html) topic.
	//
	// example:
	//
	// {"QualityControl":{"RateQuality":25,"MethodStreaming":"network"}}
	AnalysisConfig *string `json:"AnalysisConfig,omitempty" xml:"AnalysisConfig,omitempty"`
	// The input information about the preset template analysis job to be submitted. The value must be a JSON object. You must log on to the Object Storage Service (OSS) console to grant the read permissions on the specified OSS bucket to MPS. For more information, see the "Input" section of the [Parameter details](https://help.aliyun.com/document_detail/29253.html) topic.
	//
	// > The OSS bucket must reside in the same region as your MPS service.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Bucket":"example-bucket","Location":"oss-cn-hangzhou","Object":"example.flv"}
	Input        *string `json:"Input,omitempty" xml:"Input,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the MPS queue to which the job is submitted. To view the ID of the MPS queue, log on to the **MPS console*	- and choose **Global Settings*	- > **Pipelines*	- in the left-side navigation pane. If you want to enable asynchronous notifications, make sure that the MPS queue is bound to a Message Service (MNS) topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// bb558c1cc25b45309aab5be44d19****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the job in the MPS queue to which the job is submitted.
	//
	// 	- Valid values: **1 to 10**. A value of 10 indicates the highest priority.
	//
	// 	- Default value: **6**.
	//
	// example:
	//
	// 10
	Priority             *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The custom data. The custom data can contain letters, digits, and hyphens (-), and can be up to 1,024 bytes in length. The custom data cannot start with a special character.
	//
	// example:
	//
	// testid-001
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAnalysisJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAnalysisJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAnalysisJobRequest) GetAnalysisConfig() *string {
	return s.AnalysisConfig
}

func (s *SubmitAnalysisJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitAnalysisJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitAnalysisJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitAnalysisJobRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitAnalysisJobRequest) GetPriority() *string {
	return s.Priority
}

func (s *SubmitAnalysisJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitAnalysisJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitAnalysisJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAnalysisJobRequest) SetAnalysisConfig(v string) *SubmitAnalysisJobRequest {
	s.AnalysisConfig = &v
	return s
}

func (s *SubmitAnalysisJobRequest) SetInput(v string) *SubmitAnalysisJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitAnalysisJobRequest) SetOwnerAccount(v string) *SubmitAnalysisJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitAnalysisJobRequest) SetOwnerId(v int64) *SubmitAnalysisJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitAnalysisJobRequest) SetPipelineId(v string) *SubmitAnalysisJobRequest {
	s.PipelineId = &v
	return s
}

func (s *SubmitAnalysisJobRequest) SetPriority(v string) *SubmitAnalysisJobRequest {
	s.Priority = &v
	return s
}

func (s *SubmitAnalysisJobRequest) SetResourceOwnerAccount(v string) *SubmitAnalysisJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitAnalysisJobRequest) SetResourceOwnerId(v int64) *SubmitAnalysisJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitAnalysisJobRequest) SetUserData(v string) *SubmitAnalysisJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAnalysisJobRequest) Validate() error {
	return dara.Validate(s)
}
