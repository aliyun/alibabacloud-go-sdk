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
	APIKey *string `json:"APIKey,omitempty" xml:"APIKey,omitempty"`
	// This parameter is required.
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
