// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArchiveFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTaskId(v string) *ArchiveFilesResponseBody
	GetAsyncTaskId() *string
}

type ArchiveFilesResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 000e89fb-cf8f-11e9-8ab4-b6e980803a3b
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
}

func (s ArchiveFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ArchiveFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ArchiveFilesResponseBody) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *ArchiveFilesResponseBody) SetAsyncTaskId(v string) *ArchiveFilesResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *ArchiveFilesResponseBody) Validate() error {
	return dara.Validate(s)
}
