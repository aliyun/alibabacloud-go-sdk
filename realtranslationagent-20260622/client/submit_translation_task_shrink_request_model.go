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
	APIKey     *string `json:"APIKey,omitempty" xml:"APIKey,omitempty"`
	BaseTaskId *string `json:"BaseTaskId,omitempty" xml:"BaseTaskId,omitempty"`
	// This parameter is required.
	ConfigShrink      *string `json:"Config,omitempty" xml:"Config,omitempty"`
	CustomTermsShrink *string `json:"CustomTerms,omitempty" xml:"CustomTerms,omitempty"`
	TaskId            *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
