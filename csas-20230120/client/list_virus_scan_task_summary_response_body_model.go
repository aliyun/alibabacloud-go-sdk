// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanTaskSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListVirusScanTaskSummaryResponseBody
	GetRequestId() *string
	SetTasks(v []*ListVirusScanTaskSummaryResponseBodyTasks) *ListVirusScanTaskSummaryResponseBody
	GetTasks() []*ListVirusScanTaskSummaryResponseBodyTasks
}

type ListVirusScanTaskSummaryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of detection statistics for virus scan tasks.
	Tasks []*ListVirusScanTaskSummaryResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ListVirusScanTaskSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTaskSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirusScanTaskSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVirusScanTaskSummaryResponseBody) GetTasks() []*ListVirusScanTaskSummaryResponseBodyTasks {
	return s.Tasks
}

func (s *ListVirusScanTaskSummaryResponseBody) SetRequestId(v string) *ListVirusScanTaskSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirusScanTaskSummaryResponseBody) SetTasks(v []*ListVirusScanTaskSummaryResponseBodyTasks) *ListVirusScanTaskSummaryResponseBody {
	s.Tasks = v
	return s
}

func (s *ListVirusScanTaskSummaryResponseBody) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVirusScanTaskSummaryResponseBodyTasks struct {
	// The virus scan task ID.
	//
	// example:
	//
	// v1:1024772
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The number of virus files detected by the task.
	//
	// example:
	//
	// 7
	VirusFileCount *int64 `json:"VirusFileCount,omitempty" xml:"VirusFileCount,omitempty"`
}

func (s ListVirusScanTaskSummaryResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTaskSummaryResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListVirusScanTaskSummaryResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListVirusScanTaskSummaryResponseBodyTasks) GetVirusFileCount() *int64 {
	return s.VirusFileCount
}

func (s *ListVirusScanTaskSummaryResponseBodyTasks) SetTaskId(v string) *ListVirusScanTaskSummaryResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListVirusScanTaskSummaryResponseBodyTasks) SetVirusFileCount(v int64) *ListVirusScanTaskSummaryResponseBodyTasks {
	s.VirusFileCount = &v
	return s
}

func (s *ListVirusScanTaskSummaryResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
