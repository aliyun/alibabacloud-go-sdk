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
	APIKey *string `json:"APIKey,omitempty" xml:"APIKey,omitempty"`
	// This parameter is required.
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
