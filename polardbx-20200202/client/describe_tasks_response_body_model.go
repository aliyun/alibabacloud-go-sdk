// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeTasksResponseBodyItems) *DescribeTasksResponseBody
	GetItems() []*DescribeTasksResponseBodyItems
	SetPageNumber(v int32) *DescribeTasksResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeTasksResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeTasksResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeTasksResponseBody
	GetTotalRecordCount() *int32
}

type DescribeTasksResponseBody struct {
	Items []*DescribeTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// D6045D78-C119-5A17-9DEA-0F810394E008
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBody) GetItems() []*DescribeTasksResponseBodyItems {
	return s.Items
}

func (s *DescribeTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTasksResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTasksResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeTasksResponseBody) SetItems(v []*DescribeTasksResponseBodyItems) *DescribeTasksResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTasksResponseBody) SetPageNumber(v int32) *DescribeTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTasksResponseBody) SetPageRecordCount(v int32) *DescribeTasksResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeTasksResponseBody) SetRequestId(v string) *DescribeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTasksResponseBody) SetTotalRecordCount(v int32) *DescribeTasksResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeTasksResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTasksResponseBodyItems struct {
	// example:
	//
	// 2021-10-20T19:40:00Z
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// DBName
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// example:
	//
	// 2021-10-20T20:00:00Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// 80
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// ProgressInfo
	ProgressInfo *string `json:"ProgressInfo,omitempty" xml:"ProgressInfo,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	ScaleOutToken *string `json:"ScaleOutToken,omitempty" xml:"ScaleOutToken,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// multi_scale_out
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// example:
	//
	// TaskErrorCode
	TaskErrorCode *string `json:"TaskErrorCode,omitempty" xml:"TaskErrorCode,omitempty"`
	// example:
	//
	// TaskErrorMessage
	TaskErrorMessage *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	// example:
	//
	// 20089398
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTasksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyItems) GetBeginTime() *string {
	return s.BeginTime
}

func (s *DescribeTasksResponseBodyItems) GetDBName() *string {
	return s.DBName
}

func (s *DescribeTasksResponseBodyItems) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeTasksResponseBodyItems) GetProgress() *string {
	return s.Progress
}

func (s *DescribeTasksResponseBodyItems) GetProgressInfo() *string {
	return s.ProgressInfo
}

func (s *DescribeTasksResponseBodyItems) GetScaleOutToken() *string {
	return s.ScaleOutToken
}

func (s *DescribeTasksResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeTasksResponseBodyItems) GetTaskAction() *string {
	return s.TaskAction
}

func (s *DescribeTasksResponseBodyItems) GetTaskErrorCode() *string {
	return s.TaskErrorCode
}

func (s *DescribeTasksResponseBodyItems) GetTaskErrorMessage() *string {
	return s.TaskErrorMessage
}

func (s *DescribeTasksResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTasksResponseBodyItems) SetBeginTime(v string) *DescribeTasksResponseBodyItems {
	s.BeginTime = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetDBName(v string) *DescribeTasksResponseBodyItems {
	s.DBName = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetFinishTime(v string) *DescribeTasksResponseBodyItems {
	s.FinishTime = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetProgress(v string) *DescribeTasksResponseBodyItems {
	s.Progress = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetProgressInfo(v string) *DescribeTasksResponseBodyItems {
	s.ProgressInfo = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetScaleOutToken(v string) *DescribeTasksResponseBodyItems {
	s.ScaleOutToken = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetStatus(v string) *DescribeTasksResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetTaskAction(v string) *DescribeTasksResponseBodyItems {
	s.TaskAction = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetTaskErrorCode(v string) *DescribeTasksResponseBodyItems {
	s.TaskErrorCode = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetTaskErrorMessage(v string) *DescribeTasksResponseBodyItems {
	s.TaskErrorMessage = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) SetTaskId(v string) *DescribeTasksResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeTasksResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
