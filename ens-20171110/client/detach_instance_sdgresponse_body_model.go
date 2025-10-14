// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachInstanceSDGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetachInstanceSDGResponseBodyData) *DetachInstanceSDGResponseBody
	GetData() *DetachInstanceSDGResponseBodyData
	SetRequestId(v string) *DetachInstanceSDGResponseBody
	GetRequestId() *string
}

type DetachInstanceSDGResponseBody struct {
	// The returned data object.
	Data *DetachInstanceSDGResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachInstanceSDGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachInstanceSDGResponseBody) GoString() string {
	return s.String()
}

func (s *DetachInstanceSDGResponseBody) GetData() *DetachInstanceSDGResponseBodyData {
	return s.Data
}

func (s *DetachInstanceSDGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachInstanceSDGResponseBody) SetData(v *DetachInstanceSDGResponseBodyData) *DetachInstanceSDGResponseBody {
	s.Data = v
	return s
}

func (s *DetachInstanceSDGResponseBody) SetRequestId(v string) *DetachInstanceSDGResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachInstanceSDGResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetachInstanceSDGResponseBodyData struct {
	// The response message. Success is returned for a successful request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution result of the synchronization request.
	Result *DetachInstanceSDGResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s DetachInstanceSDGResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetachInstanceSDGResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetachInstanceSDGResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DetachInstanceSDGResponseBodyData) GetResult() *DetachInstanceSDGResponseBodyDataResult {
	return s.Result
}

func (s *DetachInstanceSDGResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DetachInstanceSDGResponseBodyData) SetMessage(v string) *DetachInstanceSDGResponseBodyData {
	s.Message = &v
	return s
}

func (s *DetachInstanceSDGResponseBodyData) SetResult(v *DetachInstanceSDGResponseBodyDataResult) *DetachInstanceSDGResponseBodyData {
	s.Result = v
	return s
}

func (s *DetachInstanceSDGResponseBodyData) SetSuccess(v bool) *DetachInstanceSDGResponseBodyData {
	s.Success = &v
	return s
}

func (s *DetachInstanceSDGResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetachInstanceSDGResponseBodyDataResult struct {
	// The number of failed tasks.
	//
	// example:
	//
	// 0
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// Details about failed tasks.
	FailedItems []*DetachInstanceSDGResponseBodyDataResultFailedItems `json:"FailedItems,omitempty" xml:"FailedItems,omitempty" type:"Repeated"`
	// The number of successful tasks.
	//
	// example:
	//
	// 1
	SuccessCount *string `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s DetachInstanceSDGResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s DetachInstanceSDGResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DetachInstanceSDGResponseBodyDataResult) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *DetachInstanceSDGResponseBodyDataResult) GetFailedItems() []*DetachInstanceSDGResponseBodyDataResultFailedItems {
	return s.FailedItems
}

func (s *DetachInstanceSDGResponseBodyDataResult) GetSuccessCount() *string {
	return s.SuccessCount
}

func (s *DetachInstanceSDGResponseBodyDataResult) SetFailedCount(v int64) *DetachInstanceSDGResponseBodyDataResult {
	s.FailedCount = &v
	return s
}

func (s *DetachInstanceSDGResponseBodyDataResult) SetFailedItems(v []*DetachInstanceSDGResponseBodyDataResultFailedItems) *DetachInstanceSDGResponseBodyDataResult {
	s.FailedItems = v
	return s
}

func (s *DetachInstanceSDGResponseBodyDataResult) SetSuccessCount(v string) *DetachInstanceSDGResponseBodyDataResult {
	s.SuccessCount = &v
	return s
}

func (s *DetachInstanceSDGResponseBodyDataResult) Validate() error {
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

type DetachInstanceSDGResponseBodyDataResultFailedItems struct {
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

func (s DetachInstanceSDGResponseBodyDataResultFailedItems) String() string {
	return dara.Prettify(s)
}

func (s DetachInstanceSDGResponseBodyDataResultFailedItems) GoString() string {
	return s.String()
}

func (s *DetachInstanceSDGResponseBodyDataResultFailedItems) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DetachInstanceSDGResponseBodyDataResultFailedItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachInstanceSDGResponseBodyDataResultFailedItems) SetErrMessage(v string) *DetachInstanceSDGResponseBodyDataResultFailedItems {
	s.ErrMessage = &v
	return s
}

func (s *DetachInstanceSDGResponseBodyDataResultFailedItems) SetInstanceId(v string) *DetachInstanceSDGResponseBodyDataResultFailedItems {
	s.InstanceId = &v
	return s
}

func (s *DetachInstanceSDGResponseBodyDataResultFailedItems) Validate() error {
	return dara.Validate(s)
}
