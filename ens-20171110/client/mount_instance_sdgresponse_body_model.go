// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMountInstanceSDGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *MountInstanceSDGResponseBodyData) *MountInstanceSDGResponseBody
	GetData() *MountInstanceSDGResponseBodyData
	SetRequestId(v string) *MountInstanceSDGResponseBody
	GetRequestId() *string
}

type MountInstanceSDGResponseBody struct {
	// The returned data object.
	Data *MountInstanceSDGResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F3B261DD-3858-4D3C-877D-303ADF374600
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MountInstanceSDGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MountInstanceSDGResponseBody) GoString() string {
	return s.String()
}

func (s *MountInstanceSDGResponseBody) GetData() *MountInstanceSDGResponseBodyData {
	return s.Data
}

func (s *MountInstanceSDGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MountInstanceSDGResponseBody) SetData(v *MountInstanceSDGResponseBodyData) *MountInstanceSDGResponseBody {
	s.Data = v
	return s
}

func (s *MountInstanceSDGResponseBody) SetRequestId(v string) *MountInstanceSDGResponseBody {
	s.RequestId = &v
	return s
}

func (s *MountInstanceSDGResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MountInstanceSDGResponseBodyData struct {
	// The response message. Success is returned for a successful request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution result of the synchronization request.
	Result *MountInstanceSDGResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s MountInstanceSDGResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MountInstanceSDGResponseBodyData) GoString() string {
	return s.String()
}

func (s *MountInstanceSDGResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *MountInstanceSDGResponseBodyData) GetResult() *MountInstanceSDGResponseBodyDataResult {
	return s.Result
}

func (s *MountInstanceSDGResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *MountInstanceSDGResponseBodyData) SetMessage(v string) *MountInstanceSDGResponseBodyData {
	s.Message = &v
	return s
}

func (s *MountInstanceSDGResponseBodyData) SetResult(v *MountInstanceSDGResponseBodyDataResult) *MountInstanceSDGResponseBodyData {
	s.Result = v
	return s
}

func (s *MountInstanceSDGResponseBodyData) SetSuccess(v bool) *MountInstanceSDGResponseBodyData {
	s.Success = &v
	return s
}

func (s *MountInstanceSDGResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MountInstanceSDGResponseBodyDataResult struct {
	// The number of failed tasks.
	//
	// example:
	//
	// 0
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// Details about failed tasks.
	FailedItems []*MountInstanceSDGResponseBodyDataResultFailedItems `json:"FailedItems,omitempty" xml:"FailedItems,omitempty" type:"Repeated"`
	// The number of successful tasks.
	//
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s MountInstanceSDGResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s MountInstanceSDGResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *MountInstanceSDGResponseBodyDataResult) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *MountInstanceSDGResponseBodyDataResult) GetFailedItems() []*MountInstanceSDGResponseBodyDataResultFailedItems {
	return s.FailedItems
}

func (s *MountInstanceSDGResponseBodyDataResult) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *MountInstanceSDGResponseBodyDataResult) SetFailedCount(v int64) *MountInstanceSDGResponseBodyDataResult {
	s.FailedCount = &v
	return s
}

func (s *MountInstanceSDGResponseBodyDataResult) SetFailedItems(v []*MountInstanceSDGResponseBodyDataResultFailedItems) *MountInstanceSDGResponseBodyDataResult {
	s.FailedItems = v
	return s
}

func (s *MountInstanceSDGResponseBodyDataResult) SetSuccessCount(v int64) *MountInstanceSDGResponseBodyDataResult {
	s.SuccessCount = &v
	return s
}

func (s *MountInstanceSDGResponseBodyDataResult) Validate() error {
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

type MountInstanceSDGResponseBodyDataResultFailedItems struct {
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

func (s MountInstanceSDGResponseBodyDataResultFailedItems) String() string {
	return dara.Prettify(s)
}

func (s MountInstanceSDGResponseBodyDataResultFailedItems) GoString() string {
	return s.String()
}

func (s *MountInstanceSDGResponseBodyDataResultFailedItems) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *MountInstanceSDGResponseBodyDataResultFailedItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MountInstanceSDGResponseBodyDataResultFailedItems) SetErrMessage(v string) *MountInstanceSDGResponseBodyDataResultFailedItems {
	s.ErrMessage = &v
	return s
}

func (s *MountInstanceSDGResponseBodyDataResultFailedItems) SetInstanceId(v string) *MountInstanceSDGResponseBodyDataResultFailedItems {
	s.InstanceId = &v
	return s
}

func (s *MountInstanceSDGResponseBodyDataResultFailedItems) Validate() error {
	return dara.Validate(s)
}
