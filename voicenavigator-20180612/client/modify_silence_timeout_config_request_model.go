// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySilenceTimeoutConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFinalAction(v string) *ModifySilenceTimeoutConfigRequest
	GetFinalAction() *string
	SetFinalActionParams(v string) *ModifySilenceTimeoutConfigRequest
	GetFinalActionParams() *string
	SetFinalPrompt(v string) *ModifySilenceTimeoutConfigRequest
	GetFinalPrompt() *string
	SetInstanceId(v string) *ModifySilenceTimeoutConfigRequest
	GetInstanceId() *string
	SetIntentTrigger(v string) *ModifySilenceTimeoutConfigRequest
	GetIntentTrigger() *string
	SetPrompt(v string) *ModifySilenceTimeoutConfigRequest
	GetPrompt() *string
	SetSourceType(v string) *ModifySilenceTimeoutConfigRequest
	GetSourceType() *string
	SetThreshold(v int32) *ModifySilenceTimeoutConfigRequest
	GetThreshold() *int32
	SetTimeout(v int64) *ModifySilenceTimeoutConfigRequest
	GetTimeout() *int64
}

type ModifySilenceTimeoutConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// TransferToAgent
	FinalAction *string `json:"FinalAction,omitempty" xml:"FinalAction,omitempty"`
	// example:
	//
	// {}
	FinalActionParams *string `json:"FinalActionParams,omitempty" xml:"FinalActionParams,omitempty"`
	// This parameter is required.
	FinalPrompt *string `json:"FinalPrompt,omitempty" xml:"FinalPrompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12f407b22cbe4890ac595f09985848d5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Timeout
	IntentTrigger *string `json:"IntentTrigger,omitempty" xml:"IntentTrigger,omitempty"`
	// This parameter is required.
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbotIntent
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ModifySilenceTimeoutConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySilenceTimeoutConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifySilenceTimeoutConfigRequest) GetFinalAction() *string {
	return s.FinalAction
}

func (s *ModifySilenceTimeoutConfigRequest) GetFinalActionParams() *string {
	return s.FinalActionParams
}

func (s *ModifySilenceTimeoutConfigRequest) GetFinalPrompt() *string {
	return s.FinalPrompt
}

func (s *ModifySilenceTimeoutConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifySilenceTimeoutConfigRequest) GetIntentTrigger() *string {
	return s.IntentTrigger
}

func (s *ModifySilenceTimeoutConfigRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *ModifySilenceTimeoutConfigRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ModifySilenceTimeoutConfigRequest) GetThreshold() *int32 {
	return s.Threshold
}

func (s *ModifySilenceTimeoutConfigRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *ModifySilenceTimeoutConfigRequest) SetFinalAction(v string) *ModifySilenceTimeoutConfigRequest {
	s.FinalAction = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetFinalActionParams(v string) *ModifySilenceTimeoutConfigRequest {
	s.FinalActionParams = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetFinalPrompt(v string) *ModifySilenceTimeoutConfigRequest {
	s.FinalPrompt = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetInstanceId(v string) *ModifySilenceTimeoutConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetIntentTrigger(v string) *ModifySilenceTimeoutConfigRequest {
	s.IntentTrigger = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetPrompt(v string) *ModifySilenceTimeoutConfigRequest {
	s.Prompt = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetSourceType(v string) *ModifySilenceTimeoutConfigRequest {
	s.SourceType = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetThreshold(v int32) *ModifySilenceTimeoutConfigRequest {
	s.Threshold = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) SetTimeout(v int64) *ModifySilenceTimeoutConfigRequest {
	s.Timeout = &v
	return s
}

func (s *ModifySilenceTimeoutConfigRequest) Validate() error {
	return dara.Validate(s)
}
