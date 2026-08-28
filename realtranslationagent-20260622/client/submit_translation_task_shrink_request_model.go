// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTranslationTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKey(v string) *SubmitTranslationTaskShrinkRequest
	GetAPIKey() *string
	SetBaseTaskId(v string) *SubmitTranslationTaskShrinkRequest
	GetBaseTaskId() *string
	SetConfigShrink(v string) *SubmitTranslationTaskShrinkRequest
	GetConfigShrink() *string
	SetCustomTermsShrink(v string) *SubmitTranslationTaskShrinkRequest
	GetCustomTermsShrink() *string
	SetTaskId(v string) *SubmitTranslationTaskShrinkRequest
	GetTaskId() *string
}

type SubmitTranslationTaskShrinkRequest struct {
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
	ConfigShrink *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The custom terms.
	//
	// 	Notice: Custom terms are for reference only. Actual translation results may differ. Refer to the final output for the definitive result.</notice>
	//
	// example:
	//
	// [{"SourceTerm":"dog", "TargetTerm":"dog"}]
	CustomTermsShrink *string `json:"CustomTerms,omitempty" xml:"CustomTerms,omitempty"`
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

func (s SubmitTranslationTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranslationTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitTranslationTaskShrinkRequest) GetAPIKey() *string {
	return s.APIKey
}

func (s *SubmitTranslationTaskShrinkRequest) GetBaseTaskId() *string {
	return s.BaseTaskId
}

func (s *SubmitTranslationTaskShrinkRequest) GetConfigShrink() *string {
	return s.ConfigShrink
}

func (s *SubmitTranslationTaskShrinkRequest) GetCustomTermsShrink() *string {
	return s.CustomTermsShrink
}

func (s *SubmitTranslationTaskShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitTranslationTaskShrinkRequest) SetAPIKey(v string) *SubmitTranslationTaskShrinkRequest {
	s.APIKey = &v
	return s
}

func (s *SubmitTranslationTaskShrinkRequest) SetBaseTaskId(v string) *SubmitTranslationTaskShrinkRequest {
	s.BaseTaskId = &v
	return s
}

func (s *SubmitTranslationTaskShrinkRequest) SetConfigShrink(v string) *SubmitTranslationTaskShrinkRequest {
	s.ConfigShrink = &v
	return s
}

func (s *SubmitTranslationTaskShrinkRequest) SetCustomTermsShrink(v string) *SubmitTranslationTaskShrinkRequest {
	s.CustomTermsShrink = &v
	return s
}

func (s *SubmitTranslationTaskShrinkRequest) SetTaskId(v string) *SubmitTranslationTaskShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *SubmitTranslationTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
