// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSDGsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RemoveSDGsResponseBodyData) *RemoveSDGsResponseBody
	GetData() *RemoveSDGsResponseBodyData
	SetRequestId(v string) *RemoveSDGsResponseBody
	GetRequestId() *string
}

type RemoveSDGsResponseBody struct {
	// The returned data.
	Data *RemoveSDGsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 903819D9-D18B-5C59-8DFD-20D56FD2DAE4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveSDGsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveSDGsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSDGsResponseBody) GetData() *RemoveSDGsResponseBodyData {
	return s.Data
}

func (s *RemoveSDGsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveSDGsResponseBody) SetData(v *RemoveSDGsResponseBodyData) *RemoveSDGsResponseBody {
	s.Data = v
	return s
}

func (s *RemoveSDGsResponseBody) SetRequestId(v string) *RemoveSDGsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSDGsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveSDGsResponseBodyData struct {
	// The response message. Success is returned for a successful request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution result of the synchronization request.
	Result *RemoveSDGsResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether all tasks are successful. Valid values:
	//
	// 	- **true**: All tasks are successful.
	//
	// 	- **false**: Failed tasks exist.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveSDGsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RemoveSDGsResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveSDGsResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *RemoveSDGsResponseBodyData) GetResult() *RemoveSDGsResponseBodyDataResult {
	return s.Result
}

func (s *RemoveSDGsResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveSDGsResponseBodyData) SetMessage(v string) *RemoveSDGsResponseBodyData {
	s.Message = &v
	return s
}

func (s *RemoveSDGsResponseBodyData) SetResult(v *RemoveSDGsResponseBodyDataResult) *RemoveSDGsResponseBodyData {
	s.Result = v
	return s
}

func (s *RemoveSDGsResponseBodyData) SetSuccess(v bool) *RemoveSDGsResponseBodyData {
	s.Success = &v
	return s
}

func (s *RemoveSDGsResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveSDGsResponseBodyDataResult struct {
	// The number of failed tasks.
	//
	// example:
	//
	// 0
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// Details about failed tasks.
	FailedItems []*RemoveSDGsResponseBodyDataResultFailedItems `json:"FailedItems,omitempty" xml:"FailedItems,omitempty" type:"Repeated"`
	// The number of successful tasks.
	//
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s RemoveSDGsResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s RemoveSDGsResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *RemoveSDGsResponseBodyDataResult) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *RemoveSDGsResponseBodyDataResult) GetFailedItems() []*RemoveSDGsResponseBodyDataResultFailedItems {
	return s.FailedItems
}

func (s *RemoveSDGsResponseBodyDataResult) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *RemoveSDGsResponseBodyDataResult) SetFailedCount(v int64) *RemoveSDGsResponseBodyDataResult {
	s.FailedCount = &v
	return s
}

func (s *RemoveSDGsResponseBodyDataResult) SetFailedItems(v []*RemoveSDGsResponseBodyDataResultFailedItems) *RemoveSDGsResponseBodyDataResult {
	s.FailedItems = v
	return s
}

func (s *RemoveSDGsResponseBodyDataResult) SetSuccessCount(v int64) *RemoveSDGsResponseBodyDataResult {
	s.SuccessCount = &v
	return s
}

func (s *RemoveSDGsResponseBodyDataResult) Validate() error {
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

type RemoveSDGsResponseBodyDataResultFailedItems struct {
	// The error message returned if the call failed.
	//
	// example:
	//
	// sdg not found
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// aic-xxxxx-0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RemoveSDGsResponseBodyDataResultFailedItems) String() string {
	return dara.Prettify(s)
}

func (s RemoveSDGsResponseBodyDataResultFailedItems) GoString() string {
	return s.String()
}

func (s *RemoveSDGsResponseBodyDataResultFailedItems) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *RemoveSDGsResponseBodyDataResultFailedItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveSDGsResponseBodyDataResultFailedItems) SetErrMessage(v string) *RemoveSDGsResponseBodyDataResultFailedItems {
	s.ErrMessage = &v
	return s
}

func (s *RemoveSDGsResponseBodyDataResultFailedItems) SetInstanceId(v string) *RemoveSDGsResponseBodyDataResultFailedItems {
	s.InstanceId = &v
	return s
}

func (s *RemoveSDGsResponseBodyDataResultFailedItems) Validate() error {
	return dara.Validate(s)
}
