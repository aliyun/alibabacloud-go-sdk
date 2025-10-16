// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadTaskTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDownloadTaskTypeResponseBody
	GetRequestId() *string
	SetTaskTypeArray(v []*DescribeDownloadTaskTypeResponseBodyTaskTypeArray) *DescribeDownloadTaskTypeResponseBody
	GetTaskTypeArray() []*DescribeDownloadTaskTypeResponseBodyTaskTypeArray
	SetTotalCount(v int32) *DescribeDownloadTaskTypeResponseBody
	GetTotalCount() *int32
}

type DescribeDownloadTaskTypeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B835494C-D093-5524-BBDE-BD272077B40E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task types.
	TaskTypeArray []*DescribeDownloadTaskTypeResponseBodyTaskTypeArray `json:"TaskTypeArray,omitempty" xml:"TaskTypeArray,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDownloadTaskTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadTaskTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDownloadTaskTypeResponseBody) GetTaskTypeArray() []*DescribeDownloadTaskTypeResponseBodyTaskTypeArray {
	return s.TaskTypeArray
}

func (s *DescribeDownloadTaskTypeResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDownloadTaskTypeResponseBody) SetRequestId(v string) *DescribeDownloadTaskTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDownloadTaskTypeResponseBody) SetTaskTypeArray(v []*DescribeDownloadTaskTypeResponseBodyTaskTypeArray) *DescribeDownloadTaskTypeResponseBody {
	s.TaskTypeArray = v
	return s
}

func (s *DescribeDownloadTaskTypeResponseBody) SetTotalCount(v int32) *DescribeDownloadTaskTypeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDownloadTaskTypeResponseBody) Validate() error {
	if s.TaskTypeArray != nil {
		for _, item := range s.TaskTypeArray {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDownloadTaskTypeResponseBodyTaskTypeArray struct {
	// The name of the task type.
	//
	// example:
	//
	// Internet Boundary Firewall Assets
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the task.
	//
	// example:
	//
	// InternetFirewallAsset
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDownloadTaskTypeResponseBodyTaskTypeArray) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadTaskTypeResponseBodyTaskTypeArray) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskTypeResponseBodyTaskTypeArray) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeDownloadTaskTypeResponseBodyTaskTypeArray) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeDownloadTaskTypeResponseBodyTaskTypeArray) SetTaskName(v string) *DescribeDownloadTaskTypeResponseBodyTaskTypeArray {
	s.TaskName = &v
	return s
}

func (s *DescribeDownloadTaskTypeResponseBodyTaskTypeArray) SetTaskType(v string) *DescribeDownloadTaskTypeResponseBodyTaskTypeArray {
	s.TaskType = &v
	return s
}

func (s *DescribeDownloadTaskTypeResponseBodyTaskTypeArray) Validate() error {
	return dara.Validate(s)
}
