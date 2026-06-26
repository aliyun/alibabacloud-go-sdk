// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUnrecognizingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFinalAction(v string) *ModifyUnrecognizingConfigRequest
	GetFinalAction() *string
	SetFinalActionParams(v string) *ModifyUnrecognizingConfigRequest
	GetFinalActionParams() *string
	SetFinalPrompt(v string) *ModifyUnrecognizingConfigRequest
	GetFinalPrompt() *string
	SetInstanceId(v string) *ModifyUnrecognizingConfigRequest
	GetInstanceId() *string
	SetPrompt(v string) *ModifyUnrecognizingConfigRequest
	GetPrompt() *string
	SetThreshold(v int32) *ModifyUnrecognizingConfigRequest
	GetThreshold() *int32
}

type ModifyUnrecognizingConfigRequest struct {
	// The rejection action performed after the final rejection prompt is played.
	//
	// This parameter is required.
	//
	// example:
	//
	// TransferToAgent
	FinalAction *string `json:"FinalAction,omitempty" xml:"FinalAction,omitempty"`
	// The action parameters for the rejection action, in JSON format.
	//
	// example:
	//
	// {}
	FinalActionParams *string `json:"FinalActionParams,omitempty" xml:"FinalActionParams,omitempty"`
	// The final rejection prompt. The service plays this prompt when the rejection threshold is met.
	//
	// This parameter is required.
	//
	// example:
	//
	// 再见
	FinalPrompt *string `json:"FinalPrompt,omitempty" xml:"FinalPrompt,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12f407b22cbe4890ac595f09985848d5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The rejection prompt played when the service fails to recognize user input.
	//
	// This parameter is required.
	//
	// example:
	//
	// 我不理解您的意思，请再说一遍
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The rejection threshold. The maximum number of consecutive rejections before the service triggers the rejection action.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ModifyUnrecognizingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUnrecognizingConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyUnrecognizingConfigRequest) GetFinalAction() *string {
	return s.FinalAction
}

func (s *ModifyUnrecognizingConfigRequest) GetFinalActionParams() *string {
	return s.FinalActionParams
}

func (s *ModifyUnrecognizingConfigRequest) GetFinalPrompt() *string {
	return s.FinalPrompt
}

func (s *ModifyUnrecognizingConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyUnrecognizingConfigRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *ModifyUnrecognizingConfigRequest) GetThreshold() *int32 {
	return s.Threshold
}

func (s *ModifyUnrecognizingConfigRequest) SetFinalAction(v string) *ModifyUnrecognizingConfigRequest {
	s.FinalAction = &v
	return s
}

func (s *ModifyUnrecognizingConfigRequest) SetFinalActionParams(v string) *ModifyUnrecognizingConfigRequest {
	s.FinalActionParams = &v
	return s
}

func (s *ModifyUnrecognizingConfigRequest) SetFinalPrompt(v string) *ModifyUnrecognizingConfigRequest {
	s.FinalPrompt = &v
	return s
}

func (s *ModifyUnrecognizingConfigRequest) SetInstanceId(v string) *ModifyUnrecognizingConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyUnrecognizingConfigRequest) SetPrompt(v string) *ModifyUnrecognizingConfigRequest {
	s.Prompt = &v
	return s
}

func (s *ModifyUnrecognizingConfigRequest) SetThreshold(v int32) *ModifyUnrecognizingConfigRequest {
	s.Threshold = &v
	return s
}

func (s *ModifyUnrecognizingConfigRequest) Validate() error {
	return dara.Validate(s)
}
