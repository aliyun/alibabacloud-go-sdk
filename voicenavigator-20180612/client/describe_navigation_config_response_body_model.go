// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNavigationConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGreetingConfig(v *DescribeNavigationConfigResponseBodyGreetingConfig) *DescribeNavigationConfigResponseBody
	GetGreetingConfig() *DescribeNavigationConfigResponseBodyGreetingConfig
	SetRequestId(v string) *DescribeNavigationConfigResponseBody
	GetRequestId() *string
	SetSilenceTimeoutConfig(v *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) *DescribeNavigationConfigResponseBody
	GetSilenceTimeoutConfig() *DescribeNavigationConfigResponseBodySilenceTimeoutConfig
	SetUnrecognizingConfig(v *DescribeNavigationConfigResponseBodyUnrecognizingConfig) *DescribeNavigationConfigResponseBody
	GetUnrecognizingConfig() *DescribeNavigationConfigResponseBodyUnrecognizingConfig
}

type DescribeNavigationConfigResponseBody struct {
	GreetingConfig *DescribeNavigationConfigResponseBodyGreetingConfig `json:"GreetingConfig,omitempty" xml:"GreetingConfig,omitempty" type:"Struct"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId            *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SilenceTimeoutConfig *DescribeNavigationConfigResponseBodySilenceTimeoutConfig `json:"SilenceTimeoutConfig,omitempty" xml:"SilenceTimeoutConfig,omitempty" type:"Struct"`
	UnrecognizingConfig  *DescribeNavigationConfigResponseBodyUnrecognizingConfig  `json:"UnrecognizingConfig,omitempty" xml:"UnrecognizingConfig,omitempty" type:"Struct"`
}

func (s DescribeNavigationConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNavigationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNavigationConfigResponseBody) GetGreetingConfig() *DescribeNavigationConfigResponseBodyGreetingConfig {
	return s.GreetingConfig
}

func (s *DescribeNavigationConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNavigationConfigResponseBody) GetSilenceTimeoutConfig() *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	return s.SilenceTimeoutConfig
}

func (s *DescribeNavigationConfigResponseBody) GetUnrecognizingConfig() *DescribeNavigationConfigResponseBodyUnrecognizingConfig {
	return s.UnrecognizingConfig
}

func (s *DescribeNavigationConfigResponseBody) SetGreetingConfig(v *DescribeNavigationConfigResponseBodyGreetingConfig) *DescribeNavigationConfigResponseBody {
	s.GreetingConfig = v
	return s
}

func (s *DescribeNavigationConfigResponseBody) SetRequestId(v string) *DescribeNavigationConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNavigationConfigResponseBody) SetSilenceTimeoutConfig(v *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) *DescribeNavigationConfigResponseBody {
	s.SilenceTimeoutConfig = v
	return s
}

func (s *DescribeNavigationConfigResponseBody) SetUnrecognizingConfig(v *DescribeNavigationConfigResponseBodyUnrecognizingConfig) *DescribeNavigationConfigResponseBody {
	s.UnrecognizingConfig = v
	return s
}

func (s *DescribeNavigationConfigResponseBody) Validate() error {
	if s.GreetingConfig != nil {
		if err := s.GreetingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SilenceTimeoutConfig != nil {
		if err := s.SilenceTimeoutConfig.Validate(); err != nil {
			return err
		}
	}
	if s.UnrecognizingConfig != nil {
		if err := s.UnrecognizingConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNavigationConfigResponseBodyGreetingConfig struct {
	GreetingWords *string `json:"GreetingWords,omitempty" xml:"GreetingWords,omitempty"`
	IntentTrigger *string `json:"IntentTrigger,omitempty" xml:"IntentTrigger,omitempty"`
	// example:
	//
	// chatbotIntent
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribeNavigationConfigResponseBodyGreetingConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeNavigationConfigResponseBodyGreetingConfig) GoString() string {
	return s.String()
}

func (s *DescribeNavigationConfigResponseBodyGreetingConfig) GetGreetingWords() *string {
	return s.GreetingWords
}

func (s *DescribeNavigationConfigResponseBodyGreetingConfig) GetIntentTrigger() *string {
	return s.IntentTrigger
}

func (s *DescribeNavigationConfigResponseBodyGreetingConfig) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeNavigationConfigResponseBodyGreetingConfig) SetGreetingWords(v string) *DescribeNavigationConfigResponseBodyGreetingConfig {
	s.GreetingWords = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyGreetingConfig) SetIntentTrigger(v string) *DescribeNavigationConfigResponseBodyGreetingConfig {
	s.IntentTrigger = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyGreetingConfig) SetSourceType(v string) *DescribeNavigationConfigResponseBodyGreetingConfig {
	s.SourceType = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyGreetingConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeNavigationConfigResponseBodySilenceTimeoutConfig struct {
	// example:
	//
	// HangUp
	FinalAction *string `json:"FinalAction,omitempty" xml:"FinalAction,omitempty"`
	// example:
	//
	// {}
	FinalActionParams *string `json:"FinalActionParams,omitempty" xml:"FinalActionParams,omitempty"`
	FinalPrompt       *string `json:"FinalPrompt,omitempty" xml:"FinalPrompt,omitempty"`
	// ""
	//
	// example:
	//
	// ""
	IntentTrigger *string `json:"IntentTrigger,omitempty" xml:"IntentTrigger,omitempty"`
	Prompt        *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// chatbotIntent
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 3
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// example:
	//
	// 10
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s DescribeNavigationConfigResponseBodySilenceTimeoutConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeNavigationConfigResponseBodySilenceTimeoutConfig) GoString() string {
	return s.String()
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) GetFinalAction() *string {
	return s.FinalAction
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) GetFinalActionParams() *string {
	return s.FinalActionParams
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) GetFinalPrompt() *string {
	return s.FinalPrompt
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) GetIntentTrigger() *string {
	return s.IntentTrigger
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) GetPrompt() *string {
	return s.Prompt
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) GetThreshold() *int32 {
	return s.Threshold
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) GetTimeout() *int64 {
	return s.Timeout
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetFinalAction(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.FinalAction = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetFinalActionParams(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.FinalActionParams = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetFinalPrompt(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.FinalPrompt = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetIntentTrigger(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.IntentTrigger = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetPrompt(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.Prompt = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetSourceType(v string) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.SourceType = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetThreshold(v int32) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.Threshold = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) SetTimeout(v int64) *DescribeNavigationConfigResponseBodySilenceTimeoutConfig {
	s.Timeout = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodySilenceTimeoutConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeNavigationConfigResponseBodyUnrecognizingConfig struct {
	// example:
	//
	// TransferToAgent
	FinalAction *string `json:"FinalAction,omitempty" xml:"FinalAction,omitempty"`
	// example:
	//
	// { \\"skillGroupId\\": \\"fallbackSkillGroup\\" }
	FinalActionParams *string `json:"FinalActionParams,omitempty" xml:"FinalActionParams,omitempty"`
	FinalPrompt       *string `json:"FinalPrompt,omitempty" xml:"FinalPrompt,omitempty"`
	Prompt            *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 3
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeNavigationConfigResponseBodyUnrecognizingConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeNavigationConfigResponseBodyUnrecognizingConfig) GoString() string {
	return s.String()
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) GetFinalAction() *string {
	return s.FinalAction
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) GetFinalActionParams() *string {
	return s.FinalActionParams
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) GetFinalPrompt() *string {
	return s.FinalPrompt
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) GetPrompt() *string {
	return s.Prompt
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) GetThreshold() *int32 {
	return s.Threshold
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) SetFinalAction(v string) *DescribeNavigationConfigResponseBodyUnrecognizingConfig {
	s.FinalAction = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) SetFinalActionParams(v string) *DescribeNavigationConfigResponseBodyUnrecognizingConfig {
	s.FinalActionParams = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) SetFinalPrompt(v string) *DescribeNavigationConfigResponseBodyUnrecognizingConfig {
	s.FinalPrompt = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) SetPrompt(v string) *DescribeNavigationConfigResponseBodyUnrecognizingConfig {
	s.Prompt = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) SetThreshold(v int32) *DescribeNavigationConfigResponseBodyUnrecognizingConfig {
	s.Threshold = &v
	return s
}

func (s *DescribeNavigationConfigResponseBodyUnrecognizingConfig) Validate() error {
	return dara.Validate(s)
}
