// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopySDGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CopySDGResponseBodyData) *CopySDGResponseBody
	GetData() *CopySDGResponseBodyData
	SetRequestId(v string) *CopySDGResponseBody
	GetRequestId() *string
}

type CopySDGResponseBody struct {
	// The returned data object.
	Data *CopySDGResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// XXX-XXX-XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopySDGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopySDGResponseBody) GoString() string {
	return s.String()
}

func (s *CopySDGResponseBody) GetData() *CopySDGResponseBodyData {
	return s.Data
}

func (s *CopySDGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopySDGResponseBody) SetData(v *CopySDGResponseBodyData) *CopySDGResponseBody {
	s.Data = v
	return s
}

func (s *CopySDGResponseBody) SetRequestId(v string) *CopySDGResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopySDGResponseBody) Validate() error {
	return dara.Validate(s)
}

type CopySDGResponseBodyData struct {
	// The response message. Success is returned for a successful request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution result of the synchronization request.
	Result *CopySDGResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s CopySDGResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CopySDGResponseBodyData) GoString() string {
	return s.String()
}

func (s *CopySDGResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *CopySDGResponseBodyData) GetResult() *CopySDGResponseBodyDataResult {
	return s.Result
}

func (s *CopySDGResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *CopySDGResponseBodyData) SetMessage(v string) *CopySDGResponseBodyData {
	s.Message = &v
	return s
}

func (s *CopySDGResponseBodyData) SetResult(v *CopySDGResponseBodyDataResult) *CopySDGResponseBodyData {
	s.Result = v
	return s
}

func (s *CopySDGResponseBodyData) SetSuccess(v bool) *CopySDGResponseBodyData {
	s.Success = &v
	return s
}

func (s *CopySDGResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type CopySDGResponseBodyDataResult struct {
	// The number of failed nodes.
	//
	// example:
	//
	// 0
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// Details of failed nodes.
	FailedItems []*CopySDGResponseBodyDataResultFailedItems `json:"FailedItems,omitempty" xml:"FailedItems,omitempty" type:"Repeated"`
	// The number of successful nodes.
	//
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s CopySDGResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s CopySDGResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *CopySDGResponseBodyDataResult) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *CopySDGResponseBodyDataResult) GetFailedItems() []*CopySDGResponseBodyDataResultFailedItems {
	return s.FailedItems
}

func (s *CopySDGResponseBodyDataResult) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *CopySDGResponseBodyDataResult) SetFailedCount(v int64) *CopySDGResponseBodyDataResult {
	s.FailedCount = &v
	return s
}

func (s *CopySDGResponseBodyDataResult) SetFailedItems(v []*CopySDGResponseBodyDataResultFailedItems) *CopySDGResponseBodyDataResult {
	s.FailedItems = v
	return s
}

func (s *CopySDGResponseBodyDataResult) SetSuccessCount(v int64) *CopySDGResponseBodyDataResult {
	s.SuccessCount = &v
	return s
}

func (s *CopySDGResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}

type CopySDGResponseBodyDataResultFailedItems struct {
	// The ID of the destination node.
	//
	// example:
	//
	// cn-hangzhou-xxx
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// The error message.
	//
	// example:
	//
	// region not found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s CopySDGResponseBodyDataResultFailedItems) String() string {
	return dara.Prettify(s)
}

func (s CopySDGResponseBodyDataResultFailedItems) GoString() string {
	return s.String()
}

func (s *CopySDGResponseBodyDataResultFailedItems) GetDestinationRegionId() *string {
	return s.DestinationRegionId
}

func (s *CopySDGResponseBodyDataResultFailedItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CopySDGResponseBodyDataResultFailedItems) SetDestinationRegionId(v string) *CopySDGResponseBodyDataResultFailedItems {
	s.DestinationRegionId = &v
	return s
}

func (s *CopySDGResponseBodyDataResultFailedItems) SetErrorMessage(v string) *CopySDGResponseBodyDataResultFailedItems {
	s.ErrorMessage = &v
	return s
}

func (s *CopySDGResponseBodyDataResultFailedItems) Validate() error {
	return dara.Validate(s)
}
