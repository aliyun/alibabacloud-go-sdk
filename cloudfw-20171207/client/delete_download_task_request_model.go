// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDownloadTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteDownloadTaskRequest
	GetLang() *string
	SetTaskId(v string) *DeleteDownloadTaskRequest
	GetTaskId() *string
}

type DeleteDownloadTaskRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the file download task.
	//
	// example:
	//
	// 4376
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteDownloadTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDownloadTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteDownloadTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteDownloadTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteDownloadTaskRequest) SetLang(v string) *DeleteDownloadTaskRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDownloadTaskRequest) SetTaskId(v string) *DeleteDownloadTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteDownloadTaskRequest) Validate() error {
	return dara.Validate(s)
}
