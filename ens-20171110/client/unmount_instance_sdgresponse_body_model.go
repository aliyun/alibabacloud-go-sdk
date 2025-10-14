// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnmountInstanceSDGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UnmountInstanceSDGResponseBodyData) *UnmountInstanceSDGResponseBody
	GetData() *UnmountInstanceSDGResponseBodyData
	SetRequestId(v string) *UnmountInstanceSDGResponseBody
	GetRequestId() *string
}

type UnmountInstanceSDGResponseBody struct {
	// The returned data object.
	Data *UnmountInstanceSDGResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 125B04C7-3D0D-4245-AF96-14E3758E3F06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnmountInstanceSDGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnmountInstanceSDGResponseBody) GoString() string {
	return s.String()
}

func (s *UnmountInstanceSDGResponseBody) GetData() *UnmountInstanceSDGResponseBodyData {
	return s.Data
}

func (s *UnmountInstanceSDGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnmountInstanceSDGResponseBody) SetData(v *UnmountInstanceSDGResponseBodyData) *UnmountInstanceSDGResponseBody {
	s.Data = v
	return s
}

func (s *UnmountInstanceSDGResponseBody) SetRequestId(v string) *UnmountInstanceSDGResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnmountInstanceSDGResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UnmountInstanceSDGResponseBodyData struct {
	// The response message. Success is returned for a successful request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution result of the synchronization request.
	Result *UnmountInstanceSDGResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether all tasks are successful. Valid values:
	//
	// 	- true: All tasks are successful.
	//
	// 	- false: Failed tasks exist.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnmountInstanceSDGResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UnmountInstanceSDGResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnmountInstanceSDGResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *UnmountInstanceSDGResponseBodyData) GetResult() *UnmountInstanceSDGResponseBodyDataResult {
	return s.Result
}

func (s *UnmountInstanceSDGResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UnmountInstanceSDGResponseBodyData) SetMessage(v string) *UnmountInstanceSDGResponseBodyData {
	s.Message = &v
	return s
}

func (s *UnmountInstanceSDGResponseBodyData) SetResult(v *UnmountInstanceSDGResponseBodyDataResult) *UnmountInstanceSDGResponseBodyData {
	s.Result = v
	return s
}

func (s *UnmountInstanceSDGResponseBodyData) SetSuccess(v bool) *UnmountInstanceSDGResponseBodyData {
	s.Success = &v
	return s
}

func (s *UnmountInstanceSDGResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UnmountInstanceSDGResponseBodyDataResult struct {
	// The number of failed tasks.
	//
	// example:
	//
	// 0
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// Details about failed tasks.
	FailedItems []*UnmountInstanceSDGResponseBodyDataResultFailedItems `json:"FailedItems,omitempty" xml:"FailedItems,omitempty" type:"Repeated"`
	// The number of successful tasks.
	//
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s UnmountInstanceSDGResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s UnmountInstanceSDGResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *UnmountInstanceSDGResponseBodyDataResult) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *UnmountInstanceSDGResponseBodyDataResult) GetFailedItems() []*UnmountInstanceSDGResponseBodyDataResultFailedItems {
	return s.FailedItems
}

func (s *UnmountInstanceSDGResponseBodyDataResult) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *UnmountInstanceSDGResponseBodyDataResult) SetFailedCount(v int64) *UnmountInstanceSDGResponseBodyDataResult {
	s.FailedCount = &v
	return s
}

func (s *UnmountInstanceSDGResponseBodyDataResult) SetFailedItems(v []*UnmountInstanceSDGResponseBodyDataResultFailedItems) *UnmountInstanceSDGResponseBodyDataResult {
	s.FailedItems = v
	return s
}

func (s *UnmountInstanceSDGResponseBodyDataResult) SetSuccessCount(v int64) *UnmountInstanceSDGResponseBodyDataResult {
	s.SuccessCount = &v
	return s
}

func (s *UnmountInstanceSDGResponseBodyDataResult) Validate() error {
	if s.FailedItems != nil {
		for _, item := range s.FailedItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UnmountInstanceSDGResponseBodyDataResultFailedItems struct {
	// The error message.
	//
	// example:
	//
	// sdg not found
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// aic-xxxxx-0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UnmountInstanceSDGResponseBodyDataResultFailedItems) String() string {
	return dara.Prettify(s)
}

func (s UnmountInstanceSDGResponseBodyDataResultFailedItems) GoString() string {
	return s.String()
}

func (s *UnmountInstanceSDGResponseBodyDataResultFailedItems) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UnmountInstanceSDGResponseBodyDataResultFailedItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnmountInstanceSDGResponseBodyDataResultFailedItems) SetErrMessage(v string) *UnmountInstanceSDGResponseBodyDataResultFailedItems {
	s.ErrMessage = &v
	return s
}

func (s *UnmountInstanceSDGResponseBodyDataResultFailedItems) SetInstanceId(v string) *UnmountInstanceSDGResponseBodyDataResultFailedItems {
	s.InstanceId = &v
	return s
}

func (s *UnmountInstanceSDGResponseBodyDataResultFailedItems) Validate() error {
	return dara.Validate(s)
}
