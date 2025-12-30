// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransferFileDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileIds(v []*string) *ListTransferFileDownloadUrlRequest
	GetFileIds() []*string
	SetTaskId(v string) *ListTransferFileDownloadUrlRequest
	GetTaskId() *string
}

type ListTransferFileDownloadUrlRequest struct {
	FileIds []*string `json:"FileIds,omitempty" xml:"FileIds,omitempty" type:"Repeated"`
	// example:
	//
	// trt-hffhi4nmqoi4****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListTransferFileDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransferFileDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *ListTransferFileDownloadUrlRequest) GetFileIds() []*string {
	return s.FileIds
}

func (s *ListTransferFileDownloadUrlRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTransferFileDownloadUrlRequest) SetFileIds(v []*string) *ListTransferFileDownloadUrlRequest {
	s.FileIds = v
	return s
}

func (s *ListTransferFileDownloadUrlRequest) SetTaskId(v string) *ListTransferFileDownloadUrlRequest {
	s.TaskId = &v
	return s
}

func (s *ListTransferFileDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
