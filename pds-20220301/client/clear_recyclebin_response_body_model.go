// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearRecyclebinResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTaskId(v string) *ClearRecyclebinResponseBody
	GetAsyncTaskId() *string
	SetDomainId(v string) *ClearRecyclebinResponseBody
	GetDomainId() *string
	SetDriveId(v string) *ClearRecyclebinResponseBody
	GetDriveId() *string
}

type ClearRecyclebinResponseBody struct {
	// The ID of the asynchronous task.
	//
	// You can call the GetAsyncTask operation to query the information about the asynchronous task based on the task ID.
	//
	// example:
	//
	// 13ebd3a24dba4166b1527add676ef2866051b4d5dele16
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
}

func (s ClearRecyclebinResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearRecyclebinResponseBody) GoString() string {
	return s.String()
}

func (s *ClearRecyclebinResponseBody) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *ClearRecyclebinResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *ClearRecyclebinResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *ClearRecyclebinResponseBody) SetAsyncTaskId(v string) *ClearRecyclebinResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *ClearRecyclebinResponseBody) SetDomainId(v string) *ClearRecyclebinResponseBody {
	s.DomainId = &v
	return s
}

func (s *ClearRecyclebinResponseBody) SetDriveId(v string) *ClearRecyclebinResponseBody {
	s.DriveId = &v
	return s
}

func (s *ClearRecyclebinResponseBody) Validate() error {
	return dara.Validate(s)
}
