// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTranslationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKey(v string) *CancelTranslationTaskRequest
	GetAPIKey() *string
	SetTaskId(v string) *CancelTranslationTaskRequest
	GetTaskId() *string
}

type CancelTranslationTaskRequest struct {
	// The API key, the identity of member accounts. You can obtain it from the RuiYiBao console.
	//
	// example:
	//
	// sk-5****7
	APIKey *string `json:"APIKey,omitempty" xml:"APIKey,omitempty"`
	// The task ID.
	//
	// - The TaskId is returned after a task is submitted by calling SubmitTranslationTask.
	//
	// - Only running tasks can be canceled.
	//
	// This parameter is required.
	//
	// example:
	//
	// f9c35b0453b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelTranslationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelTranslationTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelTranslationTaskRequest) GetAPIKey() *string {
	return s.APIKey
}

func (s *CancelTranslationTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelTranslationTaskRequest) SetAPIKey(v string) *CancelTranslationTaskRequest {
	s.APIKey = &v
	return s
}

func (s *CancelTranslationTaskRequest) SetTaskId(v string) *CancelTranslationTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CancelTranslationTaskRequest) Validate() error {
	return dara.Validate(s)
}
