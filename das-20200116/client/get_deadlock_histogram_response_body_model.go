// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeadlockHistogramResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetDeadlockHistogramResponseBody
	GetCode() *int64
	SetData(v []*GetDeadlockHistogramResponseBodyData) *GetDeadlockHistogramResponseBody
	GetData() []*GetDeadlockHistogramResponseBodyData
	SetMessage(v string) *GetDeadlockHistogramResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeadlockHistogramResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDeadlockHistogramResponseBody
	GetSuccess() *bool
}

type GetDeadlockHistogramResponseBody struct {
	// example:
	//
	// 200
	Code *int64                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetDeadlockHistogramResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0A74B755-98B7-59DB-8724-1321B394****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDeadlockHistogramResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeadlockHistogramResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeadlockHistogramResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetDeadlockHistogramResponseBody) GetData() []*GetDeadlockHistogramResponseBodyData {
	return s.Data
}

func (s *GetDeadlockHistogramResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeadlockHistogramResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeadlockHistogramResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDeadlockHistogramResponseBody) SetCode(v int64) *GetDeadlockHistogramResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeadlockHistogramResponseBody) SetData(v []*GetDeadlockHistogramResponseBodyData) *GetDeadlockHistogramResponseBody {
	s.Data = v
	return s
}

func (s *GetDeadlockHistogramResponseBody) SetMessage(v string) *GetDeadlockHistogramResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeadlockHistogramResponseBody) SetRequestId(v string) *GetDeadlockHistogramResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeadlockHistogramResponseBody) SetSuccess(v bool) *GetDeadlockHistogramResponseBody {
	s.Success = &v
	return s
}

func (s *GetDeadlockHistogramResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDeadlockHistogramResponseBodyData struct {
	// example:
	//
	// 1729998000000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// pc-bp1u5mas9exx7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2
	LockNumber *int32 `json:"LockNumber,omitempty" xml:"LockNumber,omitempty"`
	// example:
	//
	// pi-bp16v3824rt73****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 1729994400000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-1321B394****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 108************
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetDeadlockHistogramResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDeadlockHistogramResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeadlockHistogramResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *GetDeadlockHistogramResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDeadlockHistogramResponseBodyData) GetLockNumber() *int32 {
	return s.LockNumber
}

func (s *GetDeadlockHistogramResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *GetDeadlockHistogramResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *GetDeadlockHistogramResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetDeadlockHistogramResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDeadlockHistogramResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *GetDeadlockHistogramResponseBodyData) SetEndTime(v string) *GetDeadlockHistogramResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetDeadlockHistogramResponseBodyData) SetInstanceId(v string) *GetDeadlockHistogramResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetDeadlockHistogramResponseBodyData) SetLockNumber(v int32) *GetDeadlockHistogramResponseBodyData {
	s.LockNumber = &v
	return s
}

func (s *GetDeadlockHistogramResponseBodyData) SetNodeId(v string) *GetDeadlockHistogramResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *GetDeadlockHistogramResponseBodyData) SetStartTime(v string) *GetDeadlockHistogramResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetDeadlockHistogramResponseBodyData) SetStatus(v string) *GetDeadlockHistogramResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDeadlockHistogramResponseBodyData) SetTaskId(v string) *GetDeadlockHistogramResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetDeadlockHistogramResponseBodyData) SetUserId(v string) *GetDeadlockHistogramResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetDeadlockHistogramResponseBodyData) Validate() error {
	return dara.Validate(s)
}
