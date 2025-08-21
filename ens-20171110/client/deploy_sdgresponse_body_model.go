// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeploySDGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeploySDGResponseBodyData) *DeploySDGResponseBody
	GetData() *DeploySDGResponseBodyData
	SetRequestId(v string) *DeploySDGResponseBody
	GetRequestId() *string
}

type DeploySDGResponseBody struct {
	// The returned data object.
	Data *DeploySDGResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// A331CA96-3948-4BD2-B067-F6174F5C17EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeploySDGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeploySDGResponseBody) GoString() string {
	return s.String()
}

func (s *DeploySDGResponseBody) GetData() *DeploySDGResponseBodyData {
	return s.Data
}

func (s *DeploySDGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeploySDGResponseBody) SetData(v *DeploySDGResponseBodyData) *DeploySDGResponseBody {
	s.Data = v
	return s
}

func (s *DeploySDGResponseBody) SetRequestId(v string) *DeploySDGResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeploySDGResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeploySDGResponseBodyData struct {
	// The response message. Success is returned for a successful request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The task result.
	Result *DeploySDGResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s DeploySDGResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeploySDGResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeploySDGResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DeploySDGResponseBodyData) GetResult() *DeploySDGResponseBodyDataResult {
	return s.Result
}

func (s *DeploySDGResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DeploySDGResponseBodyData) SetMessage(v string) *DeploySDGResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeploySDGResponseBodyData) SetResult(v *DeploySDGResponseBodyDataResult) *DeploySDGResponseBodyData {
	s.Result = v
	return s
}

func (s *DeploySDGResponseBodyData) SetSuccess(v bool) *DeploySDGResponseBodyData {
	s.Success = &v
	return s
}

func (s *DeploySDGResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DeploySDGResponseBodyDataResult struct {
	// The number of failed tasks.
	//
	// example:
	//
	// 0
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// Details of failed tasks.
	FailedItems []*DeploySDGResponseBodyDataResultFailedItems `json:"FailedItems,omitempty" xml:"FailedItems,omitempty" type:"Repeated"`
	// The number of successful tasks.
	//
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s DeploySDGResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s DeploySDGResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DeploySDGResponseBodyDataResult) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *DeploySDGResponseBodyDataResult) GetFailedItems() []*DeploySDGResponseBodyDataResultFailedItems {
	return s.FailedItems
}

func (s *DeploySDGResponseBodyDataResult) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *DeploySDGResponseBodyDataResult) SetFailedCount(v int64) *DeploySDGResponseBodyDataResult {
	s.FailedCount = &v
	return s
}

func (s *DeploySDGResponseBodyDataResult) SetFailedItems(v []*DeploySDGResponseBodyDataResultFailedItems) *DeploySDGResponseBodyDataResult {
	s.FailedItems = v
	return s
}

func (s *DeploySDGResponseBodyDataResult) SetSuccessCount(v int64) *DeploySDGResponseBodyDataResult {
	s.SuccessCount = &v
	return s
}

func (s *DeploySDGResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}

type DeploySDGResponseBodyDataResultFailedItems struct {
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

func (s DeploySDGResponseBodyDataResultFailedItems) String() string {
	return dara.Prettify(s)
}

func (s DeploySDGResponseBodyDataResultFailedItems) GoString() string {
	return s.String()
}

func (s *DeploySDGResponseBodyDataResultFailedItems) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeploySDGResponseBodyDataResultFailedItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeploySDGResponseBodyDataResultFailedItems) SetErrMessage(v string) *DeploySDGResponseBodyDataResultFailedItems {
	s.ErrMessage = &v
	return s
}

func (s *DeploySDGResponseBodyDataResultFailedItems) SetInstanceId(v string) *DeploySDGResponseBodyDataResultFailedItems {
	s.InstanceId = &v
	return s
}

func (s *DeploySDGResponseBodyDataResultFailedItems) Validate() error {
	return dara.Validate(s)
}
