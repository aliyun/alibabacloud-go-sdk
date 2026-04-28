// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearRecycleBinItem interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTaskId(v string) *ClearRecycleBinItem
	GetAsyncTaskId() *string
	SetDomainId(v string) *ClearRecycleBinItem
	GetDomainId() *string
	SetDriveId(v string) *ClearRecycleBinItem
	GetDriveId() *string
	SetTaskId(v string) *ClearRecycleBinItem
	GetTaskId() *string
}

type ClearRecycleBinItem struct {
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	DomainId    *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DriveId     *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	TaskId      *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s ClearRecycleBinItem) String() string {
	return dara.Prettify(s)
}

func (s ClearRecycleBinItem) GoString() string {
	return s.String()
}

func (s *ClearRecycleBinItem) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *ClearRecycleBinItem) GetDomainId() *string {
	return s.DomainId
}

func (s *ClearRecycleBinItem) GetDriveId() *string {
	return s.DriveId
}

func (s *ClearRecycleBinItem) GetTaskId() *string {
	return s.TaskId
}

func (s *ClearRecycleBinItem) SetAsyncTaskId(v string) *ClearRecycleBinItem {
	s.AsyncTaskId = &v
	return s
}

func (s *ClearRecycleBinItem) SetDomainId(v string) *ClearRecycleBinItem {
	s.DomainId = &v
	return s
}

func (s *ClearRecycleBinItem) SetDriveId(v string) *ClearRecycleBinItem {
	s.DriveId = &v
	return s
}

func (s *ClearRecycleBinItem) SetTaskId(v string) *ClearRecycleBinItem {
	s.TaskId = &v
	return s
}

func (s *ClearRecycleBinItem) Validate() error {
	return dara.Validate(s)
}
