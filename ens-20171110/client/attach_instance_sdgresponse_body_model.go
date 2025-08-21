// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstanceSDGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AttachInstanceSDGResponseBodyData) *AttachInstanceSDGResponseBody
	GetData() *AttachInstanceSDGResponseBodyData
	SetRequestId(v string) *AttachInstanceSDGResponseBody
	GetRequestId() *string
}

type AttachInstanceSDGResponseBody struct {
	// The returned data object.
	Data *AttachInstanceSDGResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachInstanceSDGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachInstanceSDGResponseBody) GoString() string {
	return s.String()
}

func (s *AttachInstanceSDGResponseBody) GetData() *AttachInstanceSDGResponseBodyData {
	return s.Data
}

func (s *AttachInstanceSDGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachInstanceSDGResponseBody) SetData(v *AttachInstanceSDGResponseBodyData) *AttachInstanceSDGResponseBody {
	s.Data = v
	return s
}

func (s *AttachInstanceSDGResponseBody) SetRequestId(v string) *AttachInstanceSDGResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachInstanceSDGResponseBody) Validate() error {
	return dara.Validate(s)
}

type AttachInstanceSDGResponseBodyData struct {
	// The response message. Success is returned for a successful request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution result of the synchronization request.
	Result *AttachInstanceSDGResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s AttachInstanceSDGResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AttachInstanceSDGResponseBodyData) GoString() string {
	return s.String()
}

func (s *AttachInstanceSDGResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *AttachInstanceSDGResponseBodyData) GetResult() *AttachInstanceSDGResponseBodyDataResult {
	return s.Result
}

func (s *AttachInstanceSDGResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *AttachInstanceSDGResponseBodyData) SetMessage(v string) *AttachInstanceSDGResponseBodyData {
	s.Message = &v
	return s
}

func (s *AttachInstanceSDGResponseBodyData) SetResult(v *AttachInstanceSDGResponseBodyDataResult) *AttachInstanceSDGResponseBodyData {
	s.Result = v
	return s
}

func (s *AttachInstanceSDGResponseBodyData) SetSuccess(v bool) *AttachInstanceSDGResponseBodyData {
	s.Success = &v
	return s
}

func (s *AttachInstanceSDGResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type AttachInstanceSDGResponseBodyDataResult struct {
	// The number of failed tasks.
	//
	// example:
	//
	// 0
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// Details about failed tasks.
	FailedItems []*AttachInstanceSDGResponseBodyDataResultFailedItems `json:"FailedItems,omitempty" xml:"FailedItems,omitempty" type:"Repeated"`
	// The number of successful tasks.
	//
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s AttachInstanceSDGResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s AttachInstanceSDGResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *AttachInstanceSDGResponseBodyDataResult) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *AttachInstanceSDGResponseBodyDataResult) GetFailedItems() []*AttachInstanceSDGResponseBodyDataResultFailedItems {
	return s.FailedItems
}

func (s *AttachInstanceSDGResponseBodyDataResult) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *AttachInstanceSDGResponseBodyDataResult) SetFailedCount(v int64) *AttachInstanceSDGResponseBodyDataResult {
	s.FailedCount = &v
	return s
}

func (s *AttachInstanceSDGResponseBodyDataResult) SetFailedItems(v []*AttachInstanceSDGResponseBodyDataResultFailedItems) *AttachInstanceSDGResponseBodyDataResult {
	s.FailedItems = v
	return s
}

func (s *AttachInstanceSDGResponseBodyDataResult) SetSuccessCount(v int64) *AttachInstanceSDGResponseBodyDataResult {
	s.SuccessCount = &v
	return s
}

func (s *AttachInstanceSDGResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}

type AttachInstanceSDGResponseBodyDataResultFailedItems struct {
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

func (s AttachInstanceSDGResponseBodyDataResultFailedItems) String() string {
	return dara.Prettify(s)
}

func (s AttachInstanceSDGResponseBodyDataResultFailedItems) GoString() string {
	return s.String()
}

func (s *AttachInstanceSDGResponseBodyDataResultFailedItems) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *AttachInstanceSDGResponseBodyDataResultFailedItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachInstanceSDGResponseBodyDataResultFailedItems) SetErrMessage(v string) *AttachInstanceSDGResponseBodyDataResultFailedItems {
	s.ErrMessage = &v
	return s
}

func (s *AttachInstanceSDGResponseBodyDataResultFailedItems) SetInstanceId(v string) *AttachInstanceSDGResponseBodyDataResultFailedItems {
	s.InstanceId = &v
	return s
}

func (s *AttachInstanceSDGResponseBodyDataResultFailedItems) Validate() error {
	return dara.Validate(s)
}
