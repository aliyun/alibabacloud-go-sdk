// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstanceSDGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *RemoveInstanceSDGResponseBody
	GetCode() *int32
	SetData(v *RemoveInstanceSDGResponseBodyData) *RemoveInstanceSDGResponseBody
	GetData() *RemoveInstanceSDGResponseBodyData
	SetRequestId(v string) *RemoveInstanceSDGResponseBody
	GetRequestId() *string
}

type RemoveInstanceSDGResponseBody struct {
	// The returned service code. 0 indicates that the request was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data object.
	Data *RemoveInstanceSDGResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 125B04C7-3D0D-4245-AF96-14E3758E3F06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveInstanceSDGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstanceSDGResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveInstanceSDGResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RemoveInstanceSDGResponseBody) GetData() *RemoveInstanceSDGResponseBodyData {
	return s.Data
}

func (s *RemoveInstanceSDGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveInstanceSDGResponseBody) SetCode(v int32) *RemoveInstanceSDGResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveInstanceSDGResponseBody) SetData(v *RemoveInstanceSDGResponseBodyData) *RemoveInstanceSDGResponseBody {
	s.Data = v
	return s
}

func (s *RemoveInstanceSDGResponseBody) SetRequestId(v string) *RemoveInstanceSDGResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveInstanceSDGResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveInstanceSDGResponseBodyData struct {
	// The response message. Success is returned for a successful request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution result of the synchronization request.
	Result *RemoveInstanceSDGResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s RemoveInstanceSDGResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstanceSDGResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveInstanceSDGResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *RemoveInstanceSDGResponseBodyData) GetResult() *RemoveInstanceSDGResponseBodyDataResult {
	return s.Result
}

func (s *RemoveInstanceSDGResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveInstanceSDGResponseBodyData) SetMessage(v string) *RemoveInstanceSDGResponseBodyData {
	s.Message = &v
	return s
}

func (s *RemoveInstanceSDGResponseBodyData) SetResult(v *RemoveInstanceSDGResponseBodyDataResult) *RemoveInstanceSDGResponseBodyData {
	s.Result = v
	return s
}

func (s *RemoveInstanceSDGResponseBodyData) SetSuccess(v bool) *RemoveInstanceSDGResponseBodyData {
	s.Success = &v
	return s
}

func (s *RemoveInstanceSDGResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveInstanceSDGResponseBodyDataResult struct {
	// The number of failed tasks.
	//
	// example:
	//
	// 0
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// Details about the failed tasks.
	FailedItems []*RemoveInstanceSDGResponseBodyDataResultFailedItems `json:"FailedItems,omitempty" xml:"FailedItems,omitempty" type:"Repeated"`
	// The number of successful tasks.
	//
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s RemoveInstanceSDGResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstanceSDGResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *RemoveInstanceSDGResponseBodyDataResult) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *RemoveInstanceSDGResponseBodyDataResult) GetFailedItems() []*RemoveInstanceSDGResponseBodyDataResultFailedItems {
	return s.FailedItems
}

func (s *RemoveInstanceSDGResponseBodyDataResult) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *RemoveInstanceSDGResponseBodyDataResult) SetFailedCount(v int64) *RemoveInstanceSDGResponseBodyDataResult {
	s.FailedCount = &v
	return s
}

func (s *RemoveInstanceSDGResponseBodyDataResult) SetFailedItems(v []*RemoveInstanceSDGResponseBodyDataResultFailedItems) *RemoveInstanceSDGResponseBodyDataResult {
	s.FailedItems = v
	return s
}

func (s *RemoveInstanceSDGResponseBodyDataResult) SetSuccessCount(v int64) *RemoveInstanceSDGResponseBodyDataResult {
	s.SuccessCount = &v
	return s
}

func (s *RemoveInstanceSDGResponseBodyDataResult) Validate() error {
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

type RemoveInstanceSDGResponseBodyDataResultFailedItems struct {
	// The error message that is returned.
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

func (s RemoveInstanceSDGResponseBodyDataResultFailedItems) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstanceSDGResponseBodyDataResultFailedItems) GoString() string {
	return s.String()
}

func (s *RemoveInstanceSDGResponseBodyDataResultFailedItems) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *RemoveInstanceSDGResponseBodyDataResultFailedItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveInstanceSDGResponseBodyDataResultFailedItems) SetErrMessage(v string) *RemoveInstanceSDGResponseBodyDataResultFailedItems {
	s.ErrMessage = &v
	return s
}

func (s *RemoveInstanceSDGResponseBodyDataResultFailedItems) SetInstanceId(v string) *RemoveInstanceSDGResponseBodyDataResultFailedItems {
	s.InstanceId = &v
	return s
}

func (s *RemoveInstanceSDGResponseBodyDataResultFailedItems) Validate() error {
	return dara.Validate(s)
}
