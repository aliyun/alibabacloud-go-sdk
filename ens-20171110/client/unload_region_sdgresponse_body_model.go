// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnloadRegionSDGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UnloadRegionSDGResponseBodyData) *UnloadRegionSDGResponseBody
	GetData() *UnloadRegionSDGResponseBodyData
	SetRequestId(v string) *UnloadRegionSDGResponseBody
	GetRequestId() *string
}

type UnloadRegionSDGResponseBody struct {
	// The returned data object.
	Data *UnloadRegionSDGResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 125B04C7-3D0D-4245-AF96-14E3758E3F06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnloadRegionSDGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnloadRegionSDGResponseBody) GoString() string {
	return s.String()
}

func (s *UnloadRegionSDGResponseBody) GetData() *UnloadRegionSDGResponseBodyData {
	return s.Data
}

func (s *UnloadRegionSDGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnloadRegionSDGResponseBody) SetData(v *UnloadRegionSDGResponseBodyData) *UnloadRegionSDGResponseBody {
	s.Data = v
	return s
}

func (s *UnloadRegionSDGResponseBody) SetRequestId(v string) *UnloadRegionSDGResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnloadRegionSDGResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UnloadRegionSDGResponseBodyData struct {
	// The response message. Success is returned for a successful request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution result of the synchronization request.
	Result *UnloadRegionSDGResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s UnloadRegionSDGResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UnloadRegionSDGResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnloadRegionSDGResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *UnloadRegionSDGResponseBodyData) GetResult() *UnloadRegionSDGResponseBodyDataResult {
	return s.Result
}

func (s *UnloadRegionSDGResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UnloadRegionSDGResponseBodyData) SetMessage(v string) *UnloadRegionSDGResponseBodyData {
	s.Message = &v
	return s
}

func (s *UnloadRegionSDGResponseBodyData) SetResult(v *UnloadRegionSDGResponseBodyDataResult) *UnloadRegionSDGResponseBodyData {
	s.Result = v
	return s
}

func (s *UnloadRegionSDGResponseBodyData) SetSuccess(v bool) *UnloadRegionSDGResponseBodyData {
	s.Success = &v
	return s
}

func (s *UnloadRegionSDGResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UnloadRegionSDGResponseBodyDataResult struct {
	// The number of failed tasks.
	//
	// example:
	//
	// 0
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// Details about the failed tasks.
	FailedItems []*UnloadRegionSDGResponseBodyDataResultFailedItems `json:"FailedItems,omitempty" xml:"FailedItems,omitempty" type:"Repeated"`
	// The number of successful tasks.
	//
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s UnloadRegionSDGResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s UnloadRegionSDGResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *UnloadRegionSDGResponseBodyDataResult) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *UnloadRegionSDGResponseBodyDataResult) GetFailedItems() []*UnloadRegionSDGResponseBodyDataResultFailedItems {
	return s.FailedItems
}

func (s *UnloadRegionSDGResponseBodyDataResult) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *UnloadRegionSDGResponseBodyDataResult) SetFailedCount(v int64) *UnloadRegionSDGResponseBodyDataResult {
	s.FailedCount = &v
	return s
}

func (s *UnloadRegionSDGResponseBodyDataResult) SetFailedItems(v []*UnloadRegionSDGResponseBodyDataResultFailedItems) *UnloadRegionSDGResponseBodyDataResult {
	s.FailedItems = v
	return s
}

func (s *UnloadRegionSDGResponseBodyDataResult) SetSuccessCount(v int64) *UnloadRegionSDGResponseBodyDataResult {
	s.SuccessCount = &v
	return s
}

func (s *UnloadRegionSDGResponseBodyDataResult) Validate() error {
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

type UnloadRegionSDGResponseBodyDataResultFailedItems struct {
	// The ID of the destination node.
	//
	// example:
	//
	// cn-hangzhou-xxx
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// region not found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s UnloadRegionSDGResponseBodyDataResultFailedItems) String() string {
	return dara.Prettify(s)
}

func (s UnloadRegionSDGResponseBodyDataResultFailedItems) GoString() string {
	return s.String()
}

func (s *UnloadRegionSDGResponseBodyDataResultFailedItems) GetDestinationRegionId() *string {
	return s.DestinationRegionId
}

func (s *UnloadRegionSDGResponseBodyDataResultFailedItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UnloadRegionSDGResponseBodyDataResultFailedItems) SetDestinationRegionId(v string) *UnloadRegionSDGResponseBodyDataResultFailedItems {
	s.DestinationRegionId = &v
	return s
}

func (s *UnloadRegionSDGResponseBodyDataResultFailedItems) SetErrorMessage(v string) *UnloadRegionSDGResponseBodyDataResultFailedItems {
	s.ErrorMessage = &v
	return s
}

func (s *UnloadRegionSDGResponseBodyDataResultFailedItems) Validate() error {
	return dara.Validate(s)
}
