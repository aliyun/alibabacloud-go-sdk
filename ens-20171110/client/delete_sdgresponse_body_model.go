// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSDGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteSDGResponseBodyData) *DeleteSDGResponseBody
	GetData() *DeleteSDGResponseBodyData
	SetRequestId(v string) *DeleteSDGResponseBody
	GetRequestId() *string
}

type DeleteSDGResponseBody struct {
	// The returned data object.
	Data *DeleteSDGResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 86A6D421-A0C7-4C01-8648-47377CA6A2CE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSDGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSDGResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSDGResponseBody) GetData() *DeleteSDGResponseBodyData {
	return s.Data
}

func (s *DeleteSDGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSDGResponseBody) SetData(v *DeleteSDGResponseBodyData) *DeleteSDGResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSDGResponseBody) SetRequestId(v string) *DeleteSDGResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSDGResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteSDGResponseBodyData struct {
	// The response message. Success is returned for a successful request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution result of the synchronization request.
	Result *DeleteSDGResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s DeleteSDGResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteSDGResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteSDGResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DeleteSDGResponseBodyData) GetResult() *DeleteSDGResponseBodyDataResult {
	return s.Result
}

func (s *DeleteSDGResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSDGResponseBodyData) SetMessage(v string) *DeleteSDGResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeleteSDGResponseBodyData) SetResult(v *DeleteSDGResponseBodyDataResult) *DeleteSDGResponseBodyData {
	s.Result = v
	return s
}

func (s *DeleteSDGResponseBodyData) SetSuccess(v bool) *DeleteSDGResponseBodyData {
	s.Success = &v
	return s
}

func (s *DeleteSDGResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DeleteSDGResponseBodyDataResult struct {
	// The number of failed tasks.
	//
	// example:
	//
	// 0
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// Details about the failed tasks.
	FailedItems []*DeleteSDGResponseBodyDataResultFailedItems `json:"FailedItems,omitempty" xml:"FailedItems,omitempty" type:"Repeated"`
	// The number of successful tasks.
	//
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s DeleteSDGResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteSDGResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DeleteSDGResponseBodyDataResult) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *DeleteSDGResponseBodyDataResult) GetFailedItems() []*DeleteSDGResponseBodyDataResultFailedItems {
	return s.FailedItems
}

func (s *DeleteSDGResponseBodyDataResult) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *DeleteSDGResponseBodyDataResult) SetFailedCount(v int64) *DeleteSDGResponseBodyDataResult {
	s.FailedCount = &v
	return s
}

func (s *DeleteSDGResponseBodyDataResult) SetFailedItems(v []*DeleteSDGResponseBodyDataResultFailedItems) *DeleteSDGResponseBodyDataResult {
	s.FailedItems = v
	return s
}

func (s *DeleteSDGResponseBodyDataResult) SetSuccessCount(v int64) *DeleteSDGResponseBodyDataResult {
	s.SuccessCount = &v
	return s
}

func (s *DeleteSDGResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}

type DeleteSDGResponseBodyDataResultFailedItems struct {
	// The error message.
	//
	// example:
	//
	// sdg not found
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// Description
	Item *DeleteSDGResponseBodyDataResultFailedItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Struct"`
}

func (s DeleteSDGResponseBodyDataResultFailedItems) String() string {
	return dara.Prettify(s)
}

func (s DeleteSDGResponseBodyDataResultFailedItems) GoString() string {
	return s.String()
}

func (s *DeleteSDGResponseBodyDataResultFailedItems) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteSDGResponseBodyDataResultFailedItems) GetItem() *DeleteSDGResponseBodyDataResultFailedItemsItem {
	return s.Item
}

func (s *DeleteSDGResponseBodyDataResultFailedItems) SetErrMessage(v string) *DeleteSDGResponseBodyDataResultFailedItems {
	s.ErrMessage = &v
	return s
}

func (s *DeleteSDGResponseBodyDataResultFailedItems) SetItem(v *DeleteSDGResponseBodyDataResultFailedItemsItem) *DeleteSDGResponseBodyDataResultFailedItems {
	s.Item = v
	return s
}

func (s *DeleteSDGResponseBodyDataResultFailedItems) Validate() error {
	return dara.Validate(s)
}

type DeleteSDGResponseBodyDataResultFailedItemsItem struct {
	// The ID of the shared data group (SDG).
	//
	// example:
	//
	// sdg-dfet5vvvgy
	SdgId *string `json:"SdgId,omitempty" xml:"SdgId,omitempty"`
}

func (s DeleteSDGResponseBodyDataResultFailedItemsItem) String() string {
	return dara.Prettify(s)
}

func (s DeleteSDGResponseBodyDataResultFailedItemsItem) GoString() string {
	return s.String()
}

func (s *DeleteSDGResponseBodyDataResultFailedItemsItem) GetSdgId() *string {
	return s.SdgId
}

func (s *DeleteSDGResponseBodyDataResultFailedItemsItem) SetSdgId(v string) *DeleteSDGResponseBodyDataResultFailedItemsItem {
	s.SdgId = &v
	return s
}

func (s *DeleteSDGResponseBodyDataResultFailedItemsItem) Validate() error {
	return dara.Validate(s)
}
