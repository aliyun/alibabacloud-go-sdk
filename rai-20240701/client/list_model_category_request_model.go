// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentSafeImageSupported(v int32) *ListModelCategoryRequest
	GetContentSafeImageSupported() *int32
	SetContentSafeTextSupported(v int32) *ListModelCategoryRequest
	GetContentSafeTextSupported() *int32
	SetModelCategoryName(v string) *ListModelCategoryRequest
	GetModelCategoryName() *string
	SetModelSource(v string) *ListModelCategoryRequest
	GetModelSource() *string
	SetPromptAttackTextSupported(v int32) *ListModelCategoryRequest
	GetPromptAttackTextSupported() *int32
	SetSensitiveTopicTextSupported(v int32) *ListModelCategoryRequest
	GetSensitiveTopicTextSupported() *int32
}

type ListModelCategoryRequest struct {
	// example:
	//
	// False
	ContentSafeImageSupported *int32 `json:"ContentSafeImageSupported,omitempty" xml:"ContentSafeImageSupported,omitempty"`
	// example:
	//
	// True
	ContentSafeTextSupported *int32 `json:"ContentSafeTextSupported,omitempty" xml:"ContentSafeTextSupported,omitempty"`
	// example:
	//
	// demo
	ModelCategoryName *string `json:"ModelCategoryName,omitempty" xml:"ModelCategoryName,omitempty"`
	// example:
	//
	// 1
	ModelSource *string `json:"ModelSource,omitempty" xml:"ModelSource,omitempty"`
	// example:
	//
	// False
	PromptAttackTextSupported *int32 `json:"PromptAttackTextSupported,omitempty" xml:"PromptAttackTextSupported,omitempty"`
	// example:
	//
	// False
	SensitiveTopicTextSupported *int32 `json:"SensitiveTopicTextSupported,omitempty" xml:"SensitiveTopicTextSupported,omitempty"`
}

func (s ListModelCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelCategoryRequest) GoString() string {
	return s.String()
}

func (s *ListModelCategoryRequest) GetContentSafeImageSupported() *int32 {
	return s.ContentSafeImageSupported
}

func (s *ListModelCategoryRequest) GetContentSafeTextSupported() *int32 {
	return s.ContentSafeTextSupported
}

func (s *ListModelCategoryRequest) GetModelCategoryName() *string {
	return s.ModelCategoryName
}

func (s *ListModelCategoryRequest) GetModelSource() *string {
	return s.ModelSource
}

func (s *ListModelCategoryRequest) GetPromptAttackTextSupported() *int32 {
	return s.PromptAttackTextSupported
}

func (s *ListModelCategoryRequest) GetSensitiveTopicTextSupported() *int32 {
	return s.SensitiveTopicTextSupported
}

func (s *ListModelCategoryRequest) SetContentSafeImageSupported(v int32) *ListModelCategoryRequest {
	s.ContentSafeImageSupported = &v
	return s
}

func (s *ListModelCategoryRequest) SetContentSafeTextSupported(v int32) *ListModelCategoryRequest {
	s.ContentSafeTextSupported = &v
	return s
}

func (s *ListModelCategoryRequest) SetModelCategoryName(v string) *ListModelCategoryRequest {
	s.ModelCategoryName = &v
	return s
}

func (s *ListModelCategoryRequest) SetModelSource(v string) *ListModelCategoryRequest {
	s.ModelSource = &v
	return s
}

func (s *ListModelCategoryRequest) SetPromptAttackTextSupported(v int32) *ListModelCategoryRequest {
	s.PromptAttackTextSupported = &v
	return s
}

func (s *ListModelCategoryRequest) SetSensitiveTopicTextSupported(v int32) *ListModelCategoryRequest {
	s.SensitiveTopicTextSupported = &v
	return s
}

func (s *ListModelCategoryRequest) Validate() error {
	return dara.Validate(s)
}
