// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTranslationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKey(v string) *SubmitTranslationTaskRequest
	GetAPIKey() *string
	SetBaseTaskId(v string) *SubmitTranslationTaskRequest
	GetBaseTaskId() *string
	SetConfig(v *SubmitTranslationTaskRequestConfig) *SubmitTranslationTaskRequest
	GetConfig() *SubmitTranslationTaskRequestConfig
	SetCustomTerms(v []*SubmitTranslationTaskRequestCustomTerms) *SubmitTranslationTaskRequest
	GetCustomTerms() []*SubmitTranslationTaskRequestCustomTerms
	SetTaskId(v string) *SubmitTranslationTaskRequest
	GetTaskId() *string
}

type SubmitTranslationTaskRequest struct {
	// The API key that identifies the identity of the member account. You can obtain this from the RuiYiBao console.
	//
	// example:
	//
	// ***
	APIKey *string `json:"APIKey,omitempty" xml:"APIKey,omitempty"`
	// The translation task ID of a previously submitted translation task. Pass in this parameter when resubmitting a translation task.
	//
	// - You must pass in either this parameter or TaskId.
	//
	// example:
	//
	// f9c35b0453b
	BaseTaskId *string `json:"BaseTaskId,omitempty" xml:"BaseTaskId,omitempty"`
	// The translation configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//         "SourceLanguage": "zh",
	//
	//         "TargetLanguage": "en",
	//
	//         "Style": "minimal",
	//
	//         "Font": "Arial"
	//
	//     }
	Config *SubmitTranslationTaskRequestConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The custom terms.
	//
	// 	Notice: Custom terms are for reference only. Actual translation results may differ. Refer to the final output for the definitive result.</notice>
	//
	// example:
	//
	// [{"SourceTerm":"dog", "TargetTerm":"dog"}]
	CustomTerms []*SubmitTranslationTaskRequestCustomTerms `json:"CustomTerms,omitempty" xml:"CustomTerms,omitempty" type:"Repeated"`
	// The translation task ID.
	//
	// - Obtained from the TaskId returned by UploadTranslationFile.
	//
	// - You must pass in either this parameter or BaseTaskId.
	//
	// example:
	//
	// f9c35b0453b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitTranslationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranslationTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitTranslationTaskRequest) GetAPIKey() *string {
	return s.APIKey
}

func (s *SubmitTranslationTaskRequest) GetBaseTaskId() *string {
	return s.BaseTaskId
}

func (s *SubmitTranslationTaskRequest) GetConfig() *SubmitTranslationTaskRequestConfig {
	return s.Config
}

func (s *SubmitTranslationTaskRequest) GetCustomTerms() []*SubmitTranslationTaskRequestCustomTerms {
	return s.CustomTerms
}

func (s *SubmitTranslationTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitTranslationTaskRequest) SetAPIKey(v string) *SubmitTranslationTaskRequest {
	s.APIKey = &v
	return s
}

func (s *SubmitTranslationTaskRequest) SetBaseTaskId(v string) *SubmitTranslationTaskRequest {
	s.BaseTaskId = &v
	return s
}

func (s *SubmitTranslationTaskRequest) SetConfig(v *SubmitTranslationTaskRequestConfig) *SubmitTranslationTaskRequest {
	s.Config = v
	return s
}

func (s *SubmitTranslationTaskRequest) SetCustomTerms(v []*SubmitTranslationTaskRequestCustomTerms) *SubmitTranslationTaskRequest {
	s.CustomTerms = v
	return s
}

func (s *SubmitTranslationTaskRequest) SetTaskId(v string) *SubmitTranslationTaskRequest {
	s.TaskId = &v
	return s
}

func (s *SubmitTranslationTaskRequest) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	if s.CustomTerms != nil {
		for _, item := range s.CustomTerms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitTranslationTaskRequestConfig struct {
	Agent *string `json:"Agent,omitempty" xml:"Agent,omitempty"`
	// The agent ID passed to the Agent Console platform.
	//
	// example:
	//
	// app-11111111
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The font.
	//
	// - For new tasks, obtain this from UploadTranslationFile.
	//
	// - For retranslation of historical tasks, obtain this from GetTranslationTask.
	//
	// example:
	//
	// Arial
	Font *string `json:"Font,omitempty" xml:"Font,omitempty"`
	// The language of the source file.
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The translation style. Takes effect only when the translation file is a PPT file.
	//
	// example:
	//
	// minimal
	Style *string `json:"Style,omitempty" xml:"Style,omitempty"`
	// The target language.
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s SubmitTranslationTaskRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranslationTaskRequestConfig) GoString() string {
	return s.String()
}

func (s *SubmitTranslationTaskRequestConfig) GetAgent() *string {
	return s.Agent
}

func (s *SubmitTranslationTaskRequestConfig) GetAgentId() *string {
	return s.AgentId
}

func (s *SubmitTranslationTaskRequestConfig) GetFont() *string {
	return s.Font
}

func (s *SubmitTranslationTaskRequestConfig) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *SubmitTranslationTaskRequestConfig) GetStyle() *string {
	return s.Style
}

func (s *SubmitTranslationTaskRequestConfig) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *SubmitTranslationTaskRequestConfig) SetAgent(v string) *SubmitTranslationTaskRequestConfig {
	s.Agent = &v
	return s
}

func (s *SubmitTranslationTaskRequestConfig) SetAgentId(v string) *SubmitTranslationTaskRequestConfig {
	s.AgentId = &v
	return s
}

func (s *SubmitTranslationTaskRequestConfig) SetFont(v string) *SubmitTranslationTaskRequestConfig {
	s.Font = &v
	return s
}

func (s *SubmitTranslationTaskRequestConfig) SetSourceLanguage(v string) *SubmitTranslationTaskRequestConfig {
	s.SourceLanguage = &v
	return s
}

func (s *SubmitTranslationTaskRequestConfig) SetStyle(v string) *SubmitTranslationTaskRequestConfig {
	s.Style = &v
	return s
}

func (s *SubmitTranslationTaskRequestConfig) SetTargetLanguage(v string) *SubmitTranslationTaskRequestConfig {
	s.TargetLanguage = &v
	return s
}

func (s *SubmitTranslationTaskRequestConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitTranslationTaskRequestCustomTerms struct {
	// The source term.
	//
	// example:
	//
	// dog
	SourceTerm *string `json:"SourceTerm,omitempty" xml:"SourceTerm,omitempty"`
	// The target term.
	//
	// example:
	//
	// dog
	TargetTerm *string `json:"TargetTerm,omitempty" xml:"TargetTerm,omitempty"`
}

func (s SubmitTranslationTaskRequestCustomTerms) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranslationTaskRequestCustomTerms) GoString() string {
	return s.String()
}

func (s *SubmitTranslationTaskRequestCustomTerms) GetSourceTerm() *string {
	return s.SourceTerm
}

func (s *SubmitTranslationTaskRequestCustomTerms) GetTargetTerm() *string {
	return s.TargetTerm
}

func (s *SubmitTranslationTaskRequestCustomTerms) SetSourceTerm(v string) *SubmitTranslationTaskRequestCustomTerms {
	s.SourceTerm = &v
	return s
}

func (s *SubmitTranslationTaskRequestCustomTerms) SetTargetTerm(v string) *SubmitTranslationTaskRequestCustomTerms {
	s.TargetTerm = &v
	return s
}

func (s *SubmitTranslationTaskRequestCustomTerms) Validate() error {
	return dara.Validate(s)
}
