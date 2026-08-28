// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginalFileUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKey(v string) *GetOriginalFileUrlRequest
	GetAPIKey() *string
	SetTaskId(v string) *GetOriginalFileUrlRequest
	GetTaskId() *string
}

type GetOriginalFileUrlRequest struct {
	// The API key that identifies a member account. You can obtain it from the Ruiyi Bao console.
	//
	// example:
	//
	// ***
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

func (s GetOriginalFileUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOriginalFileUrlRequest) GoString() string {
	return s.String()
}

func (s *GetOriginalFileUrlRequest) GetAPIKey() *string {
	return s.APIKey
}

func (s *GetOriginalFileUrlRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetOriginalFileUrlRequest) SetAPIKey(v string) *GetOriginalFileUrlRequest {
	s.APIKey = &v
	return s
}

func (s *GetOriginalFileUrlRequest) SetTaskId(v string) *GetOriginalFileUrlRequest {
	s.TaskId = &v
	return s
}

func (s *GetOriginalFileUrlRequest) Validate() error {
	return dara.Validate(s)
}
