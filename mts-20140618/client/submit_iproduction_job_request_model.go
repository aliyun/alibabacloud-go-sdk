// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIProductionJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *SubmitIProductionJobRequest
	GetFunctionName() *string
	SetInput(v string) *SubmitIProductionJobRequest
	GetInput() *string
	SetJobParams(v string) *SubmitIProductionJobRequest
	GetJobParams() *string
	SetModelId(v string) *SubmitIProductionJobRequest
	GetModelId() *string
	SetNotifyUrl(v string) *SubmitIProductionJobRequest
	GetNotifyUrl() *string
	SetOutput(v string) *SubmitIProductionJobRequest
	GetOutput() *string
	SetOwnerAccount(v string) *SubmitIProductionJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SubmitIProductionJobRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *SubmitIProductionJobRequest
	GetPipelineId() *string
	SetResourceOwnerAccount(v string) *SubmitIProductionJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitIProductionJobRequest
	GetResourceOwnerId() *int64
	SetScheduleParams(v string) *SubmitIProductionJobRequest
	GetScheduleParams() *string
	SetUserData(v string) *SubmitIProductionJobRequest
	GetUserData() *string
}

type SubmitIProductionJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ImageCartoonize
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// example:
	//
	// oss://example-****.oss-cn-shanghai.aliyuncs.com/example.mp4
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// {"Model":"gif"}
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	// example:
	//
	// null
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// example:
	//
	// mns://125340688170****.mns.cn-beijing.aliyuncs.com/queues/example-pipeline
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// example:
	//
	// oss://example-****.oss-cn-shanghai.aliyuncs.com/iproduction/{source}-{timestamp}-{sequenceId}.srt
	Output       *string `json:"Output,omitempty" xml:"Output,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	PipelineId           *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// null
	ScheduleParams *string `json:"ScheduleParams,omitempty" xml:"ScheduleParams,omitempty"`
	// example:
	//
	// null
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitIProductionJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitIProductionJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitIProductionJobRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *SubmitIProductionJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitIProductionJobRequest) GetJobParams() *string {
	return s.JobParams
}

func (s *SubmitIProductionJobRequest) GetModelId() *string {
	return s.ModelId
}

func (s *SubmitIProductionJobRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *SubmitIProductionJobRequest) GetOutput() *string {
	return s.Output
}

func (s *SubmitIProductionJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitIProductionJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitIProductionJobRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitIProductionJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitIProductionJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitIProductionJobRequest) GetScheduleParams() *string {
	return s.ScheduleParams
}

func (s *SubmitIProductionJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitIProductionJobRequest) SetFunctionName(v string) *SubmitIProductionJobRequest {
	s.FunctionName = &v
	return s
}

func (s *SubmitIProductionJobRequest) SetInput(v string) *SubmitIProductionJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitIProductionJobRequest) SetJobParams(v string) *SubmitIProductionJobRequest {
	s.JobParams = &v
	return s
}

func (s *SubmitIProductionJobRequest) SetModelId(v string) *SubmitIProductionJobRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitIProductionJobRequest) SetNotifyUrl(v string) *SubmitIProductionJobRequest {
	s.NotifyUrl = &v
	return s
}

func (s *SubmitIProductionJobRequest) SetOutput(v string) *SubmitIProductionJobRequest {
	s.Output = &v
	return s
}

func (s *SubmitIProductionJobRequest) SetOwnerAccount(v string) *SubmitIProductionJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitIProductionJobRequest) SetOwnerId(v int64) *SubmitIProductionJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitIProductionJobRequest) SetPipelineId(v string) *SubmitIProductionJobRequest {
	s.PipelineId = &v
	return s
}

func (s *SubmitIProductionJobRequest) SetResourceOwnerAccount(v string) *SubmitIProductionJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitIProductionJobRequest) SetResourceOwnerId(v int64) *SubmitIProductionJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitIProductionJobRequest) SetScheduleParams(v string) *SubmitIProductionJobRequest {
	s.ScheduleParams = &v
	return s
}

func (s *SubmitIProductionJobRequest) SetUserData(v string) *SubmitIProductionJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitIProductionJobRequest) Validate() error {
	return dara.Validate(s)
}
