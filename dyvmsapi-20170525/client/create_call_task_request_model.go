// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *CreateCallTaskRequest
	GetBizType() *string
	SetData(v string) *CreateCallTaskRequest
	GetData() *string
	SetDataType(v string) *CreateCallTaskRequest
	GetDataType() *string
	SetFireTime(v string) *CreateCallTaskRequest
	GetFireTime() *string
	SetOwnerId(v int64) *CreateCallTaskRequest
	GetOwnerId() *int64
	SetResource(v string) *CreateCallTaskRequest
	GetResource() *string
	SetResourceOwnerAccount(v string) *CreateCallTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCallTaskRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *CreateCallTaskRequest
	GetResourceType() *string
	SetScheduleType(v string) *CreateCallTaskRequest
	GetScheduleType() *string
	SetStopTime(v string) *CreateCallTaskRequest
	GetStopTime() *string
	SetTaskName(v string) *CreateCallTaskRequest
	GetTaskName() *string
	SetTemplateCode(v string) *CreateCallTaskRequest
	GetTemplateCode() *string
	SetTemplateName(v string) *CreateCallTaskRequest
	GetTemplateName() *string
}

type CreateCallTaskRequest struct {
	// The type of the task template. Valid values:
	//
	// 	- **VMS_VOICE_TTS**: the text-to-speech (TTS) notification template.
	//
	// 	- **VMS_VOICE_CODE**: the voice notification template.
	//
	// 	- **VMS_TTS**: the voice verification code template.
	//
	// example:
	//
	// VMS_VOICE_TTS
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The called numbers.
	//
	// 	- If you set DataType to LIST, the value of Data is in the LIST format.
	//
	// 	- If you set DataType to JSON, the value of Data is in the JSON format.
	//
	// example:
	//
	// { "paramNames":["name1","name2","key3"], "calleeList":[ { "callee":"131222222", "params":["zangsan","zhangsan01","zhangsan02"] }, { "callee":"131222222", "params":["zangsan","zhangsan01","zhangsan02"] }, ] }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The type of called numbers. Valid values:
	//
	// 	- **LIST**: the called numbers that are separated by commas (,).
	//
	// 	- **JSON**: a JSON-formatted list of called numbers with template parameters.
	//
	// example:
	//
	// JSON
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// This parameter is unavailable.
	//
	// example:
	//
	// 1
	FireTime *string `json:"FireTime,omitempty" xml:"FireTime,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The calling number. Only virtual numbers are supported.
	//
	// example:
	//
	// 05516214****
	Resource             *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the calling number. Set the value to **LIST**.
	//
	// example:
	//
	// LIST
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is unavailable.
	//
	// example:
	//
	// 1
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// This parameter is unavailable.
	//
	// example:
	//
	// 1
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// The task name.
	//
	// example:
	//
	// Aliyun
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The template ID.
	//
	// example:
	//
	// TTS_2122****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The template name.
	//
	// example:
	//
	// Test Template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateCallTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCallTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateCallTaskRequest) GetBizType() *string {
	return s.BizType
}

func (s *CreateCallTaskRequest) GetData() *string {
	return s.Data
}

func (s *CreateCallTaskRequest) GetDataType() *string {
	return s.DataType
}

func (s *CreateCallTaskRequest) GetFireTime() *string {
	return s.FireTime
}

func (s *CreateCallTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCallTaskRequest) GetResource() *string {
	return s.Resource
}

func (s *CreateCallTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCallTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCallTaskRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateCallTaskRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *CreateCallTaskRequest) GetStopTime() *string {
	return s.StopTime
}

func (s *CreateCallTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateCallTaskRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateCallTaskRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateCallTaskRequest) SetBizType(v string) *CreateCallTaskRequest {
	s.BizType = &v
	return s
}

func (s *CreateCallTaskRequest) SetData(v string) *CreateCallTaskRequest {
	s.Data = &v
	return s
}

func (s *CreateCallTaskRequest) SetDataType(v string) *CreateCallTaskRequest {
	s.DataType = &v
	return s
}

func (s *CreateCallTaskRequest) SetFireTime(v string) *CreateCallTaskRequest {
	s.FireTime = &v
	return s
}

func (s *CreateCallTaskRequest) SetOwnerId(v int64) *CreateCallTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCallTaskRequest) SetResource(v string) *CreateCallTaskRequest {
	s.Resource = &v
	return s
}

func (s *CreateCallTaskRequest) SetResourceOwnerAccount(v string) *CreateCallTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCallTaskRequest) SetResourceOwnerId(v int64) *CreateCallTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCallTaskRequest) SetResourceType(v string) *CreateCallTaskRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateCallTaskRequest) SetScheduleType(v string) *CreateCallTaskRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateCallTaskRequest) SetStopTime(v string) *CreateCallTaskRequest {
	s.StopTime = &v
	return s
}

func (s *CreateCallTaskRequest) SetTaskName(v string) *CreateCallTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateCallTaskRequest) SetTemplateCode(v string) *CreateCallTaskRequest {
	s.TemplateCode = &v
	return s
}

func (s *CreateCallTaskRequest) SetTemplateName(v string) *CreateCallTaskRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateCallTaskRequest) Validate() error {
	return dara.Validate(s)
}
