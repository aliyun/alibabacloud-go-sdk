// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranslationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKey(v string) *GetTranslationTaskRequest
	GetAPIKey() *string
	SetTaskId(v string) *GetTranslationTaskRequest
	GetTaskId() *string
}

type GetTranslationTaskRequest struct {
	APIKey *string `json:"APIKey,omitempty" xml:"APIKey,omitempty"`
	// This parameter is required.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTranslationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTranslationTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTranslationTaskRequest) GetAPIKey() *string {
	return s.APIKey
}

func (s *GetTranslationTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTranslationTaskRequest) SetAPIKey(v string) *GetTranslationTaskRequest {
	s.APIKey = &v
	return s
}

func (s *GetTranslationTaskRequest) SetTaskId(v string) *GetTranslationTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetTranslationTaskRequest) Validate() error {
	return dara.Validate(s)
}
