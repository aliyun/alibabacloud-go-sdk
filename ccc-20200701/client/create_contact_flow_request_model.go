// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContactFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefinition(v string) *CreateContactFlowRequest
	GetDefinition() *string
	SetDescription(v string) *CreateContactFlowRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateContactFlowRequest
	GetInstanceId() *string
	SetName(v string) *CreateContactFlowRequest
	GetName() *string
	SetType(v string) *CreateContactFlowRequest
	GetType() *string
}

type CreateContactFlowRequest struct {
	// IVR content.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"activities":[{"type":"INCOMING_CALL","id":"e98f0d47","name":"开始","properties":{"position":{"x":263,"y":164}},"events":[{"event":"complete","next":"2d3ad2c2","edgeId":"41f7dbd0"}],"nodeIndex":0},{"type":"HANGUP","id":"bd4f37e2","name":"挂机","properties":{"position":{"x":765,"y":185}},"events":[{"event":"complete","next":null}],"nodeIndex":999},{"type":"PLAY_SAY","id":"2d3ad2c2","name":"放音","properties":{"say":"您好，欢迎来到云联络中心。","audioResourceId":"","position":{"x":485.5,"y":153.5},"audioType":"tts","audioInterrupt":false},"events":[{"event":"complete","next":"bd4f37e2","edgeId":"e1af4f1f"}],"nodeIndex":1}],"description":""}
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// Version description.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Contact flow name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 默认流程
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Flow type.<br>
	//
	// Main flow: MAIN_FLOW<br>
	//
	// Child flow: SUB_FLOW<br>
	//
	// Survey flow: SURVEY_FLOW<br><br><br>
	//
	// This parameter is required.
	//
	// example:
	//
	// MAIN_FLOW
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateContactFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContactFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateContactFlowRequest) GetDefinition() *string {
	return s.Definition
}

func (s *CreateContactFlowRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateContactFlowRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateContactFlowRequest) GetName() *string {
	return s.Name
}

func (s *CreateContactFlowRequest) GetType() *string {
	return s.Type
}

func (s *CreateContactFlowRequest) SetDefinition(v string) *CreateContactFlowRequest {
	s.Definition = &v
	return s
}

func (s *CreateContactFlowRequest) SetDescription(v string) *CreateContactFlowRequest {
	s.Description = &v
	return s
}

func (s *CreateContactFlowRequest) SetInstanceId(v string) *CreateContactFlowRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateContactFlowRequest) SetName(v string) *CreateContactFlowRequest {
	s.Name = &v
	return s
}

func (s *CreateContactFlowRequest) SetType(v string) *CreateContactFlowRequest {
	s.Type = &v
	return s
}

func (s *CreateContactFlowRequest) Validate() error {
	return dara.Validate(s)
}
