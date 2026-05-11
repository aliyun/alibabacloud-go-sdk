// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGreetingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGreetingWords(v string) *ModifyGreetingConfigRequest
	GetGreetingWords() *string
	SetInstanceId(v string) *ModifyGreetingConfigRequest
	GetInstanceId() *string
	SetIntentTrigger(v string) *ModifyGreetingConfigRequest
	GetIntentTrigger() *string
	SetSourceType(v string) *ModifyGreetingConfigRequest
	GetSourceType() *string
}

type ModifyGreetingConfigRequest struct {
	// This parameter is required.
	GreetingWords *string `json:"GreetingWords,omitempty" xml:"GreetingWords,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12f407b22cbe4890ac595f09985848d5
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntentTrigger *string `json:"IntentTrigger,omitempty" xml:"IntentTrigger,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbotIntent
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ModifyGreetingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyGreetingConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyGreetingConfigRequest) GetGreetingWords() *string {
	return s.GreetingWords
}

func (s *ModifyGreetingConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyGreetingConfigRequest) GetIntentTrigger() *string {
	return s.IntentTrigger
}

func (s *ModifyGreetingConfigRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ModifyGreetingConfigRequest) SetGreetingWords(v string) *ModifyGreetingConfigRequest {
	s.GreetingWords = &v
	return s
}

func (s *ModifyGreetingConfigRequest) SetInstanceId(v string) *ModifyGreetingConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyGreetingConfigRequest) SetIntentTrigger(v string) *ModifyGreetingConfigRequest {
	s.IntentTrigger = &v
	return s
}

func (s *ModifyGreetingConfigRequest) SetSourceType(v string) *ModifyGreetingConfigRequest {
	s.SourceType = &v
	return s
}

func (s *ModifyGreetingConfigRequest) Validate() error {
	return dara.Validate(s)
}
