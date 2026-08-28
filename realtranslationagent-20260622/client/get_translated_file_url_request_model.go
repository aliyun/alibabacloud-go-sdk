// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranslatedFileUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKey(v string) *GetTranslatedFileUrlRequest
	GetAPIKey() *string
	SetTaskId(v string) *GetTranslatedFileUrlRequest
	GetTaskId() *string
}

type GetTranslatedFileUrlRequest struct {
	// The API key that identifies a member accounts identity. You can obtain the key from the RuiYiBao console.
	//
	// example:
	//
	// sk-1***s
	APIKey *string `json:"APIKey,omitempty" xml:"APIKey,omitempty"`
	// The task ID.
	//
	// - The TaskId is returned after a task is submitted by calling SubmitTranslationTask.
	//
	// This parameter is required.
	//
	// example:
	//
	// f9c35b0453b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTranslatedFileUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTranslatedFileUrlRequest) GoString() string {
	return s.String()
}

func (s *GetTranslatedFileUrlRequest) GetAPIKey() *string {
	return s.APIKey
}

func (s *GetTranslatedFileUrlRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTranslatedFileUrlRequest) SetAPIKey(v string) *GetTranslatedFileUrlRequest {
	s.APIKey = &v
	return s
}

func (s *GetTranslatedFileUrlRequest) SetTaskId(v string) *GetTranslatedFileUrlRequest {
	s.TaskId = &v
	return s
}

func (s *GetTranslatedFileUrlRequest) Validate() error {
	return dara.Validate(s)
}
