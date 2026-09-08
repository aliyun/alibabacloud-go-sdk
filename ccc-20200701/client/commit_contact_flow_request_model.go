// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitContactFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactFlowId(v string) *CommitContactFlowRequest
	GetContactFlowId() *string
	SetDefinition(v string) *CommitContactFlowRequest
	GetDefinition() *string
	SetDescription(v string) *CommitContactFlowRequest
	GetDescription() *string
	SetDraftId(v string) *CommitContactFlowRequest
	GetDraftId() *string
	SetInstanceId(v string) *CommitContactFlowRequest
	GetInstanceId() *string
}

type CommitContactFlowRequest struct {
	// The contact flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7d706489-d06d-4a92-8666-8c9dba2c5cb1
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	// The definition of the Interactive Voice Response (IVR).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"activities":[{"type":"INCOMING_CALL","id":"e98f0d47","name":"开始","properties":{"position":{"x":263,"y":164}},"events":[{"event":"complete","next":"2d3ad2c2","edgeId":"41f7dbd0"}],"nodeIndex":0},{"type":"HANGUP","id":"bd4f37e2","name":"挂机","properties":{"position":{"x":765,"y":185}},"events":[{"event":"complete","next":null}],"nodeIndex":999},{"type":"PLAY_SAY","id":"2d3ad2c2","name":"放音","properties":{"say":"您好，欢迎来到云联络中心。","audioResourceId":"","position":{"x":485.5,"y":153.5},"audioType":"tts","audioInterrupt":false},"events":[{"event":"complete","next":"bd4f37e2","edgeId":"e1af4f1f"}],"nodeIndex":1}],"description":""}
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The version description.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The draft ID. This is the ID of the draft version that corresponds to the current contact flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// b28f74ca-5846-4496-8bbd-34fb1750798c
	DraftId *string `json:"DraftId,omitempty" xml:"DraftId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CommitContactFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s CommitContactFlowRequest) GoString() string {
	return s.String()
}

func (s *CommitContactFlowRequest) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *CommitContactFlowRequest) GetDefinition() *string {
	return s.Definition
}

func (s *CommitContactFlowRequest) GetDescription() *string {
	return s.Description
}

func (s *CommitContactFlowRequest) GetDraftId() *string {
	return s.DraftId
}

func (s *CommitContactFlowRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CommitContactFlowRequest) SetContactFlowId(v string) *CommitContactFlowRequest {
	s.ContactFlowId = &v
	return s
}

func (s *CommitContactFlowRequest) SetDefinition(v string) *CommitContactFlowRequest {
	s.Definition = &v
	return s
}

func (s *CommitContactFlowRequest) SetDescription(v string) *CommitContactFlowRequest {
	s.Description = &v
	return s
}

func (s *CommitContactFlowRequest) SetDraftId(v string) *CommitContactFlowRequest {
	s.DraftId = &v
	return s
}

func (s *CommitContactFlowRequest) SetInstanceId(v string) *CommitContactFlowRequest {
	s.InstanceId = &v
	return s
}

func (s *CommitContactFlowRequest) Validate() error {
	return dara.Validate(s)
}
