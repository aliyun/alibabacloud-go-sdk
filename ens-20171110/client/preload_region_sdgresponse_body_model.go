// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadRegionSDGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *PreloadRegionSDGResponseBodyData) *PreloadRegionSDGResponseBody
	GetData() *PreloadRegionSDGResponseBodyData
	SetRequestId(v string) *PreloadRegionSDGResponseBody
	GetRequestId() *string
}

type PreloadRegionSDGResponseBody struct {
	// The returned data object.
	Data *PreloadRegionSDGResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PreloadRegionSDGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreloadRegionSDGResponseBody) GoString() string {
	return s.String()
}

func (s *PreloadRegionSDGResponseBody) GetData() *PreloadRegionSDGResponseBodyData {
	return s.Data
}

func (s *PreloadRegionSDGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreloadRegionSDGResponseBody) SetData(v *PreloadRegionSDGResponseBodyData) *PreloadRegionSDGResponseBody {
	s.Data = v
	return s
}

func (s *PreloadRegionSDGResponseBody) SetRequestId(v string) *PreloadRegionSDGResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreloadRegionSDGResponseBody) Validate() error {
	return dara.Validate(s)
}

type PreloadRegionSDGResponseBodyData struct {
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution result of the synchronization request.
	Result *PreloadRegionSDGResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s PreloadRegionSDGResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PreloadRegionSDGResponseBodyData) GoString() string {
	return s.String()
}

func (s *PreloadRegionSDGResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *PreloadRegionSDGResponseBodyData) GetResult() *PreloadRegionSDGResponseBodyDataResult {
	return s.Result
}

func (s *PreloadRegionSDGResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *PreloadRegionSDGResponseBodyData) SetMessage(v string) *PreloadRegionSDGResponseBodyData {
	s.Message = &v
	return s
}

func (s *PreloadRegionSDGResponseBodyData) SetResult(v *PreloadRegionSDGResponseBodyDataResult) *PreloadRegionSDGResponseBodyData {
	s.Result = v
	return s
}

func (s *PreloadRegionSDGResponseBodyData) SetSuccess(v bool) *PreloadRegionSDGResponseBodyData {
	s.Success = &v
	return s
}

func (s *PreloadRegionSDGResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type PreloadRegionSDGResponseBodyDataResult struct {
	// The number of failed tasks.
	//
	// example:
	//
	// 0
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// Details about the failed tasks.
	FailedItems []*PreloadRegionSDGResponseBodyDataResultFailedItems `json:"FailedItems,omitempty" xml:"FailedItems,omitempty" type:"Repeated"`
	// The number of successful tasks.
	//
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s PreloadRegionSDGResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s PreloadRegionSDGResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *PreloadRegionSDGResponseBodyDataResult) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *PreloadRegionSDGResponseBodyDataResult) GetFailedItems() []*PreloadRegionSDGResponseBodyDataResultFailedItems {
	return s.FailedItems
}

func (s *PreloadRegionSDGResponseBodyDataResult) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *PreloadRegionSDGResponseBodyDataResult) SetFailedCount(v int64) *PreloadRegionSDGResponseBodyDataResult {
	s.FailedCount = &v
	return s
}

func (s *PreloadRegionSDGResponseBodyDataResult) SetFailedItems(v []*PreloadRegionSDGResponseBodyDataResultFailedItems) *PreloadRegionSDGResponseBodyDataResult {
	s.FailedItems = v
	return s
}

func (s *PreloadRegionSDGResponseBodyDataResult) SetSuccessCount(v int64) *PreloadRegionSDGResponseBodyDataResult {
	s.SuccessCount = &v
	return s
}

func (s *PreloadRegionSDGResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}

type PreloadRegionSDGResponseBodyDataResultFailedItems struct {
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

func (s PreloadRegionSDGResponseBodyDataResultFailedItems) String() string {
	return dara.Prettify(s)
}

func (s PreloadRegionSDGResponseBodyDataResultFailedItems) GoString() string {
	return s.String()
}

func (s *PreloadRegionSDGResponseBodyDataResultFailedItems) GetDestinationRegionId() *string {
	return s.DestinationRegionId
}

func (s *PreloadRegionSDGResponseBodyDataResultFailedItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *PreloadRegionSDGResponseBodyDataResultFailedItems) SetDestinationRegionId(v string) *PreloadRegionSDGResponseBodyDataResultFailedItems {
	s.DestinationRegionId = &v
	return s
}

func (s *PreloadRegionSDGResponseBodyDataResultFailedItems) SetErrorMessage(v string) *PreloadRegionSDGResponseBodyDataResultFailedItems {
	s.ErrorMessage = &v
	return s
}

func (s *PreloadRegionSDGResponseBodyDataResultFailedItems) Validate() error {
	return dara.Validate(s)
}
