// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperateWorkItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchOperateWorkItemsResponseBody
	GetCode() *string
	SetItems(v []*BatchOperateWorkItemsResponseBodyItems) *BatchOperateWorkItemsResponseBody
	GetItems() []*BatchOperateWorkItemsResponseBodyItems
	SetMessage(v string) *BatchOperateWorkItemsResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchOperateWorkItemsResponseBody
	GetRequestId() *string
	SetSkipped(v int64) *BatchOperateWorkItemsResponseBody
	GetSkipped() *int64
	SetSucceeded(v int64) *BatchOperateWorkItemsResponseBody
	GetSucceeded() *int64
}

type BatchOperateWorkItemsResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The results for each item. The order of the results matches the order of the items in the request.
	//
	// example:
	//
	// example
	Items []*BatchOperateWorkItemsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The number of skipped or failed operations.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"skipped,omitempty" xml:"skipped,omitempty"`
	// The number of successful operations.
	//
	// example:
	//
	// 1
	Succeeded *int64 `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}

func (s BatchOperateWorkItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchOperateWorkItemsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchOperateWorkItemsResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchOperateWorkItemsResponseBody) GetItems() []*BatchOperateWorkItemsResponseBodyItems {
	return s.Items
}

func (s *BatchOperateWorkItemsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchOperateWorkItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchOperateWorkItemsResponseBody) GetSkipped() *int64 {
	return s.Skipped
}

func (s *BatchOperateWorkItemsResponseBody) GetSucceeded() *int64 {
	return s.Succeeded
}

func (s *BatchOperateWorkItemsResponseBody) SetCode(v string) *BatchOperateWorkItemsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchOperateWorkItemsResponseBody) SetItems(v []*BatchOperateWorkItemsResponseBodyItems) *BatchOperateWorkItemsResponseBody {
	s.Items = v
	return s
}

func (s *BatchOperateWorkItemsResponseBody) SetMessage(v string) *BatchOperateWorkItemsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchOperateWorkItemsResponseBody) SetRequestId(v string) *BatchOperateWorkItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchOperateWorkItemsResponseBody) SetSkipped(v int64) *BatchOperateWorkItemsResponseBody {
	s.Skipped = &v
	return s
}

func (s *BatchOperateWorkItemsResponseBody) SetSucceeded(v int64) *BatchOperateWorkItemsResponseBody {
	s.Succeeded = &v
	return s
}

func (s *BatchOperateWorkItemsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchOperateWorkItemsResponseBodyItems struct {
	// The reason for skipping or failure.
	//
	// example:
	//
	// example
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// Indicates whether the execution was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The ID of the work item.
	//
	// example:
	//
	// example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s BatchOperateWorkItemsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s BatchOperateWorkItemsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *BatchOperateWorkItemsResponseBodyItems) GetReason() *string {
	return s.Reason
}

func (s *BatchOperateWorkItemsResponseBodyItems) GetSuccess() *bool {
	return s.Success
}

func (s *BatchOperateWorkItemsResponseBodyItems) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *BatchOperateWorkItemsResponseBodyItems) SetReason(v string) *BatchOperateWorkItemsResponseBodyItems {
	s.Reason = &v
	return s
}

func (s *BatchOperateWorkItemsResponseBodyItems) SetSuccess(v bool) *BatchOperateWorkItemsResponseBodyItems {
	s.Success = &v
	return s
}

func (s *BatchOperateWorkItemsResponseBodyItems) SetWorkItemId(v string) *BatchOperateWorkItemsResponseBodyItems {
	s.WorkItemId = &v
	return s
}

func (s *BatchOperateWorkItemsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
