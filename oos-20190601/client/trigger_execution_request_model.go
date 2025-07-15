// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *TriggerExecutionRequest
	GetClientToken() *string
	SetContent(v string) *TriggerExecutionRequest
	GetContent() *string
	SetExecutionId(v string) *TriggerExecutionRequest
	GetExecutionId() *string
	SetRegionId(v string) *TriggerExecutionRequest
	GetRegionId() *string
	SetType(v string) *TriggerExecutionRequest
	GetType() *string
}

type TriggerExecutionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// dswe2-3i0-029
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The message body to be sent to the trigger task.
	//
	// example:
	//
	// {"eventTime": "20181226T220114.058+0800", "id": "9435EAD6-3CF6-4494-8F7A-3A********77","level": "INFO","name": "Instance:StateChange","product": "ECS","regionId":"cn-hangzhou","resourceId": "acs:ecs:cn-hangzhou:169070********30:instance/i-bp1ecr********5go2go","userId": "169070********30","ver": "1.0","content": {"resourceId": "i-bp1ecr********5go2go", "resourceType": "ALIYUN::ECS::Instance","state": "Stopping"} }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the event-, alert-, or timer-triggered execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-sadw3f23rsad
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the trigger. Valid values:
	//
	// 	- Event
	//
	// 	- Alarm
	//
	// 	- Timer
	//
	// This parameter is required.
	//
	// example:
	//
	// Event
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s TriggerExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerExecutionRequest) GoString() string {
	return s.String()
}

func (s *TriggerExecutionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *TriggerExecutionRequest) GetContent() *string {
	return s.Content
}

func (s *TriggerExecutionRequest) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *TriggerExecutionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TriggerExecutionRequest) GetType() *string {
	return s.Type
}

func (s *TriggerExecutionRequest) SetClientToken(v string) *TriggerExecutionRequest {
	s.ClientToken = &v
	return s
}

func (s *TriggerExecutionRequest) SetContent(v string) *TriggerExecutionRequest {
	s.Content = &v
	return s
}

func (s *TriggerExecutionRequest) SetExecutionId(v string) *TriggerExecutionRequest {
	s.ExecutionId = &v
	return s
}

func (s *TriggerExecutionRequest) SetRegionId(v string) *TriggerExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *TriggerExecutionRequest) SetType(v string) *TriggerExecutionRequest {
	s.Type = &v
	return s
}

func (s *TriggerExecutionRequest) Validate() error {
	return dara.Validate(s)
}
