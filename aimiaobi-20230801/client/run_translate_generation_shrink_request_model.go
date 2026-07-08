// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTranslateGenerationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrompt(v string) *RunTranslateGenerationShrinkRequest
	GetPrompt() *string
	SetReferenceDataShrink(v string) *RunTranslateGenerationShrinkRequest
	GetReferenceDataShrink() *string
	SetTaskId(v string) *RunTranslateGenerationShrinkRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *RunTranslateGenerationShrinkRequest
	GetWorkspaceId() *string
}

type RunTranslateGenerationShrinkRequest struct {
	// The target language for translation. The source language is automatically detected.
	//
	// | Language           | Prompt value |
	//
	// | ------------------ | ------------ |
	//
	// | English            | English      |
	//
	// | Simplified Chinese | Chinese      |
	//
	// | Japanese           | Japanese     |
	//
	// | Korean             | Korean       |
	//
	// | Spanish            | Spanish      |
	//
	// | French             | French       |
	//
	// | Portuguese         | Portuguese   |
	//
	// | German             | German       |
	//
	// | Italian            | Italian      |
	//
	// This parameter is required.
	//
	// example:
	//
	// English
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The data required for generation.
	//
	// This parameter is required.
	ReferenceDataShrink *string `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty"`
	// Optional. The unique ID of the associated creative article.
	//
	// > You do not need to specify TaskId. The system generates one automatically. If subsequent tasks use the same TaskId, they belong to the same conversation group.
	//
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The unique ID of the Alibaba Cloud Model Studio workspace. For more information, see [Get a Workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunTranslateGenerationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunTranslateGenerationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunTranslateGenerationShrinkRequest) GetReferenceDataShrink() *string {
	return s.ReferenceDataShrink
}

func (s *RunTranslateGenerationShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunTranslateGenerationShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunTranslateGenerationShrinkRequest) SetPrompt(v string) *RunTranslateGenerationShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunTranslateGenerationShrinkRequest) SetReferenceDataShrink(v string) *RunTranslateGenerationShrinkRequest {
	s.ReferenceDataShrink = &v
	return s
}

func (s *RunTranslateGenerationShrinkRequest) SetTaskId(v string) *RunTranslateGenerationShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunTranslateGenerationShrinkRequest) SetWorkspaceId(v string) *RunTranslateGenerationShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunTranslateGenerationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
